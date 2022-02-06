package lidarr

import (
	"bytes"
	"context"
	"fmt"
	"strconv"

	"golift.io/starr"
)

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
		curr, err := l.GetHistoryPageContext(ctx, &starr.Req{PageSize: perPage, Page: page})
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
func (l *Lidarr) GetHistoryPage(params *starr.Req) (*History, error) {
	return l.GetHistoryPageContext(context.Background(), params)
}

// GetHistoryPageContext returns a single page from the Lidarr History (grabs/failures/completed).
// The page size and number is configurable with the input request parameters.
func (l *Lidarr) GetHistoryPageContext(ctx context.Context, params *starr.Req) (*History, error) {
	var history History

	err := l.GetInto(ctx, "v1/history", params.Params(), &history)
	if err != nil {
		return nil, fmt.Errorf("api.Get(history): %w", err)
	}

	return &history, nil
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

	post := bytes.NewBufferString("id=" + strconv.FormatInt(historyID, starr.Base10))

	_, err := l.Post(ctx, "v1/history/failed", nil, post)
	if err != nil {
		return fmt.Errorf("api.Post(history/failed): %w", err)
	}

	return nil
}
