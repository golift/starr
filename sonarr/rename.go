package sonarr

import (
	"context"
	"fmt"
	"net/url"

	"golift.io/starr"
)

const bpRename = APIver + "/rename"

// Rename is the /api/v3/rename endpoint.
type Rename struct {
	ID             int64    `json:"id"`
	SeriesID       int64    `json:"seriesId"`
	SeasonNumber   int64    `json:"seasonNumber"`
	EpisodeNumbers *[]int64 `json:"episodeNumbers"`
	EpisodeFileId  int64    `json:"episodeFileId"`
	ExistingPath   *string  `json:"existingPath,omitempty"`
	NewPath        *string  `json:"newPath,omitempty"`
}

// GetRename checks if the episodes in the specified series (database ID) and season need to be renamed to follow the naming format.
// If seasonNumber is set to -1, it will check all seasons at once.
func (s *Sonarr) GetRename(seriesID int64, seasonNumber int64) ([]*Rename, error) {
	return s.GetRenameContext(context.Background(), seriesID, seasonNumber)
}

// GetRenameContext checks if the episodes in the specified series (database ID) and season need to be renamed to follow the naming format.
// If seasonNumber is set to -1, it will check all seasons at once.
func (s *Sonarr) GetRenameContext(ctx context.Context, seriesID int64, seasonNumber int64) ([]*Rename, error) {
	params := make(url.Values)
	params.Set("seriesId", starr.Str(seriesID))
	if seasonNumber != -1 {
		params.Set("seasonNumber", starr.Str(seasonNumber))
	}

	var output []*Rename

	req := starr.Request{URI: bpRename, Query: params}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}
