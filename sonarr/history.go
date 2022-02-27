package sonarr

import (
	"context"
	"fmt"
	"strconv"

	"golift.io/starr"
)

// GetHistory returns the Sonarr History (grabs/failures/completed).
// WARNING: 12/30/2021 - this method changed.
// If you need control over the page, use sonarr.GetHistoryPage().
// This function simply returns the number of history records desired,
// up to the number of records present in the application.
// It grabs records in (paginated) batches of perPage, and concatenates
// them into one list.  Passing zero for records will return all of them.
func (s *Sonarr) GetHistory(records, perPage int) (*History, error) {
	return s.GetHistoryContext(context.Background(), records, perPage)
}

func (s *Sonarr) GetHistoryContext(ctx context.Context, records, perPage int) (*History, error) {
	hist := &History{Records: []*HistoryRecord{}}
	perPage = starr.SetPerPage(records, perPage)

	for page := 1; ; page++ {
		curr, err := s.GetHistoryPageContext(ctx, &starr.Req{PageSize: perPage, Page: page})
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

// GetHistoryPage returns a single page from the Sonarr History (grabs/failures/completed).
// The page size and number is configurable with the input request parameters.
func (s *Sonarr) GetHistoryPage(params *starr.Req) (*History, error) {
	return s.GetHistoryPageContext(context.Background(), params)
}

func (s *Sonarr) GetHistoryPageContext(ctx context.Context, params *starr.Req) (*History, error) {
	var history History

	_, err := s.GetInto(ctx, "v3/history", params.Params(), &history)
	if err != nil {
		return nil, fmt.Errorf("api.Get(history): %w", err)
	}

	return &history, nil
}

// Fail marks the given history item as failed by id.
func (s *Sonarr) Fail(historyID int64) error {
	return s.FailContext(context.Background(), historyID)
}

func (s *Sonarr) FailContext(ctx context.Context, historyID int64) error {
	if historyID < 1 {
		return fmt.Errorf("%w: invalid history ID: %d", starr.ErrRequestError, historyID)
	}

	// Strangely uses a POST without a payload.
	_, err := s.Post(ctx, "v3/history/failed/"+strconv.FormatInt(historyID, starr.Base10), nil, nil)
	if err != nil {
		return fmt.Errorf("api.Post(history/failed): %w", err)
	}

	return nil
}
