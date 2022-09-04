package readarr

import (
	"context"
	"fmt"
)

// MetadataProfile is the /api/v1/metadataProfile endpoint.
type MetadataProfile struct {
	ID                  int64   `json:"id"`
	Name                string  `json:"name"`
	MinPopularity       float64 `json:"minPopularity"`
	SkipMissingDate     bool    `json:"skipMissingDate"`
	SkipMissingIsbn     bool    `json:"skipMissingIsbn"`
	SkipPartsAndSets    bool    `json:"skipPartsAndSets"`
	SkipSeriesSecondary bool    `json:"skipSeriesSecondary"`
	AllowedLanguages    string  `json:"allowedLanguages,omitempty"`
}

// GetMetadataProfiles returns the metadata profiles.
func (r *Readarr) GetMetadataProfiles() ([]*MetadataProfile, error) {
	return r.GetMetadataProfilesContext(context.Background())
}

func (r *Readarr) GetMetadataProfilesContext(ctx context.Context) ([]*MetadataProfile, error) {
	var profiles []*MetadataProfile

	err := r.GetInto(ctx, "v1/metadataprofile", nil, &profiles)
	if err != nil {
		return nil, fmt.Errorf("api.Get(metadataprofile): %w", err)
	}

	return profiles, nil
}
