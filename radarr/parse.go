package radarr

import (
	"context"
	"fmt"
	"net/url"

	"golift.io/starr"
)

const bpParse = APIver + "/parse"

// ParseOutput is returned from GET /api/v3/parse.
type ParseOutput struct {
	CustomFormats     []*CustomFormatOutput `json:"customFormats,omitempty"`
	CustomFormatScore int64                 `json:"customFormatScore"`
	Languages         []*starr.Value        `json:"languages,omitempty"`
	ID                int64                 `json:"id"`
	Title             string                `json:"title,omitempty"`
	Movie             *Movie                `json:"movie,omitempty"`
	ParsedMovieInfo   map[string]any        `json:"parsedMovieInfo,omitempty"`
}

// Parse resolves a release title into parsed movie metadata.
func (r *Radarr) Parse(title string) (*ParseOutput, error) {
	return r.ParseContext(context.Background(), title)
}

// ParseContext resolves a release title into parsed movie metadata.
func (r *Radarr) ParseContext(ctx context.Context, title string) (*ParseOutput, error) {
	var output *ParseOutput

	req := starr.Request{URI: bpParse, Query: make(url.Values)}
	req.Query.Set("title", title)

	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}
