package sonarr

import (
	"context"
	"fmt"
	"path"

	"golift.io/starr"
)

const bpWanted = APIver + "/wanted"

// WantedEpisodesPage is a paged list of episodes from /api/v3/wanted/missing or /wanted/cutoff.
type WantedEpisodesPage struct {
	Page          int        `json:"page"`
	PageSize      int        `json:"pageSize"`
	SortKey       string     `json:"sortKey,omitempty"`
	SortDirection string     `json:"sortDirection,omitempty"`
	TotalRecords  int        `json:"totalRecords"`
	Records       []*Episode `json:"records"`
}

func wantedPageParams(params *starr.PageReq) *starr.PageReq {
	if params == nil {
		return &starr.PageReq{}
	}

	return params
}

// GetWantedMissingPage returns a page of missing episodes.
func (s *Sonarr) GetWantedMissingPage(params *starr.PageReq) (*WantedEpisodesPage, error) {
	return s.GetWantedMissingPageContext(context.Background(), params)
}

// GetWantedMissingPageContext returns a page of missing episodes.
func (s *Sonarr) GetWantedMissingPageContext(ctx context.Context, params *starr.PageReq) (*WantedEpisodesPage, error) {
	var output WantedEpisodesPage

	p := wantedPageParams(params)
	p.CheckSet("sortKey", "airDateUtc")

	req := starr.Request{URI: path.Join(bpWanted, "missing"), Query: p.Params()}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// GetWantedMissingEpisode returns a single missing episode by episode ID.
func (s *Sonarr) GetWantedMissingEpisode(episodeID int64) (*Episode, error) {
	return s.GetWantedMissingEpisodeContext(context.Background(), episodeID)
}

// GetWantedMissingEpisodeContext returns a single missing episode by episode ID.
func (s *Sonarr) GetWantedMissingEpisodeContext(ctx context.Context, episodeID int64) (*Episode, error) {
	var output Episode

	req := starr.Request{URI: path.Join(bpWanted, "missing", starr.Str(episodeID))}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// GetWantedCutoffPage returns a page of episodes past quality cutoff.
func (s *Sonarr) GetWantedCutoffPage(params *starr.PageReq) (*WantedEpisodesPage, error) {
	return s.GetWantedCutoffPageContext(context.Background(), params)
}

// GetWantedCutoffPageContext returns a page of episodes past quality cutoff.
func (s *Sonarr) GetWantedCutoffPageContext(ctx context.Context, params *starr.PageReq) (*WantedEpisodesPage, error) {
	var output WantedEpisodesPage

	p := wantedPageParams(params)
	p.CheckSet("sortKey", "airDateUtc")

	req := starr.Request{URI: path.Join(bpWanted, "cutoff"), Query: p.Params()}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// GetWantedCutoffEpisode returns a single cutoff-unmet episode by episode ID.
func (s *Sonarr) GetWantedCutoffEpisode(episodeID int64) (*Episode, error) {
	return s.GetWantedCutoffEpisodeContext(context.Background(), episodeID)
}

// GetWantedCutoffEpisodeContext returns a single cutoff-unmet episode by episode ID.
func (s *Sonarr) GetWantedCutoffEpisodeContext(ctx context.Context, episodeID int64) (*Episode, error) {
	var output Episode

	req := starr.Request{URI: path.Join(bpWanted, "cutoff", starr.Str(episodeID))}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}
