package prowlarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"path"

	"golift.io/starr"
)

const bpApplication = APIver + "/applications"

// ApplicationInput is used to create or update a connected application.
type ApplicationInput struct {
	ID             int64               `json:"id,omitempty"`
	Name           string              `json:"name,omitempty"`
	SyncLevel      string              `json:"syncLevel,omitempty"`
	Implementation string              `json:"implementation,omitempty"`
	ConfigContract string              `json:"configContract,omitempty"`
	AppProfileID   int64               `json:"appProfileId,omitempty"`
	Tags           []int               `json:"tags,omitempty"`
	Fields         []*starr.FieldInput `json:"fields,omitempty"`
}

// ApplicationOutput is returned from application endpoints.
type ApplicationOutput struct {
	ID                 int64                `json:"id,omitempty"`
	Name               string               `json:"name,omitempty"`
	SyncLevel          string               `json:"syncLevel,omitempty"`
	Implementation     string               `json:"implementation,omitempty"`
	ImplementationName string               `json:"implementationName,omitempty"`
	ConfigContract     string               `json:"configContract,omitempty"`
	AppProfileID       int64                `json:"appProfileId,omitempty"`
	Tags               []int                `json:"tags,omitempty"`
	Fields             []*starr.FieldOutput `json:"fields,omitempty"`
}

// GetApplications returns all connected applications.
func (p *Prowlarr) GetApplications() ([]*ApplicationOutput, error) {
	return p.GetApplicationsContext(context.Background())
}

// GetApplicationsContext returns all connected applications.
func (p *Prowlarr) GetApplicationsContext(ctx context.Context) ([]*ApplicationOutput, error) {
	var output []*ApplicationOutput

	req := starr.Request{URI: bpApplication}
	if err := p.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetApplication returns a single application.
func (p *Prowlarr) GetApplication(id int64) (*ApplicationOutput, error) {
	return p.GetApplicationContext(context.Background(), id)
}

// GetApplicationContext returns a single application.
func (p *Prowlarr) GetApplicationContext(ctx context.Context, id int64) (*ApplicationOutput, error) {
	var output ApplicationOutput

	req := starr.Request{URI: path.Join(bpApplication, starr.Str(id))}
	if err := p.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// AddApplication creates a connected application.
func (p *Prowlarr) AddApplication(app *ApplicationInput, forceSave bool) (*ApplicationOutput, error) {
	return p.AddApplicationContext(context.Background(), app, forceSave)
}

// AddApplicationContext creates a connected application.
func (p *Prowlarr) AddApplicationContext(
	ctx context.Context,
	app *ApplicationInput,
	forceSave bool,
) (*ApplicationOutput, error) {
	var output ApplicationOutput

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(app); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpApplication, err)
	}

	req := starr.Request{
		URI:   bpApplication,
		Body:  &body,
		Query: url.Values{"forceSave": []string{starr.Str(forceSave)}},
	}
	if err := p.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateApplication updates a connected application.
func (p *Prowlarr) UpdateApplication(app *ApplicationInput, forceSave bool) (*ApplicationOutput, error) {
	return p.UpdateApplicationContext(context.Background(), app, forceSave)
}

// UpdateApplicationContext updates a connected application.
func (p *Prowlarr) UpdateApplicationContext(
	ctx context.Context,
	app *ApplicationInput,
	forceSave bool,
) (*ApplicationOutput, error) {
	var output ApplicationOutput

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(app); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpApplication, err)
	}

	req := starr.Request{
		URI:   path.Join(bpApplication, starr.Str(app.ID)),
		Body:  &body,
		Query: url.Values{"forceSave": []string{starr.Str(forceSave)}},
	}
	if err := p.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteApplication removes a connected application.
func (p *Prowlarr) DeleteApplication(id int64) error {
	return p.DeleteApplicationContext(context.Background(), id)
}

// DeleteApplicationContext removes a connected application.
func (p *Prowlarr) DeleteApplicationContext(ctx context.Context, id int64) error {
	req := starr.Request{URI: path.Join(bpApplication, starr.Str(id))}
	if err := p.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}

// TestApplication tests connection settings for an application definition.
func (p *Prowlarr) TestApplication(app *ApplicationInput, forceTest bool) error {
	return p.TestApplicationContext(context.Background(), app, forceTest)
}

// TestApplicationContext tests connection settings for an application definition.
func (p *Prowlarr) TestApplicationContext(ctx context.Context, app *ApplicationInput, forceTest bool) error {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(app); err != nil {
		return fmt.Errorf("json.Marshal(%s): %w", path.Join(bpApplication, "test"), err)
	}

	var output any

	req := starr.Request{
		URI:   path.Join(bpApplication, "test"),
		Body:  &body,
		Query: url.Values{"forceTest": []string{starr.Str(forceTest)}},
	}
	if err := p.PostInto(ctx, req, &output); err != nil {
		return fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return nil
}
