package sonarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"golift.io/starr"
)

// GetCommands returns all available Sonarr commands.
// These can be used with SendCommand.
func (s *Sonarr) GetCommands() ([]*CommandResponse, error) {
	return s.GetCommandsContext(context.Background())
}

func (s *Sonarr) GetCommandsContext(ctx context.Context) ([]*CommandResponse, error) {
	var output []*CommandResponse

	if err := s.GetInto(ctx, "v3/command", nil, &output); err != nil {
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

	if err := s.PostInto(ctx, "v3/command", nil, &body, &output); err != nil {
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

	err := s.GetInto(ctx, "v3/command/"+strconv.FormatInt(commandID, starr.Base10), nil, &output)
	if err != nil {
		return nil, fmt.Errorf("api.Post(command): %w", err)
	}

	return &output, nil
}
