package radarr

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
	WhitelistedHardcodedSubs string `json:"whitelistedHardcodedSubs"`
	ID                       int64  `json:"id"`
	MaximumSize              int64  `json:"maximumSize"`
	MinimumAge               int64  `json:"minimumAge"`
	Retention                int64  `json:"retention"`
	RssSyncInterval          int64  `json:"rssSyncInterval"`
	AvailabilityDelay        int    `json:"availabilityDelay"`
	PreferIndexerFlags       bool   `json:"preferIndexerFlags"`
	AllowHardcodedSubs       bool   `json:"allowHardcodedSubs"`
}

// GetIndexerConfig returns an Indexer Config.
func (r *Radarr) GetIndexerConfig() (*IndexerConfig, error) {
	return r.GetIndexerConfigContext(context.Background())
}

// GetIndexerConfigContext returns the indexer Config.
func (r *Radarr) GetIndexerConfigContext(ctx context.Context) (*IndexerConfig, error) {
	var output IndexerConfig

	req := starr.Request{URI: bpIndexerConfig}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateIndexerConfig update the single indexerConfig.
func (r *Radarr) UpdateIndexerConfig(indexerConfig *IndexerConfig) (*IndexerConfig, error) {
	return r.UpdateIndexerConfigContext(context.Background(), indexerConfig)
}

// UpdateIndexerConfigContext update the single indexerConfig.
func (r *Radarr) UpdateIndexerConfigContext(ctx context.Context, indexerConfig *IndexerConfig) (*IndexerConfig, error) {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(indexerConfig); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpIndexerConfig, err)
	}

	var output IndexerConfig

	req := starr.Request{URI: path.Join(bpIndexerConfig, fmt.Sprint(indexerConfig.ID)), Body: &body}
	if err := r.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}
