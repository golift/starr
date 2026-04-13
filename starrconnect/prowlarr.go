package starrconnect

import (
	"encoding/json"
	"fmt"
	"time"
)

// ProwlarrRelease is indexer release metadata in a Prowlarr Grab webhook.
type ProwlarrRelease struct {
	ReleaseTitle string    `json:"releaseTitle"`
	Indexer      string    `json:"indexer"`
	Size         int64     `json:"size"` // may be empty/null
	Categories   []string  `json:"categories"`
	Genres       []string  `json:"genres"`
	IndexerFlags []string  `json:"indexerFlags"`
	PublishDate  time.Time `json:"publishDate"` // may be empty/null
}

// ProwlarrGrab is the Grab webhook payload.
type ProwlarrGrab struct {
	BaseEvent

	Release            *ProwlarrRelease `json:"release"`
	Trigger            string           `json:"trigger"`
	Source             string           `json:"source"`
	Host               string           `json:"host"`
	DownloadClient     string           `json:"downloadClient"`
	DownloadClientType string           `json:"downloadClientType"`
	DownloadID         string           `json:"downloadId"`
}

// ProwlarrTest is the Test webhook payload (base fields only).
type ProwlarrTest struct {
	BaseEvent
}

// ProwlarrDownload is reserved for a future Download webhook shape (WebhookEventType only today).
type ProwlarrDownload struct {
	BaseEvent
}

// ProwlarrRename is reserved for a future Rename webhook shape (WebhookEventType only today).
type ProwlarrRename struct {
	BaseEvent
}

// ProwlarrHealth is the Health or HealthRestored webhook payload.
type ProwlarrHealth struct {
	BaseEvent

	Level   string `json:"level"`
	Message string `json:"message"`
	Type    string `json:"type"`
	WikiURL string `json:"wikiUrl"`
}

// ProwlarrApplicationUpdate is the ApplicationUpdate webhook payload.
type ProwlarrApplicationUpdate struct {
	BaseEvent

	Message         string `json:"message"`
	PreviousVersion string `json:"previousVersion"`
	NewVersion      string `json:"newVersion"`
}

// ProwlarrEvent is a parsed Prowlarr webhook envelope plus the raw JSON body.
type ProwlarrEvent struct {
	BaseEvent

	body []byte
}

// ParseProwlarr parses the raw JSON body and returns the envelope; use Get* to decode the full payload.
func ParseProwlarr(body []byte) (*ProwlarrEvent, error) {
	var base BaseEvent
	if err := json.Unmarshal(body, &base); err != nil {
		return nil, fmt.Errorf("decoding Prowlarr event envelope: %w", err)
	}

	return &ProwlarrEvent{BaseEvent: base, body: body}, nil
}

// GetGrab decodes a Grab payload.
func (e *ProwlarrEvent) GetGrab() (*ProwlarrGrab, error) {
	return decodeWebhookPayload[ProwlarrGrab](e.body, e.EventType, EventGrab)
}

// GetTest decodes a Test payload.
func (e *ProwlarrEvent) GetTest() (*ProwlarrTest, error) {
	return decodeWebhookPayload[ProwlarrTest](e.body, e.EventType, EventTest)
}

// GetDownload decodes a Download payload (reserved; upstream may add fields later).
func (e *ProwlarrEvent) GetDownload() (*ProwlarrDownload, error) {
	return decodeWebhookPayload[ProwlarrDownload](e.body, e.EventType, EventDownload)
}

// GetRename decodes a Rename payload (reserved; upstream may add fields later).
func (e *ProwlarrEvent) GetRename() (*ProwlarrRename, error) {
	return decodeWebhookPayload[ProwlarrRename](e.body, e.EventType, EventRename)
}

// GetHealth decodes a Health payload.
func (e *ProwlarrEvent) GetHealth() (*ProwlarrHealth, error) {
	return decodeWebhookPayload[ProwlarrHealth](e.body, e.EventType, EventHealth)
}

// GetHealthRestored decodes a HealthRestored payload.
func (e *ProwlarrEvent) GetHealthRestored() (*ProwlarrHealth, error) {
	return decodeWebhookPayload[ProwlarrHealth](e.body, e.EventType, EventHealthRestored)
}

// GetApplicationUpdate decodes an ApplicationUpdate payload.
func (e *ProwlarrEvent) GetApplicationUpdate() (*ProwlarrApplicationUpdate, error) {
	return decodeWebhookPayload[ProwlarrApplicationUpdate](e.body, e.EventType, EventApplicationUpdate)
}
