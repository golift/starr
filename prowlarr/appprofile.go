package prowlarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"

	"golift.io/starr"
)

const bpAppProfile = APIver + "/appprofile"

// AppProfile is a Prowlarr application profile.
type AppProfile struct {
	ID                      int64  `json:"id,omitempty"`
	Name                    string `json:"name,omitempty"`
	EnableRss               bool   `json:"enableRss"`
	EnableAutomaticSearch   bool   `json:"enableAutomaticSearch"`
	EnableInteractiveSearch bool   `json:"enableInteractiveSearch"`
	MinimumSeeders          int    `json:"minimumSeeders,omitempty"`
}

// GetAppProfiles returns all application profiles.
func (p *Prowlarr) GetAppProfiles() ([]*AppProfile, error) {
	return p.GetAppProfilesContext(context.Background())
}

// GetAppProfilesContext returns all application profiles.
func (p *Prowlarr) GetAppProfilesContext(ctx context.Context) ([]*AppProfile, error) {
	var output []*AppProfile

	req := starr.Request{URI: bpAppProfile}
	if err := p.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetAppProfile returns a single application profile.
func (p *Prowlarr) GetAppProfile(id int64) (*AppProfile, error) {
	return p.GetAppProfileContext(context.Background(), id)
}

// GetAppProfileContext returns a single application profile.
func (p *Prowlarr) GetAppProfileContext(ctx context.Context, id int64) (*AppProfile, error) {
	var output AppProfile

	req := starr.Request{URI: path.Join(bpAppProfile, starr.Str(id))}
	if err := p.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// GetAppProfileSchema returns default application profile templates.
func (p *Prowlarr) GetAppProfileSchema() (*AppProfile, error) {
	return p.GetAppProfileSchemaContext(context.Background())
}

// GetAppProfileSchemaContext returns default application profile templates.
func (p *Prowlarr) GetAppProfileSchemaContext(ctx context.Context) (*AppProfile, error) {
	var output AppProfile

	req := starr.Request{URI: path.Join(bpAppProfile, "schema")}
	if err := p.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// AddAppProfile creates an application profile.
func (p *Prowlarr) AddAppProfile(profile *AppProfile) (*AppProfile, error) {
	return p.AddAppProfileContext(context.Background(), profile)
}

// AddAppProfileContext creates an application profile.
func (p *Prowlarr) AddAppProfileContext(ctx context.Context, profile *AppProfile) (*AppProfile, error) {
	var output AppProfile

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(profile); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpAppProfile, err)
	}

	req := starr.Request{URI: bpAppProfile, Body: &body}
	if err := p.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateAppProfile updates an application profile.
func (p *Prowlarr) UpdateAppProfile(profile *AppProfile) (*AppProfile, error) {
	return p.UpdateAppProfileContext(context.Background(), profile)
}

// UpdateAppProfileContext updates an application profile.
func (p *Prowlarr) UpdateAppProfileContext(ctx context.Context, profile *AppProfile) (*AppProfile, error) {
	var output AppProfile

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(profile); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpAppProfile, err)
	}

	req := starr.Request{URI: path.Join(bpAppProfile, starr.Str(profile.ID)), Body: &body}
	if err := p.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteAppProfile removes an application profile.
func (p *Prowlarr) DeleteAppProfile(id int64) error {
	return p.DeleteAppProfileContext(context.Background(), id)
}

// DeleteAppProfileContext removes an application profile.
func (p *Prowlarr) DeleteAppProfileContext(ctx context.Context, id int64) error {
	req := starr.Request{URI: path.Join(bpAppProfile, starr.Str(id))}
	if err := p.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
