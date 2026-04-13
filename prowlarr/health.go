package prowlarr

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
func (p *Prowlarr) GetHealth() ([]*Health, error) {
	return p.GetHealthContext(context.Background())
}

// GetHealthContext returns current health check messages.
func (p *Prowlarr) GetHealthContext(ctx context.Context) ([]*Health, error) {
	var output []*Health

	req := starr.Request{URI: bpHealth}
	if err := p.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}
