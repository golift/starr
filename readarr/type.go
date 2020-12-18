package readarr

import (
	"crypto/tls"
	"net/http"
	"time"

	"golift.io/starr"
)

type Readarr struct {
	starr.APIer
}

func New(c *starr.Config) *Readarr {
	if c.Client == nil {
		c.Client = &http.Client{ // nolint: exhaustivestruct
			Timeout: c.Timeout.Duration,
			Transport: &http.Transport{ // nolint: exhaustivestruct
				TLSClientConfig: &tls.Config{InsecureSkipVerify: !c.ValidSSL}, // nolint: gosec, exhaustivestruct
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
	Name                     string        `json:"name"`
	Path                     string        `json:"path"`
	DefaultMetadataProfileID int           `json:"defaultMetadataProfileId"`
	DefaultQualityProfileID  int           `json:"defaultQualityProfileId"`
	DefaultMonitorOption     string        `json:"defaultMonitorOption"`
	DefaultTags              []interface{} `json:"defaultTags"`
	IsCalibreLibrary         bool          `json:"isCalibreLibrary"`
	Port                     int           `json:"port"`
	OutputProfile            int           `json:"outputProfile"`
	UseSsl                   bool          `json:"useSsl"`
	Accessible               bool          `json:"accessible"`
	FreeSpace                int64         `json:"freeSpace"`
	TotalSpace               int64         `json:"totalSpace"`
	ID                       int           `json:"id"`
}

// QualityProfile is the /api/v1/qualityprofile endpoint.
type QualityProfile struct {
	Name           string `json:"name"`
	UpgradeAllowed bool   `json:"upgradeAllowed"`
	Cutoff         int    `json:"cutoff"`
	Items          []struct {
		Quality struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"quality"`
		Items   []interface{} `json:"items"`
		Allowed bool          `json:"allowed"`
	} `json:"items"`
	ID int `json:"id"`
}

// MetadataProfile is the /api/v1/metadataProfile endpoint.
type MetadataProfile struct {
	Name                string  `json:"name"`
	MinPopularity       float64 `json:"minPopularity"`
	SkipMissingDate     bool    `json:"skipMissingDate"`
	SkipMissingIsbn     bool    `json:"skipMissingIsbn"`
	SkipPartsAndSets    bool    `json:"skipPartsAndSets"`
	SkipSeriesSecondary bool    `json:"skipSeriesSecondary"`
	AllowedLanguages    string  `json:"allowedLanguages,omitempty"`
	ID                  int     `json:"id"`
}

// Book is the /api/v1/book endpoint.
type Book struct {
	Title          string          `json:"title"`
	SeriesTitle    string          `json:"seriesTitle"`
	Overview       string          `json:"overview"`
	AuthorID       int             `json:"authorId"`
	ForeignBookID  string          `json:"foreignBookId"`
	TitleSlug      string          `json:"titleSlug"`
	Monitored      bool            `json:"monitored"`
	AnyEditionOk   bool            `json:"anyEditionOk"`
	Ratings        *Ratings        `json:"ratings"`
	ReleaseDate    time.Time       `json:"releaseDate"`
	PageCount      int             `json:"pageCount"`
	Genres         []interface{}   `json:"genres"`
	Author         *BookAuthor     `json:"author,omitempty"`
	Images         []*starr.Image  `json:"images"`
	Links          []*starr.Link   `json:"links"`
	Statistics     *Statistics     `json:"statistics,omitempty"`
	Editions       []*BookEditions `json:"editions"`
	ID             int             `json:"id"`
	Disambiguation string          `json:"disambiguation,omitempty"`
}

// Ratings belong to a Book.
type Ratings struct {
	Votes      int     `json:"votes"`
	Value      float64 `json:"value"`
	Popularity float64 `json:"popularity"`
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
	Status            string         `json:"status"`
	Ended             bool           `json:"ended"`
	AuthorName        string         `json:"authorName"`
	ForeignAuthorID   string         `json:"foreignAuthorId"`
	TitleSlug         string         `json:"titleSlug"`
	Overview          string         `json:"overview"`
	Links             []*starr.Link  `json:"links"`
	Images            []*starr.Image `json:"images"`
	Path              string         `json:"path"`
	QualityProfileID  int            `json:"qualityProfileId"`
	MetadataProfileID int            `json:"metadataProfileId"`
	Monitored         bool           `json:"monitored"`
	Genres            []interface{}  `json:"genres"`
	CleanName         string         `json:"cleanName"`
	SortName          string         `json:"sortName"`
	Tags              []interface{}  `json:"tags"`
	Added             time.Time      `json:"added"`
	Ratings           *Ratings       `json:"ratings"`
	Statistics        *Statistics    `json:"statistics"`
	ID                int            `json:"id"`
}

// BookEditions is more Book meta data.
type BookEditions struct {
	BookID           int            `json:"bookId"`
	ForeignEditionID string         `json:"foreignEditionId"`
	TitleSlug        string         `json:"titleSlug"`
	Isbn13           string         `json:"isbn13"`
	Asin             string         `json:"asin"`
	Title            string         `json:"title"`
	Overview         string         `json:"overview"`
	Format           string         `json:"format"`
	IsEbook          bool           `json:"isEbook"`
	Publisher        string         `json:"publisher"`
	PageCount        int            `json:"pageCount"`
	ReleaseDate      time.Time      `json:"releaseDate"`
	Images           []*starr.Image `json:"images"`
	Links            []*starr.Link  `json:"links"`
	Ratings          *Ratings       `json:"ratings"`
	Monitored        bool           `json:"monitored"`
	ManualAdd        bool           `json:"manualAdd"`
	ID               int            `json:"id"`
}

/* These AddBook types are highly incomplete as Readarr is alpha atm and still changing. */

// AddBookInput is the input to add a book.
type AddBookInput struct {
	Monitored     bool             `json:"monitored"`
	AddOptions    AddBookOptions   `json:"addOptions"`    // Contains Search.
	Author        AddBookAuthor    `json:"author"`        // Contains Author ID
	Editions      []AddBookEdition `json:"editions"`      // contains GRID Edition ID
	ForeignBookID int              `json:"foreignBookId"` // GRID Book ID.
}

// AddBookAuthor is part of AddBookInput.
type AddBookAuthor struct {
	Monitored         bool   `json:"monitored"`         // true?
	QualityProfileID  int    `json:"qualityProfileId"`  // required
	MetadataProfileID int    `json:"metadataProfileId"` // required
	ForeignAuthorID   string `json:"foreignAuthorId"`   // required
	RootFolderPath    string `json:"rootFolderPath"`    // required
}

// AddBookOptions is part of AddBookInput.
type AddBookOptions struct {
	SearchForNewBook bool `json:"addOptions"` // true
}

// AddBookEdition is part of AddBookInput.
type AddBookEdition struct {
	Monitored        bool `json:"monitored"`        // true
	ManualAdd        bool `json:"manualAdd"`        // true
	ForeignEditionID int  `json:"foreignEditionId"` // GRID ID
}

// AddBookOutput is a placeholder until I know what this looks like.
type AddBookOutput interface{}
