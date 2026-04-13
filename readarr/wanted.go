package readarr

import (
	"context"
	"fmt"
	"path"

	"golift.io/starr"
)

const bpWanted = APIver + "/wanted"

// WantedBooksPage is a paged list of books from /api/v1/wanted/missing or /wanted/cutoff.
type WantedBooksPage struct {
	Page          int     `json:"page"`
	PageSize      int     `json:"pageSize"`
	SortKey       string  `json:"sortKey,omitempty"`
	SortDirection string  `json:"sortDirection,omitempty"`
	TotalRecords  int     `json:"totalRecords"`
	Records       []*Book `json:"records"`
}

func wantedPageParams(params *starr.PageReq) *starr.PageReq {
	if params == nil {
		return &starr.PageReq{}
	}

	return params
}

// GetWantedMissingPage returns a page of missing books.
func (r *Readarr) GetWantedMissingPage(params *starr.PageReq) (*WantedBooksPage, error) {
	return r.GetWantedMissingPageContext(context.Background(), params)
}

// GetWantedMissingPageContext returns a page of missing books.
func (r *Readarr) GetWantedMissingPageContext(ctx context.Context, params *starr.PageReq) (*WantedBooksPage, error) {
	var output WantedBooksPage

	p := wantedPageParams(params)
	p.CheckSet("sortKey", "releaseDate")

	req := starr.Request{URI: path.Join(bpWanted, "missing"), Query: p.Params()}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// GetWantedCutoffPage returns a page of books past quality cutoff.
func (r *Readarr) GetWantedCutoffPage(params *starr.PageReq) (*WantedBooksPage, error) {
	return r.GetWantedCutoffPageContext(context.Background(), params)
}

// GetWantedCutoffPageContext returns a page of books past quality cutoff.
func (r *Readarr) GetWantedCutoffPageContext(ctx context.Context, params *starr.PageReq) (*WantedBooksPage, error) {
	var output WantedBooksPage

	p := wantedPageParams(params)
	p.CheckSet("sortKey", "releaseDate")

	req := starr.Request{URI: path.Join(bpWanted, "cutoff"), Query: p.Params()}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}
