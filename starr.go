// Package starr is a library for interacting with the APIs in Radarr, Lidarr, Sonarr
// and Readarr. It consists of the main starr package and one sub package for each
// starr application. In the basic use, you create a starr Config that contains an
// API key and an App URL. Pass this into one of the other packages (like radarr),
// and it's used as an interface to make API calls.
//
// You can either call starr.New() to build an http.Client for you, or create a
// starr.Config that contains one you control. If you pass a starr.Config into
// a sub package without an http Client, it will be created for you. There are
// a lot of option to set this code up from simple and easy to more advanced.
//
// The sub package contain methods and data structures for a number of API endpoints.
// Each app has somewhere between 50 and 100 API endpoints. This library currently
// covers about 10% of those. You can retrieve things like movies, albums, series
// and books. You can retrieve server status, authors, artists and items in queues.
// You can also add new media to each application with this library.
//
package starr

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// Defaults for New().
const (
	DefaultTimeout = 10 * time.Second
)

// Errors you may receive from this package.
var (
	// ErrInvalidStatusCode is returned when the server (*arr app) returns a bad status code during an API request.
	ErrInvalidStatusCode = fmt.Errorf("invalid status code, <200||>299")
	// ErrNilClient is returned if you attempt a request with a nil http.Client.
	ErrNilClient = fmt.Errorf("http.Client must not be nil")
	// ErrNilInterface is returned by *Into() methods when a nil interface is provided.
	ErrNilInterface = fmt.Errorf("cannot unmarshal data into a nil or empty interface")
	// ErrInvalidAPIKey is returned if we know the API key didn't work.
	ErrInvalidAPIKey = fmt.Errorf("API Key may be incorrect")
	// ErrRequestError is returned when bad input is provided.
	ErrRequestError = fmt.Errorf("request error")
)

// Config is the data needed to poll Radarr or Sonarr or Lidarr or Readarr.
// At a minimum, provide a URL and API Key.
// Set ValidSSL to true if the app has a valid SSL certificate.
// HTTPUser and HTTPPass are used for Basic HTTP auth, if enabled (not common).
// Username and Password are for non-API paths with native authentication enabled.
// Timeout and ValidSSL are used to create the http Client by sub packages. You
// may set those and call New() in the sub packages to create the http.Client
// pointer, or you can create your own http.Client before calling subpackage.New().
// MaxBody is only used if a DebugLog is provided, and causes payloads to truncate.
type Config struct {
	APIKey   string                       `json:"apiKey" toml:"api_key" xml:"api_key" yaml:"apiKey"`
	URL      string                       `json:"url" toml:"url" xml:"url" yaml:"url"`
	HTTPPass string                       `json:"httpPass" toml:"http_pass" xml:"http_pass" yaml:"httpPass"`
	HTTPUser string                       `json:"httpUser" toml:"http_user" xml:"http_user" yaml:"httpUser"`
	Username string                       `json:"username" toml:"username" xml:"username" yaml:"username"`
	Password string                       `json:"password" toml:"password" xml:"password" yaml:"password"`
	Timeout  Duration                     `json:"timeout" toml:"timeout" xml:"timeout" yaml:"timeout"`
	ValidSSL bool                         `json:"validSsl" toml:"valid_ssl" xml:"valid_ssl" yaml:"validSsl"`
	MaxBody  int                          `json:"maxBody" toml:"max_body" xml:"max_body" yaml:"maxBody"`
	Client   *http.Client                 `json:"-" toml:"-" xml:"-" yaml:"-"`
	Debugf   func(string, ...interface{}) `json:"-" toml:"-" xml:"-" yaml:"-"`
	cookie   bool
}

// Duration is used to Unmarshal text into a time.Duration value.
type Duration struct{ time.Duration }

// New returns a *starr.Config pointer. This pointer is safe to modify
// further before passing it into one of the arr app New() procedures.
// Set Debugf if you want this library to print debug messages (payloads, etc).
func New(apiKey, appURL string, timeout time.Duration) *Config {
	if timeout == 0 {
		timeout = DefaultTimeout
	}

	return &Config{
		APIKey:   apiKey,
		URL:      appURL,
		HTTPUser: "",
		HTTPPass: "",
		Username: "",
		Password: "",
		MaxBody:  0,
		ValidSSL: false,
		Timeout:  Duration{Duration: timeout},
		Client:   nil, // Let each sub package handle its own client.
		Debugf:   nil,
	}
}

// UnmarshalText parses a duration type from a config file.
func (d *Duration) UnmarshalText(data []byte) error {
	var err error

	d.Duration, err = time.ParseDuration(string(data))
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}

	return nil
}

// String returns a Duration as string without trailing zero units.
func (d Duration) String() string {
	dur := d.Duration.String()
	if len(dur) > 3 && dur[len(dur)-3:] == "m0s" {
		dur = dur[:len(dur)-2]
	}

	if len(dur) > 3 && dur[len(dur)-3:] == "h0m" {
		dur = dur[:len(dur)-2]
	}

	return dur
}

// GetURL attempts to fix the URL for a starr app.
// If the url base is missing it is added; this only checks the Location header.
// You should call this once at startup and update the URL provided.
func (c *Config) GetURL() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout.Duration)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.URL, nil)
	if err != nil {
		return c.URL, fmt.Errorf("creating request: %w", err)
	}

	client := &http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: !c.ValidSSL}, // nolint:gosec
		},
	}

	req.Header.Add("X-API-Key", c.APIKey)

	resp, err := client.Do(req)
	if err != nil {
		return c.URL, fmt.Errorf("making request: %w", err)
	}
	defer resp.Body.Close()

	_, _ = io.Copy(io.Discard, resp.Body) // read the whole body to avoid memory leaks.

	location, err := resp.Location()
	if err != nil {
		return c.URL, nil //nolint:nilerr // no location header, no error returned.
	}

	if strings.Contains(location.String(), "/login") {
		return c.URL, fmt.Errorf("redirected to login page while checking URL %s: %w", c.URL, ErrInvalidAPIKey)
	}

	return location.String(), nil
}
