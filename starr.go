package starr

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// ErrInvalidStatusCode is returned when the server (*arr app) returns a bad status code during an API request.
var ErrInvalidStatusCode = fmt.Errorf("invalid status code, not 200")

// Config is the data needed to poll Radarr or Sonarr or Lidarr or Readarr.
// At a minimum, provide a URL and API Key.
// Set ValidSSL to true if the app has a valid SSL certificate.
// HTTPUser and HTTPPass are used for Basic HTTP auth, if enabled (not common).
type Config struct {
	APIKey   string   `json:"api_key" toml:"api_key" xml:"api_key" yaml:"api_key"`
	URL      string   `json:"url" toml:"url" xml:"url" yaml:"url"`
	HTTPPass string   `json:"http_pass" toml:"http_pass" xml:"http_pass" yaml:"http_pass"`
	HTTPUser string   `json:"http_user" toml:"http_user" xml:"http_user" yaml:"http_user"`
	Timeout  Duration `json:"timeout" toml:"timeout" xml:"timeout" yaml:"timeout"`
	ValidSSL bool     `json:"valid_ssl" toml:"valid_ssl" xml:"valid_ssl" yaml:"valid_ssl"`
}

// Duration is used to UnmarshalTOML into a time.Duration value.
type Duration struct{ time.Duration }

// UnmarshalText parses a duration type from a config file.
func (d *Duration) UnmarshalText(data []byte) (err error) {
	d.Duration, err = time.ParseDuration(string(data))

	return
}

// Req makes a http request, with some additions.
// path = "/query", params = "sort_by=timeleft&order=asc" (as url.Values).
func (c *Config) Req(path string, params url.Values) ([]byte, error) {
	client := &http.Client{
		Timeout: c.Timeout.Duration,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: !c.ValidSSL}, // nolint: G402
		},
	}
	path = c.fixPath(path)

	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout.Duration)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext(path): %w", err)
	}

	params["apikey"] = []string{c.APIKey}
	req.URL.RawQuery = params.Encode()

	// This app allows http auth, in addition to api key (nginx proxy).
	if auth := c.HTTPUser + ":" + c.HTTPPass; auth != ":" {
		auth = "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
		req.Header.Set("Authorization", auth)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("d.Do(req): %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed: %v (status: %v/%v): %w",
			path, resp.StatusCode, resp.Status, ErrInvalidStatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll: %w", err)
	}

	// log.Println(string(body))
	return body, nil
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
