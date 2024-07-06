package lidarr

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
	OnGrab                bool                `json:"onGrab,omitempty"`
	OnReleaseImport       bool                `json:"onReleaseImport,omitempty"`
	OnUpgrade             bool                `json:"onUpgrade,omitempty"`
	OnRename              bool                `json:"onRename,omitempty"`
	OnTrackRetag          bool                `json:"onTrackRetag,omitempty"`
	OnHealthIssue         bool                `json:"onHealthIssue,omitempty"`
	OnDownloadFailure     bool                `json:"onDownloadFailure,omitempty"`
	OnImportFailure       bool                `json:"onImportFailure,omitempty"`
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
	OnGrab                      bool                 `json:"onGrab,omitempty"`
	OnReleaseImport             bool                 `json:"onReleaseImport,omitempty"`
	OnUpgrade                   bool                 `json:"onUpgrade,omitempty"`
	OnRename                    bool                 `json:"onRename,omitempty"`
	OnTrackRetag                bool                 `json:"onTrackRetag,omitempty"`
	OnHealthIssue               bool                 `json:"onHealthIssue,omitempty"`
	OnDownloadFailure           bool                 `json:"onDownloadFailure,omitempty"`
	OnImportFailure             bool                 `json:"onImportFailure,omitempty"`
	OnApplicationUpdate         bool                 `json:"onApplicationUpdate,omitempty"`
	SupportsOnGrab              bool                 `json:"supportsOnGrab"`
	SupportsOnReleaseImport     bool                 `json:"supportsOnReleaseImport"`
	SupportsOnUpgrade           bool                 `json:"supportsOnUpgrade"`
	SupportsOnRename            bool                 `json:"supportsOnRename"`
	SupportsOnApplicationUpdate bool                 `json:"supportsOnApplicationUpdate"`
	SupportsOnDownloadFailure   bool                 `json:"supportsOnDownloadFailure"`
	SupportsOnImportFailure     bool                 `json:"supportsOnImportFailure"`
	SupportsOnTrackRetag        bool                 `json:"supportsOnTrackRetag"`
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
func (l *Lidarr) GetNotifications() ([]*NotificationOutput, error) {
	return l.GetNotificationsContext(context.Background())
}

// GetNotificationsContext returns all configured notifications.
func (l *Lidarr) GetNotificationsContext(ctx context.Context) ([]*NotificationOutput, error) {
	var output []*NotificationOutput

	req := starr.Request{URI: bpNotification}
	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetNotification returns a single notification.
func (l *Lidarr) GetNotification(notificationID int) (*NotificationOutput, error) {
	return l.GetNotificationContext(context.Background(), notificationID)
}

// GetNotificationContext returns a single notification.
func (l *Lidarr) GetNotificationContext(ctx context.Context, notificationID int) (*NotificationOutput, error) {
	var output NotificationOutput

	req := starr.Request{URI: path.Join(bpNotification, starr.Str(notificationID))}
	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// AddNotification creates a notification.
func (l *Lidarr) AddNotification(notification *NotificationInput) (*NotificationOutput, error) {
	return l.AddNotificationContext(context.Background(), notification)
}

// AddNotificationContext creates a notification.
func (l *Lidarr) AddNotificationContext(ctx context.Context, client *NotificationInput) (*NotificationOutput, error) {
	var output NotificationOutput

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(client); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpNotification, err)
	}

	req := starr.Request{URI: bpNotification, Body: &body}
	if err := l.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateNotification updates the notification.
func (l *Lidarr) UpdateNotification(notification *NotificationInput) (*NotificationOutput, error) {
	return l.UpdateNotificationContext(context.Background(), notification)
}

// UpdateNotificationContext updates the notification.
func (l *Lidarr) UpdateNotificationContext(ctx context.Context,
	client *NotificationInput,
) (*NotificationOutput, error) {
	var output NotificationOutput

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(client); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpNotification, err)
	}

	req := starr.Request{URI: path.Join(bpNotification, starr.Str(client.ID)), Body: &body}
	if err := l.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteNotification removes a single notification.
func (l *Lidarr) DeleteNotification(notificationID int64) error {
	return l.DeleteNotificationContext(context.Background(), notificationID)
}

func (l *Lidarr) DeleteNotificationContext(ctx context.Context, notificationID int64) error {
	req := starr.Request{URI: path.Join(bpNotification, starr.Str(notificationID))}
	if err := l.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
