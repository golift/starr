package sonarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"golift.io/starr"
)

// Define Base Path for Naming calls.
const bpNaming = APIver + "/config/naming"

type Naming struct {
	RenameEpisodes           bool   `json:"renameEpisodes,omitempty"`
	ReplaceIllegalCharacters bool   `json:"replaceIllegalCharacters,omitempty"`
	IncludeQuality           bool   `json:"includeQuality,omitempty"`
	IncludeSeriesTitle       bool   `json:"includeSeriesTitle,omitempty"`
	IncludeEpisodeTitle      bool   `json:"includeEpisodeTitle,omitempty"`
	ReplaceSpaces            bool   `json:"replaceSpaces,omitempty"`
	ID                       int64  `json:"id,omitempty"`
	MultiEpisodeStyle        int64  `json:"multiEpisodeStyle,omitempty"`
	Separator                string `json:"separator,omitempty"`
	NumberStyle              string `json:"numberStyle,omitempty"`
	DailyEpisodeFormat       string `json:"dailyEpisodeFormat,omitempty"`
	AnimeEpisodeFormat       string `json:"animeEpisodeFormat,omitempty"`
	SeriesFolderFormat       string `json:"seriesFolderFormat,omitempty"`
	SeasonFolderFormat       string `json:"seasonFolderFormat,omitempty"`
	SpecialsFolderFormat     string `json:"specialsFolderFormat,omitempty"`
	StandardEpisodeFormat    string `json:"standardEpisodeFormat,omitempty"`
}

// GetNaming returns the naming.
func (s *Sonarr) GetNaming() (*Naming, error) {
	return s.GetNamingContext(context.Background())
}

// GetNamingContext returns the naming.
func (s *Sonarr) GetNamingContext(ctx context.Context) (*Naming, error) {
	var output Naming

	req := starr.Request{URI: bpNaming}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateNaming updates the naming.
func (s *Sonarr) UpdateNaming(naming *Naming) (*Naming, error) {
	return s.UpdateNamingContext(context.Background(), naming)
}

// UpdateNamingContext updates the naming.
func (s *Sonarr) UpdateNamingContext(ctx context.Context, naming *Naming) (*Naming, error) {
	var output Naming

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(naming); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpNaming, err)
	}

	req := starr.Request{URI: bpNaming, Body: &body}
	if err := s.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}
