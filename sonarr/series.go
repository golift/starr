package sonarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"path"
	"strconv"
	"time"

	"golift.io/starr"
)

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
	SearchForMissingEpisodes     bool `json:"searchForMissingEpisodes"`
	SearchForCutoffUnmetEpisodes bool `json:"searchForCutoffUnmetEpisodes,omitempty"`
	IgnoreEpisodesWithFiles      bool `json:"ignoreEpisodesWithFiles,omitempty"`
	IgnoreEpisodesWithoutFiles   bool `json:"ignoreEpisodesWithoutFiles,omitempty"`
}

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
}

// Define Base Path for Series calls.
const bpSeries = APIver + "/series"

// GetAllSeries returns all configured series.
// This may not deal well with pagination atm.
func (s *Sonarr) GetAllSeries() ([]*Series, error) {
	return s.GetAllSeriesContext(context.Background())
}

func (s *Sonarr) GetAllSeriesContext(ctx context.Context) ([]*Series, error) {
	return s.GetSeriesContext(ctx, 0)
}

// GetSeries locates and returns a series by tvdbID. If tvdbID is 0, returns all series.
func (s *Sonarr) GetSeries(tvdbID int64) ([]*Series, error) {
	return s.GetSeriesContext(context.Background(), tvdbID)
}

func (s *Sonarr) GetSeriesContext(ctx context.Context, tvdbID int64) ([]*Series, error) {
	var output []*Series

	params := make(url.Values)
	if tvdbID != 0 {
		params.Add("tvdbId", strconv.FormatInt(tvdbID, starr.Base10))
	}

	if _, err := s.GetInto(ctx, bpSeries, params, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", bpSeries, err)
	}

	return output, nil
}

// UpdateSeries updates a series in place.
func (s *Sonarr) UpdateSeries(series *AddSeriesInput) (*Series, error) {
	return s.UpdateSeriesContext(context.Background(), series)
}

func (s *Sonarr) UpdateSeriesContext(ctx context.Context, series *AddSeriesInput) (*Series, error) {
	var output *Series

	params := make(url.Values)
	params.Add("moveFiles", "true")

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(series); err != nil {
		return nil, fmt.Errorf("json.Marshal(series): %w", err)
	}

	uri := path.Join(bpSeries, strconv.Itoa(int(series.ID)))
	if _, err := s.PutInto(ctx, uri, params, &body, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", uri, err)
	}

	return output, nil
}

// AddSeries adds a new series to Sonarr.
func (s *Sonarr) AddSeries(series *AddSeriesInput) (*Series, error) {
	return s.AddSeriesContext(context.Background(), series)
}

func (s *Sonarr) AddSeriesContext(ctx context.Context, series *AddSeriesInput) (*Series, error) {
	var output Series

	params := make(url.Values)
	params.Add("moveFiles", "true")

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(series); err != nil {
		return nil, fmt.Errorf("json.Marshal(series): %w", err)
	}

	if _, err := s.PostInto(ctx, bpSeries, params, &body, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", bpSeries, err)
	}

	return &output, nil
}

// GetSeriesByID locates and returns a series by DB [series] ID.
func (s *Sonarr) GetSeriesByID(seriesID int64) (*Series, error) {
	return s.GetSeriesByIDContext(context.Background(), seriesID)
}

func (s *Sonarr) GetSeriesByIDContext(ctx context.Context, seriesID int64) (*Series, error) {
	var output Series

	uri := path.Join(bpSeries, strconv.Itoa(int(seriesID)))
	if _, err := s.GetInto(ctx, uri, nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", uri, err)
	}

	return &output, nil
}

// GetSeriesLookup searches for a series [in Servarr] using a search term or a tvdbid.
// Provide a search term or a tvdbid. If you provide both, tvdbID is used.
func (s *Sonarr) GetSeriesLookup(term string, tvdbID int64) ([]*Series, error) {
	return s.GetSeriesLookupContext(context.Background(), term, tvdbID)
}

func (s *Sonarr) GetSeriesLookupContext(ctx context.Context, term string, tvdbID int64) ([]*Series, error) {
	var output []*Series

	params := make(url.Values)
	if tvdbID > 0 {
		params.Add("term", "tvdbid:"+strconv.FormatInt(tvdbID, starr.Base10))
	} else {
		params.Add("term", term)
	}

	uri := path.Join(bpSeries, "lookup")
	if _, err := s.GetInto(ctx, uri, params, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", uri, err)
	}

	return output, nil
}

// Lookup will search for series matching the specified search term.
// Searches for new shows on TheTVDB.com utilizing sonarr.tv's caching and augmentation proxy.
func (s *Sonarr) Lookup(term string) ([]*Series, error) {
	return s.LookupContext(context.Background(), term)
}

func (s *Sonarr) LookupContext(ctx context.Context, term string) ([]*Series, error) {
	return s.GetSeriesLookupContext(ctx, term, 0)
}

// DeleteSeries removes a single Series.
// deleteFiles flag defines the deleteFiles query paramenter.
// importExclude defines the addImportListExclusion query paramenter.
func (s *Sonarr) DeleteSeries(seriesID int, deleteFiles bool, importExclude bool) error {
	return s.DeleteSeriesContext(context.Background(), seriesID, deleteFiles, importExclude)
}

func (s *Sonarr) DeleteSeriesContext(ctx context.Context, seriesID int, deleteFiles bool, importExclude bool) error {
	params := make(url.Values)
	params.Add("deleteFiles", strconv.FormatBool(deleteFiles))
	params.Add("addImportListExclusion", strconv.FormatBool(importExclude))

	uri := path.Join(bpSeries, strconv.Itoa(seriesID))
	if _, err := s.Delete(ctx, uri, params); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", uri, err)
	}

	return nil
}

// DeleteSeriesDefault defines the behaviour to set deleteFiles to true and addImportListExclusion to false.
func (s *Sonarr) DeleteSeriesDefault(seriesID int) error {
	return s.DeleteSeriesContext(context.Background(), seriesID, true, false)
}
