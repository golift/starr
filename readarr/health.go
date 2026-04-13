package readarr

import (
	"context"
	"fmt"

	"golift.io/starr"
	"golift.io/starr/starrshared"
)

const bpHealth = APIver + "/health"

// Health is the /api/v1/health resource.
type Health = starrshared.Health

// GetHealth returns current health check messages.
func (r *Readarr) GetHealth() ([]*Health, error) {
	return r.GetHealthContext(context.Background())
}

// GetHealthContext returns current health check messages.
func (r *Readarr) GetHealthContext(ctx context.Context) ([]*Health, error) {
	var output []*Health

	req := starr.Request{URI: bpHealth}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}
