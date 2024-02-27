package lidarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"

	"golift.io/starr"
)

const bpQualityDefinition = APIver + "/qualitydefinition"

// QualityDefinition is the /api/v1/qualitydefinition endpoint.
type QualityDefinition struct {
	ID            int64        `json:"id"`
	Quality       *starr.Value `json:"quality"`
	Title         string       `json:"title"`
	Weight        int64        `json:"weight"`
	MinSize       float64      `json:"minSize"`
	MaxSize       float64      `json:"maxSize,omitempty"`
	PreferredSize float64      `json:"preferredSize"`
}

// GetQualityDefinitions returns all configured quality definitions.
func (l *Lidarr) GetQualityDefinitions() ([]*QualityDefinition, error) {
	return l.GetQualityDefinitionsContext(context.Background())
}

// GetQualityDefinitionsContext returns all configured quality definitions.
func (l *Lidarr) GetQualityDefinitionsContext(ctx context.Context) ([]*QualityDefinition, error) {
	var output []*QualityDefinition

	req := starr.Request{URI: bpQualityDefinition}
	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetQualityDefinition returns a single quality definition.
func (l *Lidarr) GetQualityDefinition(qualityDefinitionID int64) (*QualityDefinition, error) {
	return l.GetQualityDefinitionContext(context.Background(), qualityDefinitionID)
}

// GetQualityDefinitionContext returns a single quality definition.
func (l *Lidarr) GetQualityDefinitionContext(ctx context.Context, qdID int64) (*QualityDefinition, error) {
	var output QualityDefinition

	req := starr.Request{URI: path.Join(bpQualityDefinition, starr.Itoa(qdID))}
	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateQualityDefinition updates a quality definition.
func (l *Lidarr) UpdateQualityDefinition(definition *QualityDefinition) (*QualityDefinition, error) {
	return l.UpdateQualityDefinitionContext(context.Background(), definition)
}

// UpdateQualityDefinitionContext updates a quality definition.
func (l *Lidarr) UpdateQualityDefinitionContext(
	ctx context.Context,
	definition *QualityDefinition,
) (*QualityDefinition, error) {
	var output QualityDefinition

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(definition); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpQualityDefinition, err)
	}

	req := starr.Request{URI: path.Join(bpQualityDefinition, starr.Itoa(definition.ID)), Body: &body}
	if err := l.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateQualityDefinitions updates all quality definitions.
func (l *Lidarr) UpdateQualityDefinitions(definition []*QualityDefinition) ([]*QualityDefinition, error) {
	return l.UpdateQualityDefinitionsContext(context.Background(), definition)
}

// UpdateQualityDefinitionsContext updates all quality definitions.
func (l *Lidarr) UpdateQualityDefinitionsContext(
	ctx context.Context,
	definition []*QualityDefinition,
) ([]*QualityDefinition, error) {
	var output []*QualityDefinition

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(definition); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpQualityDefinition, err)
	}

	req := starr.Request{URI: path.Join(bpQualityDefinition, "update"), Body: &body}
	if err := l.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return output, nil
}
