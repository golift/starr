package radarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"

	"golift.io/starr"
)

const bpCustomFormat = APIver + "/customFormat"

// CustomFormatInput is the input for a new or updated CustomFormat.
type CustomFormatInput struct {
	ID                    int64                    `json:"id,omitempty"`
	Name                  string                   `json:"name"`
	IncludeCFWhenRenaming bool                     `json:"includeCustomFormatWhenRenaming"`
	Specifications        []*CustomFormatInputSpec `json:"specifications"`
}

// CustomFormatInputSpec is part of a CustomFormatInput.
type CustomFormatInputSpec struct {
	Name           string              `json:"name"`
	Implementation string              `json:"implementation"`
	Negate         bool                `json:"negate"`
	Required       bool                `json:"required"`
	Fields         []*starr.FieldInput `json:"fields"`
}

// CustomFormatOutput is the output from the CustomFormat methods.
type CustomFormatOutput struct {
	ID                    int64                     `json:"id"`
	Name                  string                    `json:"name"`
	IncludeCFWhenRenaming bool                      `json:"includeCustomFormatWhenRenaming"`
	Specifications        []*CustomFormatOutputSpec `json:"specifications"`
}

// CustomFormatOutputSpec is part of a CustomFormatOutput.
type CustomFormatOutputSpec struct {
	Name               string               `json:"name"`
	Implementation     string               `json:"implementation"`
	ImplementationName string               `json:"implementationName"`
	InfoLink           string               `json:"infoLink"`
	Negate             bool                 `json:"negate"`
	Required           bool                 `json:"required"`
	Fields             []*starr.FieldOutput `json:"fields"`
}

// GetCustomFormats returns all configured Custom Formats.
func (r *Radarr) GetCustomFormats() ([]*CustomFormatOutput, error) {
	return r.GetCustomFormatsContext(context.Background())
}

// GetCustomFormatsContext returns all configured Custom Formats.
func (r *Radarr) GetCustomFormatsContext(ctx context.Context) ([]*CustomFormatOutput, error) {
	var output []*CustomFormatOutput

	req := starr.Request{URI: bpCustomFormat}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetCustomFormat returns a single customformat.
func (r *Radarr) GetCustomFormat(customformatID int64) (*CustomFormatOutput, error) {
	return r.GetCustomFormatContext(context.Background(), customformatID)
}

// GetCustomFormatContext returns a single customformat.
func (r *Radarr) GetCustomFormatContext(ctx context.Context, customformatID int64) (*CustomFormatOutput, error) {
	var output CustomFormatOutput

	req := starr.Request{URI: path.Join(bpCustomFormat, starr.Str(customformatID))}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// AddCustomFormat creates a new custom format and returns the response (with ID).
func (r *Radarr) AddCustomFormat(format *CustomFormatInput) (*CustomFormatOutput, error) {
	return r.AddCustomFormatContext(context.Background(), format)
}

// AddCustomFormatContext creates a new custom format and returns the response (with ID).
func (r *Radarr) AddCustomFormatContext(ctx context.Context, format *CustomFormatInput) (*CustomFormatOutput, error) {
	var output CustomFormatOutput

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(format); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpCustomFormat, err)
	}

	req := starr.Request{URI: bpCustomFormat, Body: &body}
	if err := r.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateCustomFormat updates an existing custom format and returns the response.
func (r *Radarr) UpdateCustomFormat(cf *CustomFormatInput) (*CustomFormatOutput, error) {
	return r.UpdateCustomFormatContext(context.Background(), cf)
}

// UpdateCustomFormatContext updates an existing custom format and returns the response.
func (r *Radarr) UpdateCustomFormatContext(ctx context.Context,
	format *CustomFormatInput,
) (*CustomFormatOutput, error) {
	var output CustomFormatOutput

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(format); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpCustomFormat, err)
	}

	req := starr.Request{URI: path.Join(bpCustomFormat, starr.Str(format.ID)), Body: &body}
	if err := r.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteCustomFormat deletes a custom format.
func (r *Radarr) DeleteCustomFormat(cfID int64) error {
	return r.DeleteCustomFormatContext(context.Background(), cfID)
}

// DeleteCustomFormatContext deletes a custom format.
func (r *Radarr) DeleteCustomFormatContext(ctx context.Context, cfID int64) error {
	req := starr.Request{URI: path.Join(bpCustomFormat, starr.Str(cfID))}
	if err := r.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
