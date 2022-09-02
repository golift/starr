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

// QualityProfile is the /api/v3/qualityprofile endpoint.
type QualityProfile struct {
	UpgradeAllowed bool             `json:"upgradeAllowed"`
	ID             int64            `json:"id,omitempty"`
	Cutoff         int64            `json:"cutoff"`
	Name           string           `json:"name"`
	Qualities      []*starr.Quality `json:"items"`
}

// Define Base Path for Quality Profile calls.
const bpQualityProfile = APIver + "/qualityProfile"

// GetQualityProfiles returns all configured quality profiles.
func (s *Sonarr) GetQualityProfiles() ([]*QualityProfile, error) {
	return s.GetQualityProfilesContext(context.Background())
}

func (s *Sonarr) GetQualityProfilesContext(ctx context.Context) ([]*QualityProfile, error) {
	var output []*QualityProfile

	if err := s.GetInto(ctx, bpQualityProfile, nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", bpQualityProfile, err)
	}

	return output, nil
}

// GetQualityProfile returns a single quality profile.
func (s *Sonarr) GetQualityProfile(profileID int) (*QualityProfile, error) {
	return s.GetQualityProfileContext(context.Background(), profileID)
}

func (s *Sonarr) GetQualityProfileContext(ctx context.Context, profileID int) (*QualityProfile, error) {
	var output *QualityProfile

	uri := path.Join(bpQualityProfile, strconv.Itoa(profileID))
	if err := s.GetInto(ctx, uri, nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", bpQualityProfile, err)
	}

	return output, nil
}

// AddQualityProfile creates a quality profile.
func (s *Sonarr) AddQualityProfile(profile *QualityProfile) (*QualityProfile, error) {
	return s.AddQualityProfileContext(context.Background(), profile)
}

func (s *Sonarr) AddQualityProfileContext(ctx context.Context, profile *QualityProfile) (*QualityProfile, error) {
	var output QualityProfile

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(profile); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpQualityProfile, err)
	}

	if err := s.PostInto(ctx, bpQualityProfile, nil, &body, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", bpQualityProfile, err)
	}

	return &output, nil
}

// UpdateQualityProfile updates the quality profile.
func (s *Sonarr) UpdateQualityProfile(profile *QualityProfile) (*QualityProfile, error) {
	return s.UpdateQualityProfileContext(context.Background(), profile)
}

func (s *Sonarr) UpdateQualityProfileContext(ctx context.Context, profile *QualityProfile) (*QualityProfile, error) {
	var output QualityProfile

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(profile); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpQualityProfile, err)
	}

	uri := path.Join(bpQualityProfile, strconv.Itoa(int(profile.ID)))
	if err := s.PutInto(ctx, uri, nil, &body, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", bpQualityProfile, err)
	}

	return &output, nil
}

// DeleteQualityProfile removes a single quality profile.
func (s *Sonarr) DeleteQualityProfile(profileID int) error {
	return s.DeleteQualityProfileContext(context.Background(), profileID)
}

func (s *Sonarr) DeleteQualityProfileContext(ctx context.Context, profileID int) error {
	var output interface{}

	uri := path.Join(bpQualityProfile, strconv.Itoa(profileID))
	if err := s.DeleteInto(ctx, uri, nil, &output); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", bpQualityProfile, err)
	}

	return nil
}
