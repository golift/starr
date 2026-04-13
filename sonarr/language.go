package sonarr

import (
	"context"
	"fmt"
	"path"

	"golift.io/starr"
)

const bpLanguage = APIver + "/language"

// AudioLanguage is an item from /api/v3/language (episode audio languages).
type AudioLanguage struct {
	ID        int    `json:"id"`
	Name      string `json:"name,omitempty"`
	NameLower string `json:"nameLower,omitempty"`
}

// GetAudioLanguages returns all languages from the /api/v3/language endpoint.
func (s *Sonarr) GetAudioLanguages() ([]*AudioLanguage, error) {
	return s.GetAudioLanguagesContext(context.Background())
}

// GetAudioLanguagesContext returns all languages from the /api/v3/language endpoint.
func (s *Sonarr) GetAudioLanguagesContext(ctx context.Context) ([]*AudioLanguage, error) {
	var output []*AudioLanguage

	req := starr.Request{URI: bpLanguage}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetAudioLanguage returns a single language by id.
func (s *Sonarr) GetAudioLanguage(id int) (*AudioLanguage, error) {
	return s.GetAudioLanguageContext(context.Background(), id)
}

// GetAudioLanguageContext returns a single language by id.
func (s *Sonarr) GetAudioLanguageContext(ctx context.Context, id int) (*AudioLanguage, error) {
	var output AudioLanguage

	req := starr.Request{URI: path.Join(bpLanguage, starr.Str(id))}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}
