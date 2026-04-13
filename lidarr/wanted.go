package lidarr

import (
	"context"
	"fmt"
	"path"

	"golift.io/starr"
)

const bpWanted = APIver + "/wanted"

// WantedAlbumsPage is a paged list of albums from /api/v1/wanted/missing or /wanted/cutoff.
type WantedAlbumsPage struct {
	Page          int      `json:"page"`
	PageSize      int      `json:"pageSize"`
	SortKey       string   `json:"sortKey,omitempty"`
	SortDirection string   `json:"sortDirection,omitempty"`
	TotalRecords  int      `json:"totalRecords"`
	Records       []*Album `json:"records"`
}

func wantedPageParams(params *starr.PageReq) *starr.PageReq {
	if params == nil {
		return &starr.PageReq{}
	}

	return params
}

// GetWantedMissingPage returns a page of missing albums.
func (l *Lidarr) GetWantedMissingPage(params *starr.PageReq) (*WantedAlbumsPage, error) {
	return l.GetWantedMissingPageContext(context.Background(), params)
}

// GetWantedMissingPageContext returns a page of missing albums.
func (l *Lidarr) GetWantedMissingPageContext(ctx context.Context, params *starr.PageReq) (*WantedAlbumsPage, error) {
	var output WantedAlbumsPage

	p := wantedPageParams(params)
	p.CheckSet("sortKey", "releaseDate")

	req := starr.Request{URI: path.Join(bpWanted, "missing"), Query: p.Params()}
	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// GetWantedCutoffPage returns a page of albums past quality cutoff.
func (l *Lidarr) GetWantedCutoffPage(params *starr.PageReq) (*WantedAlbumsPage, error) {
	return l.GetWantedCutoffPageContext(context.Background(), params)
}

// GetWantedCutoffPageContext returns a page of albums past quality cutoff.
func (l *Lidarr) GetWantedCutoffPageContext(ctx context.Context, params *starr.PageReq) (*WantedAlbumsPage, error) {
	var output WantedAlbumsPage

	p := wantedPageParams(params)
	p.CheckSet("sortKey", "releaseDate")

	req := starr.Request{URI: path.Join(bpWanted, "cutoff"), Query: p.Params()}
	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}
