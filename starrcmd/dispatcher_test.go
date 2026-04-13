package starrcmd_test

import (
	"errors"
	"fmt"
	"testing"

	"golift.io/starr"
	"golift.io/starr/starrcmd"
)

var errDispatcherTest = errors.New("dispatcher test error")

func TestDispatcher_RunSonarrGrab(t *testing.T) {
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

	var sawTitle string

	registry := starrcmd.NewDispatcher()
	registry.Register(starr.Sonarr, starrcmd.EventGrab, func(cmd *starrcmd.CmdEvent) error {
		grab, err := cmd.GetSonarrGrab()
		if err != nil {
			return fmt.Errorf("get sonarr grab: %w", err)
		}

		sawTitle = grab.Title

		return nil
	})

	if err := registry.Run(); err != nil {
		t.Fatal(err)
	}

	if sawTitle != "This Is Us" {
		t.Fatalf("callback title: %q", sawTitle)
	}
}

func TestDispatcher_MultipleCallbacksOrder(t *testing.T) {
	t.Setenv("sonarr_eventtype", string(starrcmd.EventTest))

	var order []int

	registry := starrcmd.NewDispatcher()
	registry.Register(starr.Sonarr, starrcmd.EventTest, func(*starrcmd.CmdEvent) error {
		order = append(order, 1)

		return nil
	})
	registry.Register(starr.Sonarr, starrcmd.EventTest, func(*starrcmd.CmdEvent) error {
		order = append(order, 2)

		return nil
	})

	if err := registry.Run(); err != nil {
		t.Fatal(err)
	}

	if len(order) != 2 || order[0] != 1 || order[1] != 2 {
		t.Fatalf("order: %v", order)
	}
}

func TestDispatcher_OnUnknown(t *testing.T) {
	t.Setenv("radarr_eventtype", string(starrcmd.EventGrab))
	t.Setenv("radarr_movie_title", "X")

	var unknown bool

	registry := starrcmd.NewDispatcher()
	registry.OnUnknown = func(cmd *starrcmd.CmdEvent) error {
		if cmd.App == starr.Radarr && cmd.Type == starrcmd.EventGrab {
			unknown = true
		}

		return nil
	}

	if err := registry.Run(); err != nil {
		t.Fatal(err)
	}

	if !unknown {
		t.Fatal("OnUnknown not called")
	}
}

func TestDispatcher_NoUnknownNoOp(t *testing.T) {
	t.Setenv("prowlarr_eventtype", string(starrcmd.EventTest))

	registry := starrcmd.NewDispatcher()
	if err := registry.Run(); err != nil {
		t.Fatal(err)
	}
}

func TestDispatcher_CallbackError(t *testing.T) {
	t.Setenv("sonarr_eventtype", string(starrcmd.EventTest))

	registry := starrcmd.NewDispatcher()
	registry.Register(starr.Sonarr, starrcmd.EventTest, func(*starrcmd.CmdEvent) error {
		return errDispatcherTest
	})

	err := registry.Run()
	if err == nil || !errors.Is(err, errDispatcherTest) {
		t.Fatalf("err: %v", err)
	}
}

func TestDispatcher_NilReceiver(t *testing.T) {
	t.Parallel()

	var registry *starrcmd.Dispatcher

	if err := registry.Run(); !errors.Is(err, starrcmd.ErrNilDispatcher) {
		t.Fatalf("Run: %v", err)
	}

	if err := registry.Dispatch(&starrcmd.CmdEvent{}); !errors.Is(err, starrcmd.ErrNilDispatcher) {
		t.Fatalf("Dispatch: %v", err)
	}

	registry.Register(starr.Sonarr, starrcmd.EventTest, func(*starrcmd.CmdEvent) error { return nil }) // no panic
}

func TestDispatcher_NilCmd(t *testing.T) {
	t.Parallel()

	registry := starrcmd.NewDispatcher()
	if err := registry.Dispatch(nil); !errors.Is(err, starrcmd.ErrNilCmdEvent) {
		t.Fatalf("err: %v", err)
	}
}
