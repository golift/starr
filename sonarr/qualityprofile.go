package sonarr

import (
	"encoding/json"
	"fmt"
	"strconv"

	"golift.io/starr"
)

// GetQualityProfiles returns all configured quality profiles.
func (s *Sonarr) GetQualityProfiles() ([]*QualityProfile, error) {
	var profiles []*QualityProfile

	err := s.GetInto("v3/qualityprofile", nil, &profiles)
	if err != nil {
		return nil, fmt.Errorf("api.Get(qualityprofile): %w", err)
	}

	return profiles, nil
}

// AddQualityProfile updates a quality profile in place.
func (s *Sonarr) AddQualityProfile(profile *QualityProfile) (int64, error) {
	post, err := json.Marshal(profile)
	if err != nil {
		return 0, fmt.Errorf("json.Marshal(profile): %w", err)
	}

	var output QualityProfile

	err = s.PostInto("v3/qualityProfile", nil, post, &output)
	if err != nil {
		return 0, fmt.Errorf("api.Post(qualityProfile): %w", err)
	}

	return output.ID, nil
}

// UpdateQualityProfile updates a quality profile in place.
func (s *Sonarr) UpdateQualityProfile(profile *QualityProfile) error {
	put, err := json.Marshal(profile)
	if err != nil {
		return fmt.Errorf("json.Marshal(profile): %w", err)
	}

	_, err = s.Put("v3/qualityProfile/"+strconv.FormatInt(profile.ID, starr.Base10), nil, put)
	if err != nil {
		return fmt.Errorf("api.Put(qualityProfile): %w", err)
	}

	return nil
}
