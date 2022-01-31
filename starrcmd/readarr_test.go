//nolint:paralleltest,dupl
package starrcmd_test

import (
	"os"
	"testing"

	"golift.io/starr/starrcmd"
)

func TestReadarrApplicationUpdate(t *testing.T) {
	starrcmd.EventType = starrcmd.EventApplicationUpdate

	t.Setenv("readarr_eventtype", string(starrcmd.EventApplicationUpdate))
	t.Setenv("readarr_update_previousversion", "6.0.3.5875")
	t.Setenv("readarr_update_newversion", "6.0.4.5909")
	t.Setenv("readarr_update_message", "Readarr updated from 6.0.3.5875 to 6.0.4.5909")

	switch info, err := starrcmd.GetReadarrApplicationUpdate(); {
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
	starrcmd.EventType = starrcmd.EventHealthIssue

	t.Setenv("readarr_eventtype", string(starrcmd.EventHealthIssue))
	t.Setenv("readarr_health_issue_type", "SomeIssueTypeForReadarr")
	t.Setenv("readarr_health_issue_wiki", "https://wiki.servarr.com/readarr")
	t.Setenv("readarr_health_issue_level", "Info")
	t.Setenv("readarr_health_issue_message", "Lists unavailable due to failures: List name here")

	switch info, err := starrcmd.GetReadarrHealthIssue(); {
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
	starrcmd.EventType = starrcmd.EventTest

	t.Setenv("readarr_eventtype", string(starrcmd.EventTest))

	switch info, err := starrcmd.GetReadarrTest(); {
	case err != nil:
		t.Fatalf("got an unexpected error: %s", err)
	case info != starrcmd.ReadarrTest{}:
		t.Fatalf("got an wrong structure in return")
	}
}
