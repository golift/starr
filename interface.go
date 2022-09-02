package starr

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
// It changes once in a while, so avoid making hard dependencies on it.
type APIer interface {
	Login(ctx context.Context) error
	// Normal data, returns response. Do not use these in starr app methods.
	Get(ctx context.Context, path string, params url.Values) (*http.Response, error)
	Post(ctx context.Context, path string, params url.Values, postBody io.Reader) (*http.Response, error)
	Put(ctx context.Context, path string, params url.Values, putBody io.Reader) (*http.Response, error)
	Delete(ctx context.Context, path string, params url.Values) (*http.Response, error)
	// Normal data, unmarshals into provided interface. Use these because they close the response body.
	GetInto(ctx context.Context, path string, params url.Values, output interface{}) error
	PostInto(ctx context.Context, path string, params url.Values, postBody io.Reader, output interface{}) error
	PutInto(ctx context.Context, path string, params url.Values, putBody io.Reader, output interface{}) error
	DeleteInto(ctx context.Context, path string, params url.Values, output interface{}) error
}

// Config must satify the APIer struct.
var _ APIer = (*Config)(nil)

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

	resp, err := c.req(ctx, "/login", http.MethodPost, nil, bytes.NewBufferString(post))
	if err != nil {
		return fmt.Errorf("authenticating as user '%s' failed: %w", c.Username, err)
	}
	defer resp.Body.Close()

	_, _ = io.Copy(io.Discard, resp.Body)

	if u, _ := url.Parse(c.URL); strings.Contains(resp.Header.Get("location"), "loginFailed") ||
		len(c.Client.Jar.Cookies(u)) == 0 {
		return fmt.Errorf("%w: authenticating as user '%s' failed", ErrRequestError, c.Username)
	}

	c.cookie = true

	return nil
}

// Get makes a GET http request and returns the body.
func (c *Config) Get(ctx context.Context, path string, params url.Values) (*http.Response, error) {
	resp, err := c.req(ctx, path, http.MethodGet, params, nil)
	return resp, err
}

// Post makes a POST http request and returns the body.
func (c *Config) Post(ctx context.Context, path string, params url.Values, postBody io.Reader) (*http.Response, error) {
	resp, err := c.req(ctx, path, http.MethodPost, params, postBody)
	return resp, err
}

// Put makes a PUT http request and returns the body.
func (c *Config) Put(ctx context.Context, path string, params url.Values, putBody io.Reader) (*http.Response, error) {
	resp, err := c.req(ctx, path, http.MethodPut, params, putBody)
	return resp, err
}

// Delete makes a DELETE http request and returns the body.
func (c *Config) Delete(ctx context.Context, path string, params url.Values) (*http.Response, error) {
	resp, err := c.req(ctx, path, http.MethodDelete, params, nil)
	return resp, err
}

// GetInto performs an HTTP GET against an API path and
// unmarshals the payload into the provided pointer interface.
func (c *Config) GetInto(ctx context.Context, path string, params url.Values, output interface{}) error {
	resp, err := c.req(ctx, path, http.MethodGet, params, nil) //nolint:bodyclose
	return unmarshal(output, resp.Body, err)
}

// PostInto performs an HTTP POST against an API path and
// unmarshals the payload into the provided pointer interface.
func (c *Config) PostInto(
	ctx context.Context,
	path string,
	params url.Values,
	postBody io.Reader,
	output interface{},
) error {
	resp, err := c.req(ctx, path, http.MethodPost, params, postBody) //nolint:bodyclose
	return unmarshal(output, resp.Body, err)
}

// PutInto performs an HTTP PUT against an API path and
// unmarshals the payload into the provided pointer interface.
func (c *Config) PutInto(
	ctx context.Context,
	path string,
	params url.Values,
	putBody io.Reader,
	output interface{},
) error {
	resp, err := c.req(ctx, path, http.MethodPut, params, putBody) //nolint:bodyclose
	return unmarshal(output, resp.Body, err)
}

// DeleteInto performs an HTTP DELETE against an API path
// and unmarshals the payload into a pointer interface.
func (c *Config) DeleteInto(ctx context.Context, path string, params url.Values, output interface{}) error {
	resp, err := c.req(ctx, path, http.MethodDelete, params, nil) //nolint:bodyclose
	return unmarshal(output, resp.Body, err)
}

// unmarshal is an extra procedure to check an error and unmarshal the resp.Body payload.
func unmarshal(output interface{}, data io.ReadCloser, err error) error {
	defer data.Close()

	if err != nil {
		return err
	} else if output == nil {
		return fmt.Errorf("this is a code bug: %w", ErrNilInterface)
	} else if err = json.NewDecoder(data).Decode(output); err != nil {
		return fmt.Errorf("json parse error: %w", err)
	}

	return nil
}
