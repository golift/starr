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

// CalendarTimeFilterFormat is the Go  time format the calendar expects the filter to be in.
const CalendarTimeFilterFormat = "2006-01-02T03:04:05.000Z"

// Calendar represents the data that may be returned from the calendar endpoint.
type Calendar struct {
	ID                         int64          `json:"id"`
	SeriesID                   int64          `json:"seriesId"`
	TvdbID                     int64          `json:"tvdbId"`
	EpisodeFileID              int64          `json:"episodeFileId"`
	SeasonNumber               int            `json:"seasonNumber"`
	EpisodeNumber              int            `json:"episodeNumber"`
	Title                      string         `json:"title"`
	AirDate                    string         `json:"airDate"`
	AirDateUtc                 time.Time      `json:"airDateUtc"`
	Overview                   string         `json:"overview"`
	EpisodeFile                *EpisodeFile   `json:"episodeFile"`
	HasFile                    bool           `json:"hasFile"`
	Monitored                  bool           `json:"monitored"`
	AbsoluteEpisodeNumber      int            `json:"absoluteEpisodeNumber"`
	SceneAbsoluteEpisodeNumber int            `json:"sceneAbsoluteEpisodeNumber"`
	SceneEpisodeNumber         int            `json:"sceneEpisodeNumber"`
	SceneSeasonNumber          int            `json:"sceneSeasonNumber"`
	UnverifiedSceneNumbering   bool           `json:"unverifiedSceneNumbering"`
	EndTime                    time.Time      `json:"endTime"`
	GrabDate                   time.Time      `json:"grabDate"`
	SeriesTitle                string         `json:"seriesTitle"`
	Series                     *Series        `json:"series"`
	Images                     []*starr.Image `json:"images,omitempty"`
	Grabbed                    bool           `json:"grabbed"`
}

// CalendarInput defines the filters for fetching calendar items.
// Start and End are required. Use starr.True() and starr.False() to fill in the booleans.
type CalendarInput struct {
	Start                time.Time
	End                  time.Time
	Unmonitored          *bool
	IncludeSeries        *bool
	IncludeEpisodeFile   *bool
	IncludeEpisodeImages *bool
}

// GetCalendar returns calendars based on filters.
func (s *Sonarr) GetCalendar(filter CalendarInput) ([]*Calendar, error) {
	return s.GetCalendarContext(context.Background(), filter)
}

// GetCalendarContext returns calendars based on filters.
func (s *Sonarr) GetCalendarContext(ctx context.Context, filter CalendarInput) ([]*Calendar, error) {
	var output []*Calendar

	req := starr.Request{URI: bpCalendar, Query: make(url.Values)}

	if !filter.Start.IsZero() {
		req.Query.Add("start", filter.Start.UTC().Format(CalendarTimeFilterFormat))
	}

	if !filter.End.IsZero() {
		req.Query.Add("end", filter.End.UTC().Format(CalendarTimeFilterFormat))
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
func (s *Sonarr) GetCalendarID(calendarID int64) (*Calendar, error) {
	return s.GetCalendarIDContext(context.Background(), calendarID)
}

// GetCalendarIDContext returns a single calendar by ID.
func (s *Sonarr) GetCalendarIDContext(ctx context.Context, calendarID int64) (*Calendar, error) {
	var output *Calendar

	req := starr.Request{URI: path.Join(bpCalendar, fmt.Sprint(calendarID))}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}
