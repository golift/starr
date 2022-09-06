package radarr

import (
	"context"
	"fmt"
	"path"
	"time"

	"golift.io/starr"
)

const bpSystem = APIver + "/system"

// SystemStatus is the /api/v3/system/status endpoint.
type SystemStatus struct {
	Version           string    `json:"version"`
	BuildTime         time.Time `json:"buildTime"`
	StartupPath       string    `json:"startupPath"`
	AppData           string    `json:"appData"`
	OsName            string    `json:"osName"`
	OsVersion         string    `json:"osVersion"`
	Branch            string    `json:"branch"`
	Authentication    string    `json:"authentication"`
	SqliteVersion     string    `json:"sqliteVersion"`
	URLBase           string    `json:"urlBase"`
	RuntimeVersion    string    `json:"runtimeVersion"`
	RuntimeName       string    `json:"runtimeName"`
	MigrationVersion  int       `json:"migrationVersion"`
	IsDebug           bool      `json:"isDebug"`
	IsProduction      bool      `json:"isProduction"`
	IsAdmin           bool      `json:"isAdmin"`
	IsUserInteractive bool      `json:"isUserInteractive"`
	IsNetCore         bool      `json:"isNetCore"`
	IsMono            bool      `json:"isMono"`
	IsLinux           bool      `json:"isLinux"`
	IsOsx             bool      `json:"isOsx"`
	IsWindows         bool      `json:"isWindows"`
}

// GetSystemStatus returns system status.
func (r *Radarr) GetSystemStatus() (*SystemStatus, error) {
	return r.GetSystemStatusContext(context.Background())
}

// GetSystemStatusContext returns system status.
func (r *Radarr) GetSystemStatusContext(ctx context.Context) (*SystemStatus, error) {
	var output SystemStatus

	req := starr.Request{URI: path.Join(bpSystem, "status")}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", req, err)
	}

	return &output, nil
}

// GetBackupFiles returns all available Radarr backup files.
// Use GetBody to download a file using BackupFile.Path.
func (r *Radarr) GetBackupFiles() ([]*starr.BackupFile, error) {
	return r.GetBackupFilesContext(context.Background())
}

// GetBackupFilesContext returns all available Radarr backup files.
// Use GetBody to download a file using BackupFile.Path.
func (r *Radarr) GetBackupFilesContext(ctx context.Context) ([]*starr.BackupFile, error) {
	var output []*starr.BackupFile

	req := starr.Request{URI: path.Join(bpSystem, "backup")}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", req, err)
	}

	return output, nil
}
