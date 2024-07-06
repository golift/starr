package readarr

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
func (r *Readarr) GetRemotePathMappings() ([]*starr.RemotePathMapping, error) {
	return r.GetRemotePathMappingsContext(context.Background())
}

// GetRemotePathMappingsContext returns all configured remote path mappings.
func (r *Readarr) GetRemotePathMappingsContext(ctx context.Context) ([]*starr.RemotePathMapping, error) {
	var output []*starr.RemotePathMapping

	req := starr.Request{URI: bpRemotePathMapping}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetRemotePathMapping returns a single remote path mapping.
func (r *Readarr) GetRemotePathMapping(mappingID int64) (*starr.RemotePathMapping, error) {
	return r.GetRemotePathMappingContext(context.Background(), mappingID)
}

// GetRemotePathMappingContext returns a single remote path mapping.
func (r *Readarr) GetRemotePathMappingContext(ctx context.Context, mappingID int64) (*starr.RemotePathMapping, error) {
	var output starr.RemotePathMapping

	req := starr.Request{URI: path.Join(bpRemotePathMapping, starr.Itoa(mappingID))}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// AddRemotePathMapping creates a remote path mapping.
func (r *Readarr) AddRemotePathMapping(mapping *starr.RemotePathMapping) (*starr.RemotePathMapping, error) {
	return r.AddRemotePathMappingContext(context.Background(), mapping)
}

// AddRemotePathMappingContext creates a remote path mapping.
func (r *Readarr) AddRemotePathMappingContext(ctx context.Context,
	mapping *starr.RemotePathMapping,
) (*starr.RemotePathMapping, error) {
	var output starr.RemotePathMapping

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(mapping); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpRemotePathMapping, err)
	}

	req := starr.Request{URI: bpRemotePathMapping, Body: &body}
	if err := r.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateRemotePathMapping updates the remote path mapping.
func (r *Readarr) UpdateRemotePathMapping(mapping *starr.RemotePathMapping) (*starr.RemotePathMapping, error) {
	return r.UpdateRemotePathMappingContext(context.Background(), mapping)
}

// UpdateRemotePathMappingContext updates the remote path mapping.
func (r *Readarr) UpdateRemotePathMappingContext(ctx context.Context,
	mapping *starr.RemotePathMapping,
) (*starr.RemotePathMapping, error) {
	var output starr.RemotePathMapping

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(mapping); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpRemotePathMapping, err)
	}

	req := starr.Request{URI: path.Join(bpRemotePathMapping, starr.Itoa(mapping.ID)), Body: &body}
	if err := r.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteRemotePathMapping removes a single remote path mapping.
func (r *Readarr) DeleteRemotePathMapping(mappingID int64) error {
	return r.DeleteRemotePathMappingContext(context.Background(), mappingID)
}

// DeleteRemotePathMappingContext removes a single remote path mapping.
func (r *Readarr) DeleteRemotePathMappingContext(ctx context.Context, mappingID int64) error {
	req := starr.Request{URI: path.Join(bpRemotePathMapping, starr.Itoa(mappingID))}
	if err := r.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
