package radarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

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
