package prowlarr

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
func (p *Prowlarr) GetSystemStatus() (*SystemStatus, error) {
	return p.GetSystemStatusContext(context.Background())
}

// GetSystemStatusContext returns system status.
func (p *Prowlarr) GetSystemStatusContext(ctx context.Context) (*SystemStatus, error) {
	var output SystemStatus

	req := starr.Request{URI: path.Join(bpSystem, "status")}
	if err := p.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// GetBackupFiles returns all available Prowlarr backup files.
// Use GetBody to download a file using BackupFile.Path.
func (p *Prowlarr) GetBackupFiles() ([]*starr.BackupFile, error) {
	return p.GetBackupFilesContext(context.Background())
}

// GetBackupFilesContext returns all available Prowlarr backup files.
// Use GetBody to download a file using BackupFile.Path.
func (p *Prowlarr) GetBackupFilesContext(ctx context.Context) ([]*starr.BackupFile, error) {
	var output []*starr.BackupFile

	req := starr.Request{URI: path.Join(bpSystem, "backup")}
	if err := p.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// SystemTask is a scheduled task from /api/v1/system/task.
type SystemTask = starrshared.SystemTask

// BackupRestoreResponse is returned when restoring a backup.
type BackupRestoreResponse = starrshared.BackupRestoreResponse

// DeleteBackup deletes a backup file by ID.
func (p *Prowlarr) DeleteBackup(id int64) error {
	return p.DeleteBackupContext(context.Background(), id)
}

// DeleteBackupContext deletes a backup file by ID.
func (p *Prowlarr) DeleteBackupContext(ctx context.Context, id int64) error {
	req := starr.Request{URI: path.Join(bpSystem, "backup", starr.Str(id))}
	if err := p.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}

// RestoreBackup restores an on-disk backup by ID.
func (p *Prowlarr) RestoreBackup(id int64) (*BackupRestoreResponse, error) {
	return p.RestoreBackupContext(context.Background(), id)
}

// RestoreBackupContext restores an on-disk backup by ID.
func (p *Prowlarr) RestoreBackupContext(ctx context.Context, id int64) (*BackupRestoreResponse, error) {
	var output BackupRestoreResponse

	req := starr.Request{URI: path.Join(bpSystem, "backup", "restore", starr.Str(id))}
	if err := p.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// RestoreBackupUpload uploads a backup archive and restores it.
func (p *Prowlarr) RestoreBackupUpload(filename string, file io.Reader) (*BackupRestoreResponse, error) {
	return p.RestoreBackupUploadContext(context.Background(), filename, file)
}

// RestoreBackupUploadContext uploads a backup archive and restores it.
func (p *Prowlarr) RestoreBackupUploadContext(
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
	if err := p.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// Restart tells Prowlarr to restart.
func (p *Prowlarr) Restart() error {
	return p.RestartContext(context.Background())
}

// RestartContext tells Prowlarr to restart.
func (p *Prowlarr) RestartContext(ctx context.Context) error {
	var output any

	req := starr.Request{URI: path.Join(bpSystem, "restart")}
	if err := p.PostInto(ctx, req, &output); err != nil {
		return fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return nil
}

// Shutdown tells Prowlarr to shut down.
func (p *Prowlarr) Shutdown() error {
	return p.ShutdownContext(context.Background())
}

// ShutdownContext tells Prowlarr to shut down.
func (p *Prowlarr) ShutdownContext(ctx context.Context) error {
	var output any

	req := starr.Request{URI: path.Join(bpSystem, "shutdown")}
	if err := p.PostInto(ctx, req, &output); err != nil {
		return fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return nil
}

// GetSystemTasks returns all scheduled tasks.
func (p *Prowlarr) GetSystemTasks() ([]*SystemTask, error) {
	return p.GetSystemTasksContext(context.Background())
}

// GetSystemTasksContext returns all scheduled tasks.
func (p *Prowlarr) GetSystemTasksContext(ctx context.Context) ([]*SystemTask, error) {
	var output []*SystemTask

	req := starr.Request{URI: path.Join(bpSystem, "task")}
	if err := p.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetSystemTask returns a single scheduled task.
func (p *Prowlarr) GetSystemTask(id int64) (*SystemTask, error) {
	return p.GetSystemTaskContext(context.Background(), id)
}

// GetSystemTaskContext returns a single scheduled task.
func (p *Prowlarr) GetSystemTaskContext(ctx context.Context, id int64) (*SystemTask, error) {
	var output SystemTask

	req := starr.Request{URI: path.Join(bpSystem, "task", starr.Str(id))}
	if err := p.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}
