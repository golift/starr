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
)

// APIer is used by the sub packages to allow mocking the http methods in tests.
// This also allows consuming packages to override methods.
type APIer interface {
	Login() error // Only needed for non-API paths, like backup downloads. Requires Username and Password being set.
	Get(path string, params url.Values) (respBody []byte, err error)
	Post(path string, params url.Values, postBody []byte) (respBody []byte, err error)
	Put(path string, params url.Values, putBody []byte) (respBody []byte, err error)
	Delete(path string, params url.Values) (respBody []byte, err error)
	GetInto(path string, params url.Values, v interface{}) error
	PostInto(path string, params url.Values, postBody []byte, v interface{}) error
	PutInto(path string, params url.Values, putBody []byte, v interface{}) error
	DeleteInto(path string, params url.Values, v interface{}) error
	GetBody(ctx context.Context, path string, params url.Values) (respBody io.ReadCloser, status int, err error)
	PostBody(ctx context.Context, path string, params url.Values,
		postBody []byte) (respBody io.ReadCloser, status int, err error)
	PutBody(ctx context.Context, path string, params url.Values,
		putBody []byte) (respBody io.ReadCloser, status int, err error)
	DeleteBody(ctx context.Context, path string, params url.Values) (respBody io.ReadCloser, status int, err error)
}

// Config must satify the APIer struct.
var _ APIer = (*Config)(nil)

func (c *Config) log(code int, data, body []byte, header http.Header, path, method string, err error) {
	headers := ""

	for header, value := range header {
		for _, v := range value {
			headers += header + ": " + v + "\n"
		}
	}

	bodyStr := string(body)
	if c.MaxBody > 0 && len(bodyStr) > c.MaxBody {
		bodyStr = bodyStr[:c.MaxBody] + " <body truncated>"
	}

	dataStr := string(data)
	if c.MaxBody > 0 && len(dataStr) > c.MaxBody {
		dataStr = dataStr[:c.MaxBody] + " <data truncated>"
	}

	if len(body) > 0 {
		c.Debugf("Sent (%s) %d bytes to %s: %s\n Response: %s\n%s%s (err: %v)",
			method, len(body), path, bodyStr, http.StatusText(code), headers, dataStr, err)
	} else {
		c.Debugf("Sent (%s) to %s, Response: %s\n%s%s (err: %v)",
			method, path, http.StatusText(code), headers, dataStr, err)
	}
}

// Login POSTs to the login form in a Starr app and saves the authentication cookie for future use.
func (c *Config) Login() error {
	if c.Client.Jar == nil {
		jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
		if err != nil {
			return fmt.Errorf("cookiejar.New(publicsuffix): %w", err)
		}

		c.Client.Jar = jar
	}

	post := []byte("username=" + c.Username + "&password=" + c.Password)

	code, resp, header, err := c.body(context.Background(), "/login", http.MethodPost, nil, bytes.NewBuffer(post))
	c.log(code, nil, post, header, c.URL+"/login", http.MethodPost, err)

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
func (c *Config) Get(path string, params url.Values) ([]byte, error) {
	code, data, header, err := c.req(path, http.MethodGet, params, nil)
	c.log(code, data, nil, header, c.setPathParams(path, params), http.MethodGet, err)

	return data, err
}

// Post makes a POST http request and returns the body.
func (c *Config) Post(path string, params url.Values, postBody []byte) ([]byte, error) {
	code, data, header, err := c.req(path, http.MethodPost, params, bytes.NewBuffer(postBody))
	c.log(code, data, postBody, header, c.setPathParams(path, params), http.MethodPost, err)

	return data, err
}

// Put makes a PUT http request and returns the body.
func (c *Config) Put(path string, params url.Values, putBody []byte) ([]byte, error) {
	code, data, header, err := c.req(path, http.MethodPut, params, bytes.NewBuffer(putBody))
	c.log(code, data, putBody, header, c.setPathParams(path, params), http.MethodPut, err)

	return data, err
}

// Delete makes a DELETE http request and returns the body.
func (c *Config) Delete(path string, params url.Values) ([]byte, error) {
	code, data, header, err := c.req(path, http.MethodDelete, params, nil)
	c.log(code, data, nil, header, c.setPathParams(path, params), http.MethodDelete, err)

	return data, err
}

// GetInto performs an HTTP GET against an API path and unmarshals the payload into the provided pointer interface.
func (c *Config) GetInto(path string, params url.Values, v interface{}) error {
	data, err := c.Get(path, params)

	return unmarshal(v, data, err)
}

// PostInto performs an HTTP POST against an API path and unmarshals the payload into the provided pointer interface.
func (c *Config) PostInto(path string, params url.Values, postBody []byte, v interface{}) error {
	data, err := c.Post(path, params, postBody)

	return unmarshal(v, data, err)
}

// PutInto performs an HTTP PUT against an API path and unmarshals the payload into the provided pointer interface.
func (c *Config) PutInto(path string, params url.Values, putBody []byte, v interface{}) error {
	data, err := c.Put(path, params, putBody)

	return unmarshal(v, data, err)
}

// DeleteInto performs an HTTP DELETE against an API path and unmarshals the payload into a pointer interface.
func (c *Config) DeleteInto(path string, params url.Values, v interface{}) error {
	data, err := c.Delete(path, params)

	return unmarshal(v, data, err)
}

// GetBody makes an http request and returns the resp.Body (io.ReadCloser).
// This is useful for downloading things like backup files, but can also be used to get
// around limitations in this library. Always remember to close the io.ReadCloser.
// Before you use the returned data, check the HTTP status code.
// If it's not 200, it's possible the request had an error or was not authenticated.
func (c *Config) GetBody(ctx context.Context, path string, params url.Values) (io.ReadCloser, int, error) {
	code, data, header, err := c.body(ctx, path, http.MethodGet, params, nil)
	c.log(code, nil, nil, header, c.URL+path, http.MethodGet, err)

	return data, code, err
}

// PostBody makes a POST http request and returns the resp.Body (io.ReadCloser).
// Always remember to close the io.ReadCloser.
// Before you use the returned data, check the HTTP status code.
// If it's not 200, it's possible the request had an error or was not authenticated.
func (c *Config) PostBody(ctx context.Context, path string, params url.Values,
	postBody []byte) (io.ReadCloser, int, error) {
	code, data, header, err := c.body(ctx, path, http.MethodPost, params, bytes.NewBuffer(postBody))
	c.log(code, nil, postBody, header, c.URL+path, http.MethodPost, err)

	return data, code, err
}

// PutBody makes a PUT http request and returns the resp.Body (io.ReadCloser).
// Always remember to close the io.ReadCloser.
// Before you use the returned data, check the HTTP status code.
func (c *Config) PutBody(ctx context.Context, path string, params url.Values,
	putBody []byte) (io.ReadCloser, int, error) {
	code, data, header, err := c.body(ctx, path, http.MethodPut, params, bytes.NewBuffer(putBody))
	c.log(code, nil, putBody, header, c.URL+path, http.MethodPut, err)

	return data, code, err
}

// DeleteBody makes a DELETE http request and returns the resp.Body (io.ReadCloser).
// Always remember to close the io.ReadCloser.
// Before you use the returned data, check the HTTP status code.
// If it's not 200, it's possible the request had an error or was not authenticated.
func (c *Config) DeleteBody(ctx context.Context, path string, params url.Values) (io.ReadCloser, int, error) {
	code, data, header, err := c.body(ctx, path, http.MethodDelete, params, nil)
	c.log(code, nil, nil, header, c.URL+path, http.MethodDelete, err)

	return data, code, err
}

// unmarshal is an extra procedure to check an error and unmarshal the payload.
// This allows the methods above to have all their logic abstracted.
func unmarshal(v interface{}, data []byte, err error) error {
	if err != nil {
		return err
	} else if v == nil {
		return fmt.Errorf("this is a code bug: %w", ErrNilInterface)
	} else if err = json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("json parse error: %w", err)
	}

	return nil
}
