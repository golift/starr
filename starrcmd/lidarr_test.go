//nolint:paralleltest
package starrcmd_test

import (
	"os"
	"testing"

	"golift.io/starr/starrcmd"
)

func TestLidarrApplicationUpdate(t *testing.T) {
	t.Setenv("lidarr_eventtype", string(starrcmd.EventApplicationUpdate))
	t.Setenv("lidarr_update_previousversion", "5.0.3.5875")
	t.Setenv("lidarr_update_newversion", "5.0.4.5909")
	t.Setenv("lidarr_update_message", "Lidarr updated from 5.0.3.5875 to 5.0.4.5909")

	cmd, err := starrcmd.New()
	if err != nil {
		t.Fatalf("got an unexpected error: %s", err)
	}

	switch info, err := cmd.GetLidarrApplicationUpdate(); {
	case err != nil:
		t.Fatalf("got an unexpected error: %s", err)
	case info.Message != os.Getenv("lidarr_update_message"):
		t.Fatalf("got wrong Message? %s", info.Message)
	case info.NewVersion != "5.0.4.5909":
		t.Fatalf("got wrong new version? wanted: '5.0.4.5909' got: %s", info.Message)
	case info.PreviousVersion != "5.0.3.5875":
		t.Fatalf("got wrong Message? wanted: '5.0.3.5875' got: %s", info.Message)
	}
}

func TestLidarrHealthIssue(t *testing.T) {
	t.Setenv("lidarr_eventtype", string(starrcmd.EventHealthIssue))
	t.Setenv("lidarr_health_issue_type", "SomeIssueTypeForLidarr")
	t.Setenv("lidarr_health_issue_wiki", "https://wiki.servarr.com/lidarr")
	t.Setenv("lidarr_health_issue_level", "Womp")
	t.Setenv("lidarr_health_issue_message", "Lists unavailable due to failures: List name here")

	cmd, err := starrcmd.New()
	if err != nil {
		t.Fatalf("got an unexpected error: %s", err)
	}

	switch info, err := cmd.GetLidarrHealthIssue(); {
	case err != nil:
		t.Fatalf("got an unexpected error: %s", err)
	case info.Message != os.Getenv("lidarr_health_issue_message"):
		t.Fatalf("got wrong Message? %s", info.Message)
	case info.Wiki != "https://wiki.servarr.com/lidarr":
		t.Fatalf("got wrong wiki link? wanted: 'https://wiki.servarr.com/lidarr' got: %s", info.Wiki)
	case info.Level != "Womp":
		t.Fatalf("got wrong level? wanted: 'Womp' got: %s", info.Level)
	case info.IssueType != "SomeIssueTypeForLidarr":
		t.Fatalf("got wrong issue type? wanted: 'ImportListStatusCheck' got: %s", info.IssueType)
	}
}

func TestLidarrHealthIssueEmpty(t *testing.T) {
	t.Setenv("lidarr_eventtype", string(starrcmd.EventHealthIssue))

	cmd, err := starrcmd.New()
	if err != nil {
		t.Fatalf("got an unexpected error: %s", err)
	}

	switch info, err := cmd.GetLidarrHealthIssue(); {
	case err != nil:
		t.Fatalf("got an unexpected error: %s", err)
	case info.IssueType != "":
		t.Fatalf("got wrong issue type? wanted: '' got: %s", info.IssueType)
	}
}

func TestLidarrTest(t *testing.T) {
	t.Setenv("lidarr_eventtype", string(starrcmd.EventTest))

	cmd, err := starrcmd.New()
	if err != nil {
		t.Fatalf("got an unexpected error: %s", err)
	}

	switch info, err := cmd.GetLidarrTest(); {
	case err != nil:
		t.Fatalf("got an unexpected error: %s", err)
	case info != starrcmd.LidarrTest{}:
		t.Fatalf("got an wrong structure in return")
	}
}

func TestLidarrGrab(t *testing.T) {
	t.Setenv("lidarr_eventtype", string(starrcmd.EventGrab))
	t.Setenv("lidarr_download_client", "Deluge")
	t.Setenv("lidarr_release_albumcount", "1")
	t.Setenv("lidarr_release_size", "433061888")
	t.Setenv("lidarr_release_albumreleasedates", "4/21/2010 12:00:00 AM")
	t.Setenv("lidarr_artist_id", "262")
	t.Setenv("lidarr_artist_name", "Tom Petty and the Heartbreakers")
	t.Setenv("lidarr_artist_mbid", "f93dbc64-6f08-4033-bcc7-8a0bb4689849")
	t.Setenv("lidarr_release_indexer", "Indexilate (Prowlarr)")
	t.Setenv("lidarr_release_qualityversion", "1")
	t.Setenv("lidarr_release_quality", "FLAC")
	t.Setenv("lidarr_release_releasegroup", "Someone")
	t.Setenv("lidarr_release_title", "Tom Petty & The Heartbreakers - Mojo (2010) [FLAC (tracks + cue)]")
	t.Setenv("lidarr_release_albummbids", "75f6f410-73e6-485b-898d-6fdaea4c0266")
	t.Setenv("lidarr_download_id", "4A87D9F5F92D82DF4076463E90CC49F27077CB10")
	t.Setenv("lidarr_release_albumtitles", "Mojo")
	t.Setenv("lidarr_artist_type", "Group")

	cmd, err := starrcmd.New()
	if err != nil {
		t.Fatalf("got an unexpected error: %s", err)
	}

	switch info, err := cmd.GetLidarrGrab(); {
	case err != nil:
		t.Fatalf("got an unexpected error: %s", err)
	case info.ArtistName != "Tom Petty and the Heartbreakers":
		t.Fatalf("got wrong artists name? wanted: 'Tom Petty and the Heartbreakers' got: %s", info.ArtistName)
	}
}

func TestLidarrAlbumDownload(t *testing.T) {
	t.Setenv("lidarr_eventtype", string(starrcmd.EventAlbumDownload))
	t.Setenv("lidarr_artist_id", "12345")
	t.Setenv("lidarr_artist_name", "some artist")
	t.Setenv("lidarr_artist_path", "artist.Path)")
	t.Setenv("lidarr_artist_mbid", "artist.Metadata.Value.ForeignArtistId)")
	t.Setenv("lidarr_artist_type", "artist.Metadata.Value.Type)")
	t.Setenv("lidarr_album_id", "5432")
	t.Setenv("lidarr_album_title", "album.Title)")
	t.Setenv("lidarr_album_mbid", "album.ForeignAlbumId)")
	t.Setenv("lidarr_albumrelease_mbid", "release.ForeignReleaseId)")
	t.Setenv("lidarr_album_releasedate", "7/21/2010 1:00:00 AM")
	t.Setenv("lidarr_download_client", "message.DownloadClient ?? string.Empty)")
	t.Setenv("lidarr_download_id", "message.DownloadId ?? string.Empty)")
	t.Setenv("lidarr_addedtrackpaths", "")
	t.Setenv("lidarr_deletedpaths", "")

	cmd, err := starrcmd.New()
	if err != nil {
		t.Fatalf("got an unexpected error: %s", err)
	}

	switch info, err := cmd.GetLidarrAlbumDownload(); {
	case err != nil:
		t.Fatalf("got an unexpected error: %s", err)
	case info.ArtistName != "some artist":
		t.Fatalf("got wrong artists name? wanted: 'some artist' got: %s", info.ArtistName)
	}
}

func TestLidarrRename(t *testing.T) {
	t.Setenv("lidarr_eventtype", string(starrcmd.EventRename))
	t.Setenv("lidarr_artist_id", "34666")
	t.Setenv("lidarr_artist_name", "arti goes here")
	t.Setenv("lidarr_artist_path", "artist.Path)")
	t.Setenv("lidarr_artist_mbid", "artist.Metadata.Value.ForeignArtistId)")
	t.Setenv("lidarr_artist_type", "artist.Metadata.Value.Type)")

	cmd, err := starrcmd.New()
	if err != nil {
		t.Fatalf("got an unexpected error: %s", err)
	}

	switch info, err := cmd.GetLidarrRename(); {
	case err != nil:
		t.Fatalf("got an unexpected error: %s", err)
	case info.ArtistName != "arti goes here":
		t.Fatalf("got wrong artists name? wanted: 'arti goes here' got: %s", info.ArtistName)
	}
}

func TestLidarrTrackRetag(t *testing.T) {
	t.Setenv("lidarr_eventtype", string(starrcmd.EventTrackRetag))
	t.Setenv("lidarr_artist_id", "3473234")
	t.Setenv("lidarr_artist_name", "tushort")
	t.Setenv("lidarr_artist_path", "artist.Path)")
	t.Setenv("lidarr_artist_mbid", "artist.Metadata.Value.ForeignArtistId)")
	t.Setenv("lidarr_artist_type", "artist.Metadata.Value.Type)")
	t.Setenv("lidarr_album_id", "39739479324")
	t.Setenv("lidarr_album_title", "album.Title)")
	t.Setenv("lidarr_album_mbid", "album.ForeignAlbumId)")
	t.Setenv("lidarr_albumrelease_mbid", "release.ForeignReleaseId)")
	t.Setenv("lidarr_album_releasedate", "9/11/2001 1:00:00 AM")
	t.Setenv("lidarr_trackfile_id", "324324")
	t.Setenv("lidarr_trackfile_trackcount", "1")
	t.Setenv("lidarr_trackfile_path", "trackFile.Path)")
	t.Setenv("lidarr_trackfile_tracknumbers", "1,2,3,4")
	t.Setenv("lidarr_trackfile_tracktitles", "title1|title2")
	t.Setenv("lidarr_trackfile_quality", "trackFile.Quality.Quality.Name)")
	t.Setenv("lidarr_trackfile_qualityversion", "1")
	t.Setenv("lidarr_trackfile_releasegroup", "trackFile.ReleaseGroup ?? string.Empty)")
	t.Setenv("lidarr_trackfile_scenename", "trackFile.SceneName ?? string.Empty)")
	t.Setenv("lidarr_tags_diff", "message.Diff.ToJson())")
	t.Setenv("lidarr_tags_scrubbed", "False")

	cmd, err := starrcmd.New()
	if err != nil {
		t.Fatalf("got an unexpected error: %s", err)
	}

	switch info, err := cmd.GetLidarrTrackRetag(); {
	case err != nil:
		t.Fatalf("got an unexpected error: %s", err)
	case info.ArtistName != "tushort":
		t.Fatalf("got wrong artists name? wanted: 'tushort' got: %s", info.ArtistName)
	}
}
