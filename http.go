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

// Req makes an authenticated request to a starr application and returns the response.
// Do not forget to read and close the response Body if there is no error.
func (c *Config) Req(
	ctx context.Context,
	uri string,
	method string,
	params url.Values,
	body io.Reader,
) (*http.Response, error) {
	return c.req(ctx, c.URL+uri, method, params, body)
}

// api is an internal function to call an api path.
func (c *Config) api(
	ctx context.Context,
	uri string,
	method string,
	params url.Values,
	body io.Reader,
) (*http.Response, error) {
	return c.req(ctx, c.SetPath(uri), method, params, body)
}

// req is our abstraction method for calling a starr application.
func (c *Config) req(
	ctx context.Context,
	url string,
	method string,
	params url.Values,
	body io.Reader,
) (*http.Response, error) {
	if c.Client == nil { // we must have an http client.
		return nil, ErrNilClient
	}

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext(path): %w", err)
	}

	c.SetHeaders(req)

	if params != nil {
		req.URL.RawQuery = params.Encode()
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("httpClient.Do(req): %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, parseNon200(req, resp)
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

// SetPath makes sure the path starts with /api and returns the full URL.
func (c *Config) SetPath(uriPath string) string {
	if strings.HasPrefix(uriPath, API+"/") ||
		strings.HasPrefix(uriPath, path.Join("/", API)+"/") {
		uriPath = path.Join("/", uriPath)
	} else {
		uriPath = path.Join("/", API, uriPath)
	}

	return strings.TrimSuffix(c.URL, "/") + uriPath
}
