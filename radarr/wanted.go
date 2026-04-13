package radarr

import (
	"context"
	"fmt"
	"path"

	"golift.io/starr"
)

const bpWanted = APIver + "/wanted"

// WantedMoviesPage is a paged list of movies from /api/v3/wanted/missing or /wanted/cutoff.
type WantedMoviesPage struct {
	Page          int      `json:"page"`
	PageSize      int      `json:"pageSize"`
	SortKey       string   `json:"sortKey,omitempty"`
	SortDirection string   `json:"sortDirection,omitempty"`
	TotalRecords  int      `json:"totalRecords"`
	Records       []*Movie `json:"records"`
}

func wantedPageParams(params *starr.PageReq) *starr.PageReq {
	if params == nil {
		return &starr.PageReq{}
	}

	return params
}

// GetWantedMissingPage returns a page of missing movies.
func (r *Radarr) GetWantedMissingPage(params *starr.PageReq) (*WantedMoviesPage, error) {
	return r.GetWantedMissingPageContext(context.Background(), params)
}

// GetWantedMissingPageContext returns a page of missing movies.
func (r *Radarr) GetWantedMissingPageContext(ctx context.Context, params *starr.PageReq) (*WantedMoviesPage, error) {
	var output WantedMoviesPage

	p := wantedPageParams(params)
	p.CheckSet("sortKey", "releaseDate")

	req := starr.Request{URI: path.Join(bpWanted, "missing"), Query: p.Params()}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// GetWantedCutoffPage returns a page of movies past quality cutoff.
func (r *Radarr) GetWantedCutoffPage(params *starr.PageReq) (*WantedMoviesPage, error) {
	return r.GetWantedCutoffPageContext(context.Background(), params)
}

// GetWantedCutoffPageContext returns a page of movies past quality cutoff.
func (r *Radarr) GetWantedCutoffPageContext(ctx context.Context, params *starr.PageReq) (*WantedMoviesPage, error) {
	var output WantedMoviesPage

	p := wantedPageParams(params)
	p.CheckSet("sortKey", "releaseDate")

	req := starr.Request{URI: path.Join(bpWanted, "cutoff"), Query: p.Params()}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}
