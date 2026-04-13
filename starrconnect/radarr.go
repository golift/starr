package starrconnect

import (
	"encoding/json"
	"fmt"
	"time"
)

// Movie is movie metadata in a Radarr webhook.
type Movie struct {
	ID               int64     `json:"id"`
	Title            string    `json:"title"`
	Year             int       `json:"year"`
	FilePath         string    `json:"filePath"`
	ReleaseDate      string    `json:"releaseDate"`
	FolderPath       string    `json:"folderPath"`
	TmdbID           int64     `json:"tmdbId"`
	ImdbID           string    `json:"imdbId"`
	Overview         string    `json:"overview"`
	Genres           []string  `json:"genres"`
	Images           []Image   `json:"images"`
	Tags             []string  `json:"tags"`
	OriginalLanguage *Language `json:"originalLanguage"`
}

// RemoteMovie is a lightweight movie reference from a grab/release.
type RemoteMovie struct {
	TmdbID int64  `json:"tmdbId"`
	ImdbID string `json:"imdbId"`
	Title  string `json:"title"`
	Year   int    `json:"year"`
}

// MovieFileMediaInfo is optional media info on a movie file.
type MovieFileMediaInfo struct {
	AudioChannels         float64  `json:"audioChannels"`
	AudioCodec            string   `json:"audioCodec"`
	AudioLanguages        []string `json:"audioLanguages"`
	Height                int      `json:"height"`
	Width                 int      `json:"width"`
	Subtitles             []string `json:"subtitles"`
	VideoCodec            string   `json:"videoCodec"`
	VideoDynamicRange     string   `json:"videoDynamicRange"`
	VideoDynamicRangeType string   `json:"videoDynamicRangeType"`
}

// MovieFile is an on-disk movie file in a Radarr webhook.
type MovieFile struct {
	ID             int64               `json:"id"`
	RelativePath   string              `json:"relativePath"`
	Path           string              `json:"path"`
	Quality        string              `json:"quality"`
	QualityVersion int                 `json:"qualityVersion"`
	ReleaseGroup   string              `json:"releaseGroup"`
	SceneName      string              `json:"sceneName"`
	IndexerFlags   string              `json:"indexerFlags"`
	Size           int64               `json:"size"`
	DateAdded      time.Time           `json:"dateAdded"`
	Languages      []Language          `json:"languages"`
	MediaInfo      *MovieFileMediaInfo `json:"mediaInfo"`
	SourcePath     string              `json:"sourcePath"`
	RecycleBinPath string              `json:"recycleBinPath"`
}

// RenamedMovieFile extends RadarrMovieFile with previous paths.
type RenamedMovieFile struct {
	MovieFile

	PreviousRelativePath string `json:"previousRelativePath"`
	PreviousPath         string `json:"previousPath"`
}

// RadarrRelease is pre-grab release info (Grab / Test).
type RadarrRelease struct {
	Quality           string     `json:"quality"`
	QualityVersion    int        `json:"qualityVersion"`
	ReleaseGroup      string     `json:"releaseGroup"`
	ReleaseTitle      string     `json:"releaseTitle"`
	Indexer           string     `json:"indexer"`
	Size              int64      `json:"size"`
	CustomFormatScore int        `json:"customFormatScore"`
	CustomFormats     []string   `json:"customFormats"`
	Languages         []Language `json:"languages"`
	IndexerFlags      []string   `json:"indexerFlags"`
}

// RadarrGrabbedRelease is post-grab release info (Download / ManualInteraction).
type RadarrGrabbedRelease struct {
	ReleaseTitle string   `json:"releaseTitle"`
	Indexer      string   `json:"indexer"`
	Size         int64    `json:"size"`
	IndexerFlags []string `json:"indexerFlags"`
}

// --- Per-event payloads ---

// RadarrGrab is the Grab (and Test) webhook payload.
type RadarrGrab struct {
	BaseEvent

	Movie              *Movie            `json:"movie"`
	RemoteMovie        *RemoteMovie      `json:"remoteMovie"`
	Release            *RadarrRelease    `json:"release"`
	DownloadClient     string            `json:"downloadClient"`
	DownloadClientType string            `json:"downloadClientType"`
	DownloadID         string            `json:"downloadId"`
	CustomFormatInfo   *CustomFormatInfo `json:"customFormatInfo"`
}

// RadarrDownload is the Download webhook payload.
type RadarrDownload struct {
	BaseEvent

	Movie              *Movie                `json:"movie"`
	RemoteMovie        *RemoteMovie          `json:"remoteMovie"`
	MovieFile          *MovieFile            `json:"movieFile"`
	IsUpgrade          bool                  `json:"isUpgrade"`
	DownloadClient     string                `json:"downloadClient"`
	DownloadClientType string                `json:"downloadClientType"`
	DownloadID         string                `json:"downloadId"`
	DeletedFiles       []MovieFile           `json:"deletedFiles"`
	CustomFormatInfo   *CustomFormatInfo     `json:"customFormatInfo"`
	Release            *RadarrGrabbedRelease `json:"release"`
}

// RadarrRename is the Rename webhook payload.
type RadarrRename struct {
	BaseEvent

	Movie             *Movie             `json:"movie"`
	RenamedMovieFiles []RenamedMovieFile `json:"renamedMovieFiles"`
}

// MovieAdded is the MovieAdded webhook payload.
type MovieAdded struct {
	BaseEvent

	Movie     *Movie `json:"movie"`
	AddMethod string `json:"addMethod"`
}

// MovieDelete is the MovieDelete webhook payload.
type MovieDelete struct {
	BaseEvent

	Movie           *Movie `json:"movie"`
	DeletedFiles    bool   `json:"deletedFiles"`
	MovieFolderSize int64  `json:"movieFolderSize"`
}

// MovieFileDelete is the MovieFileDelete webhook payload.
type MovieFileDelete struct {
	BaseEvent

	Movie        *Movie     `json:"movie"`
	MovieFile    *MovieFile `json:"movieFile"`
	DeleteReason string     `json:"deleteReason"`
}

// RadarrHealth is the Health or HealthRestored webhook payload.
type RadarrHealth struct {
	BaseEvent

	Level   string `json:"level"`
	Message string `json:"message"`
	Type    string `json:"type"`
	WikiURL string `json:"wikiUrl"`
}

// RadarrApplicationUpdate is the ApplicationUpdate webhook payload.
type RadarrApplicationUpdate struct {
	BaseEvent

	Message         string `json:"message"`
	PreviousVersion string `json:"previousVersion"`
	NewVersion      string `json:"newVersion"`
}

// RadarrManualInteraction is the ManualInteractionRequired webhook payload.
type RadarrManualInteraction struct {
	BaseEvent

	Movie                  *Movie                  `json:"movie"`
	DownloadInfo           *DownloadClientItem     `json:"downloadInfo"`
	DownloadClient         string                  `json:"downloadClient"`
	DownloadClientType     string                  `json:"downloadClientType"`
	DownloadID             string                  `json:"downloadId"`
	DownloadStatus         string                  `json:"downloadStatus"`
	DownloadStatusMessages []DownloadStatusMessage `json:"downloadStatusMessages"`
	CustomFormatInfo       *CustomFormatInfo       `json:"customFormatInfo"`
	Release                *RadarrGrabbedRelease   `json:"release"`
}

// RadarrEvent is a parsed Radarr webhook envelope plus the raw JSON body.
type RadarrEvent struct {
	BaseEvent

	body []byte
}

// ParseRadarr parses the raw JSON body and returns the envelope; use Get* to decode the full payload.
func ParseRadarr(body []byte) (*RadarrEvent, error) {
	var base BaseEvent
	if err := json.Unmarshal(body, &base); err != nil {
		return nil, fmt.Errorf("decoding Radarr event envelope: %w", err)
	}

	return &RadarrEvent{BaseEvent: base, body: body}, nil
}

// GetGrab decodes a Grab or Test payload.
func (e *RadarrEvent) GetGrab() (*RadarrGrab, error) {
	return decodeWebhookPayloadEither[RadarrGrab](e.EventType, EventGrab, EventTest, e.body, "decoding RadarrGrab")
}

// GetDownload decodes a Download payload.
func (e *RadarrEvent) GetDownload() (*RadarrDownload, error) {
	return decodeWebhookPayload[RadarrDownload](e.EventType, EventDownload, e.body, "decoding RadarrDownload")
}

// GetRename decodes a Rename payload.
func (e *RadarrEvent) GetRename() (*RadarrRename, error) {
	return decodeWebhookPayload[RadarrRename](e.EventType, EventRename, e.body, "decoding RadarrRename")
}

// GetMovieAdded decodes a MovieAdded payload.
func (e *RadarrEvent) GetMovieAdded() (*MovieAdded, error) {
	return decodeWebhookPayload[MovieAdded](e.EventType, EventMovieAdded, e.body, "decoding RadarrMovieAdded")
}

// GetMovieDelete decodes a MovieDelete payload.
func (e *RadarrEvent) GetMovieDelete() (*MovieDelete, error) {
	return decodeWebhookPayload[MovieDelete](e.EventType, EventMovieDelete, e.body, "decoding RadarrMovieDelete")
}

// GetMovieFileDelete decodes a MovieFileDelete payload.
func (e *RadarrEvent) GetMovieFileDelete() (*MovieFileDelete, error) {
	return decodeWebhookPayload[MovieFileDelete](
		e.EventType, EventMovieFileDelete, e.body, "decoding RadarrMovieFileDelete")
}

// GetHealth decodes a Health payload.
func (e *RadarrEvent) GetHealth() (*RadarrHealth, error) {
	return decodeWebhookPayload[RadarrHealth](e.EventType, EventHealth, e.body, "decoding RadarrHealth")
}

// GetHealthRestored decodes a HealthRestored payload.
func (e *RadarrEvent) GetHealthRestored() (*RadarrHealth, error) {
	return decodeWebhookPayload[RadarrHealth](e.EventType, EventHealthRestored, e.body, "decoding RadarrHealth")
}

// GetApplicationUpdate decodes an ApplicationUpdate payload.
func (e *RadarrEvent) GetApplicationUpdate() (*RadarrApplicationUpdate, error) {
	return decodeWebhookPayload[RadarrApplicationUpdate](
		e.EventType, EventApplicationUpdate, e.body, "decoding RadarrApplicationUpdate")
}

// GetManualInteraction decodes a ManualInteractionRequired payload.
func (e *RadarrEvent) GetManualInteraction() (*RadarrManualInteraction, error) {
	return decodeWebhookPayload[RadarrManualInteraction](
		e.EventType, EventManualInteractionRequired, e.body, "decoding RadarrManualInteraction")
}
