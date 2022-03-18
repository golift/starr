package lidarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"golift.io/starr"
)

// QualityProfile is the /api/v1/qualityprofile endpoint.
type QualityProfile struct {
	ID             int64            `json:"id"`
	Name           string           `json:"name"`
	UpgradeAllowed bool             `json:"upgradeAllowed"`
	Cutoff         int64            `json:"cutoff"`
	Qualities      []*starr.Quality `json:"items"`
}

// GetQualityProfiles returns the quality profiles.
func (l *Lidarr) GetQualityProfiles() ([]*QualityProfile, error) {
	return l.GetQualityProfilesContext(context.Background())
}

// GetQualityProfilesContext returns the quality profiles.
func (l *Lidarr) GetQualityProfilesContext(ctx context.Context) ([]*QualityProfile, error) {
	var profiles []*QualityProfile

	_, err := l.GetInto(ctx, "v1/qualityprofile", nil, &profiles)
	if err != nil {
		return nil, fmt.Errorf("api.Get(qualityprofile): %w", err)
	}

	return profiles, nil
}

// AddQualityProfile updates a quality profile in place.
func (l *Lidarr) AddQualityProfile(profile *QualityProfile) (int64, error) {
	return l.AddQualityProfileContext(context.Background(), profile)
}

// AddQualityProfileContext updates a quality profile in place.
func (l *Lidarr) AddQualityProfileContext(ctx context.Context, profile *QualityProfile) (int64, error) {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(profile); err != nil {
		return 0, fmt.Errorf("json.Marshal(qualityProfile): %w", err)
	}

	var output QualityProfile
	if _, err := l.PostInto(ctx, "v1/qualityProfile", nil, &body, &output); err != nil {
		return 0, fmt.Errorf("api.Post(qualityProfile): %w", err)
	}

	return output.ID, nil
}

// UpdateQualityProfile updates a quality profile in place.
func (l *Lidarr) UpdateQualityProfile(profile *QualityProfile) error {
	return l.UpdateQualityProfileContext(context.Background(), profile)
}

// UpdateQualityProfileContext updates a quality profile in place.
func (l *Lidarr) UpdateQualityProfileContext(ctx context.Context, profile *QualityProfile) error {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(profile); err != nil {
		return fmt.Errorf("json.Marshal(qualityProfile): %w", err)
	}

	_, err := l.Put(ctx, "v1/qualityProfile/"+strconv.FormatInt(profile.ID, starr.Base10), nil, &body)
	if err != nil {
		return fmt.Errorf("api.Put(qualityProfile): %w", err)
	}

	return nil
}
