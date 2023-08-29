package sonarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"golift.io/starr"
)

// Define Base Path for Manual Import calls.
const bpManualImport = APIver + "/manualimport"

// ManualImportInput is the input data for a manual import request using a POST request.
type ManualImportInput struct {
	ID                int64                `json:"id"`
	Path              string               `json:"path"`
	SeriesID          int64                `json:"seriesId"`
	SeasonNumber      int                  `json:"seasonNumber"`
	Episodes          []*Episode           `json:"episodes"`
	EpisodeIDs        []int64              `json:"episodeIds"`
	Quality           *starr.Quality       `json:"quality"`
	Languages         []*starr.Value       `json:"languages"`
	ReleaseGroup      string               `json:"releaseGroup"`
	DownloadID        string               `json:"downloadId"`
	CustomFormats     []*CustomFormatInput `json:"customFormats"`
	CustomFormatScore int64                `json:"customFormatScore"`
	Rejections        []*Rejection         `json:"rejections"`
}

// ManualImportOutput is the output data for a manual import request.
type ManualImportOutput struct {
	ID                int64                 `json:"id"`
	Path              string                `json:"path"`
	RelativePath      string                `json:"relativePath"`
	FolderName        string                `json:"folderName"`
	Name              string                `json:"name"`
	Size              int                   `json:"size"`
	Series            *Series               `json:"series"`
	SeasonNumber      int                   `json:"seasonNumber"`
	Episodes          []*Episode            `json:"episodes"`
	EpisodeFileID     int64                 `json:"episodeFileId"`
	Quality           *starr.Quality        `json:"quality"`
	Languages         []*starr.Value        `json:"languages"`
	ReleaseGroup      string                `json:"releaseGroup"`
	QualityWeight     int64                 `json:"qualityWeight"`
	DownloadID        string                `json:"downloadId"`
	CustomFormats     []*CustomFormatOutput `json:"customFormats"`
	CustomFormatScore int64                 `json:"customFormatScore"`
	Rejections        []*Rejection          `json:"rejections"`
}

// Rejection is part of the manual import payload.
type Rejection struct {
	Reason string `json:"reason"`
	// permanent or temporary
	Type string `json:"type"`
}

// ManualImportParams provides the input parameters for the GET /manualimport API.
type ManualImportParams struct {
	Folder              string
	DownloadID          string
	SeriesID            int64
	SeasonNumber        int
	FilterExistingFiles bool
}

// ManualImport initiates a manual import (GET).
func (s *Sonarr) ManualImport(params *ManualImportParams) (*ManualImportOutput, error) {
	return s.ManualImportContext(context.Background(), params)
}

// ManualImportContext initiates a manual import (GET).
func (s *Sonarr) ManualImportContext(ctx context.Context, params *ManualImportParams) (*ManualImportOutput, error) {
	var output ManualImportOutput

	req := starr.Request{URI: bpManualImport, Query: make(url.Values)}
	req.Query.Add("folder", params.Folder)
	req.Query.Add("downloadId", params.DownloadID)
	req.Query.Add("seriesId", fmt.Sprint(params.SeriesID))
	req.Query.Add("seasonNumber", fmt.Sprint(params.SeasonNumber))
	req.Query.Add("filterExistingFiles", fmt.Sprint(params.FilterExistingFiles))

	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// ManualImportReprocess reprocesses a manual import (POST).
func (s *Sonarr) ManualImportReprocess(manualimport *ManualImportInput) error {
	return s.ManualImportReprocessContext(context.Background(), manualimport)
}

// ManualImportReprocessContext reprocesses a manual import (POST).
func (s *Sonarr) ManualImportReprocessContext(ctx context.Context, manualimport *ManualImportInput) error {
	var output interface{}

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(manualimport); err != nil {
		return fmt.Errorf("json.Marshal(%s): %w", bpManualImport, err)
	}

	req := starr.Request{URI: bpManualImport, Body: &body}
	if err := s.PostInto(ctx, req, &output); err != nil {
		return fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return nil
}
