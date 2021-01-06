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
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: !c.ValidSSL},
			},
		}
	}

	return &Readarr{APIer: c}
}

// Queue is the /api/v1/queue endpoint.
type Queue struct {
	Page          int          `json:"page"`
	PageSize      int          `json:"pageSize"`
	SortKey       string       `json:"sortKey"`
	SortDirection string       `json:"sortDirection"`
	TotalRecords  int          `json:"totalRecords"`
	Records       []BookRecord `json:"records"`
}

// BookRecord is a book from the queue API path.
type BookRecord struct {
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
	OutputProfile            int64  `json:"outputProfile"`
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
	Monitored         bool           `json:"monitored,omitempty"`
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
	AddOptions    *AddBookOptions   `json:"addOptions"`    // Contains Search.
	Author        *AddBookAuthor    `json:"author"`        // Contains Author ID
	Editions      []*AddBookEdition `json:"editions"`      // contains GRID Edition ID
	ForeignBookID int64             `json:"foreignBookId"` // GRID Book ID.
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
	Monitored        bool   `json:"monitored"`        // true
	ManualAdd        bool   `json:"manualAdd"`        // true
	ForeignEditionID string `json:"foreignEditionId"` // GRID ID
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
