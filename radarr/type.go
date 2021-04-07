package radarr

import (
	"crypto/tls"
	"net/http"
	"time"

	"golift.io/starr"
)

// Radarr contains all the methods to interact with a Radarr server.
type Radarr struct {
	starr.APIer
}

// New returns a Radarr object used to interact with the Radarr API.
func New(c *starr.Config) *Radarr {
	if c.Client == nil {
		//nolint:exhaustivestruct,gosec
		c.Client = &http.Client{
			Timeout: c.Timeout.Duration,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: !c.ValidSSL},
			},
		}
	}

	return &Radarr{APIer: c}
}

// SystemStatus is the /api/v1/system/status endpoint.
type SystemStatus struct {
	Version           string    `json:"version"`
	BuildTime         time.Time `json:"buildTime"`
	StartupPath       string    `json:"startupPath"`
	AppData           string    `json:"appData"`
	OsName            string    `json:"osName"`
	OsVersion         string    `json:"osVersion"`
	Branch            string    `json:"branch"`
	Authentication    string    `json:"authentication"`
	SqliteVersion     string    `json:"sqliteVersion"`
	URLBase           string    `json:"urlBase"`
	RuntimeVersion    string    `json:"runtimeVersion"`
	RuntimeName       string    `json:"runtimeName"`
	MigrationVersion  int       `json:"migrationVersion"`
	IsDebug           bool      `json:"isDebug"`
	IsProduction      bool      `json:"isProduction"`
	IsAdmin           bool      `json:"isAdmin"`
	IsUserInteractive bool      `json:"isUserInteractive"`
	IsNetCore         bool      `json:"isNetCore"`
	IsMono            bool      `json:"isMono"`
	IsLinux           bool      `json:"isLinux"`
	IsOsx             bool      `json:"isOsx"`
	IsWindows         bool      `json:"isWindows"`
}

// AddMovieInput is the input for a new movie.
type AddMovieInput struct {
	Title               string           `json:"title,omitempty"`
	TitleSlug           string           `json:"titleSlug,omitempty"`
	MinimumAvailability string           `json:"minimumAvailability,omitempty"`
	RootFolderPath      string           `json:"rootFolderPath"`
	TmdbID              int64            `json:"tmdbId"`
	QualityProfileID    int64            `json:"qualityProfileId"`
	ProfileID           int64            `json:"profileId,omitempty"`
	Year                int              `json:"year,omitempty"`
	Images              []*starr.Image   `json:"images,omitempty"`
	AddOptions          *AddMovieOptions `json:"addOptions"`
	Tags                []int            `json:"tags,omitempty"`
	Monitored           bool             `json:"monitored"`
}

// AddMovieOptions are the options for finding a new movie.
type AddMovieOptions struct {
	SearchForMovie             bool `json:"searchForMovie"`
	IgnoreEpisodesWithFiles    bool `json:"ignoreEpisodesWithFiles,omitempty"`
	IgnoreEpisodesWithoutFiles bool `json:"ignoreEpisodesWithoutFiles,omitempty"`
}

// AddMovieOutput is the data returned when adding a movier.
type AddMovieOutput struct {
	ID                    int64               `json:"id"`
	Title                 string              `json:"title"`
	OriginalTitle         string              `json:"originalTitle"`
	AlternateTitles       []*AlternativeTitle `json:"alternateTitles"`
	SecondaryYearSourceID int64               `json:"secondaryYearSourceId"`
	SortTitle             string              `json:"sortTitle"`
	SizeOnDisk            int                 `json:"sizeOnDisk"`
	Status                string              `json:"status"`
	Overview              string              `json:"overview"`
	InCinemas             time.Time           `json:"inCinemas"`
	DigitalRelease        time.Time           `json:"digitalRelease"`
	Images                []*starr.Image      `json:"images"`
	Website               string              `json:"website"`
	Year                  int                 `json:"year"`
	YouTubeTrailerID      string              `json:"youTubeTrailerId"`
	Studio                string              `json:"studio"`
	Path                  string              `json:"path"`
	QualityProfileID      int64               `json:"qualityProfileId"`
	MinimumAvailability   string              `json:"minimumAvailability"`
	FolderName            string              `json:"folderName"`
	Runtime               int                 `json:"runtime"`
	CleanTitle            string              `json:"cleanTitle"`
	ImdbID                string              `json:"imdbId"`
	TmdbID                int64               `json:"tmdbId"`
	TitleSlug             string              `json:"titleSlug"`
	Genres                []string            `json:"genres"`
	Tags                  []int               `json:"tags"`
	Added                 time.Time           `json:"added"`
	AddOptions            *AddMovieOptions    `json:"addOptions"`
	Ratings               *starr.Ratings      `json:"ratings"`
	HasFile               bool                `json:"hasFile"`
	Monitored             bool                `json:"monitored"`
	IsAvailable           bool                `json:"isAvailable"`
}

// AlternativeTitle is part of a Movie.
type AlternativeTitle struct {
	MovieID    int          `json:"movieId"`
	Title      string       `json:"title"`
	SourceType string       `json:"sourceType"`
	SourceID   int          `json:"sourceId"`
	Votes      int          `json:"votes"`
	VoteCount  int          `json:"voteCount"`
	Language   *starr.Value `json:"language"`
	ID         int          `json:"id"`
}

// RootFolder is the /rootFolder endpoint.
type RootFolder struct {
	ID              int64         `json:"id"`
	Path            string        `json:"path"`
	Accessible      bool          `json:"accessible"`
	FreeSpace       int64         `json:"freeSpace"`
	UnmappedFolders []*starr.Path `json:"unmappedFolders"`
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
	ID                  int64          `json:"id"`
	EpisodeID           int64          `json:"episodeId"`
	MovieID             int64          `json:"movieId"`
	SeriesID            int64          `json:"seriesId"`
	SourceTitle         string         `json:"sourceTitle"`
	Quality             *starr.Quality `json:"quality"`
	QualityCutoffNotMet bool           `json:"qualityCutoffNotMet"`
	Date                time.Time      `json:"date"`
	DownloadID          string         `json:"downloadId"`
	EventType           string         `json:"eventType"`
	Data                *RecordData    `json:"data"`
	Movie               *RecordMovie   `json:"movie"`
}

// RecordMovie belongs to a Record.
type RecordMovie struct {
	ID                int64          `json:"id"`
	Downloaded        bool           `json:"downloaded"`
	Monitored         bool           `json:"monitored"`
	HasFile           bool           `json:"hasFile"`
	Year              int            `json:"year"`
	ProfileID         int64          `json:"profileId"`
	Runtime           int            `json:"runtime"`
	QualityProfileID  int64          `json:"qualityProfileId"`
	SizeOnDisk        int64          `json:"sizeOnDisk"`
	Title             string         `json:"title"`
	SortTitle         string         `json:"sortTitle"`
	Status            string         `json:"status"`
	Overview          string         `json:"overview"`
	InCinemas         time.Time      `json:"inCinemas"`
	Images            []*starr.Image `json:"images"`
	Website           string         `json:"website"`
	YouTubeTrailerID  string         `json:"youTubeTrailerId"`
	Studio            string         `json:"studio"`
	Path              string         `json:"path"`
	LastInfoSync      time.Time      `json:"lastInfoSync"`
	CleanTitle        string         `json:"cleanTitle"`
	ImdbID            string         `json:"imdbId"`
	TmdbID            int64          `json:"tmdbId"`
	TitleSlug         string         `json:"titleSlug"`
	Genres            []string       `json:"genres"`
	Tags              []int          `json:"tags"`
	Added             time.Time      `json:"added"`
	Ratings           *starr.Ratings `json:"ratings"`
	AlternativeTitles []string       `json:"alternativeTitles"`
}

// RecordData belongs to a Record.
type RecordData struct {
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
}

// Queue is the /api/v3/queue endpoint.
type Queue struct {
	ID                      int64                  `json:"id"`
	Size                    float64                `json:"size"`
	Sizeleft                float64                `json:"sizeleft"`
	EstimatedCompletionTime time.Time              `json:"estimatedCompletionTime"`
	Title                   string                 `json:"title"`
	Timeleft                string                 `json:"timeleft"`
	Status                  string                 `json:"status"`
	TrackedDownloadStatus   string                 `json:"trackedDownloadStatus"`
	DownloadID              string                 `json:"downloadId"`
	Protocol                string                 `json:"protocol"`
	Movie                   *QueueMovie            `json:"movie"`
	Quality                 *starr.Quality         `json:"quality"`
	StatusMessages          []*starr.StatusMessage `json:"statusMessages"`
}

// QueueMovie is part of a Movie in the Queue.
type QueueMovie struct {
	Downloaded            bool           `json:"downloaded"`
	HasFile               bool           `json:"hasFile"`
	Monitored             bool           `json:"monitored"`
	IsAvailable           bool           `json:"isAvailable"`
	SecondaryYearSourceID int            `json:"secondaryYearSourceId"`
	Year                  int            `json:"year"`
	ProfileID             int64          `json:"profileId"`
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
	Tags                  []int          `json:"tags"`
	Images                []*starr.Image `json:"images"`
	Ratings               *starr.Ratings `json:"ratings"`
}

// Movie is the /api/v3/movie endpoint.
type Movie struct {
	ID                    int64               `json:"id"`
	Title                 string              `json:"title,omitempty"`
	Path                  string              `json:"path,omitempty"`
	MinimumAvailability   string              `json:"minimumAvailability,omitempty"`
	QualityProfileID      int64               `json:"qualityProfileId,omitempty"`
	TmdbID                int64               `json:"tmdbId,omitempty"`
	OriginalTitle         string              `json:"originalTitle,omitempty"`
	AlternateTitles       []*AlternativeTitle `json:"alternateTitles,omitempty"`
	SecondaryYearSourceID int                 `json:"secondaryYearSourceId,omitempty"`
	SortTitle             string              `json:"sortTitle,omitempty"`
	SizeOnDisk            int64               `json:"sizeOnDisk,omitempty"`
	Status                string              `json:"status,omitempty"`
	Overview              string              `json:"overview,omitempty"`
	InCinemas             time.Time           `json:"inCinemas,omitempty"`
	PhysicalRelease       time.Time           `json:"physicalRelease,omitempty"`
	DigitalRelease        time.Time           `json:"digitalRelease,omitempty"`
	Images                []*starr.Image      `json:"images,omitempty"`
	Website               string              `json:"website,omitempty"`
	Year                  int                 `json:"year,omitempty"`
	YouTubeTrailerID      string              `json:"youTubeTrailerId,omitempty"`
	Studio                string              `json:"studio,omitempty"`
	FolderName            string              `json:"folderName,omitempty"`
	Runtime               int                 `json:"runtime,omitempty"`
	CleanTitle            string              `json:"cleanTitle,omitempty"`
	ImdbID                string              `json:"imdbId,omitempty"`
	TitleSlug             string              `json:"titleSlug,omitempty"`
	Certification         string              `json:"certification,omitempty"`
	Genres                []string            `json:"genres,omitempty"`
	Tags                  []int               `json:"tags,omitempty"`
	Added                 time.Time           `json:"added,omitempty"`
	Ratings               *starr.Ratings      `json:"ratings,omitempty"`
	MovieFile             *MovieFile          `json:"movieFile,omitempty"`
	Collection            *Collection         `json:"collection,omitempty"`
	HasFile               bool                `json:"hasFile,omitempty"`
	IsAvailable           bool                `json:"isAvailable,omitempty"`
	Monitored             bool                `json:"monitored,omitempty"`
}

// Collection belongs to a Movie.
type Collection struct {
	Name   string         `json:"name"`
	TmdbID int64          `json:"tmdbId"`
	Images []*starr.Image `json:"images"`
}

// MovieFile is part of a Movie.
type MovieFile struct {
	ID                  int64          `json:"id"`
	MovieID             int64          `json:"movieId"`
	RelativePath        string         `json:"relativePath"`
	Path                string         `json:"path"`
	Size                int64          `json:"size"`
	DateAdded           time.Time      `json:"dateAdded"`
	SceneName           string         `json:"sceneName"`
	IndexerFlags        int64          `json:"indexerFlags"`
	Quality             *starr.Quality `json:"quality"`
	MediaInfo           *MediaInfo     `json:"mediaInfo"`
	QualityCutoffNotMet bool           `json:"qualityCutoffNotMet"`
	Languages           []*starr.Value `json:"languages"`
	ReleaseGroup        string         `json:"releaseGroup"`
	Edition             string         `json:"edition"`
}

// MediaInfo is part of a MovieFile.
type MediaInfo struct {
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
}

// QualityProfile is applied to Movies.
type QualityProfile struct {
	ID                int64            `json:"id"`
	Name              string           `json:"name"`
	UpgradeAllowed    bool             `json:"upgradeAllowed"`
	Cutoff            int64            `json:"cutoff"`
	Qualities         []*starr.Quality `json:"items"`
	MinFormatScore    int64            `json:"minFormatScore"`
	CutoffFormatScore int64            `json:"cutoffFormatScore"`
	FormatItems       []interface{}    `json:"formatItems"`
	Language          *starr.Value     `json:"language"`
}

type Exclusion struct {
	TMDBID int64  `json:"tmdbId"`
	Title  string `json:"movieTitle"`
	Year   int    `json:"movieYear"`
	ID     int64  `json:"id,omitempty"`
}

// CustomFormat is the api/customformat endpoint payload.
type CustomFormat struct {
	ID                    int                 `json:"id"`
	Name                  string              `json:"name"`
	IncludeCFWhenRenaming bool                `json:"includeCustomFormatWhenRenaming"`
	Specifications        []*CustomFormatSpec `json:"specifications"`
}

// CustomFormatSpec is part of a CustomFormat.
type CustomFormatSpec struct {
	Name               string               `json:"name"`
	Implementation     string               `json:"implementation"`
	Implementationname string               `json:"implementationName"`
	Infolink           string               `json:"infoLink"`
	Negate             bool                 `json:"negate"`
	Required           bool                 `json:"required"`
	Fields             []*CustomFormatField `json:"fields"`
}

// CustomFormatField is part of a CustomFormat Specification.
type CustomFormatField struct {
	Order    int         `json:"order"`
	Name     string      `json:"name"`
	Label    string      `json:"label"`
	Value    interface{} `json:"value"` // should be a string, but sometimes it's a number.
	Type     string      `json:"type"`
	Advanced bool        `json:"advanced"`
}
