package starr

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
)

/* The methods in this file satisfy the APIer
 * interface and provide basic HTTP request support. */

// Req does an HTTP GET. Deprecated, provided for legacy use.
func (c *Config) Req(path string, params url.Values) ([]byte, error) {
	return c.Get(path, params)
}

// GetInto performs an HTTP GET against an API path and unmarshals the payload into the provided pointer interface.
func (c *Config) GetInto(path string, params url.Values, v interface{}) error {
	if data, err := c.Get(path, params); err != nil {
		return err
	} else if err = json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("json parse error: %w", err)
	}

	return nil
}

// PostInto performs an HTTP POST against an API path and unmarshals the payload into the provided pointer interface.
func (c *Config) PostInto(path string, params url.Values, body []byte, v interface{}) error {
	if data, err := c.Post(path, params, body); err != nil {
		return err
	} else if err = json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("json parse error: %w", err)
	}

	return nil
}

// Get makes a GET http request and returns the body.
func (c *Config) Get(path string, params url.Values) ([]byte, error) {
	path, useParam := c.fixPath(path)

	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout.Duration)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext(path): %w", err)
	}

	resp, err := c.getResp(useParam, params, req)
	if err != nil {
		return nil, fmt.Errorf("client.Do(): %w", err)
	}
	defer resp.Body.Close()

	return c.getBody(resp)
}

// Post makes a POST http request and returns the body.
func (c *Config) Post(path string, params url.Values, body []byte) ([]byte, error) {
	path, useParam := c.fixPath(path)
	bodyReader := bytes.NewBuffer(body)

	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout.Duration)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, path, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext(path): %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.getResp(useParam, params, req)
	if err != nil {
		return nil, fmt.Errorf("client.Do(): %w", err)
	}
	defer resp.Body.Close()

	return c.getBody(resp)
}

func (c *Config) getBody(resp *http.Response) ([]byte, error) {
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll: %w", err)
	}

	// log.Println(string(body))
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		err = fmt.Errorf("failed: %v (status: %s): %w",
			resp.Request.RequestURI, resp.Status, fmt.Errorf("%w: %s", ErrInvalidStatusCode, string(b)))
	}

	return b, err
}

func (c *Config) getResp(useParam bool, params url.Values, req *http.Request) (*http.Response, error) {
	if c.Client == nil {
		return nil, ErrNilClient
	}

	if params == nil {
		params = make(url.Values)
	}

	if useParam {
		// api paths with /v1 or /v3 in them use a header.
		params.Add("apikey", c.APIKey)
	}

	req.Header.Set("X-API-Key", c.APIKey)
	req.URL.RawQuery = params.Encode()

	// This app allows http auth, in addition to api key (nginx proxy).
	if auth := c.HTTPUser + ":" + c.HTTPPass; auth != ":" {
		auth = "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
		req.Header.Set("Authorization", auth)
	}

	return c.Client.Do(req)
}

func (c *Config) fixPath(uriPath string) (string, bool) {
	if strings.Contains(uriPath, "/api") {
		uriPath = path.Join("/", uriPath)
	} else {
		uriPath = path.Join("/", "api", uriPath)
	}

	return strings.TrimSuffix(c.URL, "/") + uriPath, !strings.HasPrefix(uriPath, "/api/v")
}
