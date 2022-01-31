//nolint:paralleltest,dupl
package starrcmd_test

import (
	"os"
	"testing"

	"golift.io/starr/starrcmd"
)

func TestReadarrApplicationUpdate(t *testing.T) {
	t.Setenv("readarr_eventtype", string(starrcmd.EventApplicationUpdate))
	t.Setenv("readarr_update_previousversion", "6.0.3.5875")
	t.Setenv("readarr_update_newversion", "6.0.4.5909")
	t.Setenv("readarr_update_message", "Readarr updated from 6.0.3.5875 to 6.0.4.5909")

	cmd, err := starrcmd.New()
	if err != nil {
		t.Fatalf("got an unexpected error: %s", err)
	}

	switch info, err := cmd.GetReadarrApplicationUpdate(); {
	case err != nil:
		t.Fatalf("got an unexpected error: %s", err)
	case info.Message != os.Getenv("readarr_update_message"):
		t.Fatalf("got wrong Message? %s", info.Message)
	case info.NewVersion != "6.0.4.5909":
		t.Fatalf("got wrong new version? wanted: '6.0.4.5909' got: %s", info.Message)
	case info.PreviousVersion != "6.0.3.5875":
		t.Fatalf("got wrong Message? wanted: '6.0.3.5875' got: %s", info.Message)
	}
}

func TestReadarrHealthIssue(t *testing.T) {
	t.Setenv("readarr_eventtype", string(starrcmd.EventHealthIssue))
	t.Setenv("readarr_health_issue_type", "SomeIssueTypeForReadarr")
	t.Setenv("readarr_health_issue_wiki", "https://wiki.servarr.com/readarr")
	t.Setenv("readarr_health_issue_level", "Info")
	t.Setenv("readarr_health_issue_message", "Lists unavailable due to failures: List name here")

	cmd, err := starrcmd.New()
	if err != nil {
		t.Fatalf("got an unexpected error: %s", err)
	}

	switch info, err := cmd.GetReadarrHealthIssue(); {
	case err != nil:
		t.Fatalf("got an unexpected error: %s", err)
	case info.Message != os.Getenv("readarr_health_issue_message"):
		t.Fatalf("got wrong Message? %s", info.Message)
	case info.Wiki != "https://wiki.servarr.com/readarr":
		t.Fatalf("got wrong wiki link? wanted: 'https://wiki.servarr.com/readarr' got: %s", info.Wiki)
	case info.Level != "Info":
		t.Fatalf("got wrong level? wanted: 'Info' got: %s", info.Level)
	case info.IssueType != "SomeIssueTypeForReadarr":
		t.Fatalf("got wrong issue type? wanted: 'ImportListStatusCheck' got: %s", info.IssueType)
	}
}

func TestReadarrTest(t *testing.T) {
	t.Setenv("readarr_eventtype", string(starrcmd.EventTest))

	cmd, err := starrcmd.New()
	if err != nil {
		t.Fatalf("got an unexpected error: %s", err)
	}

	switch info, err := cmd.GetReadarrTest(); {
	case err != nil:
		t.Fatalf("got an unexpected error: %s", err)
	case info != starrcmd.ReadarrTest{}:
		t.Fatalf("got an wrong structure in return")
	}
}

func TestReadarrGrab(t *testing.T) {
	t.Setenv("readarr_eventtype", string(starrcmd.EventGrab))
	t.Setenv("readarr_author_grid", "1077326")
	t.Setenv("readarr_release_releasegroup", "BitBook")
	t.Setenv("readarr_author_name", "J.K. Rowling")
	t.Setenv("readarr_release_title", "J K Rowling - Harry Potter and the Order of the Phoenix")
	t.Setenv("readarr_release_grids", "21175582 // not sure what this looks like with 2+")
	t.Setenv("readarr_download_client", "qBittorrent")
	t.Setenv("readarr_release_size", "1279262")
	t.Setenv("readarr_release_qualityversion", "1")
	t.Setenv("readarr_release_booktitles", "Harry Potter and the Order of the Phoenix")
	t.Setenv("readarr_release_bookids", "649")
	t.Setenv("readarr_release_indexer", "InfoWars (Prowlarr)")
	t.Setenv("readarr_download_id", "3852BA2204A84185B2B43281E53BE93D56DE5C81")
	t.Setenv("readarr_release_bookcount", "1")
	t.Setenv("readarr_release_bookreleasedates", "07/10/2003 07:00:00")
	t.Setenv("readarr_release_quality", "EPUB")
	t.Setenv("readarr_author_id", "4")

	cmd, err := starrcmd.New()
	if err != nil {
		t.Fatalf("got an unexpected error: %s", err)
	}

	switch info, err := cmd.GetReadarrGrab(); {
	case err != nil:
		t.Fatalf("got an unexpected error: %s", err)
	case info.AuthorName != "J.K. Rowling":
		t.Fatalf("got an wrong author name? wanted: 'J.K. Rowling', got: %v", info.AuthorName)
	}
}

func TestReadarrBookDelete(t *testing.T) {
	t.Setenv("readarr_eventtype", string(starrcmd.EventBookDelete))
	t.Setenv("readarr_author_name", "Alyssa Cole")
	t.Setenv("readarr_book_goodreadsid", "88514853")
	t.Setenv("readarr_author_goodreadsid", "7790155")
	t.Setenv("readarr_book_title", "Unti Cole #6: A Novel")
	t.Setenv("readarr_author_path", "/books/Alyssa Cole")
	t.Setenv("readarr_book_id", "636")
	t.Setenv("readarr_book_deletedfiles", "True")
	t.Setenv("readarr_author_id", "33")

	cmd, err := starrcmd.New()
	if err != nil {
		t.Fatalf("got an unexpected error: %s", err)
	}

	switch info, err := cmd.GetReadarrBookDelete(); {
	case err != nil:
		t.Fatalf("got an unexpected error: %s", err)
	case info.AuthorName != "Alyssa Cole":
		t.Fatalf("got an wrong author name? wanted: 'Alyssa Cole', got: %v", info.AuthorName)
	}
}

func TestReadarrBookFileDelete(t *testing.T) {
	t.Setenv("readarr_eventtype", string(starrcmd.EventBookFileDelete))
	t.Setenv("readarr_delete_reason", "deleteMessage.Reason.ToString())")
	t.Setenv("readarr_author_id", "54756546")
	t.Setenv("readarr_author_name", "author.Name)")
	t.Setenv("readarr_author_goodreadsid", "34234234")
	t.Setenv("readarr_book_id", "456454345")
	t.Setenv("readarr_book_title", "book.Title)")
	t.Setenv("readarr_book_goodreadsid", "324324234")
	t.Setenv("readarr_bookfile_id", "7323445")
	t.Setenv("readarr_bookfile_path", "bookFile.Path)")
	t.Setenv("readarr_bookfile_quality", "bookFile.Quality.Quality.Name)")
	t.Setenv("readarr_bookfile_qualityversion", "1")
	t.Setenv("readarr_bookfile_releasegroup", "bookFile.ReleaseGroup ?? string.Empty)")
	t.Setenv("readarr_bookfile_scenename", "bookFile.SceneName ?? string.Empty)")
	t.Setenv("readarr_bookfile_edition_id", "213123")
	t.Setenv("readarr_bookfile_edition_name", "edition.Title)")
	t.Setenv("readarr_bookfile_edition_goodreadsid", "324234")
	t.Setenv("readarr_bookfile_edition_isbn13", "edition.Isbn13)")
	t.Setenv("readarr_bookfile_edition_asin", "edition.Asin)")

	cmd, err := starrcmd.New()
	if err != nil {
		t.Fatalf("got an unexpected error: %s", err)
	}

	switch info, err := cmd.GetReadarrBookFileDelete(); {
	case err != nil:
		t.Fatalf("got an unexpected error: %s", err)
	case info.AuthorName != "author.Name)":
		t.Fatalf("got an wrong author name? wanted: 'author.Name)', got: %v", info.AuthorName)
	}
}

func TestReadarrAuthorDelete(t *testing.T) {
	t.Setenv("readarr_eventtype", string(starrcmd.EventAuthorDelete))
	t.Setenv("readarr_author_id", "34534534")
	t.Setenv("readarr_author_name", "author.Name)")
	t.Setenv("readarr_author_path", "author.Path)")
	t.Setenv("readarr_author_goodreadsid", "234234234")
	t.Setenv("readarr_author_deletedfiles", "False")

	cmd, err := starrcmd.New()
	if err != nil {
		t.Fatalf("got an unexpected error: %s", err)
	}

	switch info, err := cmd.GetReadarrAuthorDelete(); {
	case err != nil:
		t.Fatalf("got an unexpected error: %s", err)
	case info.AuthorName != "author.Name)":
		t.Fatalf("got an wrong author name? wanted: 'author.Name)', got: %v", info.AuthorName)
	}
}

func TestReadarrRename(t *testing.T) {
	t.Setenv("readarr_eventtype", string(starrcmd.EventRename))
	t.Setenv("readarr_author_id", "16128787")
	t.Setenv("readarr_author_name", "author.Metadata.Value.Name)")
	t.Setenv("readarr_author_path", "author.Path)")
	t.Setenv("readarr_author_grid", "234234234")

	cmd, err := starrcmd.New()
	if err != nil {
		t.Fatalf("got an unexpected error: %s", err)
	}

	switch info, err := cmd.GetReadarrRename(); {
	case err != nil:
		t.Fatalf("got an unexpected error: %s", err)
	case info.AuthorName != "author.Metadata.Value.Name)":
		t.Fatalf("got an wrong author name? wanted: 'author.Metadata.Value.Name)', got: %v", info.AuthorName)
	}
}

func TestReadarrDownload(t *testing.T) {
	t.Setenv("readarr_eventtype", string(starrcmd.EventDownload))
	t.Setenv("readarr_author_id", "9182398")
	t.Setenv("readarr_author_name", "le author")
	t.Setenv("readarr_author_path", "author.Path)")
	t.Setenv("readarr_author_grid", "2234234")
	t.Setenv("readarr_book_id", "012338")
	t.Setenv("readarr_book_title", "book.Title)")
	t.Setenv("readarr_book_grid", "123123123")
	t.Setenv("readarr_book_releasedate", "09/01/2003 07:00:00")
	t.Setenv("readarr_download_client", "message.DownloadClient ?? string.Empty)")
	t.Setenv("readarr_download_id", "message.DownloadId ?? string.Empty)")
	t.Setenv("readarr_addedbookpaths", "")
	t.Setenv("readarr_deletedpaths", "")

	cmd, err := starrcmd.New()
	if err != nil {
		t.Fatalf("got an unexpected error: %s", err)
	}

	switch info, err := cmd.GetReadarrDownload(); {
	case err != nil:
		t.Fatalf("got an unexpected error: %s", err)
	case info.AuthorName != "le author":
		t.Fatalf("got an wrong author name? wanted: 'le author', got: %v", info.AuthorName)
	}
}

func TestReadarrTrackRetag(t *testing.T) {
	t.Setenv("readarr_eventtype", string(starrcmd.EventTrackRetag))
	t.Setenv("readarr_author_id", "1232131313")
	t.Setenv("readarr_author_name", "write here")
	t.Setenv("readarr_author_path", "author.Path)")
	t.Setenv("readarr_author_grid", "324324")
	t.Setenv("readarr_book_id", "676757")
	t.Setenv("readarr_book_title", "book.Title)")
	t.Setenv("readarr_book_grid", "123123123")
	t.Setenv("readarr_book_releasedate", "11/11/2003 17:00:00")
	t.Setenv("readarr_bookfile_id", "4565665")
	t.Setenv("readarr_bookfile_path", "bookFile.Path)")
	t.Setenv("readarr_bookfile_quality", "bookFile.Quality.Quality.Name)")
	t.Setenv("readarr_bookfile_qualityversion", "1")
	t.Setenv("readarr_bookfile_releasegroup", "bookFile.ReleaseGroup ?? string.Empty)")
	t.Setenv("readarr_bookfile_scenename", "bookFile.SceneName ?? string.Empty)")
	t.Setenv("readarr_tags_diff", "message.Diff.ToJson())")
	t.Setenv("readarr_tags_scrubbed", "False")

	cmd, err := starrcmd.New()
	if err != nil {
		t.Fatalf("got an unexpected error: %s", err)
	}

	switch info, err := cmd.GetReadarrTrackRetag(); {
	case err != nil:
		t.Fatalf("got an unexpected error: %s", err)
	case info.AuthorName != "write here":
		t.Fatalf("got an wrong author name? wanted: 'write here', got: %v", info.AuthorName)
	}
}
