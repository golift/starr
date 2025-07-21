package lidarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"golift.io/starr"
)

// Define Base Path for MediaManagement calls.
const bpMediaManagement = APIver + "/config/mediaManagement"

// MediaManagement represents the /config/mediamanagement endpoint.
type MediaManagement struct {
	SkipFreeSpaceCheckWhenImporting         bool   `json:"skipFreeSpaceCheckWhenImporting"`
	CopyUsingHardlinks                      bool   `json:"copyUsingHardlinks"`
	ImportExtraFiles                        bool   `json:"importExtraFiles"`
	WatchLibraryForChanges                  bool   `json:"watchLibraryForChanges"`
	AutoUnmonitorPreviouslyDownloadedTracks bool   `json:"autoUnmonitorPreviouslyDownloadedTracks"`
	CreateEmptyArtistFolders                bool   `json:"createEmptyArtistFolders"`
	DeleteEmptyFolders                      bool   `json:"deleteEmptyFolders"`
	SetPermissionsLinux                     bool   `json:"setPermissionsLinux"`
	MinimumFreeSpaceWhenImporting           int64  `json:"minimumFreeSpaceWhenImporting"`
	RecycleBinCleanupDays                   int64  `json:"recycleBinCleanupDays"`
	ID                                      int64  `json:"id"` // always 1
	RecycleBin                              string `json:"recycleBin"`
	DownloadPropersAndRepacks               string `json:"downloadPropersAndRepacks"`
	FileDate                                string `json:"fileDate"`
	RescanAfterRefresh                      string `json:"rescanAfterRefresh"`
	AllowFingerprinting                     string `json:"allowFingerprinting"`
	ChmodFolder                             string `json:"chmodFolder"`
	ChownGroup                              string `json:"chownGroup"`
	ExtraFileExtensions                     string `json:"extraFileExtensions"`
}

// GetMediaManagement returns the mediaManagement.
func (l *Lidarr) GetMediaManagement() (*MediaManagement, error) {
	return l.GetMediaManagementContext(context.Background())
}

// GetMediaManagement returns the Media Management.
func (l *Lidarr) GetMediaManagementContext(ctx context.Context) (*MediaManagement, error) {
	var output MediaManagement

	req := starr.Request{URI: bpMediaManagement}
	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateMediaManagement updates the Media Management.
func (l *Lidarr) UpdateMediaManagement(mMgt *MediaManagement) (*MediaManagement, error) {
	return l.UpdateMediaManagementContext(context.Background(), mMgt)
}

// UpdateMediaManagementContext updates the Media Management.
func (l *Lidarr) UpdateMediaManagementContext(ctx context.Context, mMgt *MediaManagement) (*MediaManagement, error) {
	var output MediaManagement

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(mMgt); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpMediaManagement, err)
	}

	req := starr.Request{URI: bpMediaManagement, Body: &body}
	if err := l.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}
