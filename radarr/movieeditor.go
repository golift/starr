package radarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"golift.io/starr"
)

const bpMovieEditor = bpMovie + "/editor"

// BulkEdit is the input for the bulk movie editor endpoint.
// You may use starr.True(), starr.False(), starr.Int64(), and starr.String() to add data to the struct members.
type BulkEdit struct {
	MovieIDs            []int64 `json:"movieIds"`
	Monitored           *bool   `json:"monitored,omitempty"`
	QualityProfileID    *int64  `json:"qualityProfileId,omitempty"`
	MinimumAvailability *string `json:"minimumAvailability,omitempty"` // tba
	RootFolderPath      *string `json:"rootFolderPath,omitempty"`      // path
	Tags                []int   `json:"tags,omitempty"`                // [0]
	ApplyTags           *string `json:"applyTags,omitempty"`           // add
	MoveFiles           *bool   `json:"moveFiles,omitempty"`
	DeleteFiles         *bool   `json:"deleteFiles,omitempty"`        // delete only
	AddImportExclusion  *bool   `json:"addImportExclusion,omitempty"` // delete only
}

// EditMovies allows bulk diting many movies at once.
func (r *Radarr) EditMovies(editMovies *BulkEdit) ([]*Movie, error) {
	return r.EditMoviesContext(context.Background(), editMovies)
}

// EditMoviesContext allows bulk diting many movies at once.
func (r *Radarr) EditMoviesContext(ctx context.Context, editMovies *BulkEdit) ([]*Movie, error) {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(editMovies); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpMovieEditor, err)
	}

	var output []*Movie

	req := starr.Request{URI: bpMovieEditor, Body: &body}
	if err := r.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return output, nil
}

// DeleteMovies bulk deletes movies. Can also mark them as excluded, and delete their files.
func (r *Radarr) DeleteMovies(deleteMovies *BulkEdit) error {
	return r.DeleteMoviesContext(context.Background(), deleteMovies)
}

// DeleteDeleteMoviesContextMovies bulk deletes movies. Can also mark them as excluded, and delete their files.
func (r *Radarr) DeleteMoviesContext(ctx context.Context, deleteMovies *BulkEdit) error {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(deleteMovies); err != nil {
		return fmt.Errorf("json.Marshal(%s): %w", bpMovieEditor, err)
	}

	req := starr.Request{URI: bpMovieEditor, Body: &body}
	if err := r.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
