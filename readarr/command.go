package readarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

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
