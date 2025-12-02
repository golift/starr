package prowlarr

import (
	"context"
	"fmt"
	"strings"

	"golift.io/starr"
)

// Prowlarr contains all the methods to interact with a Prowlarr server.
type Prowlarr struct {
	starr.APIer `json:"-"`
}

// APIver is the Prowlarr API version supported by this library.
const APIver = "v1"

// New returns a Prowlarr object used to interact with the Prowlarr API.
func New(config *starr.Config) *Prowlarr {
	if config.Client == nil {
		config.Client = starr.Client(0, false)
	}

	config.URL = strings.TrimSuffix(config.URL, "/")

	return &Prowlarr{APIer: config}
}

// bp means base path. You'll see it a lot in these files.
const bpPing = "/ping" // ping has no api or version prefix.

// Ping returns an error if the Prowlarr instance does not respond with a 200 to an HTTP /ping request.
func (p *Prowlarr) Ping() error {
	return p.PingContext(context.Background())
}

// PingContext returns an error if the Prowlarr instance does not respond with a 200 to an HTTP /ping request.
func (p *Prowlarr) PingContext(ctx context.Context) error {
	req := starr.Request{URI: bpPing}

	resp, err := p.Get(ctx, req)
	if err != nil {
		return fmt.Errorf("api.Get(%s): %w", &req, err)
	}
	defer resp.Body.Close()

	return nil
}
