package radarr

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

const bpMovie = APIver + "/movie"

// Movie is the /api/v3/movie endpoint.
type Movie struct {
	ID                    int64               `json:"id"`
	Title                 string              `json:"title,omitempty"`
	Path                  string              `json:"path,omitempty"`
	MinimumAvailability   Availability        `json:"minimumAvailability,omitempty"`
	QualityProfileID      int64               `json:"qualityProfileId,omitempty"`
	TmdbID                int64               `json:"tmdbId,omitempty"`
	OriginalTitle         string              `json:"originalTitle,omitempty"`
	AlternateTitles       []*AlternativeTitle `json:"alternateTitles,omitempty"`
	SecondaryYearSourceID int                 `json:"secondaryYearSourceId,omitempty"`
	SortTitle             string              `json:"sortTitle,omitempty"`
	SizeOnDisk            int64               `json:"sizeOnDisk,omitempty"`
	Status                string              `json:"status,omitempty"`
	Overview              string              `json:"overview,omitempty"`
	InCinemas             time.Time           `json:"inCinemas,omitempty"`
	PhysicalRelease       time.Time           `json:"physicalRelease,omitempty"`
	DigitalRelease        time.Time           `json:"digitalRelease,omitempty"`
	Images                []*starr.Image      `json:"images,omitempty"`
	Website               string              `json:"website,omitempty"`
	Year                  int                 `json:"year,omitempty"`
	YouTubeTrailerID      string              `json:"youTubeTrailerId,omitempty"`
	Studio                string              `json:"studio,omitempty"`
	FolderName            string              `json:"folderName,omitempty"`
	Runtime               int                 `json:"runtime,omitempty"`
	CleanTitle            string              `json:"cleanTitle,omitempty"`
	ImdbID                string              `json:"imdbId,omitempty"`
	TitleSlug             string              `json:"titleSlug,omitempty"`
	Certification         string              `json:"certification,omitempty"`
	Genres                []string            `json:"genres,omitempty"`
	Tags                  []int               `json:"tags,omitempty"`
	Added                 time.Time           `json:"added,omitempty"`
	Ratings               starr.OpenRatings   `json:"ratings,omitempty"`
	MovieFile             *MovieFile          `json:"movieFile,omitempty"`
	Collection            *Collection         `json:"collection,omitempty"`
	HasFile               bool                `json:"hasFile,omitempty"`
	IsAvailable           bool                `json:"isAvailable,omitempty"`
	Monitored             bool                `json:"monitored"`
	Popularity            float64             `json:"popularity"`
	OriginalLanguage      *starr.Value        `json:"originalLanguage,omitempty"`
	AddOptions            *AddMovieOptions    `json:"addOptions,omitempty"` // only available upon adding a movie.
}

// Collection belongs to a Movie.
type Collection struct {
	Name   string         `json:"name"`
	TmdbID int64          `json:"tmdbId"`
	Images []*starr.Image `json:"images"`
}

// AddMovieInput is the input for a new movie.
type AddMovieInput struct {
	Title               string           `json:"title,omitempty"`
	TitleSlug           string           `json:"titleSlug,omitempty"`
	MinimumAvailability Availability     `json:"minimumAvailability,omitempty"`
	RootFolderPath      string           `json:"rootFolderPath"`
	TmdbID              int64            `json:"tmdbId"`
	QualityProfileID    int64            `json:"qualityProfileId"`
	ProfileID           int64            `json:"profileId,omitempty"`
	Year                int              `json:"year,omitempty"`
	Images              []*starr.Image   `json:"images,omitempty"`
	AddOptions          *AddMovieOptions `json:"addOptions"`
	Tags                []int            `json:"tags,omitempty"`
	Monitored           bool             `json:"monitored"`
}

// AddMovieOptions are the options for finding a new movie.
type AddMovieOptions struct {
	SearchForMovie bool `json:"searchForMovie"`
	// Allowed values: "movieOnly", "movieAndCollection", "none"
	Monitor string `json:"monitor,omitempty"`
}

// AlternativeTitle is part of a Movie.
type AlternativeTitle struct {
	MovieMetadataID int64        `json:"movieMetadataId"`
	MovieID         int64        `json:"movieId"`
	Title           string       `json:"title"`
	SourceType      string       `json:"sourceType"`
	SourceID        int64        `json:"sourceId"`
	Votes           int          `json:"votes"`
	VoteCount       int          `json:"voteCount"`
	Language        *starr.Value `json:"language"`
	ID              int64        `json:"id"`
}

// GetMovie grabs a movie from the queue, or all movies if tmdbId is 0.
func (r *Radarr) GetMovie(tmdbID int64, excludeLocalCovers bool) ([]*Movie, error) {
	return r.GetMovieContext(context.Background(), tmdbID, excludeLocalCovers)
}

// GetMovieContext grabs a movie from the queue, or all movies if tmdbId is 0.
// excludeLocalCovers is only applicable to all movies endpoint
func (r *Radarr) GetMovieContext(ctx context.Context, tmdbID int64, excludeLocalCovers bool) ([]*Movie, error) {
	params := make(url.Values)
	if tmdbID != 0 {
        	params.Set("tmdbId", fmt.Sprint(tmdbID))
	} else {
        	params.Set("excludeLocalCovers", fmt.Sprint(excludeLocalCovers))
	}

	var output []*Movie

	req := starr.Request{URI: bpMovie, Query: params}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetMovieByID grabs a movie from the database by DB [movie] ID.
func (r *Radarr) GetMovieByID(movieID int64) (*Movie, error) {
	return r.GetMovieByIDContext(context.Background(), movieID)
}

// GetMovieByIDContext grabs a movie from the database by DB [movie] ID.
func (r *Radarr) GetMovieByIDContext(ctx context.Context, movieID int64) (*Movie, error) {
	var output Movie

	req := starr.Request{URI: path.Join(bpMovie, fmt.Sprint(movieID))}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateMovie sends a PUT request to update a movie in place.
func (r *Radarr) UpdateMovie(movieID int64, movie *Movie, moveFiles bool) (*Movie, error) {
	return r.UpdateMovieContext(context.Background(), movieID, movie, moveFiles)
}

// UpdateMovieContext sends a PUT request to update a movie in place.
func (r *Radarr) UpdateMovieContext(ctx context.Context, movieID int64, movie *Movie, moveFiles bool) (*Movie, error) {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(movie); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpMovie, err)
	}

	var output Movie

	req := starr.Request{
		URI:   path.Join(bpMovie, fmt.Sprint(movieID)),
		Query: make(url.Values),
		Body:  &body,
	}
	req.Query.Add("moveFiles", fmt.Sprint(moveFiles))

	if err := r.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// AddMovie adds a movie to the queue.
func (r *Radarr) AddMovie(movie *AddMovieInput) (*Movie, error) {
	return r.AddMovieContext(context.Background(), movie)
}

// AddMovieContext adds a movie to the queue.
func (r *Radarr) AddMovieContext(ctx context.Context, movie *AddMovieInput) (*Movie, error) {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(movie); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpMovie, err)
	}

	var output Movie

	req := starr.Request{URI: bpMovie, Query: make(url.Values), Body: &body}
	if err := r.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
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

	req := starr.Request{URI: path.Join(bpMovie, "lookup"), Query: make(url.Values)}
	req.Query.Set("term", term)

	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// LookupID will return a movie by its ID.
func (r *Radarr) LookupID(movieID int64) (*Movie, error) {
	return r.LookupIDContext(context.Background(), movieID)
}

// LookupIDContext will return a movie by its ID using a context.
func (r *Radarr) LookupIDContext(ctx context.Context, movieID int64) (*Movie, error) {
	return r.lookupSubContext(ctx, fmt.Sprint(movieID), "", "")
}

// LookupIMDB will search IMDB for the imdbId provided.
func (r *Radarr) LookupIMDB(imdbID string) (*Movie, error) {
	return r.LookupIMDBContext(context.Background(), imdbID)
}

// LookupIMDBContext will search IMDB for the imdbId provided using a context.
func (r *Radarr) LookupIMDBContext(ctx context.Context, imdbID string) (*Movie, error) {
	return r.lookupSubContext(ctx, "imdb", "imdbId", imdbID)
}

// LookupTMDB will search TMDB for the tmdbID provided.
func (r *Radarr) LookupTMDB(tmdbID int64) (*Movie, error) {
	return r.LookupTMDBContext(context.Background(), tmdbID)
}

// LookupTMDBContext will search TMDB for the tmdbID provided using a context.
func (r *Radarr) LookupTMDBContext(ctx context.Context, tmdbID int64) (*Movie, error) {
	return r.lookupSubContext(ctx, "tmdb", "tmdbId", fmt.Sprint(tmdbID))
}

// lookupSubContext abstracts lookup requests.
func (r *Radarr) lookupSubContext(ctx context.Context, sub, name, val string) (*Movie, error) {
	var output *Movie

	req := starr.Request{URI: path.Join(bpMovie, "lookup", sub), Query: make(url.Values)}

	if name != "" {
		req.Query.Set(name, val)
	}

	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// DeleteMovie removes a movie from the database. Setting deleteFiles true will delete all content for the movie.
func (r *Radarr) DeleteMovie(movieID int64, deleteFiles, addImportExclusion bool) error {
	return r.DeleteMovieContext(context.Background(), movieID, deleteFiles, addImportExclusion)
}

// DeleteMovieContext removes a movie from the database. Setting deleteFiles true will delete all content for the movie.
func (r *Radarr) DeleteMovieContext(ctx context.Context, movieID int64, deleteFiles, addImportExclusion bool) error {
	req := starr.Request{URI: path.Join(bpMovie, fmt.Sprint(movieID)), Query: make(url.Values)}
	req.Query.Set("deleteFiles", fmt.Sprint(deleteFiles))
	req.Query.Set("addImportExclusion", fmt.Sprint(addImportExclusion))

	if err := r.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
