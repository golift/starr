package starrconnect

import (
	"encoding/json"
	"fmt"
	"time"
)

// Artist is artist metadata in a Lidarr webhook.
type Artist struct {
	ID             int64    `json:"id"`
	Name           string   `json:"name"`
	Disambiguation string   `json:"disambiguation"`
	Path           string   `json:"path"`
	MBID           string   `json:"mbId"`
	Type           string   `json:"type"`
	Overview       string   `json:"overview"`
	Genres         []string `json:"genres"`
	Images         []Image  `json:"images"`
	Tags           []string `json:"tags"`
}

// Album is album metadata in a Lidarr webhook.
type Album struct {
	ID                  int64      `json:"id"`
	MBID                string     `json:"mbId"`
	Title               string     `json:"title"`
	Disambiguation      string     `json:"disambiguation"`
	Overview            string     `json:"overview"`
	AlbumType           string     `json:"albumType"`
	SecondaryAlbumTypes []string   `json:"secondaryAlbumTypes"`
	ReleaseDate         *time.Time `json:"releaseDate"`
	Genres              []string   `json:"genres"`
	Images              []Image    `json:"images"`
}

// Track is track metadata in a Lidarr webhook.
type Track struct {
	ID             int    `json:"id"`
	Title          string `json:"title"`
	TrackNumber    string `json:"trackNumber"`
	Quality        string `json:"quality"`
	QualityVersion int    `json:"qualityVersion"`
	ReleaseGroup   string `json:"releaseGroup"`
}

// TrackFile is an on-disk track file in a Lidarr webhook.
type TrackFile struct {
	ID             int64     `json:"id"`
	Path           string    `json:"path"`
	Quality        string    `json:"quality"`
	QualityVersion int       `json:"qualityVersion"`
	ReleaseGroup   string    `json:"releaseGroup"`
	SceneName      string    `json:"sceneName"`
	Size           int64     `json:"size"`
	DateAdded      time.Time `json:"dateAdded"`
}

// RenamedTrackFile extends LidarrTrackFile with the previous path.
type RenamedTrackFile struct {
	TrackFile

	PreviousPath string `json:"previousPath"`
}

// LidarrRelease is pre-grab release info (Grab / Test).
type LidarrRelease struct {
	Quality           string   `json:"quality"`
	QualityVersion    int      `json:"qualityVersion"`
	ReleaseGroup      string   `json:"releaseGroup"`
	ReleaseTitle      string   `json:"releaseTitle"`
	Indexer           string   `json:"indexer"`
	Size              int64    `json:"size"`
	CustomFormatScore int      `json:"customFormatScore"`
	CustomFormats     []string `json:"customFormats"`
}

// --- Per-event payloads ---

// LidarrGrab is the Grab (and Test) webhook payload.
type LidarrGrab struct {
	BaseEvent

	Artist             *Artist        `json:"artist"`
	Albums             []Album        `json:"albums"`
	Release            *LidarrRelease `json:"release"`
	DownloadClient     string         `json:"downloadClient"`
	DownloadClientType string         `json:"downloadClientType"`
	DownloadID         string         `json:"downloadId"`
}

// LidarrDownload is the Download webhook payload (successful import).
type LidarrDownload struct {
	BaseEvent

	Artist             *Artist     `json:"artist"`
	Album              *Album      `json:"album"`
	Tracks             []Track     `json:"tracks"`
	TrackFiles         []TrackFile `json:"trackFiles"`
	DeletedFiles       []TrackFile `json:"deletedFiles"`
	IsUpgrade          bool        `json:"isUpgrade"`
	DownloadClient     string      `json:"downloadClient"`
	DownloadClientType string      `json:"downloadClientType"`
	DownloadID         string      `json:"downloadId"`
}

// LidarrImportFailure is the ImportFailure payload; same JSON shape as LidarrDownload (album may be null).
type LidarrImportFailure = LidarrDownload

// LidarrDownloadFailure is the DownloadFailure webhook payload.
type LidarrDownloadFailure struct {
	BaseEvent

	Quality        string `json:"quality"`
	QualityVersion int    `json:"qualityVersion"`
	ReleaseTitle   string `json:"releaseTitle"`
	DownloadClient string `json:"downloadClient"`
	DownloadID     string `json:"downloadId"`
}

// LidarrRename is the Rename webhook payload.
type LidarrRename struct {
	BaseEvent

	Artist            *Artist            `json:"artist"`
	RenamedTrackFiles []RenamedTrackFile `json:"renamedTrackFiles"`
}

// LidarrRetag is the Retag webhook payload.
type LidarrRetag struct {
	BaseEvent

	Artist    *Artist    `json:"artist"`
	TrackFile *TrackFile `json:"trackFile"`
}

// ArtistAdd is the ArtistAdd webhook payload.
type ArtistAdd struct {
	BaseEvent

	Artist *Artist `json:"artist"`
}

// ArtistDelete is the ArtistDelete webhook payload.
type ArtistDelete struct {
	BaseEvent

	Artist       *Artist `json:"artist"`
	DeletedFiles bool    `json:"deletedFiles"`
}

// AlbumDelete is the AlbumDelete webhook payload.
type AlbumDelete struct {
	BaseEvent

	Artist       *Artist `json:"artist"`
	Album        *Album  `json:"album"`
	DeletedFiles bool    `json:"deletedFiles"`
}

// LidarrHealth is the Health or HealthRestored webhook payload.
type LidarrHealth struct {
	BaseEvent

	Level   string `json:"level"`
	Message string `json:"message"`
	Type    string `json:"type"`
	WikiURL string `json:"wikiUrl"`
}

// LidarrApplicationUpdate is the ApplicationUpdate webhook payload.
type LidarrApplicationUpdate struct {
	BaseEvent

	Message         string `json:"message"`
	PreviousVersion string `json:"previousVersion"`
	NewVersion      string `json:"newVersion"`
}

// LidarrEvent is a parsed Lidarr webhook envelope plus the raw JSON body.
type LidarrEvent struct {
	BaseEvent

	body []byte
}

// ParseLidarr parses the raw JSON body and returns the envelope; use Get* to decode the full payload.
func ParseLidarr(body []byte) (*LidarrEvent, error) {
	var base BaseEvent
	if err := json.Unmarshal(body, &base); err != nil {
		return nil, fmt.Errorf("decoding Lidarr event envelope: %w", err)
	}

	return &LidarrEvent{BaseEvent: base, body: body}, nil
}

// GetGrab decodes a Grab or Test payload.
func (e *LidarrEvent) GetGrab() (*LidarrGrab, error) {
	return decodeWebhookPayloadEither[LidarrGrab](e.EventType, EventGrab, EventTest, e.body, "decoding LidarrGrab")
}

// GetDownload decodes a Download payload.
func (e *LidarrEvent) GetDownload() (*LidarrDownload, error) {
	return decodeWebhookPayload[LidarrDownload](e.EventType, EventDownload, e.body, "decoding LidarrDownload")
}

// GetImportFailure decodes an ImportFailure payload.
func (e *LidarrEvent) GetImportFailure() (*LidarrImportFailure, error) {
	return decodeWebhookPayload[LidarrDownload](e.EventType, EventImportFailure, e.body, "decoding LidarrImportFailure")
}

// GetDownloadFailure decodes a DownloadFailure payload.
func (e *LidarrEvent) GetDownloadFailure() (*LidarrDownloadFailure, error) {
	return decodeWebhookPayload[LidarrDownloadFailure](
		e.EventType, EventDownloadFailure, e.body, "decoding LidarrDownloadFailure")
}

// GetRename decodes a Rename payload.
func (e *LidarrEvent) GetRename() (*LidarrRename, error) {
	return decodeWebhookPayload[LidarrRename](e.EventType, EventRename, e.body, "decoding LidarrRename")
}

// GetRetag decodes a Retag payload.
func (e *LidarrEvent) GetRetag() (*LidarrRetag, error) {
	return decodeWebhookPayload[LidarrRetag](e.EventType, EventRetag, e.body, "decoding LidarrRetag")
}

// GetArtistAdd decodes an ArtistAdd payload.
func (e *LidarrEvent) GetArtistAdd() (*ArtistAdd, error) {
	return decodeWebhookPayload[ArtistAdd](e.EventType, EventArtistAdd, e.body, "decoding LidarrArtistAdd")
}

// GetArtistDelete decodes an ArtistDelete payload.
func (e *LidarrEvent) GetArtistDelete() (*ArtistDelete, error) {
	return decodeWebhookPayload[ArtistDelete](e.EventType, EventArtistDelete, e.body, "decoding LidarrArtistDelete")
}

// GetAlbumDelete decodes an AlbumDelete payload.
func (e *LidarrEvent) GetAlbumDelete() (*AlbumDelete, error) {
	return decodeWebhookPayload[AlbumDelete](e.EventType, EventAlbumDelete, e.body, "decoding LidarrAlbumDelete")
}

// GetHealth decodes a Health payload.
func (e *LidarrEvent) GetHealth() (*LidarrHealth, error) {
	return decodeWebhookPayload[LidarrHealth](e.EventType, EventHealth, e.body, "decoding LidarrHealth")
}

// GetHealthRestored decodes a HealthRestored payload.
func (e *LidarrEvent) GetHealthRestored() (*LidarrHealth, error) {
	return decodeWebhookPayload[LidarrHealth](e.EventType, EventHealthRestored, e.body, "decoding LidarrHealth")
}

// GetApplicationUpdate decodes an ApplicationUpdate payload.
func (e *LidarrEvent) GetApplicationUpdate() (*LidarrApplicationUpdate, error) {
	return decodeWebhookPayload[LidarrApplicationUpdate](
		e.EventType, EventApplicationUpdate, e.body, "decoding LidarrApplicationUpdate")
}
