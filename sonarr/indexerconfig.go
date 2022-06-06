package sonarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"
	"strconv"
)

type IndexerConfig struct {
	ID              int64 `json:"id"`
	MaximumSize     int64 `json:"maximumSize"`
	MinimumAge      int64 `json:"minimumAge"`
	Retention       int64 `json:"retention"`
	RssSyncInterval int64 `json:"rssSyncInterval"`
}

const bpIndexerConfig = APIver + "/config/indexer"

// GetIndexerConfig returns the indexerConfig.
func (s *Sonarr) GetIndexerConfig() (*IndexerConfig, error) {
	return s.GetIndexerConfigContext(context.Background())
}

func (s *Sonarr) GetIndexerConfigContext(ctx context.Context) (*IndexerConfig, error) {
	var output *IndexerConfig

	if _, err := s.GetInto(ctx, bpIndexerConfig, nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(indexerConfig): %w", err)
	}

	return output, nil
}

// UpdateIndexerConfig update the single indexerConfig.
func (s *Sonarr) UpdateIndexerConfig(indexerConfig *IndexerConfig) (*IndexerConfig, error) {
	return s.UpdateIndexerConfigContext(context.Background(), indexerConfig)
}

func (s *Sonarr) UpdateIndexerConfigContext(ctx context.Context, indexerConfig *IndexerConfig) (*IndexerConfig, error) {
	var output IndexerConfig

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(indexerConfig); err != nil {
		return nil, fmt.Errorf("json.Marshal(indexerConfig): %w", err)
	}

	uri := path.Join(bpIndexerConfig, strconv.Itoa(int(indexerConfig.ID)))
	if _, err := s.PutInto(ctx, uri, nil, &body, &output); err != nil {
		return nil, fmt.Errorf("api.Put(indexerConfig): %w", err)
	}

	return &output, nil
}
