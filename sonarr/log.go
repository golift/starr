package sonarr

import (
	"context"
	"fmt"
	"path"
	"time"

	"golift.io/starr"
)

const bpLog = APIver + "/log"

// LogLine is one record from /api/v3/log.
type LogLine struct {
	ID            int       `json:"id"`
	Time          time.Time `json:"time"`
	Exception     string    `json:"exception,omitempty"`
	ExceptionType string    `json:"exceptionType,omitempty"`
	Level         string    `json:"level,omitempty"`
	Logger        string    `json:"logger,omitempty"`
	Message       string    `json:"message,omitempty"`
	Method        string    `json:"method,omitempty"`
}

// LogPage is a page of log lines from /api/v3/log.
type LogPage struct {
	Page          int        `json:"page"`
	PageSize      int        `json:"pageSize"`
	SortKey       string     `json:"sortKey"`
	SortDirection string     `json:"sortDirection"`
	TotalRecords  int        `json:"totalRecords"`
	Records       []*LogLine `json:"records"`
}

// LogFile describes a log file on disk.
type LogFile struct {
	Filename  string    `json:"filename,omitempty"`
	Contents  string    `json:"contents,omitempty"`
	LastWrite time.Time `json:"lastWrite,omitzero"`
}

// GetLogPage returns a page of application log lines.
func (s *Sonarr) GetLogPage(params *starr.PageReq) (*LogPage, error) {
	return s.GetLogPageContext(context.Background(), params)
}

// GetLogPageContext returns a page of application log lines.
func (s *Sonarr) GetLogPageContext(ctx context.Context, params *starr.PageReq) (*LogPage, error) {
	var output LogPage

	req := starr.Request{URI: bpLog, Query: params.Params()}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// GetLogFiles returns the list of log files.
func (s *Sonarr) GetLogFiles() ([]*LogFile, error) {
	return s.GetLogFilesContext(context.Background())
}

// GetLogFilesContext returns the list of log files.
func (s *Sonarr) GetLogFilesContext(ctx context.Context) ([]*LogFile, error) {
	var output []*LogFile

	req := starr.Request{URI: path.Join(bpLog, "file")}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetLogFile returns the contents of a named log file.
func (s *Sonarr) GetLogFile(filename string) (*LogFile, error) {
	return s.GetLogFileContext(context.Background(), filename)
}

// GetLogFileContext returns the contents of a named log file.
func (s *Sonarr) GetLogFileContext(ctx context.Context, filename string) (*LogFile, error) {
	var output LogFile

	req := starr.Request{URI: path.Join(bpLog, "file", filename)}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateLogFiles triggers a log file update/roll.
func (s *Sonarr) UpdateLogFiles() ([]*LogFile, error) {
	return s.UpdateLogFilesContext(context.Background())
}

// UpdateLogFilesContext triggers a log file update/roll.
func (s *Sonarr) UpdateLogFilesContext(ctx context.Context) ([]*LogFile, error) {
	var output []*LogFile

	req := starr.Request{URI: path.Join(bpLog, "file", "update")}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// UpdateLogFile triggers update for a specific log file.
func (s *Sonarr) UpdateLogFile(filename string) ([]*LogFile, error) {
	return s.UpdateLogFileContext(context.Background(), filename)
}

// UpdateLogFileContext triggers update for a specific log file.
func (s *Sonarr) UpdateLogFileContext(ctx context.Context, filename string) ([]*LogFile, error) {
	var output []*LogFile

	req := starr.Request{URI: path.Join(bpLog, "file", "update", filename)}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}
