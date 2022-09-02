package sonarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"
	"strconv"

	"golift.io/starr"
)

// RootFolder is the /api/v3/rootfolder endpoint.
type RootFolder struct {
	Accessible      bool          `json:"accessible,omitempty"`
	ID              int64         `json:"id,omitempty"`
	FreeSpace       int64         `json:"freeSpace,omitempty"`
	Path            string        `json:"path"`
	UnmappedFolders []*starr.Path `json:"unmappedFolders,omitempty"`
}

// Define Base Path for MediaManagement calls.
const bpRootFolder = APIver + "/rootFolder"

// GetRootFolders returns all configured root folders.
func (s *Sonarr) GetRootFolders() ([]*RootFolder, error) {
	return s.GetRootFoldersContext(context.Background())
}

func (s *Sonarr) GetRootFoldersContext(ctx context.Context) ([]*RootFolder, error) {
	var output []*RootFolder

	if err := s.GetInto(ctx, bpRootFolder, nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(rootFolder): %w", err)
	}

	return output, nil
}

// GetRootFolder returns a single root folder.
func (s *Sonarr) GetRootFolder(folderID int) (*RootFolder, error) {
	return s.GetRootFolderContext(context.Background(), folderID)
}

func (s *Sonarr) GetRootFolderContext(ctx context.Context, folderID int) (*RootFolder, error) {
	var output *RootFolder

	uri := path.Join(bpRootFolder, strconv.Itoa(folderID))
	if err := s.GetInto(ctx, uri, nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(rootFolder): %w", err)
	}

	return output, nil
}

// AddRootFolder creates a root folder.
func (s *Sonarr) AddRootFolder(folder *RootFolder) (*RootFolder, error) {
	return s.AddRootFolderContext(context.Background(), folder)
}

func (s *Sonarr) AddRootFolderContext(ctx context.Context, folder *RootFolder) (*RootFolder, error) {
	var output RootFolder

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(folder); err != nil {
		return nil, fmt.Errorf("json.Marshal(rootFolder): %w", err)
	}

	if err := s.PostInto(ctx, bpRootFolder, nil, &body, &output); err != nil {
		return nil, fmt.Errorf("api.Post(rootFolder): %w", err)
	}

	return &output, nil
}

// DeleteRootFolder removes a single root folder.
func (s *Sonarr) DeleteRootFolder(folderID int) error {
	return s.DeleteRootFolderContext(context.Background(), folderID)
}

func (s *Sonarr) DeleteRootFolderContext(ctx context.Context, folderID int) error {
	var output interface{}

	uri := path.Join(bpRootFolder, strconv.Itoa(folderID))
	if err := s.DeleteInto(ctx, uri, nil, &output); err != nil {
		return fmt.Errorf("api.Delete(rootFolder): %w", err)
	}

	return nil
}
