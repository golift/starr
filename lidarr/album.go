package lidarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"golift.io/starr"
)

// GetAlbum returns an album or all albums if mbID is "" (empty).
// mbID is the music brainz UUID for a "release-group".
func (l *Lidarr) GetAlbum(mbID string) ([]*Album, error) {
	return l.GetAlbumContext(context.Background(), mbID)
}

// GetAlbumContext returns an album or all albums if mbID is "" (empty).
// mbID is the music brainz UUID for a "release-group".
func (l *Lidarr) GetAlbumContext(ctx context.Context, mbID string) ([]*Album, error) {
	params := make(url.Values)

	if mbID != "" {
		params.Add("ForeignAlbumId", mbID)
	}

	var albums []*Album

	_, err := l.GetInto(ctx, "v1/album", params, &albums)
	if err != nil {
		return nil, fmt.Errorf("api.Get(album): %w", err)
	}

	return albums, nil
}

// GetAlbumByID returns an album by DB ID.
func (l *Lidarr) GetAlbumByID(albumID int64) (*Album, error) {
	return l.GetAlbumByIDContext(context.Background(), albumID)
}

// GetAlbumByIDContext returns an album by DB ID.
func (l *Lidarr) GetAlbumByIDContext(ctx context.Context, albumID int64) (*Album, error) {
	var album Album

	_, err := l.GetInto(ctx, "v1/album/"+strconv.FormatInt(albumID, starr.Base10), nil, &album)
	if err != nil {
		return nil, fmt.Errorf("api.Get(album): %w", err)
	}

	return &album, nil
}

// UpdateAlbum updates an album in place; the output of this is currently unknown!!!!
func (l *Lidarr) UpdateAlbum(albumID int64, album *Album) (*Album, error) {
	return l.UpdateAlbumContext(context.Background(), albumID, album)
}

// UpdateAlbumContext updates an album in place; the output of this is currently unknown!!!!
func (l *Lidarr) UpdateAlbumContext(ctx context.Context, albumID int64, album *Album) (*Album, error) {
	params := make(url.Values)
	params.Add("moveFiles", "true")

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(album); err != nil {
		return nil, fmt.Errorf("json.Marshal(album): %w", err)
	}

	var output Album

	_, err := l.PutInto(ctx, "v1/album/"+strconv.FormatInt(albumID, starr.Base10), params, &body, &output)
	if err != nil {
		return nil, fmt.Errorf("api.Put(album): %w", err)
	}

	return &output, nil
}

// AddAlbum adds a new album to Lidarr, and probably does not yet work.
func (l *Lidarr) AddAlbum(album *AddAlbumInput) (*Album, error) {
	return l.AddAlbumContext(context.Background(), album)
}

// AddAlbumContext adds a new album to Lidarr, and probably does not yet work.
func (l *Lidarr) AddAlbumContext(ctx context.Context, album *AddAlbumInput) (*Album, error) {
	if album.Releases == nil {
		album.Releases = make([]*AddAlbumInputRelease, 0)
	}

	params := make(url.Values)
	params.Add("moveFiles", "true")

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(album); err != nil {
		return nil, fmt.Errorf("json.Marshal(album): %w", err)
	}

	var output Album
	if _, err := l.PostInto(ctx, "v1/album", params, &body, &output); err != nil {
		return nil, fmt.Errorf("api.Post(album): %w", err)
	}

	return &output, nil
}

// Lookup will search for albums matching the specified search term.
func (l *Lidarr) Lookup(term string) ([]*Album, error) {
	return l.LookupContext(context.Background(), term)
}

// LookupContext will search for albums matching the specified search term.
func (l *Lidarr) LookupContext(ctx context.Context, term string) ([]*Album, error) {
	var output []*Album

	if term == "" {
		return output, nil
	}

	params := make(url.Values)
	params.Set("term", term)

	_, err := l.GetInto(ctx, "v1/album/lookup", params, &output)
	if err != nil {
		return nil, fmt.Errorf("api.Get(album/lookup): %w", err)
	}

	return output, nil
}
