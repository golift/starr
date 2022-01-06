package radarr

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// GetCustomFormats returns all configured Custom Formats.
func (r *Radarr) GetCustomFormats() ([]*CustomFormat, error) {
	var output []*CustomFormat
	if err := r.GetInto("v3/customFormat", nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(customFormat): %w", err)
	}

	return output, nil
}

// AddCustomFormat creates a new custom format and returns the response (with ID).
func (r *Radarr) AddCustomFormat(format *CustomFormat) (*CustomFormat, error) {
	var output CustomFormat

	if format == nil {
		return &output, nil
	}

	format.ID = 0 // ID must be zero when adding.

	body, err := json.Marshal(format)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal(customFormat): %w", err)
	}

	if err := r.PostInto("v3/customFormat", nil, body, &output); err != nil {
		return nil, fmt.Errorf("api.Post(customFormat): %w", err)
	}

	return &output, nil
}

// UpdateCustomFormat updates an existing custom format and returns the response.
func (r *Radarr) UpdateCustomFormat(cf *CustomFormat, cfID int) (*CustomFormat, error) {
	if cfID == 0 {
		cfID = cf.ID
	}

	body, err := json.Marshal(cf)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal(customFormat): %w", err)
	}

	var output CustomFormat
	if err := r.PutInto("v3/customFormat/"+strconv.Itoa(cfID), nil, body, &output); err != nil {
		return nil, fmt.Errorf("api.Put(customFormat): %w", err)
	}

	return &output, nil
}
