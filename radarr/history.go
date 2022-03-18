package radarr

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"golift.io/starr"
)

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
	ID                  int64          `json:"id"`
	MovieID             int64          `json:"movieId"`
	SourceTitle         string         `json:"sourceTitle"`
	Languages           []*starr.Value `json:"languages"`
	Quality             *starr.Quality `json:"quality"`
	CustomFormats       []interface{}  `json:"customFormats"`
	QualityCutoffNotMet bool           `json:"qualityCutoffNotMet"`
	Date                time.Time      `json:"date"`
	DownloadID          string         `json:"downloadId"`
	EventType           string         `json:"eventType"`
	Data                struct {
		Age                string    `json:"age"`
		AgeHours           string    `json:"ageHours"`
		AgeMinutes         string    `json:"ageMinutes"`
		DownloadClient     string    `json:"downloadClient"`
		DownloadClientName string    `json:"downloadClientName"`
		DownloadURL        string    `json:"downloadUrl"`
		DroppedPath        string    `json:"droppedPath"`
		FileID             string    `json:"fileId"`
		GUID               string    `json:"guid"`
		ImportedPath       string    `json:"importedPath"`
		Indexer            string    `json:"indexer"`
		IndexerFlags       string    `json:"indexerFlags"`
		IndexerID          string    `json:"indexerId"`
		Message            string    `json:"message"`
		NzbInfoURL         string    `json:"nzbInfoUrl"`
		Protocol           string    `json:"protocol"`
		PublishedDate      time.Time `json:"publishedDate"`
		Reason             string    `json:"reason"`
		ReleaseGroup       string    `json:"releaseGroup"`
		Size               string    `json:"size"`
		TmdbID             string    `json:"tmdbId"`
		TorrentInfoHash    string    `json:"torrentInfoHash"`
	} `json:"data"`
}

// GetHistory returns the Radarr History (grabs/failures/completed).
// WARNING: 12/30/2021 - this method changed. The second argument no longer
// controls which page is returned, but instead adjusts the pagination size.
// If you need control over the page, use radarr.GetHistoryPage().
// This function simply returns the number of history records desired,
// up to the number of records present in the application.
// It grabs records in (paginated) batches of perPage, and concatenates
// them into one list.  Passing zero for records will return all of them.
func (r *Radarr) GetHistory(records, perPage int) (*History, error) {
	return r.GetHistoryContext(context.Background(), records, perPage)
}

// GetHistoryContext returns the Radarr History (grabs/failures/completed).
func (r *Radarr) GetHistoryContext(ctx context.Context, records, perPage int) (*History, error) {
	hist := &History{Records: []*HistoryRecord{}}
	perPage = starr.SetPerPage(records, perPage)

	for page := 1; ; page++ {
		curr, err := r.GetHistoryPageContext(ctx, &starr.Req{PageSize: perPage, Page: page})
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
func (r *Radarr) GetHistoryPage(params *starr.Req) (*History, error) {
	return r.GetHistoryPageContext(context.Background(), params)
}

// GetHistoryPageContext returns a single page from the Radarr History (grabs/failures/completed).
// The page size and number is configurable with the input request parameters.
func (r *Radarr) GetHistoryPageContext(ctx context.Context, params *starr.Req) (*History, error) {
	var history History

	_, err := r.GetInto(ctx, "v3/history", params.Params(), &history)
	if err != nil {
		return nil, fmt.Errorf("api.Get(history): %w", err)
	}

	return &history, nil
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

	// Strangely uses a POST without a payload.
	_, err := r.Post(ctx, "v3/history/failed/"+strconv.FormatInt(historyID, starr.Base10), nil, nil)
	if err != nil {
		return fmt.Errorf("api.Post(history/failed): %w", err)
	}

	return nil
}
