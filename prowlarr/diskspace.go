package prowlarr

import (
	"context"
	"fmt"

	"golift.io/starr"
	"golift.io/starr/starrshared"
)

const bpDiskSpace = APIver + "/diskspace"

// DiskSpace is the /api/v1/diskspace resource.
type DiskSpace = starrshared.DiskSpace

// GetDiskSpace returns disk space information for Prowlarr paths.
func (p *Prowlarr) GetDiskSpace() ([]*DiskSpace, error) {
	return p.GetDiskSpaceContext(context.Background())
}

// GetDiskSpaceContext returns disk space information for Prowlarr paths.
func (p *Prowlarr) GetDiskSpaceContext(ctx context.Context) ([]*DiskSpace, error) {
	var output []*DiskSpace

	req := starr.Request{URI: bpDiskSpace}
	if err := p.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}
