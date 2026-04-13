package radarr

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

// AutoTagging is the /api/v3/autotagging resource.
type AutoTagging = starrshared.AutoTagging

// AutoTaggingSpecification is one rule inside an AutoTagging definition.
type AutoTaggingSpecification = starrshared.AutoTaggingSpecification

// GetAutoTaggings returns all auto tagging configurations.
func (r *Radarr) GetAutoTaggings() ([]*AutoTagging, error) {
	return r.GetAutoTaggingsContext(context.Background())
}

// GetAutoTaggingsContext returns all auto tagging configurations.
func (r *Radarr) GetAutoTaggingsContext(ctx context.Context) ([]*AutoTagging, error) {
	var output []*AutoTagging

	req := starr.Request{URI: bpAutoTagging}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetAutoTagging returns a single auto tagging configuration.
func (r *Radarr) GetAutoTagging(id int) (*AutoTagging, error) {
	return r.GetAutoTaggingContext(context.Background(), id)
}

// GetAutoTaggingContext returns a single auto tagging configuration.
func (r *Radarr) GetAutoTaggingContext(ctx context.Context, id int) (*AutoTagging, error) {
	var output AutoTagging

	req := starr.Request{URI: path.Join(bpAutoTagging, starr.Str(id))}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// GetAutoTaggingSchema returns the specification schema templates for auto tagging.
func (r *Radarr) GetAutoTaggingSchema() ([]*AutoTaggingSpecification, error) {
	return r.GetAutoTaggingSchemaContext(context.Background())
}

// GetAutoTaggingSchemaContext returns the specification schema templates for auto tagging.
func (r *Radarr) GetAutoTaggingSchemaContext(ctx context.Context) ([]*AutoTaggingSpecification, error) {
	var output []*AutoTaggingSpecification

	req := starr.Request{URI: path.Join(bpAutoTagging, "schema")}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// AddAutoTagging creates an auto tagging configuration.
func (r *Radarr) AddAutoTagging(in *AutoTagging) (*AutoTagging, error) {
	return r.AddAutoTaggingContext(context.Background(), in)
}

// AddAutoTaggingContext creates an auto tagging configuration.
func (r *Radarr) AddAutoTaggingContext(ctx context.Context, in *AutoTagging) (*AutoTagging, error) {
	var output AutoTagging

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(in); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpAutoTagging, err)
	}

	req := starr.Request{URI: bpAutoTagging, Body: &body}
	if err := r.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateAutoTagging updates an auto tagging configuration.
func (r *Radarr) UpdateAutoTagging(in *AutoTagging) (*AutoTagging, error) {
	return r.UpdateAutoTaggingContext(context.Background(), in)
}

// UpdateAutoTaggingContext updates an auto tagging configuration.
func (r *Radarr) UpdateAutoTaggingContext(ctx context.Context, input *AutoTagging) (*AutoTagging, error) {
	var output AutoTagging

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(input); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpAutoTagging, err)
	}

	req := starr.Request{URI: path.Join(bpAutoTagging, starr.Str(input.ID)), Body: &body}
	if err := r.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteAutoTagging deletes an auto tagging configuration.
func (r *Radarr) DeleteAutoTagging(id int) error {
	return r.DeleteAutoTaggingContext(context.Background(), id)
}

// DeleteAutoTaggingContext deletes an auto tagging configuration.
func (r *Radarr) DeleteAutoTaggingContext(ctx context.Context, id int) error {
	req := starr.Request{URI: path.Join(bpAutoTagging, starr.Str(id))}
	if err := r.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
