package sonarr

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"path"
	"time"

	"golift.io/starr"
	"golift.io/starr/starrshared"
)

const bpSystem = APIver + "/system"

// SystemStatus is the /api/v3/system/status endpoint.
type SystemStatus struct {
	AppData                string    `json:"appData"`
	AppName                string    `json:"appName"`
	Authentication         string    `json:"authentication"`
	Branch                 string    `json:"branch"`
	BuildTime              time.Time `json:"buildTime"`
	InstanceName           string    `json:"instanceName"`
	IsAdmin                bool      `json:"isAdmin"`
	IsDebug                bool      `json:"isDebug"`
	IsLinux                bool      `json:"isLinux"`
	IsMono                 bool      `json:"isMono"`
	IsMonoRuntime          bool      `json:"isMonoRuntime"`
	IsOsx                  bool      `json:"isOsx"`
	IsProduction           bool      `json:"isProduction"`
	IsUserInteractive      bool      `json:"isUserInteractive"`
	IsWindows              bool      `json:"isWindows"`
	Mode                   string    `json:"mode"`
	OsName                 string    `json:"osName"`
	OsVersion              string    `json:"osVersion"`
	PackageAuthor          string    `json:"packageAuthor"`
	PackageUpdateMechanism string    `json:"packageUpdateMechanism"`
	PackageVersion         string    `json:"packageVersion"`
	RuntimeName            string    `json:"runtimeName"`
	RuntimeVersion         string    `json:"runtimeVersion"`
	SqliteVersion          string    `json:"sqliteVersion"`
	StartTime              time.Time `json:"startTime"`
	StartupPath            string    `json:"startupPath"`
	URLBase                string    `json:"urlBase"`
	Version                string    `json:"version"`
}

// GetSystemStatus returns system status.
func (s *Sonarr) GetSystemStatus() (*SystemStatus, error) {
	return s.GetSystemStatusContext(context.Background())
}

// GetSystemStatusContext returns system status.
func (s *Sonarr) GetSystemStatusContext(ctx context.Context) (*SystemStatus, error) {
	var output SystemStatus

	req := starr.Request{URI: path.Join(bpSystem, "status")}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// GetBackupFiles returns all available Sonarr backup files.
// Use Get to download a file using BackupFile.Path.
func (s *Sonarr) GetBackupFiles() ([]*starr.BackupFile, error) {
	return s.GetBackupFilesContext(context.Background())
}

// GetBackupFilesContext returns all available Sonarr backup files.
// Use Get() to download a file using BackupFile.Path.
func (s *Sonarr) GetBackupFilesContext(ctx context.Context) ([]*starr.BackupFile, error) {
	var output []*starr.BackupFile

	req := starr.Request{URI: path.Join(bpSystem, "backup")}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// SystemTask is a scheduled task from /api/v3/system/task.
type SystemTask = starrshared.SystemTask

// BackupRestoreResponse is returned when restoring a backup.
type BackupRestoreResponse = starrshared.BackupRestoreResponse

// DeleteBackup deletes a backup file by ID.
func (s *Sonarr) DeleteBackup(id int64) error {
	return s.DeleteBackupContext(context.Background(), id)
}

// DeleteBackupContext deletes a backup file by ID.
func (s *Sonarr) DeleteBackupContext(ctx context.Context, id int64) error {
	req := starr.Request{URI: path.Join(bpSystem, "backup", starr.Str(id))}
	if err := s.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}

// RestoreBackup restores an on-disk backup by ID.
func (s *Sonarr) RestoreBackup(id int64) (*BackupRestoreResponse, error) {
	return s.RestoreBackupContext(context.Background(), id)
}

// RestoreBackupContext restores an on-disk backup by ID.
func (s *Sonarr) RestoreBackupContext(ctx context.Context, id int64) (*BackupRestoreResponse, error) {
	var output BackupRestoreResponse

	req := starr.Request{URI: path.Join(bpSystem, "backup", "restore", starr.Str(id))}
	if err := s.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// RestoreBackupUpload uploads a backup archive and restores it.
func (s *Sonarr) RestoreBackupUpload(filename string, file io.Reader) (*BackupRestoreResponse, error) {
	return s.RestoreBackupUploadContext(context.Background(), filename, file)
}

// RestoreBackupUploadContext uploads a backup archive and restores it.
func (s *Sonarr) RestoreBackupUploadContext(
	ctx context.Context, filename string, file io.Reader,
) (*BackupRestoreResponse, error) {
	var buf bytes.Buffer

	writer := multipart.NewWriter(&buf)

	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return nil, fmt.Errorf("creating multipart form: %w", err)
	}

	if _, err = io.Copy(part, file); err != nil {
		return nil, fmt.Errorf("writing backup to multipart form: %w", err)
	}

	if err = writer.Close(); err != nil {
		return nil, fmt.Errorf("closing multipart writer: %w", err)
	}

	var output BackupRestoreResponse

	hdr := make(http.Header)
	hdr.Set("Content-Type", writer.FormDataContentType())

	req := starr.Request{
		URI:     path.Join(bpSystem, "backup", "restore", "upload"),
		Body:    &buf,
		Headers: hdr,
	}
	if err := s.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// Restart tells Sonarr to restart.
func (s *Sonarr) Restart() error {
	return s.RestartContext(context.Background())
}

// RestartContext tells Sonarr to restart.
func (s *Sonarr) RestartContext(ctx context.Context) error {
	var output any

	req := starr.Request{URI: path.Join(bpSystem, "restart")}
	if err := s.PostInto(ctx, req, &output); err != nil {
		return fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return nil
}

// Shutdown tells Sonarr to shut down.
func (s *Sonarr) Shutdown() error {
	return s.ShutdownContext(context.Background())
}

// ShutdownContext tells Sonarr to shut down.
func (s *Sonarr) ShutdownContext(ctx context.Context) error {
	var output any

	req := starr.Request{URI: path.Join(bpSystem, "shutdown")}
	if err := s.PostInto(ctx, req, &output); err != nil {
		return fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return nil
}

// GetSystemRoutes returns the raw JSON route table (schema-less in OpenAPI).
func (s *Sonarr) GetSystemRoutes() ([]byte, error) {
	return s.GetSystemRoutesContext(context.Background())
}

// GetSystemRoutesContext returns the raw JSON route table (schema-less in OpenAPI).
func (s *Sonarr) GetSystemRoutesContext(ctx context.Context) ([]byte, error) {
	uri := starr.SetAPIPath(path.Join(bpSystem, "routes"))

	req := starr.Request{URI: uri}

	resp, err := s.Get(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body from %s: %w", uri, err)
	}

	return body, nil
}

// GetSystemDuplicateRoutes returns duplicate route definitions as raw JSON.
func (s *Sonarr) GetSystemDuplicateRoutes() ([]byte, error) {
	return s.GetSystemDuplicateRoutesContext(context.Background())
}

// GetSystemDuplicateRoutesContext returns duplicate route definitions as raw JSON.
func (s *Sonarr) GetSystemDuplicateRoutesContext(ctx context.Context) ([]byte, error) {
	uri := starr.SetAPIPath(path.Join(bpSystem, "routes", "duplicate"))

	req := starr.Request{URI: uri}

	resp, err := s.Get(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body from %s: %w", uri, err)
	}

	return body, nil
}

// GetSystemTasks returns all scheduled tasks.
func (s *Sonarr) GetSystemTasks() ([]*SystemTask, error) {
	return s.GetSystemTasksContext(context.Background())
}

// GetSystemTasksContext returns all scheduled tasks.
func (s *Sonarr) GetSystemTasksContext(ctx context.Context) ([]*SystemTask, error) {
	var output []*SystemTask

	req := starr.Request{URI: path.Join(bpSystem, "task")}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetSystemTask returns a single scheduled task.
func (s *Sonarr) GetSystemTask(id int64) (*SystemTask, error) {
	return s.GetSystemTaskContext(context.Background(), id)
}

// GetSystemTaskContext returns a single scheduled task.
func (s *Sonarr) GetSystemTaskContext(ctx context.Context, id int64) (*SystemTask, error) {
	var output SystemTask

	req := starr.Request{URI: path.Join(bpSystem, "task", starr.Str(id))}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}
