package readarr

import (
	"context"
	"fmt"

	"golift.io/starr"
	"golift.io/starr/starrshared"
)

const bpUpdate = APIver + "/update"

// UpdateChanges is the change log embedded in Update.
type UpdateChanges = starrshared.UpdateChanges

// Update is one available or installed update from /api/v1/update.
type Update = starrshared.Update

// GetUpdates returns available application updates.
func (r *Readarr) GetUpdates() ([]*Update, error) {
	return r.GetUpdatesContext(context.Background())
}

// GetUpdatesContext returns available application updates.
func (r *Readarr) GetUpdatesContext(ctx context.Context) ([]*Update, error) {
	var output []*Update

	req := starr.Request{URI: bpUpdate}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}
