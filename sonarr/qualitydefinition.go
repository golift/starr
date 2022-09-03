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
	ID       int64          `json:"id,omitempty"`
	Weight   int64          `json:"weight"` // This should not be changed.
	MinSize  float64        `json:"minSize"`
	MaxSize  float64        `json:"maxSize"`
	PrefSize float64        `json:"preferredSize"` // v4 only.
	Title    string         `json:"title"`
	Quality  *starr.Quality `json:"quality"`
}

// Define Base Path for Quality Definition calls.
const bpQualityDefinition = APIver + "/qualityDefinition"

// GetQualityDefinitions returns all configured quality definitions.
func (s *Sonarr) GetQualityDefinitions() ([]*QualityDefinition, error) {
	return s.GetQualityDefinitionsContext(context.Background())
}

func (s *Sonarr) GetQualityDefinitionsContext(ctx context.Context) ([]*QualityDefinition, error) {
	var output []*QualityDefinition

	if err := s.GetInto(ctx, bpQualityDefinition, nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(qualityDefinition): %w", err)
	}

	return output, nil
}

// GetQualityDefinition returns a single quality definition.
func (s *Sonarr) GetQualityDefinition(qualityDefinitionID int64) (*QualityDefinition, error) {
	return s.GetQualityDefinitionContext(context.Background(), qualityDefinitionID)
}

func (s *Sonarr) GetQualityDefinitionContext(ctx context.Context, qdID int64) (*QualityDefinition, error) {
	var output *QualityDefinition

	uri := path.Join(bpQualityDefinition, strconv.FormatInt(qdID, starr.Base10))
	if err := s.GetInto(ctx, uri, nil, &output); err != nil {
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
	if err := s.PutInto(ctx, uri, nil, &body, &output); err != nil {
		return nil, fmt.Errorf("api.Put(qualityDefinition): %w", err)
	}

	return &output, nil
}
