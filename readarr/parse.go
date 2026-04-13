package readarr

import (
	"context"
	"fmt"
	"net/url"

	"golift.io/starr"
)

const bpParse = APIver + "/parse"

// ParseOutput is returned from GET /api/v1/parse.
type ParseOutput struct {
	ID                int64   `json:"id"`
	Title             string  `json:"title,omitempty"`
	Author            *Author `json:"author,omitempty"`
	Books             []*Book `json:"books,omitempty"`
	CustomFormats     []any   `json:"customFormats,omitempty"`
	CustomFormatScore int64   `json:"customFormatScore"`
}

// Parse resolves a release title into parsed book metadata.
func (r *Readarr) Parse(title string) (*ParseOutput, error) {
	return r.ParseContext(context.Background(), title)
}

// ParseContext resolves a release title into parsed book metadata.
func (r *Readarr) ParseContext(ctx context.Context, title string) (*ParseOutput, error) {
	var output *ParseOutput

	req := starr.Request{URI: bpParse, Query: make(url.Values)}
	req.Query.Set("title", title)

	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}
