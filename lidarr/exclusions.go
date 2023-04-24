package lidarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"

	"golift.io/starr"
)

const bpExclusions = APIver + "/importlistexclusion"

// Exclusion is a Lidarr excluded item.
type Exclusion struct {
	ForeignID  string `json:"foreignId"`
	ArtistName string `json:"artistName"`
	ID         int64  `json:"id,omitempty"`
}

// GetExclusions returns all configured exclusions from Lidarr.
func (l *Lidarr) GetExclusions() ([]*Exclusion, error) {
	return l.GetExclusionsContext(context.Background())
}

// GetExclusionsContext returns all configured exclusions from Lidarr.
func (l *Lidarr) GetExclusionsContext(ctx context.Context) ([]*Exclusion, error) {
	var output []*Exclusion

	req := starr.Request{URI: bpExclusions}
	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// UpdateExclusion changes an exclusions in Lidarr.
func (l *Lidarr) UpdateExclusion(exclusion *Exclusion) (*Exclusion, error) {
	return l.UpdateExclusionContext(context.Background(), exclusion)
}

// UpdateExclusionContext changes an exclusions in Lidarr.
func (l *Lidarr) UpdateExclusionContext(ctx context.Context, exclusion *Exclusion) (*Exclusion, error) {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(exclusion); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpExclusions, err)
	}

	var output Exclusion

	req := starr.Request{URI: path.Join(bpExclusions, fmt.Sprint(exclusion.ID)), Body: &body}
	if err := l.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteExclusions removes exclusions from Lidarr.
func (l *Lidarr) DeleteExclusions(ids []int64) error {
	return l.DeleteExclusionsContext(context.Background(), ids)
}

// DeleteExclusionsContext removes exclusions from Lidarr.
func (l *Lidarr) DeleteExclusionsContext(ctx context.Context, ids []int64) error {
	var errs string

	for _, id := range ids {
		req := starr.Request{URI: path.Join(bpExclusions, fmt.Sprint(id))}
		if err := l.DeleteAny(ctx, req); err != nil {
			errs += fmt.Sprintf("api.Post(%s): %v ", &req, err)
		}
	}

	if errs != "" {
		return fmt.Errorf("%w: %s", starr.ErrRequestError, errs)
	}

	return nil
}
