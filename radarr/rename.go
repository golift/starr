package radarr

import (
	"context"
	"fmt"
	"net/url"

	"golift.io/starr"
)

const bpRename = APIver + "/rename"

// Rename is the /api/v3/rename endpoint.
type Rename struct {
	ID           int64  `json:"id"`
	MovieID      int64  `json:"movieId"`
	MovieFileID  int64  `json:"movieFileId"`
	ExistingPath string `json:"existingPath,omitempty"`
	NewPath      string `json:"newPath,omitempty"`
}

// GetRenames checks if the movie with the specified movieID (database ID) needs to be renamed to
// follow the naming format.
func (r *Radarr) GetRenames(movieID int64) ([]*Rename, error) {
	return r.GetRenamesContext(context.Background(), movieID)
}

// GetRenamesContext checks if the movie with the specified movieID (database ID) needs to be renamed to
// follow the naming format.
func (r *Radarr) GetRenamesContext(ctx context.Context, movieID int64) ([]*Rename, error) {
	params := make(url.Values)
	params.Set("movieId", starr.Str(movieID))

	var output []*Rename

	req := starr.Request{URI: bpRename, Query: params}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

/* Doesn't exist yet
// GetAllRenames checks if any movies need to be renamed to follow the naming format.
func (r *Radarr) GetAllRenames() ([]*Rename, error) {
	return r.GetRenamesContext(context.Background(), -1)
} */

/* Doesn't exist yet
// GetAllRenamesContext checks if any movies need to be renamed to follow the naming format.
func (r *Radarr) GetAllRenamesContext(ctx context.Context) ([]*Rename, error) {
	return r.GetRenamesContext(ctx, -1)
} */
