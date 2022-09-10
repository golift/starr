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

const bpBook = APIver + "/book"

// Book is the /api/v1/book endpoint.
type Book struct {
	Title          string         `json:"title"`
	SeriesTitle    string         `json:"seriesTitle"`
	AuthorTitle    string         `json:"authorTitle"`
	Overview       string         `json:"overview"`
	AuthorID       int64          `json:"authorId"`
	ForeignBookID  string         `json:"foreignBookId"`
	TitleSlug      string         `json:"titleSlug"`
	Monitored      bool           `json:"monitored"`
	AnyEditionOk   bool           `json:"anyEditionOk"`
	Ratings        *starr.Ratings `json:"ratings"`
	ReleaseDate    time.Time      `json:"releaseDate"`
	Added          time.Time      `json:"added"`
	PageCount      int            `json:"pageCount"`
	Genres         []string       `json:"genres"`
	Images         []*starr.Image `json:"images"`
	Links          []*starr.Link  `json:"links"`
	Statistics     *Statistics    `json:"statistics,omitempty"`
	Editions       []*Edition     `json:"editions"`
	ID             int64          `json:"id"`
	Disambiguation string         `json:"disambiguation,omitempty"`
	RemoteCover    string         `json:"remoteCover,omitempty"`
}

// BookAuthor of a Book.
type BookAuthor struct {
	ID                int64          `json:"id"`
	Status            string         `json:"status"`
	AuthorName        string         `json:"authorName"`
	ForeignAuthorID   string         `json:"foreignAuthorId"`
	TitleSlug         string         `json:"titleSlug"`
	Overview          string         `json:"overview"`
	Links             []*starr.Link  `json:"links"`
	Images            []*starr.Image `json:"images"`
	Path              string         `json:"path"`
	QualityProfileID  int64          `json:"qualityProfileId"`
	MetadataProfileID int64          `json:"metadataProfileId"`
	Genres            []interface{}  `json:"genres"`
	CleanName         string         `json:"cleanName"`
	SortName          string         `json:"sortName"`
	Tags              []int          `json:"tags"`
	Added             time.Time      `json:"added"`
	Ratings           *starr.Ratings `json:"ratings"`
	Statistics        *Statistics    `json:"statistics"`
	Monitored         bool           `json:"monitored"`
	Ended             bool           `json:"ended"`
}

// Edition is more Book meta data.
type Edition struct {
	ID               int64          `json:"id"`
	BookID           int64          `json:"bookId"`
	ForeignEditionID string         `json:"foreignEditionId"`
	TitleSlug        string         `json:"titleSlug"`
	Isbn13           string         `json:"isbn13"`
	Asin             string         `json:"asin"`
	Title            string         `json:"title"`
	Overview         string         `json:"overview"`
	Format           string         `json:"format"`
	Publisher        string         `json:"publisher"`
	PageCount        int            `json:"pageCount"`
	ReleaseDate      time.Time      `json:"releaseDate"`
	Images           []*starr.Image `json:"images"`
	Links            []*starr.Link  `json:"links"`
	Ratings          *starr.Ratings `json:"ratings"`
	Monitored        bool           `json:"monitored"`
	ManualAdd        bool           `json:"manualAdd"`
	IsEbook          bool           `json:"isEbook"`
}

// AddBookInput is the input to add a book.
type AddBookInput struct {
	Monitored     bool              `json:"monitored"`
	Tags          []int             `json:"tags"`
	AddOptions    *AddBookOptions   `json:"addOptions"`    // Contains Search.
	Author        *AddBookAuthor    `json:"author"`        // Contains Author ID
	Editions      []*AddBookEdition `json:"editions"`      // contains GRID Edition ID
	ForeignBookID string            `json:"foreignBookId"` // GRID Book ID.
}

// AddBookOptions is part of AddBookInput.
type AddBookOptions struct {
	AddType          string `json:"addType,omitempty"`
	SearchForNewBook bool   `json:"searchForNewBook"`
}

// AddBookAuthor is part of AddBookInput.
type AddBookAuthor struct {
	Monitored         bool              `json:"monitored"`         // true?
	QualityProfileID  int64             `json:"qualityProfileId"`  // required
	MetadataProfileID int64             `json:"metadataProfileId"` // required
	ForeignAuthorID   string            `json:"foreignAuthorId"`   // required
	RootFolderPath    string            `json:"rootFolderPath"`    // required
	Tags              []int             `json:"tags"`
	AddOptions        *AddAuthorOptions `json:"addOptions"`
}

// AddAuthorOptions is part of AddBookAuthor.
type AddAuthorOptions struct {
	SearchForMissingBooks bool    `json:"searchForMissingBooks"`
	Monitored             bool    `json:"monitored"`
	Monitor               string  `json:"monitor"`
	BooksToMonitor        []int64 `json:"booksToMonitor"`
}

// AddBookEdition is part of AddBookInput.
type AddBookEdition struct {
	Title            string         `json:"title"`            // Edition Title
	TitleSlug        interface{}    `json:"titleSlug"`        // Slugs are dumb
	Images           []*starr.Image `json:"images"`           // this is dumb too
	ForeignEditionID string         `json:"foreignEditionId"` // GRID ID
	Monitored        bool           `json:"monitored"`        // true
	ManualAdd        bool           `json:"manualAdd"`        // true
}

// AddBookOutput is returned when a book is added.
type AddBookOutput struct {
	ID            int64          `json:"id"`
	AuthorID      int64          `json:"authorId"`
	PageCount     int            `json:"pageCount"`
	Title         string         `json:"title"`
	SeriesTitle   string         `json:"seriesTitle"`
	Overview      string         `json:"overview"`
	ForeignBookID string         `json:"foreignBookId"`
	TitleSlug     string         `json:"titleSlug"`
	Ratings       *starr.Ratings `json:"ratings"`
	ReleaseDate   time.Time      `json:"releaseDate"`
	Genres        []interface{}  `json:"genres"`
	Author        *BookAuthor    `json:"author"`
	Images        []*starr.Image `json:"images"`
	Links         []*starr.Link  `json:"links"`
	Statistics    *Statistics    `json:"statistics"`
	Editions      []*Edition     `json:"editions"`
	Monitored     bool           `json:"monitored"`
	AnyEditionOk  bool           `json:"anyEditionOk"`
}

// GetBook returns books. All books are returned if gridID is empty.
func (r *Readarr) GetBook(gridID string) ([]*Book, error) {
	return r.GetBookContext(context.Background(), gridID)
}

// GetBookContext returns books. All books are returned if gridID is empty.
func (r *Readarr) GetBookContext(ctx context.Context, gridID string) ([]*Book, error) {
	req := starr.Request{URI: bpBook, Query: make(url.Values)}
	if gridID != "" {
		req.Query.Add("titleSlug", gridID) // this may change, but works for now.
	}

	var output []*Book

	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetBookByID returns a book.
func (r *Readarr) GetBookByID(bookID int64) (*Book, error) {
	return r.GetBookByIDContext(context.Background(), bookID)
}

// GetBookByIDContext returns a book.
func (r *Readarr) GetBookByIDContext(ctx context.Context, bookID int64) (*Book, error) {
	var output Book

	req := starr.Request{URI: path.Join(bpBook, fmt.Sprint(bookID))}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateBook updates a book in place.
func (r *Readarr) UpdateBook(bookID int64, book *Book) error {
	return r.UpdateBookContext(context.Background(), bookID, book)
}

// UpdateBookContext updates a book in place.
func (r *Readarr) UpdateBookContext(ctx context.Context, bookID int64, book *Book) error {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(book); err != nil {
		return fmt.Errorf("json.Marshal(%s): %w", bpBook, err)
	}

	req := starr.Request{
		URI:   path.Join(bpBook, fmt.Sprint(bookID)),
		Query: make(url.Values),
		Body:  &body,
	}
	req.Query.Add("moveFiles", "true")

	var output interface{} // do not know what this looks like.

	if err := r.PutInto(ctx, req, &output); err != nil {
		return fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return nil
}

// AddBook adds a new book to the library.
func (r *Readarr) AddBook(book *AddBookInput) (*AddBookOutput, error) {
	return r.AddBookContext(context.Background(), book)
}

// AddBookContext adds a new book to the library.
func (r *Readarr) AddBookContext(ctx context.Context, book *AddBookInput) (*AddBookOutput, error) {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(book); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpBook, err)
	}

	req := starr.Request{
		URI:   bpBook,
		Query: make(url.Values),
		Body:  &body,
	}
	req.Query.Add("moveFiles", "true")

	var output AddBookOutput
	if err := r.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// Lookup will search for books matching the specified search term.
func (r *Readarr) Lookup(term string) ([]*Book, error) {
	return r.LookupContext(context.Background(), term)
}

// LookupContext will search for books matching the specified search term.
func (r *Readarr) LookupContext(ctx context.Context, term string) ([]*Book, error) {
	var output []*Book

	if term == "" {
		return output, nil
	}

	req := starr.Request{URI: path.Join(bpBook, "lookup"), Query: make(url.Values)}
	req.Query.Set("term", term)

	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}
