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
