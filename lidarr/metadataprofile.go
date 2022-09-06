package lidarr

import (
	"context"
	"fmt"

	"golift.io/starr"
)

const bpMetadataProfile = APIver + "/metadataprofile"

// MetadataProfile is the /api/v1/metadataprofile endpoint.
type MetadataProfile struct {
	Name                string           `json:"name"`
	ID                  int64            `json:"id"`
	PrimaryAlbumTypes   []*AlbumType     `json:"primaryAlbumTypes"`
	SecondaryAlbumTypes []*AlbumType     `json:"secondaryAlbumTypes"`
	ReleaseStatuses     []*ReleaseStatus `json:"releaseStatuses"`
}

// AlbumType is part of MetadataProfile.
type AlbumType struct {
	AlbumType *starr.Value `json:"albumType"`
	Allowed   bool         `json:"allowed"`
}

// ReleaseStatus is part of MetadataProfile.
type ReleaseStatus struct {
	ReleaseStatus *starr.Value `json:"releaseStatus"`
	Allowed       bool         `json:"allowed"`
}

// GetMetadataProfiles returns the metadata profiles.
func (l *Lidarr) GetMetadataProfiles() ([]*MetadataProfile, error) {
	return l.GetMetadataProfilesContext(context.Background())
}

// GetMetadataProfilesContext returns the metadata profiles.
func (l *Lidarr) GetMetadataProfilesContext(ctx context.Context) ([]*MetadataProfile, error) {
	var output []*MetadataProfile

	req := starr.Request{URI: bpMetadataProfile}
	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", req, err)
	}

	return output, nil
}
