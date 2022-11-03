package prowlarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"

	"golift.io/starr"
)

// Define Base Path for notification calls.
const bpNotification = APIver + "/notification"

// NotificationInput is the input for a new or updated notification.
type NotificationInput struct {
	OnHealthIssue         bool                `json:"onHealthIssue,omitempty"`
	OnApplicationUpdate   bool                `json:"onApplicationUpdate,omitempty"`
	IncludeHealthWarnings bool                `json:"includeHealthWarnings,omitempty"`
	ID                    int64               `json:"id,omitempty"`
	Name                  string              `json:"name"`
	Implementation        string              `json:"implementation"`
	ConfigContract        string              `json:"configContract"`
	Tags                  []int               `json:"tags,omitempty"`
	Fields                []*starr.FieldInput `json:"fields"`
}

// NotificationOutput is the output from the notification methods.
type NotificationOutput struct {
	OnHealthIssue               bool                 `json:"onHealthIssue,omitempty"`
	OnApplicationUpdate         bool                 `json:"onApplicationUpdate,omitempty"`
	SupportsOnApplicationUpdate bool                 `json:"supportsOnApplicationUpdate"`
	SupportsOnHealthIssue       bool                 `json:"supportsOnHealthIssue"`
	IncludeHealthWarnings       bool                 `json:"includeHealthWarnings"`
	ID                          int64                `json:"id"`
	Name                        string               `json:"name"`
	ImplementationName          string               `json:"implementationName"`
	Implementation              string               `json:"implementation"`
	ConfigContract              string               `json:"configContract"`
	InfoLink                    string               `json:"infoLink"`
	Tags                        []int                `json:"tags"`
	Fields                      []*starr.FieldOutput `json:"fields"`
}

// GetNotifications returns all configured notifications.
func (p *Prowlarr) GetNotifications() ([]*NotificationOutput, error) {
	return p.GetNotificationsContext(context.Background())
}

// GetNotificationsContext returns all configured notifications.
func (p *Prowlarr) GetNotificationsContext(ctx context.Context) ([]*NotificationOutput, error) {
	var output []*NotificationOutput

	req := starr.Request{URI: bpNotification}
	if err := p.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetNotification returns a single notification.
func (p *Prowlarr) GetNotification(notificationID int) (*NotificationOutput, error) {
	return p.GetNotificationContext(context.Background(), notificationID)
}

// GetNotificationContext returns a single notification.
func (p *Prowlarr) GetNotificationContext(ctx context.Context, notificationID int) (*NotificationOutput, error) {
	var output NotificationOutput

	req := starr.Request{URI: path.Join(bpNotification, fmt.Sprint(notificationID))}
	if err := p.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// AddNotification creates a notification.
func (p *Prowlarr) AddNotification(notification *NotificationInput) (*NotificationOutput, error) {
	return p.AddNotificationContext(context.Background(), notification)
}

// AddNotificationContext creates a notification.
func (p *Prowlarr) AddNotificationContext(ctx context.Context, client *NotificationInput) (*NotificationOutput, error) {
	var output NotificationOutput

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(client); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpNotification, err)
	}

	req := starr.Request{URI: bpNotification, Body: &body}
	if err := p.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateNotification updates the notification.
func (p *Prowlarr) UpdateNotification(notification *NotificationInput) (*NotificationOutput, error) {
	return p.UpdateNotificationContext(context.Background(), notification)
}

// UpdateNotificationContext updates the notification.
func (p *Prowlarr) UpdateNotificationContext(ctx context.Context,
	client *NotificationInput,
) (*NotificationOutput, error) {
	var output NotificationOutput

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(client); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpNotification, err)
	}

	req := starr.Request{URI: path.Join(bpNotification, fmt.Sprint(client.ID)), Body: &body}
	if err := p.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteNotification removes a single notification.
func (p *Prowlarr) DeleteNotification(notificationID int64) error {
	return p.DeleteNotificationContext(context.Background(), notificationID)
}

func (p *Prowlarr) DeleteNotificationContext(ctx context.Context, notificationID int64) error {
	req := starr.Request{URI: path.Join(bpNotification, fmt.Sprint(notificationID))}
	if err := p.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
