package lidarr

import (
	"context"
	"fmt"

	"golift.io/starr"
)

const bpQualityDefinition = APIver + "/qualitydefinition"

// QualityDefinition is the /api/v1/qualitydefinition endpoint.
type QualityDefinition struct {
	ID      int64        `json:"id"`
	Quality *starr.Value `json:"quality"`
	Title   string       `json:"title"`
	Weight  int64        `json:"weight"`
	MinSize float64      `json:"minSize"`
	MaxSize float64      `json:"maxSize,omitempty"`
}

// GetQualityDefinition returns the Quality Definitions.
func (l *Lidarr) GetQualityDefinition() ([]*QualityDefinition, error) {
	return l.GetQualityDefinitionContext(context.Background())
}

// GetQualityDefinitionContext returns the Quality Definitions.
func (l *Lidarr) GetQualityDefinitionContext(ctx context.Context) ([]*QualityDefinition, error) {
	var output []*QualityDefinition

	req := starr.Request{URI: bpQualityDefinition}
	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}
