# Starr Connect

[![Go Reference](https://pkg.go.dev/badge/golift.io/starr/starrconnect.svg)](https://pkg.go.dev/golift.io/starr/starrconnect)

`starrconnect` parses **HTTP webhook** JSON from Sonarr, Radarr, Lidarr, and Prowlarr (Settings → Connect → Webhook). Each app exposes:

- **`Parse*([]byte)`** — read the envelope (`eventType`, `instanceName`, `applicationUrl`) and keep the raw body.
- **`Get*()` methods** on the parsed event — decode the full payload for that event (with `errors.Is` / `errors.As` against **`ErrWrongEvent`** if you call the wrong getter).
- **`*Handler` types** — implement **`http.Handler`** and dispatch POST bodies to optional **`On…`** callbacks.

For **Custom Script** env-based payloads, use package **`starrcmd`** instead ([README](../starrcmd)).

Smaller examples live in [example_test.go](example_test.go).

---

## Holistic example: one server, many webhooks

A common setup is a single HTTP server with one URL per app. Each handler only runs the callbacks you care about; omitted callbacks are no-ops. Returning an error from a callback produces **500** for that request (after logging via **`OnError`**).

```go
package main

import (
	"log"
	"net/http"
	"time"

	"golift.io/starr/starrconnect"
)

func main() {
	logger := log.Default()
	mux := http.NewServeMux()

	onErr := func(err error) { logger.Println(err) }

	mux.Handle("/hooks/sonarr", &starrconnect.SonarrHandler{
		OnGrab: func(g *starrconnect.SonarrGrab) error {
			if g.Series != nil {
				logger.Printf("sonarr %s@%s grab: %s", g.InstanceName, g.ApplicationURL, g.Series.Title)
			}
			return nil
		},
		OnDownload: func(d *starrconnect.SonarrDownload) error {
			if d.Series != nil {
				logger.Printf("sonarr download: %s", d.Series.Title)
			}
			return nil
		},
		OnImportComplete: func(ic *starrconnect.SonarrImportComplete) error {
			logger.Printf("sonarr import complete: %d files", ic.FileCount)
			return nil
		},
		OnHealth: func(h *starrconnect.SonarrHealth) error {
			logger.Printf("sonarr health [%s]: %s", h.Level, h.Message)
			return nil
		},
		OnError: onErr,
	})

	mux.Handle("/hooks/radarr", &starrconnect.RadarrHandler{
		OnGrab: func(g *starrconnect.RadarrGrab) error {
			if g.Movie != nil {
				logger.Printf("radarr grab: %s (%d)", g.Movie.Title, g.Movie.Year)
			}
			return nil
		},
		OnApplicationUpdate: func(u *starrconnect.RadarrApplicationUpdate) error {
			logger.Printf("radarr upgraded: %s -> %s", u.PreviousVersion, u.NewVersion)
			return nil
		},
		OnError: onErr,
	})

	mux.Handle("/hooks/lidarr", &starrconnect.LidarrHandler{
		OnGrab: func(g *starrconnect.LidarrGrab) error {
			if g.Artist != nil {
				logger.Printf("lidarr grab: %s", g.Artist.Name)
			}
			return nil
		},
		OnImportFailure: func(f *starrconnect.LidarrImportFailure) error {
			logger.Printf("lidarr import failure: downloadId=%s", f.DownloadID)
			return nil
		},
		OnError: onErr,
	})

	mux.Handle("/hooks/prowlarr", &starrconnect.ProwlarrHandler{
		OnGrab: func(g *starrconnect.ProwlarrGrab) error {
			if g.Release != nil {
				logger.Printf("prowlarr grab: %s via %s", g.Release.ReleaseTitle, g.Release.Indexer)
			}
			return nil
		},
		OnTest: func(t *starrconnect.ProwlarrTest) error {
			logger.Printf("prowlarr test: %s", t.InstanceName)
			return nil
		},
		OnError: onErr,
	})

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
	}
	logger.Fatal(srv.ListenAndServe())
}
```

Point each app’s webhook URL at the matching path (only **POST** is accepted). Use HTTPS and authentication in front of this service in production.

---

## Parsing without `http.Handler`

If the body arrives from a queue, file, or tests, parse by app and branch on **`EventType`**. For Sonarr **`EventDownload`**, try **`GetDownload`** first; if the payload is import-complete, that call fails with **`ErrWrongEvent`** and you should use **`GetImportComplete`** instead.

```go
import (
	"errors"

	"golift.io/starr/starrconnect"
)

func handleSonarrBody(body []byte) error {
	envelope, err := starrconnect.ParseSonarr(body)
	if err != nil {
		return err
	}

	switch envelope.EventType {
	case starrconnect.EventGrab:
		grab, err := envelope.GetGrab()
		if err != nil {
			return err
		}
		_ = grab
	case starrconnect.EventDownload:
		dl, err := envelope.GetDownload()
		if err == nil {
			_ = dl
		} else if !errors.Is(err, starrconnect.ErrWrongEvent) {
			return err
		} else {
			ic, err := envelope.GetImportComplete()
			if err != nil {
				return err
			}
			_ = ic
		}
	default:
		// ignore or log
	}

	return nil
}
```

---

## Further reading

- [Servarr webhook wiki](https://wiki.servarr.com/sonarr/settings#connections) (same idea across apps).
- Package doc comments in the Go source for payload quirks (Sonarr download shapes, Lidarr import failure, Prowlarr events emitted today, etc.).
