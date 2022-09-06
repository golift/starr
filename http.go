package starr

import (
	"context"
	"encoding/base64"
	"encoding/json"
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
	URI   string     // Required: path portion of the URL.
	Query url.Values // GET parameters work for any request type.
	Body  io.Reader  // Used in PUT, POST, DELETE. Not for GET.
}

// String turns a request into a string. Usually used in error messages.
func (r *Request) String() string {
	return r.URI
}

// Req makes an authenticated request to a starr application and returns the response.
// Do not forget to read and close the response Body if there is no error.
func (c *Config) Req(ctx context.Context, method string, req Request) (*http.Response, error) {
	return c.req(ctx, method, req)
}

// api is an internal function to call an api path.
func (c *Config) api(ctx context.Context, method string, req Request) (*http.Response, error) {
	req.URI = SetAPIPath(req.URI)
	return c.req(ctx, method, req)
}

// req is our abstraction method for calling a starr application.
func (c *Config) req(ctx context.Context, method string, req Request) (*http.Response, error) {
	if c.Client == nil { // we must have an http client.
		return nil, ErrNilClient
	}

	httpReq, err := http.NewRequestWithContext(ctx, method, strings.TrimSuffix(c.URL, "/")+req.URI, req.Body)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext(%s): %w", req.URI, err)
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
		return nil, parseNon200(httpReq, resp)
	}

	return resp, nil
}

// parseNon200 attempts to extract an error message from a non-200 response.
func parseNon200(req *http.Request, resp *http.Response) error {
	defer resp.Body.Close()

	var msg struct {
		Msg string `json:"message"`
	}

	reply, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed: %v (status: %s): %w",
			req.RequestURI, resp.Status, ErrInvalidStatusCode)
	}

	if err := json.Unmarshal(reply, &msg); err == nil {
		return fmt.Errorf("failed: %v (status: %s): %w: %s",
			req.RequestURI, resp.Status, ErrInvalidStatusCode, msg.Msg)
	}

	const maxSize = 200 // arbitrary max size

	replyStr := string(reply)
	if len(replyStr) > maxSize {
		return fmt.Errorf("failed: %v (status: %s): %w: %s",
			req.RequestURI, resp.Status, ErrInvalidStatusCode, replyStr[:maxSize])
	}

	return fmt.Errorf("failed: %v (status: %s): %w: %s",
		req.RequestURI, resp.Status, ErrInvalidStatusCode, replyStr)
}

// closeResp should be used to close requests that don't require a response body.
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
