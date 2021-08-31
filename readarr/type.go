package readarr

import (
	"crypto/tls"
	"net/http"
	"time"

	"golift.io/starr"
)

// Readarr contains all the methods to interact with a Readarr server.
type Readarr struct {
	starr.APIer
}

// New returns a Readarr object used to interact with the Readarr API.
func New(c *starr.Config) *Readarr {
	if c.Client == nil {
		//nolint:exhaustivestruct,gosec
		c.Client = &http.Client{
			Timeout: c.Timeout.Duration,
			CheckRedirect: func(r *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: !c.ValidSSL},
			},
		}
	}

	if c.Debugf == nil {
		c.Debugf = func(string, ...interface{}) {}
	}

	return &Readarr{APIer: c}
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

// QueueRecord is a book from the queue API path.
type QueueRecord struct {
	AuthorID                int64                  `json:"authorId"`
	BookID                  int64                  `json:"bookId"`
	Quality                 *starr.Quality         `json:"quality"`
	Size                    float64                `json:"size"`
	Title                   string                 `json:"title"`
	Sizeleft                float64                `json:"sizeleft"`
	Timeleft                string                 `json:"timeleft"`
	EstimatedCompletionTime time.Time              `json:"estimatedCompletionTime"`
	Status                  string                 `json:"status"`
	TrackedDownloadStatus   string                 `json:"trackedDownloadStatus,omitempty"`
	TrackedDownloadState    string                 `json:"trackedDownloadState,omitempty"`
	StatusMessages          []*starr.StatusMessage `json:"statusMessages,omitempty"`
	DownloadID              string                 `json:"downloadId,omitempty"`
	Protocol                string                 `json:"protocol"`
	DownloadClient          string                 `json:"downloadClient,omitempty"`
	Indexer                 string                 `json:"indexer"`
	OutputPath              string                 `json:"outputPath,omitempty"`
	DownloadForced          bool                   `json:"downloadForced"`
	ID                      int64                  `json:"id"`
	ErrorMessage            string                 `json:"errorMessage"`
}

// SystemStatus is the /api/v1/system/status endpoint.
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
	IsNetCore              bool      `json:"isNetCore"`
	IsMono                 bool      `json:"isMono"`
	IsLinux                bool      `json:"isLinux"`
	IsOsx                  bool      `json:"isOsx"`
	IsWindows              bool      `json:"isWindows"`
	IsDocker               bool      `json:"isDocker"`
	Mode                   string    `json:"mode"`
	Branch                 string    `json:"branch"`
	Authentication         string    `json:"authentication"`
	SqliteVersion          string    `json:"sqliteVersion"`
	MigrationVersion       int       `json:"migrationVersion"`
	URLBase                string    `json:"urlBase"`
	RuntimeVersion         string    `json:"runtimeVersion"`
	RuntimeName            string    `json:"runtimeName"`
	StartTime              time.Time `json:"startTime"`
	PackageVersion         string    `json:"packageVersion"`
	PackageAuthor          string    `json:"packageAuthor"`
	PackageUpdateMechanism string    `json:"packageUpdateMechanism"`
}

// RootFolder is the /api/v1/rootfolder endpoint.
type RootFolder struct {
	ID                       int64  `json:"id"`
	Name                     string `json:"name"`
	Path                     string `json:"path"`
	DefaultMetadataProfileID int64  `json:"defaultMetadataProfileId"`
	DefaultQualityProfileID  int64  `json:"defaultQualityProfileId"`
	DefaultMonitorOption     string `json:"defaultMonitorOption"`
	DefaultTags              []int  `json:"defaultTags"`
	Port                     int    `json:"port"`
	OutputProfile            string `json:"outputProfile"`
	UseSsl                   bool   `json:"useSsl"`
	Accessible               bool   `json:"accessible"`
	IsCalibreLibrary         bool   `json:"isCalibreLibrary"`
	FreeSpace                int64  `json:"freeSpace"`
	TotalSpace               int64  `json:"totalSpace"`
}

// QualityProfile is the /api/v1/qualityprofile endpoint.
type QualityProfile struct {
	Name           string           `json:"name"`
	UpgradeAllowed bool             `json:"upgradeAllowed"`
	Cutoff         int64            `json:"cutoff"`
	Qualities      []*starr.Quality `json:"items"`
	ID             int64            `json:"id"`
}

// MetadataProfile is the /api/v1/metadataProfile endpoint.
type MetadataProfile struct {
	ID                  int64   `json:"id"`
	Name                string  `json:"name"`
	MinPopularity       float64 `json:"minPopularity"`
	SkipMissingDate     bool    `json:"skipMissingDate"`
	SkipMissingIsbn     bool    `json:"skipMissingIsbn"`
	SkipPartsAndSets    bool    `json:"skipPartsAndSets"`
	SkipSeriesSecondary bool    `json:"skipSeriesSecondary"`
	AllowedLanguages    string  `json:"allowedLanguages,omitempty"`
}

// Author is the /api/v1/author endpoint.
type Author struct {
	ID                int64          `json:"id"`
	Status            string         `json:"status,omitempty"`
	AuthorName        string         `json:"authorName,omitempty"`
	ForeignAuthorID   string         `json:"foreignAuthorId,omitempty"`
	TitleSlug         string         `json:"titleSlug,omitempty"`
	Overview          string         `json:"overview,omitempty"`
	Links             []*starr.Link  `json:"links,omitempty"`
	Images            []*starr.Image `json:"images,omitempty"`
	Path              string         `json:"path,omitempty"`
	QualityProfileID  int            `json:"qualityProfileId,omitempty"`
	MetadataProfileID int            `json:"metadataProfileId,omitempty"`
	Genres            []interface{}  `json:"genres,omitempty"`
	CleanName         string         `json:"cleanName,omitempty"`
	SortName          string         `json:"sortName,omitempty"`
	Tags              []int          `json:"tags,omitempty"`
	Added             time.Time      `json:"added,omitempty"`
	Ratings           *starr.Ratings `json:"ratings,omitempty"`
	Statistics        *Statistics    `json:"statistics,omitempty"`
	LastBook          *AuthorBook    `json:"lastBook,omitempty"`
	NextBook          *AuthorBook    `json:"nextBook,omitempty"`
	Ended             bool           `json:"ended,omitempty"`
	Monitored         bool           `json:"monitored"`
}

// AuthorBook is part of an Author.
type AuthorBook struct {
	ID               int64           `json:"id"`
	AuthorMetadataID int             `json:"authorMetadataId"`
	ForeignBookID    string          `json:"foreignBookId"`
	TitleSlug        string          `json:"titleSlug"`
	Title            string          `json:"title"`
	ReleaseDate      time.Time       `json:"releaseDate"`
	Links            []*starr.Link   `json:"links"`
	Genres           []interface{}   `json:"genres"`
	Ratings          *starr.Ratings  `json:"ratings"`
	CleanTitle       string          `json:"cleanTitle"`
	Monitored        bool            `json:"monitored"`
	AnyEditionOk     bool            `json:"anyEditionOk"`
	LastInfoSync     time.Time       `json:"lastInfoSync"`
	Added            time.Time       `json:"added"`
	AddOptions       *AddBookOptions `json:"addOptions"`
	AuthorMetadata   *starr.IsLoaded `json:"authorMetadata"`
	Author           *starr.IsLoaded `json:"author"`
	Editions         *starr.IsLoaded `json:"editions"`
	BookFiles        *starr.IsLoaded `json:"bookFiles"`
	SeriesLinks      *starr.IsLoaded `json:"seriesLinks"`
}

// Book is the /api/v1/book endpoint.
type Book struct {
	Title          string         `json:"title"`
	SeriesTitle    string         `json:"seriesTitle"`
	Overview       string         `json:"overview"`
	AuthorID       int64          `json:"authorId"`
	ForeignBookID  string         `json:"foreignBookId"`
	TitleSlug      string         `json:"titleSlug"`
	Monitored      bool           `json:"monitored"`
	AnyEditionOk   bool           `json:"anyEditionOk"`
	Ratings        *starr.Ratings `json:"ratings"`
	ReleaseDate    time.Time      `json:"releaseDate"`
	PageCount      int            `json:"pageCount"`
	Genres         []interface{}  `json:"genres"`
	Author         *BookAuthor    `json:"author,omitempty"`
	Images         []*starr.Image `json:"images"`
	Links          []*starr.Link  `json:"links"`
	Statistics     *Statistics    `json:"statistics,omitempty"`
	Editions       []*Edition     `json:"editions"`
	ID             int64          `json:"id"`
	Disambiguation string         `json:"disambiguation,omitempty"`
}

// Statistics for a Book, or maybe an author.
type Statistics struct {
	BookCount      int     `json:"bookCount"`
	BookFileCount  int     `json:"bookFileCount"`
	TotalBookCount int     `json:"totalBookCount"`
	SizeOnDisk     int     `json:"sizeOnDisk"`
	PercentOfBooks float64 `json:"percentOfBooks"`
}

// BookAuthor of a Book.
type BookAuthor struct {
	ID                int64          `json:"id"`
	Status            string         `json:"status"`
	AuthorName        string         `json:"authorName"`
	ForeignAuthorID   string         `json:"foreignAuthorId"`
	TitleSlug         string         `json:"titleSlug"`
	Overview          string         `json:"overview"`
	Links             []*starr.Link  `json:"links"`
	Images            []*starr.Image `json:"images"`
	Path              string         `json:"path"`
	QualityProfileID  int64          `json:"qualityProfileId"`
	MetadataProfileID int64          `json:"metadataProfileId"`
	Genres            []interface{}  `json:"genres"`
	CleanName         string         `json:"cleanName"`
	SortName          string         `json:"sortName"`
	Tags              []int          `json:"tags"`
	Added             time.Time      `json:"added"`
	Ratings           *starr.Ratings `json:"ratings"`
	Statistics        *Statistics    `json:"statistics"`
	Monitored         bool           `json:"monitored"`
	Ended             bool           `json:"ended"`
}

// Edition is more Book meta data.
type Edition struct {
	ID               int64          `json:"id"`
	BookID           int64          `json:"bookId"`
	ForeignEditionID string         `json:"foreignEditionId"`
	TitleSlug        string         `json:"titleSlug"`
	Isbn13           string         `json:"isbn13"`
	Asin             string         `json:"asin"`
	Title            string         `json:"title"`
	Overview         string         `json:"overview"`
	Format           string         `json:"format"`
	Publisher        string         `json:"publisher"`
	PageCount        int            `json:"pageCount"`
	ReleaseDate      time.Time      `json:"releaseDate"`
	Images           []*starr.Image `json:"images"`
	Links            []*starr.Link  `json:"links"`
	Ratings          *starr.Ratings `json:"ratings"`
	Monitored        bool           `json:"monitored"`
	ManualAdd        bool           `json:"manualAdd"`
	IsEbook          bool           `json:"isEbook"`
}

// AddBookInput is the input to add a book.
type AddBookInput struct {
	Monitored     bool              `json:"monitored"`
	Tags          []int             `json:"tags"`
	AddOptions    *AddBookOptions   `json:"addOptions"`    // Contains Search.
	Author        *AddBookAuthor    `json:"author"`        // Contains Author ID
	Editions      []*AddBookEdition `json:"editions"`      // contains GRID Edition ID
	ForeignBookID string            `json:"foreignBookId"` // GRID Book ID.
}

// AddBookAuthor is part of AddBookInput.
type AddBookAuthor struct {
	Monitored         bool              `json:"monitored"`         // true?
	QualityProfileID  int64             `json:"qualityProfileId"`  // required
	MetadataProfileID int64             `json:"metadataProfileId"` // required
	ForeignAuthorID   string            `json:"foreignAuthorId"`   // required
	RootFolderPath    string            `json:"rootFolderPath"`    // required
	AddOptions        *AddAuthorOptions `json:"addOptions"`
}

// AddAuthorOptions is part of AddBookAuthor.
type AddAuthorOptions struct {
	SearchForMissingBooks bool    `json:"searchForMissingBooks"`
	Monitored             bool    `json:"monitored"`
	Monitor               string  `json:"monitor"`
	BooksToMonitor        []int64 `json:"booksToMonitor"`
}

// AddBookOptions is part of AddBookInput.
type AddBookOptions struct {
	AddType          string `json:"addType,omitempty"`
	SearchForNewBook bool   `json:"searchForNewBook"`
}

// AddBookEdition is part of AddBookInput.
type AddBookEdition struct {
	Title            string         `json:"title"`            // Edition Title
	TitleSlug        interface{}    `json:"titleSlug"`        // Slugs are dumb
	Images           []*starr.Image `json:"images"`           // this is dumb too
	ForeignEditionID string         `json:"foreignEditionId"` // GRID ID
	Monitored        bool           `json:"monitored"`        // true
	ManualAdd        bool           `json:"manualAdd"`        // true
}

// AddBookOutput is returned when a book is added.
type AddBookOutput struct {
	ID            int64          `json:"id"`
	AuthorID      int64          `json:"authorId"`
	PageCount     int            `json:"pageCount"`
	Title         string         `json:"title"`
	SeriesTitle   string         `json:"seriesTitle"`
	Overview      string         `json:"overview"`
	ForeignBookID string         `json:"foreignBookId"`
	TitleSlug     string         `json:"titleSlug"`
	Ratings       *starr.Ratings `json:"ratings"`
	ReleaseDate   time.Time      `json:"releaseDate"`
	Genres        []interface{}  `json:"genres"`
	Author        *BookAuthor    `json:"author"`
	Images        []*starr.Image `json:"images"`
	Links         []*starr.Link  `json:"links"`
	Statistics    *Statistics    `json:"statistics"`
	Editions      []*Edition     `json:"editions"`
	Monitored     bool           `json:"monitored"`
	AnyEditionOk  bool           `json:"anyEditionOk"`
}

// CommandReqyest goes into the /api/v1/command endpoint.
// This was created from the search command and may not support other commands yet.
type CommandRequest struct {
	Name    string  `json:"name"`
	BookIDs []int64 `json:"bookIds,omitempty"`
	BookID  int64   `json:"bookId,omitempty"`
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

// History is the /api/v1/history endpoint.
type History struct {
	Page          int             `json:"page"`
	PageSize      int             `json:"pageSize"`
	SortKey       string          `json:"sortKey"`
	SortDirection string          `json:"sortDirection"`
	TotalRecords  int             `json:"totalRecords"`
	Records       []HistoryRecord `json:"records"`
}

// HistoryRecord is part of the history. Not all items have all Data members.
// Check EventType for events you need.
type HistoryRecord struct {
	ID                  int64          `json:"id"`
	BookID              int64          `json:"bookId"`
	AuthorID            int64          `json:"authorId"`
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
