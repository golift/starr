//nolint:paralleltest,cyclop,funlen
package starrcmd_test

import (
	"os"
	"testing"

	"golift.io/starr/starrcmd"
)

func TestRadarrApplicationUpdate(t *testing.T) {
	t.Setenv("radarr_eventtype", string(starrcmd.EventApplicationUpdate))
	t.Setenv("radarr_update_previousversion", "4.0.3.5875")
	t.Setenv("radarr_update_newversion", "4.0.4.5909")
	t.Setenv("radarr_update_message", "Radarr updated from 4.0.3.5875 to 4.0.4.5909")

	cmd, err := starrcmd.New()
	if err != nil {
		t.Fatalf("got an unexpected error: %s", err)
	}

	switch info, err := cmd.GetRadarrApplicationUpdate(); {
	case err != nil:
		t.Fatalf("got an unexpected error: %s", err)
	case info.Message != os.Getenv("radarr_update_message"):
		t.Fatalf("got wrong Message? %s", info.Message)
	case info.NewVersion != "4.0.4.5909":
		t.Fatalf("got wrong new version? wanted: '4.0.4.5909' got: %s", info.Message)
	case info.PreviousVersion != "4.0.3.5875":
		t.Fatalf("got wrong Message? wanted: '4.0.3.5875' got: %s", info.Message)
	}
}

func TestRadarrHealthIssue(t *testing.T) {
	t.Setenv("radarr_eventtype", string(starrcmd.EventHealthIssue))
	t.Setenv("radarr_health_issue_type", "ImportListStatusCheck")
	t.Setenv("radarr_health_issue_wiki", "https://wiki.servarr.com/")
	t.Setenv("radarr_health_issue_level", "Warning")
	t.Setenv("radarr_health_issue_message", "Lists unavailable due to failures: List name here")

	cmd, err := starrcmd.New()
	if err != nil {
		t.Fatalf("got an unexpected error: %s", err)
	}

	switch info, err := cmd.GetRadarrHealthIssue(); {
	case err != nil:
		t.Fatalf("got an unexpected error: %s", err)
	case info.Message != os.Getenv("radarr_health_issue_message"):
		t.Fatalf("got wrong Message? %s", info.Message)
	case info.Wiki != "https://wiki.servarr.com/":
		t.Fatalf("got wrong wiki link? wanted: 'https://wiki.servarr.com/' got: %s", info.Wiki)
	case info.Level != "Warning":
		t.Fatalf("got wrong level? wanted: 'Warning' got: %s", info.Level)
	case info.IssueType != "ImportListStatusCheck":
		t.Fatalf("got wrong issue type? wanted: 'ImportListStatusCheck' got: %s", info.IssueType)
	}
}

func TestRadarrTest(t *testing.T) {
	t.Setenv("radarr_eventtype", string(starrcmd.EventTest))

	cmd, err := starrcmd.New()
	if err != nil {
		t.Fatalf("got an unexpected error: %s", err)
	}

	switch info, err := cmd.GetRadarrTest(); {
	case err != nil:
		t.Fatalf("got an unexpected error: %s", err)
	case info != starrcmd.RadarrTest{}:
		t.Fatalf("got an wrong structure in return")
	}
}

func TestRadarrGrab(t *testing.T) {
	// XXX: This isn't everything, should add the rest.
	// Also write another test that purposely doesn't add everything to catch a different regression.
	t.Setenv("radarr_eventtype", string(starrcmd.EventGrab))
	t.Setenv("radarr_release_qualityversion", "1")
	t.Setenv("radarr_movie_physical_release_date", "1/19/2006 12:00:00 AM")
	t.Setenv("radarr_release_releasegroup", "SLOWPOKE")
	t.Setenv("radarr_movie_id", "1234544")
	t.Setenv("radarr_indexerflags", "3")
	t.Setenv("radarr_movie_imdbid", "tt044817")
	t.Setenv("radarr_download_id", "E63FAFFAAA0DEE42F0846348A9C0657BC53E7AA5")
	t.Setenv("radarr_release_title", "Some damn movie")
	t.Setenv("radarr_movie_in_cinemas_date", "11/22/2005 12:00:00 AM")
	t.Setenv("radarr_movie_year", "2012")
	t.Setenv("radarr_release_indexer", "Indexinator (Prowlarr)")
	t.Setenv("radarr_movie_title", "XxX")
	t.Setenv("radarr_release_size", "123456778")
	t.Setenv("radarr_download_client", "Qbot")

	cmd, err := starrcmd.New()
	if err != nil {
		t.Fatalf("got an unexpected error: %s", err)
	}

	switch info, err := cmd.GetRadarrGrab(); {
	case err != nil:
		t.Fatalf("got an unexpected error: %s", err)
	case info.QualityVersion != int64(1):
		t.Fatalf("got wrong quality version? wanted: 1, got: %d", info.QualityVersion)
	case info.ID != int64(1234544):
		t.Fatalf("got wrong id? wanted: 1234544, got: %v", info.ID)
	case info.Year != 2012:
		t.Fatalf("got wrong year? wanted 2012: got: %v", info.Year)
	case info.ReleaseGroup != "SLOWPOKE":
		t.Fatalf("got wrong release group? wanted: SLOWPOKE, got: %v", info.ReleaseGroup)
	case info.IMDbID != "tt044817":
		t.Fatalf("got wrong imdb id? wanted: tt044817, got: %v", info.IMDbID)
	case info.IndexerFlags != 3:
		t.Fatalf("got wrong indexer flags? wanted: 3, got: %v", info.IndexerFlags)
		/*  case info.ReleaseDate != ???:
		      t.Fatalf("got wrong release date? wanted: got: %v", info.Year)
		    case info.InCinemas != ???:
		      t.Fatalf("got wrong cinema date? wanted: got: %v", info.Year) */
	case info.Title != "XxX":
		t.Fatalf("got wrong title? wanted: XxX, got: %v", info.Title)
	case info.DownloadID != "E63FAFFAAA0DEE42F0846348A9C0657BC53E7AA5":
		t.Fatalf("got wrong download id? wanted: E63FAFFAAA0DEE42F0846348A9C0657BC53E7AA5, got: %v", info.DownloadID)
	case info.DownloadClient != "Qbot":
		t.Fatalf("got wrong download client? wanted: Qbot, got: %v", info.DownloadClient)
	case info.ReleaseIndexer != "Indexinator (Prowlarr)":
		t.Fatalf("got wrong release indexer? wanted: Indexinator (Prowlarr), got: %v", info.ReleaseIndexer)
	case info.Size != 123456778:
		t.Fatalf("got wrong release size? wanted: 123456778, got: %v", info.Size)
	}
}

func TestRadarrRename(t *testing.T) {
	// This isn't everything, but it's most..
	t.Setenv("radarr_eventtype", string(starrcmd.EventRename))
	t.Setenv("radarr_movie_id", "123456")
	t.Setenv("radarr_movie_year", "2099")
	t.Setenv("radarr_movie_path", "/gohome")
	t.Setenv("radarr_movie_imdbid", "tt4444")
	t.Setenv("radarr_movie_tmdbid", "23123123")
	t.Setenv("radarr_movie_in_cinemas_date", "12/2/2025 01:21:24 AM")
	t.Setenv("radarr_movie_physical_release_date", "4/20/2023 04:21:54 PM")
	t.Setenv("radarr_moviefile_ids", "3,4,5,6,7,8")
	t.Setenv("radarr_moviefile_relativepaths", "/here|/there|/every/where")
	t.Setenv("radarr_moviefile_paths", "/movie/path")
	t.Setenv("radarr_moviefile_previousrelativepaths", "/none")
	t.Setenv("radarr_moviefile_previouspaths", "/really|/none")

	cmd, err := starrcmd.New()
	if err != nil {
		t.Fatalf("got an unexpected error: %s", err)
	}

	switch info, err := cmd.GetRadarrRename(); {
	case err != nil:
		t.Fatalf("got an unexpected error: %s", err)
	case info.ID != int64(123456):
		t.Fatalf("got wrong id? wanted: 12345, got: %v", info.ID)
	case info.Year != 2099:
		t.Fatalf("got wrong year? wanted 2013: got: %v", info.Year)
	case info.Path != "/gohome":
		t.Fatalf("got wrong path? wanted: /here, got: %v", info.Path)
	case info.IMDbID != "tt4444":
		t.Fatalf("got wrong imdb id? wanted: tt4444, got: %v", info.IMDbID)
	case info.TMDbID != 23123123:
		t.Fatalf("got wrong tmdb id? wanted: 23123123, got: %v", info.TMDbID)
		/*  case info.ReleaseDate != ???:
		      t.Fatalf("got wrong release date? wanted: got: %v", info.Year)
		    case info.InCinemas != ???:
		      t.Fatalf("got wrong cinema date? wanted: got: %v", info.Year) */
	case len(info.FileIDs) != 6 || info.FileIDs[0] != 3 || info.FileIDs[5] != 8:
		t.Fatalf("got wrong files ids? wanted: 3,4,5,6,7,8, got: %v", info.FileIDs)
	case len(info.RelativePaths) != 3 || info.RelativePaths[0] != "/here":
		t.Fatalf("got wrong relative paths?  got: %v", info.RelativePaths)
	case len(info.PreviousRelativePaths) != 1 || info.PreviousRelativePaths[0] != "/none":
		t.Fatalf("got wrong previous relative pats?  got: %v", info.PreviousRelativePaths)
	case len(info.PreviousPaths) != 2 || info.PreviousPaths[1] != "/none":
		t.Fatalf("got wrong previous paths? got: %v", info.PreviousPaths)
	}
}

func TestRadarrDownload(t *testing.T) {
	t.Setenv("radarr_eventtype", string(starrcmd.EventDownload))
	t.Setenv("radarr_movie_physical_release_date", "")
	t.Setenv("radarr_moviefile_path", "/movies/Just Go with It (2011)/Just.Go.with.It.2011.Bluray-1080p.mkv")
	t.Setenv("radarr_movie_imdbid", "tt1564367")
	t.Setenv("radarr_moviefile_scenename", "Just.Go.with.It.2011.1080p.BluRay.x264-OFT")
	t.Setenv("radarr_moviefile_id", "3594")
	t.Setenv("radarr_moviefile_releasegroup", "OFT")
	t.Setenv("radarr_download_id", "string F3D870942BFDD643488852284E917336170CEA00")
	t.Setenv("radarr_movie_in_cinemas_date", "2/10/2011 12:00:00 AM")
	t.Setenv("radarr_moviefile_sourcefolder", "/downloads/Seeding/Just.Go.with.It.2011.1080p.BluRay.x264-OFT")
	t.Setenv("radarr_movie_year", "2011")
	t.Setenv("radarr_isupgrade", "False")
	t.Setenv("radarr_movie_path", "/movies/Just Go with It (2011)")
	t.Setenv("radarr_moviefile_relativepath", "Just.Go.with.It.2011.Bluray-1080p.mkv")
	t.Setenv("radarr_download_client", "Deluge")
	t.Setenv("radarr_moviefile_sourcepath", "/downloads/Seeding/Just.Go.with.It/Just.Go.with.It.2011.x264-OFT.mkv")
	t.Setenv("radarr_movie_tmdbid", "50546")
	t.Setenv("radarr_movie_id", "924")
	t.Setenv("radarr_moviefile_quality", "Bluray-1080p")
	t.Setenv("radarr_movie_title", "Just Go with It")
	t.Setenv("radarr_moviefile_qualityversion", "1")
	t.Setenv("radarr_deletedrelativepaths", "1")
	t.Setenv("radarr_deletedpaths", "1")

	cmd, err := starrcmd.New()
	if err != nil {
		t.Fatalf("got an unexpected error: %s", err)
	}

	switch info, err := cmd.GetRadarrDownload(); {
	case err != nil:
		t.Fatalf("got an unexpected error: %s", err)
	case info.IMDbID != "tt1564367":
		t.Fatalf("got an wrong IMDBID? wanted: 'tt1564367', got: %v", info.IMDbID)
	}
}

func TestRadarrMovieFileDelete(t *testing.T) {
	t.Setenv("radarr_eventtype", string(starrcmd.EventMovieFileDelete))

	t.Setenv("radarr_moviefile_deletereason", "Upgrade")
	t.Setenv("radarr_moviefile_path", "/movies/The French Dispatch (2021)/The.French.Dispatch.2021.Bluray-720p.mkv")
	t.Setenv("radarr_moviefile_scenename", "The.French.Dispatch.2021.720p.BluRay.x264-WoAT")
	t.Setenv("radarr_movie_imdbid", "tt1564368")
	t.Setenv("radarr_moviefile_id", "3531")
	t.Setenv("radarr_moviefile_releasegroup", "WoAT")
	t.Setenv("radarr_movie_year", "2021")
	t.Setenv("radarr_movie_path", "/movies/The French Dispatch (2021)")
	t.Setenv("radarr_moviefile_relativepath", "The.French.Dispatch.2021.Bluray-720p.mkv")
	t.Setenv("radarr_moviefile_size", "3593317970")
	t.Setenv("radarr_movie_tmdbid", "542178")
	t.Setenv("radarr_movie_id", "2173")
	t.Setenv("radarr_moviefile_quality", "Bluray-720p")
	t.Setenv("radarr_movie_title", "The French Dispatch")
	t.Setenv("radarr_moviefile_qualityversion", "1")

	cmd, err := starrcmd.New()
	if err != nil {
		t.Fatalf("got an unexpected error: %s", err)
	}

	switch info, err := cmd.GetRadarrMovieFileDelete(); {
	case err != nil:
		t.Fatalf("got an unexpected error: %s", err)
	case info.IMDbID != "tt1564368":
		t.Fatalf("got an wrong IMDBID? wanted: 'tt1564368', got: %v", info.IMDbID)
	}
}

func TestRadarrMovieDelete(t *testing.T) {
	t.Setenv("radarr_eventtype", string(starrcmd.EventMovieDelete))

	t.Setenv("radarr_movie_id", "2173")
	t.Setenv("radarr_movie_title", "The French Dispatch")
	t.Setenv("radarr_movie_year", "2021")
	t.Setenv("radarr_movie_path", "/movies/The French Dispatch (2021)")
	t.Setenv("radarr_movie_imdbid", "tt1564397")
	t.Setenv("radarr_movie_tmdbid", "542178")
	t.Setenv("radarr_movie_folder_size", "3593317970")
	t.Setenv("radarr_movie_deletedfiles", "XXX: no example. Does this need a split?")

	cmd, err := starrcmd.New()
	if err != nil {
		t.Fatalf("got an unexpected error: %s", err)
	}

	switch info, err := cmd.GetRadarrMovieDelete(); {
	case err != nil:
		t.Fatalf("got an unexpected error: %s", err)
	case info.IMDbID != "tt1564397":
		t.Fatalf("got an wrong IMDBID? wanted: 'tt1564397', got: %v", info.IMDbID)
	}
}
