package sonarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"

	"golift.io/starr"
)

// Define Base Path for downloadClientConfig calls.
const bpDownloadClientConfig = APIver + "/config/downloadClient"

// DownloadClientConfig is the /api/v3/config/downloadClientConfig endpoint.
type DownloadClientConfig struct {
	EnableCompletedDownloadHandling bool   `json:"enableCompletedDownloadHandling"`
	AutoRedownloadFailed            bool   `json:"autoRedownloadFailed"`
	ID                              int64  `json:"id"`
	DownloadClientWorkingFolders    string `json:"downloadClientWorkingFolders"`
}

// GetDownloadClientConfig returns the downloadClientConfig.
func (s *Sonarr) GetDownloadClientConfig() (*DownloadClientConfig, error) {
	return s.GetDownloadClientConfigContext(context.Background())
}

// GetDownloadClientConfig returns the downloadClientConfig.
func (s *Sonarr) GetDownloadClientConfigContext(ctx context.Context) (*DownloadClientConfig, error) {
	var output DownloadClientConfig

	req := starr.Request{URI: bpDownloadClientConfig}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateDownloadClientConfig update the single downloadClientConfig.
func (s *Sonarr) UpdateDownloadClientConfig(downloadClientConfig *DownloadClientConfig) (*DownloadClientConfig, error) {
	return s.UpdateDownloadClientConfigContext(context.Background(), downloadClientConfig)
}

// UpdateDownloadClientConfig update the single downloadClientConfig.
func (s *Sonarr) UpdateDownloadClientConfigContext(ctx context.Context,
	config *DownloadClientConfig,
) (*DownloadClientConfig, error) {
	var output DownloadClientConfig

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(config); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpDownloadClientConfig, err)
	}

	req := starr.Request{URI: path.Join(bpDownloadClientConfig, fmt.Sprint(int(config.ID))), Body: &body}
	if err := s.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}
