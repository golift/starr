package radarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"

	"golift.io/starr"
)

const bpRestriction = APIver + "/restriction"

// Restriction is the input for a new or updated restriction.
type Restriction struct {
	Tags     []int  `json:"tags,omitempty"`
	Required string `json:"required,omitempty"`
	Ignored  string `json:"ignored,omitempty"`
	ID       int64  `json:"id,omitempty"`
}

// GetRestrictions returns all configured restrictions.
func (r *Radarr) GetRestrictions() ([]*Restriction, error) {
	return r.GetRestrictionsContext(context.Background())
}

// GetRestrictionsContext returns all configured restrictions.
func (r *Radarr) GetRestrictionsContext(ctx context.Context) ([]*Restriction, error) {
	var output []*Restriction

	req := starr.Request{URI: bpRestriction}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetRestriction returns a single restriction.
func (r *Radarr) GetRestriction(restrictionID int64) (*Restriction, error) {
	return r.GetRestrictionContext(context.Background(), restrictionID)
}

// GetIndGetRestrictionContextexer returns a single restriction.
func (r *Radarr) GetRestrictionContext(ctx context.Context, restrictionID int64) (*Restriction, error) {
	var output Restriction

	req := starr.Request{URI: path.Join(bpRestriction, starr.Str(restrictionID))}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// AddRestriction creates a restriction.
func (r *Radarr) AddRestriction(restriction *Restriction) (*Restriction, error) {
	return r.AddRestrictionContext(context.Background(), restriction)
}

// AddRestrictionContext creates a restriction.
func (r *Radarr) AddRestrictionContext(ctx context.Context, restriction *Restriction) (*Restriction, error) {
	var output Restriction

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(restriction); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpRestriction, err)
	}

	req := starr.Request{URI: bpRestriction, Body: &body}
	if err := r.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateRestriction updates the restriction.
func (r *Radarr) UpdateRestriction(restriction *Restriction) (*Restriction, error) {
	return r.UpdateRestrictionContext(context.Background(), restriction)
}

// UpdateRestrictionContext updates the restriction.
func (r *Radarr) UpdateRestrictionContext(ctx context.Context, restriction *Restriction) (*Restriction, error) {
	var output Restriction

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(restriction); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpRestriction, err)
	}

	req := starr.Request{URI: path.Join(bpRestriction, starr.Str(restriction.ID)), Body: &body}
	if err := r.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteRestriction removes a single restriction.
func (r *Radarr) DeleteRestriction(restrictionID int64) error {
	return r.DeleteRestrictionContext(context.Background(), restrictionID)
}

// DeleteRestrictionContext removes a single restriction.
func (r *Radarr) DeleteRestrictionContext(ctx context.Context, restrictionID int64) error {
	req := starr.Request{URI: path.Join(bpRestriction, starr.Str(restrictionID))}
	if err := r.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
