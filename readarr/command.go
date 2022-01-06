package readarr

import (
	"encoding/json"
	"fmt"
)

// GetCommands returns all available Readarr commands.
// These can be used with SendCommand.
func (r *Readarr) GetCommands() ([]*CommandResponse, error) {
	var output []*CommandResponse

	if err := r.GetInto("v1/command", nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(command): %w", err)
	}

	return output, nil
}

// SendCommand sends a command to Readarr.
func (r *Readarr) SendCommand(cmd *CommandRequest) (*CommandResponse, error) {
	var output CommandResponse

	if cmd == nil || cmd.Name == "" {
		return &output, nil
	}

	body, err := json.Marshal(cmd)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal(cmd): %w", err)
	}

	if err := r.PostInto("v1/command", nil, body, &output); err != nil {
		return nil, fmt.Errorf("api.Post(command): %w", err)
	}

	return &output, nil
}
