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

const bpMovieFile = APIver + "/moviefile"

// MovieFile is part of a Movie.
type MovieFile struct {
	ID                  int64                 `json:"id"`
	MovieID             int64                 `json:"movieId"`
	RelativePath        string                `json:"relativePath"`
	Path                string                `json:"path"`
	Size                int64                 `json:"size"`
	DateAdded           time.Time             `json:"dateAdded"`
	SceneName           string                `json:"sceneName"`
	IndexerFlags        int64                 `json:"indexerFlags"`
	Quality             *starr.Quality        `json:"quality,omitempty"`
	CustomFormats       []*CustomFormatOutput `json:"customFormats,omitempty"`
	CustomFormatScore   int                   `json:"customFormatScore"`
	MediaInfo           *MediaInfo            `json:"mediaInfo,omitempty"`
	OriginalFilePath    string                `json:"originalFilePath"`
	QualityCutoffNotMet bool                  `json:"qualityCutoffNotMet"`
	Languages           []*starr.Value        `json:"languages"`
	ReleaseGroup        string                `json:"releaseGroup"`
	Edition             string                `json:"edition"`
}

// MediaInfo is part of a MovieFile.
type MediaInfo struct {
	ID                    int64   `json:"id"`
	AudioBitrate          int     `json:"audioBitrate"`
	AudioChannels         float64 `json:"audioChannels"`
	AudioCodec            string  `json:"audioCodec"`
	AudioLanguages        string  `json:"audioLanguages"`
	AudioStreamCount      int     `json:"audioStreamCount"`
	VideoBitDepth         int     `json:"videoBitDepth"`
	VideoBitrate          int     `json:"videoBitrate"`
	VideoCodec            string  `json:"videoCodec"`
	VideoDynamicRangeType string  `json:"videoDynamicRangeType"`
	VideoFps              float64 `json:"videoFps"`
	Resolution            string  `json:"resolution"`
	RunTime               string  `json:"runTime"`
	ScanType              string  `json:"scanType"`
	Subtitles             string  `json:"subtitles"`
}

// GetMovieFile returns the movie file(s) for a movie.
func (r *Radarr) GetMovieFile(movieID int64) ([]*MovieFile, error) {
	return r.GetMovieFileContext(context.Background(), movieID)
}

// GetMovieFileContext returns the movie file(s) for a movie.
func (r *Radarr) GetMovieFileContext(ctx context.Context, movieID int64) ([]*MovieFile, error) {
	req := starr.Request{URI: bpMovieFile, Query: make(url.Values)}
	req.Query.Add("movieID", starr.Itoa(movieID))

	var output []*MovieFile
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetMovieFileByID grabs a movie from the database by DB [movieFile] ID.
func (r *Radarr) GetMovieFileByID(movieFileID int64) (*MovieFile, error) {
	return r.GetMovieFileByIDContext(context.Background(), movieFileID)
}

// GetMovieFileByIDContext grabs a movie from the database by DB [movieFile] ID.
func (r *Radarr) GetMovieFileByIDContext(ctx context.Context, movieFileID int64) (*MovieFile, error) {
	var output MovieFile

	req := starr.Request{URI: path.Join(bpMovieFile, starr.Itoa(movieFileID))}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// GetMovieFiles returns the movie file(s) requested.
func (r *Radarr) GetMovieFiles(movieFileIDs []int64) ([]*MovieFile, error) {
	return r.GetMovieFilesContext(context.Background(), movieFileIDs)
}

// GetMovieFilesContext returns the movie file(s) requested.
func (r *Radarr) GetMovieFilesContext(ctx context.Context, movieFileIDs []int64) ([]*MovieFile, error) {
	req := starr.Request{URI: bpMovieFile, Query: make(url.Values)}
	for _, id := range movieFileIDs {
		req.Query.Add("movieFileIds", starr.Itoa(id))
	}

	var output []*MovieFile
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// UpdateMovieFile updates the movie file provided.
func (r *Radarr) UpdateMovieFile(movieFile *MovieFile) (*MovieFile, error) {
	return r.UpdateMovieFileContext(context.Background(), movieFile)
}

// UpdateMovieFileContext updates the movie file provided.
func (r *Radarr) UpdateMovieFileContext(ctx context.Context, movieFile *MovieFile) (*MovieFile, error) {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(movieFile); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpMovieFile, err)
	}

	var output *MovieFile

	req := starr.Request{URI: path.Join(bpMovieFile, starr.Itoa(movieFile.ID)), Body: &body}
	if err := r.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return output, nil
}

// DeleteMovieFile deletes movie files by their IDs.
func (r *Radarr) DeleteMovieFiles(movieFileIDs ...int64) error {
	return r.DeleteMovieFilesContext(context.Background(), movieFileIDs...)
}

// DeleteMovieFileContext deletes movie files by their IDs.
func (r *Radarr) DeleteMovieFilesContext(ctx context.Context, movieFileIDs ...int64) error {
	postData := struct {
		T []int64 `json:"movieFileIds"`
	}{movieFileIDs}

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(&postData); err != nil {
		return fmt.Errorf("json.Marshal(%s): %w", bpMovieFile, err)
	}

	req := starr.Request{URI: path.Join(bpMovieFile, "bulk"), Body: &body}
	if err := r.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
