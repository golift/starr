package sonarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"
	"time"

	"golift.io/starr"
)

// Define Base Path for Block List queries.
const bpBlocklist = APIver + "/blocklist"

// BlockList represents the /api/v3/blocklist endpoint.
type BlockList struct {
	Page          int                `json:"page"`
	PageSize      int                `json:"pageSize"`
	SortKey       string             `json:"sortKey"`
	SortDirection string             `json:"sortDirection"`
	TotalRecords  int                `json:"totalRecords"`
	Records       []*BlockListRecord `json:"records"`
}

// BlockListRecord represents a single block list item.
type BlockListRecord struct {
	Series        *Series               `json:"series"`
	Quality       *starr.Quality        `json:"quality"`
	Languages     []*starr.Value        `json:"languages"`
	CustomFormats []*CustomFormatOutput `json:"customFormats"`
	EpisodeIDs    []int64               `json:"episodeIds"`
	ID            int64                 `json:"id"`
	SeriesID      int64                 `json:"seriesId"`
	Date          time.Time             `json:"date"`
	SourceTitle   string                `json:"sourceTitle"`
	Protocol      starr.Protocol        `json:"protocol"`
	Indexer       string                `json:"indexer"`
	Message       string                `json:"message"`
}

// GetBlockList returns the count of block list items requested.
// If you need control over the page, use GetBlockListPage().
func (s *Sonarr) GetBlockList(count int) (*BlockList, error) {
	return s.GetBlockListContext(context.Background(), count)
}

// GetBlockListContext returns block list items.
func (s *Sonarr) GetBlockListContext(ctx context.Context, records int) (*BlockList, error) {
	list := &BlockList{Records: []*BlockListRecord{}}
	perPage := starr.SetPerPage(records, 0)

	for page := 1; ; page++ {
		curr, err := s.GetBlockListPageContext(ctx, &starr.PageReq{PageSize: perPage, Page: page})
		if err != nil {
			return nil, err
		}

		list.Records = append(list.Records, curr.Records...)
		if len(list.Records) >= curr.TotalRecords ||
			(len(list.Records) >= records && records != 0) ||
			len(curr.Records) == 0 {
			list.PageSize = curr.TotalRecords
			list.TotalRecords = curr.TotalRecords
			list.SortDirection = curr.SortDirection
			list.SortKey = curr.SortKey

			break
		}

		perPage = starr.AdjustPerPage(records, curr.TotalRecords, len(list.Records), perPage)
	}

	return list, nil
}

// GetBlockListPage returns block list items based on filters.
func (s *Sonarr) GetBlockListPage(params *starr.PageReq) (*BlockList, error) {
	return s.GetBlockListPageContext(context.Background(), params)
}

// GetBlockListPageContext returns block list items based on filters.
func (s *Sonarr) GetBlockListPageContext(ctx context.Context, params *starr.PageReq) (*BlockList, error) {
	var output BlockList

	params.CheckSet("sortKey", "date")

	req := starr.Request{URI: bpBlocklist, Query: params.Params()}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteBlockList removes a single block list item.
func (s *Sonarr) DeleteBlockList(listID int64) error {
	return s.DeleteBlockListContext(context.Background(), listID)
}

// DeleteBlockListContext removes a single block list item with a context.
func (s *Sonarr) DeleteBlockListContext(ctx context.Context, listID int64) error {
	req := starr.Request{URI: path.Join(bpBlocklist, starr.Str(listID))}
	if err := s.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}

// DeleteBlockLists removes multiple block list items.
func (s *Sonarr) DeleteBlockLists(ids []int64) error {
	return s.DeleteBlockListsContext(context.Background(), ids)
}

// DeleteBlockListsContext removes multiple block list items with a context.
func (s *Sonarr) DeleteBlockListsContext(ctx context.Context, ids []int64) error {
	input := struct {
		IDs []int64 `json:"ids"`
	}{IDs: ids}

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(input); err != nil {
		return fmt.Errorf("json.Marshal(%s): %w", bpBlocklist, err)
	}

	req := starr.Request{URI: path.Join(bpBlocklist, "bulk"), Body: &body}
	if err := s.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
