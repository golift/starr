package readarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"golift.io/starr"
)

// Author is the /api/v1/author endpoint.
type Author struct {
	ID                int64          `json:"id"`
	Status            string         `json:"status,omitempty"`
	AuthorName        string         `json:"authorName,omitempty"`
	ForeignAuthorID   string         `json:"foreignAuthorId,omitempty"`
	TitleSlug         string         `json:"titleSlug,omitempty"`
	Overview          string         `json:"overview,omitempty"`
	Links             []*starr.Link  `json:"links,omitempty"`
	Images            []*starr.Image `json:"images,omitempty"`
	Path              string         `json:"path,omitempty"`
	QualityProfileID  int            `json:"qualityProfileId,omitempty"`
	MetadataProfileID int            `json:"metadataProfileId,omitempty"`
	Genres            []interface{}  `json:"genres,omitempty"`
	CleanName         string         `json:"cleanName,omitempty"`
	SortName          string         `json:"sortName,omitempty"`
	Tags              []int          `json:"tags,omitempty"`
	Added             time.Time      `json:"added,omitempty"`
	Ratings           *starr.Ratings `json:"ratings,omitempty"`
	Statistics        *Statistics    `json:"statistics,omitempty"`
	LastBook          *AuthorBook    `json:"lastBook,omitempty"`
	NextBook          *AuthorBook    `json:"nextBook,omitempty"`
	Ended             bool           `json:"ended,omitempty"`
	Monitored         bool           `json:"monitored"`
}

// AuthorBook is part of an Author.
type AuthorBook struct {
	ID               int64           `json:"id"`
	AuthorMetadataID int             `json:"authorMetadataId"`
	ForeignBookID    string          `json:"foreignBookId"`
	TitleSlug        string          `json:"titleSlug"`
	Title            string          `json:"title"`
	ReleaseDate      time.Time       `json:"releaseDate"`
	Links            []*starr.Link   `json:"links"`
	Genres           []interface{}   `json:"genres"`
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

	var output interface{}

	uri := "v1/author/" + strconv.FormatInt(authorID, starr.Base10)
	if err := r.PutInto(ctx, uri, params, &body, &output); err != nil {
		return fmt.Errorf("api.Put(author): %w", err)
	}

	return nil
}
