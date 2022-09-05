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
//
//nolint:lll
type APIer interface {
	Login(ctx context.Context) error
	// Normal data, returns response. Do not use these in starr app methods.
	// These methods are generally for non-api paths and will not ensure an /api uri prefix.
	Get(ctx context.Context, path string, params *Params) (*http.Response, error)    // Get request; Params are optional.
	Post(ctx context.Context, path string, params *Params) (*http.Response, error)   // Post request; Params should contain io.Reader.
	Put(ctx context.Context, path string, params *Params) (*http.Response, error)    // Put request; Params should contain io.Reader.
	Delete(ctx context.Context, path string, params *Params) (*http.Response, error) // Delete request; Params are optional.
	// Normal data, unmarshals into provided interface. Use these because they close the response body.
	GetInto(ctx context.Context, path string, params url.Values, output interface{}) error
	PostInto(ctx context.Context, path string, params url.Values, postBody io.Reader, output interface{}) error
	PutInto(ctx context.Context, path string, params url.Values, putBody io.Reader, output interface{}) error
	DeleteAny(ctx context.Context, path string, params *Params) error // Delete request; Params are optional.
}

// Config must satify the APIer struct.
var _ APIer = (*Config)(nil)

// Login POSTs to the login form in a Starr app and saves the authentication cookie for future use.
func (c *Config) Login(ctx context.Context) error {
	if c.Client.Jar == nil {
		jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
		if err != nil {
			return fmt.Errorf("cookiejar.New(publicsuffix): %w", err)
		}

		c.Client.Jar = jar
	}

	post := "username=" + c.Username + "&password=" + c.Password

	resp, err := c.api(ctx, "/login", http.MethodPost, &Params{nil, bytes.NewBufferString(post)})
	if err != nil {
		return fmt.Errorf("authenticating as user '%s' failed: %w", c.Username, err)
	}

	closeResp(resp)

	if u, _ := url.Parse(c.URL); strings.Contains(resp.Header.Get("location"), "loginFailed") ||
		len(c.Client.Jar.Cookies(u)) == 0 {
		return fmt.Errorf("%w: authenticating as user '%s' failed", ErrRequestError, c.Username)
	}

	c.cookie = true

	return nil
}

// Get makes a GET http request and returns the body.
func (c *Config) Get(ctx context.Context, path string, params *Params) (*http.Response, error) {
	return c.Req(ctx, path, http.MethodGet, params)
}

// Post makes a POST http request and returns the body.
func (c *Config) Post(ctx context.Context, path string, params *Params) (*http.Response, error) {
	return c.Req(ctx, path, http.MethodPost, params)
}

// Put makes a PUT http request and returns the body.
func (c *Config) Put(ctx context.Context, path string, params *Params) (*http.Response, error) {
	return c.Req(ctx, path, http.MethodPut, params)
}

// Delete makes a DELETE http request and returns the body.
func (c *Config) Delete(ctx context.Context, path string, params *Params) (*http.Response, error) {
	return c.Req(ctx, path, http.MethodDelete, params)
}

// GetInto performs an HTTP GET against an API path and
// unmarshals the payload into the provided pointer interface.
func (c *Config) GetInto(ctx context.Context, path string, params url.Values, output interface{}) error {
	resp, err := c.api(ctx, path, http.MethodGet, &Params{params, nil})
	return decode(output, resp, err)
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
	resp, err := c.api(ctx, path, http.MethodPost, &Params{params, postBody})
	return decode(output, resp, err)
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
	resp, err := c.api(ctx, path, http.MethodPut, &Params{params, putBody})
	return decode(output, resp, err)
}

// DeleteAny performs an HTTP DELETE against an API path, output is ignored.
func (c *Config) DeleteAny(ctx context.Context, path string, params *Params) error {
	resp, err := c.api(ctx, path, http.MethodDelete, params)
	closeResp(resp)

	return err
}

// decode is an extra procedure to check an error and decode the JSON resp.Body payload.
func decode(output interface{}, resp *http.Response, err error) error {
	if err != nil {
		return err
	} else if output == nil {
		closeResp(resp) // read the body and close it.
		return fmt.Errorf("this is a code bug: %w", ErrNilInterface)
	}

	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(output); err != nil {
		return fmt.Errorf("decoding json response body: %w", err)
	}

	return nil
}
