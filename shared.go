package starr

/* This file contains shared structs or constants for all the *arr apps. */

// StatusMessage represents the status of the item. All apps use this.
type StatusMessage struct {
	Title    string   `json:"title"`
	Messages []string `json:"messages"`
}

// Quality is a download quality attached to a movie, book, track or series.
type Quality struct {
	Quality struct {
		ID         int64  `json:"id"`
		Name       string `json:"name"`
		Source     string `json:"source,omitempty"`
		Resolution int    `json:"resolution,omitempty"`
		Modifier   string `json:"modifier,omitempty"`
	} `json:"quality"`
	Revision struct {
		Version  int64 `json:"version"`
		Real     int64 `json:"real"`
		IsRepack bool  `json:"isRepack,omitempty"`
	} `json:"revision,omitempty"`
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
