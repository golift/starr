package sonarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

/* seasonPass only seems to handle POSTs. */

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
		return fmt.Errorf("json.Marshal(seasonPass): %w", err)
	}

	if _, err := s.Post(ctx, "v3/seasonPass", nil, &body); err != nil {
		return fmt.Errorf("api.Post(seasonPass): %w", err)
	}

	return nil
}
