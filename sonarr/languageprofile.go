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

// Define Base Path for Language Profile calls.
const bpLanguageProfile = APIver + "/languageProfile"

// GetLanguageProfiles returns all configured language profiles.
func (s *Sonarr) GetLanguageProfiles() ([]*LanguageProfile, error) {
	return s.GetLanguageProfilesContext(context.Background())
}

func (s *Sonarr) GetLanguageProfilesContext(ctx context.Context) ([]*LanguageProfile, error) {
	var output []*LanguageProfile

	if err := s.GetInto(ctx, bpLanguageProfile, nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(languageProfile): %w", err)
	}

	return output, nil
}

// GetLanguageProfile returns a single language profile.
func (s *Sonarr) GetLanguageProfile(profileID int) (*LanguageProfile, error) {
	return s.GetLanguageProfileContext(context.Background(), profileID)
}

func (s *Sonarr) GetLanguageProfileContext(ctx context.Context, profileID int) (*LanguageProfile, error) {
	var output *LanguageProfile

	uri := path.Join(bpLanguageProfile, strconv.Itoa(profileID))
	if err := s.GetInto(ctx, uri, nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(languageProfile): %w", err)
	}

	return output, nil
}

// AddLanguageProfile creates a language profile.
func (s *Sonarr) AddLanguageProfile(profile *LanguageProfile) (*LanguageProfile, error) {
	return s.AddLanguageProfileContext(context.Background(), profile)
}

func (s *Sonarr) AddLanguageProfileContext(ctx context.Context, profile *LanguageProfile) (*LanguageProfile, error) {
	var output LanguageProfile

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(profile); err != nil {
		return nil, fmt.Errorf("json.Marshal(languageProfile): %w", err)
	}

	if err := s.PostInto(ctx, bpLanguageProfile, nil, &body, &output); err != nil {
		return nil, fmt.Errorf("api.Post(languageProfile): %w", err)
	}

	return &output, nil
}

// UpdateLanguageProfile updates the language profile.
func (s *Sonarr) UpdateLanguageProfile(profile *LanguageProfile) (*LanguageProfile, error) {
	return s.UpdateLanguageProfileContext(context.Background(), profile)
}

func (s *Sonarr) UpdateLanguageProfileContext(ctx context.Context, profile *LanguageProfile) (*LanguageProfile, error) {
	var output LanguageProfile

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(profile); err != nil {
		return nil, fmt.Errorf("json.Marshal(languageProfile): %w", err)
	}

	uri := path.Join(bpLanguageProfile, strconv.Itoa(int(profile.ID)))
	if err := s.PutInto(ctx, uri, nil, &body, &output); err != nil {
		return nil, fmt.Errorf("api.Put(languageProfile): %w", err)
	}

	return &output, nil
}

// DeleteLanguageProfile removes a single language profile.
func (s *Sonarr) DeleteLanguageProfile(profileID int) error {
	return s.DeleteLanguageProfileContext(context.Background(), profileID)
}

func (s *Sonarr) DeleteLanguageProfileContext(ctx context.Context, profileID int) error {
	req := &starr.Request{URI: path.Join(bpLanguageProfile, fmt.Sprint(profileID))}
	if err := s.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", req.URI, err)
	}

	return nil
}
