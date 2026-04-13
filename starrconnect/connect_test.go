package starrconnect_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"golift.io/starr/starrconnect"
)

const oversizedWebhookBody = 10<<20 + 10

var errBadSeries = errors.New("bad series")

func TestParseSonarrHealth(t *testing.T) {
	t.Parallel()

	const body = `{"eventType":"Health","instanceName":"Sonarr","applicationUrl":"http://x","level":"warning","message":"m","type":"t","wikiUrl":"w"}`

	envelope, err := starrconnect.ParseSonarr([]byte(body))
	if err != nil {
		t.Fatal(err)
	}

	if envelope.EventType != starrconnect.EventHealth {
		t.Fatalf("event: %s", envelope.EventType)
	}

	health, err := envelope.GetHealth()
	if err != nil {
		t.Fatal(err)
	}

	if health.Level != "warning" || health.Message != "m" || health.Type != "t" || health.WikiURL != "w" {
		t.Fatalf("health: %+v", health)
	}
}

func TestSonarrDownloadSingle(t *testing.T) {
	t.Parallel()

	const dl = `{"eventType":"Download","instanceName":"S","applicationUrl":"http://x","series":{"id":1,"title":"T"},"episodes":[],"episodeFile":{"id":2,"relativePath":"a.mkv","path":"/tv/a.mkv","quality":"HDTV-720p","qualityVersion":1,"releaseGroup":"","sceneName":"","size":1,"dateAdded":"2020-01-01T00:00:00Z","languages":[]},"isUpgrade":false,"downloadClient":"c","downloadClientType":"nzbget","downloadId":"id","release":{"releaseTitle":"r"}}`

	envelope, err := starrconnect.ParseSonarr([]byte(dl))
	if err != nil {
		t.Fatal(err)
	}

	if _, err := envelope.GetImportComplete(); err == nil {
		t.Fatal("expected error for single-file download")
	}

	download, err := envelope.GetDownload()
	if err != nil {
		t.Fatal(err)
	}

	if download.EpisodeFile == nil || download.EpisodeFile.ID != 2 {
		t.Fatalf("episodeFile: %+v", download.EpisodeFile)
	}
}

func TestSonarrImportCompleteBatch(t *testing.T) {
	t.Parallel()

	const icBody = `{"eventType":"Download","instanceName":"S","applicationUrl":"http://x","series":{"id":1,"title":"T"},"episodes":[],"episodeFiles":[{"id":2,"relativePath":"a.mkv","path":"/tv/a.mkv","quality":"HDTV-720p","qualityVersion":1,"releaseGroup":"","sceneName":"","size":1,"dateAdded":"2020-01-01T00:00:00Z","languages":[]}],"downloadClient":"c","downloadClientType":"nzbget","downloadId":"id","release":{"releaseTitle":"r"},"fileCount":1,"sourcePath":"/dl","destinationPath":"/tv"}`

	envelope, err := starrconnect.ParseSonarr([]byte(icBody))
	if err != nil {
		t.Fatal(err)
	}

	if _, err := envelope.GetDownload(); err == nil {
		t.Fatal("expected error for import-complete shape")
	}

	importComplete, err := envelope.GetImportComplete()
	if err != nil {
		t.Fatal(err)
	}

	if len(importComplete.EpisodeFiles) != 1 || importComplete.SourcePath != "/dl" {
		t.Fatalf("import complete: %+v", importComplete)
	}
}

func TestSonarrHandlerGrab(t *testing.T) {
	t.Parallel()

	var called bool

	handler := &starrconnect.SonarrHandler{
		OnGrab: func(grab *starrconnect.SonarrGrab) error {
			called = true

			if grab.Series == nil || grab.Series.Title != "Show" {
				return errBadSeries
			}

			return nil
		},
	}

	body := map[string]any{
		"eventType":          "Grab",
		"instanceName":       "i",
		"applicationUrl":     "http://u",
		"downloadClient":     "c",
		"downloadClientType": "nzbget",
		"downloadId":         "d",
		"series": map[string]any{
			"id": 1, "title": "Show", "titleSlug": "show", "path": "/tv", "tvdbId": 1, "tvMazeId": 0, "tmdbId": 0,
			"imdbId": "", "type": "standard", "year": 2020,
		},
		"episodes": []any{},
		"release": map[string]any{
			"quality": "HDTV-720p", "qualityVersion": 1, "releaseGroup": "g", "releaseTitle": "rt", "indexer": "i", "size": 100,
		},
	}

	raw, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequestWithContext(context.Background(), http.MethodPost, "/hook", bytes.NewReader(raw))
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status %d body %s", rec.Code, rec.Body.String())
	}

	if !called {
		t.Fatal("OnGrab not called")
	}
}

func assertProwlarrGrabFixture(t *testing.T, grab *starrconnect.ProwlarrGrab) {
	t.Helper()

	if grab.Release == nil || grab.Release.ReleaseTitle != "r" || grab.Release.Indexer != "i" {
		t.Fatalf("release: %+v", grab.Release)
	}

	if grab.Release.Size != 99 {
		t.Fatalf("size: %d", grab.Release.Size)
	}

	if grab.Trigger != "manual" || grab.DownloadClient != "c" || grab.DownloadID != "id" {
		t.Fatalf("grab: %+v", grab)
	}
}

func TestParseProwlarrGrab(t *testing.T) {
	t.Parallel()

	const body = `{"eventType":"Grab","instanceName":"P","applicationUrl":"http://p","release":{"releaseTitle":"r","indexer":"i","size":99,"categories":["TV"],"genres":["Drama"],"indexerFlags":["freeleech"],"publishDate":"2020-01-01T00:00:00Z"},"trigger":"manual","source":"s","host":"h","downloadClient":"c","downloadClientType":"nzbget","downloadId":"id"}`

	envelope, err := starrconnect.ParseProwlarr([]byte(body))
	if err != nil {
		t.Fatal(err)
	}

	if envelope.EventType != starrconnect.EventGrab {
		t.Fatalf("event: %s", envelope.EventType)
	}

	grab, err := envelope.GetGrab()
	if err != nil {
		t.Fatal(err)
	}

	assertProwlarrGrabFixture(t, grab)
}

func TestRadarrUnknownEvent(t *testing.T) {
	t.Parallel()

	handler := &starrconnect.RadarrHandler{
		OnError: func(err error) {
			if err == nil {
				t.Error("expected error")
			}
		},
	}

	req := httptest.NewRequestWithContext(
		context.Background(), http.MethodPost, "/", bytes.NewReader([]byte(`{"eventType":"Nope","instanceName":"x"}`)))
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("want 500 got %d", rec.Code)
	}
}

func TestLidarrImportFailure(t *testing.T) {
	t.Parallel()

	const body = `{"eventType":"ImportFailure","instanceName":"L","applicationUrl":"http://x","artist":{"id":1,"name":"A","disambiguation":"","path":"/m","mbId":"mbid","type":"Standard","overview":"","genres":[],"images":[],"tags":[]},"album":null,"tracks":[],"trackFiles":[],"deletedFiles":[],"isUpgrade":false,"downloadClient":"c","downloadClientType":"qbit","downloadId":"id"}`

	envelope, err := starrconnect.ParseLidarr([]byte(body))
	if err != nil {
		t.Fatal(err)
	}

	importFailure, err := envelope.GetImportFailure()
	if err != nil {
		t.Fatal(err)
	}

	if importFailure.Artist == nil || importFailure.Artist.Name != "A" {
		t.Fatalf("artist: %+v", importFailure.Artist)
	}

	if importFailure.Album != nil {
		t.Fatalf("expected nil album, got %+v", importFailure.Album)
	}
}

func TestReadRequestBodyTooLarge(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequestWithContext(
		context.Background(), http.MethodPost, "/", bytes.NewReader(bytes.Repeat([]byte("a"), oversizedWebhookBody)))
	rec := httptest.NewRecorder()

	handler := &starrconnect.SonarrHandler{}
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("want 400 got %d", rec.Code)
	}
}
