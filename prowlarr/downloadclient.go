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

// Define Base Path for download client calls.
const bpDownloadClient = APIver + "/downloadClient"

// DownloadClientInput is the input for a new or updated download client.
type DownloadClientInput struct {
	Enable         bool                `json:"enable"`
	Priority       int                 `json:"priority"`
	ID             int64               `json:"id,omitempty"`
	ConfigContract string              `json:"configContract"`
	Implementation string              `json:"implementation"`
	Name           string              `json:"name"`
	Protocol       starr.Protocol      `json:"protocol"`
	Tags           []int               `json:"tags"`
	Fields         []*starr.FieldInput `json:"fields"`
}

// DownloadClientOutput is the output from the download client methods.
type DownloadClientOutput struct {
	Enable             bool                 `json:"enable"`
	Priority           int                  `json:"priority"`
	ID                 int64                `json:"id,omitempty"`
	ConfigContract     string               `json:"configContract"`
	Implementation     string               `json:"implementation"`
	ImplementationName string               `json:"implementationName"`
	InfoLink           string               `json:"infoLink"`
	Name               string               `json:"name"`
	Protocol           starr.Protocol       `json:"protocol"`
	Tags               []int                `json:"tags"`
	Fields             []*starr.FieldOutput `json:"fields"`
}

// GetDownloadClients returns all configured download clients.
func (p *Prowlarr) GetDownloadClients() ([]*DownloadClientOutput, error) {
	return p.GetDownloadClientsContext(context.Background())
}

// GetDownloadClientsContext returns all configured download clients.
func (p *Prowlarr) GetDownloadClientsContext(ctx context.Context) ([]*DownloadClientOutput, error) {
	var output []*DownloadClientOutput

	req := starr.Request{URI: bpDownloadClient}
	if err := p.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetDownloadClient returns a single download client.
func (p *Prowlarr) GetDownloadClient(downloadclientID int64) (*DownloadClientOutput, error) {
	return p.GetDownloadClientContext(context.Background(), downloadclientID)
}

// GetDownloadClientContext returns a single download client.
func (p *Prowlarr) GetDownloadClientContext(ctx context.Context, clientID int64) (*DownloadClientOutput, error) {
	var output DownloadClientOutput

	req := starr.Request{URI: path.Join(bpDownloadClient, fmt.Sprint(clientID))}
	if err := p.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// AddDownloadClient creates a download client without testing it.
func (p *Prowlarr) AddDownloadClient(downloadclient *DownloadClientInput) (*DownloadClientOutput, error) {
	return p.AddDownloadClientContext(context.Background(), downloadclient)
}

// AddDownloadClientContext creates a download client without testing it.
func (p *Prowlarr) AddDownloadClientContext(ctx context.Context,
	client *DownloadClientInput,
) (*DownloadClientOutput, error) {
	var (
		output DownloadClientOutput
		body   bytes.Buffer
	)

	client.ID = 0
	if err := json.NewEncoder(&body).Encode(client); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpDownloadClient, err)
	}

	req := starr.Request{URI: bpDownloadClient, Body: &body, Query: url.Values{"forceSave": []string{"true"}}}
	if err := p.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// TestDownloadClient tests a download client.
func (p *Prowlarr) TestDownloadClient(client *DownloadClientInput) error {
	return p.TestDownloadClientContext(context.Background(), client)
}

// TestDownloadClientContext tests a download client.
func (p *Prowlarr) TestDownloadClientContext(ctx context.Context, client *DownloadClientInput) error {
	var output interface{}

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(client); err != nil {
		return fmt.Errorf("json.Marshal(%s): %w", bpDownloadClient, err)
	}

	req := starr.Request{URI: path.Join(bpDownloadClient, "test"), Body: &body}
	if err := p.PostInto(ctx, req, &output); err != nil {
		return fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return nil
}

// UpdateDownloadClient updates the download client.
func (p *Prowlarr) UpdateDownloadClient(client *DownloadClientInput, force bool) (*DownloadClientOutput, error) {
	return p.UpdateDownloadClientContext(context.Background(), client, force)
}

// UpdateDownloadClientContext updates the download client.
func (p *Prowlarr) UpdateDownloadClientContext(ctx context.Context,
	client *DownloadClientInput,
	force bool,
) (*DownloadClientOutput, error) {
	var output DownloadClientOutput

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(client); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpDownloadClient, err)
	}

	req := starr.Request{
		URI:   path.Join(bpDownloadClient, fmt.Sprint(client.ID)),
		Body:  &body,
		Query: url.Values{"forceSave": []string{fmt.Sprint(force)}},
	}
	if err := p.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteDownloadClient removes a single download client.
func (p *Prowlarr) DeleteDownloadClient(downloadclientID int64) error {
	return p.DeleteDownloadClientContext(context.Background(), downloadclientID)
}

// DeleteDownloadClientContext removes a single download client.
func (p *Prowlarr) DeleteDownloadClientContext(ctx context.Context, downloadclientID int64) error {
	req := starr.Request{URI: path.Join(bpDownloadClient, fmt.Sprint(downloadclientID))}
	if err := p.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
