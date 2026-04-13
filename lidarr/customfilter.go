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

const bpCustomFilter = APIver + "/customfilter"

// CustomFilter is the /api/v1/customfilter resource.
type CustomFilter = starrshared.CustomFilter

// GetCustomFilters returns all custom filters.
func (l *Lidarr) GetCustomFilters() ([]*CustomFilter, error) {
	return l.GetCustomFiltersContext(context.Background())
}

// GetCustomFiltersContext returns all custom filters.
func (l *Lidarr) GetCustomFiltersContext(ctx context.Context) ([]*CustomFilter, error) {
	var output []*CustomFilter

	req := starr.Request{URI: bpCustomFilter}
	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetCustomFilter returns a single custom filter.
func (l *Lidarr) GetCustomFilter(id int) (*CustomFilter, error) {
	return l.GetCustomFilterContext(context.Background(), id)
}

// GetCustomFilterContext returns a single custom filter.
func (l *Lidarr) GetCustomFilterContext(ctx context.Context, id int) (*CustomFilter, error) {
	var output CustomFilter

	req := starr.Request{URI: path.Join(bpCustomFilter, starr.Str(id))}
	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// AddCustomFilter creates a custom filter.
func (l *Lidarr) AddCustomFilter(in *CustomFilter) (*CustomFilter, error) {
	return l.AddCustomFilterContext(context.Background(), in)
}

// AddCustomFilterContext creates a custom filter.
func (l *Lidarr) AddCustomFilterContext(ctx context.Context, in *CustomFilter) (*CustomFilter, error) {
	var output CustomFilter

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(in); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpCustomFilter, err)
	}

	req := starr.Request{URI: bpCustomFilter, Body: &body}
	if err := l.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateCustomFilter updates a custom filter.
func (l *Lidarr) UpdateCustomFilter(in *CustomFilter) (*CustomFilter, error) {
	return l.UpdateCustomFilterContext(context.Background(), in)
}

// UpdateCustomFilterContext updates a custom filter.
func (l *Lidarr) UpdateCustomFilterContext(ctx context.Context, input *CustomFilter) (*CustomFilter, error) {
	var output CustomFilter

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(input); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpCustomFilter, err)
	}

	req := starr.Request{URI: path.Join(bpCustomFilter, starr.Str(input.ID)), Body: &body}
	if err := l.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteCustomFilter deletes a custom filter.
func (l *Lidarr) DeleteCustomFilter(id int) error {
	return l.DeleteCustomFilterContext(context.Background(), id)
}

// DeleteCustomFilterContext deletes a custom filter.
func (l *Lidarr) DeleteCustomFilterContext(ctx context.Context, id int) error {
	req := starr.Request{URI: path.Join(bpCustomFilter, starr.Str(id))}
	if err := l.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
