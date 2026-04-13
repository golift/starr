package starrshared

import "time"

// SystemTask is a scheduled task from /system/task.
type SystemTask struct {
	ID            int       `json:"id,omitempty"`
	Name          string    `json:"name,omitempty"`
	TaskName      string    `json:"taskName,omitempty"`
	Interval      int       `json:"interval,omitempty"`
	LastExecution time.Time `json:"lastExecution,omitzero"`
	LastStartTime time.Time `json:"lastStartTime,omitzero"`
	NextExecution time.Time `json:"nextExecution,omitzero"`
	LastDuration  string    `json:"lastDuration,omitempty"`
}

// BackupRestoreResponse is returned when restoring a backup.
type BackupRestoreResponse struct {
	RestartRequired bool `json:"restartRequired"`
}
