package sonarr

import (
	"time"

	"golift.io/starr"
)

// APIver is the Sonarr API version supported by this library.
const APIver = "v3"

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
	ID              int64         `json:"id"`
	Path            string        `json:"path"`
	Accessible      bool          `json:"accessible"`
	FreeSpace       int64         `json:"freeSpace"`
	UnmappedFolders []*starr.Path `json:"unmappedFolders"`
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
	ID                      int64                  `json:"id"`
	SeriesID                int64                  `json:"seriesId"`
	EpisodeID               int64                  `json:"episodeId"`
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
	ErrorMessage            string                 `json:"errorMessage"`
}

// Series the /api/v3/series endpoint.
type Series struct {
	ID                int64             `json:"id"`
	Title             string            `json:"title,omitempty"`
	AlternateTitles   []*AlternateTitle `json:"alternateTitles,omitempty"`
	SortTitle         string            `json:"sortTitle,omitempty"`
	Status            string            `json:"status,omitempty"`
	Overview          string            `json:"overview,omitempty"`
	PreviousAiring    time.Time         `json:"previousAiring,omitempty"`
	Network           string            `json:"network,omitempty"`
	Images            []*starr.Image    `json:"images,omitempty"`
	Seasons           []*Season         `json:"seasons,omitempty"`
	Year              int               `json:"year,omitempty"`
	Path              string            `json:"path,omitempty"`
	QualityProfileID  int64             `json:"qualityProfileId,omitempty"`
	LanguageProfileID int64             `json:"languageProfileId,omitempty"`
	Runtime           int               `json:"runtime,omitempty"`
	TvdbID            int64             `json:"tvdbId,omitempty"`
	TvRageID          int64             `json:"tvRageId,omitempty"`
	TvMazeID          int64             `json:"tvMazeId,omitempty"`
	FirstAired        time.Time         `json:"firstAired,omitempty"`
	SeriesType        string            `json:"seriesType,omitempty"`
	CleanTitle        string            `json:"cleanTitle,omitempty"`
	ImdbID            string            `json:"imdbId,omitempty"`
	TitleSlug         string            `json:"titleSlug,omitempty"`
	RootFolderPath    string            `json:"rootFolderPath,omitempty"`
	Certification     string            `json:"certification,omitempty"`
	Genres            []string          `json:"genres,omitempty"`
	Tags              []int             `json:"tags,omitempty"`
	Added             time.Time         `json:"added,omitempty"`
	Ratings           *starr.Ratings    `json:"ratings,omitempty"`
	Statistics        *Statistics       `json:"statistics,omitempty"`
	NextAiring        time.Time         `json:"nextAiring,omitempty"`
	AirTime           string            `json:"airTime,omitempty"`
	Ended             bool              `json:"ended,omitempty"`
	SeasonFolder      bool              `json:"seasonFolder,omitempty"`
	Monitored         bool              `json:"monitored"`
	UseSceneNumbering bool              `json:"useSceneNumbering,omitempty"`
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

// Episode is the /api/v3/episode endpoint.
type Episode struct {
	ID                       int64     `json:"id"`
	SeriesID                 int64     `json:"seriesId"`
	AbsoluteEpisodeNumber    int64     `json:"absoluteEpisodeNumber"`
	EpisodeFileID            int64     `json:"episodeFileId"`
	SeasonNumber             int64     `json:"seasonNumber"`
	EpisodeNumber            int64     `json:"episodeNumber"`
	AirDateUtc               time.Time `json:"airDateUtc"`
	AirDate                  string    `json:"airDate"`
	Title                    string    `json:"title"`
	Overview                 string    `json:"overview"`
	UnverifiedSceneNumbering bool      `json:"unverifiedSceneNumbering"`
	HasFile                  bool      `json:"hasFile"`
	Monitored                bool      `json:"monitored"`
}

// AddSeriesInput is the input for a POST to the /api/v3/series endpoint.
type AddSeriesInput struct {
	ID                int64             `json:"id,omitempty"`
	TvdbID            int64             `json:"tvdbId"`
	QualityProfileID  int64             `json:"qualityProfileId"`
	LanguageProfileID int64             `json:"languageProfileId"`
	Tags              []int             `json:"tags"`
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
	ID                int64             `json:"id"`
	Title             string            `json:"title"`
	AlternateTitles   []*AlternateTitle `json:"alternateTitles"`
	SortTitle         string            `json:"sortTitle"`
	Status            string            `json:"status"`
	Overview          string            `json:"overview"`
	Network           string            `json:"network"`
	Images            []*starr.Image    `json:"images"`
	Seasons           []*Season         `json:"seasons"`
	Year              int               `json:"year"`
	Path              string            `json:"path"`
	QualityProfileID  int64             `json:"qualityProfileId"`
	LanguageProfileID int64             `json:"languageProfileId"`
	Runtime           int               `json:"runtime"`
	TvdbID            int64             `json:"tvdbId"`
	TvRageID          int64             `json:"tvRageId"`
	TvMazeID          int64             `json:"tvMazeId"`
	FirstAired        time.Time         `json:"firstAired"`
	SeriesType        string            `json:"seriesType"`
	CleanTitle        string            `json:"cleanTitle"`
	ImdbID            string            `json:"imdbId"`
	TitleSlug         string            `json:"titleSlug"`
	RootFolderPath    string            `json:"rootFolderPath"`
	Genres            []string          `json:"genres"`
	Tags              []int             `json:"tags"`
	Added             time.Time         `json:"added"`
	AddOptions        *AddSeriesOptions `json:"addOptions"`
	Ratings           *starr.Ratings    `json:"ratings"`
	Statistics        *Statistics       `json:"statistics"`
	Ended             bool              `json:"ended"`
	SeasonFolder      bool              `json:"seasonFolder"`
	Monitored         bool              `json:"monitored"`
	UseSceneNumbering bool              `json:"useSceneNumbering"`
}

// AlternateTitle is part of a Series.
type AlternateTitle struct {
	Title        string `json:"title"`
	SeasonNumber int    `json:"seasonNumber"`
}

// History is the data from the /api/v3/history endpoint.
type History struct {
	Page          int              `json:"page"`
	PageSize      int              `json:"pageSize"`
	SortKey       string           `json:"sortKey"`
	SortDirection string           `json:"sortDirection"`
	TotalRecords  int              `json:"totalRecords"`
	Records       []*HistoryRecord `json:"records"`
}

// HistoryRecord is part of the History data.
// Not all items have all Data members. Check EventType for what you need.
type HistoryRecord struct {
	ID                   int64          `json:"id"`
	EpisodeID            int64          `json:"episodeId"`
	SeriesID             int64          `json:"seriesId"`
	SourceTitle          string         `json:"sourceTitle"`
	Language             Language       `json:"language"`
	Quality              *starr.Quality `json:"quality"`
	QualityCutoffNotMet  bool           `json:"qualityCutoffNotMet"`
	LanguageCutoffNotMet bool           `json:"languageCutoffNotMet"`
	Date                 time.Time      `json:"date"`
	DownloadID           string         `json:"downloadId,omitempty"`
	EventType            string         `json:"eventType"`
	Data                 struct {
		Age                string    `json:"age"`
		AgeHours           string    `json:"ageHours"`
		AgeMinutes         string    `json:"ageMinutes"`
		DownloadClient     string    `json:"downloadClient"`
		DownloadClientName string    `json:"downloadClientName"`
		DownloadURL        string    `json:"downloadUrl"`
		DroppedPath        string    `json:"droppedPath"`
		FileID             string    `json:"fileId"`
		GUID               string    `json:"guid"`
		ImportedPath       string    `json:"importedPath"`
		Indexer            string    `json:"indexer"`
		Message            string    `json:"message"`
		NzbInfoURL         string    `json:"nzbInfoUrl"`
		PreferredWordScore string    `json:"preferredWordScore"`
		Protocol           string    `json:"protocol"`
		PublishedDate      time.Time `json:"publishedDate"`
		Reason             string    `json:"reason"`
		ReleaseGroup       string    `json:"releaseGroup"`
		Size               string    `json:"size"`
		TorrentInfoHash    string    `json:"torrentInfoHash"`
		TvRageID           string    `json:"tvRageId"`
		TvdbID             string    `json:"tvdbId"`
	} `json:"data"`
}

// EpisodeFile is the output from the /api/v3/episodeFile endpoint.
type EpisodeFile struct {
	ID                   int64          `json:"id"`
	SeriesID             int64          `json:"seriesId"`
	SeasonNumber         int            `json:"seasonNumber"`
	RelativePath         string         `json:"relativePath"`
	Path                 string         `json:"path"`
	Size                 int64          `json:"size"`
	DateAdded            time.Time      `json:"dateAdded"`
	SceneName            string         `json:"sceneName"`
	ReleaseGroup         string         `json:"releaseGroup"`
	Language             *starr.Value   `json:"language"`
	Quality              *starr.Quality `json:"quality"`
	MediaInfo            *MediaInfo     `json:"mediaInfo"`
	QualityCutoffNotMet  bool           `json:"qualityCutoffNotMet"`
	LanguageCutoffNotMet bool           `json:"languageCutoffNotMet"`
}

// MediaInfo is part of an EpisodeFile.
type MediaInfo struct {
	AudioBitrate     int            `json:"audioBitrate"`
	AudioChannels    float64        `json:"audioChannels"`
	AudioCodec       string         `json:"audioCodec"`
	AudioLanguages   string         `json:"audioLanguages"`
	AudioStreamCount int            `json:"audioStreamCount"`
	VideoBitDepth    int            `json:"videoBitDepth"`
	VideoBitrate     int            `json:"videoBitrate"`
	VideoCodec       string         `json:"videoCodec"`
	VideoFPS         float64        `json:"videoFps"`
	Resolution       string         `json:"resolution"`
	RunTime          starr.PlayTime `json:"runTime"`
	ScanType         string         `json:"scanType"`
	Subtitles        string         `json:"subtitles"`
}
