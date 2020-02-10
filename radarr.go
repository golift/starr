package starr

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

// RadarHistory is the /api/history endpoint.
type RadarHistory struct {
	Page          int            `json:"page"`
	PageSize      int            `json:"pageSize"`
	SortKey       string         `json:"sortKey"`
	SortDirection string         `json:"sortDirection"`
	TotalRecords  int64          `json:"totalRecords"`
	Records       []*RadarRecord `json:"Records"`
}

// RadarRecord is a record in Radarr History
type RadarRecord struct {
	EpisodeID   int64  `json:"episodeId"`
	MovieID     int64  `json:"movieId"`
	SeriesID    int64  `json:"seriesId"`
	SourceTitle string `json:"sourceTitle"`
	Quality     struct {
		Quality struct {
			ID   int64  `json:"id"`
			Name string `json:"name"`
		} `json:"quality"`
		Revision struct {
			Version int64 `json:"version"`
			Real    int64 `json:"real"`
		} `json:"revision"`
	} `json:"quality"`
	QualityCutoffNotMet bool      `json:"qualityCutoffNotMet"`
	Date                time.Time `json:"date"`
	DownloadID          string    `json:"downloadId"`
	EventType           string    `json:"eventType"`
	Data                struct {
		Indexer         string    `json:"indexer"`
		NzbInfoURL      string    `json:"nzbInfoUrl"`
		ReleaseGroup    string    `json:"releaseGroup"`
		Age             string    `json:"age"`
		AgeHours        string    `json:"ageHours"`
		AgeMinutes      string    `json:"ageMinutes"`
		PublishedDate   time.Time `json:"publishedDate"`
		DownloadClient  string    `json:"downloadClient"`
		Size            string    `json:"size"`
		DownloadURL     string    `json:"downloadUrl"`
		GUID            string    `json:"guid"`
		TvdbID          string    `json:"tvdbId"`
		TvRageID        string    `json:"tvRageId"`
		Protocol        string    `json:"protocol"`
		TorrentInfoHash []string  `json:"torrentInfoHash"`
	} `json:"data"`
	Movie struct {
		Downloaded       bool      `json:"downloaded"`
		Monitored        bool      `json:"monitored"`
		HasFile          bool      `json:"hasFile"`
		Year             int       `json:"year"`
		ProfileID        int       `json:"profileId"`
		Runtime          int       `json:"runtime"`
		QualityProfileID int       `json:"qualityProfileId"`
		ID               int64     `json:"id"`
		SizeOnDisk       int64     `json:"sizeOnDisk"`
		Title            string    `json:"title"`
		SortTitle        string    `json:"sortTitle"`
		Status           string    `json:"status"`
		Overview         string    `json:"overview"`
		InCinemas        time.Time `json:"inCinemas"`
		Images           []struct {
			CoverType string `json:"coverType"`
			URL       string `json:"url"`
		} `json:"images"`
		Website          string    `json:"website"`
		YouTubeTrailerID string    `json:"youTubeTrailerId"`
		Studio           string    `json:"studio"`
		Path             string    `json:"path"`
		LastInfoSync     time.Time `json:"lastInfoSync"`
		CleanTitle       string    `json:"cleanTitle"`
		ImdbID           string    `json:"imdbId"`
		TmdbID           int64     `json:"tmdbId"`
		TitleSlug        string    `json:"titleSlug"`
		Genres           []string  `json:"genres"`
		Tags             []string  `json:"tags"`
		Added            time.Time `json:"added"`
		Ratings          struct {
			Votes int64   `json:"votes"`
			Value float64 `json:"value"`
		} `json:"ratings"`
		AlternativeTitles []string `json:"alternativeTitles"`
	} `json:"movie"`
	ID int `json:"id"`
}

// RadarQueue is the /api/queue endpoint.
type RadarQueue struct {
	ID                      int64     `json:"id"`
	Size                    float64   `json:"size"`
	Sizeleft                float64   `json:"sizeleft"`
	EstimatedCompletionTime time.Time `json:"estimatedCompletionTime"`
	Title                   string    `json:"title"`
	Timeleft                string    `json:"timeleft"`
	Status                  string    `json:"status"`
	TrackedDownloadStatus   string    `json:"trackedDownloadStatus"`
	DownloadID              string    `json:"downloadId"`
	Protocol                string    `json:"protocol"`
	Movie                   struct {
		Downloaded            bool          `json:"downloaded"`
		HasFile               bool          `json:"hasFile"`
		Monitored             bool          `json:"monitored"`
		IsAvailable           bool          `json:"isAvailable"`
		SecondaryYearSourceID int           `json:"secondaryYearSourceId"`
		Year                  int           `json:"year"`
		ProfileID             int           `json:"profileId"`
		Runtime               int           `json:"runtime"`
		QualityProfileID      int64         `json:"qualityProfileId"`
		ID                    int64         `json:"id"`
		TmdbID                int64         `json:"tmdbId"`
		SizeOnDisk            int64         `json:"sizeOnDisk"`
		InCinemas             time.Time     `json:"inCinemas"`
		PhysicalRelease       time.Time     `json:"physicalRelease"`
		LastInfoSync          time.Time     `json:"lastInfoSync"`
		Added                 time.Time     `json:"added"`
		Title                 string        `json:"title"`
		SortTitle             string        `json:"sortTitle"`
		Status                string        `json:"status"`
		Overview              string        `json:"overview"`
		Website               string        `json:"website"`
		YouTubeTrailerID      string        `json:"youTubeTrailerId"`
		Studio                string        `json:"studio"`
		Path                  string        `json:"path"`
		PathState             string        `json:"pathState"`
		MinimumAvailability   string        `json:"minimumAvailability"`
		FolderName            string        `json:"folderName"`
		CleanTitle            string        `json:"cleanTitle"`
		ImdbID                string        `json:"imdbId"`
		TitleSlug             string        `json:"titleSlug"`
		Genres                []string      `json:"genres"`
		Tags                  []string      `json:"tags"`
		AlternativeTitles     []interface{} `json:"alternativeTitles"`
		Images                []struct {
			CoverType string `json:"coverType"`
			URL       string `json:"url"`
		} `json:"images"`
		Ratings struct {
			Votes int64   `json:"votes"`
			Value float64 `json:"value"`
		} `json:"ratings"`
	} `json:"movie"`
	Quality struct {
		Quality struct {
			ID   int64  `json:"id"`
			Name string `json:"name"`
		} `json:"quality"`
		Revision struct {
			Version int64 `json:"version"`
			Real    int64 `json:"real"`
		} `json:"revision"`
	} `json:"quality"`
	StatusMessages []struct {
		Title    string   `json:"title"`
		Messages []string `json:"messages"`
	} `json:"statusMessages"`
}

// RadarrHistory returns the Radarr History (grabs/failures/completed)
func (c *Config) RadarrHistory() ([]*RadarRecord, error) {
	var history RadarHistory

	params := make(url.Values)

	params.Set("sortKey", "date")
	params.Set("sortDir", "asc")
	params.Set("page", "1")
	params.Set("pageSize", "0")

	rawJSON, err := c.Req("history", params)
	if err != nil {
		return nil, fmt.Errorf("c.Req(queue): %v", err)
	}

	if err = json.Unmarshal(rawJSON, &history); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(response): %v", err)
	}

	return history.Records, nil
}

// RadarrQueue returns the Radarr Queue (processing, but not yet imported)
func (c *Config) RadarrQueue() ([]*RadarQueue, error) {
	var queue []*RadarQueue

	params := make(url.Values)

	params.Set("sort_by", "timeleft")
	params.Set("order", "asc")

	rawJSON, err := c.Req("queue", params)
	if err != nil {
		return nil, fmt.Errorf("c.Req(queue): %v", err)
	}

	if err = json.Unmarshal(rawJSON, &queue); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(response): %v", err)
	}

	return queue, nil
}
