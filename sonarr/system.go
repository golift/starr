package sonarr

import (
	"context"
	"fmt"

	"golift.io/starr"
)

// GetSystemStatus returns system status.
func (s *Sonarr) GetSystemStatus() (*SystemStatus, error) {
	return s.GetSystemStatusContext(context.Background())
}

func (s *Sonarr) GetSystemStatusContext(ctx context.Context) (*SystemStatus, error) {
	var status SystemStatus

	err := s.GetInto(ctx, "v3/system/status", nil, &status)
	if err != nil {
		return nil, fmt.Errorf("api.Get(system/status): %w", err)
	}

	return &status, nil
}

// GetBackupFiles returns all available Sonarr backup files.
// Use GetBody to download a file using BackupFile.Path.
func (s *Sonarr) GetBackupFiles() ([]*starr.BackupFile, error) {
	return s.GetBackupFilesContext(context.Background())
}

func (s *Sonarr) GetBackupFilesContext(ctx context.Context) ([]*starr.BackupFile, error) {
	var output []*starr.BackupFile

	if err := s.GetInto(ctx, "v3/system/backup", nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(system/backup): %w", err)
	}

	return output, nil
}
