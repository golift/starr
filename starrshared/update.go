package starrshared

import "time"

// UpdateChanges is the change log embedded in Update.
type UpdateChanges struct {
	New   []string `json:"new,omitempty"`
	Fixed []string `json:"fixed,omitempty"`
}

// Update is one available or installed update from the /update endpoint.
type Update struct {
	ID          int            `json:"id,omitempty"`
	Version     string         `json:"version,omitempty"`
	Branch      string         `json:"branch,omitempty"`
	ReleaseDate time.Time      `json:"releaseDate,omitzero"`
	FileName    string         `json:"fileName,omitempty"`
	URL         string         `json:"url,omitempty"`
	Installed   bool           `json:"installed"`
	InstalledOn time.Time      `json:"installedOn,omitzero"`
	Installable bool           `json:"installable"`
	Latest      bool           `json:"latest"`
	Changes     *UpdateChanges `json:"changes,omitempty"`
	Hash        string         `json:"hash,omitempty"`
}
