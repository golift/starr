package starr

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

/* This file contains shared structs or constants for all the *arr apps. */

// App can be used to satisfy a context value key.
// It is not used in this library; provided for convenience.
type App string

// These constants are just here for convenience.
// If you add more here, add them to String() below.
const (
	Emby     App = "Emby"
	Lidarr   App = "Lidarr"
	Plex     App = "Plex"
	Prowlarr App = "Prowlarr"
	Radarr   App = "Radarr"
	Readarr  App = "Readarr"
	Sonarr   App = "Sonarr"
)

// Silly constants to avoid screwing up integer->string conversions.
const (
	Base10 = 10
)

// String turns an App name into a string.
func (a App) String() string {
	return string(a)
}

// Lower turns an App name into a lowercase string.
func (a App) Lower() string {
	return strings.ToLower(string(a))
}

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
}

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
	ID    int
	Label string
}

// Image is used in a few places.
type Image struct {
	CoverType string `json:"coverType"`
	URL       string `json:"url"`
	RemoteURL string `json:"remoteUrl,omitempty"`
	Extension string `json:"extension,omitempty"`
}

// Path is for unmanaged folder paths.
type Path struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

// Value is generic ID/Name struct applied to a few places.
type Value struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
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

// PlayTime is used in at least Sonarr, maybe other places.
// Holds a string duration converted from hh:mm:ss.
type PlayTime struct {
	Original string
	time.Duration
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
