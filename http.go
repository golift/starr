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

// req returns the body in io.ReadCloser form (read and close it yourself).
func (c *Config) req(
	ctx context.Context,
	uri string,
	method string,
	params url.Values,
	body io.Reader,
) (*http.Response, error) {
	if c.Client == nil { // we must have an http client.
		return nil, ErrNilClient
	}

	req, err := http.NewRequestWithContext(ctx, method, c.setPath(uri), body)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext(path): %w", err)
	}

	c.setHeaders(req)

	if params != nil {
		req.URL.RawQuery = params.Encode()
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("httpClient.Do(req): %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		closeResp(resp)
		return nil, fmt.Errorf("failed: %v (status: %s): %w",
			req.RequestURI, resp.Status, ErrInvalidStatusCode)
	}

	return resp, nil
}

func closeResp(resp *http.Response) {
	if resp != nil && resp.Body != nil {
		_, _ = io.ReadAll(resp.Body)
		resp.Body.Close()
	}
}

// setHeaders sets all our request headers based on method and other data.
func (c *Config) setHeaders(req *http.Request) {
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

// setPath makes sure the path starts with /api and returns the full URL.
func (c *Config) setPath(uriPath string) string {
	if strings.HasPrefix(uriPath, API+"/") ||
		strings.HasPrefix(uriPath, path.Join("/", API)+"/") {
		uriPath = path.Join("/", uriPath)
	} else {
		uriPath = path.Join("/", API, uriPath)
	}

	return strings.TrimSuffix(c.URL, "/") + uriPath
}
