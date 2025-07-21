package sonarr

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
	UseScriptImport                           bool   `json:"useScriptImport,omitempty"`
	AutoUnmonitorPreviouslyDownloadedEpisodes bool   `json:"autoUnmonitorPreviouslyDownloadedEpisodes,omitempty"`
	CopyUsingHardlinks                        bool   `json:"copyUsingHardlinks,omitempty"`
	CreateEmptySeriesFolders                  bool   `json:"createEmptySeriesFolders,omitempty"`
	DeleteEmptyFolders                        bool   `json:"deleteEmptyFolders,omitempty"`
	EnableMediaInfo                           bool   `json:"enableMediaInfo,omitempty"`
	ImportExtraFiles                          bool   `json:"importExtraFiles,omitempty"`
	SetPermissionsLinux                       bool   `json:"setPermissionsLinux,omitempty"`
	SkipFreeSpaceCheckWhenImporting           bool   `json:"skipFreeSpaceCheckWhenImporting,omitempty"`
	ID                                        int64  `json:"id"`
	MinimumFreeSpaceWhenImporting             int64  `json:"minimumFreeSpaceWhenImporting"` // 0 or empty not allowed
	RecycleBinCleanupDays                     int64  `json:"recycleBinCleanupDays,omitempty"`
	ScriptImportPath                          string `json:"scriptImportPath,omitempty"`
	ChmodFolder                               string `json:"chmodFolder,omitempty"`
	ChownGroup                                string `json:"chownGroup"` // empty string is valid
	DownloadPropersAndRepacks                 string `json:"downloadPropersAndRepacks,omitempty"`
	EpisodeTitleRequired                      string `json:"episodeTitleRequired,omitempty"`
	ExtraFileExtensions                       string `json:"extraFileExtensions,omitempty"`
	FileDate                                  string `json:"fileDate,omitempty"`
	RecycleBin                                string `json:"recycleBin"` // empty string is valid
	RescanAfterRefresh                        string `json:"rescanAfterRefresh,omitempty"`
}

// GetMediaManagement returns the mediaManagement.
func (s *Sonarr) GetMediaManagement() (*MediaManagement, error) {
	return s.GetMediaManagementContext(context.Background())
}

// GetMediaManagement returns the Media Management.
func (s *Sonarr) GetMediaManagementContext(ctx context.Context) (*MediaManagement, error) {
	var output MediaManagement

	req := starr.Request{URI: bpMediaManagement}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateMediaManagement updates the Media Management.
func (s *Sonarr) UpdateMediaManagement(mMgt *MediaManagement) (*MediaManagement, error) {
	return s.UpdateMediaManagementContext(context.Background(), mMgt)
}

// UpdateMediaManagementContext updates the Media Management.
func (s *Sonarr) UpdateMediaManagementContext(ctx context.Context, mMgt *MediaManagement) (*MediaManagement, error) {
	var output MediaManagement

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(mMgt); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpMediaManagement, err)
	}

	req := starr.Request{URI: bpMediaManagement, Body: &body}
	if err := s.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}
