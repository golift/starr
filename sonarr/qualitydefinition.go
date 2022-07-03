package sonarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"
	"strconv"

	"golift.io/starr"
)

// QualityDefinition is the /api/v3/qualitydefinition endpoint.
type QualityDefinition struct {
	ID        int64            `json:"id,omitempty"`
	Weight    int64            `json:"weight"`
	MinSize   float64          `json:"minSize"`
	MaxSize   float64          `json:"maxSize"`
	Title     string           `json:"title"`
	Qualities []*starr.Quality `json:"items"`
}

// Define Base Path for Quality Definition calls.
const bpQualityDefinition = APIver + "/qualityDefinition"

// GetQualityDefinitions returns all configured quality definitions.
func (s *Sonarr) GetQualityDefinitions() ([]*QualityDefinition, error) {
	return s.GetQualityDefinitionsContext(context.Background())
}

func (s *Sonarr) GetQualityDefinitionsContext(ctx context.Context) ([]*QualityDefinition, error) {
	var output []*QualityDefinition

	if _, err := s.GetInto(ctx, bpQualityDefinition, nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(qualityDefinition): %w", err)
	}

	return output, nil
}

// GetQualityDefinition returns a single quality definition.
func (s *Sonarr) GetQualityDefinition(qualityDefinitionID int) (*QualityDefinition, error) {
	return s.GetQualityDefinitionContext(context.Background(), qualityDefinitionID)
}

func (s *Sonarr) GetQualityDefinitionContext(ctx context.Context, qualityDefinitionID int) (*QualityDefinition, error) {
	var output *QualityDefinition

	uri := path.Join(bpQualityDefinition, strconv.Itoa(qualityDefinitionID))
	if _, err := s.GetInto(ctx, uri, nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(qualityDefinition): %w", err)
	}

	return output, nil
}

// UpdateQualityDefinition updates the quality definition.
func (s *Sonarr) UpdateQualityDefinition(definition *QualityDefinition) (*QualityDefinition, error) {
	return s.UpdateQualityDefinitionContext(context.Background(), definition)
}

func (s *Sonarr) UpdateQualityDefinitionContext(
	ctx context.Context,
	definition *QualityDefinition,
) (*QualityDefinition, error) {
	var output QualityDefinition

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(definition); err != nil {
		return nil, fmt.Errorf("json.Marshal(qualityDefinition): %w", err)
	}

	uri := path.Join(bpQualityDefinition, strconv.Itoa(int(definition.ID)))
	if _, err := s.PutInto(ctx, uri, nil, &body, &output); err != nil {
		return nil, fmt.Errorf("api.Put(qualityDefinition): %w", err)
	}

	return &output, nil
}
