package starr

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Config is the data needed to poll Radarr or Sonarr.
type Config struct {
	APIKey   string   `json:"api_key" toml:"api_key" xml:"api_key" yaml:"api_key"`
	URL      string   `json:"url" toml:"url" xml:"url" yaml:"url"`
	HTTPPass string   `json:"http_pass" toml:"http_pass" xml:"http_pass" yaml:"http_pass"`
	HTTPUser string   `json:"http_user" toml:"http_user" xml:"http_user" yaml:"http_user"`
	Timeout  Duration `json:"timeout" toml:"timeout" xml:"timeout" yaml:"timeout"`
}

// Duration is used to UnmarshalTOML into a time.Duration value.
type Duration struct{ time.Duration }

// UnmarshalText parses a duration type from a config file.
func (d *Duration) UnmarshalText(data []byte) (err error) {
	d.Duration, err = time.ParseDuration(string(data))
	return
}

// Req makes a http request, with some additions.
// path = "/query", params = "sort_by=timeleft&order=asc" (as url.Values)
func (c *Config) Req(path string, params url.Values) ([]byte, error) {
	client := &http.Client{
		Timeout:   c.Timeout.Duration,
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
	}
	path = c.fixPath(path)

	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest(path): %v", err)
	}

	params.Set("apikey", c.APIKey)

	req.URL.RawQuery = params.Encode()

	// This app allows http auth, in addition to api key (nginx proxy).
	if auth := c.HTTPUser + ":" + c.HTTPPass; auth != ":" {
		auth = "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
		req.Header.Set("Authorization", auth)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("d.Do(req): %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed: %v (status: %v/%v)", path, resp.StatusCode, resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll: %v", err)
	}

	//log.Println(string(body))
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
