package sonarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"

	"golift.io/starr"
)

const bpExclusions = APIver + "/importlistexclusion"

// Exclusion is a Sonarr excluded item.
type Exclusion struct {
	TVDBID int64  `json:"tvdbId"`
	Title  string `json:"title"`
	ID     int64  `json:"id,omitempty"`
}

// GetExclusions returns all configured exclusions from Sonarr.
func (s *Sonarr) GetExclusions() ([]*Exclusion, error) {
	return s.GetExclusionsContext(context.Background())
}

// GetExclusionsContext returns all configured exclusions from Sonarr.
func (s *Sonarr) GetExclusionsContext(ctx context.Context) ([]*Exclusion, error) {
	var output []*Exclusion

	req := starr.Request{URI: bpExclusions}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// UpdateExclusion changes an exclusions in Sonarr.
func (s *Sonarr) UpdateExclusion(exclusion *Exclusion) (*Exclusion, error) {
	return s.UpdateExclusionContext(context.Background(), exclusion)
}

// UpdateExclusionContext changes an exclusions in Sonarr.
func (s *Sonarr) UpdateExclusionContext(ctx context.Context, exclusion *Exclusion) (*Exclusion, error) {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(exclusion); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpExclusions, err)
	}

	var output Exclusion

	req := starr.Request{URI: path.Join(bpExclusions, starr.Str(exclusion.ID)), Body: &body}
	if err := s.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteExclusions removes exclusions from Sonarr.
func (s *Sonarr) DeleteExclusions(ids []int64) error {
	return s.DeleteExclusionsContext(context.Background(), ids)
}

// DeleteExclusionsContext removes exclusions from Sonarr.
func (s *Sonarr) DeleteExclusionsContext(ctx context.Context, ids []int64) error {
	var errs string

	for _, id := range ids {
		req := starr.Request{URI: path.Join(bpExclusions, starr.Str(id))}
		if err := s.DeleteAny(ctx, req); err != nil {
			errs += fmt.Sprintf("api.Post(%s): %v ", &req, err)
		}
	}

	if errs != "" {
		return fmt.Errorf("%w: %s", starr.ErrRequestError, errs)
	}

	return nil
}

// AddExclusion adds one exclusion to Sonarr.
func (s *Sonarr) AddExclusion(exclusion *Exclusion) (*Exclusion, error) {
	return s.AddExclusionContext(context.Background(), exclusion)
}

// AddExclusionContext adds one exclusion to Sonarr.
func (s *Sonarr) AddExclusionContext(ctx context.Context, exclusion *Exclusion) (*Exclusion, error) {
	exclusion.ID = 0 // if this panics, fix your code

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(exclusion); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpExclusions, err)
	}

	var output Exclusion

	req := starr.Request{URI: path.Join(bpExclusions), Body: &body}
	if err := s.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}
