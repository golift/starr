package sonarr

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
	Start                time.Time
	End                  time.Time
	Unmonitored          *bool
	IncludeSeries        *bool
	IncludeEpisodeFile   *bool
	IncludeEpisodeImages *bool
}

// GetCalendar returns calendars based on filters.
func (s *Sonarr) GetCalendar(filter Calendar) ([]*Episode, error) {
	return s.GetCalendarContext(context.Background(), filter)
}

// GetCalendarContext returns calendars based on filters.
func (s *Sonarr) GetCalendarContext(ctx context.Context, filter Calendar) ([]*Episode, error) {
	var output []*Episode

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

	if filter.IncludeSeries != nil {
		req.Query.Add("includeSeries", fmt.Sprint(*filter.IncludeSeries))
	}

	if filter.IncludeEpisodeFile != nil {
		req.Query.Add("includeEpisodeFile", fmt.Sprint(*filter.IncludeEpisodeFile))
	}

	if filter.IncludeEpisodeImages != nil {
		req.Query.Add("includeEpisodeImages", fmt.Sprint(*filter.IncludeEpisodeImages))
	}

	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetCalendarID returns a single calendar by ID.
func (s *Sonarr) GetCalendarID(calendarID int64) (*Episode, error) {
	return s.GetCalendarIDContext(context.Background(), calendarID)
}

// GetCalendarIDContext returns a single calendar by ID.
func (s *Sonarr) GetCalendarIDContext(ctx context.Context, calendarID int64) (*Episode, error) {
	var output *Episode

	req := starr.Request{URI: path.Join(bpCalendar, fmt.Sprint(calendarID))}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}
