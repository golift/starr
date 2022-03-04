package sonarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
)

// DelayProfile is the /api/v3/delayprofile endpoint.
type DelayProfile struct {
	EnableUsenet           bool   `json:"enableUsenet"`
	EnableTorrent          bool   `json:"enableTorrent"`
	BypassIfHighestQuality bool   `json:"bypassIfHighestQuality"`
	UsenetDelay            int64  `json:"usenetDelay"`
	TorrentDelay           int64  `json:"torrentDelay"`
	ID                     int64  `json:"id,omitempty"`
	Order                  int64  `json:"order"`
	Tags                   []int  `json:"tags"`
	PreferredProtocol      string `json:"preferredProtocol"`
}

// GetDelayProfiles returns all configured delay profiles.
func (s *Sonarr) GetDelayProfiles() ([]*DelayProfile, error) {
	return s.GetDelayProfilesContext(context.Background())
}

func (s *Sonarr) GetDelayProfilesContext(ctx context.Context) ([]*DelayProfile, error) {
	var output []*DelayProfile

	if _, err := s.GetInto(ctx, "v3/delayProfile", nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(delayProfile): %w", err)
	}

	return output, nil
}

// GetDelayProfile returns a single delay profile.
func (s *Sonarr) GetDelayProfile(profileID int) (*DelayProfile, error) {
	return s.GetDelayProfileContext(context.Background(), profileID)
}

func (s *Sonarr) GetDelayProfileContext(ctx context.Context, profileID int) (*DelayProfile, error) {
	var output *DelayProfile

	id := strconv.Itoa(profileID)
	if _, err := s.GetInto(ctx, "v3/delayProfile/"+id, nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(delayProfile): %w", err)
	}

	return output, nil
}

// AddDelayProfile creates a delay profile.
func (s *Sonarr) AddDelayProfile(profile *DelayProfile) (*DelayProfile, error) {
	return s.AddDelayProfileContext(context.Background(), profile)
}

func (s *Sonarr) AddDelayProfileContext(ctx context.Context, profile *DelayProfile) (*DelayProfile, error) {
	var output DelayProfile

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(profile); err != nil {
		return nil, fmt.Errorf("json.Marshal(delayProfile): %w", err)
	}

	if _, err := s.PostInto(ctx, "v3/delayProfile", nil, &body, &output); err != nil {
		return nil, fmt.Errorf("api.Post(delayProfile): %w", err)
	}

	return &output, nil
}

// UpdateDelayProfile updates the delay profile.
func (s *Sonarr) UpdateDelayProfile(profile *DelayProfile) (*DelayProfile, error) {
	return s.UpdateDelayProfileContext(context.Background(), profile)
}

func (s *Sonarr) UpdateDelayProfileContext(ctx context.Context, profile *DelayProfile) (*DelayProfile, error) {
	var output DelayProfile

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(profile); err != nil {
		return nil, fmt.Errorf("json.Marshal(delayProfile): %w", err)
	}

	id := strconv.Itoa(int(profile.ID))
	if _, err := s.PutInto(ctx, "v3/delayProfile/"+id, nil, &body, &output); err != nil {
		return nil, fmt.Errorf("api.Put(delayProfile): %w", err)
	}

	return &output, nil
}

// DeleteDelayProfile removes a single delay profile.
func (s *Sonarr) DeleteDelayProfile(profileID int) error {
	return s.DeleteDelayProfileContext(context.Background(), profileID)
}

func (s *Sonarr) DeleteDelayProfileContext(ctx context.Context, profileID int) error {
	id := strconv.Itoa(profileID)
	if _, err := s.Delete(ctx, "v3/delayProfile/"+id, nil); err != nil {
		return fmt.Errorf("api.Delete(delayProfile): %w", err)
	}

	return nil
}