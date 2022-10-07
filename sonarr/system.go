package sonarr

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
