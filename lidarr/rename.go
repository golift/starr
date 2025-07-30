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
	ID           int64   `json:"id"`
	ArtistID     int64   `json:"artistId"`
	AlbumID      int64   `json:"albumId"`
	TrackNumbers []int64 `json:"trackNumbers"`
	TrackFileID  int64   `json:"trackFileId"`
	ExistingPath string  `json:"existingPath,omitempty"`
	NewPath      string  `json:"newPath,omitempty"`
}

// GetRenames checks if the tracks by the specified artist (database ID) on the specified album (database ID)
// need to be renamed to follow the naming format. If albumID is set to -1, it will check all albums at once.
func (l *Lidarr) GetRenames(artistID int64, albumID int64) ([]*Rename, error) {
	return l.GetRenamesContext(context.Background(), artistID, albumID)
}

// GetRenamesContext checks if the tracks by the specified artist (database ID) on the specified album (database ID)
// need to be renamed to follow the naming format. If albumID is set to -1, it will check all albums at once.
func (l *Lidarr) GetRenamesContext(ctx context.Context, artistID int64, albumID int64) ([]*Rename, error) {
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

// GetArtistRenames checks if the tracks by the specified artist (database ID) need to be renamed to
// follow the naming format.
func (l *Lidarr) GetArtistRenames(artistID int64) ([]*Rename, error) {
	return l.GetRenamesContext(context.Background(), artistID, -1)
}

// GetArtistRenamesContext checks if the tracks by the specified artist (database ID) need to be renamed to
// follow the naming format.
func (l *Lidarr) GetArtistRenamesContext(ctx context.Context, artistID int64) ([]*Rename, error) {
	return l.GetRenamesContext(ctx, artistID, -1)
}

/* Doesn't exist yet
// GetAllRenames checks if any tracks need to be renamed to follow the naming format.
func (l *Lidarr) GetAllRenames() ([]*Rename, error) {
	return l.GetRenamesContext(context.Background(), -1, -1)
} */

/* Doesn't exist yet
// GetAllRenamesContext checks if any tracks need to be renamed to follow the naming format.
func (l *Lidarr) GetAllRenamesContext(ctx context.Context) ([]*Rename, error) {
	return l.GetRenamesContext(ctx, -1, -1)
} */
