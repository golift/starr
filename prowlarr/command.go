package prowlarr

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

// CommandRequest is sent to POST /api/v1/command.
type CommandRequest struct {
	Name string         `json:"name"`
	Body map[string]any `json:"body,omitempty"`
}

// CommandResponse is returned from command endpoints.
type CommandResponse struct {
	ID                  int64          `json:"id"`
	Name                string         `json:"name"`
	CommandName         string         `json:"commandName"`
	Message             string         `json:"message,omitempty"`
	Priority            string         `json:"priority"`
	Status              string         `json:"status"`
	Queued              time.Time      `json:"queued"`
	Started             time.Time      `json:"started,omitzero"`
	Ended               time.Time      `json:"ended,omitzero"`
	StateChangeTime     time.Time      `json:"stateChangeTime,omitzero"`
	LastExecutionTime   time.Time      `json:"lastExecutionTime,omitzero"`
	Duration            string         `json:"duration,omitempty"`
	Trigger             string         `json:"trigger"`
	SendUpdatesToClient bool           `json:"sendUpdatesToClient"`
	UpdateScheduledTask bool           `json:"updateScheduledTask"`
	Body                map[string]any `json:"body"`
}

// GetCommands returns queued and recent commands.
func (p *Prowlarr) GetCommands() ([]*CommandResponse, error) {
	return p.GetCommandsContext(context.Background())
}

// GetCommandsContext returns queued and recent commands.
func (p *Prowlarr) GetCommandsContext(ctx context.Context) ([]*CommandResponse, error) {
	var output []*CommandResponse

	req := starr.Request{URI: bpCommand}
	if err := p.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// SendCommand queues a command.
func (p *Prowlarr) SendCommand(cmd *CommandRequest) (*CommandResponse, error) {
	return p.SendCommandContext(context.Background(), cmd)
}

// SendCommandContext queues a command.
func (p *Prowlarr) SendCommandContext(ctx context.Context, cmd *CommandRequest) (*CommandResponse, error) {
	var output CommandResponse

	if cmd == nil || cmd.Name == "" {
		return &output, nil
	}

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(cmd); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpCommand, err)
	}

	req := starr.Request{URI: bpCommand, Body: &body}
	if err := p.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// GetCommandStatus returns a command by id.
func (p *Prowlarr) GetCommandStatus(commandID int64) (*CommandResponse, error) {
	return p.GetCommandStatusContext(context.Background(), commandID)
}

// GetCommandStatusContext returns a command by id.
func (p *Prowlarr) GetCommandStatusContext(ctx context.Context, commandID int64) (*CommandResponse, error) {
	var output CommandResponse

	if commandID == 0 {
		return &output, nil
	}

	req := starr.Request{URI: path.Join(bpCommand, starr.Str(commandID))}
	if err := p.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteCommand removes a queued command.
func (p *Prowlarr) DeleteCommand(commandID int64) error {
	return p.DeleteCommandContext(context.Background(), commandID)
}

// DeleteCommandContext removes a queued command.
func (p *Prowlarr) DeleteCommandContext(ctx context.Context, commandID int64) error {
	req := starr.Request{URI: path.Join(bpCommand, starr.Str(commandID))}
	if err := p.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
