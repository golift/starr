package lidarr

import (
	"context"
	"fmt"

	"golift.io/starr"
)

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
	var profiles []*MetadataProfile

	err := l.GetInto(ctx, "v1/metadataprofile", nil, &profiles)
	if err != nil {
		return nil, fmt.Errorf("api.Get(metadataprofile): %w", err)
	}

	return profiles, nil
}
