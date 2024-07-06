package prowlarr

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"golift.io/starr"
)

const bpSearch = APIver + "/search"

// Search is the output from the Prowlarr search endpoint.
type Search struct {
	GUID         string         `json:"guid"`
	Age          int64          `json:"age"`
	AgeHours     float64        `json:"ageHours"`
	AgeMinutes   float64        `json:"ageMinutes"`
	Size         int64          `json:"size"`
	Files        int            `json:"files"`
	Grabs        int            `json:"grabs"`
	IndexerID    int64          `json:"indexerId"`
	Indexer      string         `json:"indexer"`
	Title        string         `json:"title"`
	SortTitle    string         `json:"sortTitle"`
	ImdbID       int64          `json:"imdbId"`
	TmdbID       int64          `json:"tmdbId"`
	TvdbID       int64          `json:"tvdbId"`
	TvMazeID     int64          `json:"tvMazeId"`
	PublishDate  time.Time      `json:"publishDate"`
	CommentURL   string         `json:"commentUrl"`
	DownloadURL  string         `json:"downloadUrl"`
	InfoURL      string         `json:"infoUrl"`
	IndexerFlags []string       `json:"indexerFlags"`
	Categories   []*Category    `json:"categories"`
	Protocol     starr.Protocol `json:"protocol"`
	FileName     string         `json:"fileName"`
	InfoHash     string         `json:"infoHash"`
	Seeders      int            `json:"seeders"`
	Leechers     int            `json:"leechers"`
}

// Category is part of the Search output.
type Category struct {
	ID            int64       `json:"id"`
	Name          string      `json:"name"`
	SubCategories []*Category `json:"subCategories"`
}

// SearchInput is the input to the search endpoint.
type SearchInput struct {
	Query      string  `json:"query"` // Query is required. Fill it in.
	Type       string  `json:"type"`  // defaults to "search" if left empty
	IndexerIDs []int64 `json:"indexerIds"`
	Categories []int64 `json:"categories"`
	Limit      int     `json:"limit"`  // Defaults to 100 if left empty or less than 1.
	Offset     int     `json:"offset"` // Skip this many records.
}

// Search the Prowlarr indexers for media and content. Must provide a Query in the SearchInput.
func (p *Prowlarr) Search(search SearchInput) ([]*Search, error) {
	return p.SearchContext(context.Background(), search)
}

// SearchContext searches the Prowlarr indexers for media and content.
func (p *Prowlarr) SearchContext(ctx context.Context, search SearchInput) ([]*Search, error) {
	const defaultSearchLimit = 100

	if search.Type == "" {
		search.Type = "search"
	}

	if search.Limit < 1 {
		search.Limit = defaultSearchLimit
	}

	if search.Limit < 0 {
		search.Limit = 0
	}

	req := starr.Request{URI: bpSearch, Query: make(url.Values)}
	req.Query.Set("query", search.Query)
	req.Query.Set("type", search.Type)
	req.Query.Set("limit", starr.Itoa(int64(search.Limit)))
	req.Query.Set("offset", starr.Itoa(int64(search.Offset)))

	for _, val := range search.Categories {
		req.Query.Add("categories", starr.Itoa(val))
	}

	for _, val := range search.IndexerIDs {
		req.Query.Add("indexerIds", starr.Itoa(val))
	}

	var output []*Search
	if err := p.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}
