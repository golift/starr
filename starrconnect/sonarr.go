package starrconnect

import (
	"encoding/json"
	"fmt"
	"time"
)

// Series is series metadata in a Sonarr webhook.
type Series struct {
	ID               int64     `json:"id"`
	Title            string    `json:"title"`
	TitleSlug        string    `json:"titleSlug"`
	Path             string    `json:"path"`
	TvdbID           int64     `json:"tvdbId"`
	TvMazeID         int64     `json:"tvMazeId"`
	TmdbID           int64     `json:"tmdbId"`
	ImdbID           string    `json:"imdbId"`
	MalIDs           []int64   `json:"malIds"`
	AniListIDs       []int64   `json:"aniListIds"`
	Type             string    `json:"type"`
	Year             int       `json:"year"`
	Genres           []string  `json:"genres"`
	Images           []Image   `json:"images"`
	Tags             []string  `json:"tags"`
	OriginalLanguage *Language `json:"originalLanguage"`
	OriginalCountry  string    `json:"originalCountry"`
}

// Episode is episode metadata in a Sonarr webhook.
type Episode struct {
	ID            int64      `json:"id"`
	EpisodeNumber int        `json:"episodeNumber"`
	SeasonNumber  int        `json:"seasonNumber"`
	Title         string     `json:"title"`
	Overview      string     `json:"overview"`
	AirDate       string     `json:"airDate"`
	AirDateUtc    *time.Time `json:"airDateUtc"`
	SeriesID      int64      `json:"seriesId"`
	TvdbID        int64      `json:"tvdbId"`
	FinaleType    string     `json:"finaleType"`
}

// EpisodeFileMediaInfo is optional media info on an episode file.
type EpisodeFileMediaInfo struct {
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

// EpisodeFile is an on-disk episode file in a Sonarr webhook.
type EpisodeFile struct {
	ID             int64                 `json:"id"`
	RelativePath   string                `json:"relativePath"`
	Path           string                `json:"path"`
	Quality        string                `json:"quality"`
	QualityVersion int                   `json:"qualityVersion"`
	ReleaseGroup   string                `json:"releaseGroup"`
	SceneName      string                `json:"sceneName"`
	Size           int64                 `json:"size"`
	DateAdded      time.Time             `json:"dateAdded"`
	Languages      []Language            `json:"languages"`
	MediaInfo      *EpisodeFileMediaInfo `json:"mediaInfo"`
	SourcePath     string                `json:"sourcePath"`
	RecycleBinPath string                `json:"recycleBinPath"`
}

// RenamedEpisodeFile extends SonarrEpisodeFile with previous paths.
type RenamedEpisodeFile struct {
	EpisodeFile

	PreviousRelativePath string `json:"previousRelativePath"`
	PreviousPath         string `json:"previousPath"`
}

// SonarrRelease is pre-grab release info (Grab / Test).
type SonarrRelease struct {
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

// SonarrGrabbedRelease is post-grab release info (Download / ManualInteraction / ImportComplete).
type SonarrGrabbedRelease struct {
	ReleaseTitle string   `json:"releaseTitle"`
	Indexer      string   `json:"indexer"`
	Size         *int64   `json:"size"`
	IndexerFlags []string `json:"indexerFlags"`
	ReleaseType  string   `json:"releaseType"`
}

// --- Per-event payloads ---

// SonarrGrab is the Grab (and Test) webhook payload.
type SonarrGrab struct {
	BaseEvent

	Series             *Series           `json:"series"`
	Episodes           []Episode         `json:"episodes"`
	Release            *SonarrRelease    `json:"release"`
	DownloadClient     string            `json:"downloadClient"`
	DownloadClientType string            `json:"downloadClientType"`
	DownloadID         string            `json:"downloadId"`
	CustomFormatInfo   *CustomFormatInfo `json:"customFormatInfo"`
}

// SonarrDownload is the Download webhook for a single imported episode file.
type SonarrDownload struct {
	BaseEvent

	Series             *Series               `json:"series"`
	Episodes           []Episode             `json:"episodes"`
	EpisodeFile        *EpisodeFile          `json:"episodeFile"`
	IsUpgrade          bool                  `json:"isUpgrade"`
	DownloadClient     string                `json:"downloadClient"`
	DownloadClientType string                `json:"downloadClientType"`
	DownloadID         string                `json:"downloadId"`
	DeletedFiles       []EpisodeFile         `json:"deletedFiles"`
	CustomFormatInfo   *CustomFormatInfo     `json:"customFormatInfo"`
	Release            *SonarrGrabbedRelease `json:"release"`
}

// SonarrImportComplete is the Download webhook when a batch import completes (episodeFiles set).
type SonarrImportComplete struct {
	BaseEvent

	Series             *Series               `json:"series"`
	Episodes           []Episode             `json:"episodes"`
	EpisodeFiles       []EpisodeFile         `json:"episodeFiles"`
	DownloadClient     string                `json:"downloadClient"`
	DownloadClientType string                `json:"downloadClientType"`
	DownloadID         string                `json:"downloadId"`
	Release            *SonarrGrabbedRelease `json:"release"`
	FileCount          int                   `json:"fileCount"`
	SourcePath         string                `json:"sourcePath"`
	DestinationPath    string                `json:"destinationPath"`
}

// SonarrRename is the Rename webhook payload.
type SonarrRename struct {
	BaseEvent

	Series              *Series              `json:"series"`
	RenamedEpisodeFiles []RenamedEpisodeFile `json:"renamedEpisodeFiles"`
}

// SeriesAdd is the SeriesAdd webhook payload.
type SeriesAdd struct {
	BaseEvent

	Series *Series `json:"series"`
}

// SeriesDelete is the SeriesDelete webhook payload.
type SeriesDelete struct {
	BaseEvent

	Series       *Series `json:"series"`
	DeletedFiles bool    `json:"deletedFiles"`
}

// EpisodeFileDelete is the EpisodeFileDelete webhook payload.
type EpisodeFileDelete struct {
	BaseEvent

	Series       *Series      `json:"series"`
	Episodes     []Episode    `json:"episodes"`
	EpisodeFile  *EpisodeFile `json:"episodeFile"`
	DeleteReason string       `json:"deleteReason"`
}

// SonarrHealth is the Health or HealthRestored webhook payload.
type SonarrHealth struct {
	BaseEvent

	Level   string `json:"level"`
	Message string `json:"message"`
	Type    string `json:"type"`
	WikiURL string `json:"wikiUrl"`
}

// SonarrApplicationUpdate is the ApplicationUpdate webhook payload.
type SonarrApplicationUpdate struct {
	BaseEvent

	Message         string `json:"message"`
	PreviousVersion string `json:"previousVersion"`
	NewVersion      string `json:"newVersion"`
}

// SonarrManualInteraction is the ManualInteractionRequired webhook payload.
type SonarrManualInteraction struct {
	BaseEvent

	Series                 *Series                 `json:"series"`
	Episodes               []Episode               `json:"episodes"`
	DownloadInfo           *DownloadClientItem     `json:"downloadInfo"`
	DownloadClient         string                  `json:"downloadClient"`
	DownloadClientType     string                  `json:"downloadClientType"`
	DownloadID             string                  `json:"downloadId"`
	DownloadStatus         string                  `json:"downloadStatus"`
	DownloadStatusMessages []DownloadStatusMessage `json:"downloadStatusMessages"`
	CustomFormatInfo       *CustomFormatInfo       `json:"customFormatInfo"`
	Release                *SonarrGrabbedRelease   `json:"release"`
}

// SonarrEvent is a parsed Sonarr webhook envelope plus the raw JSON body.
type SonarrEvent struct {
	BaseEvent

	body []byte
}

// ParseSonarr parses the raw JSON body and returns the envelope; use Get* to decode the full payload.
func ParseSonarr(body []byte) (*SonarrEvent, error) {
	var base BaseEvent
	if err := json.Unmarshal(body, &base); err != nil {
		return nil, fmt.Errorf("decoding Sonarr event envelope: %w", err)
	}

	return &SonarrEvent{BaseEvent: base, body: body}, nil
}

// GetGrab decodes a Grab or Test payload (Test uses the same shape as Grab).
func (e *SonarrEvent) GetGrab() (*SonarrGrab, error) {
	return decodeWebhookPayload[SonarrGrab](e.body, e.EventType, EventGrab, EventTest)
}

// GetDownload decodes a single-file Download payload (not import-complete batch).
func (e *SonarrEvent) GetDownload() (*SonarrDownload, error) {
	if e.EventType != EventDownload {
		return nil, fmt.Errorf("%w: got %q want %q", ErrWrongEvent, e.EventType, EventDownload)
	}

	if sonarrIsImportCompleteBody(e.body) {
		return nil, fmt.Errorf("%w: payload is import-complete (episodeFiles), use GetImportComplete", ErrWrongEvent)
	}

	return decodeWebhookPayload[SonarrDownload](e.body, e.EventType, EventDownload)
}

// GetImportComplete decodes a batch Download (import complete) payload.
func (e *SonarrEvent) GetImportComplete() (*SonarrImportComplete, error) {
	if e.EventType != EventDownload {
		return nil, fmt.Errorf("%w: got %q want %q", ErrWrongEvent, e.EventType, EventDownload)
	}

	if !sonarrIsImportCompleteBody(e.body) {
		return nil, fmt.Errorf("%w: payload is single-file import (episodeFile), use GetDownload", ErrWrongEvent)
	}

	return decodeWebhookPayload[SonarrImportComplete](e.body, e.EventType, EventDownload)
}

// GetRename decodes a Rename payload.
func (e *SonarrEvent) GetRename() (*SonarrRename, error) {
	return decodeWebhookPayload[SonarrRename](e.body, e.EventType, EventRename)
}

// GetSeriesAdd decodes a SeriesAdd payload.
func (e *SonarrEvent) GetSeriesAdd() (*SeriesAdd, error) {
	return decodeWebhookPayload[SeriesAdd](e.body, e.EventType, EventSeriesAdd)
}

// GetSeriesDelete decodes a SeriesDelete payload.
func (e *SonarrEvent) GetSeriesDelete() (*SeriesDelete, error) {
	return decodeWebhookPayload[SeriesDelete](e.body, e.EventType, EventSeriesDelete)
}

// GetEpisodeFileDelete decodes an EpisodeFileDelete payload.
func (e *SonarrEvent) GetEpisodeFileDelete() (*EpisodeFileDelete, error) {
	return decodeWebhookPayload[EpisodeFileDelete](e.body, e.EventType, EventEpisodeFileDelete)
}

// GetHealth decodes a Health payload.
func (e *SonarrEvent) GetHealth() (*SonarrHealth, error) {
	return decodeWebhookPayload[SonarrHealth](e.body, e.EventType, EventHealth)
}

// GetHealthRestored decodes a HealthRestored payload.
func (e *SonarrEvent) GetHealthRestored() (*SonarrHealth, error) {
	return decodeWebhookPayload[SonarrHealth](e.body, e.EventType, EventHealthRestored)
}

// GetApplicationUpdate decodes an ApplicationUpdate payload.
func (e *SonarrEvent) GetApplicationUpdate() (*SonarrApplicationUpdate, error) {
	return decodeWebhookPayload[SonarrApplicationUpdate](e.body, e.EventType, EventApplicationUpdate)
}

// GetManualInteraction decodes a ManualInteractionRequired payload.
func (e *SonarrEvent) GetManualInteraction() (*SonarrManualInteraction, error) {
	return decodeWebhookPayload[SonarrManualInteraction](e.body, e.EventType, EventManualInteractionRequired)
}
