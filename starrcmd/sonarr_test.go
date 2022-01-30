//nolint:paralleltest
package starrcmd_test

import (
	"os"
	"testing"

	"golift.io/starr/starrcmd"
)

func TestSonarrHealthIssue(t *testing.T) {
	starrcmd.EventType = starrcmd.EventHealthIssue

	t.Setenv("sonarr_eventtype", string(starrcmd.EventHealthIssue))
	t.Setenv("sonarr_health_issue_type", "ImportListStatusCheck")
	t.Setenv("sonarr_health_issue_wiki", "https://wiki.servarr.com/")
	t.Setenv("sonarr_health_issue_level", "Warning")
	t.Setenv("sonarr_health_issue_message", "Lists unavailable due to failures:Listnamehere")

	switch info, err := starrcmd.GetSonarrHealthIssue(); {
	case err != nil:
		t.Fatalf("got an unexpected error: %s", err)
	case info.Message != os.Getenv("sonarr_health_issue_message"):
		t.Fatalf("got wrong Message? %s", info.Message)
	case info.Wiki != "https://wiki.servarr.com/":
		t.Fatalf("got wrong wiki link? wanted: 'https://wiki.servarr.com/' got: %s", info.Wiki)
	case info.Level != "Warning":
		t.Fatalf("got wrong level? wanted: 'Warning' got: %s", info.Level)
	case info.IssueType != "ImportListStatusCheck":
		t.Fatalf("got wrong issue type? wanted: 'ImportListStatusCheck' got: %s", info.IssueType)
	}
}

// XXX: this test could use a bit more love.
func TestSonarrDownload(t *testing.T) {
	starrcmd.EventType = starrcmd.EventDownload

	// Only testing a few members here. Expand this if you need more tests!
	t.Setenv("sonarr_eventtype", string(starrcmd.EventDownload))
	t.Setenv("sonarr_series_title", "Le Title")
	t.Setenv("sonarr_series_id", "1234")
	t.Setenv("sonarr_isupgrade", "True")
	t.Setenv("sonarr_episodefile_episodeairdatesutc", "1/21/2022 2:00:00 PM,1/21/2022 2:12:00 PM")
	t.Setenv("sonarr_deletedpaths", "/path1|/path2")
	t.Setenv("sonarr_episodefile_episodeairdates", "2022-01-21,2022-01-21")
	t.Setenv("sonarr_episodefile_episodetitles", "Title 1|Title 2")

	switch info, err := starrcmd.GetSonarrDownload(); {
	default:
		// fmt.Println(info.EpisodeAirDatesUTC)
	case err != nil:
		t.Fatalf("got an unexpected error: %s", err)
	case info.DownloadClient != "":
		t.Fatalf("got wrong download client? expected: <blank>, got: %v", info.DownloadClient)
	case !info.IsUpgrade:
		t.Fatalf("got wrong upgrade bool? expected: true, got: %v", info.IsUpgrade)
	}
}
