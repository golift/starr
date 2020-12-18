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
	"strings"
	"time"
)

var (
	// ErrInvalidStatusCode is returned when the server (*arr app) returns a bad status code during an API request.
	ErrInvalidStatusCode = fmt.Errorf("invalid status code, <200||>299")
	// ErrNilClient is returned if you attempt a request with a nil http.Client.
	ErrNilClient = fmt.Errorf("http.Client must not be nil")
)

// Config is the data needed to poll Radarr or Sonarr or Lidarr or Readarr.
// At a minimum, provide a URL and API Key.
// Set ValidSSL to true if the app has a valid SSL certificate.
// HTTPUser and HTTPPass are used for Basic HTTP auth, if enabled (not common).
type Config struct {
	APIKey   string       `json:"api_key" toml:"api_key" xml:"api_key" yaml:"api_key"`
	URL      string       `json:"url" toml:"url" xml:"url" yaml:"url"`
	HTTPPass string       `json:"http_pass" toml:"http_pass" xml:"http_pass" yaml:"http_pass"`
	HTTPUser string       `json:"http_user" toml:"http_user" xml:"http_user" yaml:"http_user"`
	Timeout  Duration     `json:"timeout" toml:"timeout" xml:"timeout" yaml:"timeout"`
	ValidSSL bool         `json:"valid_ssl" toml:"valid_ssl" xml:"valid_ssl" yaml:"valid_ssl"`
	Client   *http.Client `json:"-" toml:"-" xml:"-" yaml:"-"` // required!
}

// Duration is used to UnmarshalTOML into a time.Duration value.
type Duration struct{ time.Duration }

// UnmarshalText parses a duration type from a config file.
func (d *Duration) UnmarshalText(data []byte) (err error) {
	d.Duration, err = time.ParseDuration(string(data))

	return
}

// Req does an HTTP GET. Provided for legacy apps.
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
	path = c.fixPath(path)

	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout.Duration)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext(path): %w", err)
	}

	resp, err := c.getResp(path, params, req)
	if err != nil {
		return nil, fmt.Errorf("client.Do(): %w", err)
	}
	defer resp.Body.Close()

	return c.getBody(resp)
}

// Post makes a POST http request and returns the body.
func (c *Config) Post(path string, params url.Values, body []byte) ([]byte, error) {
	path = c.fixPath(path)
	bodyReader := bytes.NewBuffer(body)

	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout.Duration)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, path, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext(path): %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.getResp(path, params, req)
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
			resp.Request.RequestURI, resp.Status, ErrInvalidStatusCode)
	}

	return b, err
}

func (c *Config) getResp(path string, params url.Values, req *http.Request) (*http.Response, error) {
	if c.Client == nil {
		return nil, ErrNilClient
	}

	if params == nil {
		params = make(url.Values)
	}

	if strings.HasPrefix(path, "/api/v") {
		// api paths with /v1 or /v3 in them use a header.
		req.Header.Set("X-API-Key", c.APIKey)
	} else {
		params.Add("apikey", c.APIKey)
	}

	req.URL.RawQuery = params.Encode()

	// This app allows http auth, in addition to api key (nginx proxy).
	if auth := c.HTTPUser + ":" + c.HTTPPass; auth != ":" {
		auth = "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
		req.Header.Set("Authorization", auth)
	}

	return c.Client.Do(req)
}

func (c *Config) fixPath(path string) string {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	if strings.HasSuffix(c.URL, "/") {
		return c.URL + "api" + path
	}

	return c.URL + "/api" + path
}
