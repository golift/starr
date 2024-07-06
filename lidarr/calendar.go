package lidarr

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
type Calendar struct {
	Start         time.Time
	End           time.Time
	Unmonitored   bool
	IncludeArtist bool
}

// GetCalendar returns calendars based on filters.
func (l *Lidarr) GetCalendar(filter Calendar) ([]*Album, error) {
	return l.GetCalendarContext(context.Background(), filter)
}

// GetCalendarContext returns calendars based on filters.
func (l *Lidarr) GetCalendarContext(ctx context.Context, filter Calendar) ([]*Album, error) {
	var output []*Album

	req := starr.Request{URI: bpCalendar, Query: make(url.Values)}
	req.Query.Add("unmonitored", starr.Itoa(filter.Unmonitored))
	req.Query.Add("includeArtist", starr.Itoa(filter.IncludeArtist))

	if !filter.Start.IsZero() {
		req.Query.Add("start", filter.Start.UTC().Format(starr.CalendarTimeFilterFormat))
	}

	if !filter.End.IsZero() {
		req.Query.Add("end", filter.End.UTC().Format(starr.CalendarTimeFilterFormat))
	}

	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetCalendarID returns a single calendar by ID.
func (l *Lidarr) GetCalendarID(calendarID int64) (*Album, error) {
	return l.GetCalendarIDContext(context.Background(), calendarID)
}

// GetCalendarIDContext returns a single calendar by ID.
func (l *Lidarr) GetCalendarIDContext(ctx context.Context, calendarID int64) (*Album, error) {
	var output *Album

	req := starr.Request{URI: path.Join(bpCalendar, starr.Itoa(calendarID))}
	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}
