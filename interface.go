package starr

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
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
	GetInitializeJS(ctx context.Context) (*InitializeJS, error)
	// Login is used for non-API paths, like downloading backups or the initialize.js file.
	Login(ctx context.Context) error
	// Normal data, returns response. Do not use these in starr app methods.
	// These methods are generally for non-api paths and will not ensure an /api uri prefix.
	Get(ctx context.Context, req Request) (*http.Response, error)    // Get request; Params are optional.
	Post(ctx context.Context, req Request) (*http.Response, error)   // Post request; Params should contain io.Reader.
	Put(ctx context.Context, req Request) (*http.Response, error)    // Put request; Params should contain io.Reader.
	Delete(ctx context.Context, req Request) (*http.Response, error) // Delete request; Params are optional.
	// Normal data, unmarshals into provided interface. Use these because they close the response body.
	GetInto(ctx context.Context, req Request, output interface{}) error  // API GET Request.
	PostInto(ctx context.Context, req Request, output interface{}) error // API POST Request.
	PutInto(ctx context.Context, req Request, output interface{}) error  // API PUT Request.
	DeleteAny(ctx context.Context, req Request) error                    // API Delete request.
}

// Config must satify the APIer struct.
var _ APIer = (*Config)(nil)

// InitializeJS is the data contained in the initialize.js file.
type InitializeJS struct {
	App          string
	APIRoot      string
	APIKey       string
	Release      string
	Version      string
	InstanceName string
	Theme        string
	Branch       string
	Analytics    string
	UserHash     string
	URLBase      string
	IsProduction bool
}

// Login POSTs to the login form in a Starr app and saves the authentication cookie for future use.
func (c *Config) Login(ctx context.Context) error {
	if c.Client.Jar == nil {
		jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
		if err != nil {
			return fmt.Errorf("cookiejar.New(publicsuffix): %w", err)
		}

		c.Client.Jar = jar
	}

	params := make(url.Values)
	params.Add("username", c.Username)
	params.Add("password", c.Password)

	req := Request{URI: "/login", Body: bytes.NewBufferString(params.Encode())}
	codeErr := &ReqError{}

	resp, err := c.req(ctx, http.MethodPost, req)
	if err != nil {
		if !errors.As(err, &codeErr) { // pointer to a pointer, yup.
			return fmt.Errorf("invalid reply authenticating as user '%s': %w", c.Username, err)
		}
	} else {
		// Protect a nil map in case we don't get an error (which should be impossible).
		codeErr.Header = resp.Header
	}

	closeResp(resp)

	if u, _ := url.Parse(c.URL); strings.Contains(codeErr.Get("location"), "loginFailed") ||
		len(c.Client.Jar.Cookies(u)) == 0 {
		return fmt.Errorf("%w: authenticating as user '%s' failed", ErrRequestError, c.Username)
	}

	c.cookie = true

	return nil
}

// Get makes a GET http request and returns the body.
func (c *Config) Get(ctx context.Context, req Request) (*http.Response, error) {
	return c.Req(ctx, http.MethodGet, req)
}

// Post makes a POST http request and returns the body.
func (c *Config) Post(ctx context.Context, req Request) (*http.Response, error) {
	return c.Req(ctx, http.MethodPost, req)
}

// Put makes a PUT http request and returns the body.
func (c *Config) Put(ctx context.Context, req Request) (*http.Response, error) {
	return c.Req(ctx, http.MethodPut, req)
}

// Delete makes a DELETE http request and returns the body.
func (c *Config) Delete(ctx context.Context, req Request) (*http.Response, error) {
	return c.Req(ctx, http.MethodDelete, req)
}

// GetInto performs an HTTP GET against an API path and
// unmarshals the payload into the provided pointer interface.
func (c *Config) GetInto(ctx context.Context, req Request, output interface{}) error {
	resp, err := c.api(ctx, http.MethodGet, req)
	return decode(output, resp, err)
}

// PostInto performs an HTTP POST against an API path and
// unmarshals the payload into the provided pointer interface.
func (c *Config) PostInto(ctx context.Context, req Request, output interface{}) error {
	resp, err := c.api(ctx, http.MethodPost, req)
	return decode(output, resp, err)
}

// PutInto performs an HTTP PUT against an API path and
// unmarshals the payload into the provided pointer interface.
func (c *Config) PutInto(ctx context.Context, req Request, output interface{}) error {
	resp, err := c.api(ctx, http.MethodPut, req)
	return decode(output, resp, err)
}

// DeleteAny performs an HTTP DELETE against an API path, output is ignored.
func (c *Config) DeleteAny(ctx context.Context, req Request) error {
	resp, err := c.api(ctx, http.MethodDelete, req)
	closeResp(resp)

	return err
}

// decode is an extra procedure to check an error and decode the JSON resp.Body payload.
func decode(output interface{}, resp *http.Response, err error) error {
	if err != nil {
		return err
	} else if output == nil {
		closeResp(resp) // read the body and close it.
		return fmt.Errorf("this is a Starr library bug: %w", ErrNilInterface)
	}

	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(output); err != nil {
		return fmt.Errorf("decoding Starr JSON response body: %w", err)
	}

	return nil
}

// GetInitializeJS returns the data from the initialize.js file.
// If the instance requires authentication, you must call Login() before this method.
func (c *Config) GetInitializeJS(ctx context.Context) (*InitializeJS, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.URL+"initialize.js", nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext(initialize.js): %w", err)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("httpClient.Do(req): %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		_, _ = io.Copy(io.Discard, resp.Body)
		return nil, &ReqError{Code: resp.StatusCode}
	}

	return readInitializeJS(resp.Body)
}

func readInitializeJS(input io.Reader) (*InitializeJS, error) { //nolint:cyclop
	output := &InitializeJS{}
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		switch split := strings.Fields(scanner.Text()); {
		case len(split) < 2: //nolint:gomnd
			continue
		case split[0] == "apiRoot:":
			output.APIRoot = strings.Trim(split[1], `"',`)
		case split[0] == "apiKey:":
			output.APIKey = strings.Trim(split[1], `"',`)
		case split[0] == "version:":
			output.Version = strings.Trim(split[1], `"',`)
		case split[0] == "release:":
			output.Release = strings.Trim(split[1], `"',`)
		case split[0] == "instanceName:":
			output.InstanceName = strings.Trim(split[1], `"',`)
		case split[0] == "theme:":
			output.Theme = strings.Trim(split[1], `"',`)
		case split[0] == "branch:":
			output.Branch = strings.Trim(split[1], `"',`)
		case split[0] == "analytics:":
			output.Analytics = strings.Trim(split[1], `"',`)
		case split[0] == "userHash:":
			output.UserHash = strings.Trim(split[1], `"',`)
		case split[0] == "urlBase:":
			output.URLBase = strings.Trim(split[1], `"',`)
		case split[0] == "isProduction:":
			output.IsProduction = strings.Trim(split[1], `"',`) == "true"
		case strings.HasPrefix(split[0], "window."):
			output.App = strings.TrimPrefix(split[0], "window.")
		}
	}

	if err := scanner.Err(); err != nil {
		return output, fmt.Errorf("scanning HTTP response: %w", err)
	}

	return output, nil
}
