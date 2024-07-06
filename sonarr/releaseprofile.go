package sonarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"

	"golift.io/starr"
)

// Define Base Path for Release Profile calls.
const bpReleaseProfile = APIver + "/releaseProfile"

// ReleaseProfile defines a release profile's data from Sonarr.
type ReleaseProfile struct {
	Name            string            `json:"name"`
	Enabled         bool              `json:"enabled"`
	Required        []string          `json:"required"`
	Ignored         []string          `json:"ignored"`
	IndexerID       int64             `json:"indexerId"`
	Tags            []int             `json:"tags"`
	ID              int64             `json:"id,omitempty"`
	IncPrefOnRename *bool             `json:"includePreferredWhenRenaming,omitempty"` // V3 only, removed from v4.
	Preferred       []*starr.KeyValue `json:"preferred,omitempty"`                    // V3 only, removed from v4.
}

// GetReleaseProfiles returns all configured release profiles.
func (s *Sonarr) GetReleaseProfiles() ([]*ReleaseProfile, error) {
	return s.GetReleaseProfilesContext(context.Background())
}

// GetReleaseProfilesContext returns all configured release profiles.
func (s *Sonarr) GetReleaseProfilesContext(ctx context.Context) ([]*ReleaseProfile, error) {
	var output []*ReleaseProfile

	req := starr.Request{URI: bpReleaseProfile}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetReleaseProfile returns a single release profile.
func (s *Sonarr) GetReleaseProfile(profileID int64) (*ReleaseProfile, error) {
	return s.GetReleaseProfileContext(context.Background(), profileID)
}

// GetReleaseProfileContext returns a single release profile.
func (s *Sonarr) GetReleaseProfileContext(ctx context.Context, profileID int64) (*ReleaseProfile, error) {
	var output ReleaseProfile

	req := starr.Request{URI: path.Join(bpReleaseProfile, starr.Itoa(profileID))}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// AddReleaseProfile creates a release profile.
func (s *Sonarr) AddReleaseProfile(profile *ReleaseProfile) (*ReleaseProfile, error) {
	return s.AddReleaseProfileContext(context.Background(), profile)
}

// AddReleaseProfileContext creates a release profile.
func (s *Sonarr) AddReleaseProfileContext(ctx context.Context, profile *ReleaseProfile) (*ReleaseProfile, error) {
	var output ReleaseProfile

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(profile); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpReleaseProfile, err)
	}

	req := starr.Request{URI: bpReleaseProfile, Body: &body}
	if err := s.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateReleaseProfile updates the release profile.
func (s *Sonarr) UpdateReleaseProfile(profile *ReleaseProfile) (*ReleaseProfile, error) {
	return s.UpdateReleaseProfileContext(context.Background(), profile)
}

// UpdateReleaseProfileContext updates the release profile.
func (s *Sonarr) UpdateReleaseProfileContext(ctx context.Context, profile *ReleaseProfile) (*ReleaseProfile, error) {
	var output ReleaseProfile

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(profile); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpReleaseProfile, err)
	}

	req := starr.Request{URI: path.Join(bpReleaseProfile, starr.Itoa(profile.ID)), Body: &body}
	if err := s.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteReleaseProfile removes a single release profile.
func (s *Sonarr) DeleteReleaseProfile(profileID int64) error {
	return s.DeleteReleaseProfileContext(context.Background(), profileID)
}

// DeleteReleaseProfileContext removes a single release profile.
func (s *Sonarr) DeleteReleaseProfileContext(ctx context.Context, profileID int64) error {
	req := starr.Request{URI: path.Join(bpReleaseProfile, starr.Itoa(profileID))}
	if err := s.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
