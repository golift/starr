// Package debuglog provides a RoundTripper you can put into
// an HTTP client Transport to log requests made with that client.
package debuglog

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"strings"
)

// Config is the input data for the logger.
type Config struct {
	MaxBody int                          // Limit payloads to this many bytes. 0=unlimited
	Debugf  func(string, ...interface{}) // This is where logs go.
	Caller  Caller                       // This can be used for byte counters.
}

// Caller is a callback function you may use to collect statistics.
type Caller func(status, method string, sentBytes, rcvdBytes int, err error)

// LoggingRoundTripper allows us to use a datacounter to log http request data.
type LoggingRoundTripper struct {
	next   http.RoundTripper // The next Transport to call after logging.
	config *Config
}

type fakeCloser struct {
	io.Reader
	CloseFn func() error
	Body    *bytes.Buffer
	Sent    *bytes.Buffer
	Method  string
	URL     string
	Status  string
	Header  http.Header
	*Config
}

// NewLoggingRoundTripper returns a round tripper to log requests counts and response sizes.
func NewLoggingRoundTripper(config Config, next http.RoundTripper) *LoggingRoundTripper {
	if next == nil {
		next = http.DefaultTransport
	}

	if config.Debugf == nil {
		config.Debugf = log.Printf
	}

	return &LoggingRoundTripper{
		next:   next,
		config: &config,
	}
}

// RoundTrip satisfies the http.RoundTripper interface.
func (rt *LoggingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	buf := bytes.Buffer{}
	if req.Body != nil {
		sent := io.TeeReader(req.Body, &buf)
		req.Body = io.NopCloser(sent)

		defer req.Body.Close()
	}

	resp, err := rt.next.RoundTrip(req)
	if err != nil {
		if rt.config.Caller != nil {
			// Send this report now since .Close() will never be called.
			rt.config.Caller("000 Failed", req.Method, buf.Len(), 0, err)
		}

		return resp, err //nolint:wrapcheck
	}

	resp.Body = rt.newFakeCloser(resp, &buf)

	return resp, nil
}

func (rt *LoggingRoundTripper) newFakeCloser(resp *http.Response, sent *bytes.Buffer) io.ReadCloser {
	var buf bytes.Buffer

	return &fakeCloser{
		CloseFn: resp.Body.Close,
		Reader:  io.TeeReader(resp.Body, &buf),
		Body:    &buf,
		Method:  resp.Request.Method,
		Status:  resp.Status,
		URL:     resp.Request.URL.String(),
		Sent:    sent,
		Header:  resp.Header,
		Config:  rt.config,
	}
}

// Close closes the response Body, logs the request, and fires the call back.
func (f *fakeCloser) Close() error {
	sentBytes, rcvdBytes := f.logRequest()
	if f.Caller != nil {
		f.Caller(f.Status, f.Method, sentBytes, rcvdBytes, nil)
	}

	return f.CloseFn()
}

func (f *fakeCloser) logRequest() (int, int) {
	var (
		sent      string
		rcvd      string
		headers   string
		rcvdBytes = f.Body.Len()
		sentBytes = f.Sent.Len()
	)

	if f.MaxBody > 0 && sentBytes > f.MaxBody {
		sent = string(f.Sent.Bytes()[:f.MaxBody]) + " <data truncated>"
	} else {
		sent = f.Sent.String()
	}

	switch ctype := f.Header.Get("content-type"); {
	case !strings.Contains(ctype, "json"):
		rcvd = "<data not logged, content-type: " + ctype + ">"
	case f.MaxBody > 0 && rcvdBytes > f.MaxBody:
		rcvd = string(f.Body.Bytes()[:f.MaxBody]) + " <body truncated>"
	default:
		rcvd = f.Body.String()
	}

	for header, value := range f.Header {
		for _, v := range value {
			headers += header + ": " + v + "\n"
		}
	}

	if sentBytes > 0 {
		f.Debugf("Sent (%s) %d bytes to %s: %s\n Response: %s %d bytes\n%s%s)",
			f.Method, sentBytes, f.URL, sent, f.Status, rcvdBytes, headers, rcvd)
	} else {
		f.Debugf("Sent (%s) to %s, Response: %s %d bytes\n%s%s",
			f.Method, f.URL, f.Status, rcvdBytes, headers, rcvd)
	}

	return sentBytes, rcvdBytes
}
