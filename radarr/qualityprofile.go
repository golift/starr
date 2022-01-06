package radarr

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"golift.io/starr"
)

// GetQualityProfiles returns all configured quality profiles.
func (r *Radarr) GetQualityProfiles() ([]*QualityProfile, error) {
	return r.GetQualityProfilesContext(context.Background())
}

// GetQualityProfilesContext returns all configured quality profiles.
func (r *Radarr) GetQualityProfilesContext(ctx context.Context) ([]*QualityProfile, error) {
	var profiles []*QualityProfile

	err := r.GetInto(ctx, "v3/qualityProfile", nil, &profiles)
	if err != nil {
		return nil, fmt.Errorf("api.Get(qualityProfile): %w", err)
	}

	return profiles, nil
}

// AddQualityProfile updates a quality profile in place.
func (r *Radarr) AddQualityProfile(profile *QualityProfile) (int64, error) {
	return r.AddQualityProfileContext(context.Background(), profile)
}

// AddQualityProfileContext updates a quality profile in place.
func (r *Radarr) AddQualityProfileContext(ctx context.Context, profile *QualityProfile) (int64, error) {
	post, err := json.Marshal(profile)
	if err != nil {
		return 0, fmt.Errorf("json.Marshal(profile): %w", err)
	}

	var output QualityProfile

	err = r.PostInto(ctx, "v3/qualityProfile", nil, post, &output)
	if err != nil {
		return 0, fmt.Errorf("api.Post(qualityProfile): %w", err)
	}

	return output.ID, nil
}

// UpdateQualityProfile updates a quality profile in place.
func (r *Radarr) UpdateQualityProfile(profile *QualityProfile) error {
	return r.UpdateQualityProfileContext(context.Background(), profile)
}

// UpdateQualityProfileContext updates a quality profile in place.
func (r *Radarr) UpdateQualityProfileContext(ctx context.Context, profile *QualityProfile) error {
	put, err := json.Marshal(profile)
	if err != nil {
		return fmt.Errorf("json.Marshal(profile): %w", err)
	}

	_, err = r.Put(ctx, "v3/qualityProfile/"+strconv.FormatInt(profile.ID, starr.Base10), nil, put)
	if err != nil {
		return fmt.Errorf("api.Put(qualityProfile): %w", err)
	}

	return nil
}
