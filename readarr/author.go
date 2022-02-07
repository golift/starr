package readarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"

	"golift.io/starr"
)

// GetAuthorByID returns an author.
func (r *Readarr) GetAuthorByID(authorID int64) (*Author, error) {
	return r.GetAuthorByIDContext(context.Background(), authorID)
}

// GetAuthorByIDContext returns an author.
func (r *Readarr) GetAuthorByIDContext(ctx context.Context, authorID int64) (*Author, error) {
	var author Author

	err := r.GetInto(ctx, "v1/author/"+strconv.FormatInt(authorID, starr.Base10), nil, &author)
	if err != nil {
		return nil, fmt.Errorf("api.Get(author): %w", err)
	}

	return &author, nil
}

// UpdateAuthor updates an author in place.
func (r *Readarr) UpdateAuthor(authorID int64, author *Author) error {
	return r.UpdateAuthorContext(context.Background(), authorID, author)
}

// UpdateAuthorContext updates an author in place.
func (r *Readarr) UpdateAuthorContext(ctx context.Context, authorID int64, author *Author) error {
	params := make(url.Values)
	params.Add("moveFiles", "true")

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(author); err != nil {
		return fmt.Errorf("json.Marshal(author): %w", err)
	}

	b, err := r.Put(ctx, "v1/author/"+strconv.FormatInt(authorID, starr.Base10), params, &body)
	if err != nil {
		return fmt.Errorf("api.Put(author): %w", err)
	}

	log.Println("SHOW THIS TO CAPTAIN plz:", string(b))

	return nil
}
