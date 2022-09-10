package radarr

import (
	"context"
	"fmt"

	"golift.io/starr"
)

const bpRootFolder = APIver + "/rootFolder"

// RootFolder is the /rootFolder endpoint.
type RootFolder struct {
	ID              int64         `json:"id"`
	Path            string        `json:"path"`
	Accessible      bool          `json:"accessible"`
	FreeSpace       int64         `json:"freeSpace"`
	UnmappedFolders []*starr.Path `json:"unmappedFolders"`
}

// GetRootFolders returns all configured root folders.
func (r *Radarr) GetRootFolders() ([]*RootFolder, error) {
	return r.GetRootFoldersContext(context.Background())
}

// GetRootFoldersContext returns all configured root folders.
func (r *Radarr) GetRootFoldersContext(ctx context.Context) ([]*RootFolder, error) {
	var output []*RootFolder

	req := starr.Request{URI: bpRootFolder}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}
