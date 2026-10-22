package sonarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"

	"golift.io/starr"
)

// Define Base Path for Release Profile calls.
const bpReleaseProfile = APIver + "/releaseProfile"

// ReleaseProfile defines a release profile's data from Sonarr.
// A nil IndexerIDs falls back to IndexerID. An empty list means all indexers.
// Zero is never a valid indexer id.
type ReleaseProfile struct {
	Name                                   string            `json:"name"`
	Enabled                                bool              `json:"enabled"`
	AirDateRestriction                     bool              `json:"airDateRestriction,omitempty"`
	AllowSeasonPackWithoutAllEpisodesAired bool              `json:"allowSeasonPackWithoutAllEpisodesAired,omitempty"`
	Required                               []string          `json:"required"`
	Ignored                                []string          `json:"ignored"`
	IndexerID                              int64             `json:"indexerId"`
	IndexerIDs                             []int64           `json:"indexerIds,omitempty"`
	AirDateGracePeriod                     int               `json:"airDateGracePeriod,omitempty"`
	Tags                                   []int             `json:"tags"`
	ExcludedTags                           []int             `json:"excludedTags,omitempty"`
	ID                                     int64             `json:"id,omitempty"`
	IncPrefOnRename                        *bool             `json:"includePreferredWhenRenaming,omitempty"` // v3 only.
	Preferred                              []*starr.KeyValue `json:"preferred,omitempty"`                    // v3 only.
}

// GetReleaseProfiles returns all configured release profiles.
func (s *Sonarr) GetReleaseProfiles() ([]*ReleaseProfile, error) {
	return s.GetReleaseProfilesContext(context.Background())
}

// GetReleaseProfilesContext returns all configured release profiles.
func (s *Sonarr) GetReleaseProfilesContext(ctx context.Context) ([]*ReleaseProfile, error) {
	var output []*ReleaseProfile

	req := starr.Request{URI: bpReleaseProfile}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetReleaseProfile returns a single release profile.
func (s *Sonarr) GetReleaseProfile(profileID int64) (*ReleaseProfile, error) {
	return s.GetReleaseProfileContext(context.Background(), profileID)
}

// GetReleaseProfileContext returns a single release profile.
func (s *Sonarr) GetReleaseProfileContext(ctx context.Context, profileID int64) (*ReleaseProfile, error) {
	var output ReleaseProfile

	req := starr.Request{URI: path.Join(bpReleaseProfile, starr.Str(profileID))}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// AddReleaseProfile creates a release profile.
func (s *Sonarr) AddReleaseProfile(profile *ReleaseProfile) (*ReleaseProfile, error) {
	return s.AddReleaseProfileContext(context.Background(), profile)
}

// AddReleaseProfileContext creates a release profile.
func (s *Sonarr) AddReleaseProfileContext(ctx context.Context, profile *ReleaseProfile) (*ReleaseProfile, error) {
	var output ReleaseProfile

	body, err := s.releaseProfileBody(profile)
	if err != nil {
		return nil, err
	}

	req := starr.Request{URI: bpReleaseProfile, Body: body}
	if err := s.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateReleaseProfile updates the release profile.
func (s *Sonarr) UpdateReleaseProfile(profile *ReleaseProfile) (*ReleaseProfile, error) {
	return s.UpdateReleaseProfileContext(context.Background(), profile)
}

// UpdateReleaseProfileContext updates the release profile.
func (s *Sonarr) UpdateReleaseProfileContext(ctx context.Context, profile *ReleaseProfile) (*ReleaseProfile, error) {
	var output ReleaseProfile

	body, err := s.releaseProfileBody(profile)
	if err != nil {
		return nil, err
	}

	req := starr.Request{URI: path.Join(bpReleaseProfile, starr.Str(profile.ID)), Body: body}
	if err := s.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteReleaseProfile removes a single release profile.
func (s *Sonarr) DeleteReleaseProfile(profileID int64) error {
	return s.DeleteReleaseProfileContext(context.Background(), profileID)
}

// DeleteReleaseProfileContext removes a single release profile.
func (s *Sonarr) DeleteReleaseProfileContext(ctx context.Context, profileID int64) error {
	req := starr.Request{URI: path.Join(bpReleaseProfile, starr.Str(profileID))}
	if err := s.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}

func (s *Sonarr) releaseProfileBody(profile *ReleaseProfile) (*bytes.Buffer, error) {
	var (
		body    bytes.Buffer
		payload any = profile.v3Payload()
	)

	if s.APIVersion() == APIv5 {
		payload = profile.v5Payload()
	}

	if err := json.NewEncoder(&body).Encode(payload); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpReleaseProfile, err)
	}

	return &body, nil
}

// UnmarshalJSON accepts indexerId and indexerIds. Zero is dropped so "all indexers" stays empty.
func (p *ReleaseProfile) UnmarshalJSON(data []byte) error {
	type profile ReleaseProfile

	var raw profile
	if err := json.Unmarshal(data, &raw); err != nil {
		return fmt.Errorf("json.Unmarshal(releaseProfile): %w", err)
	}

	*p = ReleaseProfile(raw)
	p.IndexerIDs = normalizeIndexerIDs(p.IndexerIDs, p.IndexerID)

	return nil
}

// v3Payload sends indexerId and keeps the v3-only preferred fields.
func (p *ReleaseProfile) v3Payload() *ReleaseProfile {
	if p == nil {
		return nil
	}

	clone := *p
	clone.IndexerIDs = nil

	return &clone
}

// v5Payload sends indexerIds and never indexerId. An empty list means all indexers.
func (p *ReleaseProfile) v5Payload() releaseProfileV5 {
	if p == nil {
		return releaseProfileV5{
			IndexerIDs:   []int64{},
			Required:     []string{},
			Ignored:      []string{},
			Tags:         []int{},
			ExcludedTags: []int{},
		}
	}

	return releaseProfileV5{
		Name:                                   p.Name,
		Enabled:                                p.Enabled,
		Required:                               nonNilStrings(p.Required),
		Ignored:                                nonNilStrings(p.Ignored),
		IndexerIDs:                             p.indexerIDs(),
		Tags:                                   nonNilInts(p.Tags),
		ExcludedTags:                           nonNilInts(p.ExcludedTags),
		ID:                                     p.ID,
		AirDateRestriction:                     p.AirDateRestriction,
		AirDateGracePeriod:                     p.AirDateGracePeriod,
		AllowSeasonPackWithoutAllEpisodesAired: p.AllowSeasonPackWithoutAllEpisodesAired,
	}
}

func (p *ReleaseProfile) indexerIDs() []int64 {
	if p == nil {
		return []int64{}
	}

	ids := normalizeIndexerIDs(p.IndexerIDs, p.IndexerID)
	if len(ids) == 0 {
		return []int64{}
	}

	return ids
}

type releaseProfileV5 struct {
	Name                                   string   `json:"name"`
	Enabled                                bool     `json:"enabled"`
	AirDateRestriction                     bool     `json:"airDateRestriction"`
	AllowSeasonPackWithoutAllEpisodesAired bool     `json:"allowSeasonPackWithoutAllEpisodesAired"`
	Required                               []string `json:"required"`
	Ignored                                []string `json:"ignored"`
	IndexerIDs                             []int64  `json:"indexerIds"`
	AirDateGracePeriod                     int      `json:"airDateGracePeriod"`
	Tags                                   []int    `json:"tags"`
	ExcludedTags                           []int    `json:"excludedTags"`
	ID                                     int64    `json:"id,omitempty"`
}

// normalizeIndexerIDs drops non-positive ids. A nil list falls back to single
// when that id is positive. A non-nil list is kept as given, so an explicit
// empty list stays empty and means all indexers.
func normalizeIndexerIDs(ids []int64, single int64) []int64 {
	if ids == nil {
		if single > 0 {
			return []int64{single}
		}

		return nil
	}

	cleaned := make([]int64, 0, len(ids))
	for _, id := range ids {
		if id > 0 {
			cleaned = append(cleaned, id)
		}
	}

	return cleaned
}

func nonNilInts(vals []int) []int {
	if vals == nil {
		return []int{}
	}

	return vals
}

func nonNilStrings(vals []string) []string {
	if vals == nil {
		return []string{}
	}

	return vals
}
