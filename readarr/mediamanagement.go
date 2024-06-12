package readarr

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
	SkipFreeSpaceCheckWhenImporting        bool   `json:"skipFreeSpaceCheckWhenImporting"`
	AutoUnmonitorPreviouslyDownloadedBooks bool   `json:"autoUnmonitorPreviouslyDownloadedBooks"`
	SetPermissionsLinux                    bool   `json:"setPermissionsLinux"`
	CreateEmptyAuthorFolders               bool   `json:"createEmptyAuthorFolders"`
	DeleteEmptyFolders                     bool   `json:"deleteEmptyFolders"`
	WatchLibraryForChanges                 bool   `json:"watchLibraryForChanges"`
	CopyUsingHardlinks                     bool   `json:"copyUsingHardlinks"`
	ImportExtraFiles                       bool   `json:"importExtraFiles"`
	MinimumFreeSpaceWhenImporting          int64  `json:"minimumFreeSpaceWhenImporting"`
	ID                                     int64  `json:"id"` // always 1
	RecycleBinCleanupDays                  int64  `json:"recycleBinCleanupDays"`
	RecycleBin                             string `json:"recycleBin"`
	DownloadPropersAndRepacks              string `json:"downloadPropersAndRepacks"`
	FileDate                               string `json:"fileDate"`
	RescanAfterRefresh                     string `json:"rescanAfterRefresh"`
	AllowFingerprinting                    string `json:"allowFingerprinting"`
	ChmodFolder                            string `json:"chmodFolder"`
	ChownGroup                             string `json:"chownGroup"`
	ExtraFileExtensions                    string `json:"extraFileExtensions"`
}

// GetMediaManagement returns the media management.
func (r *Readarr) GetMediaManagement() (*MediaManagement, error) {
	return r.GetMediaManagementContext(context.Background())
}

// GetMediaManagement returns the media management.
func (r *Readarr) GetMediaManagementContext(ctx context.Context) (*MediaManagement, error) {
	var output MediaManagement

	req := starr.Request{URI: bpMediaManagement}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateMediaManagement updates the media management.
func (r *Readarr) UpdateMediaManagement(mMgt *MediaManagement) (*MediaManagement, error) {
	return r.UpdateMediaManagementContext(context.Background(), mMgt)
}

// UpdateMediaManagementContext updates the media management.
func (r *Readarr) UpdateMediaManagementContext(ctx context.Context, mMgt *MediaManagement) (*MediaManagement, error) {
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
