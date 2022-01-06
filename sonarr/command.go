package sonarr

import (
	"encoding/json"
	"fmt"
	"strconv"

	"golift.io/starr"
)

// GetCommands returns all available Sonarr commands.
// These can be used with SendCommand.
func (s *Sonarr) GetCommands() ([]*CommandResponse, error) {
	var output []*CommandResponse

	if err := s.GetInto("v3/command", nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(command): %w", err)
	}

	return output, nil
}

// SendCommand sends a command to Sonarr.
func (s *Sonarr) SendCommand(cmd *CommandRequest) (*CommandResponse, error) {
	var output CommandResponse

	if cmd == nil || cmd.Name == "" {
		return &output, nil
	}

	body, err := json.Marshal(cmd)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal(cmd): %w", err)
	}

	if err := s.PostInto("v3/command", nil, body, &output); err != nil {
		return nil, fmt.Errorf("api.Post(command): %w", err)
	}

	return &output, nil
}

// GetCommandStatus returns the status of an already started command.
func (s *Sonarr) GetCommandStatus(commandID int64) (*CommandResponse, error) {
	var output CommandResponse

	if commandID == 0 {
		return &output, nil
	}

	err := s.GetInto("v3/command/"+strconv.FormatInt(commandID, starr.Base10), nil, &output)
	if err != nil {
		return nil, fmt.Errorf("api.Post(command): %w", err)
	}

	return &output, nil
}
