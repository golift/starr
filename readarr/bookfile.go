package readarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"path"
	"time"

	"golift.io/starr"
)

const bpBookFile = APIver + "/bookfile"

// BookFile represents the data from the bookfile endpoint.
type BookFile struct {
	AuthorID            int64          `json:"authorId"`
	BookID              int64          `json:"bookId"`
	Path                string         `json:"path"`
	Size                int            `json:"size,omitempty"`
	DateAdded           time.Time      `json:"dateAdded,omitempty"`
	Quality             *starr.Quality `json:"quality"`
	QualityWeight       int            `json:"qualityWeight,omitempty"`
	QualityCutoffNotMet bool           `json:"qualityCutoffNotMet"`
	ID                  int64          `json:"id"`
	AudioTags           *AudioTags     `json:"audioTags,omitempty"`
}

/* I've never seen audio tags in the wild. */

// AudioTags are part of a bookfile.
type AudioTags struct {
	Title          string          `json:"title"`
	CleanTitle     string          `json:"cleanTitle"`
	Authors        []string        `json:"authors"`
	AuthorTitle    string          `json:"authorTitle"`
	BookTitle      string          `json:"bookTitle"`
	SeriesTitle    string          `json:"seriesTitle"`
	SeriesIndex    string          `json:"seriesIndex"`
	ISBN           string          `json:"isbn"`
	ASIN           string          `json:"asin"`
	GoodreadsID    string          `json:"goodreadsId"`
	AuthorMBID     string          `json:"authorMBId"`
	BookMBID       string          `json:"bookMBId"`
	ReleaseMBID    string          `json:"releaseMBId"`
	RecordingMBID  string          `json:"recordingMBId"`
	TrackMBID      string          `json:"trackMBId"`
	DiscNumber     int             `json:"discNumber"`
	DiscCount      int             `json:"discCount"`
	Country        *AudioCountry   `json:"country"`
	Year           int             `json:"year"`
	Publisher      string          `json:"publisher"`
	Label          string          `json:"label"`
	Source         string          `json:"source"`
	CatalogNumber  string          `json:"catalogNumber"`
	Disambiguation string          `json:"disambiguation"`
	Duration       *starr.TimeSpan `json:"duration"`
	Quality        *starr.Quality  `json:"quality"`
	MediaInfo      *AudioMediaInfo `json:"mediaInfo"`
	TrackNumbers   []int           `json:"trackNumbers"`
	Language       string          `json:"language"`
	ReleaseGroup   string          `json:"releaseGroup"`
	ReleaseHash    string          `json:"releaseHash"`
}

// AudioMediaInfo is part of AudioTags.
type AudioMediaInfo struct {
	AudioFormat     string `json:"audioFormat"`
	AudioBitrate    int    `json:"audioBitrate"`
	AudioChannels   int    `json:"audioChannels"`
	AudioBits       int    `json:"audioBits"`
	AudioSampleRate int    `json:"audioSampleRate"`
}

// AudioCountry is part of AudioTags.
type AudioCountry struct {
	TwoLetterCode string `json:"twoLetterCode"`
	Name          string `json:"name"`
}

// GetBookFilesForAuthor returns the book files for an author.
func (r *Readarr) GetBookFilesForAuthor(authorID int64) ([]*BookFile, error) {
	return r.GetBookFilesForAuthorContext(context.Background(), authorID)
}

// GetBookFilesForAuthorContext returns the book files for an author.
func (r *Readarr) GetBookFilesForAuthorContext(ctx context.Context, authorID int64) ([]*BookFile, error) {
	var output []*BookFile

	req := starr.Request{URI: bpBookFile, Query: make(url.Values)}
	req.Query.Add("authorId", starr.Itoa(authorID))

	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetBookFilesForBook returns the book files for a book or books.
func (r *Readarr) GetBookFilesForBook(bookID ...int64) ([]*BookFile, error) {
	return r.GetBookFilesForBookContext(context.Background(), bookID...)
}

// GetBookFilesForBookContext returns the book files for a book or books.
func (r *Readarr) GetBookFilesForBookContext(ctx context.Context, bookID ...int64) ([]*BookFile, error) {
	var output []*BookFile

	req := starr.Request{URI: bpBookFile, Query: make(url.Values)}

	for _, id := range bookID {
		req.Query.Add("bookId", starr.Itoa(id))
	}

	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetBookFiles returns the requested book files by ID.
func (r *Readarr) GetBookFiles(bookFileIDs []int64) ([]*BookFile, error) {
	return r.GetBookFilesContext(context.Background(), bookFileIDs)
}

// GetBookFilesContext returns the requested book files by their IDs.
func (r *Readarr) GetBookFilesContext(ctx context.Context, bookFileIDs []int64) ([]*BookFile, error) {
	var output []*BookFile

	if len(bookFileIDs) == 0 {
		return output, nil
	}

	req := starr.Request{URI: bpBookFile, Query: make(url.Values)}

	for _, fileID := range bookFileIDs {
		req.Query.Add("bookFileIds", starr.Itoa(fileID))
	}

	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// UpdateBookFile updates a book file.
func (r *Readarr) UpdateBookFile(bookFile *BookFile) (*BookFile, error) {
	return r.UpdateBookFileContext(context.Background(), bookFile)
}

// UpdateBookFileContext updates a book file.
func (r *Readarr) UpdateBookFileContext(ctx context.Context, bookFile *BookFile) (*BookFile, error) {
	var output BookFile

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(bookFile); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpBookFile, err)
	}

	req := starr.Request{URI: path.Join(bpBookFile, starr.Itoa(bookFile.ID)), Body: &body}
	if err := r.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteBookFile deletes a book file.
func (r *Readarr) DeleteBookFile(bookFileID int64) error {
	return r.DeleteBookFileContext(context.Background(), bookFileID)
}

// DeleteBookFileContext deletes a book file.
func (r *Readarr) DeleteBookFileContext(ctx context.Context, bookFileID int64) error {
	req := starr.Request{URI: path.Join(bpBookFile, starr.Itoa(bookFileID))}
	if err := r.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}

// DeleteBookFiles bulk deletes book files by their IDs.
func (r *Readarr) DeleteBookFiles(bookFileIDs []int64) error {
	return r.DeleteBookFilesContext(context.Background(), bookFileIDs)
}

// DeleteBookFilesContext bulk deletes book files by their IDs.
func (r *Readarr) DeleteBookFilesContext(ctx context.Context, bookFileIDs []int64) error {
	postData := struct {
		T []int64 `json:"bookFileIds"`
	}{bookFileIDs}

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(&postData); err != nil {
		return fmt.Errorf("json.Marshal(%s): %w", bpBookFile, err)
	}

	req := starr.Request{URI: path.Join(bpBookFile, "bulk"), Body: &body}
	if err := r.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
