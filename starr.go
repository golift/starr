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
	"fmt"
	"net/http"
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
)

// Config is the data needed to poll Radarr or Sonarr or Lidarr or Readarr.
// At a minimum, provide a URL and API Key.
// Set ValidSSL to true if the app has a valid SSL certificate.
// HTTPUser and HTTPPass are used for Basic HTTP auth, if enabled (not common).
// Timeout and ValidSSL are used to create the http Client by sub packages. You
// may set those and call New() in the sub packages to create the http.Client
// pointer, or you can create your own http.Client before calling subpackage.New().
type Config struct {
	APIKey   string       `json:"api_key" toml:"api_key" xml:"api_key" yaml:"api_key"`
	URL      string       `json:"url" toml:"url" xml:"url" yaml:"url"`
	HTTPPass string       `json:"http_pass" toml:"http_pass" xml:"http_pass" yaml:"http_pass"`
	HTTPUser string       `json:"http_user" toml:"http_user" xml:"http_user" yaml:"http_user"`
	Timeout  Duration     `json:"timeout" toml:"timeout" xml:"timeout" yaml:"timeout"`
	ValidSSL bool         `json:"valid_ssl" toml:"valid_ssl" xml:"valid_ssl" yaml:"valid_ssl"`
	Client   *http.Client `json:"-" toml:"-" xml:"-" yaml:"-"`
}

// Duration is used to Unmarshal text into a time.Duration value.
type Duration struct{ time.Duration }

// New returns a *starr.Config pointer. This pointer is safe to modify
// further before passing it into one of the arr app New() procedures.
func New(apiKey, appURL string, timeout time.Duration) *Config {
	if timeout == 0 {
		timeout = DefaultTimeout
	}

	return &Config{
		APIKey:   apiKey,
		URL:      appURL,
		HTTPUser: "",
		HTTPPass: "",
		ValidSSL: false,
		Timeout:  Duration{Duration: timeout},
		Client:   nil, // Let each sub package handle its own client.
	}
}

// UnmarshalText parses a duration type from a config file.
func (d *Duration) UnmarshalText(data []byte) (err error) {
	d.Duration, err = time.ParseDuration(string(data))

	return
}
