package lidarr

import (
	"context"
	"fmt"
	"path"
	"time"

	"golift.io/starr"
)

const bpSystem = APIver + "/system"

// SystemStatus is the /api/v1/system/status endpoint.
type SystemStatus struct {
	AppData                string    `json:"appData"`
	AppName                string    `json:"appName"`
	Authentication         string    `json:"authentication"`
	Branch                 string    `json:"branch"`
	BuildTime              time.Time `json:"buildTime"`
	InstanceName           string    `json:"instanceName"`
	IsAdmin                bool      `json:"isAdmin"`
	IsDebug                bool      `json:"isDebug"`
	IsDocker               bool      `json:"isDocker"`
	IsLinux                bool      `json:"isLinux"`
	IsNetCore              bool      `json:"isNetCore"`
	IsOsx                  bool      `json:"isOsx"`
	IsProduction           bool      `json:"isProduction"`
	IsUserInteractive      bool      `json:"isUserInteractive"`
	IsWindows              bool      `json:"isWindows"`
	MigrationVersion       int64     `json:"migrationVersion"`
	Mode                   string    `json:"mode"`
	OsName                 string    `json:"osName"`
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
func (l *Lidarr) GetSystemStatus() (*SystemStatus, error) {
	return l.GetSystemStatusContext(context.Background())
}

// GetSystemStatusContext returns system status.
func (l *Lidarr) GetSystemStatusContext(ctx context.Context) (*SystemStatus, error) {
	var output SystemStatus

	req := starr.Request{URI: path.Join(bpSystem, "status")}
	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// GetBackupFiles returns all available Lidarr backup files.
// Use GetBody to download a file using BackupFile.Path.
func (l *Lidarr) GetBackupFiles() ([]*starr.BackupFile, error) {
	return l.GetBackupFilesContext(context.Background())
}

// GetBackupFilesContext returns all available Lidarr backup files.
// Use GetBody to download a file using BackupFile.Path.
func (l *Lidarr) GetBackupFilesContext(ctx context.Context) ([]*starr.BackupFile, error) {
	var output []*starr.BackupFile

	req := starr.Request{URI: path.Join(bpSystem, "backup")}
	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}
