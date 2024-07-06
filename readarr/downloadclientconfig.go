package readarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"

	"golift.io/starr"
)

// Define Base Path for download client config calls.
const bpDownloadClientConfig = APIver + "/config/downloadClient"

// DownloadClientConfig is the /api/v1/config/downloadClientConfig endpoint.
type DownloadClientConfig struct {
	EnableCompletedDownloadHandling bool   `json:"enableCompletedDownloadHandling"`
	AutoRedownloadFailed            bool   `json:"autoRedownloadFailed"`
	ID                              int64  `json:"id"`
	DownloadClientWorkingFolders    string `json:"downloadClientWorkingFolders"`
	RemoveCompletedDownloads        bool   `json:"removeCompletedDownloads"`
	RemoveFailedDownloads           bool   `json:"removeFailedDownloads"`
}

// GetDownloadClientConfig returns the download client config.
func (r *Readarr) GetDownloadClientConfig() (*DownloadClientConfig, error) {
	return r.GetDownloadClientConfigContext(context.Background())
}

// GetDownloadClientConfig returns the download client config.
func (r *Readarr) GetDownloadClientConfigContext(ctx context.Context) (*DownloadClientConfig, error) {
	var output DownloadClientConfig

	req := starr.Request{URI: bpDownloadClientConfig}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateDownloadClientConfig update the single download client config.
func (r *Readarr) UpdateDownloadClientConfig(downloadConfig *DownloadClientConfig) (*DownloadClientConfig, error) {
	return r.UpdateDownloadClientConfigContext(context.Background(), downloadConfig)
}

// UpdateDownloadClientConfig update the single download client config.
func (r *Readarr) UpdateDownloadClientConfigContext(ctx context.Context,
	config *DownloadClientConfig,
) (*DownloadClientConfig, error) {
	var output DownloadClientConfig

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(config); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpDownloadClientConfig, err)
	}

	req := starr.Request{URI: path.Join(bpDownloadClientConfig, starr.Str(config.ID)), Body: &body}
	if err := r.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}
