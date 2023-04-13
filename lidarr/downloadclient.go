package lidarr

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

// DownloadClientOutput is the output from the download client methods.
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

// GetDownloadClients returns all configured download clients.
func (l *Lidarr) GetDownloadClients() ([]*DownloadClientOutput, error) {
	return l.GetDownloadClientsContext(context.Background())
}

// GetDownloadClientsContext returns all configured download clients.
func (l *Lidarr) GetDownloadClientsContext(ctx context.Context) ([]*DownloadClientOutput, error) {
	var output []*DownloadClientOutput

	req := starr.Request{URI: bpDownloadClient}
	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetDownloadClient returns a single download client.
func (l *Lidarr) GetDownloadClient(downloadclientID int64) (*DownloadClientOutput, error) {
	return l.GetDownloadClientContext(context.Background(), downloadclientID)
}

// GetDownloadClientContext returns a single download client.
func (l *Lidarr) GetDownloadClientContext(ctx context.Context, downloadclientID int64) (*DownloadClientOutput, error) {
	var output DownloadClientOutput

	req := starr.Request{URI: path.Join(bpDownloadClient, fmt.Sprint(downloadclientID))}
	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// AddDownloadClient creates a download client.
func (l *Lidarr) AddDownloadClient(downloadclient *DownloadClientInput) (*DownloadClientOutput, error) {
	return l.AddDownloadClientContext(context.Background(), downloadclient)
}

// AddDownloadClientContext creates a download client.
func (l *Lidarr) AddDownloadClientContext(ctx context.Context,
	client *DownloadClientInput,
) (*DownloadClientOutput, error) {
	var output DownloadClientOutput

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(client); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpDownloadClient, err)
	}

	req := starr.Request{URI: bpDownloadClient, Body: &body}
	if err := l.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateDownloadClient updates the download client.
func (l *Lidarr) UpdateDownloadClient(downloadclient *DownloadClientInput, force bool) (*DownloadClientOutput, error) {
	return l.UpdateDownloadClientContext(context.Background(), downloadclient, force)
}

// UpdateDownloadClientContext updates the download client.
func (l *Lidarr) UpdateDownloadClientContext(ctx context.Context,
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
	if err := l.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteDownloadClient removes a single download client.
func (l *Lidarr) DeleteDownloadClient(downloadclientID int64) error {
	return l.DeleteDownloadClientContext(context.Background(), downloadclientID)
}

// DeleteDownloadClientContext removes a single download client.
func (l *Lidarr) DeleteDownloadClientContext(ctx context.Context, downloadclientID int64) error {
	req := starr.Request{URI: path.Join(bpDownloadClient, fmt.Sprint(downloadclientID))}
	if err := l.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
