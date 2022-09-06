package readarr

import (
	"context"
	"fmt"

	"golift.io/starr"
)

const bpRootFolder = APIver + "/rootFolder"

// RootFolder is the /api/v1/rootfolder endpoint.
type RootFolder struct {
	ID                       int64  `json:"id"`
	Name                     string `json:"name"`
	Path                     string `json:"path"`
	DefaultMetadataProfileID int64  `json:"defaultMetadataProfileId"`
	DefaultQualityProfileID  int64  `json:"defaultQualityProfileId"`
	DefaultMonitorOption     string `json:"defaultMonitorOption"`
	DefaultTags              []int  `json:"defaultTags"`
	Port                     int    `json:"port"`
	OutputProfile            string `json:"outputProfile"`
	UseSsl                   bool   `json:"useSsl"`
	Accessible               bool   `json:"accessible"`
	IsCalibreLibrary         bool   `json:"isCalibreLibrary"`
	FreeSpace                int64  `json:"freeSpace"`
	TotalSpace               int64  `json:"totalSpace"`
}

// GetRootFolders returns all configured root folders.
func (r *Readarr) GetRootFolders() ([]*RootFolder, error) {
	return r.GetRootFoldersContext(context.Background())
}

// GetRootFoldersContext returns all configured root folders.
func (r *Readarr) GetRootFoldersContext(ctx context.Context) ([]*RootFolder, error) {
	var output []*RootFolder

	req := starr.Request{URI: bpRootFolder}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", req, err)
	}

	return output, nil
}
