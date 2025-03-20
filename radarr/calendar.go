package radarr

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"golift.io/starr"
)

// Define Base Path for Calendar queries.
const bpCalendar = APIver + "/calendar"

// Calendar defines the filters for fetching calendar items.
// Start and End are required. Use starr.True() and starr.False() to fill in the booleans.
type Calendar struct {
	Start       time.Time
	End         time.Time
	Unmonitored bool
}

// GetCalendar returns calendars based on filters.
func (r *Radarr) GetCalendar(filter Calendar) ([]*Movie, error) {
	return r.GetCalendarContext(context.Background(), filter)
}

// GetCalendarContext returns calendars based on filters.
func (r *Radarr) GetCalendarContext(ctx context.Context, filter Calendar) ([]*Movie, error) {
	var output []*Movie

	req := starr.Request{URI: bpCalendar, Query: make(url.Values)}
	req.Query.Add("unmonitored", starr.Str(filter.Unmonitored))

	if !filter.Start.IsZero() {
		req.Query.Add("start", filter.Start.UTC().Format(starr.CalendarTimeFilterFormat))
	}

	if !filter.End.IsZero() {
		req.Query.Add("end", filter.End.UTC().Format(starr.CalendarTimeFilterFormat))
	}

	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

/*
// I could not get this method to work in Radar..
// https://discord.com/channels/264387956343570434/264387994302021632/1058509574107312168

// GetCalendarID returns a single calendar by ID.
func (r *Radarr) GetCalendarID(calendarID int64) (*Movie, error) {
	return r.GetCalendarIDContext(context.Background(), calendarID)
}

// GetCalendarIDContext returns a single calendar by ID.
func (r *Radarr) GetCalendarIDContext(ctx context.Context, calendarID int64) (*Movie, error) {
	var output *Movie

	req := starr.Request{URI: path.Join(bpCalendar, starr.Str(calendarID))}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
} /**/
