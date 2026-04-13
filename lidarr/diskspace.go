package lidarr

import (
	"context"
	"fmt"

	"golift.io/starr"
	"golift.io/starr/starrshared"
)

const bpDiskSpace = APIver + "/diskspace"

// DiskSpace is the /api/v1/diskspace resource.
type DiskSpace = starrshared.DiskSpace

// GetDiskSpace returns disk space information for Lidarr paths.
func (l *Lidarr) GetDiskSpace() ([]*DiskSpace, error) {
	return l.GetDiskSpaceContext(context.Background())
}

// GetDiskSpaceContext returns disk space information for Lidarr paths.
func (l *Lidarr) GetDiskSpaceContext(ctx context.Context) ([]*DiskSpace, error) {
	var output []*DiskSpace

	req := starr.Request{URI: bpDiskSpace}
	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}
