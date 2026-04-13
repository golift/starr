package sonarr

import (
	"context"
	"fmt"
	"time"

	"golift.io/starr"
)

const bpUpdate = APIver + "/update"

// UpdateChanges is the change log embedded in Update.
type UpdateChanges struct {
	New   []string `json:"new,omitempty"`
	Fixed []string `json:"fixed,omitempty"`
}

// Update is one available or installed update from /api/v3/update.
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

// GetUpdates returns available application updates.
func (s *Sonarr) GetUpdates() ([]*Update, error) {
	return s.GetUpdatesContext(context.Background())
}

// GetUpdatesContext returns available application updates.
func (s *Sonarr) GetUpdatesContext(ctx context.Context) ([]*Update, error) {
	var output []*Update

	req := starr.Request{URI: bpUpdate}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}
