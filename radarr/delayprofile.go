package radarr

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
	EnableUsenet           bool           `json:"enableUsenet,omitempty"`
	EnableTorrent          bool           `json:"enableTorrent,omitempty"`
	BypassIfHighestQuality bool           `json:"bypassIfHighestQuality,omitempty"`
	UsenetDelay            int64          `json:"usenetDelay,omitempty"`
	TorrentDelay           int64          `json:"torrentDelay,omitempty"`
	ID                     int64          `json:"id,omitempty"`
	Order                  int64          `json:"order,omitempty"`
	Tags                   []int          `json:"tags"`
	PreferredProtocol      starr.Protocol `json:"preferredProtocol,omitempty"`
}

// GetDelayProfiles returns all configured delay profiles.
func (r *Radarr) GetDelayProfiles() ([]*DelayProfile, error) {
	return r.GetDelayProfilesContext(context.Background())
}

// GetDelayProfilesContext returns all configured delay profiles.
func (r *Radarr) GetDelayProfilesContext(ctx context.Context) ([]*DelayProfile, error) {
	var output []*DelayProfile

	req := starr.Request{URI: bpDelayProfile}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetDelayProfile returns a single delay profile.
func (r *Radarr) GetDelayProfile(profileID int64) (*DelayProfile, error) {
	return r.GetDelayProfileContext(context.Background(), profileID)
}

// GetDelayProfileContext returns a single delay profile.
func (r *Radarr) GetDelayProfileContext(ctx context.Context, profileID int64) (*DelayProfile, error) {
	var output DelayProfile

	req := starr.Request{URI: path.Join(bpDelayProfile, fmt.Sprint(profileID))}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// AddDelayProfile creates a delay profile.
// AddDelayProfile doesn't take into account the "order" field sent on creation.
// Order will be set to first available. This can only be edited via UpdateDelayProfile later on.
func (r *Radarr) AddDelayProfile(profile *DelayProfile) (*DelayProfile, error) {
	return r.AddDelayProfileContext(context.Background(), profile)
}

// AddDelayProfileContext creates a delay profile.
func (r *Radarr) AddDelayProfileContext(ctx context.Context, profile *DelayProfile) (*DelayProfile, error) {
	var output DelayProfile

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(profile); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpDelayProfile, err)
	}

	req := starr.Request{URI: bpDelayProfile, Body: &body}
	if err := r.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateDelayProfile updates the delay profile.
func (r *Radarr) UpdateDelayProfile(profile *DelayProfile) (*DelayProfile, error) {
	return r.UpdateDelayProfileContext(context.Background(), profile)
}

// UpdateDelayProfileContext updates the delay profile.
func (r *Radarr) UpdateDelayProfileContext(ctx context.Context, profile *DelayProfile) (*DelayProfile, error) {
	var output DelayProfile

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(profile); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpDelayProfile, err)
	}

	req := starr.Request{URI: path.Join(bpDelayProfile, fmt.Sprint(profile.ID)), Body: &body}
	if err := r.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteDelayProfile removes a single delay profile.
func (r *Radarr) DeleteDelayProfile(profileID int64) error {
	return r.DeleteDelayProfileContext(context.Background(), profileID)
}

// DeleteDelayProfileContext removes a single delay profile.
func (r *Radarr) DeleteDelayProfileContext(ctx context.Context, profileID int64) error {
	req := starr.Request{URI: path.Join(bpDelayProfile, fmt.Sprint(profileID))}
	if err := r.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
