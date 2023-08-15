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

const bpAuthor = APIver + "/author"

// Author is the /api/v1/author endpoint.
type Author struct {
	ID                  int64          `json:"id"`
	Status              string         `json:"status,omitempty"`
	AuthorName          string         `json:"authorName,omitempty"`
	ForeignAuthorID     string         `json:"foreignAuthorId,omitempty"`
	TitleSlug           string         `json:"titleSlug,omitempty"`
	Overview            string         `json:"overview,omitempty"`
	Links               []*starr.Link  `json:"links,omitempty"`
	Images              []*starr.Image `json:"images,omitempty"`
	Path                string         `json:"path,omitempty"`
	QualityProfileID    int            `json:"qualityProfileId,omitempty"`
	MetadataProfileID   int            `json:"metadataProfileId,omitempty"`
	Genres              []string       `json:"genres,omitempty"`
	CleanName           string         `json:"cleanName,omitempty"`
	SortName            string         `json:"sortName,omitempty"`
	Tags                []int          `json:"tags,omitempty"`
	Added               time.Time      `json:"added,omitempty"`
	Ratings             *starr.Ratings `json:"ratings,omitempty"`
	Statistics          *Statistics    `json:"statistics,omitempty"`
	LastBook            *AuthorBook    `json:"lastBook,omitempty"`
	NextBook            *AuthorBook    `json:"nextBook,omitempty"`
	Ended               bool           `json:"ended,omitempty"`
	Monitored           bool           `json:"monitored"`
	AuthorMetadataID    int64          `json:"authorMetadataId"`
	AuthorNameLastFirst string         `json:"authorNameLastFirst"`
	MonitorNewItems     string         `json:"monitorNewItems"`
	SortNameLastFirst   string         `json:"sortNameLastFirst"`
}

// AuthorBook is part of an Author, and is very different from a normal Book type.
type AuthorBook struct {
	ID               int64           `json:"id"`
	AuthorMetadataID int             `json:"authorMetadataId"`
	ForeignBookID    string          `json:"foreignBookId"`
	TitleSlug        string          `json:"titleSlug"`
	Title            string          `json:"title"`
	ReleaseDate      time.Time       `json:"releaseDate"`
	Links            []*starr.Link   `json:"links"`
	Genres           []string        `json:"genres"`
	Ratings          *starr.Ratings  `json:"ratings"`
	CleanTitle       string          `json:"cleanTitle"`
	Monitored        bool            `json:"monitored"`
	AnyEditionOk     bool            `json:"anyEditionOk"`
	LastInfoSync     time.Time       `json:"lastInfoSync"`
	Added            time.Time       `json:"added"`
	AddOptions       *AddBookOptions `json:"addOptions"`
	AuthorMetadata   *starr.IsLoaded `json:"authorMetadata"`
	Author           *starr.IsLoaded `json:"author"`
	Editions         *starr.IsLoaded `json:"editions"`
	BookFiles        *starr.IsLoaded `json:"bookFiles"`
	SeriesLinks      *starr.IsLoaded `json:"seriesLinks"`
}

// Statistics for a Book, or maybe an author.
type Statistics struct {
	BookCount          int     `json:"bookCount"`
	BookFileCount      int     `json:"bookFileCount"`
	TotalBookCount     int     `json:"totalBookCount"`
	SizeOnDisk         int     `json:"sizeOnDisk"`
	PercentOfBooks     float64 `json:"percentOfBooks"`
	AvailableBookCount int     `json:"availableBookCount"`
}

// GetAuthorByID returns an author.
func (r *Readarr) GetAuthorByID(authorID int64) (*Author, error) {
	return r.GetAuthorByIDContext(context.Background(), authorID)
}

// GetAuthorByIDContext returns an author.
func (r *Readarr) GetAuthorByIDContext(ctx context.Context, authorID int64) (*Author, error) {
	var output Author

	req := starr.Request{URI: path.Join(bpAuthor, fmt.Sprint(authorID))}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateAuthor updates an author in place.
func (r *Readarr) UpdateAuthor(author *Author, moveFiles bool) (*Author, error) {
	return r.UpdateAuthorContext(context.Background(), author, moveFiles)
}

// UpdateAuthorContext updates an author in place.
func (r *Readarr) UpdateAuthorContext(ctx context.Context, author *Author, moveFiles bool) (*Author, error) {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(author); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpAuthor, err)
	}

	var output Author

	req := starr.Request{
		URI:   path.Join(bpAuthor, fmt.Sprint(author.ID)),
		Query: make(url.Values),
		Body:  &body,
	}
	req.Query.Add("moveFiles", fmt.Sprint(moveFiles))

	if err := r.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteAuthor removes an Author from the database. Setting deleteFiles true will delete all content for the Author.
func (r *Readarr) DeleteAuthor(authorID int64, deleteFiles, addImportExclusion bool) error {
	return r.DeleteAuthorContext(context.Background(), authorID, deleteFiles, addImportExclusion)
}

// DeleteAuthorContext removes na Author from the database. Setting deleteFiles true will delete all content for the Author.
func (r *Readarr) DeleteAuthorContext(ctx context.Context, authorID int64, deleteFiles, addImportExclusion bool) error {
	req := starr.Request{URI: path.Join(bpAuthor, fmt.Sprint(authorID)), Query: make(url.Values)}
	req.Query.Set("deleteFiles", fmt.Sprint(deleteFiles))
	req.Query.Set("addImportListExclusion", fmt.Sprint(addImportExclusion))

	if err := r.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
