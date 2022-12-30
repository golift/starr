package lidarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"

	"golift.io/starr"
)

const bpIndexerConfig = APIver + "/config/indexer"

// IndexerConfig represents the /config/indexer endpoint.
type IndexerConfig struct {
	ID              int64 `json:"id"`
	MaximumSize     int64 `json:"maximumSize"`
	MinimumAge      int64 `json:"minimumAge"`
	Retention       int64 `json:"retention"`
	RssSyncInterval int64 `json:"rssSyncInterval"`
}

// GetIndexerConfig returns an Indexer Config.
func (l *Lidarr) GetIndexerConfig() (*IndexerConfig, error) {
	return l.GetIndexerConfigContext(context.Background())
}

// GetIndexerConfigContext returns the indexer Config.
func (l *Lidarr) GetIndexerConfigContext(ctx context.Context) (*IndexerConfig, error) {
	var output IndexerConfig

	req := starr.Request{URI: bpIndexerConfig}
	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateIndexerConfig update the single indexerConfig.
func (l *Lidarr) UpdateIndexerConfig(indexerConfig *IndexerConfig) (*IndexerConfig, error) {
	return l.UpdateIndexerConfigContext(context.Background(), indexerConfig)
}

// UpdateIndexerConfigContext update the single indexerConfig.
func (l *Lidarr) UpdateIndexerConfigContext(ctx context.Context, indexerConfig *IndexerConfig) (*IndexerConfig, error) {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(indexerConfig); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpIndexerConfig, err)
	}

	var output IndexerConfig

	req := starr.Request{URI: path.Join(bpIndexerConfig, fmt.Sprint(indexerConfig.ID)), Body: &body}
	if err := l.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}
