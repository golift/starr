package starr

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"golift.io/starr/debuglog"
)

// App can be used to satisfy a context value key.
// It is not used in this library; provided for convenience.
type App string

// These constants are just here for convenience.
const (
	Emby     App = "Emby"
	Lidarr   App = "Lidarr"
	Plex     App = "Plex"
	Prowlarr App = "Prowlarr"
	Radarr   App = "Radarr"
	Readarr  App = "Readarr"
	Sonarr   App = "Sonarr"
	Whisparr App = "Whisparr"
)

// String turns an App name into a string.
func (a App) String() string {
	return string(a)
}

// Lower turns an App name into a lowercase string.
func (a App) Lower() string {
	return strings.ToLower(string(a))
}

// Client returns the default client, and is used if one is not passed in.
func Client(timeout time.Duration, verifySSL bool) *http.Client {
	return &http.Client{
		Timeout: timeout,
		CheckRedirect: func(_ *http.Request, _ []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: !verifySSL}, //nolint:gosec
		},
	}
}

// ClientWithDebug returns an http client with a debug logger enabled.
func ClientWithDebug(timeout time.Duration, verifySSL bool, logConfig debuglog.Config) *http.Client {
	client := Client(timeout, verifySSL)
	client.Transport = debuglog.NewLoggingRoundTripper(logConfig, client.Transport)

	return client
}

// Itoa converts an int64 to a string.
// Deprecated: Use starr.Str() instead.
func Itoa(v int64) string {
	return Str(v)
}

// Str converts numbers and booleans to a string.
func Str[I int | int64 | float64 | bool](val I) string {
	const (
		base10 = 10
		bits64 = 64
	)

	switch val := interface{}(val).(type) {
	case int:
		return strconv.Itoa(val)
	case bool:
		return strconv.FormatBool(val)
	case int64:
		return strconv.FormatInt(val, base10)
	case float64:
		return strconv.FormatFloat(val, 'f', -1, bits64)
	default:
		return fmt.Sprint(val)
	}
}

// String returns a pointer to a string.
func String(s string) *string {
	return &s
}

// True returns a pointer to a true boolean.
func True() *bool {
	s := true
	return &s
}

// False returns a pointer to a false boolean.
func False() *bool {
	s := false
	return &s
}

// Int64 returns a pointer to the provided integer.
func Int64(s int64) *int64 {
	return &s
}
