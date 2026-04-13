package prowlarr

import (
	"context"
	"fmt"
	"net/url"
	"path"
	"time"

	"golift.io/starr"
)

const bpHistory = APIver + "/history"

// HistoryPage is a paged history response.
type HistoryPage struct {
	Page          int              `json:"page"`
	PageSize      int              `json:"pageSize"`
	SortKey       string           `json:"sortKey,omitempty"`
	SortDirection string           `json:"sortDirection,omitempty"`
	TotalRecords  int              `json:"totalRecords"`
	Records       []*HistoryRecord `json:"records"`
}

// HistoryRecord is one history entry.
type HistoryRecord struct {
	ID          int64          `json:"id"`
	SourceTitle string         `json:"sourceTitle,omitempty"`
	Date        time.Time      `json:"date"`
	EventType   string         `json:"eventType,omitempty"`
	DownloadID  string         `json:"downloadId,omitempty"`
	Data        map[string]any `json:"data,omitempty"`
}

// GetHistoryPage returns a page of history.
func (p *Prowlarr) GetHistoryPage(params *starr.PageReq) (*HistoryPage, error) {
	return p.GetHistoryPageContext(context.Background(), params)
}

// GetHistoryPageContext returns a page of history.
func (p *Prowlarr) GetHistoryPageContext(ctx context.Context, params *starr.PageReq) (*HistoryPage, error) {
	if params == nil {
		params = &starr.PageReq{}
	}

	var output HistoryPage

	req := starr.Request{URI: bpHistory, Query: params.Params()}
	if err := p.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// GetHistorySince returns history since a date.
func (p *Prowlarr) GetHistorySince(date time.Time, eventType string) ([]*HistoryRecord, error) {
	return p.GetHistorySinceContext(context.Background(), date, eventType)
}

// GetHistorySinceContext returns history since a date.
func (p *Prowlarr) GetHistorySinceContext(
	ctx context.Context,
	date time.Time,
	eventType string,
) ([]*HistoryRecord, error) {
	params := make(url.Values)
	params.Set("date", date.UTC().Format(time.RFC3339))

	if eventType != "" {
		params.Set("eventType", eventType)
	}

	var output []*HistoryRecord

	req := starr.Request{URI: path.Join(bpHistory, "since"), Query: params}
	if err := p.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetHistoryByIndexer returns history for an indexer.
func (p *Prowlarr) GetHistoryByIndexer(indexerID int64, eventType string, limit int) ([]*HistoryRecord, error) {
	return p.GetHistoryByIndexerContext(context.Background(), indexerID, eventType, limit)
}

// GetHistoryByIndexerContext returns history for an indexer.
func (p *Prowlarr) GetHistoryByIndexerContext(
	ctx context.Context, indexerID int64, eventType string, limit int,
) ([]*HistoryRecord, error) {
	params := make(url.Values)
	params.Set("indexerId", starr.Str(indexerID))

	if eventType != "" {
		params.Set("eventType", eventType)
	}

	if limit > 0 {
		params.Set("limit", starr.Str(limit))
	}

	var output []*HistoryRecord

	req := starr.Request{URI: path.Join(bpHistory, "indexer"), Query: params}
	if err := p.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}
