package readarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"

	"golift.io/starr"
)

const bpExclusions = APIver + "/importlistexclusion"

// Exclusion is a Readarr excluded item.
type Exclusion struct {
	ForeignID  string `json:"foreignId"`
	AuthorName string `json:"authorName"`
	ID         int64  `json:"id,omitempty"`
}

// GetExclusions returns all configured exclusions from Readarr.
func (r *Readarr) GetExclusions() ([]*Exclusion, error) {
	return r.GetExclusionsContext(context.Background())
}

// GetExclusionsContext returns all configured exclusions from Readarr.
func (r *Readarr) GetExclusionsContext(ctx context.Context) ([]*Exclusion, error) {
	var output []*Exclusion

	req := starr.Request{URI: bpExclusions}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// UpdateExclusion changes an exclusions in Readarr.
func (r *Readarr) UpdateExclusion(exclusion *Exclusion) (*Exclusion, error) {
	return r.UpdateExclusionContext(context.Background(), exclusion)
}

// UpdateExclusionContext changes an exclusions in Readarr.
func (r *Readarr) UpdateExclusionContext(ctx context.Context, exclusion *Exclusion) (*Exclusion, error) {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(exclusion); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpExclusions, err)
	}

	var output Exclusion

	req := starr.Request{URI: path.Join(bpExclusions, starr.Itoa(exclusion.ID)), Body: &body}
	if err := r.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteExclusions removes exclusions from Readarr.
func (r *Readarr) DeleteExclusions(ids []int64) error {
	return r.DeleteExclusionsContext(context.Background(), ids)
}

// DeleteExclusionsContext removes exclusions from Readarr.
func (r *Readarr) DeleteExclusionsContext(ctx context.Context, ids []int64) error {
	var errs string

	for _, id := range ids {
		req := starr.Request{URI: path.Join(bpExclusions, starr.Itoa(id))}
		if err := r.DeleteAny(ctx, req); err != nil {
			errs += fmt.Sprintf("api.Post(%s): %v ", &req, err)
		}
	}

	if errs != "" {
		return fmt.Errorf("%w: %s", starr.ErrRequestError, errs)
	}

	return nil
}

// AddExclusion adds one exclusion to Readarr.
func (r *Readarr) AddExclusion(exclusion *Exclusion) (*Exclusion, error) {
	return r.AddExclusionContext(context.Background(), exclusion)
}

// AddExclusionContext adds one exclusion to Readarr.
func (r *Readarr) AddExclusionContext(ctx context.Context, exclusion *Exclusion) (*Exclusion, error) {
	exclusion.ID = 0 // if this panics, fix your code

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(exclusion); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpExclusions, err)
	}

	var output Exclusion

	req := starr.Request{URI: path.Join(bpExclusions), Body: &body}
	if err := r.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}
