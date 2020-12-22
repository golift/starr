package lidarr

import (
	"crypto/tls"
	"net/http"
	"time"

	"golift.io/starr"
)

// Lidarr contains all the methods to interact with a Lidarr server.
type Lidarr struct {
	a starr.APIer
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

	return &Lidarr{c}
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

// Artist represents the /api/v1/artist endpoint.
type Artist struct {
	ID                int64          `json:"id"`
	Status            string         `json:"status"`
	LastInfoSync      time.Time      `json:"lastInfoSync"`
	ArtistName        string         `json:"artistName"`
	ForeignArtistID   string         `json:"foreignArtistId"`
	TadbID            int64          `json:"tadbId"`
	DiscogsID         int64          `json:"discogsId"`
	Overview          string         `json:"overview"`
	ArtistType        string         `json:"artistType,omitempty"`
	Disambiguation    string         `json:"disambiguation"`
	Links             []*starr.Link  `json:"links"`
	Images            []*starr.Image `json:"images"`
	Path              string         `json:"path"`
	QualityProfileID  int64          `json:"qualityProfileId"`
	MetadataProfileID int64          `json:"metadataProfileId"`
	Genres            []string       `json:"genres"`
	CleanName         string         `json:"cleanName"`
	SortName          string         `json:"sortName"`
	Tags              []interface{}  `json:"tags"`
	Added             time.Time      `json:"added"`
	Ratings           *starr.Ratings `json:"ratings"`
	Statistics        *Statistics    `json:"statistics"`
	LastAlbum         *Album         `json:"lastAlbum,omitempty"`
	NextAlbum         *Album         `json:"nextAlbum,omitempty"`
	Ended             bool           `json:"ended"`
	AlbumFolder       bool           `json:"albumFolder"`
	Monitored         bool           `json:"monitored"`
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
	ID                 int64          `json:"id"`
	ArtistID           int64          `json:"artistId"`
	ProfileID          int64          `json:"profileId"`
	ArtistMetadataID   int64          `json:"artistMetadataId,omitempty"`
	ForeignAlbumID     string         `json:"foreignAlbumId"`
	OldForeignAlbumIds []interface{}  `json:"oldForeignAlbumIds,omitempty"`
	Title              string         `json:"title"`
	Overview           string         `json:"overview,omitempty"`
	Disambiguation     string         `json:"disambiguation"`
	ReleaseDate        time.Time      `json:"releaseDate"`
	Images             []*starr.Image `json:"images"`
	Links              []*starr.Link  `json:"links"`
	Genres             []string       `json:"genres"`
	AlbumType          string         `json:"albumType"`
	CleanTitle         string         `json:"cleanTitle,omitempty"`
	LastInfoSync       time.Time      `json:"lastInfoSync,omitempty"`
	Added              time.Time      `json:"added,omitempty"`
	SecondaryTypes     []interface{}  `json:"secondaryTypes"`
	Duration           int            `json:"duration,omitempty"`
	MediumCount        int            `json:"mediumCount,omitempty"`
	Ratings            *starr.Ratings `json:"ratings"`
	Releases           []*Releases    `json:"releases,omitempty"`
	Media              []*Media       `json:"media,omitempty"`
	Statistics         *Statistics    `json:"statistics,omitempty"`
	ArtistMetadata     *IsLoaded      `json:"artistMetadata"`
	AlbumReleases      *IsLoaded      `json:"albumReleases"`
	Artist             *IsLoaded      `json:"artist"`
	Monitored          bool           `json:"monitored"`
	AnyReleaseOk       bool           `json:"anyReleaseOk"`
}

// IsLoaded is a generic struct used in a few places.
type IsLoaded struct {
	IsLoaded bool `json:"isLoaded"`
}

// Releases is part of an Album.
type Releases struct {
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

// AddAlbumInput is currently unknown.
type AddAlbumInput interface{}

// AddAlbumOutput is currently unknown.
type AddAlbumOutput interface{}
