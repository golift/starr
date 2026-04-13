package sonarr

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"path"

	"golift.io/starr"
)

const bpMediaCover = APIver + "/mediacover"

// GetMediaCover downloads a series media cover image (jpg, png, or gif).
func (s *Sonarr) GetMediaCover(seriesID int64, filename string) ([]byte, error) {
	return s.GetMediaCoverContext(context.Background(), seriesID, filename)
}

// GetMediaCoverContext downloads a series media cover image (jpg, png, or gif).
func (s *Sonarr) GetMediaCoverContext(ctx context.Context, seriesID int64, filename string) ([]byte, error) {
	uri := starr.SetAPIPath(path.Join(bpMediaCover, starr.Str(seriesID), url.PathEscape(filename)))

	req := starr.Request{URI: uri}

	resp, err := s.Get(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body from %s: %w", uri, err)
	}

	return body, nil
}
