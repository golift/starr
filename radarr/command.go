package radarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// CommandRequest goes into the /api/v3/command endpoint.
// This was created from the search command and may not support other commands yet.
type CommandRequest struct {
	Name     string  `json:"name"`
	MovieIDs []int64 `json:"movieIds,omitempty"`
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

// GetCommands returns all available Radarr commands.
func (r *Radarr) GetCommands() ([]*CommandResponse, error) {
	return r.GetCommandsContext(context.Background())
}

// GetCommandsContext returns all available Radarr commands.
func (r *Radarr) GetCommandsContext(ctx context.Context) ([]*CommandResponse, error) {
	var output []*CommandResponse

	if err := r.GetInto(ctx, "v3/command", nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(command): %w", err)
	}

	return output, nil
}

// SendCommand sends a command to Radarr.
func (r *Radarr) SendCommand(cmd *CommandRequest) (*CommandResponse, error) {
	return r.SendCommandContext(context.Background(), cmd)
}

// SendCommandContext sends a command to Radarr.
func (r *Radarr) SendCommandContext(ctx context.Context, cmd *CommandRequest) (*CommandResponse, error) {
	var output CommandResponse

	if cmd == nil || cmd.Name == "" {
		return &output, nil
	}

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(cmd); err != nil {
		return nil, fmt.Errorf("json.Marshal(cmd): %w", err)
	}

	if err := r.PostInto(ctx, "v3/command", nil, &body, &output); err != nil {
		return nil, fmt.Errorf("api.Post(command): %w", err)
	}

	return &output, nil
}
