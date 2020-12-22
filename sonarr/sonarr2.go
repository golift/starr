package sonarr

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"golift.io/starr"
)

/* This is all deprecated and will be removed in the future. Switch to v3. */

// QueueV2 is the /api/queue endpoint.
type QueueV2 struct {
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
		SeasonFolder      bool           `json:"seasonFolder"`
		Monitored         bool           `json:"monitored"`
		UseSceneNumbering bool           `json:"useSceneNumbering"`
		Year              int            `json:"year"`
		ProfileID         int            `json:"profileId"`
		Runtime           int            `json:"runtime"`
		QualityProfileID  int            `json:"qualityProfileId"`
		ID                int            `json:"id"`
		SeasonCount       int            `json:"seasonCount"`
		TvdbID            int64          `json:"tvdbId"`
		TvRageID          int64          `json:"tvRageId"`
		TvMazeID          int64          `json:"tvMazeId"`
		FirstAired        time.Time      `json:"firstAired"`
		LastInfoSync      time.Time      `json:"lastInfoSync"`
		Added             time.Time      `json:"added"`
		Path              string         `json:"path"`
		SeriesType        string         `json:"seriesType"`
		CleanTitle        string         `json:"cleanTitle"`
		ImdbID            string         `json:"imdbId"`
		TitleSlug         string         `json:"titleSlug"`
		Certification     string         `json:"certification"`
		Genres            []string       `json:"genres"`
		Tags              []interface{}  `json:"tags"`
		Title             string         `json:"title"`
		SortTitle         string         `json:"sortTitle"`
		Status            string         `json:"status"`
		Overview          string         `json:"overview"`
		Network           string         `json:"network"`
		AirTime           string         `json:"airTime"`
		Images            []*starr.Image `json:"images"`
		Seasons           []struct {
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
	Quality        *starr.Quality         `json:"quality"`
	StatusMessages []*starr.StatusMessage `json:"statusMessages"`
}

// GetQueueV2 returns the Sonarr Queue.
func (s *Sonarr) GetQueueV2() ([]*QueueV2, error) {
	params := make(url.Values)
	params.Set("sort_by", "timeleft")
	params.Set("order", "asc")

	var queue []*QueueV2
	if err := s.a.GetInto("queue", params, &queue); err != nil {
		return queue, fmt.Errorf("api.Get(queue): %w", err)
	}

	return queue, nil
}

// SystemStatusV2 is the /api/system/status endpoint.
type SystemStatusV2 struct {
	Version           string    `json:"version"`
	BuildTime         time.Time `json:"buildTime"`
	IsDebug           bool      `json:"isDebug"`
	IsProduction      bool      `json:"isProduction"`
	IsAdmin           bool      `json:"isAdmin"`
	IsUserInteractive bool      `json:"isUserInteractive"`
	StartupPath       string    `json:"startupPath"`
	AppData           string    `json:"appData"`
	OsName            string    `json:"osName"`
	OsVersion         string    `json:"osVersion"`
	IsMonoRuntime     bool      `json:"isMonoRuntime"`
	IsMono            bool      `json:"isMono"`
	IsLinux           bool      `json:"isLinux"`
	IsOsx             bool      `json:"isOsx"`
	IsWindows         bool      `json:"isWindows"`
	Branch            string    `json:"branch"`
	Authentication    string    `json:"authentication"`
	SqliteVersion     string    `json:"sqliteVersion"`
	URLBase           string    `json:"urlBase"`
	RuntimeVersion    string    `json:"runtimeVersion"`
	RuntimeName       string    `json:"runtimeName"`
}

// GetSystemStatusV2 returns system status.
func (s *Sonarr) GetSystemStatusV2() (*SystemStatusV2, error) {
	var status *SystemStatusV2
	if err := s.a.GetInto("system/status", nil, status); err != nil {
		return status, fmt.Errorf("api.Get(system/status): %w", err)
	}

	return status, nil
}

// QualityProfileV2 is the /api/profile endpoint.
type QualityProfileV2 struct {
	Name   string `json:"name"`
	Cutoff struct {
		ID         int    `json:"id"`
		Name       string `json:"name"`
		Source     string `json:"source"`
		Resolution int    `json:"resolution"`
	} `json:"cutoff"`
	Items []struct {
		Quality struct {
			ID         int    `json:"id"`
			Name       string `json:"name"`
			Source     string `json:"source"`
			Resolution int    `json:"resolution"`
		} `json:"quality"`
		Allowed bool `json:"allowed"`
	} `json:"items"`
	Language string `json:"language"`
	ID       int    `json:"id"`
}

// GetQualityProfilesV2 returns all configured quality profiles.
func (s *Sonarr) GetQualityProfilesV2() ([]*QualityProfileV2, error) {
	var profiles []*QualityProfileV2
	if err := s.a.GetInto("profile", nil, &profiles); err != nil {
		return nil, fmt.Errorf("api.Get(profile): %w", err)
	}

	return profiles, nil
}

// RootFolderV2 comes from /api/rootfolder path.
type RootFolderV2 struct {
	Path            string `json:"path"`
	FreeSpace       int64  `json:"freeSpace"`
	TotalSpace      int64  `json:"totalSpace"`
	UnmappedFolders []struct {
		Name string `json:"name"`
		Path string `json:"path"`
	} `json:"unmappedFolders"`
	ID int `json:"id"`
}

// RootFoldersV2 returns all configured root folders.
func (s *Sonarr) GetRootFoldersV2() ([]*RootFolderV2, error) {
	var folders []*RootFolderV2
	if err := s.a.GetInto("rootFolder", nil, &folders); err != nil {
		return nil, fmt.Errorf("api.Get(rootFolder): %w", err)
	}

	return folders, nil
}

// SeriesV2 is the /api/series endpoint data.
type SeriesV2 struct {
	Title             string        `json:"title"`
	AlternateTitles   []interface{} `json:"alternateTitles"`
	SortTitle         string        `json:"sortTitle"`
	SeasonCount       int           `json:"seasonCount"`
	TotalEpisodeCount int           `json:"totalEpisodeCount"`
	EpisodeCount      int           `json:"episodeCount"`
	EpisodeFileCount  int           `json:"episodeFileCount"`
	SizeOnDisk        int64         `json:"sizeOnDisk"`
	Status            string        `json:"status"`
	Overview          string        `json:"overview"`
	PreviousAiring    time.Time     `json:"previousAiring,omitempty"`
	Network           string        `json:"network"`
	Images            []struct {
		CoverType string `json:"coverType"`
		URL       string `json:"url"`
	} `json:"images"`
	Seasons           []SeasonV2    `json:"seasons"`
	Year              int           `json:"year"`
	Path              string        `json:"path"`
	ProfileID         int           `json:"profileId"`
	SeasonFolder      bool          `json:"seasonFolder"`
	Monitored         bool          `json:"monitored"`
	UseSceneNumbering bool          `json:"useSceneNumbering"`
	Runtime           int           `json:"runtime"`
	TvdbID            int           `json:"tvdbId"`
	TvRageID          int           `json:"tvRageId"`
	TvMazeID          int           `json:"tvMazeId"`
	FirstAired        time.Time     `json:"firstAired"`
	LastInfoSync      time.Time     `json:"lastInfoSync"`
	SeriesType        string        `json:"seriesType"`
	CleanTitle        string        `json:"cleanTitle"`
	ImdbID            string        `json:"imdbId,omitempty"`
	TitleSlug         string        `json:"titleSlug"`
	Certification     string        `json:"certification,omitempty"`
	Genres            []string      `json:"genres"`
	Tags              []interface{} `json:"tags"`
	Added             time.Time     `json:"added"`
	Ratings           struct {
		Votes int     `json:"votes"`
		Value float64 `json:"value"`
	} `json:"ratings"`
	QualityProfileID int       `json:"qualityProfileId"`
	ID               int       `json:"id"`
	NextAiring       time.Time `json:"nextAiring,omitempty"`
	AirTime          string    `json:"airTime,omitempty"`
}

type SeasonV2 struct {
	SeasonNumber int  `json:"seasonNumber"`
	Monitored    bool `json:"monitored"`
	Statistics   *struct {
		EpisodeFileCount  int     `json:"episodeFileCount"`
		EpisodeCount      int     `json:"episodeCount"`
		TotalEpisodeCount int     `json:"totalEpisodeCount"`
		SizeOnDisk        int     `json:"sizeOnDisk"`
		PercentOfEpisodes float64 `json:"percentOfEpisodes"`
	} `json:"statistics,omitempty"`
}

// GetAllSeriesV2 returns all configured series.
func (s *Sonarr) GetAllSeriesV2() ([]*SeriesV2, error) {
	var series []*SeriesV2
	if err := s.a.GetInto("series", nil, &series); err != nil {
		return nil, fmt.Errorf("api.Get(series): %w", err)
	}

	return series, nil
}

// SeriesLookupV2 is the /api/series/lookup path.
type SeriesLookupV2 struct {
	Title       string `json:"title"`
	SortTitle   string `json:"sortTitle"`
	SeasonCount int    `json:"seasonCount"`
	Status      string `json:"status"`
	Overview    string `json:"overview"`
	Network     string `json:"network"`
	AirTime     string `json:"airTime"`
	Images      []struct {
		CoverType string `json:"coverType"`
		URL       string `json:"url"`
	} `json:"images"`
	RemotePoster      string        `json:"remotePoster"`
	Seasons           []SeasonV2    `json:"seasons"`
	Year              int           `json:"year"`
	ProfileID         int           `json:"profileId"`
	SeasonFolder      bool          `json:"seasonFolder"`
	Monitored         bool          `json:"monitored"`
	UseSceneNumbering bool          `json:"useSceneNumbering"`
	Runtime           int           `json:"runtime"`
	TvdbID            int           `json:"tvdbId"`
	TvRageID          int           `json:"tvRageId"`
	TvMazeID          int           `json:"tvMazeId"`
	FirstAired        time.Time     `json:"firstAired"`
	SeriesType        string        `json:"seriesType"`
	CleanTitle        string        `json:"cleanTitle"`
	ImdbID            string        `json:"imdbId"`
	TitleSlug         string        `json:"titleSlug"`
	Certification     string        `json:"certification"`
	Genres            []string      `json:"genres"`
	Tags              []interface{} `json:"tags"`
	Added             time.Time     `json:"added"`
	Ratings           struct {
		Votes int     `json:"votes"`
		Value float64 `json:"value"`
	} `json:"ratings"`
	QualityProfileID int `json:"qualityProfileId"`
}

// GetSeriesLookupV2 searches for a series using a search term or a tvdbid.
// Provide a search term or a tvdbid. If you provide both, tvdbID is used.
func (s *Sonarr) GetSeriesLookupV2(term string, tvdbID int) ([]*SeriesLookupV2, error) {
	params := make(url.Values)

	if tvdbID > 0 {
		params.Add("term", "tvdbid:"+strconv.Itoa(tvdbID))
	} else {
		params.Add("term", term)
	}

	var series []*SeriesLookupV2
	if err := s.a.GetInto("series/lookup", nil, &series); err != nil {
		return nil, fmt.Errorf("api.Get(series/lookup): %w", err)
	}

	return series, nil
}
