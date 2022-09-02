// Package debuglog provides a RoundTripper you can put into
// an HTTP client Transport to log requests made with that client.
package debuglog

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

// LogConfig is the input data for the logger.
type LogConfig struct {
	MaxBody int                            // Limit payloads to this many bytes. 0=unlimited
	Debugf  func(string, ...interface{})   // This is where logs go.
	Caller  func(sentBytes, rcvdBytes int) // This can be used for byte counters.
}

// LoggingRoundTripper allows us to use a datacounter to log http request data.
type LoggingRoundTripper struct {
	next   http.RoundTripper // The next Transport to call after logging.
	config *LogConfig
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
	*LogConfig
}

// NewLoggingRoundTripper returns a round tripper to log requests counts and response sizes.
func NewLoggingRoundTripper(config LogConfig, next http.RoundTripper) *LoggingRoundTripper {
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
	defer req.Body.Close()

	buf := bytes.Buffer{}
	if req.Body != nil {
		sent := io.TeeReader(req.Body, &buf)
		req.Body = io.NopCloser(sent)
	}

	resp, err := rt.next.RoundTrip(req)
	if resp == nil || resp.Body == nil {
		return resp, err //nolint:wrapcheck
	}

	resp.Body = rt.newFakeCloser(resp, &buf)

	return resp, err //nolint:wrapcheck
}

func (rt *LoggingRoundTripper) newFakeCloser(resp *http.Response, sent *bytes.Buffer) io.ReadCloser {
	var buf bytes.Buffer

	return &fakeCloser{
		CloseFn:   resp.Body.Close,
		Reader:    io.TeeReader(resp.Body, &buf),
		Body:      &buf,
		Method:    resp.Request.Method,
		Status:    resp.Status,
		URL:       resp.Request.URL.String(),
		Sent:      sent,
		Header:    resp.Header,
		LogConfig: rt.config,
	}
}

func (f *fakeCloser) Close() error {
	sentBytes, rcvdBytes := f.logRequest()
	if f.Caller != nil {
		f.Caller(sentBytes, rcvdBytes)
	}

	return f.CloseFn()
}

func (f *fakeCloser) logRequest() (int, int) {
	var (
		headers   = ""
		rcvd      = ""
		sent      = ""
		rcvdBytes = f.Body.Len()
		sentBytes = f.Sent.Len()
	)

	for header, value := range f.Header {
		for _, v := range value {
			headers += header + ": " + v + "\n"
		}
	}

	if f.MaxBody > 0 && rcvdBytes > f.MaxBody {
		rcvd = string(f.Body.Bytes()[:f.MaxBody]) + " <body truncated>"
	}

	if f.MaxBody > 0 && sentBytes > f.MaxBody {
		sent = string(f.Sent.Bytes()[:f.MaxBody]) + " <data truncated>"
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
