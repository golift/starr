package starr

// Radarr v3 structs

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

// AddMovie is the input for a new movie.
type AddMovie struct {
	Title               string          `json:"title"`
	TitleSlug           string          `json:"titleSlug"`
	TmdbID              int             `json:"tmdbId"`
	Images              interface{}     `json:"images"`
	Monitored           bool            `json:"monitored"`
	QualityProfileID    int             `json:"qualityProfileId"`
	ProfileID           int             `json:"profileId"`
	Year                int             `json:"year"`
	MinimumAvailability string          `json:"minimumAvailability"`
	RootFolderPath      string          `json:"rootFolderPath"`
	AddMovieOptions     AddMovieOptions `json:"addOptions"`
}

// AddMovieOptions are the options for finding a new movie.
type AddMovieOptions struct {
	SearchForMovie bool `json:"searchForMovie"`
}

// Radar3History is the /api/history endpoint.
type Radar3History struct {
	Page          int             `json:"page"`
	PageSize      int             `json:"pageSize"`
	SortKey       string          `json:"sortKey"`
	SortDirection string          `json:"sortDirection"`
	TotalRecords  int64           `json:"totalRecords"`
	Records       []*Radar3Record `json:"Records"`
}

// Radar3Record is a record in Radarr History
type Radar3Record struct {
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

// Radar3Queue is the /api/v3/queue endpoint.
type Radar3Queue struct {
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
		Downloaded            bool      `json:"downloaded"`
		HasFile               bool      `json:"hasFile"`
		Monitored             bool      `json:"monitored"`
		IsAvailable           bool      `json:"isAvailable"`
		SecondaryYearSourceID int       `json:"secondaryYearSourceId"`
		Year                  int       `json:"year"`
		ProfileID             int       `json:"profileId"`
		Runtime               int       `json:"runtime"`
		QualityProfileID      int64     `json:"qualityProfileId"`
		ID                    int64     `json:"id"`
		TmdbID                int64     `json:"tmdbId"`
		SizeOnDisk            int64     `json:"sizeOnDisk"`
		InCinemas             time.Time `json:"inCinemas"`
		PhysicalRelease       time.Time `json:"physicalRelease"`
		LastInfoSync          time.Time `json:"lastInfoSync"`
		Added                 time.Time `json:"added"`
		Title                 string    `json:"title"`
		SortTitle             string    `json:"sortTitle"`
		Status                string    `json:"status"`
		Overview              string    `json:"overview"`
		Website               string    `json:"website"`
		YouTubeTrailerID      string    `json:"youTubeTrailerId"`
		Studio                string    `json:"studio"`
		Path                  string    `json:"path"`
		PathState             string    `json:"pathState"`
		MinimumAvailability   string    `json:"minimumAvailability"`
		FolderName            string    `json:"folderName"`
		CleanTitle            string    `json:"cleanTitle"`
		ImdbID                string    `json:"imdbId"`
		TitleSlug             string    `json:"titleSlug"`
		Genres                []string  `json:"genres"`
		Tags                  []string  `json:"tags"`
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

// Radar3Movie is the /api/v3/movie endpoint.
type Radar3Movie struct {
	Title           string `json:"title"`
	OriginalTitle   string `json:"originalTitle"`
	AlternateTitles []struct {
		SourceType string `json:"sourceType"`
		MovieID    int    `json:"movieId"`
		Title      string `json:"title"`
		SourceID   int    `json:"sourceId"`
		Votes      int    `json:"votes"`
		VoteCount  int    `json:"voteCount"`
		Language   struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"language"`
		ID int `json:"id"`
	} `json:"alternateTitles"`
	SecondaryYearSourceID int       `json:"secondaryYearSourceId"`
	SortTitle             string    `json:"sortTitle"`
	SizeOnDisk            int64     `json:"sizeOnDisk"`
	Status                string    `json:"status"`
	Overview              string    `json:"overview"`
	InCinemas             time.Time `json:"inCinemas"`
	PhysicalRelease       time.Time `json:"physicalRelease"`
	DigitalRelease        time.Time `json:"digitalRelease"`
	Images                []struct {
		CoverType string `json:"coverType"`
		URL       string `json:"url"`
		RemoteURL string `json:"remoteUrl"`
	} `json:"images"`
	Website             string        `json:"website"`
	Year                int           `json:"year"`
	HasFile             bool          `json:"hasFile"`
	YouTubeTrailerID    string        `json:"youTubeTrailerId"`
	Studio              string        `json:"studio"`
	Path                string        `json:"path"`
	QualityProfileID    int           `json:"qualityProfileId"`
	Monitored           bool          `json:"monitored"`
	MinimumAvailability string        `json:"minimumAvailability"`
	IsAvailable         bool          `json:"isAvailable"`
	FolderName          string        `json:"folderName"`
	Runtime             int           `json:"runtime"`
	CleanTitle          string        `json:"cleanTitle"`
	ImdbID              string        `json:"imdbId"`
	TmdbID              int           `json:"tmdbId"`
	TitleSlug           string        `json:"titleSlug"`
	Certification       string        `json:"certification"`
	Genres              []string      `json:"genres"`
	Tags                []interface{} `json:"tags"`
	Added               time.Time     `json:"added"`
	Ratings             struct {
		Votes int     `json:"votes"`
		Value float64 `json:"value"`
	} `json:"ratings"`
	MovieFile struct {
		MovieID      int       `json:"movieId"`
		RelativePath string    `json:"relativePath"`
		Path         string    `json:"path"`
		Size         int64     `json:"size"`
		DateAdded    time.Time `json:"dateAdded"`
		SceneName    string    `json:"sceneName"`
		IndexerFlags int       `json:"indexerFlags"`
		Quality      struct {
			Quality struct {
				ID         int    `json:"id"`
				Name       string `json:"name"`
				Source     string `json:"source"`
				Resolution int    `json:"resolution"`
				Modifier   string `json:"modifier"`
			} `json:"quality"`
			Revision struct {
				Version  int  `json:"version"`
				Real     int  `json:"real"`
				IsRepack bool `json:"isRepack"`
			} `json:"revision"`
		} `json:"quality"`
		MediaInfo struct {
			AudioAdditionalFeatures string  `json:"audioAdditionalFeatures"`
			AudioBitrate            int     `json:"audioBitrate"`
			AudioChannels           float64 `json:"audioChannels"`
			AudioCodec              string  `json:"audioCodec"`
			AudioLanguages          string  `json:"audioLanguages"`
			AudioStreamCount        int     `json:"audioStreamCount"`
			VideoBitDepth           int     `json:"videoBitDepth"`
			VideoBitrate            int     `json:"videoBitrate"`
			VideoCodec              string  `json:"videoCodec"`
			VideoFps                float64 `json:"videoFps"`
			Resolution              string  `json:"resolution"`
			RunTime                 string  `json:"runTime"`
			ScanType                string  `json:"scanType"`
			Subtitles               string  `json:"subtitles"`
		} `json:"mediaInfo"`
		QualityCutoffNotMet bool `json:"qualityCutoffNotMet"`
		Languages           []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"languages"`
		ReleaseGroup string `json:"releaseGroup"`
		Edition      string `json:"edition"`
		ID           int    `json:"id"`
	} `json:"movieFile"`
	Collection struct {
		Name   string        `json:"name"`
		TmdbID int           `json:"tmdbId"`
		Images []interface{} `json:"images"`
	} `json:"collection"`
	ID int `json:"id"`
}

// Radarr3History returns the Radarr History (grabs/failures/completed)
func (c *Config) Radarr3History() ([]*Radar3Record, error) {
	var history Radar3History

	params := make(url.Values)

	params.Set("sortKey", "date")
	params.Set("sortDir", "asc")
	params.Set("page", "1")
	params.Set("pageSize", "0")

	rawJSON, err := c.Req("v3/history", params)
	if err != nil {
		return nil, fmt.Errorf("c.Req(queue): %w", err)
	}

	if err = json.Unmarshal(rawJSON, &history); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(response): %w", err)
	}

	return history.Records, nil
}

// Radarr3Queue returns the Radarr Queue (processing, but not yet imported)
func (c *Config) Radarr3Queue() ([]*Radar3Queue, error) {
	var queue []*Radar3Queue

	params := make(url.Values)

	params.Set("sort_by", "timeleft")
	params.Set("order", "asc")

	rawJSON, err := c.Req("v3/queue", params)
	if err != nil {
		return nil, fmt.Errorf("c.Req(queue): %w", err)
	}

	if err = json.Unmarshal(rawJSON, &queue); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(response): %w", err)
	}

	return queue, nil
}

// Radarr3Movie grabs a movie from the queue, or all movies if tmdbId is empty.
func (c *Config) Radarr3Movie(tmdbID int) ([]*Radar3Movie, error) {
	var movie []*Radar3Movie

	params := make(url.Values)

	params.Set("tmdbId", strconv.Itoa(tmdbID))

	rawJSON, err := c.Req("v3/movie", params)
	if err != nil {
		return nil, fmt.Errorf("c.Req(movie): %w", err)
	}

	if err = json.Unmarshal(rawJSON, &movie); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(response): %w", err)
	}

	return movie, nil
}

// Radarr3AddMovie adds a movie to the queue.
func (c *Config) Radarr3AddMovie(movie *AddMovie) error {
	body, err := json.Marshal(movie)
	if err != nil {
		return fmt.Errorf("json.Marshal(movie): %w", err)
	}

	if _, err = c.Req("v3/movie", nil, body...); err != nil {
		return fmt.Errorf("c.Req(movie): %w", err)
	}

	return nil
}
