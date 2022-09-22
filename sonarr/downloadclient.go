package sonarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"

	"golift.io/starr"
)

// Define Base Path for downloadClient calls.
const bpDownloadClient = APIver + "/downloadClient"

// DownloadClientInput is the input for a new or updated downladClient.
type DownloadClientInput struct {
	Enable                   bool                `json:"enable"`
	RemoveCompletedDownloads bool                `json:"removeCompletedDownloads"`
	RemoveFailedDownloads    bool                `json:"removeFailedDownloads"`
	Priority                 int                 `json:"priority"`
	ID                       int64               `json:"id,omitempty"`
	ConfigContract           string              `json:"configContract"`
	Implementation           string              `json:"implementation"`
	Name                     string              `json:"name"`
	Protocol                 string              `json:"protocol"`
	Tags                     []int               `json:"tags"`
	Fields                   []*starr.FieldInput `json:"fields"`
}

// DownloadClientOutput is the output from the downladClient methods.
type DownloadClientOutput struct {
	Enable                   bool                 `json:"enable"`
	RemoveCompletedDownloads bool                 `json:"removeCompletedDownloads"`
	RemoveFailedDownloads    bool                 `json:"removeFailedDownloads"`
	Priority                 int                  `json:"priority"`
	ID                       int64                `json:"id,omitempty"`
	ConfigContract           string               `json:"configContract"`
	Implementation           string               `json:"implementation"`
	ImplementationName       string               `json:"implementationName"`
	InfoLink                 string               `json:"infoLink"`
	Name                     string               `json:"name"`
	Protocol                 string               `json:"protocol"`
	Tags                     []int                `json:"tags"`
	Fields                   []*starr.FieldOutput `json:"fields"`
}

// GetDownloadClients returns all configured downloadclients.
func (s *Sonarr) GetDownloadClients() ([]*DownloadClientOutput, error) {
	return s.GetDownloadClientsContext(context.Background())
}

// GetDownloadClientsContext returns all configured downloadclients.
func (s *Sonarr) GetDownloadClientsContext(ctx context.Context) ([]*DownloadClientOutput, error) {
	var output []*DownloadClientOutput

	req := starr.Request{URI: bpDownloadClient}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get%s): %w", &req, err)
	}

	return output, nil
}

// GetDownloadClient returns a single downloadclient.
func (s *Sonarr) GetDownloadClient(downloadclientID int) (*DownloadClientOutput, error) {
	return s.GetDownloadClientContext(context.Background(), downloadclientID)
}

// GetDownloadClientContext returns a single downloadclient.
func (s *Sonarr) GetDownloadClientContext(ctx context.Context, downloadclientID int) (*DownloadClientOutput, error) {
	var output DownloadClientOutput

	req := starr.Request{URI: path.Join(bpDownloadClient, fmt.Sprint(downloadclientID))}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get%s): %w", &req, err)
	}

	return &output, nil
}

// AddDownloadClient creates a downloadclient.
func (s *Sonarr) AddDownloadClient(downloadclient *DownloadClientInput) (*DownloadClientOutput, error) {
	return s.AddDownloadClientContext(context.Background(), downloadclient)
}

// AddDownloadClientContext creates a downloadclient.
func (s *Sonarr) AddDownloadClientContext(ctx context.Context,
	client *DownloadClientInput,
) (*DownloadClientOutput, error) {
	var output DownloadClientOutput

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(client); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpDownloadClient, err)
	}

	req := starr.Request{URI: bpDownloadClient, Body: &body}
	if err := s.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateDownloadClient updates the downloadclient.
func (s *Sonarr) UpdateDownloadClient(downloadclient *DownloadClientInput) (*DownloadClientOutput, error) {
	return s.UpdateDownloadClientContext(context.Background(), downloadclient)
}

// UpdateDownloadClientContext updates the downloadclient.
func (s *Sonarr) UpdateDownloadClientContext(ctx context.Context,
	client *DownloadClientInput,
) (*DownloadClientOutput, error) {
	var output DownloadClientOutput

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(client); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpDownloadClient, err)
	}

	req := starr.Request{URI: path.Join(bpDownloadClient, fmt.Sprint(int(client.ID))), Body: &body}
	if err := s.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteDownloadClient removes a single downloadclient.
func (s *Sonarr) DeleteDownloadClient(downloadclientID int) error {
	return s.DeleteDownloadClientContext(context.Background(), downloadclientID)
}

// DeleteDownloadClientContext removes a single downloadclient.
func (s *Sonarr) DeleteDownloadClientContext(ctx context.Context, downloadclientID int) error {
	req := starr.Request{URI: path.Join(bpDownloadClient, fmt.Sprint(downloadclientID))}
	if err := s.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete%s): %w", &req, err)
	}

	return nil
}
