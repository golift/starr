//nolint:lll
package starrcmd

/*
All 9 Readarr events accounted for; 1/30/2022.
https://github.com/Readarr/Readarr/blob/develop/src/NzbDrone.Core/Notifications/CustomScript/CustomScript.cs
*/

import (
	"time"
)

// ReadarrApplicationUpdate is the ApplicationUpdate event.
type ReadarrApplicationUpdate struct {
	PreviousVersion string `env:"readarr_update_previousversion"` // 4.0.3.5875
	NewVersion      string `env:"readarr_update_newversion"`      // 4.0.4.5909
	Message         string `env:"readarr_update_message"`         // Readarr updated from 4.0.3.5875 to 4.0.4.5909
}

// ReadarrHealthIssue is the HealthIssue event.
type ReadarrHealthIssue struct {
	Message   string `env:"readarr_health_issue_message"` // Lists unavailable due to failures: List name here
	IssueType string `env:"readarr_health_issue_type"`    // ImportListStatusCheck
	Wiki      string `env:"readarr_health_issue_wiki"`    // https://wiki.servarr.com/
	Level     string `env:"readarr_health_issue_level"`   // Warning
}

// ReadarrGrab is the Grab event.
type ReadarrGrab struct {
	AuthorGRID     int64       `env:"readarr_author_grid"`                // 1077326
	ReleaseGroup   string      `env:"readarr_release_releasegroup"`       // BitBook
	AuthorName     string      `env:"readarr_author_name"`                // J.K. Rowling
	ReleaseTitle   string      `env:"readarr_release_title"`              // J K Rowling - Harry Potter and the Order of the Phoenix 2012 Retail EPUB eBook-BitBook
	GRIDs          string      `env:"readarr_release_grids"`              // 21175582 // not sure what this looks like with 2+
	DownloadClient string      `env:"readarr_download_client"`            // qBittorrent
	Size           int64       `env:"readarr_release_size"`               // 1279262
	QualityVersion string      `env:"readarr_release_qualityversion"`     // 1
	Titles         []string    `env:"readarr_release_booktitles,|"`       // Harry Potter and the Order of the Phoenix
	IDs            []int64     `env:"readarr_release_bookids,|"`          // 649
	ReleaseIndexer string      `env:"readarr_release_indexer"`            // InfoWars (Prowlarr)
	DownloadID     string      `env:"readarr_download_id"`                // 3852BA2204A84185B2B43281E53BE93D56DE5C81
	BookCount      int         `env:"readarr_release_bookcount"`          // 1
	ReleaseDates   []time.Time `env:"readarr_release_bookreleasedates,,"` // 07/10/2003 07:00:00
	Quality        string      `env:"readarr_release_quality"`            // EPUB
	AuthorID       int64       `env:"readarr_author_id"`                  // 4
}

// ReadarrBookDelete is the BookDelete event.
type ReadarrBookDelete struct {
	AuthorName   string `env:"readarr_author_name"`        // Alyssa Cole
	GrID         int64  `env:"readarr_book_goodreadsid"`   // 88514853
	AuthorGrID   int64  `env:"readarr_author_goodreadsid"` // 7790155
	Title        string `env:"readarr_book_title"`         // Unti Cole #6: A Novel
	Path         string `env:"readarr_author_path"`        // /books/Alyssa Cole
	ID           int64  `env:"readarr_book_id"`            // 636
	DeletedFiles bool   `env:"readarr_book_deletedfiles"`  // True
	AuthorID     string `env:"readarr_author_id"`          // 33
}

// ReadarrBookFileDelete is the BookFileDelete event.
type ReadarrBookFileDelete struct {
	Reason         string `env:"readarr_delete_reason"`                // deleteMessage.Reason.ToString())
	AuthorID       int64  `env:"readarr_author_id"`                    // author.Id.ToString())
	AuthorName     string `env:"readarr_author_name"`                  // author.Name)
	AuthorGrID     int64  `env:"readarr_author_goodreadsid"`           // author.ForeignAuthorId)
	ID             string `env:"readarr_book_id"`                      // book.Id.ToString())
	Title          string `env:"readarr_book_title"`                   // book.Title)
	GrID           int64  `env:"readarr_book_goodreadsid"`             // book.ForeignBookId)
	FileID         int64  `env:"readarr_bookfile_id"`                  // bookFile.Id.ToString())
	Path           string `env:"readarr_bookfile_path"`                // bookFile.Path)
	Quality        string `env:"readarr_bookfile_quality"`             // bookFile.Quality.Quality.Name)
	QualityVersion int64  `env:"readarr_bookfile_qualityversion"`      // bookFile.Quality.Revision.Version.ToString())
	ReleaseGroup   string `env:"readarr_bookfile_releasegroup"`        // bookFile.ReleaseGroup ?? string.Empty)
	SceneName      string `env:"readarr_bookfile_scenename"`           // bookFile.SceneName ?? string.Empty)
	EditionID      int64  `env:"readarr_bookfile_edition_id"`          // edition.Id.ToString())
	EditionName    string `env:"readarr_bookfile_edition_name"`        // edition.Title)
	EditionGrID    int64  `env:"readarr_bookfile_edition_goodreadsid"` // edition.ForeignEditionId)
	EditionISBN13  string `env:"readarr_bookfile_edition_isbn13"`      // edition.Isbn13)
	EditionASIN    string `env:"readarr_bookfile_edition_asin"`        // edition.Asin)
}

// ReadarrAuthorDelete is the AuthorDelete event.
type ReadarrAuthorDelete struct {
	AuthorID     int64  `env:"readarr_author_id"`           // author.Id.ToString())
	AuthorName   string `env:"readarr_author_name"`         // author.Name)
	Path         string `env:"readarr_author_path"`         // author.Path)
	AuthorGrID   int64  `env:"readarr_author_goodreadsid"`  // author.ForeignAuthorId)
	DeletedFiles bool   `env:"readarr_author_deletedfiles"` // deleteMessage.DeletedFiles.ToString())
}

// ReadarrRename is the Rename event.
type ReadarrRename struct {
	AuthorID   int64  `env:"readarr_author_id"`   // author.Id.ToString())
	AuthorName string `env:"readarr_author_name"` // author.Metadata.Value.Name)
	Path       string `env:"readarr_author_path"` // author.Path)
	AuthorGrID int64  `env:"readarr_author_grid"` // author.Metadata.Value.ForeignAuthorId)
}

// ReadarrDownload is Download event.
type ReadarrDownload struct {
	AuthorID       int64    `env:"readarr_author_id"`        // author.Id.ToString())
	AuthorName     string   `env:"readarr_author_name"`      // author.Metadata.Value.Name)
	Path           string   `env:"readarr_author_path"`      // author.Path)
	AuthorGrID     int64    `env:"readarr_author_grid"`      // author.Metadata.Value.ForeignAuthorId)
	ID             int64    `env:"readarr_book_id"`          // book.Id.ToString())
	Title          string   `env:"readarr_book_title"`       // book.Title)
	GrID           int64    `env:"readarr_book_grid"`        // book.Editions.Value.Single(e => e.Monitored).ForeignEditionId.ToString())
	ReleaseDate    string   `env:"readarr_book_releasedate"` // book.ReleaseDate.ToString())
	DownloadClient string   `env:"readarr_download_client"`  // message.DownloadClient ?? string.Empty)
	DownloadID     int64    `env:"readarr_download_id"`      // message.DownloadId ?? string.Empty)
	AddedBookPaths []string `env:"readarr_addedbookpaths,|"` // string.Join("|", message.BookFiles.Select(e => e.Path)))
	DeletedPaths   []string `env:"readarr_deletedpaths,|"`   // string.Join("|", message.OldFiles.Select(e => e.Path)))
}

// ReadarrTrackRetag is the TrackRetag event.
type ReadarrTrackRetag struct {
	AuthorID       int64     `env:"readarr_author_id"`               // author.Id.ToString())
	AuthorName     string    `env:"readarr_author_name"`             // author.Metadata.Value.Name)
	Path           string    `env:"readarr_author_path"`             // author.Path)
	AuthorGRId     int64     `env:"readarr_author_grid"`             // author.Metadata.Value.ForeignAuthorId)
	ID             int64     `env:"readarr_book_id"`                 // book.Id.ToString())
	Title          string    `env:"readarr_book_title"`              // book.Title)
	GrID           int64     `env:"readarr_book_grid"`               // book.Editions.Value.Single(e => e.Monitored).ForeignEditionId.ToString())
	ReleaseDate    time.Time `env:"readarr_book_releasedate"`        // book.ReleaseDate.ToString())
	FileID         int64     `env:"readarr_bookfile_id"`             // bookFile.Id.ToString())
	FilePath       string    `env:"readarr_bookfile_path"`           // bookFile.Path)
	Quality        string    `env:"readarr_bookfile_quality"`        // bookFile.Quality.Quality.Name)
	QualityVersion int64     `env:"readarr_bookfile_qualityversion"` // bookFile.Quality.Revision.Version.ToString())
	ReleaseGroup   string    `env:"readarr_bookfile_releasegroup"`   // bookFile.ReleaseGroup ?? string.Empty)
	SceneName      string    `env:"readarr_bookfile_scenename"`      // bookFile.SceneName ?? string.Empty)
	TagsDiff       string    `env:"readarr_tags_diff"`               // message.Diff.ToJson())
	Scrubbed       bool      `env:"readarr_tags_scrubbed"`           // message.Scrubbed.ToString())
}

// ReadarrTest has no members.
type ReadarrTest struct{}

// GetReadarrApplicationUpdate returns the ApplicationUpdate event data.
func (c *CmdEvent) GetReadarrApplicationUpdate() (output ReadarrApplicationUpdate, err error) {
	return output, c.get(EventApplicationUpdate, &output)
}

// GetReadarrHealthIssue returns the ApplicationUpdate event data.
func (c *CmdEvent) GetReadarrHealthIssue() (output ReadarrHealthIssue, err error) {
	return output, c.get(EventHealthIssue, &output)
}

// GetReadarrGrab returns the Grab event data.
func (c *CmdEvent) GetReadarrGrab() (output ReadarrGrab, err error) {
	return output, c.get(EventGrab, &output)
}

// GetReadarrBookDelete returns the BookDelete event data.
func (c *CmdEvent) GetReadarrBookDelete() (output ReadarrBookDelete, err error) {
	return output, c.get(EventBookDelete, &output)
}

// GetReadarrBookFileDelete returns the BookFileDelete event data.
func (c *CmdEvent) GetReadarrBookFileDelete() (output ReadarrBookFileDelete, err error) {
	return output, c.get(EventBookFileDelete, &output)
}

// GetReadarrDownload returns the Download event data.
func (c *CmdEvent) GetReadarrDownload() (output ReadarrDownload, err error) {
	return output, c.get(EventDownload, &output)
}

// GetReadarrRename returns the Rename event data.
func (c *CmdEvent) GetReadarrRename() (output ReadarrRename, err error) {
	return output, c.get(EventRename, &output)
}

// GetReadarrTrackRetag returns the TrackRetag event data.
func (c *CmdEvent) GetReadarrTrackRetag() (output ReadarrTrackRetag, err error) {
	return output, c.get(EventTrackRetag, &output)
}

// GetReadarrTest returns the ApplicationUpdate event data.
func (c *CmdEvent) GetReadarrTest() (output ReadarrTest, err error) {
	return output, c.get(EventTest, &output)
}
