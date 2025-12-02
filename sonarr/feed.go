package sonarr

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
const bpFeed = "feed/" + APIver + "/calendar/sonarr.ics"

// Feed is the /feed/v3/calendar endpoint.
type Feed struct {
	// Default Value: 7
	PastDays int `json:"pastDays"`
	// Default Value: 28
	FutureDays int   `json:"futureDays"`
	Tags       []int `json:"tags"`
	// Include unmonitored tv shows in the iCal feed.
	Unmonitored bool `json:"unmonitored"`
	// Only include premieres in the iCal feed.
	PremieresOnly bool `json:"premieresOnly"`
	// Events will appear as all day events in your calendar.
	AsAllDay bool `json:"asAllDay"`
}

// GetFeed returns the Calendar ICS feed file.
func (r *Sonarr) GetFeed(filter Feed) ([]byte, error) {
	return r.GetFeedContext(context.Background(), filter)
}

// GetFeedContext returns the Calendar ICS feed file.
func (r *Sonarr) GetFeedContext(ctx context.Context, filter Feed) ([]byte, error) {
	tags := make([]string, len(filter.Tags))
	for idx, tag := range filter.Tags {
		tags[idx] = strconv.Itoa(tag)
	}

	req := starr.Request{URI: bpFeed, Query: url.Values{
		"unmonitored":   {strconv.FormatBool(filter.Unmonitored)},
		"pastDays":      {strconv.Itoa(filter.PastDays)},
		"futureDays":    {strconv.Itoa(filter.FutureDays)},
		"tags":          {strings.Join(tags, ",")},
		"asAllDay":      {strconv.FormatBool(filter.AsAllDay)},
		"premieresOnly": {strconv.FormatBool(filter.PremieresOnly)},
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
