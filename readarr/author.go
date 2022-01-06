package readarr

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"

	"golift.io/starr"
)

// GetAuthorByID returns an author.
func (r *Readarr) GetAuthorByID(authorID int64) (*Author, error) {
	var author Author

	err := r.GetInto("v1/author/"+strconv.FormatInt(authorID, starr.Base10), nil, &author)
	if err != nil {
		return nil, fmt.Errorf("api.Get(author): %w", err)
	}

	return &author, nil
}

// UpdateAuthor updates an author in place.
func (r *Readarr) UpdateAuthor(authorID int64, author *Author) error {
	put, err := json.Marshal(author)
	if err != nil {
		return fmt.Errorf("json.Marshal(author): %w", err)
	}

	params := make(url.Values)
	params.Add("moveFiles", "true")

	b, err := r.Put("v1/author/"+strconv.FormatInt(authorID, starr.Base10), params, put)
	if err != nil {
		return fmt.Errorf("api.Put(author): %w", err)
	}

	log.Println("SHOW THIS TO CAPTAIN plz:", string(b))

	return nil
}
