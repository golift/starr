package radarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
)

// CustomFormat is the api/customformat endpoint payload.
type CustomFormat struct {
	ID                    int                 `json:"id"`
	Name                  string              `json:"name"`
	IncludeCFWhenRenaming bool                `json:"includeCustomFormatWhenRenaming"`
	Specifications        []*CustomFormatSpec `json:"specifications"`
}

// CustomFormatSpec is part of a CustomFormat.
type CustomFormatSpec struct {
	Name               string               `json:"name"`
	Implementation     string               `json:"implementation"`
	Implementationname string               `json:"implementationName"`
	Infolink           string               `json:"infoLink"`
	Negate             bool                 `json:"negate"`
	Required           bool                 `json:"required"`
	Fields             []*CustomFormatField `json:"fields"`
}

// CustomFormatField is part of a CustomFormat Specification.
type CustomFormatField struct {
	Order    int         `json:"order"`
	Name     string      `json:"name"`
	Label    string      `json:"label"`
	Value    interface{} `json:"value"` // should be a string, but sometimes it's a number.
	Type     string      `json:"type"`
	Advanced bool        `json:"advanced"`
}

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
