// Package debuglog provides a RoundTripper you can put into
// an HTTP client Transport to log requests made with that client.
// This has been proven useful for finding starr app API payloads,
// and as a general debug log wrapper for an integrating application.
package debuglog

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

// Config is the input data for the logger.
type Config struct {
	// Limit logged JSON payloads to this many bytes. 0=unlimited
	MaxBody int
	// This is where logs go. If not set they go to log.Printf.
	Debugf func(string, ...interface{})
	// This can be used for byte counters, but is optional otherwise.
	Caller Caller
	// Any strings in this list are replaced with <recated> in the log output.
	// Useful for hiding api keys and passwords from debug logs. String must be 4+ chars.
	Redact []string
}

const minRedactChars = 4

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
	Elapsed time.Duration
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

	start := time.Now()

	resp, err := rt.next.RoundTrip(req)
	if err != nil {
		if rt.config.Caller != nil {
			// Send this report now since .Close() will never be called.
			rt.config.Caller("000 Failed", req.Method, buf.Len(), 0, err)
		}

		return resp, err //nolint:wrapcheck
	}

	resp.Body = rt.newFakeCloser(resp, &buf, start)

	return resp, nil
}

func (rt *LoggingRoundTripper) newFakeCloser(resp *http.Response, sent *bytes.Buffer, start time.Time) io.ReadCloser {
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
		Elapsed: time.Since(start),
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
		sentBytes = f.Sent.Len()
		rcvdBytes = f.Body.Len()
		sent      = f.Sent.String()
		rcvd      = f.Body.String()
	)

	if f.MaxBody > 0 && len(sent) > f.MaxBody {
		sent = sent[:f.MaxBody] + " <data truncated>"
	}

	switch ctype := f.Header.Get("content-type"); {
	case !strings.Contains(ctype, "json"):
		// We only log JSON. Need something else? Ask!
		rcvd = "<data not logged, content-type: " + ctype + ">"
	case f.MaxBody > 0 && len(rcvd) > f.MaxBody:
		rcvd = rcvd[:f.MaxBody] + " <body truncated>"
	}

	if sentBytes > 0 {
		f.redactLog("Sent (%s) %d bytes to %s in %s: %s\n Response: %s %d bytes\n%s%s)",
			f.Method, sentBytes, f.URL, f.Elapsed.Round(time.Millisecond),
			sent, f.Status, rcvdBytes, f.headers(), rcvd)
	} else {
		f.redactLog("Sent (%s) to %s in %s, Response: %s %d bytes\n%s%s",
			f.Method, f.URL, f.Elapsed.Round(time.Millisecond),
			f.Status, rcvdBytes, f.headers(), rcvd)
	}

	return sentBytes, rcvdBytes
}

func (f *fakeCloser) headers() string {
	var headers string

	for header, values := range f.Header {
		for _, value := range values {
			headers += header + ": " + value + "\n"
		}
	}

	return headers
}

func (f *fakeCloser) redactLog(msg string, format ...interface{}) {
	msg = fmt.Sprintf(msg, format...)

	for _, redact := range f.Redact {
		if len(redact) >= minRedactChars {
			msg = strings.ReplaceAll(msg, redact, "<redacted>")
		}
	}

	f.Debugf(msg)
}
