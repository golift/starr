package readarr

import (
	"golift.io/starr"
)

// APIver is the Readarr API version supported by this library.
const APIver = "v1"

// Readarr contains all the methods to interact with a Readarr server.
type Readarr struct {
	starr.APIer
}

// Filter values are integers. Given names for ease of discovery.
// https://github.com/Readarr/Readarr/blob/de72cfcaaa22495c7ce9fcb596a93beff6efb3d6/src/NzbDrone.Core/History/EntityHistory.cs#L31-L43
//
//nolint:lll
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
		config.Client = starr.Client(config.Timeout.Duration, config.ValidSSL)
	}

	return &Readarr{APIer: config}
}
