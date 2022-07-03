package starr

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"reflect"
	"strings"
)

// API is the beginning of every API path.
const API = "api"

/* The methods in this file provide assumption-ridden HTTP calls for Starr apps. */

// Req makes an http request and returns the body in []byte form (already read).
func (c *Config) Req(
	ctx context.Context,
	path string,
	method string,
	params url.Values,
	body io.Reader,
) (int, []byte, http.Header, error) {
	req, err := c.newReq(ctx, c.SetPath(path), method, params, body)
	if err != nil {
		return 0, nil, nil, err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return 0, nil, nil, fmt.Errorf("httpClient.Do(req): %w", err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, resp.Header, fmt.Errorf("ioutil.ReadAll: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return resp.StatusCode, respBody, resp.Header, fmt.Errorf("failed: %v (status: %s): %w: %s",
			resp.Request.RequestURI, resp.Status, ErrInvalidStatusCode, string(respBody))
	}

	return resp.StatusCode, respBody, resp.Header, nil
}

// body returns the body in io.ReadCloser form (read and close it yourself).
func (c *Config) body(
	ctx context.Context,
	uri string,
	method string,
	params url.Values,
	body io.Reader,
) (int, io.ReadCloser, http.Header, error) {
	req, err := c.newReq(ctx, c.URL+uri, method, params, body)
	if err != nil {
		return 0, nil, nil, err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return 0, nil, nil, fmt.Errorf("httpClient.Do(req): %w", err)
	}

	return resp.StatusCode, resp.Body, resp.Header, nil
}

func (c *Config) newReq(
	ctx context.Context,
	path string,
	method string,
	params url.Values,
	body io.Reader,
) (*http.Request, error) {
	if c.Client == nil { // we must have an http client.
		return nil, ErrNilClient
	}

	req, err := http.NewRequestWithContext(ctx, method, path, body)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext(path): %w", err)
	}

	c.SetHeaders(req)

	if params != nil {
		req.URL.RawQuery = params.Encode()
	}

	return req, nil
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
