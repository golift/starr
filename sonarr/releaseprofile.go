package sonarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"
	"strconv"

	"golift.io/starr"
)

// ReleaseProfile defines a release profile's data from Sonarr.
type ReleaseProfile struct {
	Enabled         bool              `json:"enabled"`
	IncPrefOnRename bool              `json:"includePreferredWhenRenaming"`
	ID              int64             `json:"id,omitempty"`
	IndexerID       int64             `json:"indexerId"`
	Name            string            `json:"name"`
	Tags            []int             `json:"tags"`
	Required        []string          `json:"required"`
	Ignored         []string          `json:"ignored"`
	Preferred       []*starr.KeyValue `json:"preferred"`
}

// Define Base Path for Release Profile calls.
const bpReleaseProfile = APIver + "/releaseProfile"

// GetReleaseProfiles returns all configured release profiles.
func (s *Sonarr) GetReleaseProfiles() ([]*ReleaseProfile, error) {
	return s.GetReleaseProfilesContext(context.Background())
}

func (s *Sonarr) GetReleaseProfilesContext(ctx context.Context) ([]*ReleaseProfile, error) {
	var output []*ReleaseProfile

	if err := s.GetInto(ctx, bpReleaseProfile, nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", bpReleaseProfile, err)
	}

	return output, nil
}

// GetReleaseProfile returns a single release profile.
func (s *Sonarr) GetReleaseProfile(profileID int) (*ReleaseProfile, error) {
	return s.GetReleaseProfileContext(context.Background(), profileID)
}

func (s *Sonarr) GetReleaseProfileContext(ctx context.Context, profileID int) (*ReleaseProfile, error) {
	var output *ReleaseProfile

	uri := path.Join(bpReleaseProfile, strconv.Itoa(profileID))
	if err := s.GetInto(ctx, uri, nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", bpReleaseProfile, err)
	}

	return output, nil
}

// AddReleaseProfile creates a release profile.
func (s *Sonarr) AddReleaseProfile(profile *ReleaseProfile) (*ReleaseProfile, error) {
	return s.AddReleaseProfileContext(context.Background(), profile)
}

func (s *Sonarr) AddReleaseProfileContext(ctx context.Context, profile *ReleaseProfile) (*ReleaseProfile, error) {
	var output ReleaseProfile

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(profile); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpReleaseProfile, err)
	}

	if err := s.PostInto(ctx, bpReleaseProfile, nil, &body, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", bpReleaseProfile, err)
	}

	return &output, nil
}

// UpdateReleaseProfile updates the release profile.
func (s *Sonarr) UpdateReleaseProfile(profile *ReleaseProfile) (*ReleaseProfile, error) {
	return s.UpdateReleaseProfileContext(context.Background(), profile)
}

func (s *Sonarr) UpdateReleaseProfileContext(ctx context.Context, profile *ReleaseProfile) (*ReleaseProfile, error) {
	var output ReleaseProfile

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(profile); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpReleaseProfile, err)
	}

	uri := path.Join(bpReleaseProfile, strconv.Itoa(int(profile.ID)))
	if err := s.PutInto(ctx, uri, nil, &body, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", bpReleaseProfile, err)
	}

	return &output, nil
}

// DeleteReleaseProfile removes a single release profile.
func (s *Sonarr) DeleteReleaseProfile(profileID int) error {
	return s.DeleteReleaseProfileContext(context.Background(), profileID)
}

func (s *Sonarr) DeleteReleaseProfileContext(ctx context.Context, profileID int) error {
	req := &starr.Request{URI: path.Join(bpReleaseProfile, fmt.Sprint(profileID))}
	if err := s.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", req.URI, err)
	}

	return nil
}
