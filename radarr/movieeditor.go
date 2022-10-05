package radarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"golift.io/starr"
)

const bpMovieEditor = bpMovie + "/editor"

// EditMovies is the input for the bulk movie editor endpoint.
// You may use starr.True(), starr.False(), starr.Int64(), and starr.String() to add data to the struct members.
type EditMovies struct {
	MovieIDs            []int64 `json:"movieIds"`
	Monitored           *bool   `json:"monitored,omitempty"`
	QualityProfileID    *int64  `json:"qualityProfileId,omitempty"`
	MinimumAvailability *string `json:"minimumAvailability,omitempty"` // tba
	RootFolderPath      *string `json:"rootFolderPath,omitempty"`      // path
	Tags                []int   `json:"tags,omitempty"`                // [0]
	ApplyTags           *string `json:"applyTags,omitempty"`           // add
	MoveFiles           *bool   `json:"moveFiles,omitempty"`
	DeleteFiles         *bool   `json:"deleteFiles,omitempty"`
	AddImportExclusion  *bool   `json:"addImportExclusion,omitempty"`
}

func (r *Radarr) EditMovies(editMovies *EditMovies) ([]*Movie, error) {
	return r.EditMoviesContext(context.Background(), editMovies)
}

func (r *Radarr) EditMoviesContext(ctx context.Context, editMovies *EditMovies) ([]*Movie, error) {
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

func (r *Radarr) DeleteMovies(movieIDs []int64) error {
	return r.DeleteMoviesContext(context.Background(), movieIDs)
}

func (r *Radarr) DeleteMoviesContext(ctx context.Context, movieIDs []int64) error {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(map[string]interface{}{"movieIds": movieIDs}); err != nil {
		return fmt.Errorf("json.Marshal(%s): %w", bpMovieEditor, err)
	}

	req := starr.Request{URI: bpMovieEditor, Body: &body}
	if err := r.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
