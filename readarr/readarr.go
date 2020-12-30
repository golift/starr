package readarr

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// GetQueue returns the Readarr Queue (processing, but not yet imported).
func (r *Readarr) GetQueue(maxRecords int) (*Queue, error) {
	if maxRecords < 1 {
		maxRecords = 1
	}

	params := make(url.Values)
	params["pageSize"] = []string{strconv.Itoa(maxRecords)}

	var queue Queue

	err := r.GetInto("v1/queue", params, &queue)
	if err != nil {
		return nil, fmt.Errorf("api.Get(queue): %w", err)
	}

	return &queue, nil
}

// GetSystemStatus returns system status.
func (r *Readarr) GetSystemStatus() (*SystemStatus, error) {
	var status SystemStatus

	err := r.GetInto("v1/system/status", nil, &status)
	if err != nil {
		return &status, fmt.Errorf("api.Get(system/status): %w", err)
	}

	return &status, nil
}

// GetRootFolders returns all configured root folders.
func (r *Readarr) GetRootFolders() ([]*RootFolder, error) {
	var folders []*RootFolder

	err := r.GetInto("v1/rootFolder", nil, &folders)
	if err != nil {
		return nil, fmt.Errorf("api.Get(rootFolder): %w", err)
	}

	return folders, nil
}

// GetMetadataProfiles returns the metadata profiles.
func (r *Readarr) GetMetadataProfiles() ([]*MetadataProfile, error) {
	var profiles []*MetadataProfile

	err := r.GetInto("v1/metadataprofile", nil, &profiles)
	if err != nil {
		return nil, fmt.Errorf("api.Get(metadataprofile): %w", err)
	}

	return profiles, nil
}

// GetQualityProfiles returns the quality profiles.
func (r *Readarr) GetQualityProfiles() ([]*QualityProfile, error) {
	var profiles []*QualityProfile

	err := r.GetInto("v1/qualityprofile", nil, &profiles)
	if err != nil {
		return nil, fmt.Errorf("api.Get(qualityprofile): %w", err)
	}

	return profiles, nil
}

// GetBook returns books. All if gridID is 0.
func (r *Readarr) GetBook(gridID int64) ([]*Book, error) {
	params := make(url.Values)

	if gridID > 0 {
		params.Add("titleSlug", strconv.FormatInt(gridID, 10)) // this may change, but works for now.
	}

	var books []*Book

	err := r.GetInto("v1/book", params, &books)
	if err != nil {
		return nil, fmt.Errorf("api.Get(book): %w", err)
	}

	return books, nil
}

// GetBookByID return a book.
func (r *Readarr) GetBookByID(bookID int64) (*Book, error) {
	var book Book

	err := r.GetInto("v1/book/"+strconv.FormatInt(bookID, 10), nil, &book)
	if err != nil {
		return nil, fmt.Errorf("api.Get(book): %w", err)
	}

	return &book, nil
}

// UpdateBook adds a new book to the library.
func (r *Readarr) UpdateBook(bookID int64, book *Book) error {
	put, err := json.Marshal(book)
	if err != nil {
		return fmt.Errorf("json.Marshal(book): %w", err)
	}

	params := make(url.Values)
	params.Add("moveFiles", "true")

	b, err := r.Put("v1/book/"+strconv.FormatInt(bookID, 10), params, put)
	if err != nil {
		return fmt.Errorf("api.Put(book): %w", err)
	}

	fmt.Println("SHOW THIS TO CAPTAIN plz:", string(b))

	return nil
}

// AddBook adds a new book to the library.
func (r *Readarr) AddBook(book *AddBookInput) (*AddBookOutput, error) {
	body, err := json.Marshal(book)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal(book): %w", err)
	}

	params := make(url.Values)
	params.Add("moveFiles", "true")

	var output AddBookOutput

	err = r.PostInto("v1/book", params, body, &output)
	if err != nil {
		return nil, fmt.Errorf("api.Post(book): %w", err)
	}

	return &output, nil
}
