package lidarr

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

// GetRemotePathMappings returns all configured remote path mappings.
func (l *Lidarr) GetRemotePathMappings() ([]*starr.RemotePathMapping, error) {
	return l.GetRemotePathMappingsContext(context.Background())
}

// GetRemotePathMappingsContext returns all configured remote path mappings.
func (l *Lidarr) GetRemotePathMappingsContext(ctx context.Context) ([]*starr.RemotePathMapping, error) {
	var output []*starr.RemotePathMapping

	req := starr.Request{URI: bpRemotePathMapping}
	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetRemotePathMapping returns a single remote path mapping.
func (l *Lidarr) GetRemotePathMapping(mappingID int64) (*starr.RemotePathMapping, error) {
	return l.GetRemotePathMappingContext(context.Background(), mappingID)
}

// GetRemotePathMappingContext returns a single remote path mapping.
func (l *Lidarr) GetRemotePathMappingContext(ctx context.Context, mappingID int64) (*starr.RemotePathMapping, error) {
	var output starr.RemotePathMapping

	req := starr.Request{URI: path.Join(bpRemotePathMapping, fmt.Sprint(mappingID))}
	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// AddRemotePathMapping creates a remote path mapping.
func (l *Lidarr) AddRemotePathMapping(mapping *starr.RemotePathMapping) (*starr.RemotePathMapping, error) {
	return l.AddRemotePathMappingContext(context.Background(), mapping)
}

// AddRemotePathMappingContext creates a remote path mapping.
func (l *Lidarr) AddRemotePathMappingContext(ctx context.Context,
	mapping *starr.RemotePathMapping,
) (*starr.RemotePathMapping, error) {
	var output starr.RemotePathMapping

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(mapping); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpRemotePathMapping, err)
	}

	req := starr.Request{URI: bpRemotePathMapping, Body: &body}
	if err := l.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateRemotePathMapping updates the remote path mapping.
func (l *Lidarr) UpdateRemotePathMapping(mapping *starr.RemotePathMapping) (*starr.RemotePathMapping, error) {
	return l.UpdateRemotePathMappingContext(context.Background(), mapping)
}

// UpdateRemotePathMappingContext updates the remote path mapping.
func (l *Lidarr) UpdateRemotePathMappingContext(ctx context.Context,
	mapping *starr.RemotePathMapping,
) (*starr.RemotePathMapping, error) {
	var output starr.RemotePathMapping

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(mapping); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpRemotePathMapping, err)
	}

	req := starr.Request{URI: path.Join(bpRemotePathMapping, fmt.Sprint(mapping.ID)), Body: &body}
	if err := l.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteRemotePathMapping removes a single remote path mapping.
func (l *Lidarr) DeleteRemotePathMapping(mappingID int64) error {
	return l.DeleteRemotePathMappingContext(context.Background(), mappingID)
}

// DeleteRemotePathMappingContext removes a single remote path mapping.
func (l *Lidarr) DeleteRemotePathMappingContext(ctx context.Context, mappingID int64) error {
	req := starr.Request{URI: path.Join(bpRemotePathMapping, fmt.Sprint(mappingID))}
	if err := l.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
