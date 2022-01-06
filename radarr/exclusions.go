package radarr

import (
	"encoding/json"
	"fmt"
	"strconv"

	"golift.io/starr"
)

// GetExclusions returns all configured exclusions from Radarr.
func (r *Radarr) GetExclusions() ([]*Exclusion, error) {
	var exclusions []*Exclusion

	err := r.GetInto("v3/exclusions", nil, &exclusions)
	if err != nil {
		return nil, fmt.Errorf("api.Get(exclusions): %w", err)
	}

	return exclusions, nil
}

// DeleteExclusions removes exclusions from Radarr.
func (r *Radarr) DeleteExclusions(ids []int64) error {
	var errs string

	for _, id := range ids {
		_, err := r.Delete("v3/exclusions/"+strconv.FormatInt(id, starr.Base10), nil)
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
	for i := range exclusions {
		exclusions[i].ID = 0
	}

	body, err := json.Marshal(exclusions)
	if err != nil {
		return fmt.Errorf("json.Marshal(movie): %w", err)
	}

	_, err = r.Post("v3/exclusions/bulk", nil, body)
	if err != nil {
		return fmt.Errorf("api.Post(exclusions): %w", err)
	}

	return nil
}
