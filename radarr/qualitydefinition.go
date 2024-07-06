package radarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"

	"golift.io/starr"
)

// QualityDefinition is the /api/v3/qualitydefinition endpoint.
type QualityDefinition struct {
	ID       int64              `json:"id,omitempty"`
	Weight   int64              `json:"weight"` // This should not be changed.
	MinSize  float64            `json:"minSize"`
	MaxSize  float64            `json:"maxSize"`
	PrefSize float64            `json:"preferredSize"`
	Title    string             `json:"title"`
	Quality  *starr.BaseQuality `json:"quality"`
}

// Define Base Path for Quality Definition calls.
const bpQualityDefinition = APIver + "/qualityDefinition"

// GetQualityDefinitions returns all configured quality definitions.
func (r *Radarr) GetQualityDefinitions() ([]*QualityDefinition, error) {
	return r.GetQualityDefinitionsContext(context.Background())
}

// GetQualityDefinitionsContext returns all configured quality definitions.
func (r *Radarr) GetQualityDefinitionsContext(ctx context.Context) ([]*QualityDefinition, error) {
	var output []*QualityDefinition

	req := starr.Request{URI: bpQualityDefinition}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetQualityDefinition returns a single quality definition.
func (r *Radarr) GetQualityDefinition(qualityDefinitionID int64) (*QualityDefinition, error) {
	return r.GetQualityDefinitionContext(context.Background(), qualityDefinitionID)
}

// GetQualityDefinitionContext returns a single quality definition.
func (r *Radarr) GetQualityDefinitionContext(ctx context.Context, qdID int64) (*QualityDefinition, error) {
	var output QualityDefinition

	req := starr.Request{URI: path.Join(bpQualityDefinition, starr.Str(qdID))}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateQualityDefinition updates a quality definition.
func (r *Radarr) UpdateQualityDefinition(definition *QualityDefinition) (*QualityDefinition, error) {
	return r.UpdateQualityDefinitionContext(context.Background(), definition)
}

// UpdateQualityDefinitionContext updates a quality definition.
func (r *Radarr) UpdateQualityDefinitionContext(
	ctx context.Context,
	definition *QualityDefinition,
) (*QualityDefinition, error) {
	var output QualityDefinition

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(definition); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpQualityDefinition, err)
	}

	req := starr.Request{URI: path.Join(bpQualityDefinition, starr.Str(definition.ID)), Body: &body}
	if err := r.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateQualityDefinitions updates all quality definitions.
func (r *Radarr) UpdateQualityDefinitions(definition []*QualityDefinition) ([]*QualityDefinition, error) {
	return r.UpdateQualityDefinitionsContext(context.Background(), definition)
}

// UpdateQualityDefinitionsContext updates all quality definitions.
func (r *Radarr) UpdateQualityDefinitionsContext(
	ctx context.Context,
	definition []*QualityDefinition,
) ([]*QualityDefinition, error) {
	var output []*QualityDefinition

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(definition); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpQualityDefinition, err)
	}

	req := starr.Request{URI: path.Join(bpQualityDefinition, "update"), Body: &body}
	if err := r.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return output, nil
}
