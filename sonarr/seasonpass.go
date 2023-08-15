package sonarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"golift.io/starr"
)

/* seasonPass only seems to handle POSTs. */

const bpSeasonPass = APIver + "/seasonPass"

// SeasonPass is the input payload for a seasonPass update.
type SeasonPass struct {
	Series            []*MonitoredSeries `json:"series"`
	MonitoringOptions *MonitoringOptions `json:"monitoringOptions"`
}

// MonitoringOptions is part of the SeasonPass payload.
type MonitoringOptions struct {
	// Valid values for Monitor are: all, future, missing, existing, firstSeason, latestSeason, and none.
	Monitor string `json:"monitor"`
}

// MonitoredSeries is part of the SeasonPass payload.
type MonitoredSeries struct {
	ID        int  `json:"id"`
	Monitored bool `json:"monitored"`
}

// UpdateSeasonPass allows monitoring many Series and episodes at once.
func (s *Sonarr) UpdateSeasonPass(seasonPass *SeasonPass) error {
	return s.UpdateSeasonPassContext(context.Background(), seasonPass)
}

// UpdateSeasonPassContext allows monitoring many series and episodes at once.
func (s *Sonarr) UpdateSeasonPassContext(ctx context.Context, seasonPass *SeasonPass) error {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(seasonPass); err != nil {
		return fmt.Errorf("json.Marshal(%s): %w", bpSeasonPass, err)
	}

	var output interface{} // any ok

	req := starr.Request{URI: bpSeasonPass, Body: &body}
	if err := s.PostInto(ctx, req, &output); err != nil {
		return fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return nil
}
