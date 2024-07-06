package sonarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"path"
	"time"

	"golift.io/starr"
)

// Define Base Path for Series calls.
const bpSeries = APIver + "/series"

// AddSeriesInput is the input for /api/v3/series endpoint.
type AddSeriesInput struct {
	Monitored         bool           `json:"monitored"`
	SeasonFolder      bool           `json:"seasonFolder,omitempty"`
	UseSceneNumbering bool           `json:"useSceneNumbering,omitempty"`
	ID                int64          `json:"id,omitempty"`
	LanguageProfileID int64          `json:"languageProfileId,omitempty"`
	QualityProfileID  int64          `json:"qualityProfileId,omitempty"`
	TvdbID            int64          `json:"tvdbId,omitempty"`
	ImdbID            string         `json:"imdbId,omitempty"`
	TvMazeID          int64          `json:"tvMazeId,omitempty"`
	TvRageID          int64          `json:"tvRageId,omitempty"`
	Path              string         `json:"path,omitempty"`
	SeriesType        string         `json:"seriesType,omitempty"`
	Title             string         `json:"title,omitempty"`
	TitleSlug         string         `json:"titleSlug,omitempty"`
	RootFolderPath    string         `json:"rootFolderPath,omitempty"`
	Tags              []int          `json:"tags,omitempty"`
	Seasons           []*Season      `json:"seasons,omitempty"`
	Images            []*starr.Image `json:"images,omitempty"`
	// to be used only on POST, not for PUT
	AddOptions *AddSeriesOptions `json:"addOptions,omitempty"`
}

// Series is the output of /api/v3/series endpoint.
type Series struct {
	Ended             bool              `json:"ended,omitempty"`
	Monitored         bool              `json:"monitored"`
	SeasonFolder      bool              `json:"seasonFolder,omitempty"`
	UseSceneNumbering bool              `json:"useSceneNumbering,omitempty"`
	Runtime           int               `json:"runtime,omitempty"`
	Year              int               `json:"year,omitempty"`
	ID                int64             `json:"id,omitempty"`
	LanguageProfileID int64             `json:"languageProfileId,omitempty"`
	QualityProfileID  int64             `json:"qualityProfileId,omitempty"`
	TvdbID            int64             `json:"tvdbId,omitempty"`
	TvMazeID          int64             `json:"tvMazeId,omitempty"`
	TvRageID          int64             `json:"tvRageId,omitempty"`
	AirTime           string            `json:"airTime,omitempty"`
	Certification     string            `json:"certification,omitempty"`
	CleanTitle        string            `json:"cleanTitle,omitempty"`
	ImdbID            string            `json:"imdbId,omitempty"`
	Network           string            `json:"network,omitempty"`
	Overview          string            `json:"overview,omitempty"`
	Path              string            `json:"path,omitempty"`
	SeriesType        string            `json:"seriesType,omitempty"`
	SortTitle         string            `json:"sortTitle,omitempty"`
	Status            string            `json:"status,omitempty"`
	Title             string            `json:"title,omitempty"`
	TitleSlug         string            `json:"titleSlug,omitempty"`
	RootFolderPath    string            `json:"rootFolderPath,omitempty"`
	Added             time.Time         `json:"added,omitempty"`
	FirstAired        time.Time         `json:"firstAired,omitempty"`
	NextAiring        time.Time         `json:"nextAiring,omitempty"`
	PreviousAiring    time.Time         `json:"previousAiring,omitempty"`
	Ratings           *starr.Ratings    `json:"ratings,omitempty"`
	Statistics        *Statistics       `json:"statistics,omitempty"`
	Tags              []int             `json:"tags,omitempty"`
	Genres            []string          `json:"genres,omitempty"`
	AlternateTitles   []*AlternateTitle `json:"alternateTitles,omitempty"`
	Seasons           []*Season         `json:"seasons,omitempty"`
	Images            []*starr.Image    `json:"images,omitempty"`
}

// AddSeriesOptions is part of AddSeriesInput.
type AddSeriesOptions struct {
	SearchForMissingEpisodes     bool        `json:"searchForMissingEpisodes"`
	SearchForCutoffUnmetEpisodes bool        `json:"searchForCutoffUnmetEpisodes,omitempty"`
	IgnoreEpisodesWithFiles      bool        `json:"ignoreEpisodesWithFiles,omitempty"`
	IgnoreEpisodesWithoutFiles   bool        `json:"ignoreEpisodesWithoutFiles,omitempty"`
	Monitor                      MonitorType `json:"monitor,omitempty"`
}

// MonitorType is part of the AddSeriesOptions.
type MonitorType string

// These are the possible values for the monitor option when adding a new series.
const (
	MonitorUnknown           MonitorType = "unknown"
	MonitorAll               MonitorType = "all"
	MonitorFuture            MonitorType = "future"
	MonitorMissing           MonitorType = "missing"
	MonitorExisting          MonitorType = "existing"
	MonitorFirstSeason       MonitorType = "firstSeason"
	MonitorLastSeason        MonitorType = "lastSeason"
	MonitorLatestSeason      MonitorType = "latestSeason" // obsolete
	MonitorPilot             MonitorType = "pilot"
	MonitorRecent            MonitorType = "recent"
	MonitorMonitorSpecials   MonitorType = "monitorSpecials"
	MonitorUnmonitorSpecials MonitorType = "unmonitorSpecials"
	MonitorNone              MonitorType = "none"
	MonitorSkip              MonitorType = "skip"
)

// AlternateTitle is part of a AddSeriesInput.
type AlternateTitle struct {
	SeasonNumber int    `json:"seasonNumber"`
	Title        string `json:"title"`
}

// Season is part of AddSeriesInput and Queue and used in a few places.
type Season struct {
	Monitored    bool        `json:"monitored"`
	SeasonNumber int         `json:"seasonNumber"`
	Statistics   *Statistics `json:"statistics,omitempty"`
}

// Statistics is part of AddSeriesInput and Queue.
type Statistics struct {
	SeasonCount       int       `json:"seasonCount"`
	EpisodeFileCount  int       `json:"episodeFileCount"`
	EpisodeCount      int       `json:"episodeCount"`
	TotalEpisodeCount int       `json:"totalEpisodeCount"`
	SizeOnDisk        int64     `json:"sizeOnDisk"`
	PercentOfEpisodes float64   `json:"percentOfEpisodes"`
	PreviousAiring    time.Time `json:"previousAiring"`
	NextAiring        time.Time `json:"nextAiring"`
	ReleaseGroups     []string  `json:"releaseGroups"`
}

// GetAllSeries returns all configured series.
// This may not deal well with pagination atm, let us know?
func (s *Sonarr) GetAllSeries() ([]*Series, error) {
	return s.GetAllSeriesContext(context.Background())
}

// GetAllSeriesContext returns all configured series.
// This may not deal well with pagination atm, let us know?
func (s *Sonarr) GetAllSeriesContext(ctx context.Context) ([]*Series, error) {
	return s.GetSeriesContext(ctx, 0)
}

// GetSeries locates and returns a series by tvdbID. If tvdbID is 0, returns all series.
func (s *Sonarr) GetSeries(tvdbID int64) ([]*Series, error) {
	return s.GetSeriesContext(context.Background(), tvdbID)
}

// GetSeriesContext locates and returns a series by tvdbID. If tvdbID is 0, returns all series.
func (s *Sonarr) GetSeriesContext(ctx context.Context, tvdbID int64) ([]*Series, error) {
	var output []*Series

	req := starr.Request{URI: bpSeries, Query: make(url.Values)}
	if tvdbID != 0 {
		req.Query.Add("tvdbId", starr.Itoa(tvdbID))
	}

	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// UpdateSeries updates a series in place.
func (s *Sonarr) UpdateSeries(series *AddSeriesInput, moveFiles bool) (*Series, error) {
	return s.UpdateSeriesContext(context.Background(), series, moveFiles)
}

// UpdateSeriesContext updates a series in place.
func (s *Sonarr) UpdateSeriesContext(ctx context.Context, series *AddSeriesInput, moveFiles bool) (*Series, error) {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(series); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpSeries, err)
	}

	var output Series

	req := starr.Request{
		URI:   path.Join(bpSeries, starr.Itoa(series.ID)),
		Query: make(url.Values),
		Body:  &body,
	}
	req.Query.Add("moveFiles", starr.Itoa(moveFiles))

	if err := s.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// AddSeries adds a new series to Sonarr.
func (s *Sonarr) AddSeries(series *AddSeriesInput) (*Series, error) {
	return s.AddSeriesContext(context.Background(), series)
}

// AddSeriesContext adds a new series to Sonarr.
func (s *Sonarr) AddSeriesContext(ctx context.Context, series *AddSeriesInput) (*Series, error) {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(series); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpSeries, err)
	}

	var output Series

	req := starr.Request{URI: bpSeries, Query: make(url.Values), Body: &body}
	if err := s.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// GetSeriesByID locates and returns a series by DB [series] ID.
func (s *Sonarr) GetSeriesByID(seriesID int64) (*Series, error) {
	return s.GetSeriesByIDContext(context.Background(), seriesID)
}

// GetSeriesByIDContext locates and returns a series by DB [series] ID.
func (s *Sonarr) GetSeriesByIDContext(ctx context.Context, seriesID int64) (*Series, error) {
	var output Series

	req := starr.Request{URI: path.Join(bpSeries, starr.Itoa(seriesID))}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// GetSeriesLookup searches for a series [in Servarr] using a search term or a tvdbid.
// Provide a search term or a tvdbid. If you provide both, tvdbID is used.
func (s *Sonarr) GetSeriesLookup(term string, tvdbID int64) ([]*Series, error) {
	return s.GetSeriesLookupContext(context.Background(), term, tvdbID)
}

// GetSeriesLookupContext searches for a series [in Servarr] using a search term or a tvdbid.
// Provide a search term or a tvdbid. If you provide both, tvdbID is used.
func (s *Sonarr) GetSeriesLookupContext(ctx context.Context, term string, tvdbID int64) ([]*Series, error) {
	var output []*Series

	req := starr.Request{URI: path.Join(bpSeries, "lookup"), Query: make(url.Values)}
	if tvdbID > 0 {
		req.Query.Add("term", "tvdbid:"+starr.Itoa(tvdbID))
	} else {
		req.Query.Add("term", term)
	}

	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// Lookup will search for series matching the specified search term.
// Searches for new shows on TheTVDB.com utilizing sonarr.tv's caching and augmentation proxy.
func (s *Sonarr) Lookup(term string) ([]*Series, error) {
	return s.LookupContext(context.Background(), term)
}

// Lookup will search for series matching the specified search term.
// Searches for new shows on TheTVDB.com utilizing sonarr.tv's caching and augmentation proxy.
func (s *Sonarr) LookupContext(ctx context.Context, term string) ([]*Series, error) {
	return s.GetSeriesLookupContext(ctx, term, 0)
}

// DeleteSeries removes a single Series.
// deleteFiles flag defines the deleteFiles query parameter.
// importExclude defines the addImportListExclusion query parameter.
func (s *Sonarr) DeleteSeries(seriesID int, deleteFiles bool, importExclude bool) error {
	return s.DeleteSeriesContext(context.Background(), seriesID, deleteFiles, importExclude)
}

// DeleteSeries removes a single Series.
// deleteFiles flag defines the deleteFiles query parameter.
// importExclude defines the addImportListExclusion query parameter.
func (s *Sonarr) DeleteSeriesContext(ctx context.Context, seriesID int, deleteFiles bool, importExclude bool) error {
	req := starr.Request{URI: path.Join(bpSeries, starr.Itoa(seriesID)), Query: make(url.Values)}
	req.Query.Add("deleteFiles", starr.Itoa(deleteFiles))
	req.Query.Add("addImportListExclusion", starr.Itoa(importExclude))

	if err := s.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}

// DeleteSeriesDefault defines the behaviour to set deleteFiles to true and addImportListExclusion to false.
func (s *Sonarr) DeleteSeriesDefault(seriesID int) error {
	return s.DeleteSeriesContext(context.Background(), seriesID, true, false)
}
