//nolint:lll
package starrcmd

/*
All events accounted for; 1/30/2022
https://github.com/Radarr/Radarr/blob/develop/src/NzbDrone.Core/Notifications/CustomScript/CustomScript.cs
*/

import (
	"time"
)

// RadarrApplicationUpdate is the ApplicationUpdate event.
type RadarrApplicationUpdate struct {
	PreviousVersion string `env:"radarr_update_previousversion"` // 4.0.3.5875
	NewVersion      string `env:"radarr_update_newversion"`      // 4.0.4.5909
	Message         string `env:"radarr_update_message"`         // Radarr updated from 4.0.3.5875 to 4.0.4.5909
}

// RadarrDownload is the Download event.
type RadarrDownload struct {
	ReleaseDate          time.Time `env:"radarr_movie_physical_release_date"`
	InCinemas            time.Time `env:"radarr_movie_in_cinemas_date"`  // 2/10/2011 12:00:00 AM
	FilePath             string    `env:"radarr_moviefile_path"`         // /movies/Just Go with It (2011)/Just.Go.with.It.2011.Bluray-1080p.mkv
	IMDbID               string    `env:"radarr_movie_imdbid"`           // tt1564367
	SceneName            string    `env:"radarr_moviefile_scenename"`    // Just.Go.with.It.2011.1080p.BluRay.x264-OFT
	ReleaseGroup         string    `env:"radarr_moviefile_releasegroup"` // OFT
	DownloadID           string    `env:"radarr_download_id"`            // string F3D870942BFDD643488852284E917336170CEA00
	SourceFolder         string    `env:"radarr_moviefile_sourcefolder"` // /downloads/Seeding/Just.Go.with.It.2011.1080p.BluRay.x264-OFT
	Path                 string    `env:"radarr_movie_path"`             // /movies/Just Go with It (2011)
	RelativePath         string    `env:"radarr_moviefile_relativepath"` // Just.Go.with.It.2011.Bluray-1080p.mkv
	DownloadClient       string    `env:"radarr_download_client"`        // Deluge
	SourcePath           string    `env:"radarr_moviefile_sourcepath"`   // /downloads/Seeding/Just.Go.with.It.2011.1080p.BluRay.x264-OFT/Just.Go.with.It.2011.1080p.BluRay.x264-OFT.mkv
	Quality              string    `env:"radarr_moviefile_quality"`      // Bluray-1080p
	Title                string    `env:"radarr_movie_title"`            // Just Go with It
	DeletedRelativePaths []string  `env:"radarr_deletedrelativepaths,|"`
	DeletedPaths         []string  `env:"radarr_deletedpaths,|"`
	FileID               int64     `env:"radarr_moviefile_id"`             // 3594
	Year                 int       `env:"radarr_movie_year"`               // 2011
	TMDbID               int64     `env:"radarr_movie_tmdbid"`             // 50546
	ID                   int64     `env:"radarr_movie_id"`                 // 924
	QualityVersion       int64     `env:"radarr_moviefile_qualityversion"` // 1
	IsUpgrade            bool      `env:"radarr_isupgrade"`                // False
}

// RadarrGrab is the Grab event.
type RadarrGrab struct {
	ReleaseDate    time.Time `env:"radarr_movie_physical_release_date"` // 1/19/2006 12:00:00 AM
	InCinemas      time.Time `env:"radarr_movie_in_cinemas_date"`       // 11/22/2005 12:00:00 AM
	ReleaseGroup   string    `env:"radarr_release_releasegroup"`        // SLOT
	IMDbID         string    `env:"radarr_movie_imdbid"`                // tt0448172
	DownloadID     string    `env:"radarr_download_id"`                 // E63FAFFAAA0DEE42F0846348A9C0657BC53E7AA5
	ReleaseTitle   string    `env:"radarr_release_title"`               // 8MM 2 2005 1080p BluRay x264
	Quality        string    `env:"radarr_release_quality"`             // Bluray-1080p
	DownloadClient string    `env:"radarr_download_client"`             // Deluge
	ReleaseIndexer string    `env:"radarr_release_indexer"`             // Inexilator (Prowlarr)
	Title          string    `env:"radarr_movie_title"`                 // 8MM 2
	QualityVersion int64     `env:"radarr_release_qualityversion"`      // 1
	IndexerFlags   int64     `env:"radarr_indexerflags"`                // 0
	Size           int64     `env:"radarr_release_size"`                // 2158221056
	Year           int       `env:"radarr_movie_year"`                  // 2005
	TMDbID         int64     `env:"radarr_movie_tmdbid"`                // 7295
	ID             int64     `env:"radarr_movie_id"`                    // 339
}

// RadarrHealthIssue is the HealthIssue event.
type RadarrHealthIssue struct {
	Message   string `env:"radarr_health_issue_message"` // Lists unavailable due to failures: List name here
	IssueType string `env:"radarr_health_issue_type"`    // ImportListStatusCheck
	Wiki      string `env:"radarr_health_issue_wiki"`    // https://wiki.servarr.com/radarr/system#lists-are-unavailable-due-to-failures
	Level     string `env:"radarr_health_issue_level"`   // Warning
}

// RadarrMovieFileDelete is the MovieFileDelete event.
type RadarrMovieFileDelete struct {
	Reason         string `env:"radarr_moviefile_deletereason"`   // Upgrade
	FilePath       string `env:"radarr_moviefile_path"`           // /movies/The French Dispatch (2021)/The.French.Dispatch.2021.Bluray-720p.mkv
	SceneName      string `env:"radarr_moviefile_scenename"`      // The.French.Dispatch.2021.720p.BluRay.x264-WoAT
	IMDbID         string `env:"radarr_movie_imdbid"`             // tt8847712
	ReleaseGroup   string `env:"radarr_moviefile_releasegroup"`   // WoAT
	Path           string `env:"radarr_movie_path"`               // /movies/The French Dispatch (2021)
	RelativePath   string `env:"radarr_moviefile_relativepath"`   // The.French.Dispatch.2021.Bluray-720p.mkv
	TMDbID         string `env:"radarr_movie_tmdbid"`             // 542178
	Quality        string `env:"radarr_moviefile_quality"`        // Bluray-720p
	Title          string `env:"radarr_movie_title"`              // The French Dispatch
	FileID         int64  `env:"radarr_moviefile_id"`             // 3531
	Year           int    `env:"radarr_movie_year"`               // 2021
	Size           int64  `env:"radarr_moviefile_size"`           // 3593317970
	ID             int64  `env:"radarr_movie_id"`                 // 2173
	QualityVersion int64  `env:"radarr_moviefile_qualityversion"` // 1
}

// RadarrMovieDelete is the MovieDelete event.
type RadarrMovieDelete struct {
	Title       string `env:"radarr_movie_title"`        // The French Dispatch
	Path        string `env:"radarr_movie_path"`         // /movies/The French Dispatch (2021)
	IMDbID      string `env:"radarr_movie_imdbid"`       // tt8847712
	DeleteFiles string `env:"radarr_movie_deletedfiles"` // XXX: no example. Does this need a split?
	ID          int64  `env:"radarr_movie_id"`           // 2173
	Year        int    `env:"radarr_movie_year"`         // 2021
	TMDbID      int64  `env:"radarr_movie_tmdbid"`       // 542178
	Size        int64  `env:"radarr_movie_folder_size"`  // 3593317970
}

// RadarrRename is the Rename event.
type RadarrRename struct {
	InCinemas             time.Time `env:"radarr_movie_in_cinemas_date"` // 11/22/2005 12:00:00 AM
	ReleaseDate           time.Time `env:"radarr_movie_physical_release_date"`
	Path                  string    `env:"radarr_movie_path"`   // /movies/The French Dispatch (2021)
	IMDbID                string    `env:"radarr_movie_imdbid"` // tt8847712
	FileIDs               []int64   `env:"radarr_moviefile_ids,,"`
	RelativePaths         []string  `env:"radarr_moviefile_relativepaths,|"`
	Paths                 []string  `env:"radarr_moviefile_paths,|"`
	PreviousRelativePaths []string  `env:"radarr_moviefile_previousrelativepaths,|"`
	PreviousPaths         []string  `env:"radarr_moviefile_previouspaths,|"`
	ID                    int64     `env:"radarr_movie_id"`     // 2173
	Year                  int       `env:"radarr_movie_year"`   // 2021
	TMDbID                int64     `env:"radarr_movie_tmdbid"` // 542178
}

// RadarrTest has no members.
type RadarrTest struct{}

// GetRadarrHealthIssue returns the HealthIssue event data.
func (c *CmdEvent) GetRadarrHealthIssue() (output RadarrHealthIssue, err error) {
	return output, c.get(EventHealthIssue, &output)
}

// GetRadarrApplicationUpdate returns the ApplicationUpdate event data.
func (c *CmdEvent) GetRadarrApplicationUpdate() (output RadarrApplicationUpdate, err error) {
	return output, c.get(EventApplicationUpdate, &output)
}

// GetRadarrDownload returns the Download event data.
func (c *CmdEvent) GetRadarrDownload() (output RadarrDownload, err error) {
	return output, c.get(EventDownload, &output)
}

// GetRadarrGrab returns the Grab event data.
func (c *CmdEvent) GetRadarrGrab() (output RadarrGrab, err error) {
	return output, c.get(EventGrab, &output)
}

// GetRadarrMovieFileDelete returns the MovieFileDelete event data.
func (c *CmdEvent) GetRadarrMovieFileDelete() (output RadarrMovieFileDelete, err error) {
	return output, c.get(EventMovieFileDelete, &output)
}

// GetRadarrTest returns the Test event data.
func (c *CmdEvent) GetRadarrTest() (output RadarrTest, err error) {
	return output, c.get(EventTest, &output)
}

// GetRadarrMovieDelete returns the MovieDelete event data.
func (c *CmdEvent) GetRadarrMovieDelete() (output RadarrMovieDelete, err error) {
	return output, c.get(EventMovieDelete, &output)
}

// GetRadarrRename returns the Rename event data.
func (c *CmdEvent) GetRadarrRename() (output RadarrRename, err error) {
	return output, c.get(EventRename, &output)
}
