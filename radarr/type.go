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
func New(config *starr.Config) *Radarr {
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

	return &Radarr{APIer: config}
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

// Queue is the /api/v3/queue endpoint.
type Queue struct {
	Page          int            `json:"page"`
	PageSize      int            `json:"pageSize"`
	SortKey       string         `json:"sortKey"`
	SortDirection string         `json:"sortDirection"`
	TotalRecords  int            `json:"totalRecords"`
	Records       []*QueueRecord `json:"records"`
}

// QueueRecord is part of the activity Queue.
type QueueRecord struct {
	MovieID                 int64                  `json:"movieId"`
	Languages               []*starr.Value         `json:"languages"`
	Quality                 *starr.Quality         `json:"quality"`
	CustomFormats           []interface{}          `json:"customFormats"` // probably []int64
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
	ID                      int64                  `json:"id"`
	ErrorMessage            string                 `json:"errorMessage"`
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
	Monitored             bool                `json:"monitored"`
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
	FormatItems       []*FormatItem    `json:"formatItems,omitempty"`
	Language          *starr.Value     `json:"language"`
}

// FormatItem is part of a QualityProfile.
type FormatItem struct {
	Format int    `json:"format"`
	Name   string `json:"name"`
	Score  int    `json:"score"`
}

// Exclusion is a Radarr excluded item.
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

// CommandRequest goes into the /api/v3/command endpoint.
// This was created from the search command and may not support other commands yet.
type CommandRequest struct {
	Name     string  `json:"name"`
	MovieIDs []int64 `json:"movieIds,omitempty"`
}

// CommandResponse comes from the /api/v3/command endpoint.
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

// History is the /api/v3/history endpoint.
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
	ID                  int64          `json:"id"`
	MovieID             int64          `json:"movieId"`
	SourceTitle         string         `json:"sourceTitle"`
	Languages           []*starr.Value `json:"languages"`
	Quality             *starr.Quality `json:"quality"`
	CustomFormats       []interface{}  `json:"customFormats"`
	QualityCutoffNotMet bool           `json:"qualityCutoffNotMet"`
	Date                time.Time      `json:"date"`
	DownloadID          string         `json:"downloadId"`
	EventType           string         `json:"eventType"`
	Data                struct {
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
		IndexerFlags       string    `json:"indexerFlags"`
		IndexerID          string    `json:"indexerId"`
		Message            string    `json:"message"`
		NzbInfoURL         string    `json:"nzbInfoUrl"`
		Protocol           string    `json:"protocol"`
		PublishedDate      time.Time `json:"publishedDate"`
		Reason             string    `json:"reason"`
		ReleaseGroup       string    `json:"releaseGroup"`
		Size               string    `json:"size"`
		TmdbID             string    `json:"tmdbId"`
		TorrentInfoHash    string    `json:"torrentInfoHash"`
	} `json:"data"`
}

// ImportList represents the api/v3/importlist endpoint.
type ImportList struct {
	ID                  int64    `json:"id"`
	Name                string   `json:"name"`
	Enabled             bool     `json:"enabled"`
	EnableAuto          bool     `json:"enableAuto"`
	ShouldMonitor       bool     `json:"shouldMonitor"`
	SearchOnAdd         bool     `json:"searchOnAdd"`
	RootFolderPath      string   `json:"rootFolderPath"`
	QualityProfileID    int64    `json:"qualityProfileId"`
	MinimumAvailability string   `json:"minimumAvailability"`
	ListType            string   `json:"listType"`
	ListOrder           int64    `json:"listOrder"`
	Fields              []*Field `json:"fields"`
	ImplementationName  string   `json:"implementationName"`
	Implementation      string   `json:"implementation"`
	ConfigContract      string   `json:"configContract"`
	InfoLink            string   `json:"infoLink"`
	Tags                []int    `json:"tags"`
}

// Field is currently only part of ImportList.
type Field struct {
	Name          string          `json:"name"`
	Value         interface{}     `json:"value"` // sometimes number, sometimes string. 'Type' may tell you.
	Label         string          `json:"label"`
	HelpText      string          `json:"helpText"`
	Type          string          `json:"type"`
	Order         int64           `json:"order"`
	Advanced      bool            `json:"advanced"`
	SelectOptions []*SelectOption `json:"selectOptions,omitempty"`
}

// SelectOption is part of a Field from an ImportList.
type SelectOption struct {
	Value        int    `json:"value"`
	Name         string `json:"name"`
	Order        int    `json:"order"`
	DividerAfter bool   `json:"dividerAfter"`
}
