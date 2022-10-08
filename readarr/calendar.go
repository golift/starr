package readarr

import (
	"context"
	"fmt"
	"net/url"
	"path"
	"time"

	"golift.io/starr"
)

// Define Base Path for Calendar queries.
const bpCalendar = APIver + "/calendar"

// Calendar defines the filters for fetching calendar items.
// Start and End are required. Use starr.True() and starr.False() to fill in the booleans.
type Calendar struct {
	Start         time.Time
	End           time.Time
	Unmonitored   *bool
	IncludeAuthor *bool
}

// GetCalendar returns calendars based on filters.
func (r *Readarr) GetCalendar(filter Calendar) ([]*Book, error) {
	return r.GetCalendarContext(context.Background(), filter)
}

// GetCalendarContext returns calendars based on filters.
func (r *Readarr) GetCalendarContext(ctx context.Context, filter Calendar) ([]*Book, error) {
	var output []*Book

	req := starr.Request{URI: bpCalendar, Query: make(url.Values)}

	if !filter.Start.IsZero() {
		req.Query.Add("start", filter.Start.UTC().Format(starr.CalendarTimeFilterFormat))
	}

	if !filter.End.IsZero() {
		req.Query.Add("end", filter.End.UTC().Format(starr.CalendarTimeFilterFormat))
	}

	if filter.Unmonitored != nil {
		req.Query.Add("unmonitored", fmt.Sprint(*filter.Unmonitored))
	}

	if filter.IncludeAuthor != nil {
		req.Query.Add("includeAuthor", fmt.Sprint(*filter.IncludeAuthor))
	}

	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetCalendarID returns a single calendar by ID.
func (r *Readarr) GetCalendarID(calendarID int64) (*Book, error) {
	return r.GetCalendarIDContext(context.Background(), calendarID)
}

// GetCalendarIDContext returns a single calendar by ID.
func (r *Readarr) GetCalendarIDContext(ctx context.Context, calendarID int64) (*Book, error) {
	var output *Book

	req := starr.Request{URI: path.Join(bpCalendar, fmt.Sprint(calendarID))}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}
