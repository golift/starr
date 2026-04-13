package readarr

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

// SystemStatus is the /api/v1/system/status endpoint.
type SystemStatus struct {
	AppData                string    `json:"appData"`
	AppName                string    `json:"appName"`
	Authentication         string    `json:"authentication"`
	Branch                 string    `json:"branch"`
	BuildTime              time.Time `json:"buildTime"`
	DatabaseType           string    `json:"databaseType"`
	DatabaseVersion        string    `json:"databaseVersion"`
	InstanceName           string    `json:"instanceName"`
	IsAdmin                bool      `json:"isAdmin"`
	IsDebug                bool      `json:"isDebug"`
	IsDocker               bool      `json:"isDocker"`
	IsLinux                bool      `json:"isLinux"`
	IsMono                 bool      `json:"isMono"`
	IsNetCore              bool      `json:"isNetCore"`
	IsOsx                  bool      `json:"isOsx"`
	IsProduction           bool      `json:"isProduction"`
	IsUserInteractive      bool      `json:"isUserInteractive"`
	IsWindows              bool      `json:"isWindows"`
	MigrationVersion       int64     `json:"migrationVersion"`
	Mode                   string    `json:"mode"`
	OsName                 string    `json:"osName"`
	OsVersion              string    `json:"osVersion"`
	PackageAuthor          string    `json:"packageAuthor"`
	PackageUpdateMechanism string    `json:"packageUpdateMechanism"`
	PackageVersion         string    `json:"packageVersion"`
	RuntimeName            string    `json:"runtimeName"`
	RuntimeVersion         string    `json:"runtimeVersion"`
	StartTime              time.Time `json:"startTime"`
	StartupPath            string    `json:"startupPath"`
	URLBase                string    `json:"urlBase"`
	Version                string    `json:"version"`
}

// GetSystemStatus returns system status.
func (r *Readarr) GetSystemStatus() (*SystemStatus, error) {
	return r.GetSystemStatusContext(context.Background())
}

// GetSystemStatusContext returns system status.
func (r *Readarr) GetSystemStatusContext(ctx context.Context) (*SystemStatus, error) {
	var output SystemStatus

	req := starr.Request{URI: path.Join(bpSystem, "status")}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// GetBackupFiles returns all available Readarr backup files.
// Use Get() to download a file using BackupFile.Path.
func (r *Readarr) GetBackupFiles() ([]*starr.BackupFile, error) {
	return r.GetBackupFilesContext(context.Background())
}

// GetBackupFilesContext returns all available Readarr backup files.
// Use Get() to download a file using BackupFile.Path.
func (r *Readarr) GetBackupFilesContext(ctx context.Context) ([]*starr.BackupFile, error) {
	var output []*starr.BackupFile

	req := starr.Request{URI: path.Join(bpSystem, "backup")}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// SystemTask is a scheduled task from /api/v1/system/task.
type SystemTask = starrshared.SystemTask

// BackupRestoreResponse is returned when restoring a backup.
type BackupRestoreResponse = starrshared.BackupRestoreResponse

// DeleteBackup deletes a backup file by ID.
func (r *Readarr) DeleteBackup(id int64) error {
	return r.DeleteBackupContext(context.Background(), id)
}

// DeleteBackupContext deletes a backup file by ID.
func (r *Readarr) DeleteBackupContext(ctx context.Context, id int64) error {
	req := starr.Request{URI: path.Join(bpSystem, "backup", starr.Str(id))}
	if err := r.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}

// RestoreBackup restores an on-disk backup by ID.
func (r *Readarr) RestoreBackup(id int64) (*BackupRestoreResponse, error) {
	return r.RestoreBackupContext(context.Background(), id)
}

// RestoreBackupContext restores an on-disk backup by ID.
func (r *Readarr) RestoreBackupContext(ctx context.Context, id int64) (*BackupRestoreResponse, error) {
	var output BackupRestoreResponse

	req := starr.Request{URI: path.Join(bpSystem, "backup", "restore", starr.Str(id))}
	if err := r.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// RestoreBackupUpload uploads a backup archive and restores it.
func (r *Readarr) RestoreBackupUpload(filename string, file io.Reader) (*BackupRestoreResponse, error) {
	return r.RestoreBackupUploadContext(context.Background(), filename, file)
}

// RestoreBackupUploadContext uploads a backup archive and restores it.
func (r *Readarr) RestoreBackupUploadContext(
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
	if err := r.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// Restart tells Readarr to restart.
func (r *Readarr) Restart() error {
	return r.RestartContext(context.Background())
}

// RestartContext tells Readarr to restart.
func (r *Readarr) RestartContext(ctx context.Context) error {
	var output any

	req := starr.Request{URI: path.Join(bpSystem, "restart")}
	if err := r.PostInto(ctx, req, &output); err != nil {
		return fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return nil
}

// Shutdown tells Readarr to shut down.
func (r *Readarr) Shutdown() error {
	return r.ShutdownContext(context.Background())
}

// ShutdownContext tells Readarr to shut down.
func (r *Readarr) ShutdownContext(ctx context.Context) error {
	var output any

	req := starr.Request{URI: path.Join(bpSystem, "shutdown")}
	if err := r.PostInto(ctx, req, &output); err != nil {
		return fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return nil
}

// GetSystemTasks returns all scheduled tasks.
func (r *Readarr) GetSystemTasks() ([]*SystemTask, error) {
	return r.GetSystemTasksContext(context.Background())
}

// GetSystemTasksContext returns all scheduled tasks.
func (r *Readarr) GetSystemTasksContext(ctx context.Context) ([]*SystemTask, error) {
	var output []*SystemTask

	req := starr.Request{URI: path.Join(bpSystem, "task")}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetSystemTask returns a single scheduled task.
func (r *Readarr) GetSystemTask(id int64) (*SystemTask, error) {
	return r.GetSystemTaskContext(context.Background(), id)
}

// GetSystemTaskContext returns a single scheduled task.
func (r *Readarr) GetSystemTaskContext(ctx context.Context, id int64) (*SystemTask, error) {
	var output SystemTask

	req := starr.Request{URI: path.Join(bpSystem, "task", starr.Str(id))}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}
