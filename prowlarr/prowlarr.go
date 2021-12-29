package prowlarr

import (
	"fmt"

	"golift.io/starr"
)

// GetBackupFiles returns all available Prowlarr backup files.
// Use GetBody to download a file using BackupFile.Path.
func (r *Prowlarr) GetBackupFiles() ([]*starr.BackupFile, error) {
	var output []*starr.BackupFile

	if err := r.GetInto("v1/system/backup", nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(system/backup): %w", err)
	}

	return output, nil
}
