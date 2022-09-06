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
	MinimumAvailability   string              `json:"minimumAvailability,omitempty"`
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
	Ratings               *starr.Ratings      `json:"ratings,omitempty"`
	MovieFile             *MovieFile          `json:"movieFile,omitempty"`
	Collection            *Collection         `json:"collection,omitempty"`
	HasFile               bool                `json:"hasFile,omitempty"`
	IsAvailable           bool                `json:"isAvailable,omitempty"`
	Monitored             bool                `json:"monitored"`
}

// Collection belongs to a Movie.
type Collection struct {
	Name   string         `json:"name"`
	TmdbID int64          `json:"tmdbId"`
	Images []*starr.Image `json:"images"`
}

// MovieFile is part of a Movie.
type MovieFile struct {
	ID                  int64          `json:"id"`
	MovieID             int64          `json:"movieId"`
	RelativePath        string         `json:"relativePath"`
	Path                string         `json:"path"`
	Size                int64          `json:"size"`
	DateAdded           time.Time      `json:"dateAdded"`
	SceneName           string         `json:"sceneName"`
	IndexerFlags        int64          `json:"indexerFlags"`
	Quality             *starr.Quality `json:"quality"`
	MediaInfo           *MediaInfo     `json:"mediaInfo"`
	QualityCutoffNotMet bool           `json:"qualityCutoffNotMet"`
	Languages           []*starr.Value `json:"languages"`
	ReleaseGroup        string         `json:"releaseGroup"`
	Edition             string         `json:"edition"`
}

// MediaInfo is part of a MovieFile.
type MediaInfo struct {
	AudioAdditionalFeatures string  `json:"audioAdditionalFeatures"`
	AudioBitrate            int     `json:"audioBitrate"`
	AudioChannels           float64 `json:"audioChannels"`
	AudioCodec              string  `json:"audioCodec"`
	AudioLanguages          string  `json:"audioLanguages"`
	AudioStreamCount        int     `json:"audioStreamCount"`
	VideoBitDepth           int     `json:"videoBitDepth"`
	VideoBitrate            int     `json:"videoBitrate"`
	VideoCodec              string  `json:"videoCodec"`
	VideoFps                float64 `json:"videoFps"`
	Resolution              string  `json:"resolution"`
	RunTime                 string  `json:"runTime"`
	ScanType                string  `json:"scanType"`
	Subtitles               string  `json:"subtitles"`
}

// AddMovieInput is the input for a new movie.
type AddMovieInput struct {
	Title               string           `json:"title,omitempty"`
	TitleSlug           string           `json:"titleSlug,omitempty"`
	MinimumAvailability string           `json:"minimumAvailability,omitempty"`
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
}

// AddMovieOutput is the data returned when adding a movier.
type AddMovieOutput struct {
	ID                    int64               `json:"id"`
	Title                 string              `json:"title"`
	OriginalTitle         string              `json:"originalTitle"`
	AlternateTitles       []*AlternativeTitle `json:"alternateTitles"`
	SecondaryYearSourceID int64               `json:"secondaryYearSourceId"`
	SortTitle             string              `json:"sortTitle"`
	SizeOnDisk            int                 `json:"sizeOnDisk"`
	Status                string              `json:"status"`
	Overview              string              `json:"overview"`
	InCinemas             time.Time           `json:"inCinemas"`
	DigitalRelease        time.Time           `json:"digitalRelease"`
	Images                []*starr.Image      `json:"images"`
	Website               string              `json:"website"`
	Year                  int                 `json:"year"`
	YouTubeTrailerID      string              `json:"youTubeTrailerId"`
	Studio                string              `json:"studio"`
	Path                  string              `json:"path"`
	QualityProfileID      int64               `json:"qualityProfileId"`
	MinimumAvailability   string              `json:"minimumAvailability"`
	FolderName            string              `json:"folderName"`
	Runtime               int                 `json:"runtime"`
	CleanTitle            string              `json:"cleanTitle"`
	ImdbID                string              `json:"imdbId"`
	TmdbID                int64               `json:"tmdbId"`
	TitleSlug             string              `json:"titleSlug"`
	Genres                []string            `json:"genres"`
	Tags                  []int               `json:"tags"`
	Added                 time.Time           `json:"added"`
	AddOptions            *AddMovieOptions    `json:"addOptions"`
	Ratings               *starr.Ratings      `json:"ratings"`
	HasFile               bool                `json:"hasFile"`
	Monitored             bool                `json:"monitored"`
	IsAvailable           bool                `json:"isAvailable"`
}

// AlternativeTitle is part of a Movie.
type AlternativeTitle struct {
	MovieID    int          `json:"movieId"`
	Title      string       `json:"title"`
	SourceType string       `json:"sourceType"`
	SourceID   int          `json:"sourceId"`
	Votes      int          `json:"votes"`
	VoteCount  int          `json:"voteCount"`
	Language   *starr.Value `json:"language"`
	ID         int          `json:"id"`
}

// GetMovie grabs a movie from the queue, or all movies if tmdbId is 0.
func (r *Radarr) GetMovie(tmdbID int64) ([]*Movie, error) {
	return r.GetMovieContext(context.Background(), tmdbID)
}

// GetMovieContext grabs a movie from the queue, or all movies if tmdbId is 0.
func (r *Radarr) GetMovieContext(ctx context.Context, tmdbID int64) ([]*Movie, error) {
	params := make(url.Values)
	if tmdbID != 0 {
		params.Set("tmdbId", fmt.Sprint(tmdbID))
	}

	var output []*Movie

	req := starr.Request{URI: bpMovie, Query: params}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", req, err)
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
		return nil, fmt.Errorf("api.Get(%s): %w", req, err)
	}

	return &output, nil
}

// UpdateMovie sends a PUT request to update a movie in place.
func (r *Radarr) UpdateMovie(movieID int64, movie *Movie) error {
	return r.UpdateMovieContext(context.Background(), movieID, movie)
}

// UpdateMovieContext sends a PUT request to update a movie in place.
func (r *Radarr) UpdateMovieContext(ctx context.Context, movieID int64, movie *Movie) error {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(movie); err != nil {
		return fmt.Errorf("json.Marshal(%s): %w", bpMovie, err)
	}

	var output interface{} // not sure what this looks like.

	req := starr.Request{
		URI:   path.Join(bpMovie, fmt.Sprint(movieID)),
		Query: make(url.Values),
		Body:  &body,
	}
	req.Query.Add("moveFiles", "true")

	if err := r.PutInto(ctx, req, &output); err != nil {
		return fmt.Errorf("api.Put(%s): %w", req, err)
	}

	return nil
}

// AddMovie adds a movie to the queue.
func (r *Radarr) AddMovie(movie *AddMovieInput) (*AddMovieOutput, error) {
	return r.AddMovieContext(context.Background(), movie)
}

// AddMovieContext adds a movie to the queue.
func (r *Radarr) AddMovieContext(ctx context.Context, movie *AddMovieInput) (*AddMovieOutput, error) {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(movie); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpMovie, err)
	}

	var output AddMovieOutput

	req := starr.Request{URI: bpMovie, Query: make(url.Values), Body: &body}
	req.Query.Add("moveFiles", "true")

	if err := r.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", req, err)
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
		return nil, fmt.Errorf("api.Get(%s): %w", req, err)
	}

	return output, nil
}
