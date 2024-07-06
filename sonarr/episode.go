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

const bpEpisode = APIver + "/episode"

// Episode is the /api/v3/episode endpoint.
type Episode struct {
	AbsoluteEpisodeNumber    int            `json:"absoluteEpisodeNumber"`
	SeasonNumber             int            `json:"seasonNumber"`
	EpisodeNumber            int            `json:"episodeNumber"`
	ID                       int64          `json:"id"`
	SeriesID                 int64          `json:"seriesId"`
	TvdbID                   int64          `json:"tvdbId"`
	EpisodeFileID            int64          `json:"episodeFileId"`
	AirDateUtc               time.Time      `json:"airDateUtc"`
	AirDate                  string         `json:"airDate"`
	Title                    string         `json:"title"`
	Overview                 string         `json:"overview"`
	UnverifiedSceneNumbering bool           `json:"unverifiedSceneNumbering"`
	HasFile                  bool           `json:"hasFile"`
	Monitored                bool           `json:"monitored"`
	Images                   []*starr.Image `json:"images"`
	Series                   *Series        `json:"series"`
}

// GetEpisode represents the input parameters for an episode api request.
type GetEpisode struct {
	// Set seriesID to get episodes for a specific series. Set to zero to get all episodes.
	SeriesID int64
	// Set seasonNumber to get episodes for a specific season. Set to zero to get all episodes.
	SeasonNumber int
	// Set episodeIds to get episodes for a specific set of ID's. Set to zero to get all episodes.
	EpisodeIDs []int64
	// Set episodeFileId to get episodes for a specific file. Set to zero to get all episodes.
	EpisodeFileID int64
	// Set includeImages to include images for each episode.
	IncludeImages bool
}

// GetSeriesEpisodes returns all episodes matching the provided filters
// You can get series IDs from GetAllSeries() and GetSeries().
func (s *Sonarr) GetSeriesEpisodes(getEpisode *GetEpisode) ([]*Episode, error) {
	return s.GetSeriesEpisodesContext(context.Background(), getEpisode)
}

// GetSeriesEpisodesContext returns all episodes matching the provided filters.
// You can get series IDs from GetAllSeries() and GetSeries().
func (s *Sonarr) GetSeriesEpisodesContext(ctx context.Context, getEpisode *GetEpisode) ([]*Episode, error) {
	var output []*Episode

	params := make(url.Values)

	if getEpisode.SeriesID > 0 {
		params.Set("seriesId", starr.Str(getEpisode.SeriesID))
	}

	if getEpisode.SeasonNumber > 0 {
		params.Set("seasonNumber", strconv.Itoa(getEpisode.SeasonNumber))
	}

	for _, id := range getEpisode.EpisodeIDs {
		params.Add("episodeIds", starr.Str(id))
	}

	if getEpisode.EpisodeFileID > 0 {
		params.Set("episodeFileId", starr.Str(getEpisode.EpisodeFileID))
	}

	if getEpisode.IncludeImages {
		params.Set("includeImages", "true")
	}

	req := starr.Request{URI: bpEpisode, Query: params}

	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetEpisodeByID locates and returns an episode by DB [episode] ID.
func (s *Sonarr) GetEpisodeByID(episodeID int64) (*Episode, error) {
	return s.GetEpisodeByIDContext(context.Background(), episodeID)
}

// GetEpisodeByIDContext locates and returns an episode by DB [episode] ID.
func (s *Sonarr) GetEpisodeByIDContext(ctx context.Context, episodeID int64) (*Episode, error) {
	var output Episode

	req := starr.Request{URI: path.Join(bpEpisode, starr.Str(episodeID))}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// MonitorEpisode sends a request to monitor (true) or unmonitor (false) a list of episodes by ID.
// You can get episode IDs from GetSeriesEpisodes().
func (s *Sonarr) MonitorEpisode(episodeIDs []int64, monitor bool) ([]*Episode, error) {
	return s.MonitorEpisodeContext(context.Background(), episodeIDs, monitor)
}

// MonitorEpisodeContext sends a request to monitor (true) or unmonitor (false) a list of episodes by ID.
// You can get episode IDs from GetSeriesEpisodes().
func (s *Sonarr) MonitorEpisodeContext(ctx context.Context, episodeIDs []int64, monitor bool) ([]*Episode, error) {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(&struct {
		E []int64 `json:"episodeIds"`
		M bool    `json:"monitored"`
	}{E: episodeIDs, M: monitor}); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpEpisode, err)
	}

	var output []*Episode

	req := starr.Request{URI: path.Join(bpEpisode, "monitor"), Body: &body}
	if err := s.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return output, nil
}
