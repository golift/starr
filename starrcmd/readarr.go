//nolint:lll
package starrcmd

/*
INCOMPLETE.
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

// ReadarrTest has no members.
type ReadarrTest struct{}

// GetReadarrApplicationUpdate returns the ApplicationUpdate event data.
func GetReadarrApplicationUpdate() (output ReadarrApplicationUpdate, err error) {
	return output, get(EventApplicationUpdate, &output)
}

// GetReadarrHealthIssue returns the ApplicationUpdate event data.
func GetReadarrHealthIssue() (output ReadarrHealthIssue, err error) {
	return output, get(EventHealthIssue, &output)
}

// GetReadarrGrab returns the Grab event data.
func GetReadarrGrab() (output ReadarrGrab, err error) {
	return output, get(EventGrab, &output)
}

// GetReadarrTest returns the ApplicationUpdate event data.
func GetReadarrTest() (output ReadarrTest, err error) {
	return output, get(EventTest, &output)
}
