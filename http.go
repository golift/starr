package starr

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
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
	req, err := c.newReq(ctx, c.URL+uri, method, params, body)
	if err != nil {
		return nil, err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("httpClient.Do(req): %w", err)
	}

	return resp, nil
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

	c.setHeaders(req)

	if params != nil {
		req.URL.RawQuery = params.Encode()
	}

	return req, nil
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
