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

func TestSonarrGrab(t *testing.T) {
	t.Setenv("sonarr_eventtype", string(starrcmd.EventGrab))
	t.Setenv("sonarr_release_quality", "HDTV-720p")
	t.Setenv("sonarr_series_title", "This Is Us")
	t.Setenv("sonarr_release_qualityversion", "1")
	t.Setenv("sonarr_series_id", "47")
	t.Setenv("sonarr_release_episodenumbers", "4")
	t.Setenv("sonarr_release_episodecount", "1")
	t.Setenv("sonarr_download_client", "NZBGet")
	t.Setenv("sonarr_release_episodeairdates", "2022-01-25")
	t.Setenv("sonarr_release_episodetitles", "Don't Let Me Keep You")
	t.Setenv("sonarr_release_title", "This.is.Us.S06E04.720p.HDTV.x264-SYNCOPY")
	t.Setenv("sonarr_download_id", "a87bda3c0e7f40a1b8fa011b421a5201")
	t.Setenv("sonarr_release_indexer", "Indexor (Prowlarr)")
	t.Setenv("sonarr_series_type", "Standard")
	t.Setenv("sonarr_release_size", "885369406")
	t.Setenv("sonarr_series_tvdbid", "311714")
	t.Setenv("sonarr_series_tvmazeid", "17128")
	t.Setenv("sonarr_release_releasegroup", "SYNCOPY")
	t.Setenv("sonarr_release_seasonnumber", "6")
	t.Setenv("sonarr_release_absoluteepisodenumbers", "92")
	t.Setenv("sonarr_series_imdbid", "tt5555260")
	t.Setenv("sonarr_release_episodeairdatesutc", "1/26/2022 2:00:00 AM")

	cmd, err := starrcmd.New()
	if err != nil {
		t.Fatalf("got an unexpected error: %s", err)
	}

	switch info, err := cmd.GetSonarrGrab(); {
	default:
		// fmt.Println(info.EpisodeAirDatesUTC)
	case err != nil:
		t.Fatalf("got an unexpected error: %s", err)
	case info.DownloadClient != "NZBGet":
		t.Fatalf("got wrong download client? expected: <blank>, got: %v", info.DownloadClient)
	}
}

func TestSonarrRename(t *testing.T) {
	t.Setenv("sonarr_eventtype", string(starrcmd.EventRename))

	t.Setenv("sonarr_series_id", "12345")
	t.Setenv("sonarr_series_title", "series.Title")
	t.Setenv("sonarr_series_path", "series.Path")
	t.Setenv("sonarr_series_tvdbid", "324")
	t.Setenv("sonarr_series_tvmazeid", "23455")
	t.Setenv("sonarr_series_imdbid", "iasdashdgaisdhaidhadji")
	t.Setenv("sonarr_series_type", "anime")
	t.Setenv("sonarr_episodefile_ids", "1,2,3,4")
	t.Setenv("sonarr_episodefile_relativepaths", "foo")
	t.Setenv("sonarr_episodefile_paths", "stuff")
	t.Setenv("sonarr_episodefile_previousrelativepaths", "more stuff")
	t.Setenv("sonarr_episodefile_previouspaths", "all the stuff")

	cmd, err := starrcmd.New()
	if err != nil {
		t.Fatalf("got an unexpected error: %s", err)
	}

	switch info, err := cmd.GetSonarrRename(); {
	default:
		// fmt.Println(info.EpisodeAirDatesUTC)
	case err != nil:
		t.Fatalf("got an unexpected error: %s", err)
	case info.ID != 12345:
		t.Fatalf("got wrong ID? expected: 12345, got: %v", info.ID)
	}
}

// SonarrSeriesDelete is the SeriesDelete event.
func TestSonarrSeriesDelete(t *testing.T) {
	t.Setenv("sonarr_eventtype", string(starrcmd.EventSeriesDelete))

	t.Setenv("sonarr_series_id", "2323")
	t.Setenv("sonarr_series_title", "series.Title")
	t.Setenv("sonarr_series_path", "series.Path")
	t.Setenv("sonarr_series_tvdbid", "232323")
	t.Setenv("sonarr_series_tvmazeid", "343434")
	t.Setenv("sonarr_series_imdbid", "series.ImdbId ?? string.Empty)")
	t.Setenv("sonarr_series_type", "animar")
	t.Setenv("sonarr_series_deletedfiles", "false")

	cmd, err := starrcmd.New()
	if err != nil {
		t.Fatalf("got an unexpected error: %s", err)
	}

	switch info, err := cmd.GetSonarrSeriesDelete(); {
	default:
		// fmt.Println(info.EpisodeAirDatesUTC)
	case err != nil:
		t.Fatalf("got an unexpected error: %s", err)
	case info.ID != 2323:
		t.Fatalf("got wrong ID? expected: 2323, got: %v", info.ID)
	}
}

// SonarrEpisodeFileDelete is the EpisodeFileDelete event.
func TestSonarrEpisodeFileDelete(t *testing.T) {
	t.Setenv("sonarr_eventtype", string(starrcmd.EventEpisodeFileDelete))

	t.Setenv("sonarr_episodefile_deletereason", "deleteMessage.Reason.ToString())")
	t.Setenv("sonarr_series_id", "31212")
	t.Setenv("sonarr_series_title", "series.Title")
	t.Setenv("sonarr_series_path", "series.Path")
	t.Setenv("sonarr_series_tvdbid", "1234444")
	t.Setenv("sonarr_series_tvmazeid", "44332")
	t.Setenv("sonarr_series_imdbid", "series.ImdbId ?? string.Empty)")
	t.Setenv("sonarr_series_type", "anime")
	t.Setenv("sonarr_episodefile_id", "12121222")
	t.Setenv("sonarr_episodefile_episodecount", "2")
	t.Setenv("sonarr_episodefile_relativepath", "episodeFile.RelativePath")
	t.Setenv("sonarr_episodefile_path", "Path.Combine(series.Path, episodeFile.RelativePath))")
	t.Setenv("sonarr_episodefile_episodeids", "1,2,3,4")
	t.Setenv("sonarr_episodefile_seasonnumber", "episodeFile.SeasonNumber.ToString())")
	t.Setenv("sonarr_episodefile_episodenumbers", "1,2,3,4")
	t.Setenv("sonarr_episodefile_episodeairdates", "")
	t.Setenv("sonarr_episodefile_episodeairdatesutc", "")
	t.Setenv("sonarr_episodefile_episodetitles", "title1|title2")
	t.Setenv("sonarr_episodefile_quality", "")
	t.Setenv("sonarr_episodefile_qualityversion", "no idea")
	t.Setenv("sonarr_episodefile_releasegroup", "DOPE")
	t.Setenv("sonarr_episodefile_scenename", "dirty harry")

	cmd, err := starrcmd.New()
	if err != nil {
		t.Fatalf("got an unexpected error: %s", err)
	}

	switch info, err := cmd.GetSonarrEpisodeFileDelete(); {
	default:
		// fmt.Println(info.EpisodeAirDatesUTC)
	case err != nil:
		t.Fatalf("got an unexpected error: %s", err)
	case info.ID != 31212:
		t.Fatalf("got wrong ID? expected: 12345, got: %v", info.ID)
	}
}
