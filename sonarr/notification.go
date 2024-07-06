package sonarr

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
	OnGrab                        bool                `json:"onGrab,omitempty"`
	OnDownload                    bool                `json:"onDownload,omitempty"`
	OnUpgrade                     bool                `json:"onUpgrade,omitempty"`
	OnRename                      bool                `json:"onRename,omitempty"`
	OnSeriesDelete                bool                `json:"onSeriesDelete,omitempty"`
	OnEpisodeFileDelete           bool                `json:"onEpisodeFileDelete,omitempty"`
	OnEpisodeFileDeleteForUpgrade bool                `json:"onEpisodeFileDeleteForUpgrade,omitempty"`
	OnHealthIssue                 bool                `json:"onHealthIssue,omitempty"`
	OnApplicationUpdate           bool                `json:"onApplicationUpdate,omitempty"`
	IncludeHealthWarnings         bool                `json:"includeHealthWarnings,omitempty"`
	ID                            int64               `json:"id,omitempty"`
	Name                          string              `json:"name"`
	Implementation                string              `json:"implementation"`
	ConfigContract                string              `json:"configContract"`
	Tags                          []int               `json:"tags,omitempty"`
	Fields                        []*starr.FieldInput `json:"fields"`
}

// NotificationOutput is the output from the notification methods.
type NotificationOutput struct {
	OnGrab                                bool                 `json:"onGrab"`
	OnDownload                            bool                 `json:"onDownload"`
	OnUpgrade                             bool                 `json:"onUpgrade"`
	OnRename                              bool                 `json:"onRename"`
	OnSeriesDelete                        bool                 `json:"onSeriesDelete"`
	OnEpisodeFileDelete                   bool                 `json:"onEpisodeFileDelete"`
	OnEpisodeFileDeleteForUpgrade         bool                 `json:"onEpisodeFileDeleteForUpgrade"`
	OnHealthIssue                         bool                 `json:"onHealthIssue"`
	OnApplicationUpdate                   bool                 `json:"onApplicationUpdate"`
	SupportsOnGrab                        bool                 `json:"supportsOnGrab"`
	SupportsOnDownload                    bool                 `json:"supportsOnDownload"`
	SupportsOnUpgrade                     bool                 `json:"supportsOnUpgrade"`
	SupportsOnRename                      bool                 `json:"supportsOnRename"`
	SupportsOnSeriesDelete                bool                 `json:"supportsOnSeriesDelete"`
	SupportsOnEpisodeFileDelete           bool                 `json:"supportsOnEpisodeFileDelete"`
	SupportsOnEpisodeFileDeleteForUpgrade bool                 `json:"supportsOnEpisodeFileDeleteForUpgrade"`
	SupportsOnHealthIssue                 bool                 `json:"supportsOnHealthIssue"`
	SupportsOnApplicationUpdate           bool                 `json:"supportsOnApplicationUpdate"`
	IncludeHealthWarnings                 bool                 `json:"includeHealthWarnings"`
	ID                                    int64                `json:"id"`
	Name                                  string               `json:"name"`
	ImplementationName                    string               `json:"implementationName"`
	Implementation                        string               `json:"implementation"`
	ConfigContract                        string               `json:"configContract"`
	InfoLink                              string               `json:"infoLink"`
	Tags                                  []int                `json:"tags"`
	Fields                                []*starr.FieldOutput `json:"fields"`
}

// GetNotifications returns all configured notifications.
func (s *Sonarr) GetNotifications() ([]*NotificationOutput, error) {
	return s.GetNotificationsContext(context.Background())
}

// GetNotificationsContext returns all configured notifications.
func (s *Sonarr) GetNotificationsContext(ctx context.Context) ([]*NotificationOutput, error) {
	var output []*NotificationOutput

	req := starr.Request{URI: bpNotification}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetNotification returns a single notification.
func (s *Sonarr) GetNotification(notificationID int) (*NotificationOutput, error) {
	return s.GetNotificationContext(context.Background(), notificationID)
}

// GetNotificationContext returns a single notification.
func (s *Sonarr) GetNotificationContext(ctx context.Context, notificationID int) (*NotificationOutput, error) {
	var output NotificationOutput

	req := starr.Request{URI: path.Join(bpNotification, starr.Str(notificationID))}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// AddNotification creates a notification.
func (s *Sonarr) AddNotification(notification *NotificationInput) (*NotificationOutput, error) {
	return s.AddNotificationContext(context.Background(), notification)
}

// AddNotificationContext creates a notification.
func (s *Sonarr) AddNotificationContext(ctx context.Context, client *NotificationInput) (*NotificationOutput, error) {
	var output NotificationOutput

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(client); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpNotification, err)
	}

	req := starr.Request{URI: bpNotification, Body: &body}
	if err := s.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateNotification updates the notification.
func (s *Sonarr) UpdateNotification(notification *NotificationInput) (*NotificationOutput, error) {
	return s.UpdateNotificationContext(context.Background(), notification)
}

// UpdateNotificationContext updates the notification.
func (s *Sonarr) UpdateNotificationContext(ctx context.Context,
	client *NotificationInput,
) (*NotificationOutput, error) {
	var output NotificationOutput

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(client); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpNotification, err)
	}

	req := starr.Request{URI: path.Join(bpNotification, starr.Str(client.ID)), Body: &body}
	if err := s.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteNotification removes a single notification.
func (s *Sonarr) DeleteNotification(notificationID int64) error {
	return s.DeleteNotificationContext(context.Background(), notificationID)
}

func (s *Sonarr) DeleteNotificationContext(ctx context.Context, notificationID int64) error {
	req := starr.Request{URI: path.Join(bpNotification, starr.Str(notificationID))}
	if err := s.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
