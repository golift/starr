package readarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"
	"strconv"

	"golift.io/starr"
)

const bpQualityProfile = APIver + "/qualityProfile"

// QualityProfile is the /api/v1/qualityprofile endpoint.
type QualityProfile struct {
	Name           string           `json:"name"`
	UpgradeAllowed bool             `json:"upgradeAllowed"`
	Cutoff         int64            `json:"cutoff"`
	Qualities      []*starr.Quality `json:"items"`
	ID             int64            `json:"id"`
}

// GetQualityProfiles returns the quality profiles.
func (r *Readarr) GetQualityProfiles() ([]*QualityProfile, error) {
	return r.GetQualityProfilesContext(context.Background())
}

func (r *Readarr) GetQualityProfilesContext(ctx context.Context) ([]*QualityProfile, error) {
	var profiles []*QualityProfile

	err := r.GetInto(ctx, bpQualityProfile, nil, &profiles)
	if err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", bpQualityProfile, err)
	}

	return profiles, nil
}

// AddQualityProfile updates a quality profile in place.
func (r *Readarr) AddQualityProfile(profile *QualityProfile) (int64, error) {
	return r.AddQualityProfileContext(context.Background(), profile)
}

func (r *Readarr) AddQualityProfileContext(ctx context.Context, profile *QualityProfile) (int64, error) {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(profile); err != nil {
		return 0, fmt.Errorf("json.Marshal(%s): %w", bpQualityProfile, err)
	}

	var output QualityProfile
	if err := r.PostInto(ctx, bpQualityProfile, nil, &body, &output); err != nil {
		return 0, fmt.Errorf("api.Post(%s): %w", bpQualityProfile, err)
	}

	return output.ID, nil
}

// UpdateQualityProfile updates a quality profile in place.
func (r *Readarr) UpdateQualityProfile(profile *QualityProfile) error {
	return r.UpdateQualityProfileContext(context.Background(), profile)
}

func (r *Readarr) UpdateQualityProfileContext(ctx context.Context, profile *QualityProfile) error {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(profile); err != nil {
		return fmt.Errorf("json.Marshal(%s): %w", bpQualityProfile, err)
	}

	var output interface{}

	uri := path.Join(bpQualityProfile, strconv.FormatInt(profile.ID, starr.Base10))
	if err := r.PutInto(ctx, uri, nil, &body, &output); err != nil {
		return fmt.Errorf("api.Put(%s): %w", bpQualityProfile, err)
	}

	return nil
}

// DeleteQualityProfile deletes a quality profile.
func (r *Readarr) DeleteQualityProfile(profileID int64) error {
	return r.DeleteQualityProfileContext(context.Background(), profileID)
}

// DeleteQualityProfileContext deletes a quality profile.
func (r *Readarr) DeleteQualityProfileContext(ctx context.Context, profileID int64) error {
	req := starr.Request{URI: path.Join(bpQualityProfile, fmt.Sprint(profileID))}
	if err := r.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", req.URI, err)
	}

	return nil
}
