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
func New(c *starr.Config) *Lidarr {
	if c.Client == nil {
		//nolint:exhaustivestruct,gosec
		c.Client = &http.Client{
			Timeout: c.Timeout.Duration,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: !c.ValidSSL},
			},
		}
	}

	return &Lidarr{APIer: c}
}

// Queue is the /api/v1/queue endpoint.
type Queue struct {
	Page          int       `json:"page"`
	PageSize      int       `json:"pageSize"`
	SortKey       string    `json:"sortKey"`
	SortDirection string    `json:"sortDirection"`
	TotalRecords  int       `json:"totalRecords"`
	Records       []*Record `json:"records"`
}

// Record represents the records returns by the /api/v1/queue endpoint.
type Record struct {
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
	ID                int64          `json:"id"`
	Status            string         `json:"status,omitempty"`
	LastInfoSync      time.Time      `json:"lastInfoSync,omitempty"`
	ArtistName        string         `json:"artistName,omitempty"`
	ForeignArtistID   string         `json:"foreignArtistId,omitempty"`
	TadbID            int64          `json:"tadbId,omitempty"`
	DiscogsID         int64          `json:"discogsId,omitempty"`
	Overview          string         `json:"overview,omitempty"`
	ArtistType        string         `json:"artistType,omitempty"`
	Disambiguation    string         `json:"disambiguation,omitempty"`
	Links             []*starr.Link  `json:"links,omitempty"`
	Images            []*starr.Image `json:"images,omitempty"`
	Path              string         `json:"path,omitempty"`
	QualityProfileID  int64          `json:"qualityProfileId,omitempty"`
	MetadataProfileID int64          `json:"metadataProfileId,omitempty"`
	Genres            []string       `json:"genres,omitempty"`
	CleanName         string         `json:"cleanName,omitempty"`
	SortName          string         `json:"sortName,omitempty"`
	Tags              []interface{}  `json:"tags,omitempty"`
	Added             time.Time      `json:"added,omitempty"`
	Ratings           *starr.Ratings `json:"ratings,omitempty"`
	Statistics        *Statistics    `json:"statistics,omitempty"`
	LastAlbum         *Album         `json:"lastAlbum,omitempty"`
	NextAlbum         *Album         `json:"nextAlbum,omitempty"`
	Ended             bool           `json:"ended,omitempty"`
	AlbumFolder       bool           `json:"albumFolder,omitempty"`
	Monitored         bool           `json:"monitored,omitempty"`
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
	ID             int64          `json:"id"`
	Title          string         `json:"title"`
	Disambiguation string         `json:"disambiguation"`
	Overview       string         `json:"overview"`
	ArtistID       int64          `json:"artistId"`
	ForeignAlbumID string         `json:"foreignAlbumId"`
	Monitored      bool           `json:"monitored"`
	AnyReleaseOk   bool           `json:"anyReleaseOk"`
	ProfileID      int64          `json:"profileId"`
	Duration       int            `json:"duration"`
	AlbumType      string         `json:"albumType"`
	SecondaryTypes []interface{}  `json:"secondaryTypes"`
	MediumCount    int            `json:"mediumCount"`
	Ratings        *starr.Ratings `json:"ratings"`
	ReleaseDate    time.Time      `json:"releaseDate"`
	Releases       []*Release     `json:"releases"`
	Genres         []interface{}  `json:"genres"`
	Media          []*Media       `json:"media"`
	Artist         *Artist        `json:"artist"`
	Links          []*starr.Link  `json:"links"`
	Images         []*starr.Image `json:"images"`
	Statistics     *Statistics    `json:"statistics"`
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

// XXX: fix these.

// AddArtistInput is currently unknown.
type AddArtistInput struct {
	RootFolderPath    string            `json:"rootFolderPath"`
	QualityProfileID  int               `json:"qualityProfileId"`
	MetadataProfileID int               `json:"metadataProfileId"`
	ForeignArtistID   string            `json:"foreignArtistId"`
	Monitored         bool              `json:"monitored"`
	AddOptions        *ArtistAddOptions `json:"addOptions"`
}

// AddArtistOutput is currently unknown.
type AddArtistOutput struct {
	Status            string            `json:"status"`
	Ended             bool              `json:"ended"`
	ArtistName        string            `json:"artistName"`
	ForeignArtistID   string            `json:"foreignArtistId"`
	TadbID            int               `json:"tadbId"`
	DiscogsID         int               `json:"discogsId"`
	Overview          string            `json:"overview"`
	Disambiguation    string            `json:"disambiguation"`
	Links             []*starr.Link     `json:"links"`
	Images            []*starr.Image    `json:"images"`
	Path              string            `json:"path"`
	QualityProfileID  int               `json:"qualityProfileId"`
	MetadataProfileID int               `json:"metadataProfileId"`
	AlbumFolder       bool              `json:"albumFolder"`
	Monitored         bool              `json:"monitored"`
	Genres            []string          `json:"genres"`
	CleanName         string            `json:"cleanName"`
	SortName          string            `json:"sortName"`
	Tags              []interface{}     `json:"tags"`
	Added             time.Time         `json:"added"`
	Ratings           *starr.Ratings    `json:"ratings"`
	Statistics        *Statistics       `json:"statistics"`
	ID                int               `json:"id"`
	AddOptions        *ArtistAddOptions `json:"addOptions"`
	RootFolderPath    string            `json:"rootFolderPath"`
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
	Artist         *AddArtistInput         `json:"artist"`
}

// AddAlbumInputRelease is part of AddAlbumInput.
type AddAlbumInputRelease struct {
	ForeignReleaseID string   `json:"foreignReleaseId"`
	Title            string   `json:"title"`
	Media            []*Media `json:"media"`
	Monitored        bool     `json:"monitored"`
}

// AddAlbumOutput is currently unknown.
type AddAlbumOutput struct {
	Title          string           `json:"title"`
	Disambiguation string           `json:"disambiguation"`
	Overview       string           `json:"overview"`
	ArtistID       int              `json:"artistId"`
	ForeignAlbumID string           `json:"foreignAlbumId"`
	Monitored      bool             `json:"monitored"`
	AnyReleaseOk   bool             `json:"anyReleaseOk"`
	ProfileID      int              `json:"profileId"`
	Duration       int              `json:"duration"`
	AlbumType      string           `json:"albumType"`
	SecondaryTypes []interface{}    `json:"secondaryTypes"`
	MediumCount    int              `json:"mediumCount"`
	Ratings        *starr.Ratings   `json:"ratings"`
	ReleaseDate    time.Time        `json:"releaseDate"`
	Releases       []*Release       `json:"releases"`
	Genres         []string         `json:"genres"`
	Media          []*Media         `json:"media"`
	Artist         *AddArtistOutput `json:"artist"`
	Images         []*starr.Image   `json:"images"`
	Links          []*starr.Link    `json:"links"`
	RemoteCover    string           `json:"remoteCover"`
	AddOptions     *AlbumAddOptions `json:"addOptions"`
}

// AlbumAddOptions is part of an Album.
type AlbumAddOptions struct {
	SearchForNewAlbum bool `json:"searchForNewAlbum,omitempty"`
}
