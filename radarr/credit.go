package radarr

import (
	"context"
	"fmt"
	"net/url"
	"path"

	"golift.io/starr"
)

const bpCredit = APIver + "/credit"

// Credit is a cast/crew credit from /api/v3/credit.
type Credit struct {
	ID              int64          `json:"id,omitempty"`
	CreditID        string         `json:"creditId,omitempty"`
	PersonName      string         `json:"personName,omitempty"`
	Job             string         `json:"job,omitempty"`
	Department      string         `json:"department,omitempty"`
	Character       string         `json:"character,omitempty"`
	Order           int            `json:"order,omitempty"`
	ProfilePath     string         `json:"profilePath,omitempty"`
	PersonTmdbID    int64          `json:"personTmdbId,omitempty"`
	MovieID         int64          `json:"movieId,omitempty"`
	MovieMetadataID int64          `json:"movieMetadataId,omitempty"`
	Type            string         `json:"type,omitempty"`
	Images          []*starr.Image `json:"images,omitempty"`
}

// GetCredits returns credits for a movie or movie metadata.
func (r *Radarr) GetCredits(movieID, movieMetadataID int64) ([]*Credit, error) {
	return r.GetCreditsContext(context.Background(), movieID, movieMetadataID)
}

// GetCreditsContext returns credits for a movie or movie metadata.
func (r *Radarr) GetCreditsContext(ctx context.Context, movieID, movieMetadataID int64) ([]*Credit, error) {
	params := make(url.Values)
	if movieID != 0 {
		params.Set("movieId", starr.Str(movieID))
	}

	if movieMetadataID != 0 {
		params.Set("movieMetadataId", starr.Str(movieMetadataID))
	}

	var output []*Credit

	req := starr.Request{URI: bpCredit, Query: params}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetCredit returns a single credit by id.
func (r *Radarr) GetCredit(creditID int64) (*Credit, error) {
	return r.GetCreditContext(context.Background(), creditID)
}

// GetCreditContext returns a single credit by id.
func (r *Radarr) GetCreditContext(ctx context.Context, creditID int64) (*Credit, error) {
	var output Credit

	req := starr.Request{URI: path.Join(bpCredit, starr.Str(creditID))}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}
