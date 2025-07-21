package sonarr

import (
	"context"
	"fmt"
	"strings"

	"golift.io/starr"
)

// APIver is the Sonarr API version supported by this library.
const APIver = "v3"

// Sonarr contains all the methods to interact with a Sonarr server.
type Sonarr struct {
	starr.APIer
}

// Filter values are integers. Given names for ease of discovery.
// https://github.com/Sonarr/Sonarr/blob/0cb8d93069d6310abd39ee2fe73219e17aa83fe6/src/NzbDrone.Core/History/EpisodeHistory.cs#L34-L41
//
//nolint:lll
const (
	FilterUnknown starr.Filtering = iota
	FilterGrabbed
	FilterSeriesFolderImported
	FilterDownloadFolderImported
	FilterDownloadFailed
	FilterDeleted
	FilterRenamed
	FilterImportFailed
)

// New returns a Sonarr object used to interact with the Sonarr API.
func New(config *starr.Config) *Sonarr {
	if config.Client == nil {
		config.Client = starr.Client(0, false)
	}

	config.URL = strings.TrimSuffix(config.URL, "/")

	return &Sonarr{APIer: config}
}

// bp means base path. You'll see it a lot in these files.
const bpPing = "/ping" // ping has no api or version prefix.

// Ping returns an error if the starr instance does not respond with a 200 to an HTTP /ping request.
func (s *Sonarr) Ping() error {
	return s.PingContext(context.Background())
}

// PingContext returns an error if the starr instance does not respond with a 200 to an HTTP /ping request.
func (s *Sonarr) PingContext(ctx context.Context) error {
	req := starr.Request{URI: bpPing}

	resp, err := s.Get(ctx, req)
	if err != nil {
		return fmt.Errorf("api.Get(%s): %w", &req, err)
	}
	defer resp.Body.Close()

	return nil
}
