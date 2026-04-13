package sonarr

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"path"

	"golift.io/starr"
)

const bpFilesystem = APIver + "/filesystem"

// FilesystemQuery is the query for /api/v3/filesystem.
type FilesystemQuery struct {
	Path                               string
	IncludeFiles                       bool
	AllowFoldersWithoutTrailingSlashes bool
}

// Values builds query parameters for the filesystem browser.
func (q *FilesystemQuery) Values() url.Values {
	val := make(url.Values)
	if q == nil {
		return val
	}

	if q.Path != "" {
		val.Set("path", q.Path)
	}

	val.Set("includeFiles", starr.Str(q.IncludeFiles))
	val.Set("allowFoldersWithoutTrailingSlashes", starr.Str(q.AllowFoldersWithoutTrailingSlashes))

	return val
}

// BrowseFilesystem lists files and folders for a path.
func (s *Sonarr) BrowseFilesystem(query *FilesystemQuery) ([]*starr.Path, error) {
	return s.BrowseFilesystemContext(context.Background(), query)
}

// BrowseFilesystemContext lists files and folders for a path.
func (s *Sonarr) BrowseFilesystemContext(ctx context.Context, query *FilesystemQuery) ([]*starr.Path, error) {
	var output []*starr.Path

	req := starr.Request{URI: bpFilesystem, Query: query.Values()}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// BrowseFilesystemMediaFiles lists media files under a path.
func (s *Sonarr) BrowseFilesystemMediaFiles(pathName string) ([]*starr.Path, error) {
	return s.BrowseFilesystemMediaFilesContext(context.Background(), pathName)
}

// BrowseFilesystemMediaFilesContext lists media files under a path.
func (s *Sonarr) BrowseFilesystemMediaFilesContext(ctx context.Context, pathName string) ([]*starr.Path, error) {
	var output []*starr.Path

	params := make(url.Values)
	if pathName != "" {
		params.Set("path", pathName)
	}

	req := starr.Request{URI: path.Join(bpFilesystem, "mediafiles"), Query: params}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetFilesystemType returns the raw response body from /api/v3/filesystem/type.
func (s *Sonarr) GetFilesystemType(pathName string) ([]byte, error) {
	return s.GetFilesystemTypeContext(context.Background(), pathName)
}

// GetFilesystemTypeContext returns the raw response body from /api/v3/filesystem/type.
func (s *Sonarr) GetFilesystemTypeContext(ctx context.Context, pathName string) ([]byte, error) {
	params := make(url.Values)
	params.Set("path", pathName)

	req := starr.Request{URI: path.Join(bpFilesystem, "type"), Query: params}

	resp, err := s.Get(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading HTTP response body: %w", err)
	}

	return body, nil
}
