package prowlarr

import (
	"crypto/tls"
	"net/http"

	"golift.io/starr"
)

// Prowlarr contains all the methods to interact with a Prowlarr server.
type Prowlarr struct {
	starr.APIer
}

// New returns a Prowlarr object used to interact with the Prowlarr API.
func New(config *starr.Config) *Prowlarr {
	if config.Client == nil {
		//nolint:exhaustivestruct,gosec
		config.Client = &http.Client{
			Timeout: config.Timeout.Duration,
			CheckRedirect: func(r *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: !config.ValidSSL},
			},
		}
	}

	if config.Debugf == nil {
		config.Debugf = func(string, ...interface{}) {}
	}

	return &Prowlarr{APIer: config}
}
