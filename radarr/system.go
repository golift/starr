package radarr

import (
	"context"
	"fmt"

	"golift.io/starr"
)

// GetSystemStatus returns system status.
func (r *Radarr) GetSystemStatus() (*SystemStatus, error) {
	return r.GetSystemStatusContext(context.Background())
}

// GetSystemStatusContext returns system status.
func (r *Radarr) GetSystemStatusContext(ctx context.Context) (*SystemStatus, error) {
	var status SystemStatus

	_, err := r.GetInto(ctx, "v3/system/status", nil, &status)
	if err != nil {
		return nil, fmt.Errorf("api.Get(system/status): %w", err)
	}

	return &status, nil
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

	if _, err := r.GetInto(ctx, "v3/system/backup", nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(system/backup): %w", err)
	}

	return output, nil
}
