package starr

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"reflect"
	"strings"
)

// API is the beginning of every API path.
const API = "api"

/* The methods in this file provide assumption-ridden HTTP calls for Starr apps. */

// Request contains the GET and/or POST values for an HTTP request.
type Request struct {
	URI   string     // path portion of the URL.
	Query url.Values // GET parameters work for any request type.
	Body  io.Reader  // Used in PUT, POST, DELETE. Not for GET.
	uri   string     // derived URI.
}

// Req makes an authenticated request to a starr application and returns the response.
// Do not forget to read and close the response Body if there is no error.
func (c *Config) Req(ctx context.Context, method string, req *Request) (*http.Response, error) {
	return c.req(ctx, method, req)
}

// api is an internal function to call an api path.
func (c *Config) api(ctx context.Context, method string, req *Request) (*http.Response, error) {
	req.uri = SetAPIPath(req.URI)
	return c.req(ctx, method, req)
}

// req is our abstraction method for calling a starr application.
func (c *Config) req(ctx context.Context, method string, req *Request) (*http.Response, error) {
	if c.Client == nil { // we must have an http client.
		return nil, ErrNilClient
	}

	if req == nil {
		return nil, fmt.Errorf("%w: request uri required", ErrRequestError)
	}

	if req.uri == "" {
		req.uri = req.URI
	}

	httpReq, err := http.NewRequestWithContext(ctx, method, strings.TrimSuffix(c.URL, "/")+req.uri, req.Body)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext(%s): %w", req.uri, err)
	}

	c.SetHeaders(httpReq)

	if req.Query != nil {
		httpReq.URL.RawQuery = req.Query.Encode()
	}

	resp, err := c.Client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("httpClient.Do(req): %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		closeResp(resp)
		return nil, fmt.Errorf("failed: %v (status: %s): %w",
			httpReq.RequestURI, resp.Status, ErrInvalidStatusCode)
	}

	return resp, nil
}

func closeResp(resp *http.Response) {
	if resp != nil && resp.Body != nil {
		_, _ = io.ReadAll(resp.Body)
		resp.Body.Close()
	}
}

// SetHeaders sets all our request headers based on method and other data.
func (c *Config) SetHeaders(req *http.Request) {
	// This app allows http auth, in addition to api key (nginx proxy).
	if auth := c.HTTPUser + ":" + c.HTTPPass; auth != ":" {
		req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(auth)))
	}

	if req.Body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	if req.Method == http.MethodPost && strings.HasSuffix(req.URL.RequestURI(), "/login") {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else {
		req.Header.Set("Accept", "application/json")
	}

	req.Header.Set("User-Agent", "go-starr: https://"+reflect.TypeOf(Config{}).PkgPath()) //nolint:exhaustivestruct
	req.Header.Set("X-API-Key", c.APIKey)
}

// SetAPIPath makes sure the path starts with /api.
func SetAPIPath(uriPath string) string {
	if strings.HasPrefix(uriPath, API+"/") ||
		strings.HasPrefix(uriPath, path.Join("/", API)+"/") {
		return path.Join("/", uriPath)
	}

	return path.Join("/", API, uriPath)
}
