package sonarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"

	"golift.io/starr"
)

// Define Base Path for Language Profile calls.
const bpLanguageProfile = APIver + "/languageProfile"

// LanguageProfile is the /api/v3/languageprofile endpoint.
type LanguageProfile struct {
	UpgradeAllowed bool         `json:"upgradeAllowed"`
	ID             int64        `json:"id,omitempty"`
	Name           string       `json:"name"`
	Cutoff         *starr.Value `json:"cutoff"`
	Languages      []Language   `json:"languages"`
}

// Language is part of LanguageProfile.
type Language struct {
	Allowed  bool         `json:"allowed"`
	Language *starr.Value `json:"language"`
}

// GetLanguageProfiles returns all configured language profiles.
func (s *Sonarr) GetLanguageProfiles() ([]*LanguageProfile, error) {
	return s.GetLanguageProfilesContext(context.Background())
}

// GetLanguageProfilesContext returns all configured language profiles.
func (s *Sonarr) GetLanguageProfilesContext(ctx context.Context) ([]*LanguageProfile, error) {
	var output []*LanguageProfile

	req := starr.Request{URI: bpLanguageProfile}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetLanguageProfile returns a single language profile.
func (s *Sonarr) GetLanguageProfile(profileID int) (*LanguageProfile, error) {
	return s.GetLanguageProfileContext(context.Background(), profileID)
}

// GetLanguageProfileContext returns a single language profile.
func (s *Sonarr) GetLanguageProfileContext(ctx context.Context, profileID int) (*LanguageProfile, error) {
	var output LanguageProfile

	req := starr.Request{URI: path.Join(bpLanguageProfile, fmt.Sprint(profileID))}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// AddLanguageProfile creates a language profile.
func (s *Sonarr) AddLanguageProfile(profile *LanguageProfile) (*LanguageProfile, error) {
	return s.AddLanguageProfileContext(context.Background(), profile)
}

// AddLanguageProfileContext creates a language profile.
func (s *Sonarr) AddLanguageProfileContext(ctx context.Context, profile *LanguageProfile) (*LanguageProfile, error) {
	var output LanguageProfile

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(profile); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpLanguageProfile, err)
	}

	req := starr.Request{URI: bpLanguageProfile, Body: &body}
	if err := s.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateLanguageProfile updates the language profile.
func (s *Sonarr) UpdateLanguageProfile(profile *LanguageProfile) (*LanguageProfile, error) {
	return s.UpdateLanguageProfileContext(context.Background(), profile)
}

// UpdateLanguageProfileContext updates the language profile.
func (s *Sonarr) UpdateLanguageProfileContext(ctx context.Context, profile *LanguageProfile) (*LanguageProfile, error) {
	var output LanguageProfile

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(profile); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpLanguageProfile, err)
	}

	req := starr.Request{URI: path.Join(bpLanguageProfile, fmt.Sprint(int(profile.ID))), Body: &body}
	if err := s.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteLanguageProfile removes a single language profile.
func (s *Sonarr) DeleteLanguageProfile(profileID int) error {
	return s.DeleteLanguageProfileContext(context.Background(), profileID)
}

// DeleteLanguageProfileContext removes a single language profile.
func (s *Sonarr) DeleteLanguageProfileContext(ctx context.Context, profileID int) error {
	req := starr.Request{URI: path.Join(bpLanguageProfile, fmt.Sprint(profileID))}
	if err := s.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
