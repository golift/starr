package sonarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"
	"time"

	"golift.io/starr"
)

const bpCommand = APIver + "/command"

// CommandRequest goes into the /api/v3/command endpoint.
// This was created from the search command and may not support other commands yet.
type CommandRequest struct {
	SeasonNumber int     `json:"seasonNumber,omitempty"`
	SeriesID     int64   `json:"seriesId,omitempty"`
	EpisodeID    int64   `json:"episodeId,omitempty"`
	Name         string  `json:"name"`
	Files        []int64 `json:"files,omitempty"` // RenameFiles only
	SeriesIDs    []int64 `json:"seriesIds,omitempty"`
	EpisodeIDs   []int64 `json:"episodeIds,omitempty"`
}

// CommandResponse comes from the /api/v3/command endpoint.
type CommandResponse struct {
	ID                  int64                  `json:"id"`
	Name                string                 `json:"name"`
	CommandName         string                 `json:"commandName"`
	Message             string                 `json:"message,omitempty"`
	Priority            string                 `json:"priority"`
	Status              string                 `json:"status"`
	Queued              time.Time              `json:"queued"`
	Started             time.Time              `json:"started,omitempty"`
	Ended               time.Time              `json:"ended,omitempty"`
	StateChangeTime     time.Time              `json:"stateChangeTime,omitempty"`
	LastExecutionTime   time.Time              `json:"lastExecutionTime,omitempty"`
	Duration            string                 `json:"duration,omitempty"`
	Trigger             string                 `json:"trigger"`
	SendUpdatesToClient bool                   `json:"sendUpdatesToClient"`
	UpdateScheduledTask bool                   `json:"updateScheduledTask"`
	Body                map[string]interface{} `json:"body"`
}

// GetCommands returns all available Sonarr commands.
// These can be used with SendCommand.
func (s *Sonarr) GetCommands() ([]*CommandResponse, error) {
	return s.GetCommandsContext(context.Background())
}

// GetCommands returns all available Sonarr commands.
// These can be used with SendCommand.
func (s *Sonarr) GetCommandsContext(ctx context.Context) ([]*CommandResponse, error) {
	var output []*CommandResponse

	req := starr.Request{URI: bpCommand}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// SendCommand sends a command to Sonarr.
func (s *Sonarr) SendCommand(cmd *CommandRequest) (*CommandResponse, error) {
	return s.SendCommandContext(context.Background(), cmd)
}

// SendCommandContext sends a command to Sonarr.
func (s *Sonarr) SendCommandContext(ctx context.Context, cmd *CommandRequest) (*CommandResponse, error) {
	var output CommandResponse

	if cmd == nil || cmd.Name == "" {
		return &output, nil
	}

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(cmd); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpCommand, err)
	}

	req := starr.Request{URI: bpCommand, Body: &body}
	if err := s.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// GetCommandStatus returns the status of an already started command.
func (s *Sonarr) GetCommandStatus(commandID int64) (*CommandResponse, error) {
	return s.GetCommandStatusContext(context.Background(), commandID)
}

// GetCommandStatusContext returns the status of an already started command.
func (s *Sonarr) GetCommandStatusContext(ctx context.Context, commandID int64) (*CommandResponse, error) {
	var output CommandResponse

	if commandID == 0 {
		return &output, nil
	}

	req := starr.Request{URI: path.Join(bpCommand, starr.Itoa(commandID))}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}
