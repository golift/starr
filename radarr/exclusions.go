package radarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"golift.io/starr"
)

// Exclusion is a Radarr excluded item.
type Exclusion struct {
	TMDBID int64  `json:"tmdbId"`
	Title  string `json:"movieTitle"`
	Year   int    `json:"movieYear"`
	ID     int64  `json:"id,omitempty"`
}

// GetExclusions returns all configured exclusions from Radarr.
func (r *Radarr) GetExclusions() ([]*Exclusion, error) {
	return r.GetExclusionsContext(context.Background())
}

// GetExclusionsContext returns all configured exclusions from Radarr.
func (r *Radarr) GetExclusionsContext(ctx context.Context) ([]*Exclusion, error) {
	var exclusions []*Exclusion

	_, err := r.GetInto(ctx, "v3/exclusions", nil, &exclusions)
	if err != nil {
		return nil, fmt.Errorf("api.Get(exclusions): %w", err)
	}

	return exclusions, nil
}

// DeleteExclusions removes exclusions from Radarr.
func (r *Radarr) DeleteExclusions(ids []int64) error {
	return r.DeleteExclusionsContext(context.Background(), ids)
}

// DeleteExclusionsContext removes exclusions from Radarr.
func (r *Radarr) DeleteExclusionsContext(ctx context.Context, ids []int64) error {
	var errs string

	for _, id := range ids {
		_, err := r.Delete(ctx, "v3/exclusions/"+strconv.FormatInt(id, starr.Base10), nil)
		if err != nil {
			errs += err.Error() + " "
		}
	}

	if errs != "" {
		return fmt.Errorf("%w: %s", starr.ErrRequestError, errs)
	}

	return nil
}

// AddExclusions adds an exclusion to Radarr.
func (r *Radarr) AddExclusions(exclusions []*Exclusion) error {
	return r.AddExclusionsContext(context.Background(), exclusions)
}

// AddExclusionsContext adds an exclusion to Radarr.
func (r *Radarr) AddExclusionsContext(ctx context.Context, exclusions []*Exclusion) error {
	for i := range exclusions {
		exclusions[i].ID = 0
	}

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(exclusions); err != nil {
		return fmt.Errorf("json.Marshal(exclusions): %w", err)
	}

	_, err := r.Post(ctx, "v3/exclusions/bulk", nil, &body)
	if err != nil {
		return fmt.Errorf("api.Post(exclusions): %w", err)
	}

	return nil
}
