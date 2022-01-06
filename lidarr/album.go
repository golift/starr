package lidarr

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"golift.io/starr"
)

// GetAlbum returns an album or all albums if mbID is "" (empty).
// mbID is the music brainz UUID for a "release-group".
func (l *Lidarr) GetAlbum(mbID string) ([]*Album, error) {
	params := make(url.Values)

	if mbID != "" {
		params.Add("ForeignAlbumId", mbID)
	}

	var albums []*Album

	err := l.GetInto("v1/album", params, &albums)
	if err != nil {
		return nil, fmt.Errorf("api.Get(album): %w", err)
	}

	return albums, nil
}

// GetAlbumByID returns an album by DB ID.
func (l *Lidarr) GetAlbumByID(albumID int64) (*Album, error) {
	var album Album

	err := l.GetInto("v1/album/"+strconv.FormatInt(albumID, starr.Base10), nil, &album)
	if err != nil {
		return nil, fmt.Errorf("api.Get(album): %w", err)
	}

	return &album, nil
}

// UpdateAlbum updates an album in place; the output of this is currently unknown!!!!
func (l *Lidarr) UpdateAlbum(albumID int64, album *Album) (*Album, error) {
	put, err := json.Marshal(album)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal(album): %w", err)
	}

	params := make(url.Values)
	params.Add("moveFiles", "true")

	var output Album

	err = l.PutInto("v1/album/"+strconv.FormatInt(albumID, starr.Base10), params, put, &output)
	if err != nil {
		return nil, fmt.Errorf("api.Put(album): %w", err)
	}

	return &output, nil
}

// AddAlbum adds a new album to Lidarr, and probably does not yet work.
func (l *Lidarr) AddAlbum(album *AddAlbumInput) (*Album, error) {
	if album.Releases == nil {
		album.Releases = make([]*AddAlbumInputRelease, 0)
	}

	body, err := json.Marshal(album)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal(album): %w", err)
	}

	params := make(url.Values)
	params.Add("moveFiles", "true")

	var output Album

	err = l.PostInto("v1/album", params, body, &output)
	if err != nil {
		return nil, fmt.Errorf("api.Post(album): %w", err)
	}

	return &output, nil
}

// Lookup will search for albums matching the specified search term.
func (l *Lidarr) Lookup(term string) ([]*Album, error) {
	var output []*Album

	if term == "" {
		return output, nil
	}

	params := make(url.Values)
	params.Set("term", term)

	err := l.GetInto("v1/album/lookup", params, &output)
	if err != nil {
		return nil, fmt.Errorf("api.Get(album/lookup): %w", err)
	}

	return output, nil
}
