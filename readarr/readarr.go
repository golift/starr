package readarr

import (
	"crypto/tls"
	"net/http"

	"golift.io/starr"
)

// APIver is the Readarr API version supported by this library.
const APIver = "v1"

// Readarr contains all the methods to interact with a Readarr server.
type Readarr struct {
	starr.APIer
}

// Filter values are integers. Given names for ease of discovery.
//nolint:lll
// https://github.com/Readarr/Readarr/blob/de72cfcaaa22495c7ce9fcb596a93beff6efb3d6/src/NzbDrone.Core/History/EntityHistory.cs#L31-L43
const (
	FilterAll starr.Filtering = iota
	FilterGrabbed
	_ // 2 is unused
	FilterBookFileImported
	FilterDownloadFailed
	FilterDeleted
	FilterRenamed
	FilterImportFailed
	FilterDownloadImported
	FilterRetagged
	FilterIgnored
)

// New returns a Readarr object used to interact with the Readarr API.
func New(config *starr.Config) *Readarr {
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

	return &Readarr{APIer: config}
}
