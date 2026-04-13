package lidarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"

	"golift.io/starr"
	"golift.io/starr/starrshared"
)

const bpAutoTagging = APIver + "/autotagging"

// AutoTagging is the /api/v1/autotagging resource.
type AutoTagging = starrshared.AutoTagging

// AutoTaggingSpecification is one rule inside an AutoTagging definition.
type AutoTaggingSpecification = starrshared.AutoTaggingSpecification

// GetAutoTaggings returns all auto tagging configurations.
func (l *Lidarr) GetAutoTaggings() ([]*AutoTagging, error) {
	return l.GetAutoTaggingsContext(context.Background())
}

// GetAutoTaggingsContext returns all auto tagging configurations.
func (l *Lidarr) GetAutoTaggingsContext(ctx context.Context) ([]*AutoTagging, error) {
	var output []*AutoTagging

	req := starr.Request{URI: bpAutoTagging}
	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetAutoTagging returns a single auto tagging configuration.
func (l *Lidarr) GetAutoTagging(id int) (*AutoTagging, error) {
	return l.GetAutoTaggingContext(context.Background(), id)
}

// GetAutoTaggingContext returns a single auto tagging configuration.
func (l *Lidarr) GetAutoTaggingContext(ctx context.Context, id int) (*AutoTagging, error) {
	var output AutoTagging

	req := starr.Request{URI: path.Join(bpAutoTagging, starr.Str(id))}
	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// GetAutoTaggingSchema returns the specification schema templates for auto tagging.
func (l *Lidarr) GetAutoTaggingSchema() ([]*AutoTaggingSpecification, error) {
	return l.GetAutoTaggingSchemaContext(context.Background())
}

// GetAutoTaggingSchemaContext returns the specification schema templates for auto tagging.
func (l *Lidarr) GetAutoTaggingSchemaContext(ctx context.Context) ([]*AutoTaggingSpecification, error) {
	var output []*AutoTaggingSpecification

	req := starr.Request{URI: path.Join(bpAutoTagging, "schema")}
	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// AddAutoTagging creates an auto tagging configuration.
func (l *Lidarr) AddAutoTagging(in *AutoTagging) (*AutoTagging, error) {
	return l.AddAutoTaggingContext(context.Background(), in)
}

// AddAutoTaggingContext creates an auto tagging configuration.
func (l *Lidarr) AddAutoTaggingContext(ctx context.Context, in *AutoTagging) (*AutoTagging, error) {
	var output AutoTagging

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(in); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpAutoTagging, err)
	}

	req := starr.Request{URI: bpAutoTagging, Body: &body}
	if err := l.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateAutoTagging updates an auto tagging configuration.
func (l *Lidarr) UpdateAutoTagging(in *AutoTagging) (*AutoTagging, error) {
	return l.UpdateAutoTaggingContext(context.Background(), in)
}

// UpdateAutoTaggingContext updates an auto tagging configuration.
func (l *Lidarr) UpdateAutoTaggingContext(ctx context.Context, input *AutoTagging) (*AutoTagging, error) {
	var output AutoTagging

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(input); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpAutoTagging, err)
	}

	req := starr.Request{URI: path.Join(bpAutoTagging, starr.Str(input.ID)), Body: &body}
	if err := l.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteAutoTagging deletes an auto tagging configuration.
func (l *Lidarr) DeleteAutoTagging(id int) error {
	return l.DeleteAutoTaggingContext(context.Background(), id)
}

// DeleteAutoTaggingContext deletes an auto tagging configuration.
func (l *Lidarr) DeleteAutoTaggingContext(ctx context.Context, id int) error {
	req := starr.Request{URI: path.Join(bpAutoTagging, starr.Str(id))}
	if err := l.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
