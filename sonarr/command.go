package sonarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"golift.io/starr"
)

// CommandRequest goes into the /api/v3/command endpoint.
// This was created from the search command and may not support other commands yet.
type CommandRequest struct {
	Name      string  `json:"name"`
	Files     []int64 `json:"files,omitempty"` // RenameFiles only
	SeriesIDs []int64 `json:"seriesIds,omitempty"`
	SeriesID  int64   `json:"seriesId,omitempty"`
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

func (s *Sonarr) GetCommandsContext(ctx context.Context) ([]*CommandResponse, error) {
	var output []*CommandResponse

	if _, err := s.GetInto(ctx, "v3/command", nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(command): %w", err)
	}

	return output, nil
}

// SendCommand sends a command to Sonarr.
func (s *Sonarr) SendCommand(cmd *CommandRequest) (*CommandResponse, error) {
	return s.SendCommandContext(context.Background(), cmd)
}

func (s *Sonarr) SendCommandContext(ctx context.Context, cmd *CommandRequest) (*CommandResponse, error) {
	var output CommandResponse

	if cmd == nil || cmd.Name == "" {
		return &output, nil
	}

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(cmd); err != nil {
		return nil, fmt.Errorf("json.Marshal(cmd): %w", err)
	}

	if _, err := s.PostInto(ctx, "v3/command", nil, &body, &output); err != nil {
		return nil, fmt.Errorf("api.Post(command): %w", err)
	}

	return &output, nil
}

// GetCommandStatus returns the status of an already started command.
func (s *Sonarr) GetCommandStatus(commandID int64) (*CommandResponse, error) {
	return s.GetCommandStatusContext(context.Background(), commandID)
}

func (s *Sonarr) GetCommandStatusContext(ctx context.Context, commandID int64) (*CommandResponse, error) {
	var output CommandResponse

	if commandID == 0 {
		return &output, nil
	}

	_, err := s.GetInto(ctx, "v3/command/"+strconv.FormatInt(commandID, starr.Base10), nil, &output)
	if err != nil {
		return nil, fmt.Errorf("api.Post(command): %w", err)
	}

	return &output, nil
}
