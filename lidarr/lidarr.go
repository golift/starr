package lidarr

import (
	"golift.io/starr"
)

// APIver is the Lidarr API version supported by this library.
const APIver = "v1"

// Lidarr contains all the methods to interact with a Lidarr server.
type Lidarr struct {
	starr.APIer
}

// Filter values are integers. Given names for ease of discovery.
// https://github.com/Lidarr/Lidarr/blob/c2adf078345f81012ddb5d2f384e2ee45ff7f1af/src/NzbDrone.Core/History/History.cs#L35-L45
//
//nolint:lll
const (
	FilterUnknown starr.Filtering = iota
	FilterGrabbed
	FilterArtistFolderImported
	FilterTrackFileImported
	FilterDownloadFailed
	FilterDeleted
	FilterRenamed
	FilterImportFailed
	FilterDownloadImported
	FilterRetagged
	FilterIgnored
)

// New returns a Lidarr object used to interact with the Lidarr API.
func New(config *starr.Config) *Lidarr {
	if config.Client == nil {
		config.Client = starr.Client(0, false)
	}

	return &Lidarr{APIer: config}
}
