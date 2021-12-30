package starr

import (
	"net/url"
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

// Silly constants to not screw up integer->string conversions.
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
// Readarr does not use Name or ID in this struct.
type Quality struct {
	Name     string           `json:"name,omitempty"`
	ID       int              `json:"id,omitempty"`
	Quality  *BaseQuality     `json:"quality,omitempty"`
	Items    []*Quality       `json:"items"`
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
}

// Req is the input to search requests that have page-able responses.
// These are turned into HTTP parameters.
type Req struct {
	PageSize   int    // 10 is default if not provided.
	Page       int    // 1 or higher
	SortKey    string // date, timeleft, others?
	SortDir    string // asc, desc
	url.Values        // Additional values that may be set.
}

// Params returns a brand new url.Values with all request parameters combined.
func (r *Req) Params() url.Values {
	params := make(url.Values)

	if r.Page > 0 {
		params.Set("page", strconv.Itoa(r.Page))
	} else {
		params.Set("page", "1")
	}

	if r.PageSize > 0 {
		params.Set("pageSize", strconv.Itoa(r.PageSize))
	} else {
		params.Set("pageSize", "10")
	}

	if r.SortKey != "" {
		params.Set("sortKey", r.SortKey)
	} else {
		params.Set("sortKey", "date") // timeleft
	}

	if r.SortDir != "" {
		params.Set("sortDir", r.SortDir)
	} else {
		params.Set("sortDir", "asc") // desc
	}

	for k, v := range r.Values {
		for _, val := range v {
			params.Set(k, val)
		}
	}

	return params
}

// Encode turns our request parameters into a URI string.
func (r *Req) Encode() string {
	return r.Params().Encode()
}

// CheckSet sets a request parameter if it's not already set.
func (r *Req) CheckSet(key, value string) { //nolint:cyclop
	switch strings.ToLower(key) {
	case "page":
		if r.Page == 0 {
			r.Page, _ = strconv.Atoi(value)
		}
	case "pagesize":
		if r.PageSize == 0 {
			r.PageSize, _ = strconv.Atoi(value)
		}
	case "sortkey":
		if r.SortKey == "" {
			r.SortKey = value
		}
	case "sortdir":
		if r.SortDir == "" {
			r.SortDir = value
		}
	default:
		if r.Values == nil || r.Values.Get(key) == "" {
			r.Values.Set(key, value)
		}
	}
}

// Set sets a request parameter.
func (r *Req) Set(key, value string) {
	switch strings.ToLower(key) {
	case "page":
		r.Page, _ = strconv.Atoi(value)
	case "pagesize":
		r.PageSize, _ = strconv.Atoi(value)
	case "sortkey":
		r.SortKey = value
	case "sortdir":
		r.SortDir = value
	default:
		if r.Values == nil {
			r.Values = make(url.Values)
		}

		r.Values.Set(key, value)
	}
}

// SetPerPage returns a proper perPage value that is not equal to zero,
// and not larger than the record count desired. If the count is zero, then
// perPage can be anything other than zero.
// This is used by paginated methods in the starr modules.
func SetPerPage(records, perPage int) int {
	const perPageDefault = 500

	if perPage <= 1 {
		if records > perPageDefault || records == 0 {
			perPage = perPageDefault
		} else {
			perPage = records
		}
	} else if perPage > records && records != 0 {
		perPage = records
	}

	return perPage
}

// AdjustPerPage to make sure we don't go over, or ask for more records than exist.
// This is used by paginated methods in the starr modules.
// 'records' is the number requested, 'total' is the number in the app,
// 'collected' is how many we have so far, and 'perPage' is the current perPage setting.
func AdjustPerPage(records, total, collected, perPage int) int {
	if d := records - collected; perPage > d && d > 0 {
		perPage = d
	}

	if d := total - collected; perPage > d {
		perPage = d
	}

	return perPage
}
