package radarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"golift.io/starr"
)

// Define Base Path for Naming calls.
const bpNaming = APIver + "/config/naming"

// Naming represents the config/naming endpoint in Radarr.
type Naming struct {
	RenameMovies             bool   `json:"renameMovies,omitempty"`
	ReplaceIllegalCharacters bool   `json:"replaceIllegalCharacters,omitempty"`
	ID                       int64  `json:"id"` // ID must always be 1 (Oct 10, 2022)
	ColonReplacementFormat   CRF    `json:"colonReplacementFormat,omitempty"`
	StandardMovieFormat      string `json:"standardMovieFormat"` // required
	MovieFolderFormat        string `json:"movieFolderFormat"`   // required
}

// CRF is ColonReplacementFormat, for naming config.
type CRF string

// These are all of the possible Colon Replacement Formats (for naming config) in Radarr.
const (
	ColonDelete                    CRF = "delete"
	ColonReplaceWithDash           CRF = "dash"
	ColonReplaceWithSpaceDash      CRF = "spaceDash"
	ColonReplaceWithSpaceDashSpace CRF = "spaceDashSpace"
)

// GetNaming returns the file naming rules.
func (r *Radarr) GetNaming() (*Naming, error) {
	return r.GetNamingContext(context.Background())
}

// GetNamingContext returns the file naming rules.
func (r *Radarr) GetNamingContext(ctx context.Context) (*Naming, error) {
	var output Naming

	req := starr.Request{URI: bpNaming}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateNaming updates the file naming rules.
func (r *Radarr) UpdateNaming(naming *Naming) (*Naming, error) {
	return r.UpdateNamingContext(context.Background(), naming)
}

// UpdateNamingContext updates the file naming rules.
func (r *Radarr) UpdateNamingContext(ctx context.Context, naming *Naming) (*Naming, error) {
	var (
		output Naming
		body   bytes.Buffer
	)

	naming.ID = 1
	if err := json.NewEncoder(&body).Encode(naming); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpNaming, err)
	}

	req := starr.Request{URI: bpNaming, Body: &body}
	if err := r.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}
