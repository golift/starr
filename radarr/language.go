package radarr

import (
	"context"
	"fmt"
	"path"

	"golift.io/starr"
)

const bpLanguage = APIver + "/language"

// Language is an item from /api/v3/language.
type Language struct {
	ID   int    `json:"id"`
	Name string `json:"name,omitempty"`
}

// GetLanguages returns all movie languages.
func (r *Radarr) GetLanguages() ([]*Language, error) {
	return r.GetLanguagesContext(context.Background())
}

// GetLanguagesContext returns all movie languages.
func (r *Radarr) GetLanguagesContext(ctx context.Context) ([]*Language, error) {
	var output []*Language

	req := starr.Request{URI: bpLanguage}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetLanguage returns a single language by id.
func (r *Radarr) GetLanguage(id int) (*Language, error) {
	return r.GetLanguageContext(context.Background(), id)
}

// GetLanguageContext returns a single language by id.
func (r *Radarr) GetLanguageContext(ctx context.Context, id int) (*Language, error) {
	var output Language

	req := starr.Request{URI: path.Join(bpLanguage, starr.Str(id))}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}
