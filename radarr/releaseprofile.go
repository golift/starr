package radarr

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

// ReleaseProfile defines a release profile's data from Radarr. v4 only.
type ReleaseProfile struct {
	Name      string   `json:"name"`
	Enabled   bool     `json:"enabled"`
	Required  []string `json:"required"`
	Ignored   []string `json:"ignored"`
	IndexerID int64    `json:"indexerId"`
	Tags      []int    `json:"tags"`
	ID        int64    `json:"id,omitempty"`
}

// GetReleaseProfiles returns all configured release profiles.
func (r *Radarr) GetReleaseProfiles() ([]*ReleaseProfile, error) {
	return r.GetReleaseProfilesContext(context.Background())
}

// GetReleaseProfilesContext returns all configured release profiles.
func (r *Radarr) GetReleaseProfilesContext(ctx context.Context) ([]*ReleaseProfile, error) {
	var output []*ReleaseProfile

	req := starr.Request{URI: bpReleaseProfile}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetReleaseProfile returns a single release profile.
func (r *Radarr) GetReleaseProfile(profileID int64) (*ReleaseProfile, error) {
	return r.GetReleaseProfileContext(context.Background(), profileID)
}

// GetReleaseProfileContext returns a single release profile.
func (r *Radarr) GetReleaseProfileContext(ctx context.Context, profileID int64) (*ReleaseProfile, error) {
	var output ReleaseProfile

	req := starr.Request{URI: path.Join(bpReleaseProfile, starr.Itoa(profileID))}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// AddReleaseProfile creates a release profile.
func (r *Radarr) AddReleaseProfile(profile *ReleaseProfile) (*ReleaseProfile, error) {
	return r.AddReleaseProfileContext(context.Background(), profile)
}

// AddReleaseProfileContext creates a release profile.
func (r *Radarr) AddReleaseProfileContext(ctx context.Context, profile *ReleaseProfile) (*ReleaseProfile, error) {
	var output ReleaseProfile

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(profile); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpReleaseProfile, err)
	}

	req := starr.Request{URI: bpReleaseProfile, Body: &body}
	if err := r.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateReleaseProfile updates the release profile.
func (r *Radarr) UpdateReleaseProfile(profile *ReleaseProfile) (*ReleaseProfile, error) {
	return r.UpdateReleaseProfileContext(context.Background(), profile)
}

// UpdateReleaseProfileContext updates the release profile.
func (r *Radarr) UpdateReleaseProfileContext(ctx context.Context, profile *ReleaseProfile) (*ReleaseProfile, error) {
	var output ReleaseProfile

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(profile); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpReleaseProfile, err)
	}

	req := starr.Request{URI: path.Join(bpReleaseProfile, starr.Itoa(profile.ID)), Body: &body}
	if err := r.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteReleaseProfile removes a single release profile.
func (r *Radarr) DeleteReleaseProfile(profileID int64) error {
	return r.DeleteReleaseProfileContext(context.Background(), profileID)
}

// DeleteReleaseProfileContext removes a single release profile.
func (r *Radarr) DeleteReleaseProfileContext(ctx context.Context, profileID int64) error {
	req := starr.Request{URI: path.Join(bpReleaseProfile, starr.Itoa(profileID))}
	if err := r.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
