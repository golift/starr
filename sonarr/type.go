package sonarr

import (
	"crypto/tls"
	"net/http"
	"time"

	"golift.io/starr"
)

type Sonarr struct {
	starr.APIer
}

func New(c *starr.Config) *Sonarr {
	if c.Client == nil {
		c.Client = &http.Client{ // nolint: exhaustivestruct
			Timeout: c.Timeout.Duration,
			Transport: &http.Transport{ // nolint: exhaustivestruct
				TLSClientConfig: &tls.Config{InsecureSkipVerify: !c.ValidSSL}, // nolint: gosec, exhaustivestruct
			},
		}
	}

	return &Sonarr{APIer: c}
}

// QualityProfile is the /api/v3/qualityprofile endpoint.
type QualityProfile struct {
	Name           string           `json:"name"`
	UpgradeAllowed bool             `json:"upgradeAllowed"`
	Cutoff         int              `json:"cutoff"`
	Qualities      []*starr.Quality `json:"items"`
	ID             int              `json:"id"`
}

// SystemStatus is the /api/v3/system/status endpoint.
type SystemStatus struct {
	Version                string    `json:"version"`
	BuildTime              time.Time `json:"buildTime"`
	IsDebug                bool      `json:"isDebug"`
	IsProduction           bool      `json:"isProduction"`
	IsAdmin                bool      `json:"isAdmin"`
	IsUserInteractive      bool      `json:"isUserInteractive"`
	StartupPath            string    `json:"startupPath"`
	AppData                string    `json:"appData"`
	OsName                 string    `json:"osName"`
	OsVersion              string    `json:"osVersion"`
	IsMonoRuntime          bool      `json:"isMonoRuntime"`
	IsMono                 bool      `json:"isMono"`
	IsLinux                bool      `json:"isLinux"`
	IsOsx                  bool      `json:"isOsx"`
	IsWindows              bool      `json:"isWindows"`
	Mode                   string    `json:"mode"`
	Branch                 string    `json:"branch"`
	Authentication         string    `json:"authentication"`
	SqliteVersion          string    `json:"sqliteVersion"`
	URLBase                string    `json:"urlBase"`
	RuntimeVersion         string    `json:"runtimeVersion"`
	RuntimeName            string    `json:"runtimeName"`
	StartTime              time.Time `json:"startTime"`
	PackageVersion         string    `json:"packageVersion"`
	PackageAuthor          string    `json:"packageAuthor"`
	PackageUpdateMechanism string    `json:"packageUpdateMechanism"`
}

// RootFolder is the /api/v3/rootfolder endpoint.
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

// Queue is the /api/v3/queue endpoint.
type Queue struct {
	Page          int            `json:"page"`
	PageSize      int            `json:"pageSize"`
	SortKey       string         `json:"sortKey"`
	SortDirection string         `json:"sortDirection"`
	TotalRecords  int            `json:"totalRecords"`
	Records       []*QueueRecord `json:"records"`
}

// QueueRecord is part of Queue.
type QueueRecord struct {
	SeriesID                int                    `json:"seriesId"`
	EpisodeID               int                    `json:"episodeId"`
	Language                *starr.Value           `json:"language"`
	Quality                 *starr.Quality         `json:"quality"`
	Size                    float64                `json:"size"`
	Title                   string                 `json:"title"`
	Sizeleft                float64                `json:"sizeleft"`
	Timeleft                string                 `json:"timeleft"`
	EstimatedCompletionTime time.Time              `json:"estimatedCompletionTime"`
	Status                  string                 `json:"status"`
	TrackedDownloadStatus   string                 `json:"trackedDownloadStatus"`
	TrackedDownloadState    string                 `json:"trackedDownloadState"`
	StatusMessages          []*starr.StatusMessage `json:"statusMessages"`
	DownloadID              string                 `json:"downloadId"`
	Protocol                string                 `json:"protocol"`
	DownloadClient          string                 `json:"downloadClient"`
	Indexer                 string                 `json:"indexer"`
	OutputPath              string                 `json:"outputPath"`
	ID                      int                    `json:"id"`
}

// Series the /api/v3/series endpoint.
type Series struct {
	ID                int            `json:"id"`
	Title             string         `json:"title"`
	AlternateTitles   []interface{}  `json:"alternateTitles"`
	SortTitle         string         `json:"sortTitle"`
	Status            string         `json:"status"`
	Overview          string         `json:"overview"`
	PreviousAiring    time.Time      `json:"previousAiring,omitempty"`
	Network           string         `json:"network"`
	Images            []*starr.Image `json:"images"`
	Seasons           []*Season      `json:"seasons"`
	Year              int            `json:"year"`
	Path              string         `json:"path"`
	QualityProfileID  int            `json:"qualityProfileId"`
	LanguageProfileID int            `json:"languageProfileId"`
	Runtime           int            `json:"runtime"`
	TvdbID            int            `json:"tvdbId"`
	TvRageID          int            `json:"tvRageId"`
	TvMazeID          int            `json:"tvMazeId"`
	FirstAired        time.Time      `json:"firstAired"`
	SeriesType        string         `json:"seriesType"`
	CleanTitle        string         `json:"cleanTitle"`
	ImdbID            string         `json:"imdbId,omitempty"`
	TitleSlug         string         `json:"titleSlug"`
	RootFolderPath    string         `json:"rootFolderPath"`
	Certification     string         `json:"certification,omitempty"`
	Genres            []string       `json:"genres"`
	Tags              []interface{}  `json:"tags"`
	Added             time.Time      `json:"added"`
	Ratings           *starr.Ratings `json:"ratings"`
	Statistics        *Statistics    `json:"statistics"`
	NextAiring        time.Time      `json:"nextAiring,omitempty"`
	AirTime           string         `json:"airTime,omitempty"`
	Ended             bool           `json:"ended"`
	SeasonFolder      bool           `json:"seasonFolder"`
	Monitored         bool           `json:"monitored"`
	UseSceneNumbering bool           `json:"useSceneNumbering"`
}

// Statistics is part of Queue.
type Statistics struct {
	SeasonCount       int       `json:"seasonCount"`
	PreviousAiring    time.Time `json:"previousAiring"`
	EpisodeFileCount  int       `json:"episodeFileCount"`
	EpisodeCount      int       `json:"episodeCount"`
	TotalEpisodeCount int       `json:"totalEpisodeCount"`
	SizeOnDisk        int64     `json:"sizeOnDisk"`
	PercentOfEpisodes float64   `json:"percentOfEpisodes"`
}

// Season is part of Queue and used in a few places.
type Season struct {
	SeasonNumber int         `json:"seasonNumber"`
	Monitored    bool        `json:"monitored"`
	Statistics   *Statistics `json:"statistics,omitempty"`
}

// SeriesLookup is the /api/v3/series/lookup endpoint.
type SeriesLookup struct {
	Title             string         `json:"title"`
	SortTitle         string         `json:"sortTitle"`
	Status            string         `json:"status"`
	Overview          string         `json:"overview"`
	Network           string         `json:"network"`
	AirTime           string         `json:"airTime"`
	Images            []*starr.Image `json:"images"`
	RemotePoster      string         `json:"remotePoster"`
	Seasons           []*Season      `json:"seasons"`
	Year              int            `json:"year"`
	QualityProfileID  int            `json:"qualityProfileId"`
	LanguageProfileID int            `json:"languageProfileId"`
	Runtime           int            `json:"runtime"`
	TvdbID            int            `json:"tvdbId"`
	TvRageID          int            `json:"tvRageId"`
	TvMazeID          int            `json:"tvMazeId"`
	FirstAired        time.Time      `json:"firstAired"`
	SeriesType        string         `json:"seriesType"`
	CleanTitle        string         `json:"cleanTitle"`
	ImdbID            string         `json:"imdbId"`
	TitleSlug         string         `json:"titleSlug"`
	Folder            string         `json:"folder"`
	Certification     string         `json:"certification"`
	Genres            []string       `json:"genres"`
	Tags              []interface{}  `json:"tags"`
	Added             time.Time      `json:"added"`
	Ratings           *starr.Ratings `json:"ratings"`
	Statistics        *Statistics    `json:"statistics"`
	Ended             bool           `json:"ended"`
	SeasonFolder      bool           `json:"seasonFolder"`
	Monitored         bool           `json:"monitored"`
	UseSceneNumbering bool           `json:"useSceneNumbering"`
}

// LanguageProfile is the /api/v3/languageprofile endpoint.
type LanguageProfile struct {
	Name           string       `json:"name"`
	UpgradeAllowed bool         `json:"upgradeAllowed"`
	Cutoff         *starr.Value `json:"cutoff"`
	Languages      Languages    `json:"languages"`
	ID             int          `json:"id"`
}

// Language is part of LanguageProfile.
type Languages []struct {
	Language *starr.Value `json:"language"`
	Allowed  bool         `json:"allowed"`
}

// AddSeriesInput is the input for a POST to the /api/v3/series endpoint.
type AddSeriesInput struct {
	ID                int               `json:"id,omitempty"`
	TvdbID            int               `json:"tvdbId"`
	QualityProfileID  int               `json:"qualityProfileId"`
	LanguageProfileID int               `json:"languageProfileID"`
	RootFolderPath    string            `json:"rootFolderPath"`
	Title             string            `json:"title,omitempty"`
	SeriesType        string            `json:"seriesType,omitempty"`
	Seasons           []*Season         `json:"seasons"`
	AddOptions        *AddSeriesOptions `json:"addOptions"`
	SeasonFolder      bool              `json:"seasonFolder"`
	Monitored         bool              `json:"monitored"`
}

// AddSeriesOptions is part of AddSeriesInput.
type AddSeriesOptions struct {
	SearchForMissingEpisodes     bool `json:"searchForMissingEpisodes"`
	SearchForCutoffUnmetEpisodes bool `json:"searchForCutoffUnmetEpisodes,omitempty"`
	IgnoreEpisodesWithFiles      bool `json:"ignoreEpisodesWithFiles,omitempty"`
	IgnoreEpisodesWithoutFiles   bool `json:"ignoreEpisodesWithoutFiles,omitempty"`
}

// AddSeriesOutput is currently an unknown format.
type AddSeriesOutput struct {
	ID                int               `json:"id"`
	Title             string            `json:"title"`
	AlternateTitles   []interface{}     `json:"alternateTitles"`
	SortTitle         string            `json:"sortTitle"`
	Status            string            `json:"status"`
	Overview          string            `json:"overview"`
	Network           string            `json:"network"`
	Images            []*starr.Image    `json:"images"`
	Seasons           []*Season         `json:"seasons"`
	Year              int               `json:"year"`
	Path              string            `json:"path"`
	QualityProfileID  int               `json:"qualityProfileId"`
	LanguageProfileID int               `json:"languageProfileId"`
	Runtime           int               `json:"runtime"`
	TvdbID            int               `json:"tvdbId"`
	TvRageID          int               `json:"tvRageId"`
	TvMazeID          int               `json:"tvMazeId"`
	FirstAired        time.Time         `json:"firstAired"`
	SeriesType        string            `json:"seriesType"`
	CleanTitle        string            `json:"cleanTitle"`
	ImdbID            string            `json:"imdbId"`
	TitleSlug         string            `json:"titleSlug"`
	RootFolderPath    string            `json:"rootFolderPath"`
	Genres            []string          `json:"genres"`
	Tags              []interface{}     `json:"tags"`
	Added             time.Time         `json:"added"`
	AddOptions        *AddSeriesOptions `json:"addOptions"`
	Ratings           *starr.Ratings    `json:"ratings"`
	Statistics        *Statistics       `json:"statistics"`
	Ended             bool              `json:"ended"`
	SeasonFolder      bool              `json:"seasonFolder"`
	Monitored         bool              `json:"monitored"`
	UseSceneNumbering bool              `json:"useSceneNumbering"`
}
