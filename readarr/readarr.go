package readarr

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"

	"golift.io/starr"
)

// GetTags returns all the tags.
func (r *Readarr) GetTags() ([]*starr.Tag, error) {
	var tags []*starr.Tag

	err := r.GetInto("v1/tag", nil, &tags)
	if err != nil {
		return nil, fmt.Errorf("api.Get(tag): %w", err)
	}

	return tags, nil
}

// UpdateTag updates the label for a tag.
func (r *Readarr) UpdateTag(tagID int, label string) (int, error) {
	body, err := json.Marshal(&starr.Tag{Label: label, ID: tagID})
	if err != nil {
		return 0, fmt.Errorf("json.Marshal(tag): %w", err)
	}

	var tag starr.Tag
	if err = r.PutInto("v1/tag", nil, body, &tag); err != nil {
		return tag.ID, fmt.Errorf("api.Put(tag): %w", err)
	}

	return tag.ID, nil
}

// AddTag adds a tag or returns the ID for an existing tag.
func (r *Readarr) AddTag(label string) (int, error) {
	body, err := json.Marshal(&starr.Tag{Label: label, ID: 0})
	if err != nil {
		return 0, fmt.Errorf("json.Marshal(tag): %w", err)
	}

	var tag starr.Tag
	if err = r.PostInto("v1/tag", nil, body, &tag); err != nil {
		return tag.ID, fmt.Errorf("api.Post(tag): %w", err)
	}

	return tag.ID, nil
}

// GetQueue returns the Readarr Queue (processing, but not yet imported).
func (r *Readarr) GetQueue(maxRecords int) (*Queue, error) {
	if maxRecords < 1 {
		maxRecords = 1
	}

	params := make(url.Values)
	params.Set("pageSize", strconv.Itoa(maxRecords))
	params.Set("includeUnknownAuthorItems", "true")

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

// AddQualityProfile updates a quality profile in place.
func (r *Readarr) AddQualityProfile(profile *QualityProfile) (int64, error) {
	post, err := json.Marshal(profile)
	if err != nil {
		return 0, fmt.Errorf("json.Marshal(profile): %w", err)
	}

	var output QualityProfile

	err = r.PostInto("v1/qualityProfile", nil, post, &output)
	if err != nil {
		return 0, fmt.Errorf("api.Post(qualityProfile): %w", err)
	}

	return output.ID, nil
}

// UpdateQualityProfile updates a quality profile in place.
func (r *Readarr) UpdateQualityProfile(profile *QualityProfile) error {
	put, err := json.Marshal(profile)
	if err != nil {
		return fmt.Errorf("json.Marshal(profile): %w", err)
	}

	_, err = r.Put("v1/qualityProfile/"+strconv.FormatInt(profile.ID, starr.Base10), nil, put)
	if err != nil {
		return fmt.Errorf("api.Put(qualityProfile): %w", err)
	}

	return nil
}

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

// GetBook returns books. All if gridID is empty.
func (r *Readarr) GetBook(gridID string) ([]*Book, error) {
	params := make(url.Values)

	if gridID != "" {
		params.Add("titleSlug", gridID) // this may change, but works for now.
	}

	var books []*Book

	err := r.GetInto("v1/book", params, &books)
	if err != nil {
		return nil, fmt.Errorf("api.Get(book): %w", err)
	}

	return books, nil
}

// GetBookByID returns a book.
func (r *Readarr) GetBookByID(bookID int64) (*Book, error) {
	var book Book

	err := r.GetInto("v1/book/"+strconv.FormatInt(bookID, starr.Base10), nil, &book)
	if err != nil {
		return nil, fmt.Errorf("api.Get(book): %w", err)
	}

	return &book, nil
}

// UpdateBook updates a book in place.
func (r *Readarr) UpdateBook(bookID int64, book *Book) error {
	put, err := json.Marshal(book)
	if err != nil {
		return fmt.Errorf("json.Marshal(book): %w", err)
	}

	params := make(url.Values)
	params.Add("moveFiles", "true")

	b, err := r.Put("v1/book/"+strconv.FormatInt(bookID, starr.Base10), params, put)
	if err != nil {
		return fmt.Errorf("api.Put(book): %w", err)
	}

	log.Println("SHOW THIS TO CAPTAIN plz:", string(b))

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

// GetCommands returns all available Readarr commands.
// These can be used with SendCommand.
func (r *Readarr) GetCommands() ([]*CommandResponse, error) {
	var output []*CommandResponse

	if err := r.GetInto("v1/command", nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(command): %w", err)
	}

	return output, nil
}

// SendCommand sends a command to Readarr.
func (r *Readarr) SendCommand(cmd *CommandRequest) (*CommandResponse, error) {
	if cmd == nil || cmd.Name == "" {
		return nil, nil
	}

	body, err := json.Marshal(cmd)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal(cmd): %w", err)
	}

	var output CommandResponse

	if err := r.PostInto("v1/command", nil, body, &output); err != nil {
		return nil, fmt.Errorf("api.Post(command): %w", err)
	}

	return &output, nil
}

// GetHistory returns the last few items from the history endpoint.
func (r *Readarr) GetHistory(maxRecords int) (*History, error) {
	if maxRecords < 1 {
		maxRecords = 1
	}

	params := make(url.Values)
	params.Set("pageSize", strconv.Itoa(maxRecords))

	var history History

	err := r.GetInto("v1/history", params, &history)
	if err != nil {
		return nil, fmt.Errorf("api.Get(history): %w", err)
	}

	return &history, nil
}

// GetBackupFiles returns all available Readarr backup files.
// Use GetBody to download a file using BackupFile.Path.
func (r *Readarr) GetBackupFiles() ([]*starr.BackupFile, error) {
	var output []*starr.BackupFile

	if err := r.GetInto("v1/system/backup", nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(system/backup): %w", err)
	}

	return output, nil
}
