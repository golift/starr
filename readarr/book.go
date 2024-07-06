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

// Book is the /api/v1/book endpoint among others, and gets used across this package.
type Book struct {
	Added          time.Time      `json:"added"`
	AnyEditionOk   bool           `json:"anyEditionOk"`
	AuthorID       int64          `json:"authorId"`
	AuthorTitle    string         `json:"authorTitle"`
	Disambiguation string         `json:"disambiguation,omitempty"`
	Editions       []*Edition     `json:"editions"`
	ForeignBookID  string         `json:"foreignBookId"`
	Genres         []string       `json:"genres"`
	ID             int64          `json:"id"`
	Images         []*starr.Image `json:"images"`
	Links          []*starr.Link  `json:"links"`
	Monitored      bool           `json:"monitored"`
	Grabbed        bool           `json:"grabbed"`
	Overview       string         `json:"overview"`
	PageCount      int            `json:"pageCount"`
	Ratings        *starr.Ratings `json:"ratings"`
	ReleaseDate    time.Time      `json:"releaseDate"`
	RemoteCover    string         `json:"remoteCover,omitempty"`
	SeriesTitle    string         `json:"seriesTitle"`
	Statistics     *Statistics    `json:"statistics,omitempty"`
	Title          string         `json:"title"`
	TitleSlug      string         `json:"titleSlug"`
	Author         *Author        `json:"author"`
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
	TitleSlug        string         `json:"titleSlug"`        // Slugs are dumb
	Images           []*starr.Image `json:"images"`           // this is dumb too
	ForeignEditionID string         `json:"foreignEditionId"` // GRID ID
	Monitored        bool           `json:"monitored"`        // true
	ManualAdd        bool           `json:"manualAdd"`        // true
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

	req := starr.Request{URI: path.Join(bpBook, starr.Str(bookID))}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateBook updates a book in place.
func (r *Readarr) UpdateBook(bookID int64, book *Book, moveFiles bool) error {
	return r.UpdateBookContext(context.Background(), bookID, book, moveFiles)
}

// UpdateBookContext updates a book in place.
func (r *Readarr) UpdateBookContext(ctx context.Context, bookID int64, book *Book, moveFiles bool) error {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(book); err != nil {
		return fmt.Errorf("json.Marshal(%s): %w", bpBook, err)
	}

	req := starr.Request{
		URI:   path.Join(bpBook, starr.Str(bookID)),
		Query: make(url.Values),
		Body:  &body,
	}
	req.Query.Add("moveFiles", starr.Str(moveFiles))

	var output interface{} // do not know what this looks like.

	if err := r.PutInto(ctx, req, &output); err != nil {
		return fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return nil
}

// AddBook adds a new book to the library.
func (r *Readarr) AddBook(book *AddBookInput) (*Book, error) {
	return r.AddBookContext(context.Background(), book)
}

// AddBookContext adds a new book to the library.
func (r *Readarr) AddBookContext(ctx context.Context, book *AddBookInput) (*Book, error) {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(book); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpBook, err)
	}

	req := starr.Request{
		URI:   bpBook,
		Query: make(url.Values),
		Body:  &body,
	}

	var output Book
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

// DeleteBook removes a Book from the database.
// Setting deleteFiles true will delete all content for the Book.
func (r *Readarr) DeleteBook(bookID int64, deleteFiles, addImportExclusion bool) error {
	return r.DeleteBookContext(context.Background(), bookID, deleteFiles, addImportExclusion)
}

// DeleteBookContext removes a Book from the database.
// Setting deleteFiles true will delete all content for the Book.
func (r *Readarr) DeleteBookContext(ctx context.Context, bookID int64, deleteFiles, addImportExclusion bool) error {
	req := starr.Request{URI: path.Join(bpBook, starr.Str(bookID)), Query: make(url.Values)}
	req.Query.Set("deleteFiles", starr.Str(deleteFiles))
	req.Query.Set("addImportListExclusion", starr.Str(addImportExclusion))

	if err := r.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
