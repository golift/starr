package readarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"golift.io/starr"
)

// Define Base Path for Naming calls.
const bpNaming = APIver + "/config/naming"

// Naming represents the config/naming endpoint in Readarr.
type Naming struct {
	RenameBooks              bool   `json:"renameBooks"`
	ReplaceIllegalCharacters bool   `json:"replaceIllegalCharacters"`
	IncludeAuthorName        bool   `json:"includeAuthorName"`
	IncludeBookTitle         bool   `json:"includeBookTitle"`
	IncludeQuality           bool   `json:"includeQuality"`
	ReplaceSpaces            bool   `json:"replaceSpaces"`
	ColonReplacementFormat   CRF    `json:"colonReplacementFormat"`
	ID                       int64  `json:"id"`
	StandardBookFormat       string `json:"standardBookFormat"`
	AuthorFolderFormat       string `json:"authorFolderFormat"`
}

// CRF is ColonReplacementFormat, for naming config.
type CRF int

// These are all of the possible Colon Replacement Formats (for naming config) in Readarr.
const (
	ColonDelete CRF = iota
	ColonReplaceWithDash
	ColonReplaceWithSpaceDash
	ColonReplaceWithSpaceDashSpace
	ColonSmartReplace
)

// GetNaming returns the file naming rules.
func (r *Readarr) GetNaming() (*Naming, error) {
	return r.GetNamingContext(context.Background())
}

// GetNamingContext returns the file naming rules.
func (r *Readarr) GetNamingContext(ctx context.Context) (*Naming, error) {
	var output Naming

	req := starr.Request{URI: bpNaming}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateNaming updates the file naming rules.
func (r *Readarr) UpdateNaming(naming *Naming) (*Naming, error) {
	return r.UpdateNamingContext(context.Background(), naming)
}

// UpdateNamingContext updates the file naming rules.
func (r *Readarr) UpdateNamingContext(ctx context.Context, naming *Naming) (*Naming, error) {
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
