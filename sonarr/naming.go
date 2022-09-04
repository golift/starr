package sonarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

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

// Define Base Path for Naming calls.
const bpNaming = APIver + "/config/naming"

// GetNaming returns the naming.
func (s *Sonarr) GetNaming() (*Naming, error) {
	return s.GetNamingContext(context.Background())
}

func (s *Sonarr) GetNamingContext(ctx context.Context) (*Naming, error) {
	var output *Naming

	if err := s.GetInto(ctx, bpNaming, nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(naming): %w", err)
	}

	return output, nil
}

// UpdateNaming updates the naming.
func (s *Sonarr) UpdateNaming(naming *Naming) (*Naming, error) {
	return s.UpdateNamingContext(context.Background(), naming)
}

func (s *Sonarr) UpdateNamingContext(ctx context.Context, naming *Naming) (*Naming, error) {
	var output Naming

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(naming); err != nil {
		return nil, fmt.Errorf("json.Marshal(naming): %w", err)
	}

	if err := s.PutInto(ctx, bpNaming, nil, &body, &output); err != nil {
		return nil, fmt.Errorf("api.Put(naming): %w", err)
	}

	return &output, nil
}
