package radarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
)

// GetCustomFormats returns all configured Custom Formats.
func (r *Radarr) GetCustomFormats() ([]*CustomFormat, error) {
	return r.GetCustomFormatsContext(context.Background())
}

// GetCustomFormatsContext returns all configured Custom Formats.
func (r *Radarr) GetCustomFormatsContext(ctx context.Context) ([]*CustomFormat, error) {
	var output []*CustomFormat
	if _, err := r.GetInto(ctx, "v3/customFormat", nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(customFormat): %w", err)
	}

	return output, nil
}

// AddCustomFormat creates a new custom format and returns the response (with ID).
func (r *Radarr) AddCustomFormat(format *CustomFormat) (*CustomFormat, error) {
	return r.AddCustomFormatContext(context.Background(), format)
}

// AddCustomFormatContext creates a new custom format and returns the response (with ID).
func (r *Radarr) AddCustomFormatContext(ctx context.Context, format *CustomFormat) (*CustomFormat, error) {
	var output CustomFormat

	if format == nil {
		return &output, nil
	}

	format.ID = 0 // ID must be zero when adding.

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(format); err != nil {
		return nil, fmt.Errorf("json.Marshal(customFormat): %w", err)
	}

	if _, err := r.PostInto(ctx, "v3/customFormat", nil, &body, &output); err != nil {
		return nil, fmt.Errorf("api.Post(customFormat): %w", err)
	}

	return &output, nil
}

// UpdateCustomFormat updates an existing custom format and returns the response.
func (r *Radarr) UpdateCustomFormat(cf *CustomFormat, cfID int) (*CustomFormat, error) {
	return r.UpdateCustomFormatContext(context.Background(), cf, cfID)
}

// UpdateCustomFormatContext updates an existing custom format and returns the response.
func (r *Radarr) UpdateCustomFormatContext(ctx context.Context, format *CustomFormat, cfID int) (*CustomFormat, error) {
	if cfID == 0 {
		cfID = format.ID
	}

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(format); err != nil {
		return nil, fmt.Errorf("json.Marshal(customFormat): %w", err)
	}

	var output CustomFormat
	if _, err := r.PutInto(ctx, "v3/customFormat/"+strconv.Itoa(cfID), nil, &body, &output); err != nil {
		return nil, fmt.Errorf("api.Put(customFormat): %w", err)
	}

	return &output, nil
}
