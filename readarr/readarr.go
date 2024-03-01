package readarr

import (
	"context"
	"fmt"
	"strings"

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
		config.Client = starr.Client(0, false)
	}

	config.URL = strings.TrimSuffix(config.URL, "/")

	return &Readarr{APIer: config}
}

// bp means base path. You'll see it a lot in these files.
const bpPing = "/ping" // ping has no api or version prefix.

// Ping returns an error if the starr instance does not respond with a 200 to an HTTP /ping request.
func (r *Readarr) Ping() error {
	return r.PingContext(context.Background())
}

// PingContext returns an error if the starr instance does not respond with a 200 to an HTTP /ping request.
func (r *Readarr) PingContext(ctx context.Context) error {
	req := starr.Request{URI: bpPing}

	resp, err := r.Get(ctx, req)
	if err != nil {
		return fmt.Errorf("api.Get(%s): %w", &req, err)
	}
	defer resp.Body.Close()

	return nil
}
