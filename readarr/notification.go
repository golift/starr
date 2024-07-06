package readarr

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
	OnGrab                     bool                `json:"onGrab,omitempty"`
	OnReleaseImport            bool                `json:"onReleaseImport,omitempty"`
	OnUpgrade                  bool                `json:"onUpgrade,omitempty"`
	OnRename                   bool                `json:"onRename,omitempty"`
	OnAuthorDelete             bool                `json:"onAuthorDelete,omitempty"`
	OnBookDelete               bool                `json:"onBookDelete,omitempty"`
	OnBookFileDelete           bool                `json:"onBookFileDelete,omitempty"`
	OnBookFileDeleteForUpgrade bool                `json:"onBookFileDeleteForUpgrade,omitempty"`
	OnHealthIssue              bool                `json:"onHealthIssue,omitempty"`
	OnDownloadFailure          bool                `json:"onDownloadFailure,omitempty"`
	OnImportFailure            bool                `json:"onImportFailure,omitempty"`
	OnBookRetag                bool                `json:"onBookRetag,omitempty"`
	OnApplicationUpdate        bool                `json:"onApplicationUpdate,omitempty"`
	IncludeHealthWarnings      bool                `json:"includeHealthWarnings,omitempty"`
	ID                         int64               `json:"id,omitempty"`
	Name                       string              `json:"name"`
	Implementation             string              `json:"implementation"`
	ConfigContract             string              `json:"configContract"`
	Tags                       []int               `json:"tags,omitempty"`
	Fields                     []*starr.FieldInput `json:"fields"`
}

// NotificationOutput is the output from the notification methods.
type NotificationOutput struct {
	OnGrab                             bool                 `json:"onGrab,omitempty"`
	OnReleaseImport                    bool                 `json:"onReleaseImport,omitempty"`
	OnUpgrade                          bool                 `json:"onUpgrade,omitempty"`
	OnRename                           bool                 `json:"onRename,omitempty"`
	OnAuthorDelete                     bool                 `json:"onAuthorDelete,omitempty"`
	OnBookDelete                       bool                 `json:"onBookDelete,omitempty"`
	OnBookFileDelete                   bool                 `json:"onBookFileDelete,omitempty"`
	OnBookFileDeleteForUpgrade         bool                 `json:"onBookFileDeleteForUpgrade,omitempty"`
	OnHealthIssue                      bool                 `json:"onHealthIssue,omitempty"`
	OnDownloadFailure                  bool                 `json:"onDownloadFailure,omitempty"`
	OnImportFailure                    bool                 `json:"onImportFailure,omitempty"`
	OnBookRetag                        bool                 `json:"onBookRetag,omitempty"`
	OnApplicationUpdate                bool                 `json:"onApplicationUpdate,omitempty"`
	SupportsOnGrab                     bool                 `json:"supportsOnGrab"`
	SupportsOnReleaseImport            bool                 `json:"supportsOnReleaseImport"`
	SupportsOnUpgrade                  bool                 `json:"supportsOnUpgrade"`
	SupportsOnRename                   bool                 `json:"supportsOnRename"`
	SupportsOnAuthorDelete             bool                 `json:"supportsOnAuthorDelete"`
	SupportsOnBookDelete               bool                 `json:"supportsOnBookDelete"`
	SupportsOnBookFileDelete           bool                 `json:"supportsOnBookFileDelete"`
	SupportsOnBookFileDeleteForUpgrade bool                 `json:"supportsOnBookFileDeleteForUpgrade"`
	SupportsOnApplicationUpdate        bool                 `json:"supportsOnApplicationUpdate"`
	SupportsOnDownloadFailure          bool                 `json:"supportsOnDownloadFailure"`
	SupportsOnImportFailure            bool                 `json:"supportsOnImportFailure"`
	SupportsOnBookRetag                bool                 `json:"supportsOnBookRetag"`
	SupportsOnHealthIssue              bool                 `json:"supportsOnHealthIssue"`
	IncludeHealthWarnings              bool                 `json:"includeHealthWarnings"`
	ID                                 int64                `json:"id"`
	Name                               string               `json:"name"`
	ImplementationName                 string               `json:"implementationName"`
	Implementation                     string               `json:"implementation"`
	ConfigContract                     string               `json:"configContract"`
	InfoLink                           string               `json:"infoLink"`
	Tags                               []int                `json:"tags"`
	Fields                             []*starr.FieldOutput `json:"fields"`
}

// GetNotifications returns all configured notifications.
func (r *Readarr) GetNotifications() ([]*NotificationOutput, error) {
	return r.GetNotificationsContext(context.Background())
}

// GetNotificationsContext returns all configured notifications.
func (r *Readarr) GetNotificationsContext(ctx context.Context) ([]*NotificationOutput, error) {
	var output []*NotificationOutput

	req := starr.Request{URI: bpNotification}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetNotification returns a single notification.
func (r *Readarr) GetNotification(notificationID int) (*NotificationOutput, error) {
	return r.GetNotificationContext(context.Background(), notificationID)
}

// GetNotificationContext returns a single notification.
func (r *Readarr) GetNotificationContext(ctx context.Context, notificationID int) (*NotificationOutput, error) {
	var output NotificationOutput

	req := starr.Request{URI: path.Join(bpNotification, starr.Str(notificationID))}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// AddNotification creates a notification.
func (r *Readarr) AddNotification(notification *NotificationInput) (*NotificationOutput, error) {
	return r.AddNotificationContext(context.Background(), notification)
}

// AddNotificationContext creates a notification.
func (r *Readarr) AddNotificationContext(ctx context.Context, client *NotificationInput) (*NotificationOutput, error) {
	var output NotificationOutput

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(client); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpNotification, err)
	}

	req := starr.Request{URI: bpNotification, Body: &body}
	if err := r.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateNotification updates the notification.
func (r *Readarr) UpdateNotification(notification *NotificationInput) (*NotificationOutput, error) {
	return r.UpdateNotificationContext(context.Background(), notification)
}

// UpdateNotificationContext updates the notification.
func (r *Readarr) UpdateNotificationContext(ctx context.Context,
	client *NotificationInput,
) (*NotificationOutput, error) {
	var output NotificationOutput

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(client); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpNotification, err)
	}

	req := starr.Request{URI: path.Join(bpNotification, starr.Str(client.ID)), Body: &body}
	if err := r.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteNotification removes a single notification.
func (r *Readarr) DeleteNotification(notificationID int64) error {
	return r.DeleteNotificationContext(context.Background(), notificationID)
}

func (r *Readarr) DeleteNotificationContext(ctx context.Context, notificationID int64) error {
	req := starr.Request{URI: path.Join(bpNotification, starr.Str(notificationID))}
	if err := r.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
