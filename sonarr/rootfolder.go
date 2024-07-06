package sonarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"

	"golift.io/starr"
)

// Define Base Path for MediaManagement calls.
const bpRootFolder = APIver + "/rootFolder"

// RootFolder is the /api/v3/rootfolder endpoint.
type RootFolder struct {
	Accessible      bool          `json:"accessible,omitempty"`
	ID              int64         `json:"id,omitempty"`
	FreeSpace       int64         `json:"freeSpace,omitempty"`
	Path            string        `json:"path"`
	UnmappedFolders []*starr.Path `json:"unmappedFolders,omitempty"`
}

// GetRootFolders returns all configured root folders.
func (s *Sonarr) GetRootFolders() ([]*RootFolder, error) {
	return s.GetRootFoldersContext(context.Background())
}

// GetRootFoldersContext returns all configured root folders.
func (s *Sonarr) GetRootFoldersContext(ctx context.Context) ([]*RootFolder, error) {
	var output []*RootFolder

	req := starr.Request{URI: bpRootFolder}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetRootFolder returns a single root folder.
func (s *Sonarr) GetRootFolder(folderID int64) (*RootFolder, error) {
	return s.GetRootFolderContext(context.Background(), folderID)
}

// GetRootFolderContext returns a single root folder.
func (s *Sonarr) GetRootFolderContext(ctx context.Context, folderID int64) (*RootFolder, error) {
	var output RootFolder

	req := starr.Request{URI: path.Join(bpRootFolder, starr.Str(folderID))}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// AddRootFolder creates a root folder.
func (s *Sonarr) AddRootFolder(folder *RootFolder) (*RootFolder, error) {
	return s.AddRootFolderContext(context.Background(), folder)
}

// AddRootFolderContext creates a root folder.
func (s *Sonarr) AddRootFolderContext(ctx context.Context, folder *RootFolder) (*RootFolder, error) {
	var output RootFolder

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(folder); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpRootFolder, err)
	}

	req := starr.Request{URI: bpRootFolder, Body: &body}
	if err := s.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteRootFolder removes a single root folder.
func (s *Sonarr) DeleteRootFolder(folderID int64) error {
	return s.DeleteRootFolderContext(context.Background(), folderID)
}

// DeleteRootFolderContext removes a single root folder.
func (s *Sonarr) DeleteRootFolderContext(ctx context.Context, folderID int64) error {
	req := starr.Request{URI: path.Join(bpRootFolder, starr.Str(folderID))}
	if err := s.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
