package prowlarr

import (
	"fmt"

	"golift.io/starr"
)

// GetSystemStatus returns system status.
func (p *Prowlarr) GetSystemStatus() (*SystemStatus, error) {
	var status SystemStatus

	err := p.GetInto("v1/system/status", nil, &status)
	if err != nil {
		return nil, fmt.Errorf("api.Get(system/status): %w", err)
	}

	return &status, nil
}

// GetBackupFiles returns all available Prowlarr backup files.
// Use GetBody to download a file using BackupFile.Path.
func (p *Prowlarr) GetBackupFiles() ([]*starr.BackupFile, error) {
	var output []*starr.BackupFile

	if err := p.GetInto("v1/system/backup", nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(system/backup): %w", err)
	}

	return output, nil
}
