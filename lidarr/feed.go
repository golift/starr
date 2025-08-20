package lidarr

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"strconv"
	"strings"

	"golift.io/starr"
)

// This is not an /api path.
const bpFeed = "feed/" + APIver + "/calendar/lidarr.ics"

// Feed is the /feed/v1/calendar endpoint.
type Feed struct {
	// Default Value: 7
	PastDays int `json:"pastDays"`
	// Default Value: 28
	FutureDays int   `json:"futureDays"`
	Tags       []int `json:"tags"`
	// Include unmonitored albums in the iCal feed.
	Unmonitored bool `json:"unmonitored"`
}

// GetFeed returns the Calendar ICS feed file.
func (r *Lidarr) GetFeed(filter Feed) ([]byte, error) {
	return r.GetFeedContext(context.Background(), filter)
}

// GetFeedContext returns the Calendar ICS feed file.
func (r *Lidarr) GetFeedContext(ctx context.Context, filter Feed) ([]byte, error) {
	tags := make([]string, len(filter.Tags))
	for idx, tag := range filter.Tags {
		tags[idx] = strconv.Itoa(tag)
	}

	req := starr.Request{URI: bpFeed, Query: url.Values{
		"unmonitored": {strconv.FormatBool(filter.Unmonitored)},
		"pastDays":    {strconv.Itoa(filter.PastDays)},
		"futureDays":  {strconv.Itoa(filter.FutureDays)},
		"tags":        {strings.Join(tags, ",")},
	}}

	resp, err := r.Get(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("http.Get(%s): %w", &req, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll: %w", err)
	}

	return body, nil
}
