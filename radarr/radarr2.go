package radarr

import (
	"fmt"
	"net/url"
)

/* This is all deprecated and will be removed in the future. Switch to v3. */

// GetQueueV2 returns the Radarr Queue (processing, but not yet imported).
func (r *Radarr) GetQueueV2() ([]*Queue, error) {
	params := make(url.Values)
	params.Set("sort_by", "timeleft")
	params.Set("order", "asc")

	var queue []*Queue
	if err := r.GetInto("queue", params, &queue); err != nil {
		return nil, fmt.Errorf("api.Get(queue): %w", err)
	}

	return queue, nil
}
