package lidarr

import (
	"crypto/tls"
	"net/http"
	"time"

	"golift.io/starr"
)

// Lidarr contains all the methods to interact with a Lidarr server.
type Lidarr struct {
	starr.APIer
}

// New returns a Lidarr object used to interact with the Lidarr API.
func New(config *starr.Config) *Lidarr {
	if config.Client == nil {
		//nolint:exhaustivestruct,gosec
		config.Client = &http.Client{
			Timeout: config.Timeout.Duration,
			CheckRedirect: func(r *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: !config.ValidSSL},
			},
		}
	}

	if config.Debugf == nil {
		config.Debugf = func(string, ...interface{}) {}
	}

	return &Lidarr{APIer: config}
}

// Queue is the /api/v1/queue endpoint.
type Queue struct {
	Page          int            `json:"page"`
	PageSize      int            `json:"pageSize"`
	SortKey       string         `json:"sortKey"`
	SortDirection string         `json:"sortDirection"`
	TotalRecords  int            `json:"totalRecords"`
	Records       []*QueueRecord `json:"records"`
}

// QueueRecord represents the records returns by the /api/v1/queue endpoint.
type QueueRecord struct {
	ArtistID                int64                  `json:"artistId"`
	AlbumID                 int64                  `json:"albumId"`
	Quality                 *starr.Quality         `json:"quality"`
	Size                    float64                `json:"size"`
	Title                   string                 `json:"title"`
	Sizeleft                float64                `json:"sizeleft"`
	Timeleft                string                 `json:"timeleft"`
	EstimatedCompletionTime time.Time              `json:"estimatedCompletionTime"`
	Status                  string                 `json:"status"`
	TrackedDownloadStatus   string                 `json:"trackedDownloadStatus"`
	StatusMessages          []*starr.StatusMessage `json:"statusMessages"`
	DownloadID              string                 `json:"downloadId"`
	Protocol                string                 `json:"protocol"`
	DownloadClient          string                 `json:"downloadClient"`
	Indexer                 string                 `json:"indexer"`
	OutputPath              string                 `json:"outputPath"`
	DownloadForced          bool                   `json:"downloadForced"`
	ID                      int64                  `json:"id"`
	ErrorMessage            string                 `json:"errorMessage"`
}

// QualityProfile is the /api/v1/qualityprofile endpoint.
type QualityProfile struct {
	ID             int64            `json:"id"`
	Name           string           `json:"name"`
	UpgradeAllowed bool             `json:"upgradeAllowed"`
	Cutoff         int64            `json:"cutoff"`
	Qualities      []*starr.Quality `json:"items"`
}

// MetadataProfile is the /api/v1/metadataprofile endpoint.
type MetadataProfile struct {
	Name                string           `json:"name"`
	ID                  int64            `json:"id"`
	PrimaryAlbumTypes   []*AlbumType     `json:"primaryAlbumTypes"`
	SecondaryAlbumTypes []*AlbumType     `json:"secondaryAlbumTypes"`
	ReleaseStatuses     []*ReleaseStatus `json:"releaseStatuses"`
}

// AlbumType is part of MetadataProfile.
type AlbumType struct {
	AlbumType *starr.Value `json:"albumType"`
	Allowed   bool         `json:"allowed"`
}

// ReleaseStatus is part of MetadataProfile.
type ReleaseStatus struct {
	ReleaseStatus *starr.Value `json:"releaseStatus"`
	Allowed       bool         `json:"allowed"`
}

// RootFolder is the /api/v1/rootfolder endpoint.
type RootFolder struct {
	ID              int64         `json:"id"`
	Path            string        `json:"path"`
	FreeSpace       int64         `json:"freeSpace"`
	TotalSpace      int64         `json:"totalSpace"`
	UnmappedFolders []*starr.Path `json:"unmappedFolders"`
}

// QualityDefinition is the /api/v1/qualitydefinition endpoint.
type QualityDefinition struct {
	ID      int64        `json:"id"`
	Quality *starr.Value `json:"quality"`
	Title   string       `json:"title"`
	Weight  int64        `json:"weight"`
	MinSize float64      `json:"minSize"`
	MaxSize float64      `json:"maxSize,omitempty"`
}

// SystemStatus is the /api/v1/system/status endpoint.
type SystemStatus struct {
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
	IsDocker          bool      `json:"isDocker"`
	Mode              string    `json:"mode"`
	Branch            string    `json:"branch"`
	Authentication    string    `json:"authentication"`
	SqliteVersion     string    `json:"sqliteVersion"`
	MigrationVersion  int64     `json:"migrationVersion"`
	URLBase           string    `json:"urlBase"`
	RuntimeVersion    string    `json:"runtimeVersion"`
	RuntimeName       string    `json:"runtimeName"`
	StartTime         time.Time `json:"startTime"`
}

// Artist represents the /api/v1/artist endpoint, and it's part of an Album.
type Artist struct {
	ID                int64             `json:"id"`
	Status            string            `json:"status,omitempty"`
	LastInfoSync      time.Time         `json:"lastInfoSync,omitempty"`
	ArtistName        string            `json:"artistName,omitempty"`
	ForeignArtistID   string            `json:"foreignArtistId,omitempty"`
	TadbID            int64             `json:"tadbId,omitempty"`
	DiscogsID         int64             `json:"discogsId,omitempty"`
	QualityProfileID  int64             `json:"qualityProfileId,omitempty"`
	MetadataProfileID int64             `json:"metadataProfileId,omitempty"`
	Overview          string            `json:"overview,omitempty"`
	ArtistType        string            `json:"artistType,omitempty"`
	Disambiguation    string            `json:"disambiguation,omitempty"`
	RootFolderPath    string            `json:"rootFolderPath,omitempty"`
	Path              string            `json:"path,omitempty"`
	CleanName         string            `json:"cleanName,omitempty"`
	SortName          string            `json:"sortName,omitempty"`
	Links             []*starr.Link     `json:"links,omitempty"`
	Images            []*starr.Image    `json:"images,omitempty"`
	Genres            []string          `json:"genres,omitempty"`
	Tags              []int             `json:"tags,omitempty"`
	Added             time.Time         `json:"added,omitempty"`
	Ratings           *starr.Ratings    `json:"ratings,omitempty"`
	Statistics        *Statistics       `json:"statistics,omitempty"`
	LastAlbum         *Album            `json:"lastAlbum,omitempty"`
	NextAlbum         *Album            `json:"nextAlbum,omitempty"`
	AddOptions        *ArtistAddOptions `json:"addOptions,omitempty"`
	AlbumFolder       bool              `json:"albumFolder,omitempty"`
	Monitored         bool              `json:"monitored"`
	Ended             bool              `json:"ended,omitempty"`
}

// Statistics is part of Artist.
type Statistics struct {
	AlbumCount      int     `json:"albumCount,omitempty"`
	TrackFileCount  int     `json:"trackFileCount"`
	TrackCount      int     `json:"trackCount"`
	TotalTrackCount int     `json:"totalTrackCount"`
	SizeOnDisk      int     `json:"sizeOnDisk"`
	PercentOfTracks float64 `json:"percentOfTracks"`
}

// Album is the /api/v1/album endpoint.
type Album struct {
	ID             int64            `json:"id,omitempty"`
	Title          string           `json:"title"`
	Disambiguation string           `json:"disambiguation"`
	Overview       string           `json:"overview"`
	ArtistID       int64            `json:"artistId"`
	ForeignAlbumID string           `json:"foreignAlbumId"`
	ProfileID      int64            `json:"profileId"`
	Duration       int              `json:"duration"`
	AlbumType      string           `json:"albumType"`
	SecondaryTypes []interface{}    `json:"secondaryTypes"`
	MediumCount    int              `json:"mediumCount"`
	Ratings        *starr.Ratings   `json:"ratings"`
	ReleaseDate    time.Time        `json:"releaseDate"`
	Releases       []*Release       `json:"releases"`
	Genres         []interface{}    `json:"genres"`
	Media          []*Media         `json:"media"`
	Artist         *Artist          `json:"artist"`
	Links          []*starr.Link    `json:"links"`
	Images         []*starr.Image   `json:"images"`
	Statistics     *Statistics      `json:"statistics"`
	RemoteCover    string           `json:"remoteCover,omitempty"`
	AddOptions     *AlbumAddOptions `json:"addOptions,omitempty"`
	Monitored      bool             `json:"monitored"`
	AnyReleaseOk   bool             `json:"anyReleaseOk"`
}

// Release is part of an Album.
type Release struct {
	ID               int64    `json:"id"`
	AlbumID          int64    `json:"albumId"`
	ForeignReleaseID string   `json:"foreignReleaseId"`
	Title            string   `json:"title"`
	Status           string   `json:"status"`
	Duration         int      `json:"duration"`
	TrackCount       int      `json:"trackCount"`
	Media            []*Media `json:"media"`
	MediumCount      int      `json:"mediumCount"`
	Disambiguation   string   `json:"disambiguation"`
	Country          []string `json:"country"`
	Label            []string `json:"label"`
	Format           string   `json:"format"`
	Monitored        bool     `json:"monitored"`
}

// Media is part of an Album.
type Media struct {
	MediumNumber int64  `json:"mediumNumber"`
	MediumName   string `json:"mediumName"`
	MediumFormat string `json:"mediumFormat"`
}

// ArtistAddOptions is part of an artist and an album.
type ArtistAddOptions struct {
	Monitor                string `json:"monitor,omitempty"`
	Monitored              bool   `json:"monitored,omitempty"`
	SearchForMissingAlbums bool   `json:"searchForMissingAlbums,omitempty"`
}

// AddAlbumInput is currently unknown.
type AddAlbumInput struct {
	ForeignAlbumID string                  `json:"foreignAlbumId"`
	Monitored      bool                    `json:"monitored"`
	Releases       []*AddAlbumInputRelease `json:"releases"`
	AddOptions     *AlbumAddOptions        `json:"addOptions"`
	Artist         *Artist                 `json:"artist"`
}

// AddAlbumInputRelease is part of AddAlbumInput.
type AddAlbumInputRelease struct {
	ForeignReleaseID string   `json:"foreignReleaseId"`
	Title            string   `json:"title"`
	Media            []*Media `json:"media"`
	Monitored        bool     `json:"monitored"`
}

// AlbumAddOptions is part of an Album.
type AlbumAddOptions struct {
	SearchForNewAlbum bool `json:"searchForNewAlbum,omitempty"`
}

// CommandRequest goes into the /api/v1/command endpoint.
// This was created from the search command and may not support other commands yet.
type CommandRequest struct {
	Name     string  `json:"name"`
	AlbumIDs []int64 `json:"albumIds,omitempty"`
	AlbumID  int64   `json:"albumId,omitempty"`
}

// CommandResponse comes from the /api/v1/command endpoint.
type CommandResponse struct {
	ID                  int64                  `json:"id"`
	Name                string                 `json:"name"`
	CommandName         string                 `json:"commandName"`
	Message             string                 `json:"message,omitempty"`
	Priority            string                 `json:"priority"`
	Status              string                 `json:"status"`
	Queued              time.Time              `json:"queued"`
	Started             time.Time              `json:"started,omitempty"`
	Ended               time.Time              `json:"ended,omitempty"`
	StateChangeTime     time.Time              `json:"stateChangeTime,omitempty"`
	LastExecutionTime   time.Time              `json:"lastExecutionTime,omitempty"`
	Duration            string                 `json:"duration,omitempty"`
	Trigger             string                 `json:"trigger"`
	SendUpdatesToClient bool                   `json:"sendUpdatesToClient"`
	UpdateScheduledTask bool                   `json:"updateScheduledTask"`
	Body                map[string]interface{} `json:"body"`
}

// History represents the /api/v1/history endpoint.
type History struct {
	Page          int              `json:"page"`
	PageSize      int              `json:"pageSize"`
	SortKey       string           `json:"sortKey"`
	SortDirection string           `json:"sortDirection"`
	TotalRecords  int              `json:"totalRecords"`
	Records       []*HistoryRecord `json:"records"`
}

// HistoryRecord is part of the history. Not all items have all Data members.
// Check EventType for events you need.
type HistoryRecord struct {
	ID                  int64          `json:"id"`
	AlbumID             int64          `json:"albumId"`
	ArtistID            int64          `json:"artistId"`
	TrackID             int64          `json:"trackId"`
	SourceTitle         string         `json:"sourceTitle"`
	Quality             *starr.Quality `json:"quality"`
	QualityCutoffNotMet bool           `json:"qualityCutoffNotMet"`
	Date                time.Time      `json:"date"`
	DownloadID          string         `json:"downloadId"`
	EventType           string         `json:"eventType"`
	Data                struct {
		Age             string    `json:"age"`
		AgeHours        string    `json:"ageHours"`
		AgeMinutes      string    `json:"ageMinutes"`
		DownloadClient  string    `json:"downloadClient"`
		DownloadForced  string    `json:"downloadForced"`
		DownloadURL     string    `json:"downloadUrl"`
		DroppedPath     string    `json:"droppedPath"`
		GUID            string    `json:"guid"`
		ImportedPath    string    `json:"importedPath"`
		Indexer         string    `json:"indexer"`
		Message         string    `json:"message"`
		NzbInfoURL      string    `json:"nzbInfoUrl"`
		Protocol        string    `json:"protocol"`
		PublishedDate   time.Time `json:"publishedDate"`
		Reason          string    `json:"reason"`
		ReleaseGroup    string    `json:"releaseGroup"`
		Size            string    `json:"size"`
		StatusMessages  string    `json:"statusMessages"`
		TorrentInfoHash string    `json:"torrentInfoHash"`
	} `json:"data"`
}
