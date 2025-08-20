package readarr

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
const bpFeed = "feed/" + APIver + "/calendar/readarr.ics"

// Feed is the /feed/v3/calendar endpoint.
type Feed struct {
	// Default Value: 7
	PastDays int `json:"pastDays"`
	// Default Value: 28
	FutureDays int   `json:"futureDays"`
	Tags       []int `json:"tags"`
	// Include unmonitored books in the iCal feed.
	Unmonitored bool `json:"unmonitored"`
}

func (r *Readarr) GetFeed(filter Feed) ([]byte, error) {
	return r.GetFeedContext(context.Background(), filter)
}

// GetFeedContext returns the Readarr Feed.
func (r *Readarr) GetFeedContext(ctx context.Context, filter Feed) ([]byte, error) {
	var (
		tags []string
	)

	for _, tag := range filter.Tags {
		tags = append(tags, strconv.Itoa(tag))
	}

	query := url.Values{
		"unmonitored": {strconv.FormatBool(filter.Unmonitored)},
		"pastDays":    {strconv.Itoa(filter.PastDays)},
		"futureDays":  {strconv.Itoa(filter.FutureDays)},
		"tags":        {strings.Join(tags, ",")},
	}

	resp, err := r.Get(ctx, starr.Request{URI: bpFeed, Query: query})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll: %w", err)
	}

	return body, nil
}
