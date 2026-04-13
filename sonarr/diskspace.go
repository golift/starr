package sonarr

import (
	"context"
	"fmt"

	"golift.io/starr"
	"golift.io/starr/starrshared"
)

const bpDiskSpace = APIver + "/diskspace"

// DiskSpace is the /api/v3/diskspace resource.
type DiskSpace = starrshared.DiskSpace

// GetDiskSpace returns disk space information for Sonarr paths.
func (s *Sonarr) GetDiskSpace() ([]*DiskSpace, error) {
	return s.GetDiskSpaceContext(context.Background())
}

// GetDiskSpaceContext returns disk space information for Sonarr paths.
func (s *Sonarr) GetDiskSpaceContext(ctx context.Context) ([]*DiskSpace, error) {
	var output []*DiskSpace

	req := starr.Request{URI: bpDiskSpace}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}
