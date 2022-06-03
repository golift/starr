package sonarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

type MediaManagement struct {
	AutoUnmonitorPreviouslyDownloadedEpisodes bool   `json:"autoUnmonitorPreviouslyDownloadedEpisodes,omitempty"`
	CopyUsingHardlinks                        bool   `json:"copyUsingHardlinks,omitempty"`
	CreateEmptySeriesFolders                  bool   `json:"createEmptySeriesFolders,omitempty"`
	DeleteEmptyFolders                        bool   `json:"deleteEmptyFolders,omitempty"`
	EnableMediaInfo                           bool   `json:"enableMediaInfo,omitempty"`
	ImportExtraFiles                          bool   `json:"importExtraFiles,omitempty"`
	SetPermissionsLinux                       bool   `json:"setPermissionsLinux,omitempty"`
	SkipFreeSpaceCheckWhenImporting           bool   `json:"skipFreeSpaceCheckWhenImporting,omitempty"`
	ID                                        int64  `json:"id,omitempty"`
	MinimumFreeSpaceWhenImporting             int64  `json:"minimumFreeSpaceWhenImporting,omitempty"`
	RecycleBinCleanupDays                     int64  `json:"recycleBinCleanupDays,omitempty"`
	ChmodFolder                               string `json:"chmodFolder,omitempty"`
	ChownGroup                                string `json:"chownGroup,omitempty"`
	DownloadPropersAndRepacks                 string `json:"downloadPropersAndRepacks,omitempty"`
	EpisodeTitleRequired                      string `json:"episodeTitleRequired,omitempty"`
	ExtraFileExtensions                       string `json:"extraFileExtensions,omitempty"`
	FileDate                                  string `json:"fileDate,omitempty"`
	RecycleBin                                string `json:"recycleBin,omitempty"`
	RescanAfterRefresh                        string `json:"rescanAfterRefresh,omitempty"`
}

// Define Base Path for MediaManagement calls.
const bpMediaManagement = APIver + "/config/mediaManagement"

// GetMediaManagement returns the mediaManagement.
func (s *Sonarr) GetMediaManagement() (*MediaManagement, error) {
	return s.GetMediaManagementContext(context.Background())
}

func (s *Sonarr) GetMediaManagementContext(ctx context.Context) (*MediaManagement, error) {
	var output *MediaManagement

	if _, err := s.GetInto(ctx, bpMediaManagement, nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(mediaManagement): %w", err)
	}

	return output, nil
}

// UpdateMediaManagement updates the mediaManagement.
func (s *Sonarr) UpdateMediaManagement(mMgt *MediaManagement) (*MediaManagement, error) {
	return s.UpdateMediaManagementContext(context.Background(), mMgt)
}

func (s *Sonarr) UpdateMediaManagementContext(ctx context.Context, mMgt *MediaManagement) (*MediaManagement, error) {
	var output MediaManagement

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(mMgt); err != nil {
		return nil, fmt.Errorf("json.Marshal(mediaManagement): %w", err)
	}

	if _, err := s.PutInto(ctx, bpMediaManagement, nil, &body, &output); err != nil {
		return nil, fmt.Errorf("api.Put(mediaManagement): %w", err)
	}

	return &output, nil
}
