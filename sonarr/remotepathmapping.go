package sonarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"

	"golift.io/starr"
)

// Define Base Path for remote path mapping calls.
const bpRemotePathMapping = APIver + "/remotePathMapping"

// RemotePathMapping is the /api/v3/remotePathMapping endpoint.
type RemotePathMapping struct {
	ID         int64  `json:"id,omitempty"`
	Host       string `json:"host"`
	RemotePath string `json:"remotePath"`
	LocalPath  string `json:"localPath"`
}

// GetRemotePathMappings returns all configured remote path mappings.
func (s *Sonarr) GetRemotePathMappings() ([]*RemotePathMapping, error) {
	return s.GetRemotePathMappingsContext(context.Background())
}

// GetRemotePathMappingsContext returns all configured remote path mappings.
func (s *Sonarr) GetRemotePathMappingsContext(ctx context.Context) ([]*RemotePathMapping, error) {
	var output []*RemotePathMapping

	req := starr.Request{URI: bpRemotePathMapping}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetRemotePathMapping returns a single remote path mapping.
func (s *Sonarr) GetRemotePathMapping(mappingID int64) (*RemotePathMapping, error) {
	return s.GetRemotePathMappingContext(context.Background(), mappingID)
}

// GetRemotePathMappingContext returns a single remote path mapping.
func (s *Sonarr) GetRemotePathMappingContext(ctx context.Context, mappingID int64) (*RemotePathMapping, error) {
	var output RemotePathMapping

	req := starr.Request{URI: path.Join(bpRemotePathMapping, fmt.Sprint(mappingID))}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// AddRemotePathMapping creates a remote path mapping.
func (s *Sonarr) AddRemotePathMapping(mapping *RemotePathMapping) (*RemotePathMapping, error) {
	return s.AddRemotePathMappingContext(context.Background(), mapping)
}

// AddRemotePathMappingContext creates a remote path mapping.
func (s *Sonarr) AddRemotePathMappingContext(ctx context.Context,
	mapping *RemotePathMapping,
) (*RemotePathMapping, error) {
	var output RemotePathMapping

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(mapping); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpRemotePathMapping, err)
	}

	req := starr.Request{URI: bpRemotePathMapping, Body: &body}
	if err := s.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateRemotePathMapping updates the remote path mapping.
func (s *Sonarr) UpdateRemotePathMapping(mapping *RemotePathMapping) (*RemotePathMapping, error) {
	return s.UpdateRemotePathMappingContext(context.Background(), mapping)
}

// UpdateRemotePathMappingContext updates the remote path mapping.
func (s *Sonarr) UpdateRemotePathMappingContext(ctx context.Context,
	mapping *RemotePathMapping,
) (*RemotePathMapping, error) {
	var output RemotePathMapping

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(mapping); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpRemotePathMapping, err)
	}

	req := starr.Request{URI: path.Join(bpRemotePathMapping, fmt.Sprint(mapping.ID)), Body: &body}
	if err := s.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteRemotePathMapping removes a single remote path mapping.
func (s *Sonarr) DeleteRemotePathMapping(mappingID int64) error {
	return s.DeleteRemotePathMappingContext(context.Background(), mappingID)
}

// DeleteRemotePathMappingContext removes a single remote path mapping.
func (s *Sonarr) DeleteRemotePathMappingContext(ctx context.Context, mappingID int64) error {
	req := starr.Request{URI: path.Join(bpRemotePathMapping, fmt.Sprint(mappingID))}
	if err := s.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
