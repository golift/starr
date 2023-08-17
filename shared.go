package starr

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"golift.io/starr/debuglog"
)

/* This file contains shared structs or constants for all the *arr apps. */

// App can be used to satisfy a context value key.
// It is not used in this library; provided for convenience.
type App string

// These constants are just here for convenience.
const (
	Emby     App = "Emby"
	Lidarr   App = "Lidarr"
	Plex     App = "Plex"
	Prowlarr App = "Prowlarr"
	Radarr   App = "Radarr"
	Readarr  App = "Readarr"
	Sonarr   App = "Sonarr"
	Whisparr App = "Whisparr"
)

// String turns an App name into a string.
func (a App) String() string {
	return string(a)
}

// Lower turns an App name into a lowercase string.
func (a App) Lower() string {
	return strings.ToLower(string(a))
}

// Client returns the default client, and is used if one is not passed in.
func Client(timeout time.Duration, verifySSL bool) *http.Client {
	return &http.Client{
		Timeout: timeout,
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: !verifySSL}, //nolint:gosec
		},
	}
}

// ClientWithDebug returns an http client with a debug logger enabled.
func ClientWithDebug(timeout time.Duration, verifySSL bool, logConfig debuglog.Config) *http.Client {
	client := Client(timeout, verifySSL)
	client.Transport = debuglog.NewLoggingRoundTripper(logConfig, client.Transport)

	return client
}

// CalendarTimeFilterFormat is the Go time format the calendar expects the filter to be in.
const CalendarTimeFilterFormat = "2006-01-02T03:04:05.000Z"

// StatusMessage represents the status of the item. All apps use this.
type StatusMessage struct {
	Title    string   `json:"title"`
	Messages []string `json:"messages"`
}

// BaseQuality is a base quality profile.
type BaseQuality struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Source     string `json:"source,omitempty"`
	Resolution int    `json:"resolution,omitempty"`
	Modifier   string `json:"modifier,omitempty"`
}

// Quality is a download quality profile attached to a movie, book, track or series.
// It may contain 1 or more profiles.
// Sonarr nor Readarr use Name or ID in this struct.
type Quality struct {
	Name     string           `json:"name,omitempty"`
	ID       int              `json:"id,omitempty"`
	Quality  *BaseQuality     `json:"quality,omitempty"`
	Items    []*Quality       `json:"items,omitempty"`
	Allowed  bool             `json:"allowed"`
	Revision *QualityRevision `json:"revision,omitempty"` // Not sure which app had this....
}

// QualityRevision is probably used in Sonarr.
type QualityRevision struct {
	Version  int64 `json:"version"`
	Real     int64 `json:"real"`
	IsRepack bool  `json:"isRepack,omitempty"`
}

// Ratings belong to a few types.
type Ratings struct {
	Votes      int64   `json:"votes"`
	Value      float64 `json:"value"`
	Popularity float64 `json:"popularity,omitempty"`
	Type       string  `json:"type,omitempty"`
}

// OpenRatings is a ratings type that has a source and type.
type OpenRatings map[string]Ratings

// IsLoaded is a generic struct used in a few places.
type IsLoaded struct {
	IsLoaded bool `json:"isLoaded"`
}

// Link is used in a few places.
type Link struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}

// Tag may be applied to nearly anything.
type Tag struct {
	ID    int    `json:"id,omitempty"`
	Label string `json:"label"`
}

// Image is used in a few places.
type Image struct {
	CoverType string `json:"coverType"`
	URL       string `json:"url,omitempty"`
	RemoteURL string `json:"remoteUrl,omitempty"`
	Extension string `json:"extension,omitempty"`
}

// Path is for unmanaged folder paths.
type Path struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

// RemotePathMapping is the remotePathMapping endpoint.
type RemotePathMapping struct {
	ID         int64  `json:"id,omitempty"`
	Host       string `json:"host"`
	RemotePath string `json:"remotePath"`
	LocalPath  string `json:"localPath"`
}

// Value is generic ID/Name struct applied to a few places.
type Value struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// FieldOutput is generic Name/Value struct applied to a few places.
type FieldOutput struct {
	Advanced                    bool            `json:"advanced,omitempty"`
	Order                       int64           `json:"order,omitempty"`
	HelpLink                    string          `json:"helpLink,omitempty"`
	HelpText                    string          `json:"helpText,omitempty"`
	Hidden                      string          `json:"hidden,omitempty"`
	Label                       string          `json:"label,omitempty"`
	Name                        string          `json:"name"`
	SelectOptionsProviderAction string          `json:"selectOptionsProviderAction,omitempty"`
	Type                        string          `json:"type,omitempty"`
	Privacy                     string          `json:"privacy"`
	Value                       interface{}     `json:"value,omitempty"`
	SelectOptions               []*SelectOption `json:"selectOptions,omitempty"`
}

// FieldInput is generic Name/Value struct applied to a few places.
type FieldInput struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value,omitempty"`
}

// SelectOption is part of Field.
type SelectOption struct {
	DividerAfter bool   `json:"dividerAfter,omitempty"`
	Order        int64  `json:"order"`
	Value        int64  `json:"value"`
	Hint         string `json:"hint"`
	Name         string `json:"name"`
}

// KeyValue is yet another reusable generic type.
type KeyValue struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

// BackupFile comes from the system/backup paths in all apps.
type BackupFile struct {
	Name string    `json:"name"`
	Path string    `json:"path"`
	Type string    `json:"type"`
	Time time.Time `json:"time"`
	ID   int64     `json:"id"`
	Size int64     `json:"size"`
}

// QueueDeleteOpts are the extra inputs when deleting an item from the Activity Queue.
// Set these appropriately for your expectations. All inputs are the same in all apps.
// Providing this input to the QueueDelete methods is optional; nil sets the defaults shown.
type QueueDeleteOpts struct {
	// Default True, use starr.False() to change it.
	RemoveFromClient *bool
	// Default False
	BlockList bool
	// Default False
	SkipRedownload bool
}

// Values turns delete options into http get query parameters.
func (o *QueueDeleteOpts) Values() url.Values {
	params := make(url.Values)
	params.Set("removeFromClient", "true")

	if o == nil {
		return params
	}

	params.Set("blocklist", fmt.Sprint(o.BlockList))
	params.Set("skipRedownload", fmt.Sprint(o.SkipRedownload))

	if o.RemoveFromClient != nil {
		params.Set("removeFromClient", fmt.Sprint(*o.RemoveFromClient))
	}

	return params
}

// PlayTime is used in at least Sonarr, maybe other places.
// Holds a string duration converted from hh:mm:ss.
type PlayTime struct { //nolint:musttag
	Original string
	time.Duration
}

// FormatItem is part of a quality profile.
type FormatItem struct {
	Format int64  `json:"format"`
	Name   string `json:"name"`
	Score  int64  `json:"score"`
}

// UnmarshalJSON parses a run time duration in format hh:mm:ss.
func (d *PlayTime) UnmarshalJSON(b []byte) error {
	d.Original = strings.Trim(string(b), `"'`)

	switch parts := strings.Split(d.Original, ":"); len(parts) {
	case 3: //nolint:gomnd // hh:mm:ss
		h, _ := strconv.Atoi(parts[0])
		m, _ := strconv.Atoi(parts[1])
		s, _ := strconv.Atoi(parts[2])
		d.Duration = (time.Hour * time.Duration(h)) + (time.Minute * time.Duration(m)) + (time.Second * time.Duration(s))
	case 2: //nolint:gomnd // mm:ss
		m, _ := strconv.Atoi(parts[0])
		s, _ := strconv.Atoi(parts[1])
		d.Duration = (time.Minute * time.Duration(m)) + (time.Second * time.Duration(s))
	case 1: // ss
		s, _ := strconv.Atoi(parts[0])
		d.Duration += (time.Second * time.Duration(s))
	}

	return nil
}

func (d *PlayTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.Original + `"`), nil
}

var _ json.Unmarshaler = (*PlayTime)(nil)

// ApplyTags is an enum used as an input for Bulk editors, and perhaps other places.
type ApplyTags string

// ApplyTags enum constants. Use these as inputs for "ApplyTags" member values.
// Schema doc'd here: https://radarr.video/docs/api/#/MovieEditor/put_api_v3_movie_editor
const (
	TagsAdd     ApplyTags = "add"
	TagsRemove  ApplyTags = "remove"
	TagsReplace ApplyTags = "replace"
)

// Ptr returns a pointer to an apply tags value. Useful for a BulkEdit struct.
func (a ApplyTags) Ptr() *ApplyTags {
	return &a
}

// TimeSpan is part of AudioTags and possibly used other places.
type TimeSpan struct {
	Ticks             int64 `json:"ticks"`
	Days              int64 `json:"days"`
	Hours             int64 `json:"hours"`
	Milliseconds      int64 `json:"milliseconds"`
	Minutes           int64 `json:"minutes"`
	Seconds           int64 `json:"seconds"`
	TotalDays         int64 `json:"totalDays"`
	TotalHours        int64 `json:"totalHours"`
	TotalMilliseconds int64 `json:"totalMilliseconds"`
	TotalMinutes      int64 `json:"totalMinutes"`
	TotalSeconds      int64 `json:"totalSeconds"`
}
