package sonarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"golift.io/starr"
)

// GetReleaseProfiles returns all configured release profiles.
func (s *Sonarr) GetReleaseProfiles() ([]*ReleaseProfile, error) {
	return s.GetReleaseProfilesContext(context.Background())
}

func (s *Sonarr) GetReleaseProfilesContext(ctx context.Context) ([]*ReleaseProfile, error) {
	var profiles []*ReleaseProfile

	err := s.GetInto(ctx, "v3/releaseProfile", nil, &profiles)
	if err != nil {
		return nil, fmt.Errorf("api.Get(releaseProfile): %w", err)
	}

	return profiles, nil
}

// AddReleaseProfile updates a release profile in place.
func (s *Sonarr) AddReleaseProfile(profile *ReleaseProfile) (int64, error) {
	return s.AddReleaseProfileContext(context.Background(), profile)
}

func (s *Sonarr) AddReleaseProfileContext(ctx context.Context, profile *ReleaseProfile) (int64, error) {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(profile); err != nil {
		return 0, fmt.Errorf("json.Marshal(profile): %w", err)
	}

	var output ReleaseProfile
	if err := s.PostInto(ctx, "v3/releaseProfile", nil, &body, &output); err != nil {
		return 0, fmt.Errorf("api.Post(releaseProfile): %w", err)
	}

	return output.ID, nil
}

// UpdateReleaseProfile updates a release profile in place.
func (s *Sonarr) UpdateReleaseProfile(profile *ReleaseProfile) error {
	return s.UpdateReleaseProfileContext(context.Background(), profile)
}

func (s *Sonarr) UpdateReleaseProfileContext(ctx context.Context, profile *ReleaseProfile) error {
	put, err := json.Marshal(profile)
	if err != nil {
		return fmt.Errorf("json.Marshal(profile): %w", err)
	}

	_, err = s.Put(ctx, "v3/releaseProfile/"+strconv.FormatInt(profile.ID, starr.Base10), nil, put)
	if err != nil {
		return fmt.Errorf("api.Put(releaseProfile): %w", err)
	}

	return nil
}
