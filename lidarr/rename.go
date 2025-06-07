package lidarr

import (
	"context"
	"fmt"
	"net/url"

	"golift.io/starr"
)

const bpRename = APIver + "/rename"

// Rename is the /api/v1/rename endpoint.
type Rename struct {
	ID           int64    `json:"id"`
	ArtistID     int64    `json:"artistId"`
	AlbumID      int64    `json:"albumId"`
	TrackNumbers *[]int64 `json:"trackNumbers"`
	TrackFileID  int64    `json:"trackFileId"`
	ExistingPath *string  `json:"existingPath,omitempty"`
	NewPath      *string  `json:"newPath,omitempty"`
}

// GetRename checks if the tracks by the specified artist (database ID) on the specified album (database ID) need to
// be renamed to follow the naming format. If albumID is set to -1, it will check all albums at once.
func (l *Lidarr) GetRename(artistID int64, albumID int64) ([]*Rename, error) {
	return l.GetRenameContext(context.Background(), artistID, albumID)
}

// GetRenameContext checks if the tracks by the specified artist (database ID) on the specified album (database ID) need to
// be renamed to follow the naming format. If albumID is set to -1, it will check all albums at once.
func (l *Lidarr) GetRenameContext(ctx context.Context, artistID int64, albumID int64) ([]*Rename, error) {
	params := make(url.Values)
	params.Set("artistId", starr.Str(artistID))
	if albumID != -1 {
		params.Set("albumId", starr.Str(albumID))
	}

	var output []*Rename

	req := starr.Request{URI: bpRename, Query: params}
	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}
