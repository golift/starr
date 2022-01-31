//nolint:paralleltest
package starrcmd_test

import (
	"os"
	"testing"

	"golift.io/starr/starrcmd"
)

func TestSonarrApplicationUpdate(t *testing.T) {
	t.Setenv("sonarr_eventtype", string(starrcmd.EventApplicationUpdate))
	t.Setenv("sonarr_update_previousversion", "2.0.3.5875")
	t.Setenv("sonarr_update_newversion", "2.0.4.5909")
	t.Setenv("sonarr_update_message", "Sonarr updated from 2.0.3.5875 to 2.0.4.5909")

	cmd, err := starrcmd.New()
	if err != nil {
		t.Fatalf("got an unexpected error: %s", err)
	}

	switch info, err := cmd.GetSonarrApplicationUpdate(); {
	case err != nil:
		t.Fatalf("got an unexpected error: %s", err)
	case info.Message != os.Getenv("sonarr_update_message"):
		t.Fatalf("got wrong Message? %s", info.Message)
	case info.NewVersion != "2.0.4.5909":
		t.Fatalf("got wrong new version? wanted: '2.0.4.5909' got: %s", info.Message)
	case info.PreviousVersion != "2.0.3.5875":
		t.Fatalf("got wrong Message? wanted: '2.0.3.5875' got: %s", info.Message)
	}
}

func TestSonarrHealthIssue(t *testing.T) {
	t.Setenv("sonarr_eventtype", string(starrcmd.EventHealthIssue))
	t.Setenv("sonarr_health_issue_type", "SomeIssueTypeForSonarr")
	t.Setenv("sonarr_health_issue_wiki", "https://wiki.servarr.com/sonarr")
	t.Setenv("sonarr_health_issue_level", "Error")
	t.Setenv("sonarr_health_issue_message", "Lists unavailable due to failures: List name here")

	cmd, err := starrcmd.New()
	if err != nil {
		t.Fatalf("got an unexpected error: %s", err)
	}

	switch info, err := cmd.GetSonarrHealthIssue(); {
	case err != nil:
		t.Fatalf("got an unexpected error: %s", err)
	case info.Message != os.Getenv("sonarr_health_issue_message"):
		t.Fatalf("got wrong Message? %s", info.Message)
	case info.Wiki != "https://wiki.servarr.com/sonarr":
		t.Fatalf("got wrong wiki link? wanted: 'https://wiki.servarr.com/sonarr' got: %s", info.Wiki)
	case info.Level != "Error":
		t.Fatalf("got wrong level? wanted: 'Error' got: %s", info.Level)
	case info.IssueType != "SomeIssueTypeForSonarr":
		t.Fatalf("got wrong issue type? wanted: 'ImportListStatusCheck' got: %s", info.IssueType)
	}
}

func TestSonarrTest(t *testing.T) {
	t.Setenv("sonarr_eventtype", string(starrcmd.EventTest))

	cmd, err := starrcmd.New()
	if err != nil {
		t.Fatalf("got an unexpected error: %s", err)
	}

	switch info, err := cmd.GetSonarrTest(); {
	case err != nil:
		t.Fatalf("got an unexpected error: %s", err)
	case info != starrcmd.SonarrTest{}:
		t.Fatalf("got an wrong structure in return")
	}
}

// XXX: this test could use a bit more love.
func TestSonarrDownload(t *testing.T) {
	// Only testing a few members here. Expand this if you need more tests!
	t.Setenv("sonarr_eventtype", string(starrcmd.EventDownload))
	t.Setenv("sonarr_series_title", "Le Title")
	t.Setenv("sonarr_series_id", "1234")
	t.Setenv("sonarr_isupgrade", "True")
	t.Setenv("sonarr_episodefile_episodeairdatesutc", "1/21/2022 2:00:00 PM,1/21/2022 2:12:00 PM")
	t.Setenv("sonarr_deletedpaths", "/path1|/path2")
	t.Setenv("sonarr_episodefile_episodeairdates", "2022-01-21,2022-01-21")
	t.Setenv("sonarr_episodefile_episodetitles", "Title 1|Title 2")

	cmd, err := starrcmd.New()
	if err != nil {
		t.Fatalf("got an unexpected error: %s", err)
	}

	switch info, err := cmd.GetSonarrDownload(); {
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
