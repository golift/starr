package radarr

import (
	"context"
	"fmt"
	"path"
	"time"

	"golift.io/starr"
)

const bpHistory = APIver + "/history"

// History is the /api/v3/history endpoint.
type History struct {
	Page          int              `json:"page"`
	PageSize      int              `json:"pageSize"`
	SortKey       string           `json:"sortKey"`
	SortDirection string           `json:"sortDirection"`
	TotalRecords  int              `json:"totalRecords"`
	Records       []*HistoryRecord `json:"records"`
}

// HistoryRecord is part of the History data.
// Not all items have all Data members. Check EventType for what you need.
type HistoryRecord struct {
	ID                  int64                 `json:"id"`
	MovieID             int64                 `json:"movieId"`
	SourceTitle         string                `json:"sourceTitle"`
	Languages           []*starr.Value        `json:"languages"`
	Quality             *starr.Quality        `json:"quality"`
	CustomFormats       []*CustomFormatOutput `json:"customFormats"`
	QualityCutoffNotMet bool                  `json:"qualityCutoffNotMet"`
	Date                time.Time             `json:"date"`
	DownloadID          string                `json:"downloadId"`
	EventType           string                `json:"eventType"`
	Data                struct {
		Age                string         `json:"age"`
		AgeHours           string         `json:"ageHours"`
		AgeMinutes         string         `json:"ageMinutes"`
		DownloadClient     string         `json:"downloadClient"`
		DownloadClientName string         `json:"downloadClientName"`
		DownloadURL        string         `json:"downloadUrl"`
		DroppedPath        string         `json:"droppedPath"`
		FileID             string         `json:"fileId"`
		GUID               string         `json:"guid"`
		ImportedPath       string         `json:"importedPath"`
		Indexer            string         `json:"indexer"`
		IndexerFlags       string         `json:"indexerFlags"`
		IndexerID          string         `json:"indexerId"`
		Message            string         `json:"message"`
		NzbInfoURL         string         `json:"nzbInfoUrl"`
		Protocol           starr.Protocol `json:"protocol"`
		PublishedDate      time.Time      `json:"publishedDate"`
		Reason             string         `json:"reason"`
		ReleaseGroup       string         `json:"releaseGroup"`
		Size               string         `json:"size"`
		TmdbID             string         `json:"tmdbId"`
		TorrentInfoHash    string         `json:"torrentInfoHash"`
	} `json:"data"`
}

// GetHistory returns the Radarr History (grabs/failures/completed).
// If you need control over the page, use radarr.GetHistoryPage().
// This function simply returns the number of history records desired,
// up to the number of records present in the application.
// It grabs records in (paginated) batches of perPage, and concatenates
// them into one list. Passing zero for records will return all of them.
func (r *Radarr) GetHistory(records, perPage int) (*History, error) {
	return r.GetHistoryContext(context.Background(), records, perPage)
}

// GetHistoryContext returns the Radarr History (grabs/failures/completed).
func (r *Radarr) GetHistoryContext(ctx context.Context, records, perPage int) (*History, error) {
	hist := &History{Records: []*HistoryRecord{}}
	perPage = starr.SetPerPage(records, perPage)

	for page := 1; ; page++ {
		curr, err := r.GetHistoryPageContext(ctx, &starr.PageReq{PageSize: perPage, Page: page})
		if err != nil {
			return nil, err
		}

		hist.Records = append(hist.Records, curr.Records...)
		if len(hist.Records) >= curr.TotalRecords ||
			(len(hist.Records) >= records && records != 0) ||
			len(curr.Records) == 0 {
			hist.PageSize = curr.TotalRecords
			hist.TotalRecords = curr.TotalRecords
			hist.SortDirection = curr.SortDirection
			hist.SortKey = curr.SortKey

			break
		}

		perPage = starr.AdjustPerPage(records, curr.TotalRecords, len(hist.Records), perPage)
	}

	return hist, nil
}

// GetHistoryPage returns a single page from the Radarr History (grabs/failures/completed).
// The page size and number is configurable with the input request parameters.
func (r *Radarr) GetHistoryPage(params *starr.PageReq) (*History, error) {
	return r.GetHistoryPageContext(context.Background(), params)
}

// GetHistoryPageContext returns a single page from the Radarr History (grabs/failures/completed).
// The page size and number is configurable with the input request parameters.
func (r *Radarr) GetHistoryPageContext(ctx context.Context, params *starr.PageReq) (*History, error) {
	var output History

	req := starr.Request{URI: bpHistory, Query: params.Params()}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// Fail marks the given history item as failed by id.
func (r *Radarr) Fail(historyID int64) error {
	return r.FailContext(context.Background(), historyID)
}

// FailContext marks the given history item as failed by id.
func (r *Radarr) FailContext(ctx context.Context, historyID int64) error {
	if historyID < 1 {
		return fmt.Errorf("%w: invalid history ID: %d", starr.ErrRequestError, historyID)
	}

	var output interface{} // any ok

	// Strangely uses a POST without a payload.
	req := starr.Request{URI: path.Join(bpHistory, "failed", starr.Itoa(historyID))}
	if err := r.PostInto(ctx, req, &output); err != nil {
		return fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return nil
}
