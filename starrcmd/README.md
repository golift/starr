# Starr Command

[![Go Reference](https://pkg.go.dev/badge/golift.io/starr/starrcmd.svg)](https://pkg.go.dev/golift.io/starr/starrcmd)

`starrcmd` reads **Custom Script** hooks from Radarr, Sonarr, Lidarr, Readarr, and Prowlarr. When an app runs your executable, it sets **`{app}_eventtype`** and many other environment variables. This package:

- Discovers which app fired using **`starrcmd.New()`** (checks `radarr_eventtype`, `sonarr_eventtype`, `lidarr_eventtype`, `readarr_eventtype`, `prowlarr_eventtype`).
- Exposes typed structs whose **`env:"..."`** tags map to those variable names.
- Provides **`Get…()` methods on `*CmdEvent`** that verify **`cmd.Type`** and fill the struct from the environment.
- Optionally routes events with a **`Dispatcher`**: typed helpers like **`OnRadarrGrab(func(RadarrGrab) error)`** (one per app/event), or low-level **`Register(starr.App, Event, func(*CmdEvent) error)`**, then **`Run()`** (wraps **`New()`**) or **`Dispatch(cmd)`** for tests.

For **HTTP Webhook** JSON instead of env vars, use package **[starrconnect](../starrconnect)**.

Configure scripts under **Settings → Connect → Custom Script** in each app. See the upstream [Custom Scripts](https://wiki.servarr.com/radarr/custom-scripts) wiki (same pattern across apps).

Smaller patterns live in [example_test.go](example_test.go). Callback routing is covered in [dispatcher_test.go](dispatcher_test.go).

---

## Dispatcher (recommended)

Typed **`On{App}{Event}`** methods register for a single **`(starr.App, Event)`** pair, **`Run()`** / **`Dispatch`** unmarshals env vars into the payload struct, then calls your handler. You can still use **`Register`** when you need **`*CmdEvent`** without a dedicated helper. Callbacks for the same pair run in registration order. If nothing matches, **`OnUnknown`** runs when set; otherwise it is a no-op success. The first callback error is returned from **`Run`** / **`Dispatch`**.

```go
package main

import (
	"fmt"
	"log"

	"golift.io/starr/starrcmd"
)

func main() {
	registry := starrcmd.NewDispatcher()
	registry.OnRadarrGrab(func(grab starrcmd.RadarrGrab) error {
		fmt.Println(grab.Title)
		return nil
	})
	registry.OnUnknown = func(cmd *starrcmd.CmdEvent) error {
		log.Printf("unhandled: %s %s", cmd.App, cmd.Type)
		return nil
	}
	if err := registry.Run(); err != nil {
		log.Fatal(err)
	}
}
```

Only one app sets **`{app}_eventtype`** per invocation, so **`Run()`** matches at most one **`(app, event)`** — but you typically register handlers for every app your binary is wired to; the rest stay idle until that app fires.

```go
package main

import (
	"fmt"
	"log"

	"golift.io/starr/starrcmd"
)

func main() {
	registry := starrcmd.NewDispatcher()

	registry.OnSonarrGrab(func(g starrcmd.SonarrGrab) error {
		fmt.Println("sonarr grab:", g.Title, g.ReleaseTitle)
		return nil
	})
	registry.OnSonarrDownload(func(d starrcmd.SonarrDownload) error {
		fmt.Println("sonarr download:", d.Title, d.EpisodePath)
		return nil
	})
	registry.OnLidarrAlbumDownload(func(d starrcmd.LidarrAlbumDownload) error {
		fmt.Println("lidarr album download:", d.ArtistName, d.Title)
		return nil
	})
	registry.OnRadarrHealthIssue(func(h starrcmd.RadarrHealthIssue) error {
		fmt.Println("radarr health:", h.Level, h.Message)
		return nil
	})

	registry.OnUnknown = func(cmd *starrcmd.CmdEvent) error {
		log.Printf("no handler for %s %s", cmd.App, cmd.Type)
		return nil
	}
	// This can go in a routine.
	if err := registry.Run(); err != nil {
		log.Fatal(err)
	}
}
```

For unit tests, build a **`CmdEvent`** (or use **`New()`** with **`t.Setenv`**) and call **`Dispatch(cmd)`** instead of **`Run()`**. A nil **`Dispatcher`** or nil **`cmd`** yields **`ErrNilDispatcher`** / **`ErrNilCmdEvent`**.

---

## Holistic example: one binary, every app

**This is how we used to do it before the dispatcher above. You can still do this, but the dispatcher is cleaner.**

Your script is usually a single binary on disk, registered in more than one app. Call **`New()`**
once, branch on **`cmd.App`**, then on **`cmd.Type`**, and use the matching **`Get*`** for that app.
Wrong event for the chosen getter returns **`ErrInvalidEvent`** (wrap with **`errors.Is`** if you
want to fall through).

```go
package main

import (
	"errors"
	"fmt"
	"log"

	"golift.io/starr"
	"golift.io/starr/starrcmd"
)

func main() {
	cmd, err := starrcmd.New()
	if err != nil {
		log.Fatalf("not invoked from a Starr custom script (missing *_eventtype): %v", err)
	}

	switch cmd.App {
	case starr.Radarr:
		handleRadarr(cmd)
	case starr.Sonarr:
		handleSonarr(cmd)
	case starr.Lidarr:
		handleLidarr(cmd)
	case starr.Readarr:
		handleReadarr(cmd)
	case starr.Prowlarr:
		handleProwlarr(cmd)
	default:
		log.Fatalf("unknown app: %v", cmd.App)
	}
}

func handleRadarr(cmd *starrcmd.CmdEvent) {
	switch cmd.Type {
	case starrcmd.EventGrab:
		grab, err := cmd.GetRadarrGrab()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("grab:", grab.Title, grab.ReleaseTitle)
	case starrcmd.EventDownload:
		dl, err := cmd.GetRadarrDownload()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("download:", dl.Title, dl.FilePath)
	case starrcmd.EventHealthIssue:
		h, err := cmd.GetRadarrHealthIssue()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("health:", h.Level, h.Message)
	case starrcmd.EventTest:
		// optional: _, _ = cmd.GetRadarrTest()
		return
	default:
		fmt.Println("unhandled Radarr event:", cmd.Type)
	}
}

func handleSonarr(cmd *starrcmd.CmdEvent) {
	switch cmd.Type {
	case starrcmd.EventGrab:
		grab, err := cmd.GetSonarrGrab()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("grab:", grab.Title, grab.ReleaseTitle)
	case starrcmd.EventDownload:
		dl, err := cmd.GetSonarrDownload()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("download:", dl.Title, dl.EpisodePath)
	default:
		fmt.Println("unhandled Sonarr event:", cmd.Type)
	}
}

func handleLidarr(cmd *starrcmd.CmdEvent) {
	switch cmd.Type {
	case starrcmd.EventGrab:
		grab, err := cmd.GetLidarrGrab()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("grab:", grab.ArtistName, grab.ReleaseTitle)
	case starrcmd.EventAlbumDownload:
		dl, err := cmd.GetLidarrAlbumDownload()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("album download:", dl.ArtistName, dl.Title, dl.Path)
	default:
		fmt.Println("unhandled Lidarr event:", cmd.Type)
	}
}

func handleReadarr(cmd *starrcmd.CmdEvent) {
	_, err := cmd.GetReadarrGrab()
	if err != nil && !errors.Is(err, starrcmd.ErrInvalidEvent) {
		log.Fatal(err)
	}
	if err == nil {
		fmt.Println("readarr grab (example)")
		return
	}
	fmt.Println("unhandled or non-grab Readarr event:", cmd.Type)
}

func handleProwlarr(cmd *starrcmd.CmdEvent) {
	switch cmd.Type {
	case starrcmd.EventHealthIssue:
		h, err := cmd.GetProwlarrHealthIssue()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("prowlarr health:", h.Message)
	default:
		fmt.Println("unhandled Prowlarr event:", cmd.Type)
	}
}
```

Use **`NewMust()`** only if a missing event should panic; **`NewMustNoPanic()`** returns an empty **`CmdEvent`** when nothing is set (see package docs).

---

## Testing and env vars

In tests, set the right **`{app}_eventtype`** and any **`env:`** keys your structs need. Slice fields use a **split character** in the struct tag (for example **`",,"`** or **`"|"`**); omitting it where the parser expects one can **panic**—see **`parser.go`** / **`config.go`** developer notes and the existing `*_test.go` files for patterns.

---

## Further reading

- Package **`config.go`** comments: date formats, supported field types, and slice rules.
- **`starrconnect`** for JSON webhooks: [starrconnect](../starrconnect).
