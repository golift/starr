package radarr

import (
	"context"
	"fmt"
	"net/url"
	"path"

	"golift.io/starr"
)

const bpAltTitle = APIver + "/alttitle"

// GetAlternativeTitles returns alternative titles for a movie.
func (r *Radarr) GetAlternativeTitles(movieID, movieMetadataID int64) ([]*AlternativeTitle, error) {
	return r.GetAlternativeTitlesContext(context.Background(), movieID, movieMetadataID)
}

// GetAlternativeTitlesContext returns alternative titles for a movie.
func (r *Radarr) GetAlternativeTitlesContext(
	ctx context.Context,
	movieID, movieMetadataID int64,
) ([]*AlternativeTitle, error) {
	params := make(url.Values)
	if movieID != 0 {
		params.Set("movieId", starr.Str(movieID))
	}

	if movieMetadataID != 0 {
		params.Set("movieMetadataId", starr.Str(movieMetadataID))
	}

	var output []*AlternativeTitle

	req := starr.Request{URI: bpAltTitle, Query: params}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetAlternativeTitle returns a single alternative title by id.
func (r *Radarr) GetAlternativeTitle(altTitleID int64) (*AlternativeTitle, error) {
	return r.GetAlternativeTitleContext(context.Background(), altTitleID)
}

// GetAlternativeTitleContext returns a single alternative title by id.
func (r *Radarr) GetAlternativeTitleContext(ctx context.Context, altTitleID int64) (*AlternativeTitle, error) {
	var output AlternativeTitle

	req := starr.Request{URI: path.Join(bpAltTitle, starr.Str(altTitleID))}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}
