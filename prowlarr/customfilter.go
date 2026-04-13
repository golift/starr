package prowlarr

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
func (p *Prowlarr) GetCustomFilters() ([]*CustomFilter, error) {
	return p.GetCustomFiltersContext(context.Background())
}

// GetCustomFiltersContext returns all custom filters.
func (p *Prowlarr) GetCustomFiltersContext(ctx context.Context) ([]*CustomFilter, error) {
	var output []*CustomFilter

	req := starr.Request{URI: bpCustomFilter}
	if err := p.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetCustomFilter returns a single custom filter.
func (p *Prowlarr) GetCustomFilter(id int) (*CustomFilter, error) {
	return p.GetCustomFilterContext(context.Background(), id)
}

// GetCustomFilterContext returns a single custom filter.
func (p *Prowlarr) GetCustomFilterContext(ctx context.Context, id int) (*CustomFilter, error) {
	var output CustomFilter

	req := starr.Request{URI: path.Join(bpCustomFilter, starr.Str(id))}
	if err := p.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// AddCustomFilter creates a custom filter.
func (p *Prowlarr) AddCustomFilter(in *CustomFilter) (*CustomFilter, error) {
	return p.AddCustomFilterContext(context.Background(), in)
}

// AddCustomFilterContext creates a custom filter.
func (p *Prowlarr) AddCustomFilterContext(ctx context.Context, in *CustomFilter) (*CustomFilter, error) {
	var output CustomFilter

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(in); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpCustomFilter, err)
	}

	req := starr.Request{URI: bpCustomFilter, Body: &body}
	if err := p.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateCustomFilter updates a custom filter.
func (p *Prowlarr) UpdateCustomFilter(in *CustomFilter) (*CustomFilter, error) {
	return p.UpdateCustomFilterContext(context.Background(), in)
}

// UpdateCustomFilterContext updates a custom filter.
func (p *Prowlarr) UpdateCustomFilterContext(ctx context.Context, input *CustomFilter) (*CustomFilter, error) {
	var output CustomFilter

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(input); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpCustomFilter, err)
	}

	req := starr.Request{URI: path.Join(bpCustomFilter, starr.Str(input.ID)), Body: &body}
	if err := p.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteCustomFilter deletes a custom filter.
func (p *Prowlarr) DeleteCustomFilter(id int) error {
	return p.DeleteCustomFilterContext(context.Background(), id)
}

// DeleteCustomFilterContext deletes a custom filter.
func (p *Prowlarr) DeleteCustomFilterContext(ctx context.Context, id int) error {
	req := starr.Request{URI: path.Join(bpCustomFilter, starr.Str(id))}
	if err := p.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
