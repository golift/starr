package sonarr

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
func (s *Sonarr) GetIndexerConfig() (*IndexerConfig, error) {
	return s.GetIndexerConfigContext(context.Background())
}

// GetIndexerConfigContext returns the indexer Config.
func (s *Sonarr) GetIndexerConfigContext(ctx context.Context) (*IndexerConfig, error) {
	var output IndexerConfig

	req := starr.Request{URI: bpIndexerConfig}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateIndexerConfig update the single indexerConfig.
func (s *Sonarr) UpdateIndexerConfig(indexerConfig *IndexerConfig) (*IndexerConfig, error) {
	return s.UpdateIndexerConfigContext(context.Background(), indexerConfig)
}

// UpdateIndexerConfigContext update the single indexerConfig.
func (s *Sonarr) UpdateIndexerConfigContext(ctx context.Context, indexerConfig *IndexerConfig) (*IndexerConfig, error) {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(indexerConfig); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpIndexerConfig, err)
	}

	var output IndexerConfig

	req := starr.Request{URI: path.Join(bpIndexerConfig, starr.Str(indexerConfig.ID)), Body: &body}
	if err := s.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}
