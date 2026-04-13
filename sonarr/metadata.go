package sonarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"path"

	"golift.io/starr"
)

const bpMetadata = APIver + "/metadata"

// MetadataProviderMessage is the provider message object on metadata consumers.
type MetadataProviderMessage struct {
	Message string `json:"message,omitempty"`
	Type    string `json:"type,omitempty"`
}

// MetadataOutput is the output from /api/v3/metadata (MetadataResource).
type MetadataOutput struct {
	ID                 int64                    `json:"id,omitempty"`
	Name               string                   `json:"name,omitempty"`
	Fields             []*starr.FieldOutput     `json:"fields,omitempty"`
	ImplementationName string                   `json:"implementationName,omitempty"`
	Implementation     string                   `json:"implementation,omitempty"`
	ConfigContract     string                   `json:"configContract,omitempty"`
	InfoLink           string                   `json:"infoLink,omitempty"`
	Message            *MetadataProviderMessage `json:"message,omitempty"`
	Tags               []int                    `json:"tags,omitempty"`
	Presets            []*MetadataOutput        `json:"presets,omitempty"`
	Enable             bool                     `json:"enable"`
}

// MetadataInput is the input for creating or updating metadata consumers.
type MetadataInput struct {
	ID             int64               `json:"id,omitempty"`
	Name           string              `json:"name,omitempty"`
	Fields         []*starr.FieldInput `json:"fields,omitempty"`
	Implementation string              `json:"implementation,omitempty"`
	ConfigContract string              `json:"configContract,omitempty"`
	Tags           []int               `json:"tags,omitempty"`
	Enable         bool                `json:"enable"`
}

// GetMetadata returns all configured metadata consumers.
func (s *Sonarr) GetMetadata() ([]*MetadataOutput, error) {
	return s.GetMetadataContext(context.Background())
}

// GetMetadataContext returns all configured metadata consumers.
func (s *Sonarr) GetMetadataContext(ctx context.Context) ([]*MetadataOutput, error) {
	var output []*MetadataOutput

	req := starr.Request{URI: bpMetadata}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetMetadataByID returns a single metadata consumer.
func (s *Sonarr) GetMetadataByID(id int64) (*MetadataOutput, error) {
	return s.GetMetadataByIDContext(context.Background(), id)
}

// GetMetadataByIDContext returns a single metadata consumer.
func (s *Sonarr) GetMetadataByIDContext(ctx context.Context, id int64) (*MetadataOutput, error) {
	var output MetadataOutput

	req := starr.Request{URI: path.Join(bpMetadata, starr.Str(id))}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// GetMetadataSchema returns metadata consumer templates.
func (s *Sonarr) GetMetadataSchema() ([]*MetadataOutput, error) {
	return s.GetMetadataSchemaContext(context.Background())
}

// GetMetadataSchemaContext returns metadata consumer templates.
func (s *Sonarr) GetMetadataSchemaContext(ctx context.Context) ([]*MetadataOutput, error) {
	var output []*MetadataOutput

	req := starr.Request{URI: path.Join(bpMetadata, "schema")}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// AddMetadata creates a metadata consumer.
func (s *Sonarr) AddMetadata(input *MetadataInput, forceSave bool) (*MetadataOutput, error) {
	return s.AddMetadataContext(context.Background(), input, forceSave)
}

// AddMetadataContext creates a metadata consumer.
func (s *Sonarr) AddMetadataContext(
	ctx context.Context, input *MetadataInput, forceSave bool,
) (*MetadataOutput, error) {
	var output MetadataOutput

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(input); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpMetadata, err)
	}

	q := url.Values{}
	if forceSave {
		q.Set("forceSave", "true")
	}

	req := starr.Request{URI: bpMetadata, Body: &body, Query: q}
	if err := s.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateMetadata updates a metadata consumer.
func (s *Sonarr) UpdateMetadata(input *MetadataInput, forceSave bool) (*MetadataOutput, error) {
	return s.UpdateMetadataContext(context.Background(), input, forceSave)
}

// UpdateMetadataContext updates a metadata consumer.
func (s *Sonarr) UpdateMetadataContext(
	ctx context.Context, input *MetadataInput, forceSave bool,
) (*MetadataOutput, error) {
	var output MetadataOutput

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(input); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpMetadata, err)
	}

	params := url.Values{}
	if forceSave {
		params.Set("forceSave", "true")
	}

	uri := path.Join(bpMetadata, starr.Str(input.ID))

	req := starr.Request{URI: uri, Body: &body, Query: params}
	if err := s.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteMetadata deletes a metadata consumer.
func (s *Sonarr) DeleteMetadata(id int64) error {
	return s.DeleteMetadataContext(context.Background(), id)
}

// DeleteMetadataContext deletes a metadata consumer.
func (s *Sonarr) DeleteMetadataContext(ctx context.Context, id int64) error {
	req := starr.Request{URI: path.Join(bpMetadata, starr.Str(id))}
	if err := s.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}

// MetadataAction runs a named action on a metadata consumer.
func (s *Sonarr) MetadataAction(name string, input *MetadataInput) error {
	return s.MetadataActionContext(context.Background(), name, input)
}

// MetadataActionContext runs a named action on a metadata consumer.
func (s *Sonarr) MetadataActionContext(ctx context.Context, name string, input *MetadataInput) error {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(input); err != nil {
		return fmt.Errorf("json.Marshal(%s): %w", bpMetadata, err)
	}

	var output any

	req := starr.Request{URI: path.Join(bpMetadata, "action", path.Base(name)), Body: &body}
	if err := s.PostInto(ctx, req, &output); err != nil {
		return fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return nil
}

// TestMetadata tests a metadata consumer configuration.
func (s *Sonarr) TestMetadata(input *MetadataInput, forceTest bool) error {
	return s.TestMetadataContext(context.Background(), input, forceTest)
}

// TestMetadataContext tests a metadata consumer configuration.
func (s *Sonarr) TestMetadataContext(ctx context.Context, input *MetadataInput, forceTest bool) error {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(input); err != nil {
		return fmt.Errorf("json.Marshal(%s): %w", bpMetadata, err)
	}

	query := url.Values{}
	if forceTest {
		query.Set("forceTest", "true")
	}

	var output any

	req := starr.Request{URI: path.Join(bpMetadata, "test"), Body: &body, Query: query}
	if err := s.PostInto(ctx, req, &output); err != nil {
		return fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return nil
}

// TestAllMetadata tests all metadata consumers.
func (s *Sonarr) TestAllMetadata() error {
	return s.TestAllMetadataContext(context.Background())
}

// TestAllMetadataContext tests all metadata consumers.
func (s *Sonarr) TestAllMetadataContext(ctx context.Context) error {
	var output any

	req := starr.Request{URI: path.Join(bpMetadata, "testall")}
	if err := s.PostInto(ctx, req, &output); err != nil {
		return fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return nil
}
