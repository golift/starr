package sonarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"

	"golift.io/starr"
)

const bpCustomFilter = APIver + "/customfilter"

// CustomFilter is the /api/v3/customfilter resource.
type CustomFilter struct {
	ID      int               `json:"id,omitempty"`
	Type    string            `json:"type,omitempty"`
	Label   string            `json:"label,omitempty"`
	Filters []json.RawMessage `json:"filters,omitempty"`
}

// GetCustomFilters returns all custom filters.
func (s *Sonarr) GetCustomFilters() ([]*CustomFilter, error) {
	return s.GetCustomFiltersContext(context.Background())
}

// GetCustomFiltersContext returns all custom filters.
func (s *Sonarr) GetCustomFiltersContext(ctx context.Context) ([]*CustomFilter, error) {
	var output []*CustomFilter

	req := starr.Request{URI: bpCustomFilter}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetCustomFilter returns a single custom filter.
func (s *Sonarr) GetCustomFilter(id int) (*CustomFilter, error) {
	return s.GetCustomFilterContext(context.Background(), id)
}

// GetCustomFilterContext returns a single custom filter.
func (s *Sonarr) GetCustomFilterContext(ctx context.Context, id int) (*CustomFilter, error) {
	var output CustomFilter

	req := starr.Request{URI: path.Join(bpCustomFilter, starr.Str(id))}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// AddCustomFilter creates a custom filter.
func (s *Sonarr) AddCustomFilter(in *CustomFilter) (*CustomFilter, error) {
	return s.AddCustomFilterContext(context.Background(), in)
}

// AddCustomFilterContext creates a custom filter.
func (s *Sonarr) AddCustomFilterContext(ctx context.Context, in *CustomFilter) (*CustomFilter, error) {
	var output CustomFilter

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(in); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpCustomFilter, err)
	}

	req := starr.Request{URI: bpCustomFilter, Body: &body}
	if err := s.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateCustomFilter updates a custom filter.
func (s *Sonarr) UpdateCustomFilter(in *CustomFilter) (*CustomFilter, error) {
	return s.UpdateCustomFilterContext(context.Background(), in)
}

// UpdateCustomFilterContext updates a custom filter.
func (s *Sonarr) UpdateCustomFilterContext(ctx context.Context, input *CustomFilter) (*CustomFilter, error) {
	var output CustomFilter

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(input); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpCustomFilter, err)
	}

	req := starr.Request{URI: path.Join(bpCustomFilter, starr.Str(input.ID)), Body: &body}
	if err := s.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteCustomFilter deletes a custom filter.
func (s *Sonarr) DeleteCustomFilter(id int) error {
	return s.DeleteCustomFilterContext(context.Background(), id)
}

// DeleteCustomFilterContext deletes a custom filter.
func (s *Sonarr) DeleteCustomFilterContext(ctx context.Context, id int) error {
	req := starr.Request{URI: path.Join(bpCustomFilter, starr.Str(id))}
	if err := s.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
