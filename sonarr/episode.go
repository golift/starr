package sonarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"golift.io/starr"
)

// GetSeriesEpisodes returns all episodes for a series by series ID.
// You can get series IDs from GetAllSeries() and GetSeries().
func (s *Sonarr) GetSeriesEpisodes(seriesID int64) ([]*Episode, error) {
	return s.GetSeriesEpisodesContext(context.Background(), seriesID)
}

func (s *Sonarr) GetSeriesEpisodesContext(ctx context.Context, seriesID int64) ([]*Episode, error) {
	var output []*Episode

	params := make(url.Values)
	params.Add("seriesId", strconv.FormatInt(seriesID, starr.Base10))

	err := s.GetInto(ctx, "v3/episode?seriesId", params, &output)
	if err != nil {
		return nil, fmt.Errorf("api.Get(episode): %w", err)
	}

	return output, nil
}

// MonitorEpisode sends a request to monitor (true) or unmonitor (false) a list of episodes by ID.
// You can get episode IDs from GetSeriesEpisodes().
func (s *Sonarr) MonitorEpisode(episodeIDs []int64, monitor bool) ([]*Episode, error) {
	return s.MonitorEpisodeContext(context.Background(), episodeIDs, monitor)
}

func (s *Sonarr) MonitorEpisodeContext(ctx context.Context, episodeIDs []int64, monitor bool) ([]*Episode, error) {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(&struct {
		E []int64 `json:"episodeIds"`
		M bool    `json:"monitored"`
	}{E: episodeIDs, M: monitor}); err != nil {
		return nil, fmt.Errorf("json.Marshal(episodeIDs): %w", err)
	}

	var output []*Episode
	if err := s.PutInto(ctx, "v3/episode/monitor", nil, &body, &output); err != nil {
		return nil, fmt.Errorf("api.Put(episode/monitor): %w", err)
	}

	return output, nil
}
