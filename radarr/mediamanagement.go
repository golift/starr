package radarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"golift.io/starr"
)

// Define Base Path for media management calls.
const bpMediaManagement = APIver + "/config/mediaManagement"

// MediaManagement represents the /config/mediaManagement endpoint.
type MediaManagement struct {
	AutoRenameFolders                       bool   `json:"autoRenameFolders,omitempty"`
	AutoUnmonitorPreviouslyDownloadedMovies bool   `json:"autoUnmonitorPreviouslyDownloadedMovies,omitempty"`
	CopyUsingHardlinks                      bool   `json:"copyUsingHardlinks,omitempty"`
	CreateEmptyMovieFolders                 bool   `json:"createEmptyMovieFolders,omitempty"`
	DeleteEmptyFolders                      bool   `json:"deleteEmptyFolders,omitempty"`
	EnableMediaInfo                         bool   `json:"enableMediaInfo,omitempty"`
	ImportExtraFiles                        bool   `json:"importExtraFiles,omitempty"`
	PathsDefaultStatic                      bool   `json:"pathsDefaultStatic,omitempty"`
	SetPermissionsLinux                     bool   `json:"setPermissionsLinux,omitempty"`
	SkipFreeSpaceCheckWhenImporting         bool   `json:"skipFreeSpaceCheckWhenImporting,omitempty"`
	ID                                      int64  `json:"id,omitempty"`
	MinimumFreeSpaceWhenImporting           int64  `json:"minimumFreeSpaceWhenImporting,omitempty"`
	RecycleBinCleanupDays                   int64  `json:"recycleBinCleanupDays,omitempty"`
	ChmodFolder                             string `json:"chmodFolder,omitempty"`
	ChownGroup                              string `json:"chownGroup,omitempty"`
	DownloadPropersAndRepacks               string `json:"downloadPropersAndRepacks,omitempty"`
	ExtraFileExtensions                     string `json:"extraFileExtensions,omitempty"`
	FileDate                                string `json:"fileDate,omitempty"`
	RecycleBin                              string `json:"recycleBin,omitempty"`
	RescanAfterRefresh                      string `json:"rescanAfterRefresh,omitempty"`
}

// GetMediaManagement returns the media management.
func (r *Radarr) GetMediaManagement() (*MediaManagement, error) {
	return r.GetMediaManagementContext(context.Background())
}

// GetMediaManagement returns the media management.
func (r *Radarr) GetMediaManagementContext(ctx context.Context) (*MediaManagement, error) {
	var output MediaManagement

	req := starr.Request{URI: bpMediaManagement}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateMediaManagement updates the media management.
func (r *Radarr) UpdateMediaManagement(mMgt *MediaManagement) (*MediaManagement, error) {
	return r.UpdateMediaManagementContext(context.Background(), mMgt)
}

// UpdateMediaManagementContext updates the media management.
func (r *Radarr) UpdateMediaManagementContext(ctx context.Context, mMgt *MediaManagement) (*MediaManagement, error) {
	var output MediaManagement

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(mMgt); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpMediaManagement, err)
	}

	req := starr.Request{URI: bpMediaManagement, Body: &body}
	if err := r.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}
