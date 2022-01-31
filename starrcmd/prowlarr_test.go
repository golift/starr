//nolint:paralleltest,dupl
package starrcmd_test

import (
	"os"
	"testing"

	"golift.io/starr/starrcmd"
)

func TestProwlarrApplicationUpdate(t *testing.T) {
	starrcmd.EventType = starrcmd.EventApplicationUpdate

	t.Setenv("prowlarr_eventtype", string(starrcmd.EventApplicationUpdate))
	t.Setenv("prowlarr_update_previousversion", "4.0.3.5875")
	t.Setenv("prowlarr_update_newversion", "4.0.4.5909")
	t.Setenv("prowlarr_update_message", "Prowlarr updated from 4.0.3.5875 to 4.0.4.5909")

	switch info, err := starrcmd.GetProwlarrApplicationUpdate(); {
	case err != nil:
		t.Fatalf("got an unexpected error: %s", err)
	case info.Message != os.Getenv("prowlarr_update_message"):
		t.Fatalf("got wrong Message? %s", info.Message)
	case info.NewVersion != "4.0.4.5909":
		t.Fatalf("got wrong new version? wanted: '4.0.4.5909' got: %s", info.Message)
	case info.PreviousVersion != "4.0.3.5875":
		t.Fatalf("got wrong Message? wanted: '4.0.3.5875' got: %s", info.Message)
	}
}

func TestProwlarrHealthIssue(t *testing.T) {
	starrcmd.EventType = starrcmd.EventHealthIssue

	t.Setenv("prowlarr_eventtype", string(starrcmd.EventHealthIssue))
	t.Setenv("prowlarr_health_issue_type", "SomeIssueTypeForProwlarr")
	t.Setenv("prowlarr_health_issue_wiki", "https://wiki.servarr.com/prowlarr")
	t.Setenv("prowlarr_health_issue_level", "Error")
	t.Setenv("prowlarr_health_issue_message", "Lists unavailable due to failures: List name here")

	switch info, err := starrcmd.GetProwlarrHealthIssue(); {
	case err != nil:
		t.Fatalf("got an unexpected error: %s", err)
	case info.Message != os.Getenv("prowlarr_health_issue_message"):
		t.Fatalf("got wrong Message? %s", info.Message)
	case info.Wiki != "https://wiki.servarr.com/prowlarr":
		t.Fatalf("got wrong wiki link? wanted: 'https://wiki.servarr.com/prowlarr' got: %s", info.Wiki)
	case info.Level != "Error":
		t.Fatalf("got wrong level? wanted: 'Warning' got: %s", info.Level)
	case info.IssueType != "SomeIssueTypeForProwlarr":
		t.Fatalf("got wrong issue type? wanted: 'ImportListStatusCheck' got: %s", info.IssueType)
	}
}

func TestProwlarrTest(t *testing.T) {
	starrcmd.EventType = starrcmd.EventTest

	t.Setenv("prowlarr_eventtype", string(starrcmd.EventTest))

	switch info, err := starrcmd.GetProwlarrTest(); {
	case err != nil:
		t.Fatalf("got an unexpected error: %s", err)
	case info != starrcmd.ProwlarrTest{}:
		t.Fatalf("got an wrong structure in return")
	}
}
