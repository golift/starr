package sonarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"

	"golift.io/starr"
)

// Define Base Path for Quality Profile calls.
const bpQualityProfile = APIver + "/qualityProfile"

// QualityProfile is the /api/v3/qualityprofile endpoint.
type QualityProfile struct {
	UpgradeAllowed    bool                `json:"upgradeAllowed"`
	ID                int64               `json:"id"`
	Cutoff            int64               `json:"cutoff"`
	Name              string              `json:"name"`
	Qualities         []*starr.Quality    `json:"items"`
	MinFormatScore    int64               `json:"minFormatScore"`        // v4 only.
	CutoffFormatScore int64               `json:"cutoffFormatScore"`     // v4 only.
	FormatItems       []*starr.FormatItem `json:"formatItems,omitempty"` // v4 only.
	Language          *starr.Value        `json:"language,omitempty"`    // v4 only.
}

// GetQualityProfiles returns all configured quality profiles.
func (s *Sonarr) GetQualityProfiles() ([]*QualityProfile, error) {
	return s.GetQualityProfilesContext(context.Background())
}

// GetQualityProfilesContext returns all configured quality profiles.
func (s *Sonarr) GetQualityProfilesContext(ctx context.Context) ([]*QualityProfile, error) {
	var output []*QualityProfile

	req := starr.Request{URI: bpQualityProfile}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetQualityProfile returns a single quality profile.
func (s *Sonarr) GetQualityProfile(profileID int64) (*QualityProfile, error) {
	return s.GetQualityProfileContext(context.Background(), profileID)
}

// GetQualityProfileContext returns a single quality profile.
func (s *Sonarr) GetQualityProfileContext(ctx context.Context, profileID int64) (*QualityProfile, error) {
	var output QualityProfile

	req := starr.Request{URI: path.Join(bpQualityProfile, starr.Str(profileID))}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// AddQualityProfile creates a quality profile.
func (s *Sonarr) AddQualityProfile(profile *QualityProfile) (*QualityProfile, error) {
	return s.AddQualityProfileContext(context.Background(), profile)
}

// AddQualityProfileContext creates a quality profile.
func (s *Sonarr) AddQualityProfileContext(ctx context.Context, profile *QualityProfile) (*QualityProfile, error) {
	var (
		output QualityProfile
		body   bytes.Buffer
	)

	profile.ID = 0
	if err := json.NewEncoder(&body).Encode(profile); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpQualityProfile, err)
	}

	req := starr.Request{URI: bpQualityProfile, Body: &body}
	if err := s.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateQualityProfile updates the quality profile.
func (s *Sonarr) UpdateQualityProfile(profile *QualityProfile) (*QualityProfile, error) {
	return s.UpdateQualityProfileContext(context.Background(), profile)
}

// UpdateQualityProfileContext updates the quality profile.
func (s *Sonarr) UpdateQualityProfileContext(ctx context.Context, profile *QualityProfile) (*QualityProfile, error) {
	var output QualityProfile

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(profile); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpQualityProfile, err)
	}

	req := starr.Request{URI: path.Join(bpQualityProfile, starr.Str(profile.ID)), Body: &body}
	if err := s.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteQualityProfile removes a single quality profile.
func (s *Sonarr) DeleteQualityProfile(profileID int64) error {
	return s.DeleteQualityProfileContext(context.Background(), profileID)
}

// DeleteQualityProfileContext removes a single quality profile.
func (s *Sonarr) DeleteQualityProfileContext(ctx context.Context, profileID int64) error {
	req := starr.Request{URI: path.Join(bpQualityProfile, starr.Str(profileID))}
	if err := s.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
