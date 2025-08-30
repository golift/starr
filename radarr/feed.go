package radarr

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
const bpFeed = "feed/" + APIver + "/calendar/radarr.ics"

type ReleaseType string

const (
	ReleaseTypeCinema   ReleaseType = "cinemaRelease"
	ReleaseTypeDigital  ReleaseType = "digitalRelease"
	ReleaseTypePhysical ReleaseType = "physicalRelease"
)

// Feed is the /feed/v3/calendar endpoint.
type Feed struct {
	// Default Value: 7
	PastDays int `json:"pastDays"`
	// Default Value: 28
	FutureDays int   `json:"futureDays"`
	Tags       []int `json:"tags"`
	// Include unmonitored movies in the iCal feed.
	Unmonitored  bool          `json:"unmonitored"`
	ReleaseTypes []ReleaseType `json:"releaseTypes"`
	// Events will appear as all day events in your calendar.
	AsAllDay bool `json:"asAllDay"`
}

// GetFeed returns the Calendar ICS feed file.
func (r *Radarr) GetFeed(filter Feed) ([]byte, error) {
	return r.GetFeedContext(context.Background(), filter)
}

// GetFeedContext returns the Calendar ICS feed file.
func (r *Radarr) GetFeedContext(ctx context.Context, filter Feed) ([]byte, error) {
	tags := make([]string, len(filter.Tags))
	for idx, tag := range filter.Tags {
		tags[idx] = strconv.Itoa(tag)
	}

	releaseTypes := make([]string, len(filter.ReleaseTypes))
	for idx, releaseType := range filter.ReleaseTypes {
		releaseTypes[idx] = string(releaseType)
	}

	req := starr.Request{URI: bpFeed, Query: url.Values{
		"unmonitored":  {strconv.FormatBool(filter.Unmonitored)},
		"pastDays":     {strconv.Itoa(filter.PastDays)},
		"futureDays":   {strconv.Itoa(filter.FutureDays)},
		"tags":         {strings.Join(tags, ",")},
		"asAllDay":     {strconv.FormatBool(filter.AsAllDay)},
		"releaseTypes": {strings.Join(releaseTypes, ",")},
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
