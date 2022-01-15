package readarr

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"

	"golift.io/starr"
)

// GetBook returns books. All if gridID is empty.
func (r *Readarr) GetBook(gridID string) ([]*Book, error) {
	return r.GetBookContext(context.Background(), gridID)
}

func (r *Readarr) GetBookContext(ctx context.Context, gridID string) ([]*Book, error) {
	params := make(url.Values)

	if gridID != "" {
		params.Add("titleSlug", gridID) // this may change, but works for now.
	}

	var books []*Book

	err := r.GetInto(ctx, "v1/book", params, &books)
	if err != nil {
		return nil, fmt.Errorf("api.Get(book): %w", err)
	}

	return books, nil
}

// GetBookByID returns a book.
func (r *Readarr) GetBookByID(bookID int64) (*Book, error) {
	return r.GetBookByIDContext(context.Background(), bookID)
}

func (r *Readarr) GetBookByIDContext(ctx context.Context, bookID int64) (*Book, error) {
	var book Book

	err := r.GetInto(ctx, "v1/book/"+strconv.FormatInt(bookID, starr.Base10), nil, &book)
	if err != nil {
		return nil, fmt.Errorf("api.Get(book): %w", err)
	}

	return &book, nil
}

// UpdateBook updates a book in place.
func (r *Readarr) UpdateBook(bookID int64, book *Book) error {
	return r.UpdateBookContext(context.Background(), bookID, book)
}

func (r *Readarr) UpdateBookContext(ctx context.Context, bookID int64, book *Book) error {
	put, err := json.Marshal(book)
	if err != nil {
		return fmt.Errorf("json.Marshal(book): %w", err)
	}

	params := make(url.Values)
	params.Add("moveFiles", "true")

	b, err := r.Put(ctx, "v1/book/"+strconv.FormatInt(bookID, starr.Base10), params, put)
	if err != nil {
		return fmt.Errorf("api.Put(book): %w", err)
	}

	log.Println("SHOW THIS TO CAPTAIN plz:", string(b))

	return nil
}

// AddBook adds a new book to the library.
func (r *Readarr) AddBook(book *AddBookInput) (*AddBookOutput, error) {
	return r.AddBookContext(context.Background(), book)
}

func (r *Readarr) AddBookContext(ctx context.Context, book *AddBookInput) (*AddBookOutput, error) {
	body, err := json.Marshal(book)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal(book): %w", err)
	}

	params := make(url.Values)
	params.Add("moveFiles", "true")

	var output AddBookOutput

	err = r.PostInto(ctx, "v1/book", params, body, &output)
	if err != nil {
		return nil, fmt.Errorf("api.Post(book): %w", err)
	}

	return &output, nil
}

// Lookup will search for books matching the specified search term.
func (r *Readarr) Lookup(term string) ([]*Book, error) {
	return r.LookupContext(context.Background(), term)
}

func (r *Readarr) LookupContext(ctx context.Context, term string) ([]*Book, error) {
	var output []*Book

	if term == "" {
		return output, nil
	}

	params := make(url.Values)
	params.Set("term", term)

	err := r.GetInto(ctx, "v1/book/lookup", params, &output)
	if err != nil {
		return nil, fmt.Errorf("api.Get(book/lookup): %w", err)
	}

	return output, nil
}
