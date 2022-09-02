package prowlarr

import (
	"golift.io/starr"
)

// Prowlarr contains all the methods to interact with a Prowlarr server.
type Prowlarr struct {
	starr.APIer
}

// APIver is the Prowlarr API version supported by this library.
const APIver = "v1"

// New returns a Prowlarr object used to interact with the Prowlarr API.
func New(config *starr.Config) *Prowlarr {
	if config.Client == nil {
		config.Client = starr.Client(0, false)
	}

	return &Prowlarr{APIer: config}
}
