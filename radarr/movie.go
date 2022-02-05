package radarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"golift.io/starr"
)

// GetMovie grabs a movie from the queue, or all movies if tmdbId is 0.
func (r *Radarr) GetMovie(tmdbID int64) ([]*Movie, error) {
	return r.GetMovieContext(context.Background(), tmdbID)
}

// GetMovieContext grabs a movie from the queue, or all movies if tmdbId is 0.
func (r *Radarr) GetMovieContext(ctx context.Context, tmdbID int64) ([]*Movie, error) {
	params := make(url.Values)
	if tmdbID != 0 {
		params.Set("tmdbId", strconv.FormatInt(tmdbID, starr.Base10))
	}

	var movie []*Movie

	err := r.GetInto(ctx, "v3/movie", params, &movie)
	if err != nil {
		return nil, fmt.Errorf("api.Get(movie): %w", err)
	}

	return movie, nil
}

// GetMovieByID grabs a movie from the database by DB [movie] ID.
func (r *Radarr) GetMovieByID(movieID int64) (*Movie, error) {
	return r.GetMovieByIDContext(context.Background(), movieID)
}

// GetMovieByIDContext grabs a movie from the database by DB [movie] ID.
func (r *Radarr) GetMovieByIDContext(ctx context.Context, movieID int64) (*Movie, error) {
	var movie Movie

	err := r.GetInto(ctx, "v3/movie/"+strconv.FormatInt(movieID, starr.Base10), nil, &movie)
	if err != nil {
		return nil, fmt.Errorf("api.Get(movie): %w", err)
	}

	return &movie, nil
}

// UpdateMovie sends a PUT request to update a movie in place.
func (r *Radarr) UpdateMovie(movieID int64, movie *Movie) error {
	return r.UpdateMovieContext(context.Background(), movieID, movie)
}

// UpdateMovieContext sends a PUT request to update a movie in place.
func (r *Radarr) UpdateMovieContext(ctx context.Context, movieID int64, movie *Movie) error {
	put, err := json.Marshal(movie)
	if err != nil {
		return fmt.Errorf("json.Marshal(movie): %w", err)
	}

	params := make(url.Values)
	params.Add("moveFiles", "true")

	_, err = r.Put(ctx, "v3/movie/"+strconv.FormatInt(movieID, starr.Base10), params, put)
	if err != nil {
		return fmt.Errorf("api.Put(movie): %w", err)
	}

	return nil
}

// AddMovie adds a movie to the queue.
func (r *Radarr) AddMovie(movie *AddMovieInput) (*AddMovieOutput, error) {
	return r.AddMovieContext(context.Background(), movie)
}

// AddMovieContext adds a movie to the queue.
func (r *Radarr) AddMovieContext(ctx context.Context, movie *AddMovieInput) (*AddMovieOutput, error) {
	params := make(url.Values)
	params.Add("moveFiles", "true")

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(movie); err != nil {
		return nil, fmt.Errorf("json.Marshal(movie): %w", err)
	}

	var output AddMovieOutput
	if err := r.PostInto(ctx, "v3/movie", params, &body, &output); err != nil {
		return nil, fmt.Errorf("api.Post(movie): %w", err)
	}

	return &output, nil
}

// Lookup will search for movies matching the specified search term.
func (r *Radarr) Lookup(term string) ([]*Movie, error) {
	return r.LookupContext(context.Background(), term)
}

// LookupContext will search for movies matching the specified search term.
func (r *Radarr) LookupContext(ctx context.Context, term string) ([]*Movie, error) {
	var output []*Movie

	if term == "" {
		return output, nil
	}

	params := make(url.Values)
	params.Set("term", term)

	err := r.GetInto(ctx, "v3/movie/lookup", params, &output)
	if err != nil {
		return nil, fmt.Errorf("api.Get(movie/lookup): %w", err)
	}

	return output, nil
}
