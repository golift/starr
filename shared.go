package starr

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

/* This file contains shared structs and constants for all the Starr apps. */

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
	Value                       any             `json:"value,omitempty"`
	SelectOptions               []*SelectOption `json:"selectOptions,omitempty"`
}

// FieldInput is generic Name/Value struct applied to a few places.
type FieldInput struct {
	Name  string `json:"name"`
	Value any    `json:"value,omitempty"`
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
	// Default False
	ChangeCategory bool
}

// Values turns delete options into http get query parameters.
func (o *QueueDeleteOpts) Values() url.Values {
	params := make(url.Values)
	params.Set("removeFromClient", "true")

	if o == nil {
		return params
	}

	params.Set("blocklist", Str(o.BlockList))
	params.Set("skipRedownload", Str(o.SkipRedownload))
	params.Set("changeCategory", Str(o.ChangeCategory))

	if o.RemoveFromClient != nil {
		params.Set("removeFromClient", Str(*o.RemoveFromClient))
	}

	return params
}

// FormatItem is part of a quality profile.
type FormatItem struct {
	Format int64  `json:"format"`
	Name   string `json:"name"`
	Score  int64  `json:"score"`
}

// PlayTime is used in at least Sonarr, maybe other places.
// Holds a string duration converted from hh:mm:ss.
type PlayTime struct {
	time.Duration

	Original string
}

var (
	_ json.Unmarshaler = (*PlayTime)(nil)
	_ json.Marshaler   = (*PlayTime)(nil)
)

// UnmarshalJSON parses a run time duration in format hh:mm:ss or hh:mm:ss.fraction.
func (d *PlayTime) UnmarshalJSON(b []byte) error {
	d.Original = strings.Trim(string(b), `"'`)

	switch parts := strings.Split(d.Original, ":"); len(parts) {
	case 3: //nolint:mnd // hh:mm:ss or hh:mm:ss.fraction
		h, _ := strconv.Atoi(parts[0])
		m, _ := strconv.Atoi(parts[1])
		s, _ := strconv.ParseFloat(parts[2], 64)
		d.Duration = (time.Hour * time.Duration(h)) + (time.Minute * time.Duration(m)) + time.Duration(s*float64(time.Second))
	case 2: //nolint:mnd // mm:ss or mm:ss.fraction
		m, _ := strconv.Atoi(parts[0])
		s, _ := strconv.ParseFloat(parts[1], 64)
		d.Duration = (time.Minute * time.Duration(m)) + time.Duration(s*float64(time.Second))
	case 1: // ss or ss.fraction
		s, _ := strconv.ParseFloat(parts[0], 64)
		d.Duration = time.Duration(s * float64(time.Second))
	}

	return nil
}

// MarshalJSON marshals the PlayTime to JSON.
//
//nolint:wrapcheck,mnd // no value added, seconds per hour, etc.
func (d *PlayTime) MarshalJSON() ([]byte, error) {
	s := d.Original
	if s != "" {
		return json.Marshal(s)
	}

	// Format Duration as hh:mm:ss or hh:mm:ss.fraction to match API shape.
	total := d.Seconds()
	if total == 0 {
		return json.Marshal("00:00:00")
	}

	hours := int(total / 3600)
	mins := int((total - float64(hours*3600)) / 60)
	secs := total - float64(hours*3600) - float64(mins*60)

	return json.Marshal(fmt.Sprintf("%02d:%02d:%s", hours, mins, strconv.FormatFloat(secs, 'f', -1, 64)))
}

// ApplyTags is an enum used as an input for Bulk editors, and perhaps other places.
type ApplyTags string

// ApplyTags enum constants. Use these as inputs for "ApplyTags" member values.
// Schema doc'd here: https://radarr.video/docs/api/#/MovieEditor/put_api_v3_movie_editor
const (
	TagsAdd     ApplyTags = "add"
	TagsRemove  ApplyTags = "remove"
	TagsReplace ApplyTags = "replace"
)

// TimeSpan is used when a starr API returns a duration as an object (ticks, days, hours, etc.).
// For ParsedTrackInfo/audioTags, the APIs use a string with format "date-span" instead; use string there.
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

// Protocol used to download media. Comes with enum constants.
type Protocol string

// These are all the starr-supported protocols.
const (
	ProtocolUnknown Protocol = "unknown"
	ProtocolUsenet  Protocol = "usenet"
	ProtocolTorrent Protocol = "torrent"
)

// BulkIndexer is the input to UpdateIndexers on all apps except Prowlarr.
// Use the starr.True/False/Ptr() funcs to create the pointers.
type BulkIndexer struct {
	IDs                     []int64   `json:"ids"`
	Tags                    []int     `json:"tags,omitempty"`
	ApplyTags               ApplyTags `json:"applyTags,omitempty"`
	EnableRss               *bool     `json:"enableRss,omitempty"`
	EnableAutomaticSearch   *bool     `json:"enableAutomaticSearch,omitempty"`
	EnableInteractiveSearch *bool     `json:"enableInteractiveSearch,omitempty"`
	Priority                *int64    `json:"priority,omitempty"`
}
