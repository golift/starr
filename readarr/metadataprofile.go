package readarr

import (
	"context"
	"fmt"

	"golift.io/starr"
)

const bpMetadataProfile = APIver + "/metadataprofile"

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

// GetMetadataProfilesContext returns the metadata profiles.
func (r *Readarr) GetMetadataProfilesContext(ctx context.Context) ([]*MetadataProfile, error) {
	var output []*MetadataProfile

	req := starr.Request{URI: bpMetadataProfile}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}
