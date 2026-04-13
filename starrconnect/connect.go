// Package starrconnect unmarshals HTTP webhook JSON payloads from Sonarr, Radarr, Lidarr, and Prowlarr.
// Configure webhooks in each app under Settings → Connect → Webhook.
//
// For Custom Script (environment variables) instead of HTTP webhooks, see package starrcmd.
package starrconnect

/* Notes to future developers of this module:
- IDs should be int64 where they may exceed int32; C# int maps to Go int where safe.
- Sizes are int64 (bytes).
- JSON uses camelCase property names (Servarr Newtonsoft / System.Text.Json settings).
- Sonarr sends eventType "Download" for both single-file import and import-complete; use
  episodeFile vs episodeFiles in the JSON to tell them apart (see sonarr.go).
- Lidarr sends eventType "ImportFailure" with the same shape as "Download" but album is null.
- Health payload "level" is a string enum: ok, notice, warning, error.
- Prowlarr currently emits Test, Grab, Health, HealthRestored, and ApplicationUpdate; Download and
  Rename exist on WebhookEventType but have no dedicated payload types in upstream yet.
- Add fixtures from real apps when extending types; upstream DTOs live under
  NzbDrone.Core/Notifications/Webhook/ in each app's repo.
*/

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const maxBodyBytes = 10 << 20 // 10 MiB

var (
	// ErrWrongEvent is returned by Get* methods when the envelope eventType does not match.
	ErrWrongEvent = errors.New("starrconnect: event type mismatch")
	// ErrUnknownEvent is returned when the JSON eventType is not recognized for this app.
	ErrUnknownEvent = errors.New("starrconnect: unknown event type")
	// ErrBodyTooLarge is returned when the webhook POST exceeds maxBodyBytes.
	ErrBodyTooLarge = errors.New("starrconnect: request body too large")
)

// EventType is the webhook eventType field (JSON string).
type EventType string

// EventType constants.
const (
	EventTest                      EventType = "Test"
	EventGrab                      EventType = "Grab"
	EventDownload                  EventType = "Download"
	EventRename                    EventType = "Rename"
	EventHealth                    EventType = "Health"
	EventHealthRestored            EventType = "HealthRestored"
	EventApplicationUpdate         EventType = "ApplicationUpdate"
	EventManualInteractionRequired EventType = "ManualInteractionRequired"
	EventSeriesAdd                 EventType = "SeriesAdd"
	EventSeriesDelete              EventType = "SeriesDelete"
	EventEpisodeFileDelete         EventType = "EpisodeFileDelete"
	EventMovieAdded                EventType = "MovieAdded"
	EventMovieDelete               EventType = "MovieDelete"
	EventMovieFileDelete           EventType = "MovieFileDelete"
	EventDownloadFailure           EventType = "DownloadFailure"
	EventImportFailure             EventType = "ImportFailure"
	EventArtistAdd                 EventType = "ArtistAdd"
	EventArtistDelete              EventType = "ArtistDelete"
	EventAlbumDelete               EventType = "AlbumDelete"
	EventRetag                     EventType = "Retag"
)

// BaseEvent is embedded in every concrete webhook payload.
type BaseEvent struct {
	EventType      EventType `json:"eventType"`
	InstanceName   string    `json:"instanceName"`
	ApplicationURL string    `json:"applicationUrl"`
}

// Image is a cover / media image reference in webhook payloads.
type Image struct {
	CoverType string `json:"coverType"`
	URL       string `json:"url"`
	RemoteURL string `json:"remoteUrl"`
}

// Language is an audio/subtitle language entry (id + name).
type Language struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// CustomFormat is a scored custom format on a release.
type CustomFormat struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// CustomFormatInfo groups custom formats and score (Sonarr/Radarr).
type CustomFormatInfo struct {
	CustomFormats     []CustomFormat `json:"customFormats"`
	CustomFormatScore int            `json:"customFormatScore"`
}

// DownloadStatusMessage is a tracked-download status message block.
type DownloadStatusMessage struct {
	Title    string   `json:"title"`
	Messages []string `json:"messages"`
}

// DownloadClientItem describes the download client queue item (e.g. manual interaction).
type DownloadClientItem struct {
	Quality        string `json:"quality"`
	QualityVersion int    `json:"qualityVersion"`
	Title          string `json:"title"`
	Indexer        string `json:"indexer"`
	Size           int64  `json:"size"`
}

func readRequestBody(httpReq *http.Request) ([]byte, error) {
	if httpReq.Body == nil {
		return nil, fmt.Errorf("nil body: %w", io.ErrUnexpectedEOF)
	}

	defer httpReq.Body.Close()

	data, err := io.ReadAll(io.LimitReader(httpReq.Body, maxBodyBytes+1))
	if err != nil {
		return nil, fmt.Errorf("read body: %w", err)
	}

	if int64(len(data)) > maxBodyBytes {
		return nil, fmt.Errorf("%w", ErrBodyTooLarge)
	}

	return data, nil
}

// decodeWebhookPayload checks got == want, then unmarshals body into *T (Servarr Get* helpers).
func decodeWebhookPayload[T any](got, want EventType, body []byte, decodeLabel string) (*T, error) {
	if got != want {
		return nil, fmt.Errorf("%w: got %q want %q", ErrWrongEvent, got, want)
	}

	var out T

	if err := json.Unmarshal(body, &out); err != nil {
		return nil, fmt.Errorf("%s: %w", decodeLabel, err)
	}

	return &out, nil
}

// decodeWebhookPayloadEither checks got is a or b, then unmarshals body into *T (e.g. Grab + Test).
func decodeWebhookPayloadEither[T any](got, a, b EventType, body []byte, decodeLabel string) (*T, error) {
	if got != a && got != b { // Usually "Grab" or "Test"
		return nil, fmt.Errorf("%w: got %q want %q or %q", ErrWrongEvent, got, a, b)
	}

	var out T

	if err := json.Unmarshal(body, &out); err != nil {
		return nil, fmt.Errorf("%s: %w", decodeLabel, err)
	}

	return &out, nil
}

// webhookHTTPError attaches an HTTP status and client-facing message to an error chain.
type webhookHTTPError struct {
	Status  int
	Message string
	Cause   error
}

func (e *webhookHTTPError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Cause)
	}

	return e.Message
}

func (e *webhookHTTPError) Unwrap() error {
	return e.Cause
}

func webhookErr(status int, public string, cause error) error {
	return &webhookHTTPError{Status: status, Message: public, Cause: cause}
}

func writeWebhookHTTPError(responseWriter http.ResponseWriter, log func(error), err error) {
	var httpErr *webhookHTTPError
	if errors.As(err, &httpErr) {
		if log != nil && httpErr.Cause != nil {
			log(httpErr.Cause)
		}

		http.Error(responseWriter, httpErr.Message, httpErr.Status)

		return
	}

	if log != nil {
		log(err)
	}

	http.Error(responseWriter, "handler error", http.StatusInternalServerError)
}

// runWebhookCallback runs an optional user callback after a successful typed decode.
func runWebhookCallback[T any](callback func(*T) error, decode func() (*T, error)) error {
	if callback == nil {
		return nil
	}

	v, err := decode()
	if err != nil {
		return webhookErr(http.StatusInternalServerError, "handler error", err)
	}

	if err := callback(v); err != nil {
		return webhookErr(http.StatusInternalServerError, "handler error", err)
	}

	return nil
}
