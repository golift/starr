package sonarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"

	"golift.io/starr"
)

// Define Base Path for Quality Definition calls.
const bpQualityDefinition = APIver + "/qualityDefinition"

// QualityDefinition is the /api/v3/qualitydefinition endpoint.
type QualityDefinition struct {
	ID       int64              `json:"id,omitempty"`
	Weight   int64              `json:"weight"` // This should not be changed.
	MinSize  float64            `json:"minSize"`
	MaxSize  float64            `json:"maxSize"`
	PrefSize float64            `json:"preferredSize"` // v4 only.
	Title    string             `json:"title"`
	Quality  *starr.BaseQuality `json:"quality"`
}

// GetQualityDefinitions returns all configured quality definitions.
func (s *Sonarr) GetQualityDefinitions() ([]*QualityDefinition, error) {
	return s.GetQualityDefinitionsContext(context.Background())
}

// GetQualityDefinitionsContext returns all configured quality definitions.
func (s *Sonarr) GetQualityDefinitionsContext(ctx context.Context) ([]*QualityDefinition, error) {
	var output []*QualityDefinition

	req := starr.Request{URI: bpQualityDefinition}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetQualityDefinition returns a single quality definition.
func (s *Sonarr) GetQualityDefinition(qualityDefinitionID int64) (*QualityDefinition, error) {
	return s.GetQualityDefinitionContext(context.Background(), qualityDefinitionID)
}

// GetQualityDefinitionContext returns a single quality definition.
func (s *Sonarr) GetQualityDefinitionContext(ctx context.Context, qdID int64) (*QualityDefinition, error) {
	var output QualityDefinition

	req := starr.Request{URI: path.Join(bpQualityDefinition, starr.Str(qdID))}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateQualityDefinition updates a quality definition.
func (s *Sonarr) UpdateQualityDefinition(definition *QualityDefinition) (*QualityDefinition, error) {
	return s.UpdateQualityDefinitionContext(context.Background(), definition)
}

// UpdateQualityDefinitionContext updates a quality definition.
func (s *Sonarr) UpdateQualityDefinitionContext(
	ctx context.Context,
	definition *QualityDefinition,
) (*QualityDefinition, error) {
	var output QualityDefinition

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(definition); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpQualityDefinition, err)
	}

	req := starr.Request{URI: path.Join(bpQualityDefinition, starr.Str(definition.ID)), Body: &body}
	if err := s.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateQualityDefinition updates all quality definitions.
func (s *Sonarr) UpdateQualityDefinitions(definitions []*QualityDefinition) ([]*QualityDefinition, error) {
	return s.UpdateQualityDefinitionsContext(context.Background(), definitions)
}

// UpdateQualityDefinitionsContext updates all quality definitions.
func (s *Sonarr) UpdateQualityDefinitionsContext(
	ctx context.Context,
	definitions []*QualityDefinition,
) ([]*QualityDefinition, error) {
	var output []*QualityDefinition

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(definitions); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpQualityDefinition, err)
	}

	req := starr.Request{URI: path.Join(bpQualityDefinition, "update"), Body: &body}
	if err := s.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return output, nil
}
