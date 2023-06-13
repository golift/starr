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
	OnGrab                      bool                `json:"onGrab"`
	OnHealthIssue               bool                `json:"onHealthIssue"`
	OnHealthRestored            bool                `json:"onHealthRestored"`
	OnApplicationUpdate         bool                `json:"onApplicationUpdate"`
	SupportsOnGrab              bool                `json:"supportsOnGrab"`
	IncludeManualGrabs          bool                `json:"includeManualGrabs"`
	SupportsOnHealthIssue       bool                `json:"supportsOnHealthIssue"`
	SupportsOnHealthRestored    bool                `json:"supportsOnHealthRestored"`
	IncludeHealthWarnings       bool                `json:"includeHealthWarnings"`
	SupportsOnApplicationUpdate bool                `json:"supportsOnApplicationUpdate"`
	ID                          int64               `json:"id,omitempty"` // update only
	Name                        string              `json:"name"`
	ImplementationName          string              `json:"implementationName"`
	Implementation              string              `json:"implementation"`
	ConfigContract              string              `json:"configContract"`
	InfoLink                    string              `json:"infoLink"`
	Tags                        []int               `json:"tags"`
	Fields                      []*starr.FieldInput `json:"fields"`
}

// NotificationOutput is the output from the notification methods.
type NotificationOutput struct {
	OnGrab                      bool                 `json:"onGrab"`
	OnHealthIssue               bool                 `json:"onHealthIssue"`
	OnHealthRestored            bool                 `json:"onHealthRestored"`
	OnApplicationUpdate         bool                 `json:"onApplicationUpdate"`
	SupportsOnGrab              bool                 `json:"supportsOnGrab"`
	IncludeManualGrabs          bool                 `json:"includeManualGrabs"`
	SupportsOnHealthIssue       bool                 `json:"supportsOnHealthIssue"`
	SupportsOnHealthRestored    bool                 `json:"supportsOnHealthRestored"`
	IncludeHealthWarnings       bool                 `json:"includeHealthWarnings"`
	SupportsOnApplicationUpdate bool                 `json:"supportsOnApplicationUpdate"`
	ID                          int64                `json:"id"`
	Name                        string               `json:"name"`
	ImplementationName          string               `json:"implementationName"`
	Implementation              string               `json:"implementation"`
	ConfigContract              string               `json:"configContract"`
	InfoLink                    string               `json:"infoLink"`
	Tags                        []int                `json:"tags"`
	Fields                      []*starr.FieldOutput `json:"fields"`
	Message                     struct {
		Message string `json:"message"` // this is a weird place for a message
		Type    string `json:"type"`
	} `json:"message"`
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
