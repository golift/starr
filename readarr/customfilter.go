package readarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"

	"golift.io/starr"
	"golift.io/starr/starrshared"
)

const bpCustomFilter = APIver + "/customfilter"

// CustomFilter is the /api/v1/customfilter resource.
type CustomFilter = starrshared.CustomFilter

// GetCustomFilters returns all custom filters.
func (r *Readarr) GetCustomFilters() ([]*CustomFilter, error) {
	return r.GetCustomFiltersContext(context.Background())
}

// GetCustomFiltersContext returns all custom filters.
func (r *Readarr) GetCustomFiltersContext(ctx context.Context) ([]*CustomFilter, error) {
	var output []*CustomFilter

	req := starr.Request{URI: bpCustomFilter}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetCustomFilter returns a single custom filter.
func (r *Readarr) GetCustomFilter(id int) (*CustomFilter, error) {
	return r.GetCustomFilterContext(context.Background(), id)
}

// GetCustomFilterContext returns a single custom filter.
func (r *Readarr) GetCustomFilterContext(ctx context.Context, id int) (*CustomFilter, error) {
	var output CustomFilter

	req := starr.Request{URI: path.Join(bpCustomFilter, starr.Str(id))}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// AddCustomFilter creates a custom filter.
func (r *Readarr) AddCustomFilter(in *CustomFilter) (*CustomFilter, error) {
	return r.AddCustomFilterContext(context.Background(), in)
}

// AddCustomFilterContext creates a custom filter.
func (r *Readarr) AddCustomFilterContext(ctx context.Context, in *CustomFilter) (*CustomFilter, error) {
	var output CustomFilter

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(in); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpCustomFilter, err)
	}

	req := starr.Request{URI: bpCustomFilter, Body: &body}
	if err := r.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateCustomFilter updates a custom filter.
func (r *Readarr) UpdateCustomFilter(in *CustomFilter) (*CustomFilter, error) {
	return r.UpdateCustomFilterContext(context.Background(), in)
}

// UpdateCustomFilterContext updates a custom filter.
func (r *Readarr) UpdateCustomFilterContext(ctx context.Context, input *CustomFilter) (*CustomFilter, error) {
	var output CustomFilter

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(input); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpCustomFilter, err)
	}

	req := starr.Request{URI: path.Join(bpCustomFilter, starr.Str(input.ID)), Body: &body}
	if err := r.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteCustomFilter deletes a custom filter.
func (r *Readarr) DeleteCustomFilter(id int) error {
	return r.DeleteCustomFilterContext(context.Background(), id)
}

// DeleteCustomFilterContext deletes a custom filter.
func (r *Readarr) DeleteCustomFilterContext(ctx context.Context, id int) error {
	req := starr.Request{URI: path.Join(bpCustomFilter, starr.Str(id))}
	if err := r.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
