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

// GetArtist returns an artist or all artists.
func (l *Lidarr) GetArtist(mbID string) ([]*Artist, error) {
	return l.GetArtistContext(context.Background(), mbID)
}

// GetArtistContext returns an artist or all artists.
func (l *Lidarr) GetArtistContext(ctx context.Context, mbID string) ([]*Artist, error) {
	params := make(url.Values)

	if mbID != "" {
		params.Add("mbId", mbID)
	}

	var artist []*Artist

	err := l.GetInto(ctx, "v1/artist", params, &artist)
	if err != nil {
		return artist, fmt.Errorf("api.Get(artist): %w", err)
	}

	return artist, nil
}

// GetArtistByID returns an artist from an ID.
func (l *Lidarr) GetArtistByID(artistID int64) (*Artist, error) {
	return l.GetArtistByIDContext(context.Background(), artistID)
}

// GetArtistByIDContext returns an artist from an ID.
func (l *Lidarr) GetArtistByIDContext(ctx context.Context, artistID int64) (*Artist, error) {
	var artist Artist

	err := l.GetInto(ctx, "v1/artist/"+strconv.FormatInt(artistID, starr.Base10), nil, &artist)
	if err != nil {
		return &artist, fmt.Errorf("api.Get(artist): %w", err)
	}

	return &artist, nil
}

// AddArtist adds a new artist to Lidarr, and probably does not yet work.
func (l *Lidarr) AddArtist(artist *Artist) (*Artist, error) {
	return l.AddArtistContext(context.Background(), artist)
}

// AddArtistContext adds a new artist to Lidarr, and probably does not yet work.
func (l *Lidarr) AddArtistContext(ctx context.Context, artist *Artist) (*Artist, error) {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(artist); err != nil {
		return nil, fmt.Errorf("json.Marshal(artist): %w", err)
	}

	params := make(url.Values)
	params.Add("moveFiles", "true")

	var output Artist

	err := l.PostInto(ctx, "v1/artist", params, &body, &output)
	if err != nil {
		return nil, fmt.Errorf("api.Post(artist): %w", err)
	}

	return &output, nil
}

// UpdateArtist updates an artist in place.
func (l *Lidarr) UpdateArtist(artist *Artist) (*Artist, error) {
	return l.UpdateArtistContext(context.Background(), artist)
}

// UpdateArtistContext updates an artist in place.
func (l *Lidarr) UpdateArtistContext(ctx context.Context, artist *Artist) (*Artist, error) {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(artist); err != nil {
		return nil, fmt.Errorf("json.Marshal(artist): %w", err)
	}

	params := make(url.Values)
	params.Add("moveFiles", "true")

	var output Artist

	err := l.PutInto(ctx, "v1/artist/"+strconv.FormatInt(artist.ID, starr.Base10), params, &body, &output)
	if err != nil {
		return nil, fmt.Errorf("api.Put(artist): %w", err)
	}

	return &output, nil
}
