package starr

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

// SonarQueue is the /api/queue endpoint.
type SonarQueue struct {
	ID                      int64     `json:"id"`
	EstimatedCompletionTime time.Time `json:"estimatedCompletionTime"`
	Size                    float64   `json:"size"`
	Sizeleft                float64   `json:"sizeleft"`
	Title                   string    `json:"title"`
	Timeleft                string    `json:"timeleft"`
	Status                  string    `json:"status"`
	TrackedDownloadStatus   string    `json:"trackedDownloadStatus"`
	DownloadID              string    `json:"downloadId"`
	Protocol                string    `json:"protocol"`
	Series                  struct {
		SeasonFolder      bool          `json:"seasonFolder"`
		Monitored         bool          `json:"monitored"`
		UseSceneNumbering bool          `json:"useSceneNumbering"`
		Year              int           `json:"year"`
		ProfileID         int           `json:"profileId"`
		Runtime           int           `json:"runtime"`
		QualityProfileID  int           `json:"qualityProfileId"`
		ID                int           `json:"id"`
		SeasonCount       int           `json:"seasonCount"`
		TvdbID            int64         `json:"tvdbId"`
		TvRageID          int64         `json:"tvRageId"`
		TvMazeID          int64         `json:"tvMazeId"`
		FirstAired        time.Time     `json:"firstAired"`
		LastInfoSync      time.Time     `json:"lastInfoSync"`
		Added             time.Time     `json:"added"`
		Path              string        `json:"path"`
		SeriesType        string        `json:"seriesType"`
		CleanTitle        string        `json:"cleanTitle"`
		ImdbID            string        `json:"imdbId"`
		TitleSlug         string        `json:"titleSlug"`
		Certification     string        `json:"certification"`
		Genres            []string      `json:"genres"`
		Tags              []interface{} `json:"tags"` // used to []string and now []int (sonarr v3)
		Title             string        `json:"title"`
		SortTitle         string        `json:"sortTitle"`
		Status            string        `json:"status"`
		Overview          string        `json:"overview"`
		Network           string        `json:"network"`
		AirTime           string        `json:"airTime"`
		Images            []struct {
			CoverType string `json:"coverType"`
			URL       string `json:"url"`
		} `json:"images"`
		Seasons []struct {
			SeasonNumber int  `json:"seasonNumber"`
			Monitored    bool `json:"monitored"`
		} `json:"seasons"`
		Ratings struct {
			Votes int64   `json:"votes"`
			Value float64 `json:"value"`
		} `json:"ratings"`
	} `json:"series"`
	Episode struct {
		HasFile                  bool      `json:"hasFile"`
		Monitored                bool      `json:"monitored"`
		UnverifiedSceneNumbering bool      `json:"unverifiedSceneNumbering"`
		SeriesID                 int       `json:"seriesId"`
		EpisodeFileID            int       `json:"episodeFileId"`
		SeasonNumber             int       `json:"seasonNumber"`
		EpisodeNumber            int       `json:"episodeNumber"`
		AbsoluteEpisodeNumber    int       `json:"absoluteEpisodeNumber"`
		ID                       int64     `json:"id"`
		AirDateUtc               time.Time `json:"airDateUtc"`
		Title                    string    `json:"title"`
		AirDate                  string    `json:"airDate"`
		Overview                 string    `json:"overview"`
	} `json:"episode"`
	Quality struct {
		Quality struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"quality"`
		Revision struct {
			Version int `json:"version"`
			Real    int `json:"real"`
		} `json:"revision"`
	} `json:"quality"`

	StatusMessages []struct {
		Title    string   `json:"title"`
		Messages []string `json:"messages"`
	} `json:"statusMessages"`
}

// SonarrQueue returns the Sonarr Queue
func (c *Config) SonarrQueue() ([]*SonarQueue, error) {
	var queue []*SonarQueue

	params := make(url.Values)

	params.Set("sort_by", "timeleft")
	params.Set("order", "asc")

	rawJSON, err := c.Req("queue", params)
	if err != nil {
		return queue, fmt.Errorf("c.Req(queue): %v", err)
	}

	if err = json.Unmarshal(rawJSON, &queue); err != nil {
		return queue, fmt.Errorf("json.Unmarshal(response): %v", err)
	}

	return queue, nil
}
