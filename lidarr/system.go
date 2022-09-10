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
	Version           string    `json:"version"`
	BuildTime         time.Time `json:"buildTime"`
	IsDebug           bool      `json:"isDebug"`
	IsProduction      bool      `json:"isProduction"`
	IsAdmin           bool      `json:"isAdmin"`
	IsUserInteractive bool      `json:"isUserInteractive"`
	StartupPath       string    `json:"startupPath"`
	AppData           string    `json:"appData"`
	OsName            string    `json:"osName"`
	OsVersion         string    `json:"osVersion"`
	IsMonoRuntime     bool      `json:"isMonoRuntime"`
	IsMono            bool      `json:"isMono"`
	IsLinux           bool      `json:"isLinux"`
	IsOsx             bool      `json:"isOsx"`
	IsWindows         bool      `json:"isWindows"`
	IsDocker          bool      `json:"isDocker"`
	Mode              string    `json:"mode"`
	Branch            string    `json:"branch"`
	Authentication    string    `json:"authentication"`
	SqliteVersion     string    `json:"sqliteVersion"`
	MigrationVersion  int64     `json:"migrationVersion"`
	URLBase           string    `json:"urlBase"`
	RuntimeVersion    string    `json:"runtimeVersion"`
	RuntimeName       string    `json:"runtimeName"`
	StartTime         time.Time `json:"startTime"`
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
