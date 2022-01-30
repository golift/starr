//nolint:lll
package starrcmd

/*
INCOMPLETE.
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

// LidarrTest has no members.
type LidarrTest struct{}

// GetLidarrApplicationUpdate returns the ApplicationUpdate event data.
func GetLidarrApplicationUpdate() (output LidarrApplicationUpdate, err error) {
	return output, get(EventApplicationUpdate, &output)
}

// GetLidarrHealthIssue returns the ApplicationUpdate event data.
func GetLidarrHealthIssue() (output LidarrHealthIssue, err error) {
	return output, get(EventHealthIssue, &output)
}

// GetLidarrGrab returns the Grab event data.
func GetLidarrGrab() (output LidarrGrab, err error) {
	return output, get(EventGrab, &output)
}

// GetLidarrTest returns the ApplicationUpdate event data.
func GetLidarrTest() (output LidarrTest, err error) {
	return output, get(EventTest, &output)
}
