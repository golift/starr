package radarr

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

	if err := r.GetInto(ctx, bpQualityDefinition, nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", bpQualityDefinition, err)
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

	uri := path.Join(bpQualityDefinition, strconv.FormatInt(qdID, starr.Base10))
	if err := r.GetInto(ctx, uri, nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", uri, err)
	}

	return &output, nil
}

// UpdateQualityDefinitions updates the quality definition.
func (r *Radarr) UpdateQualityDefinitions(definition []*QualityDefinition) ([]*QualityDefinition, error) {
	return r.UpdateQualityDefinitionsContext(context.Background(), definition)
}

// UpdateQualityDefinitionsContext updates the quality definition.
func (r *Radarr) UpdateQualityDefinitionsContext(
	ctx context.Context,
	definition []*QualityDefinition,
) ([]*QualityDefinition, error) {
	var output []*QualityDefinition

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(definition); err != nil {
		return nil, fmt.Errorf("json.Marshal(qualityDefinition): %w", err)
	}

	uri := path.Join(bpQualityDefinition, "update")
	if err := r.PutInto(ctx, uri, nil, &body, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", uri, err)
	}

	return output, nil
}
