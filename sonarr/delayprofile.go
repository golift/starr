package sonarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"

	"golift.io/starr"
)

// Define Base Path for Delay Profile calls.
const bpDelayProfile = APIver + "/delayProfile"

// DelayProfile is the /api/v3/delayprofile endpoint.
type DelayProfile struct {
	EnableUsenet           bool   `json:"enableUsenet,omitempty"`
	EnableTorrent          bool   `json:"enableTorrent,omitempty"`
	BypassIfHighestQuality bool   `json:"bypassIfHighestQuality,omitempty"`
	UsenetDelay            int64  `json:"usenetDelay,omitempty"`
	TorrentDelay           int64  `json:"torrentDelay,omitempty"`
	ID                     int64  `json:"id,omitempty"`
	Order                  int64  `json:"order,omitempty"`
	Tags                   []int  `json:"tags"`
	PreferredProtocol      string `json:"preferredProtocol,omitempty"`
}

// GetDelayProfiles returns all configured delay profiles.
func (s *Sonarr) GetDelayProfiles() ([]*DelayProfile, error) {
	return s.GetDelayProfilesContext(context.Background())
}

// GetDelayProfilesContext returns all configured delay profiles.
func (s *Sonarr) GetDelayProfilesContext(ctx context.Context) ([]*DelayProfile, error) {
	var output []*DelayProfile

	req := starr.Request{URI: bpDelayProfile}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetDelayProfile returns a single delay profile.
func (s *Sonarr) GetDelayProfile(profileID int64) (*DelayProfile, error) {
	return s.GetDelayProfileContext(context.Background(), profileID)
}

// GetDelayProfileContext returns a single delay profile.
func (s *Sonarr) GetDelayProfileContext(ctx context.Context, profileID int64) (*DelayProfile, error) {
	var output DelayProfile

	req := starr.Request{URI: path.Join(bpDelayProfile, fmt.Sprint(profileID))}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// AddDelayProfile creates a delay profile.
func (s *Sonarr) AddDelayProfile(profile *DelayProfile) (*DelayProfile, error) {
	return s.AddDelayProfileContext(context.Background(), profile)
}

// AddDelayProfileContext creates a delay profile.
func (s *Sonarr) AddDelayProfileContext(ctx context.Context, profile *DelayProfile) (*DelayProfile, error) {
	var output DelayProfile

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(profile); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpDelayProfile, err)
	}

	req := starr.Request{URI: bpDelayProfile, Body: &body}
	if err := s.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateDelayProfile updates the delay profile.
func (s *Sonarr) UpdateDelayProfile(profile *DelayProfile) (*DelayProfile, error) {
	return s.UpdateDelayProfileContext(context.Background(), profile)
}

// UpdateDelayProfileContext updates the delay profile.
func (s *Sonarr) UpdateDelayProfileContext(ctx context.Context, profile *DelayProfile) (*DelayProfile, error) {
	var output DelayProfile

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(profile); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpDelayProfile, err)
	}

	req := starr.Request{URI: path.Join(bpDelayProfile, fmt.Sprint(profile.ID)), Body: &body}
	if err := s.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteDelayProfile removes a single delay profile.
func (s *Sonarr) DeleteDelayProfile(profileID int64) error {
	return s.DeleteDelayProfileContext(context.Background(), profileID)
}

// DeleteDelayProfileContext removes a single delay profile.
func (s *Sonarr) DeleteDelayProfileContext(ctx context.Context, profileID int64) error {
	req := starr.Request{URI: path.Join(bpDelayProfile, fmt.Sprint(profileID))}
	if err := s.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
