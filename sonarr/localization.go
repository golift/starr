package sonarr

import (
	"context"
	"fmt"
	"path"

	"golift.io/starr"
)

const bpLocalization = APIver + "/localization"

// Localization is the /api/v3/localization resource.
type Localization struct {
	ID      int               `json:"id"`
	Strings map[string]string `json:"strings,omitempty"`
}

// GetLocalization returns the default localization dictionary.
func (s *Sonarr) GetLocalization() (*Localization, error) {
	return s.GetLocalizationContext(context.Background())
}

// GetLocalizationContext returns the default localization dictionary.
func (s *Sonarr) GetLocalizationContext(ctx context.Context) (*Localization, error) {
	var output Localization

	req := starr.Request{URI: bpLocalization}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// GetLocalizationByID returns a localization dictionary by id.
func (s *Sonarr) GetLocalizationByID(id int) (*Localization, error) {
	return s.GetLocalizationByIDContext(context.Background(), id)
}

// GetLocalizationByIDContext returns a localization dictionary by id.
func (s *Sonarr) GetLocalizationByIDContext(ctx context.Context, id int) (*Localization, error) {
	var output Localization

	req := starr.Request{URI: path.Join(bpLocalization, starr.Str(id))}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// UILanguage is an item from /api/v3/localization/language.
type UILanguage struct {
	Identifier string `json:"identifier,omitempty"`
}

// GetLocalizationLanguages returns available UI languages.
func (s *Sonarr) GetLocalizationLanguages() ([]*UILanguage, error) {
	return s.GetLocalizationLanguagesContext(context.Background())
}

// GetLocalizationLanguagesContext returns available UI languages.
func (s *Sonarr) GetLocalizationLanguagesContext(ctx context.Context) ([]*UILanguage, error) {
	var output []*UILanguage

	req := starr.Request{URI: path.Join(bpLocalization, "language")}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}
