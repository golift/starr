package radarr

import (
	"crypto/tls"
	"net/http"
	"time"

	"golift.io/starr"
)

type Radarr struct {
	starr.APIer
}

func New(c *starr.Config) *Radarr {
	if c.Client == nil {
		c.Client = &http.Client{ // nolint: exhaustivestruct
			Timeout: c.Timeout.Duration,
			Transport: &http.Transport{ // nolint: exhaustivestruct
				TLSClientConfig: &tls.Config{InsecureSkipVerify: !c.ValidSSL}, // nolint: gosec, exhaustivestruct
			},
		}
	}

	return &Radarr{APIer: c}
}

// AddMovieInput is the input for a new movie.
type AddMovieInput struct {
	Title               string          `json:"title,omitempty"`
	TitleSlug           string          `json:"titleSlug,omitempty"`
	MinimumAvailability string          `json:"minimumAvailability,omitempty"`
	RootFolderPath      string          `json:"rootFolderPath"`
	TmdbID              int             `json:"tmdbId"`
	QualityProfileID    int             `json:"qualityProfileId"`
	ProfileID           int             `json:"profileId,omitempty"`
	Year                int             `json:"year,omitempty"`
	Images              []*starr.Image  `json:"images,omitempty"`
	AddOptions          AddMovieOptions `json:"addOptions"`
	Monitored           bool            `json:"monitored,omitempty"`
}

// AddMovieOptions are the options for finding a new movie.
type AddMovieOptions struct {
	SearchForMovie bool `json:"searchForMovie"`
}

// AddMovieOutput is the data returned when adding a movier.
type AddMovieOutput struct {
	Title                 string         `json:"title"`
	OriginalTitle         string         `json:"originalTitle"`
	AlternateTitles       []interface{}  `json:"alternateTitles"`
	SecondaryYearSourceID int            `json:"secondaryYearSourceId"`
	SortTitle             string         `json:"sortTitle"`
	SizeOnDisk            int            `json:"sizeOnDisk"`
	Status                string         `json:"status"`
	Overview              string         `json:"overview"`
	InCinemas             time.Time      `json:"inCinemas"`
	DigitalRelease        time.Time      `json:"digitalRelease"`
	Images                []*starr.Image `json:"images"`
	Website               string         `json:"website"`
	Year                  int            `json:"year"`
	HasFile               bool           `json:"hasFile"`
	YouTubeTrailerID      string         `json:"youTubeTrailerId"`
	Studio                string         `json:"studio"`
	Path                  string         `json:"path"`
	QualityProfileID      int            `json:"qualityProfileId"`
	Monitored             bool           `json:"monitored"`
	MinimumAvailability   string         `json:"minimumAvailability"`
	IsAvailable           bool           `json:"isAvailable"`
	FolderName            string         `json:"folderName"`
	Runtime               int            `json:"runtime"`
	CleanTitle            string         `json:"cleanTitle"`
	ImdbID                string         `json:"imdbId"`
	TmdbID                int            `json:"tmdbId"`
	TitleSlug             string         `json:"titleSlug"`
	Genres                []string       `json:"genres"`
	Tags                  []interface{}  `json:"tags"`
	Added                 time.Time      `json:"added"`
	AddOptions            struct {
		SearchForMovie             bool `json:"searchForMovie"`
		IgnoreEpisodesWithFiles    bool `json:"ignoreEpisodesWithFiles"`
		IgnoreEpisodesWithoutFiles bool `json:"ignoreEpisodesWithoutFiles"`
	} `json:"addOptions"`
	Ratings struct {
		Votes int     `json:"votes"`
		Value float64 `json:"value"`
	} `json:"ratings"`
	ID int `json:"id"`
}

// RootFolder is the /rootFolder endpoint.
type RootFolder struct {
	Path            string `json:"path"`
	Accessible      bool   `json:"accessible"`
	FreeSpace       int64  `json:"freeSpace"`
	UnmappedFolders []struct {
		Name string `json:"name"`
		Path string `json:"path"`
	} `json:"unmappedFolders"`
	ID int `json:"id"`
}

// History is the /api/history endpoint.
type History struct {
	Page          int       `json:"page"`
	PageSize      int       `json:"pageSize"`
	SortKey       string    `json:"sortKey"`
	SortDirection string    `json:"sortDirection"`
	TotalRecords  int64     `json:"totalRecords"`
	Records       []*Record `json:"Records"`
}

// Record is a record in Radarr History.
type Record struct {
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
		Downloaded       bool           `json:"downloaded"`
		Monitored        bool           `json:"monitored"`
		HasFile          bool           `json:"hasFile"`
		Year             int            `json:"year"`
		ProfileID        int            `json:"profileId"`
		Runtime          int            `json:"runtime"`
		QualityProfileID int            `json:"qualityProfileId"`
		ID               int64          `json:"id"`
		SizeOnDisk       int64          `json:"sizeOnDisk"`
		Title            string         `json:"title"`
		SortTitle        string         `json:"sortTitle"`
		Status           string         `json:"status"`
		Overview         string         `json:"overview"`
		InCinemas        time.Time      `json:"inCinemas"`
		Images           []*starr.Image `json:"images"`
		Website          string         `json:"website"`
		YouTubeTrailerID string         `json:"youTubeTrailerId"`
		Studio           string         `json:"studio"`
		Path             string         `json:"path"`
		LastInfoSync     time.Time      `json:"lastInfoSync"`
		CleanTitle       string         `json:"cleanTitle"`
		ImdbID           string         `json:"imdbId"`
		TmdbID           int64          `json:"tmdbId"`
		TitleSlug        string         `json:"titleSlug"`
		Genres           []string       `json:"genres"`
		Tags             []string       `json:"tags"`
		Added            time.Time      `json:"added"`
		Ratings          struct {
			Votes int64   `json:"votes"`
			Value float64 `json:"value"`
		} `json:"ratings"`
		AlternativeTitles []string `json:"alternativeTitles"`
	} `json:"movie"`
	ID int `json:"id"`
}

// Queue is the /api/v3/queue endpoint.
type Queue struct {
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
		Downloaded            bool           `json:"downloaded"`
		HasFile               bool           `json:"hasFile"`
		Monitored             bool           `json:"monitored"`
		IsAvailable           bool           `json:"isAvailable"`
		SecondaryYearSourceID int            `json:"secondaryYearSourceId"`
		Year                  int            `json:"year"`
		ProfileID             int            `json:"profileId"`
		Runtime               int            `json:"runtime"`
		QualityProfileID      int64          `json:"qualityProfileId"`
		ID                    int64          `json:"id"`
		TmdbID                int64          `json:"tmdbId"`
		SizeOnDisk            int64          `json:"sizeOnDisk"`
		InCinemas             time.Time      `json:"inCinemas"`
		PhysicalRelease       time.Time      `json:"physicalRelease"`
		LastInfoSync          time.Time      `json:"lastInfoSync"`
		Added                 time.Time      `json:"added"`
		Title                 string         `json:"title"`
		SortTitle             string         `json:"sortTitle"`
		Status                string         `json:"status"`
		Overview              string         `json:"overview"`
		Website               string         `json:"website"`
		YouTubeTrailerID      string         `json:"youTubeTrailerId"`
		Studio                string         `json:"studio"`
		Path                  string         `json:"path"`
		PathState             string         `json:"pathState"`
		MinimumAvailability   string         `json:"minimumAvailability"`
		FolderName            string         `json:"folderName"`
		CleanTitle            string         `json:"cleanTitle"`
		ImdbID                string         `json:"imdbId"`
		TitleSlug             string         `json:"titleSlug"`
		Genres                []string       `json:"genres"`
		Tags                  []string       `json:"tags"`
		Images                []*starr.Image `json:"images"`
		Ratings               struct {
			Votes int64   `json:"votes"`
			Value float64 `json:"value"`
		} `json:"ratings"`
	} `json:"movie"`
	Quality        *starr.Quality         `json:"quality"`
	StatusMessages []*starr.StatusMessage `json:"statusMessages"`
}

// Movie is the /api/v3/movie endpoint.
type Movie struct {
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
	SecondaryYearSourceID int            `json:"secondaryYearSourceId"`
	SortTitle             string         `json:"sortTitle"`
	SizeOnDisk            int64          `json:"sizeOnDisk"`
	Status                string         `json:"status"`
	Overview              string         `json:"overview"`
	InCinemas             time.Time      `json:"inCinemas"`
	PhysicalRelease       time.Time      `json:"physicalRelease"`
	DigitalRelease        time.Time      `json:"digitalRelease"`
	Images                []*starr.Image `json:"images"`
	Website               string         `json:"website"`
	Year                  int            `json:"year"`
	HasFile               bool           `json:"hasFile"`
	YouTubeTrailerID      string         `json:"youTubeTrailerId"`
	Studio                string         `json:"studio"`
	Path                  string         `json:"path"`
	QualityProfileID      int            `json:"qualityProfileId"`
	Monitored             bool           `json:"monitored"`
	MinimumAvailability   string         `json:"minimumAvailability"`
	IsAvailable           bool           `json:"isAvailable"`
	FolderName            string         `json:"folderName"`
	Runtime               int            `json:"runtime"`
	CleanTitle            string         `json:"cleanTitle"`
	ImdbID                string         `json:"imdbId"`
	TmdbID                int            `json:"tmdbId"`
	TitleSlug             string         `json:"titleSlug"`
	Certification         string         `json:"certification"`
	Genres                []string       `json:"genres"`
	Tags                  []interface{}  `json:"tags"`
	Added                 time.Time      `json:"added"`
	Ratings               struct {
		Votes int     `json:"votes"`
		Value float64 `json:"value"`
	} `json:"ratings"`
	MovieFile struct {
		MovieID      int            `json:"movieId"`
		RelativePath string         `json:"relativePath"`
		Path         string         `json:"path"`
		Size         int64          `json:"size"`
		DateAdded    time.Time      `json:"dateAdded"`
		SceneName    string         `json:"sceneName"`
		IndexerFlags int            `json:"indexerFlags"`
		Quality      *starr.Quality `json:"quality"`
		MediaInfo    struct {
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
		Name   string         `json:"name"`
		TmdbID int            `json:"tmdbId"`
		Images []*starr.Image `json:"images"`
	} `json:"collection"`
	ID int `json:"id"`
}

type QualityProfile struct {
	Name           string `json:"name"`
	UpgradeAllowed bool   `json:"upgradeAllowed"`
	Cutoff         int    `json:"cutoff"`
	Items          []struct {
		Quality struct {
			ID         int    `json:"id"`
			Name       string `json:"name"`
			Source     string `json:"source"`
			Resolution int    `json:"resolution"`
			Modifier   string `json:"modifier"`
		} `json:"quality,omitempty"`
		Items   []interface{} `json:"items"`
		Allowed bool          `json:"allowed"`
		Name    string        `json:"name,omitempty"`
		ID      int           `json:"id,omitempty"`
	} `json:"items"`
	MinFormatScore    int           `json:"minFormatScore"`
	CutoffFormatScore int           `json:"cutoffFormatScore"`
	FormatItems       []interface{} `json:"formatItems"`
	Language          struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"language"`
	ID int `json:"id"`
}
