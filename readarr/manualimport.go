package readarr

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
	ID                      int64          `json:"id"`
	Path                    string         `json:"path"`
	Name                    string         `json:"name"`
	AuthorID                int64          `json:"authorID"`
	BookID                  int64          `json:"bookID"`
	ForeignEditionID        int64          `json:"foreignEditionId"`
	Quality                 *starr.Quality `json:"quality"`
	ReleaseGroup            string         `json:"releaseGroup"`
	DownloadID              string         `json:"downloadId"`
	AdditionalFile          bool           `json:"additionalFile"`
	ReplaceExistingFiles    bool           `json:"replaceExistingFiles"`
	DisableReleaseSwitching bool           `json:"disableReleaseSwitching"`
	Rejections              []*Rejection   `json:"rejections"`
}

// ManualImportOutput is the output data for a manual import request.
type ManualImportOutput struct {
	ID                      int64          `json:"id"`
	Path                    string         `json:"path"`
	Name                    string         `json:"name"`
	Size                    int            `json:"size"`
	Author                  *Author        `json:"author"`
	Book                    *Book          `json:"book"`
	ForeignEditionID        int64          `json:"foreignEditionId"`
	Quality                 *starr.Quality `json:"quality"`
	ReleaseGroup            string         `json:"releaseGroup"`
	QualityWeight           int64          `json:"qualityWeight"`
	DownloadID              string         `json:"downloadId"`
	AudioTags               *AudioTags     `json:"audioTags"`
	AdditionalFile          bool           `json:"additionalFile"`
	ReplaceExistingFiles    bool           `json:"replaceExistingFiles"`
	DisableReleaseSwitching bool           `json:"disableReleaseSwitching"`
	Rejections              []*Rejection   `json:"rejections"`
}

// Rejection is part of the manual import payload.
type Rejection struct {
	Reason string `json:"reason"`
	// permanent or temporary
	Type string `json:"type"`
}

// ManualImportParams provides the input parameters for the GET /manualimport API.
type ManualImportParams struct {
	Folder               string
	DownloadID           string
	AuthorID             int64
	ReplaceExistingFiles bool
	FilterExistingFiles  bool
}

// ManualImport initiates a manual import (GET).
func (r *Readarr) ManualImport(params *ManualImportParams) (*ManualImportOutput, error) {
	return r.ManualImportContext(context.Background(), params)
}

// ManualImportContext initiates a manual import (GET).
func (r *Readarr) ManualImportContext(ctx context.Context, params *ManualImportParams) (*ManualImportOutput, error) {
	req := starr.Request{URI: bpManualImport, Query: make(url.Values)}
	req.Query.Add("folder", params.Folder)
	req.Query.Add("downloadId", params.DownloadID)
	req.Query.Add("authorId", fmt.Sprint(params.AuthorID))
	req.Query.Add("replaceExistingFiles", fmt.Sprint(params.ReplaceExistingFiles))
	req.Query.Add("filterExistingFiles", fmt.Sprint(params.FilterExistingFiles))

	var output ManualImportOutput
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// ManualImportReprocess reprocesses a manual import (POST).
func (r *Readarr) ManualImportReprocess(manualimport *ManualImportInput) error {
	return r.ManualImportReprocessContext(context.Background(), manualimport)
}

// ManualImportReprocessContext reprocesses a manual import (POST).
func (r *Readarr) ManualImportReprocessContext(ctx context.Context, manualimport *ManualImportInput) error {
	var output interface{}

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(manualimport); err != nil {
		return fmt.Errorf("json.Marshal(%s): %w", bpManualImport, err)
	}

	req := starr.Request{URI: bpManualImport, Body: &body}
	if err := r.PostInto(ctx, req, &output); err != nil {
		return fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return nil
}
