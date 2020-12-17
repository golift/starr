package readarr

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// GetQueue returns the Readarr Queue (processing, but not yet imported).
func (r *Readarr) GetQueue(maxRecords int) (*Queue, error) {
	var queue *Queue

	if maxRecords < 1 {
		maxRecords = 1
	}

	params := make(url.Values)
	params["pageSize"] = []string{strconv.Itoa(maxRecords)}

	rawJSON, err := r.config.Req("v1/queue", params)
	if err != nil {
		return nil, fmt.Errorf("c.Req(queue): %w", err)
	}

	if err = json.Unmarshal(rawJSON, &queue); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(response): %w", err)
	}

	return queue, nil
}

// GetSystemStatus returns system status.
func (r *Readarr) GetSystemStatus() (*SystemStatus, error) {
	var status *SystemStatus

	rawJSON, err := r.config.Req("v1/system/status", nil)
	if err != nil {
		return status, fmt.Errorf("c.Req(system/status): %w", err)
	}

	if err = json.Unmarshal(rawJSON, &status); err != nil {
		return status, fmt.Errorf("json.Unmarshal(response): %w", err)
	}

	return status, nil
}

// GetRootFolders returns all configured root folders.
func (r *Readarr) GetRootFolders() ([]*RootFolder, error) {
	var folders []*RootFolder

	rawJSON, err := r.config.Req("v1/rootFolder", nil)
	if err != nil {
		return nil, fmt.Errorf("c.Req(rootFolder): %w", err)
	}

	if err = json.Unmarshal(rawJSON, &folders); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(response): %w", err)
	}

	return folders, nil
}

// GetMetadataProfiles returns the metadata profiles.
func (r *Readarr) GetMetadataProfiles() ([]*MetadataProfile, error) {
	var profiles []*MetadataProfile

	rawJSON, err := r.config.Req("v1/metadataprofile", nil)
	if err != nil {
		return nil, fmt.Errorf("c.Req(metadataprofile): %w", err)
	}

	if err = json.Unmarshal(rawJSON, &profiles); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(response): %w", err)
	}

	return profiles, nil
}

// GetQualityProfiles returns the quality profiles.
func (r *Readarr) GetQualityProfiles() ([]*QualityProfile, error) {
	var profiles []*QualityProfile

	rawJSON, err := r.config.Req("v1/qualityprofile", nil)
	if err != nil {
		return nil, fmt.Errorf("c.Req(qualityprofile): %w", err)
	}

	if err = json.Unmarshal(rawJSON, &profiles); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(response): %w", err)
	}

	return profiles, nil
}

// GetBook returns books. All if gridID is 0.
func (r *Readarr) GetBook(gridID int) ([]*Book, error) {
	var books []*Book

	params := make(url.Values)

	if gridID > 0 {
		params.Add("titleSlug", strconv.Itoa(gridID)) // this may change, but works for now.
	}

	rawJSON, err := r.config.Req("v1/book", params)
	if err != nil {
		return nil, fmt.Errorf("c.Req(book): %w", err)
	}

	if err = json.Unmarshal(rawJSON, &books); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(response): %w", err)
	}

	return books, nil
}

// AddBook adds a new book to the library.
func (r *Readarr) AddBook(book *AddBookInput) (*AddBookOutput, error) {
	body, err := json.Marshal(book)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal(book): %w", err)
	}

	params := make(url.Values)
	params.Add("moveFiles", "true")

	rawJSON, err := r.config.Req("v1/book", params, body...)
	if err != nil {
		return nil, fmt.Errorf("c.Req(book): %w", err)
	}

	var bookOutput *AddBookOutput

	if err = json.Unmarshal(rawJSON, &bookOutput); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(response): %w", err)
	}

	return bookOutput, nil
}
