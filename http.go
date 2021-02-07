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

/* The methods in this file provided assumption-ridden HTTP calls for Starr apps. */

func (c *Config) req(path, method string, params url.Values, body io.Reader) ([]byte, error) {
	if c.Client == nil { // we must have an http client.
		return nil, ErrNilClient
	}

	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout.Duration)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, method, c.setPathParams(path, params), body)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext(path): %w", err)
	}

	c.setHeaders(req)
	req.URL.RawQuery = params.Encode()

	return c.getBody(req)
}

// setHeaders sets all our request headers.
func (c *Config) setHeaders(req *http.Request) {
	// This app allows http auth, in addition to api key (nginx proxy).
	if auth := c.HTTPUser + ":" + c.HTTPPass; auth != ":" {
		req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(auth)))
	}

	if req.Body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "go-starr: https://"+reflect.TypeOf(Config{}).PkgPath())
	req.Header.Set("X-API-Key", c.APIKey)
}

// getBody makes an http request and returns the response body if there are no errors.
func (c *Config) getBody(req *http.Request) ([]byte, error) {
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("httpClient.Do(req): %w", err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll: %w", err)
	}

	// fmt.Println(string(b))
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return b, fmt.Errorf("failed: %v (status: %s): %w: %s",
			resp.Request.RequestURI, resp.Status, ErrInvalidStatusCode, string(b))
	}

	return b, nil
}

// setPathParams makes sure the path starts with /api and returns the full URL.
// Also makes sure params is not nil (so it can be encoded later).
// Sets the apikey as a path parameter for use by older radarr/sonarr versions.
func (c *Config) setPathParams(uriPath string, params url.Values) string {
	if strings.Contains(uriPath, "api/") {
		uriPath = path.Join("/", uriPath)
	} else {
		uriPath = path.Join("/", "api", uriPath)
	}

	if params == nil {
		params = make(url.Values)
	}

	if !strings.HasPrefix(uriPath, "/api/v") {
		// api paths with /v1 or /v3 in them use a header instead.
		params.Add("apikey", c.APIKey)
	}

	return strings.TrimSuffix(c.URL, "/") + uriPath
}
