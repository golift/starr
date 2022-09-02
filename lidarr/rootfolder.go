package lidarr

import (
	"context"
	"fmt"

	"golift.io/starr"
)

// RootFolder is the /api/v1/rootfolder endpoint.
type RootFolder struct {
	ID              int64         `json:"id"`
	Path            string        `json:"path"`
	FreeSpace       int64         `json:"freeSpace"`
	TotalSpace      int64         `json:"totalSpace"`
	UnmappedFolders []*starr.Path `json:"unmappedFolders"`
}

// GetRootFolders returns all configured root folders.
func (l *Lidarr) GetRootFolders() ([]*RootFolder, error) {
	return l.GetRootFoldersContext(context.Background())
}

// GetRootFoldersContext returns all configured root folders.
func (l *Lidarr) GetRootFoldersContext(ctx context.Context) ([]*RootFolder, error) {
	var folders []*RootFolder

	err := l.GetInto(ctx, "v1/rootFolder", nil, &folders)
	if err != nil {
		return nil, fmt.Errorf("api.Get(rootFolder): %w", err)
	}

	return folders, nil
}
