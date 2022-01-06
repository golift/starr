package readarr

import (
	"fmt"

	"golift.io/starr"
)

// GetSystemStatus returns system status.
func (r *Readarr) GetSystemStatus() (*SystemStatus, error) {
	var status SystemStatus

	err := r.GetInto("v1/system/status", nil, &status)
	if err != nil {
		return &status, fmt.Errorf("api.Get(system/status): %w", err)
	}

	return &status, nil
}

// GetBackupFiles returns all available Readarr backup files.
// Use GetBody to download a file using BackupFile.Path.
func (r *Readarr) GetBackupFiles() ([]*starr.BackupFile, error) {
	var output []*starr.BackupFile

	if err := r.GetInto("v1/system/backup", nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(system/backup): %w", err)
	}

	return output, nil
}
