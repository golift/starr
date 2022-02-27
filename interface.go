package starr

// The following go:generate directive re-creates the mocks for this API when you run go generate.
//go:generate mockgen -package mocks -destination=mocks/apier.go golift.io/starr APIer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"golang.org/x/net/publicsuffix"
	"golift.io/datacounter"
)

// APIer is used by the sub packages to allow mocking the http methods in tests.
// It changes once in a while, so avoid making hard dependencies on it.
type APIer interface {
	Login(ctx context.Context) error
	// Normal data, returns response body.
	Get(ctx context.Context, path string, params url.Values) (respBody []byte, err error)
	Post(ctx context.Context, path string, params url.Values, postBody io.Reader) (respBody []byte, err error)
	Put(ctx context.Context, path string, params url.Values, putBody io.Reader) (respBody []byte, err error)
	Delete(ctx context.Context, path string, params url.Values) (respBody []byte, err error)
	// Normal data, unmarshals into provided interface.
	GetInto(ctx context.Context, path string, params url.Values, output interface{}) (int64, error)
	PostInto(ctx context.Context, path string, params url.Values, postBody io.Reader, output interface{}) (int64, error)
	PutInto(ctx context.Context, path string, params url.Values, putBody io.Reader, output interface{}) (int64, error)
	DeleteInto(ctx context.Context, path string, params url.Values, output interface{}) (int64, error)
	// Body methods.
	GetBody(ctx context.Context, path string, params url.Values) (respBody io.ReadCloser, status int, err error)
	PostBody(ctx context.Context, path string, params url.Values,
		postBody io.Reader) (respBody io.ReadCloser, status int, err error)
	PutBody(ctx context.Context, path string, params url.Values,
		putBody io.Reader) (respBody io.ReadCloser, status int, err error)
	DeleteBody(ctx context.Context, path string, params url.Values) (respBody io.ReadCloser, status int, err error)
}

// Config must satify the APIer struct.
var _ APIer = (*Config)(nil)

// log the request. Do not call this if c.Debugf is nil.
func (c *Config) log(code int, data, body string, header http.Header, path, method string, err error) {
	headers := ""

	for header, value := range header {
		for _, v := range value {
			headers += header + ": " + v + "\n"
		}
	}

	if c.MaxBody > 0 && len(body) > c.MaxBody {
		body = body[:c.MaxBody] + " <body truncated>"
	}

	if c.MaxBody > 0 && len(data) > c.MaxBody {
		data = data[:c.MaxBody] + " <data truncated>"
	}

	if len(body) > 0 {
		c.Debugf("Sent (%s) %d bytes to %s: %s\n Response: %s\n%s%s (err: %v)",
			method, len(body), path, body, http.StatusText(code), headers, data, err)
	} else {
		c.Debugf("Sent (%s) to %s, Response: %s\n%s%s (err: %v)",
			method, path, http.StatusText(code), headers, data, err)
	}
}

// LoginC POSTs to the login form in a Starr app and saves the authentication cookie for future use.
func (c *Config) Login(ctx context.Context) error {
	if c.Client.Jar == nil {
		jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
		if err != nil {
			return fmt.Errorf("cookiejar.New(publicsuffix): %w", err)
		}

		c.Client.Jar = jar
	}

	post := "username=" + c.Username + "&password=" + c.Password

	code, resp, header, err := c.body(ctx, "/login", http.MethodPost, nil, bytes.NewBufferString(post))
	c.log(code, "", post, header, c.URL+"/login", http.MethodPost, err)

	if err != nil {
		return fmt.Errorf("authenticating as user '%s' failed: %w", c.Username, err)
	}
	defer resp.Close()

	_, _ = io.Copy(io.Discard, resp)

	if u, _ := url.Parse(c.URL); strings.Contains(header.Get("location"), "loginFailed") ||
		len(c.Client.Jar.Cookies(u)) == 0 {
		return fmt.Errorf("%w: authenticating as user '%s' failed", ErrRequestError, c.Username)
	}

	c.cookie = true

	return nil
}

// Get makes a GET http request and returns the body.
func (c *Config) Get(ctx context.Context, path string, params url.Values) ([]byte, error) {
	code, data, header, err := c.Req(ctx, path, http.MethodGet, params, nil)

	if c.Debugf != nil { // log the request.
		c.log(code, string(data), "", header, c.SetPathParams(path, params), http.MethodGet, err)
	}

	return data, err
}

// Post makes a POST http request and returns the body.
func (c *Config) Post(ctx context.Context, path string, params url.Values, postBody io.Reader) ([]byte, error) {
	if c.Debugf == nil { // no log, pass it through.
		_, data, _, err := c.Req(ctx, path, http.MethodPost, params, postBody)

		return data, err
	}

	var buf bytes.Buffer                // split the reader for request and log.
	tee := io.TeeReader(postBody, &buf) // must read tee first.
	code, data, header, err := c.Req(ctx, path, http.MethodPost, params, tee)
	c.log(code, string(data), buf.String(), header, c.SetPathParams(path, params), http.MethodPost, err)

	return data, err
}

// Put makes a PUT http request and returns the body.
func (c *Config) Put(ctx context.Context, path string, params url.Values, putBody io.Reader) ([]byte, error) {
	if c.Debugf == nil { // no log, pass it through.
		_, data, _, err := c.Req(ctx, path, http.MethodPut, params, putBody)

		return data, err
	}

	var buf bytes.Buffer               // split the reader for request and log.
	tee := io.TeeReader(putBody, &buf) // must read tee first.
	code, data, header, err := c.Req(ctx, path, http.MethodPut, params, tee)
	c.log(code, string(data), buf.String(), header, c.SetPathParams(path, params), http.MethodPut, err)

	return data, err
}

// Delete makes a DELETE http request and returns the body.
func (c *Config) Delete(ctx context.Context, path string, params url.Values) ([]byte, error) {
	code, data, header, err := c.Req(ctx, path, http.MethodDelete, params, nil)

	if c.Debugf != nil {
		c.log(code, string(data), "", header, c.SetPathParams(path, params), http.MethodDelete, err)
	}

	return data, err
}

// GetInto performs an HTTP GET against an API path and
// unmarshals the payload into the provided pointer interface.
func (c *Config) GetInto(ctx context.Context, path string, params url.Values, output interface{}) (int64, error) {
	if c.Debugf == nil { // no log, pass it through.
		_, data, _, err := c.body(ctx, path, http.MethodGet, params, nil)

		return unmarshalBody(output, data, err)
	}

	data, err := c.Get(ctx, path, params) // log the request

	return unmarshal(output, data, err)
}

// PostInto performs an HTTP POST against an API path and
// unmarshals the payload into the provided pointer interface.
func (c *Config) PostInto(ctx context.Context,
	path string, params url.Values, postBody io.Reader, output interface{}) (int64, error) {
	if c.Debugf == nil { // no log, pass it through.
		_, data, _, err := c.body(ctx, path, http.MethodPost, params, postBody)

		return unmarshalBody(output, data, err)
	}

	data, err := c.Post(ctx, path, params, postBody) // log the request

	return unmarshal(output, data, err)
}

// PutInto performs an HTTP PUT against an API path and
// unmarshals the payload into the provided pointer interface.
func (c *Config) PutInto(ctx context.Context,
	path string, params url.Values, putBody io.Reader, output interface{}) (int64, error) {
	if c.Debugf == nil { // no log, pass it through.
		_, data, _, err := c.body(ctx, path, http.MethodPut, params, putBody)

		return unmarshalBody(output, data, err)
	}

	data, err := c.Put(ctx, path, params, putBody) // log the request

	return unmarshal(output, data, err)
}

// DeleteInto performs an HTTP DELETE against an API path
// and unmarshals the payload into a pointer interface.
func (c *Config) DeleteInto(ctx context.Context, path string, params url.Values, output interface{}) (int64, error) {
	if c.Debugf == nil { // no log, pass it through.
		_, data, _, err := c.body(ctx, path, http.MethodDelete, params, nil)

		return unmarshalBody(output, data, err)
	}

	data, err := c.Delete(ctx, path, params) // log the request

	return unmarshal(output, data, err)
}

// GetBody makes an http request and returns the resp.Body (io.ReadCloser).
// This is useful for downloading things like backup files, but can also be used to get
// around limitations in this library. Always remember to close the io.ReadCloser.
// Before you use the returned data, check the HTTP status code.
// If it's not 200, it's possible the request had an error or was not authenticated.
func (c *Config) GetBody(ctx context.Context, path string, params url.Values) (io.ReadCloser, int, error) {
	code, data, header, err := c.body(ctx, path, http.MethodGet, params, nil)

	if c.Debugf != nil {
		c.log(code, "", "", header, c.URL+path, http.MethodGet, err)
	}

	return data, code, err
}

// PostBody makes a POST http request and returns the resp.Body (io.ReadCloser).
// Always remember to close the io.ReadCloser.
// Before you use the returned data, check the HTTP status code.
// If it's not 200, it's possible the request had an error or was not authenticated.
func (c *Config) PostBody(ctx context.Context,
	path string, params url.Values, postBody io.Reader) (io.ReadCloser, int, error) {
	code, data, header, err := c.body(ctx, path, http.MethodPost, params, postBody)

	if c.Debugf != nil {
		c.log(code, "", "", header, c.URL+path, http.MethodPost, err)
	}

	return data, code, err
}

// PutBody makes a PUT http request and returns the resp.Body (io.ReadCloser).
// Always remember to close the io.ReadCloser.
// Before you use the returned data, check the HTTP status code.
func (c *Config) PutBody(ctx context.Context,
	path string, params url.Values, putBody io.Reader) (io.ReadCloser, int, error) {
	code, data, header, err := c.body(ctx, path, http.MethodPut, params, putBody)

	if c.Debugf != nil {
		c.log(code, "", "", header, c.URL+path, http.MethodPut, err)
	}

	return data, code, err
}

// DeleteBody makes a DELETE http request and returns the resp.Body (io.ReadCloser).
// Always remember to close the io.ReadCloser.
// Before you use the returned data, check the HTTP status code.
// If it's not 200, it's possible the request had an error or was not authenticated.
func (c *Config) DeleteBody(ctx context.Context, path string, params url.Values) (io.ReadCloser, int, error) {
	code, data, header, err := c.body(ctx, path, http.MethodDelete, params, nil)

	if c.Debugf != nil {
		c.log(code, "", "", header, c.URL+path, http.MethodDelete, err)
	}

	return data, code, err
}

// unmarshal is an extra procedure to check an error and unmarshal the payload.
// This allows the methods above to have all their logic abstracted.
func unmarshal(v interface{}, data []byte, err error) (int64, error) {
	if err != nil {
		return int64(len(data)), err
	} else if v == nil {
		return int64(len(data)), fmt.Errorf("this is a code bug: %w", ErrNilInterface)
	} else if err = json.Unmarshal(data, v); err != nil {
		return int64(len(data)), fmt.Errorf("json parse error: %w", err)
	}

	return int64(len(data)), nil
}

// unmarshalBody is an extra procedure to check an error and unmarshal the resp.Body payload.
// This version unmarshals the resp.Body directly.
func unmarshalBody(output interface{}, data io.ReadCloser, err error) (int64, error) {
	defer data.Close()

	counter := datacounter.NewReaderCounter(data)

	if err != nil {
		return int64(counter.Count()), err
	} else if output == nil {
		return int64(counter.Count()), fmt.Errorf("this is a code bug: %w", ErrNilInterface)
	} else if err = json.NewDecoder(counter).Decode(output); err != nil {
		return int64(counter.Count()), fmt.Errorf("json parse error: %w", err)
	}

	return int64(counter.Count()), nil
}
