package lidarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"golift.io/starr"
)

// Define Base Path for Naming calls.
const bpNaming = APIver + "/config/naming"

// CRF is ColonReplacementFormat, for naming config.
type CRF int

// These are all of the possible Colon Replacement Formats (for naming config) in Lidarr.
const (
	ColonDelete CRF = iota
	ColonReplaceWithDash
	ColonReplaceWithSpaceDash
	ColonReplaceWithSpaceDashSpace
	ColonSmartReplace
)

// Naming represents the config/naming endpoint in Lidarr.
type Naming struct {
	RenameTracks             bool   `json:"renameTracks"`
	ReplaceIllegalCharacters bool   `json:"replaceIllegalCharacters"`
	IncludeArtistName        bool   `json:"includeArtistName"`
	IncludeAlbumTitle        bool   `json:"includeAlbumTitle"`
	IncludeQuality           bool   `json:"includeQuality"`
	ReplaceSpaces            bool   `json:"replaceSpaces"`
	ColonReplacementFormat   CRF    `json:"colonReplacementFormat"`
	ID                       int64  `json:"id"`
	StandardTrackFormat      string `json:"standardTrackFormat"`
	MultiDiscTrackFormat     string `json:"multiDiscTrackFormat"`
	ArtistFolderFormat       string `json:"artistFolderFormat"`
}

// GetNaming returns the file naming rules.
func (l *Lidarr) GetNaming() (*Naming, error) {
	return l.GetNamingContext(context.Background())
}

// GetNamingContext returns the file naming rules.
func (l *Lidarr) GetNamingContext(ctx context.Context) (*Naming, error) {
	var output Naming

	req := starr.Request{URI: bpNaming}
	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateNaming updates the file naming rules.
func (l *Lidarr) UpdateNaming(naming *Naming) (*Naming, error) {
	return l.UpdateNamingContext(context.Background(), naming)
}

// UpdateNamingContext updates the file naming rules.
func (l *Lidarr) UpdateNamingContext(ctx context.Context, naming *Naming) (*Naming, error) {
	var output Naming

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(naming); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpNaming, err)
	}

	req := starr.Request{URI: bpNaming, Body: &body}
	if err := l.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}
