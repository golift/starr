package radarr

import (
	"golift.io/starr"
)

// APIver is the Radarr API version supported by this library.
const APIver = "v3"

// Radarr contains all the methods to interact with a Radarr server.
type Radarr struct {
	starr.APIer
}

// Filter values are integers. Given names for ease of discovery.
// https://github.com/Radarr/Radarr/blob/2bca1a71a2ed5130ea642343cb76250f3bf5bc4e/src/NzbDrone.Core/History/History.cs#L33-L44
//
//nolint:lll
const (
	FilterUnknown starr.Filtering = iota
	FilterGrabbed
	_ // 2 is unused
	FilterDownloadFolderImported
	FilterDownloadFailed
	_ // 5 is unused. FilterDeleted
	FilterFileDeleted
	_ // FilterFolderImported // not used yet, 1/17/2022
	FilterRenamed
	FilterIgnored
)

// New returns a Radarr object used to interact with the Radarr API.
func New(config *starr.Config) *Radarr {
	if config.Client == nil {
		config.Client = starr.Client(config.Timeout.Duration, config.ValidSSL)
	}

	return &Radarr{APIer: config}
}
