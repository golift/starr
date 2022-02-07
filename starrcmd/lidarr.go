//nolint:lll
package starrcmd

/*
All 7 Lidarr events are accounted for; 1/30/2022.
https://github.com/Lidarr/Lidarr/blob/develop/src/NzbDrone.Core/Notifications/CustomScript/CustomScript.cs
*/

import (
	"time"
)

// LidarrApplicationUpdate is the ApplicationUpdate event.
type LidarrApplicationUpdate struct {
	PreviousVersion string `env:"lidarr_update_previousversion"` // 4.0.3.5875
	NewVersion      string `env:"lidarr_update_newversion"`      // 4.0.4.5909
	Message         string `env:"lidarr_update_message"`         // Lidarr updated from 4.0.3.5875 to 4.0.4.5909
}

// LidarrHealthIssue is the HealthIssue event.
type LidarrHealthIssue struct {
	Message   string `env:"lidarr_health_issue_message"` // Lists unavailable due to failures: List name here
	IssueType string `env:"lidarr_health_issue_type"`    // ImportListStatusCheck
	Wiki      string `env:"lidarr_health_issue_wiki"`    // https://wiki.servarr.com/lidarr/
	Level     string `env:"lidarr_health_issue_level"`   // Warning
}

// LidarrGrab is the Grab event.
type LidarrGrab struct {
	DownloadClient string      `env:"lidarr_download_client"`             // Deluge
	AlbumCount     int         `env:"lidarr_release_albumcount"`          // 1
	Size           int64       `env:"lidarr_release_size"`                // 433061888
	ReleaseDates   []time.Time `env:"lidarr_release_albumreleasedates,,"` // 4/21/2010 12:00:00 AM
	ArtistID       int64       `env:"lidarr_artist_id"`                   // 262
	ArtistName     string      `env:"lidarr_artist_name"`                 // Tom Petty and the Heartbreakers
	MBID           string      `env:"lidarr_artist_mbid"`                 // f93dbc64-6f08-4033-bcc7-8a0bb4689849
	Indexer        string      `env:"lidarr_release_indexer"`             // Indexilate (Prowlarr)
	QualityVerson  int64       `env:"lidarr_release_qualityversion"`      // 1
	Quality        string      `env:"lidarr_release_quality"`             // FLAC
	ReleaseGroup   string      `env:"lidarr_release_releasegroup"`        //
	ReleaseTitle   string      `env:"lidarr_release_title"`               // Tom Petty & The Heartbreakers - Mojo (2010) [FLAC (tracks + cue)]
	AlbumMBIDs     []string    `env:"lidarr_release_albummbids,|"`        // 75f6f410-73e6-485b-898d-6fdaea4c0266
	DownloadID     string      `env:"lidarr_download_id"`                 // 4A87D9F5F92D82DF4076463E90CC49F27077CB10
	Titles         []string    `env:"lidarr_release_albumtitles,|"`       // Mojo
	ArtistType     string      `env:"lidarr_artist_type"`                 // Group
}

// LidarrAlbumDownload is the AlbumDownload event.
type LidarrAlbumDownload struct {
	ArtistID         int64     `env:"lidarr_artist_id"`         // artist.Id.ToString())
	ArtistName       string    `env:"lidarr_artist_name"`       // artist.Metadata.Value.Name)
	Path             string    `env:"lidarr_artist_path"`       // artist.Path)
	ArtistMBID       string    `env:"lidarr_artist_mbid"`       // artist.Metadata.Value.ForeignArtistId)
	ArtistType       string    `env:"lidarr_artist_type"`       // artist.Metadata.Value.Type)
	AlbumID          int64     `env:"lidarr_album_id"`          // album.Id.ToString())
	Title            string    `env:"lidarr_album_title"`       // album.Title)
	MBID             string    `env:"lidarr_album_mbid"`        // album.ForeignAlbumId)
	AlbumReleaseMBID string    `env:"lidarr_albumrelease_mbid"` // release.ForeignReleaseId)
	ReleaseDate      time.Time `env:"lidarr_album_releasedate"` // album.ReleaseDate.ToString())
	DownloadClient   string    `env:"lidarr_download_client"`   // message.DownloadClient ?? string.Empty)
	DownloadID       string    `env:"lidarr_download_id"`       // message.DownloadId ?? string.Empty)
	AddedTrackPaths  []string  `env:"lidarr_addedtrackpaths,|"` // string.Join("|", message.TrackFiles.Select(e => e.Path)))
	DeletedPaths     []string  `env:"lidarr_deletedpaths,|"`    // string.Join("|", message.OldFiles.Select(e => e.Path)))

}

// LidarrRename is the Rename event.
type LidarrRename struct {
	ArtistID   int64  `env:"lidarr_artist_id"`   // artist.Id.ToString())
	ArtistName string `env:"lidarr_artist_name"` // artist.Metadata.Value.Name)
	Path       string `env:"lidarr_artist_path"` // artist.Path)
	ArtistMBID string `env:"lidarr_artist_mbid"` // artist.Metadata.Value.ForeignArtistId)
	ArtistType string `env:"lidarr_artist_type"` // artist.Metadata.Value.Type)
}

// LidarrTrackRetag is the TrackRetag event.
type LidarrTrackRetag struct {
	ArtistID         int64     `env:"lidarr_artist_id"`                // artist.Id.ToString())
	ArtistName       string    `env:"lidarr_artist_name"`              // artist.Metadata.Value.Name)
	Path             string    `env:"lidarr_artist_path"`              // artist.Path)
	ArtistMBID       string    `env:"lidarr_artist_mbid"`              // artist.Metadata.Value.ForeignArtistId)
	ArtistType       string    `env:"lidarr_artist_type"`              // artist.Metadata.Value.Type)
	ID               int64     `env:"lidarr_album_id"`                 // album.Id.ToString())
	Title            string    `env:"lidarr_album_title"`              // album.Title)
	MBID             string    `env:"lidarr_album_mbid"`               // album.ForeignAlbumId)
	AlbumReleaseMBID string    `env:"lidarr_albumrelease_mbid"`        // release.ForeignReleaseId)
	ReleaseDate      time.Time `env:"lidarr_album_releasedate"`        // album.ReleaseDate.ToString())
	FileID           int64     `env:"lidarr_trackfile_id"`             // trackFile.Id.ToString())
	TrackCount       string    `env:"lidarr_trackfile_trackcount"`     // trackFile.Tracks.Value.Count.ToString())
	FilePath         string    `env:"lidarr_trackfile_path"`           // trackFile.Path)
	TrackNumbers     []int     `env:"lidarr_trackfile_tracknumbers,,"` // string.Join(",", trackFile.Tracks.Value.Select(e => e.TrackNumber)))
	TrackTitles      []string  `env:"lidarr_trackfile_tracktitles,|"`  // string.Join("|", trackFile.Tracks.Value.Select(e => e.Title)))
	Quality          string    `env:"lidarr_trackfile_quality"`        // trackFile.Quality.Quality.Name)
	QualityVersion   int64     `env:"lidarr_trackfile_qualityversion"` // trackFile.Quality.Revision.Version.ToString())
	ReleaseGroup     string    `env:"lidarr_trackfile_releasegroup"`   // trackFile.ReleaseGroup ?? string.Empty)
	SceneName        string    `env:"lidarr_trackfile_scenename"`      // trackFile.SceneName ?? string.Empty)
	TagsDiff         string    `env:"lidarr_tags_diff"`                // message.Diff.ToJson())
	TagsScrubbed     bool      `env:"lidarr_tags_scrubbed"`            // message.Scrubbed.ToString())
}

// LidarrTest has no members.
type LidarrTest struct{}

// GetLidarrApplicationUpdate returns the ApplicationUpdate event data.
func (c *CmdEvent) GetLidarrApplicationUpdate() (output LidarrApplicationUpdate, err error) {
	return output, c.get(EventApplicationUpdate, &output)
}

// GetLidarrHealthIssue returns the ApplicationUpdate event data.
func (c *CmdEvent) GetLidarrHealthIssue() (output LidarrHealthIssue, err error) {
	return output, c.get(EventHealthIssue, &output)
}

// GetLidarrGrab returns the Grab event data.
func (c *CmdEvent) GetLidarrGrab() (output LidarrGrab, err error) {
	return output, c.get(EventGrab, &output)
}

// GetLidarrAlbumDownload returns the AlbumDownload event data.
func (c *CmdEvent) GetLidarrAlbumDownload() (output LidarrAlbumDownload, err error) {
	return output, c.get(EventAlbumDownload, &output)
}

// GetLidarrRename returns the Rename event data.
func (c *CmdEvent) GetLidarrRename() (output LidarrRename, err error) {
	return output, c.get(EventRename, &output)
}

// GetLidarrTrackRetag returns the TrackRetag event data.
func (c *CmdEvent) GetLidarrTrackRetag() (output LidarrTrackRetag, err error) {
	return output, c.get(EventTrackRetag, &output)
}

// GetLidarrTest returns the ApplicationUpdate event data.
func (c *CmdEvent) GetLidarrTest() (output LidarrTest, err error) {
	return output, c.get(EventTest, &output)
}
