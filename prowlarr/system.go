package prowlarr

import (
	"context"
	"fmt"

	"golift.io/starr"
)

// GetSystemStatus returns system status.
func (p *Prowlarr) GetSystemStatus() (*SystemStatus, error) {
	return p.GetSystemStatusContext(context.Background())
}

// GetSystemStatusContext returns system status.
func (p *Prowlarr) GetSystemStatusContext(ctx context.Context) (*SystemStatus, error) {
	var status SystemStatus

	_, err := p.GetInto(ctx, "v1/system/status", nil, &status)
	if err != nil {
		return nil, fmt.Errorf("api.Get(system/status): %w", err)
	}

	return &status, nil
}

// GetBackupFiles returns all available Prowlarr backup files.
// Use GetBody to download a file using BackupFile.Path.
func (p *Prowlarr) GetBackupFiles() ([]*starr.BackupFile, error) {
	return p.GetBackupFilesContext(context.Background())
}

// GetBackupFiles returns all available Prowlarr backup files.
// Use GetBody to download a file using BackupFile.Path.
func (p *Prowlarr) GetBackupFilesContext(ctx context.Context) ([]*starr.BackupFile, error) {
	var output []*starr.BackupFile

	if _, err := p.GetInto(ctx, "v1/system/backup", nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(system/backup): %w", err)
	}

	return output, nil
}
