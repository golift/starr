package readarr

import (
	"context"
	"fmt"
	"net/url"

	"golift.io/starr"
)

const bpRename = APIver + "/rename"

// Rename is the /api/v1/rename endpoint.
type Rename struct {
	ID           int64  `json:"id"`
	AuthorID     int64  `json:"authorId"`
	BookID       int64  `json:"bookId"`
	BookFileID   int64  `json:"bookFileId"`
	ExistingPath string `json:"existingPath,omitempty"`
	NewPath      string `json:"newPath,omitempty"`
}

// GetRenames checks if the specified book (database ID) from the author (database ID) needs to be renamed to
// follow the naming format. If bookId is set to -1, it will check all books at once.
func (r *Readarr) GetRenames(authorID int64, bookID int64) ([]*Rename, error) {
	return r.GetRenamesContext(context.Background(), authorID, bookID)
}

// GetRenamesContext checks if the specified book (database ID) from the author (database ID) needs to be renamed to
// follow the naming format. If bookId is set to -1, it will check all books at once.
func (r *Readarr) GetRenamesContext(ctx context.Context, authorID int64, bookID int64) ([]*Rename, error) {
	params := make(url.Values)
	params.Set("authorId", starr.Str(authorID))

	if bookID != -1 {
		params.Set("bookId", starr.Str(bookID))
	}

	var output []*Rename

	req := starr.Request{URI: bpRename, Query: params}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetAuthorRenames checks if the books from the specified author (database ID) need to be renamed to
// follow the naming format.
func (r *Readarr) GetAuthorRenames(authorID int64) ([]*Rename, error) {
	return r.GetRenamesContext(context.Background(), authorID, -1)
}

// GetAuthorRenamesContext checks if the books from the specified author (database ID) need to be renamed to
// follow the naming format.
func (r *Readarr) GetAuthorRenamesContext(ctx context.Context, authorID int64) ([]*Rename, error) {
	return r.GetRenamesContext(ctx, authorID, -1)
}

/* Doesn't exist yet
// GetAllRenames checks if any books need to be renamed to follow the naming format.
func (r *Readarr) GetAllRenames() ([]*Rename, error) {
	return r.GetRenamesContext(context.Background(), -1, -1)
} */

/* Doesn't exist yet
// GetAllRenamesContext checks if any books need to be renamed to follow the naming format.
func (r *Readarr) GetAllRenamesContext(ctx context.Context) ([]*Rename, error) {
	return r.GetRenamesContext(ctx, -1, -1)
} */
