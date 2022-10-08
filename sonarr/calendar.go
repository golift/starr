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

// Calendar represents the data that may be returned from the calendar endpoint.
type Calendar struct {
	AbsoluteEpisodeNumber      int            `json:"absoluteEpisodeNumber"`
	AirDate                    string         `json:"airDate"`
	AirDateUtc                 time.Time      `json:"airDateUtc"`
	EndTime                    time.Time      `json:"endTime"`
	EpisodeFile                *EpisodeFile   `json:"episodeFile"`
	EpisodeFileID              int64          `json:"episodeFileId"`
	EpisodeNumber              int            `json:"episodeNumber"`
	GrabDate                   time.Time      `json:"grabDate"`
	Grabbed                    bool           `json:"grabbed"`
	HasFile                    bool           `json:"hasFile"`
	ID                         int64          `json:"id"`
	Images                     []*starr.Image `json:"images,omitempty"`
	Monitored                  bool           `json:"monitored"`
	Overview                   string         `json:"overview"`
	SceneAbsoluteEpisodeNumber int            `json:"sceneAbsoluteEpisodeNumber"`
	SceneEpisodeNumber         int            `json:"sceneEpisodeNumber"`
	SceneSeasonNumber          int            `json:"sceneSeasonNumber"`
	SeasonNumber               int            `json:"seasonNumber"`
	Series                     *Series        `json:"series"`
	SeriesID                   int64          `json:"seriesId"`
	SeriesTitle                string         `json:"seriesTitle"`
	Title                      string         `json:"title"`
	TvdbID                     int64          `json:"tvdbId"`
	UnverifiedSceneNumbering   bool           `json:"unverifiedSceneNumbering"`
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
