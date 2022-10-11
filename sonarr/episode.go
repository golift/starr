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

const bpEpisode = APIver + "/episode"

// Episode is the /api/v3/episode endpoint.
type Episode struct {
	ID                       int64          `json:"id"`
	SeriesID                 int64          `json:"seriesId"`
	TvdbID                   int64          `json:"tvdbId"`
	AbsoluteEpisodeNumber    int64          `json:"absoluteEpisodeNumber"`
	EpisodeFileID            int64          `json:"episodeFileId"`
	SeasonNumber             int64          `json:"seasonNumber"`
	EpisodeNumber            int64          `json:"episodeNumber"`
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

// GetSeriesEpisodes returns all episodes for a series by series ID.
// You can get series IDs from GetAllSeries() and GetSeries().
func (s *Sonarr) GetSeriesEpisodes(seriesID int64) ([]*Episode, error) {
	return s.GetSeriesEpisodesContext(context.Background(), seriesID)
}

// GetSeriesEpisodesContext returns all episodes for a series by series ID.
// You can get series IDs from GetAllSeries() and GetSeries().
func (s *Sonarr) GetSeriesEpisodesContext(ctx context.Context, seriesID int64) ([]*Episode, error) {
	var output []*Episode

	req := starr.Request{URI: bpEpisode, Query: make(url.Values)}
	req.Query.Add("seriesId", fmt.Sprint(seriesID))

	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
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
