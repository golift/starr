package lidarr

import (
	"bytes"
	"context"
	"fmt"
	"path"
	"time"

	"golift.io/starr"
)

const bpHistory = APIver + "/history"

// History represents the /api/v1/history endpoint.
type History struct {
	Page          int              `json:"page"`
	PageSize      int              `json:"pageSize"`
	SortKey       string           `json:"sortKey"`
	SortDirection string           `json:"sortDirection"`
	TotalRecords  int              `json:"totalRecords"`
	Records       []*HistoryRecord `json:"records"`
}

// HistoryRecord is part of the history. Not all items have all Data members.
// Check EventType for events you need.
type HistoryRecord struct {
	ID                  int64          `json:"id"`
	AlbumID             int64          `json:"albumId"`
	ArtistID            int64          `json:"artistId"`
	TrackID             int64          `json:"trackId"`
	SourceTitle         string         `json:"sourceTitle"`
	Quality             *starr.Quality `json:"quality"`
	QualityCutoffNotMet bool           `json:"qualityCutoffNotMet"`
	Date                time.Time      `json:"date"`
	DownloadID          string         `json:"downloadId"`
	EventType           string         `json:"eventType"`
	Data                struct {
		Age             string    `json:"age"`
		AgeHours        string    `json:"ageHours"`
		AgeMinutes      string    `json:"ageMinutes"`
		DownloadClient  string    `json:"downloadClient"`
		DownloadForced  string    `json:"downloadForced"`
		DownloadURL     string    `json:"downloadUrl"`
		DroppedPath     string    `json:"droppedPath"`
		GUID            string    `json:"guid"`
		ImportedPath    string    `json:"importedPath"`
		Indexer         string    `json:"indexer"`
		Message         string    `json:"message"`
		NzbInfoURL      string    `json:"nzbInfoUrl"`
		Protocol        string    `json:"protocol"`
		PublishedDate   time.Time `json:"publishedDate"`
		Reason          string    `json:"reason"`
		ReleaseGroup    string    `json:"releaseGroup"`
		Size            string    `json:"size"`
		StatusMessages  string    `json:"statusMessages"`
		TorrentInfoHash string    `json:"torrentInfoHash"`
	} `json:"data"`
}

// GetHistory returns the Lidarr History (grabs/failures/completed).
// WARNING: 12/30/2021 - this method changed.
// If you need control over the page, use lidarr.GetHistoryPage().
// This function simply returns the number of history records desired,
// up to the number of records present in the application.
// It grabs records in (paginated) batches of perPage, and concatenates
// them into one list.  Passing zero for records will return all of them.
func (l *Lidarr) GetHistory(records, perPage int) (*History, error) {
	return l.GetHistoryContext(context.Background(), records, perPage)
}

// GetHistoryContext returns the Lidarr History (grabs/failures/completed).
func (l *Lidarr) GetHistoryContext(ctx context.Context, records, perPage int) (*History, error) {
	hist := &History{Records: []*HistoryRecord{}}
	perPage = starr.SetPerPage(records, perPage)

	for page := 1; ; page++ {
		curr, err := l.GetHistoryPageContext(ctx, &starr.PageReq{PageSize: perPage, Page: page})
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

// GetHistoryPage returns a single page from the Lidarr History (grabs/failures/completed).
// The page size and number is configurable with the input request parameters.
func (l *Lidarr) GetHistoryPage(params *starr.PageReq) (*History, error) {
	return l.GetHistoryPageContext(context.Background(), params)
}

// GetHistoryPageContext returns a single page from the Lidarr History (grabs/failures/completed).
// The page size and number is configurable with the input request parameters.
func (l *Lidarr) GetHistoryPageContext(ctx context.Context, params *starr.PageReq) (*History, error) {
	var output History

	req := starr.Request{URI: bpHistory, Query: params.Params()}
	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", req, err)
	}

	return &output, nil
}

// Fail marks the given history item as failed by id.
func (l *Lidarr) Fail(historyID int64) error {
	return l.FailContext(context.Background(), historyID)
}

// FailContext marks the given history item as failed by id.
func (l *Lidarr) FailContext(ctx context.Context, historyID int64) error {
	if historyID < 1 {
		return fmt.Errorf("%w: invalid history ID: %d", starr.ErrRequestError, historyID)
	}

	var output interface{}

	req := starr.Request{
		URI:  path.Join(bpHistory, "failed"),
		Body: bytes.NewBufferString("id=" + fmt.Sprint(historyID)),
	}
	if err := l.PostInto(ctx, req, &output); err != nil {
		return fmt.Errorf("api.Post(%s): %w", req, err)
	}

	return nil
}
