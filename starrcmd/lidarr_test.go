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
