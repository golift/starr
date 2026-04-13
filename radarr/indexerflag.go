package radarr

import (
	"context"
	"fmt"

	"golift.io/starr"
)

const bpIndexerFlag = APIver + "/indexerflag"

// IndexerFlag is the /api/v3/indexerflag resource.
type IndexerFlag struct {
	ID        int    `json:"id"`
	Name      string `json:"name,omitempty"`
	NameLower string `json:"nameLower,omitempty"`
}

// GetIndexerFlags returns all indexer flags.
func (r *Radarr) GetIndexerFlags() ([]*IndexerFlag, error) {
	return r.GetIndexerFlagsContext(context.Background())
}

// GetIndexerFlagsContext returns all indexer flags.
func (r *Radarr) GetIndexerFlagsContext(ctx context.Context) ([]*IndexerFlag, error) {
	var output []*IndexerFlag

	req := starr.Request{URI: bpIndexerFlag}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}
