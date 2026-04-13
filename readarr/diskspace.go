package readarr

import (
	"context"
	"fmt"

	"golift.io/starr"
	"golift.io/starr/starrshared"
)

const bpDiskSpace = APIver + "/diskspace"

// DiskSpace is the /api/v1/diskspace resource.
type DiskSpace = starrshared.DiskSpace

// GetDiskSpace returns disk space information for Readarr paths.
func (r *Readarr) GetDiskSpace() ([]*DiskSpace, error) {
	return r.GetDiskSpaceContext(context.Background())
}

// GetDiskSpaceContext returns disk space information for Readarr paths.
func (r *Readarr) GetDiskSpaceContext(ctx context.Context) ([]*DiskSpace, error) {
	var output []*DiskSpace

	req := starr.Request{URI: bpDiskSpace}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}
