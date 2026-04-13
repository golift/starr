package lidarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"

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
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetMetadataProfile returns a single metadata profile.
func (l *Lidarr) GetMetadataProfile(profileID int64) (*MetadataProfile, error) {
	return l.GetMetadataProfileContext(context.Background(), profileID)
}

// GetMetadataProfileContext returns a single metadata profile.
func (l *Lidarr) GetMetadataProfileContext(ctx context.Context, profileID int64) (*MetadataProfile, error) {
	var output MetadataProfile

	req := starr.Request{URI: path.Join(bpMetadataProfile, starr.Str(profileID))}
	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// AddMetadataProfile creates a metadata profile.
func (l *Lidarr) AddMetadataProfile(profile *MetadataProfile) (*MetadataProfile, error) {
	return l.AddMetadataProfileContext(context.Background(), profile)
}

// AddMetadataProfileContext creates a metadata profile.
func (l *Lidarr) AddMetadataProfileContext(ctx context.Context, profile *MetadataProfile) (*MetadataProfile, error) {
	var output MetadataProfile

	profile.ID = 0

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(profile); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpMetadataProfile, err)
	}

	req := starr.Request{URI: bpMetadataProfile, Body: &body}
	if err := l.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateMetadataProfile updates a metadata profile.
func (l *Lidarr) UpdateMetadataProfile(profile *MetadataProfile) (*MetadataProfile, error) {
	return l.UpdateMetadataProfileContext(context.Background(), profile)
}

// UpdateMetadataProfileContext updates a metadata profile.
func (l *Lidarr) UpdateMetadataProfileContext(ctx context.Context, profile *MetadataProfile) (*MetadataProfile, error) {
	var output MetadataProfile

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(profile); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpMetadataProfile, err)
	}

	req := starr.Request{URI: path.Join(bpMetadataProfile, starr.Str(profile.ID)), Body: &body}
	if err := l.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteMetadataProfile deletes a metadata profile.
func (l *Lidarr) DeleteMetadataProfile(profileID int64) error {
	return l.DeleteMetadataProfileContext(context.Background(), profileID)
}

// DeleteMetadataProfileContext deletes a metadata profile.
func (l *Lidarr) DeleteMetadataProfileContext(ctx context.Context, profileID int64) error {
	req := starr.Request{URI: path.Join(bpMetadataProfile, starr.Str(profileID))}
	if err := l.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
