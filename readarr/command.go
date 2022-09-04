package readarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// CommandRequest goes into the /api/v1/command endpoint.
// This was created from the search command and may not support other commands yet.
type CommandRequest struct {
	Name    string  `json:"name"`
	BookIDs []int64 `json:"bookIds,omitempty"`
	BookID  int64   `json:"bookId,omitempty"`
}

// CommandResponse comes from the /api/v1/command endpoint.
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

// GetCommands returns all available Readarr commands.
// These can be used with SendCommand.
func (r *Readarr) GetCommands() ([]*CommandResponse, error) {
	return r.GetCommandsContext(context.Background())
}

func (r *Readarr) GetCommandsContext(ctx context.Context) ([]*CommandResponse, error) {
	var output []*CommandResponse

	if err := r.GetInto(ctx, "v1/command", nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(command): %w", err)
	}

	return output, nil
}

// SendCommand sends a command to Readarr.
func (r *Readarr) SendCommand(cmd *CommandRequest) (*CommandResponse, error) {
	return r.SendCommandContext(context.Background(), cmd)
}

func (r *Readarr) SendCommandContext(ctx context.Context, cmd *CommandRequest) (*CommandResponse, error) {
	var output CommandResponse

	if cmd == nil || cmd.Name == "" {
		return &output, nil
	}

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(cmd); err != nil {
		return nil, fmt.Errorf("json.Marshal(cmd): %w", err)
	}

	if err := r.PostInto(ctx, "v1/command", nil, &body, &output); err != nil {
		return nil, fmt.Errorf("api.Post(command): %w", err)
	}

	return &output, nil
}
