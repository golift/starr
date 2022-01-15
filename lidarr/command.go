package lidarr

import (
	"context"
	"encoding/json"
	"fmt"
)

// GetCommands returns all available Lidarr commands.
func (l *Lidarr) GetCommands() ([]*CommandResponse, error) {
	return l.GetCommandsContext(context.Background())
}

// GetCommandsContext returns all available Lidarr commands.
func (l *Lidarr) GetCommandsContext(ctx context.Context) ([]*CommandResponse, error) {
	var output []*CommandResponse

	if err := l.GetInto(ctx, "v1/command", nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(command): %w", err)
	}

	return output, nil
}

// SendCommand sends a command to Lidarr.
func (l *Lidarr) SendCommand(cmd *CommandRequest) (*CommandResponse, error) {
	return l.SendCommandContext(context.Background(), cmd)
}

// SendCommandContext sends a command to Lidarr.
func (l *Lidarr) SendCommandContext(ctx context.Context, cmd *CommandRequest) (*CommandResponse, error) {
	var output CommandResponse

	if cmd == nil || cmd.Name == "" {
		return &output, nil
	}

	body, err := json.Marshal(cmd)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal(cmd): %w", err)
	}

	if err := l.PostInto(ctx, "v1/command", nil, body, &output); err != nil {
		return nil, fmt.Errorf("api.Post(command): %w", err)
	}

	return &output, nil
}
