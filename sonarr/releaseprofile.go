package sonarr

import (
	"encoding/json"
	"fmt"
	"strconv"

	"golift.io/starr"
)

// GetReleaseProfiles returns all configured release profiles.
func (s *Sonarr) GetReleaseProfiles() ([]*ReleaseProfile, error) {
	var profiles []*ReleaseProfile

	err := s.GetInto("v3/releaseProfile", nil, &profiles)
	if err != nil {
		return nil, fmt.Errorf("api.Get(releaseProfile): %w", err)
	}

	return profiles, nil
}

// AddReleaseProfile updates a release profile in place.
func (s *Sonarr) AddReleaseProfile(profile *ReleaseProfile) (int64, error) {
	post, err := json.Marshal(profile)
	if err != nil {
		return 0, fmt.Errorf("json.Marshal(profile): %w", err)
	}

	var output ReleaseProfile

	err = s.PostInto("v3/releaseProfile", nil, post, &output)
	if err != nil {
		return 0, fmt.Errorf("api.Post(releaseProfile): %w", err)
	}

	return output.ID, nil
}

// UpdateReleaseProfile updates a release profile in place.
func (s *Sonarr) UpdateReleaseProfile(profile *ReleaseProfile) error {
	put, err := json.Marshal(profile)
	if err != nil {
		return fmt.Errorf("json.Marshal(profile): %w", err)
	}

	_, err = s.Put("v3/releaseProfile/"+strconv.FormatInt(profile.ID, starr.Base10), nil, put)
	if err != nil {
		return fmt.Errorf("api.Put(releaseProfile): %w", err)
	}

	return nil
}
