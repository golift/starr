package lidarr

import (
	"context"
	"fmt"
	"net/url"

	"golift.io/starr"
)

const bpTrack = APIver + "/track"

// Track is an album track.
type Track struct {
	ArtistID            int64          `json:"artistId"`
	ForeignTrackID      string         `json:"foreignTrackId"`
	ForeignRecordingID  string         `json:"foreignRecordingId"`
	TrackFileID         int64          `json:"trackFileId"`
	AlbumID             int64          `json:"albumId"`
	Explicit            bool           `json:"explicit"`
	AbsoluteTrackNumber int            `json:"absoluteTrackNumber"`
	TrackNumber         string         `json:"trackNumber"`
	Title               string         `json:"title"`
	Duration            int            `json:"duration"`
	MediumNumber        int            `json:"mediumNumber"`
	HasFile             bool           `json:"hasFile"`
	Ratings             *starr.Ratings `json:"ratings"`
	Grabbed             bool           `json:"grabbed"`
	ID                  int64          `json:"id"`
	Artist              *Artist        `json:"artist"`    // probably empty.
	TrackFile           *TrackFile     `json:"trackFile"` // probably empty.
}

// GetTracks by their IDs.
func (l *Lidarr) GetTracks(trackID ...int64) ([]*Track, error) {
	return l.GetTracksContext(context.Background(), trackID...)
}

// GetTracksContext gets track files by their IDs using a provided context.
func (l *Lidarr) GetTracksContext(ctx context.Context, trackID ...int64) ([]*Track, error) {
	req := starr.Request{URI: bpTrack, Query: make(url.Values)}
	for _, id := range trackID {
		req.Query.Add("trackIds", starr.Str(id))
	}

	var output []*Track
	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetTracksByAlbum gets track files using an album ID.
func (l *Lidarr) GetTracksByAlbum(albumID int64) ([]*Track, error) {
	return l.GetTracksByAlbumContext(context.Background(), albumID)
}

// GetTracksByAlbumContext gets track files using an album ID.
func (l *Lidarr) GetTracksByAlbumContext(ctx context.Context, albumID int64) ([]*Track, error) {
	req := starr.Request{URI: bpTrack, Query: make(url.Values)}
	req.Query.Add("albumId", starr.Str(albumID))

	var output []*Track
	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetTracksByArtist gets track files using an artist ID.
func (l *Lidarr) GetTracksByArtist(artistID int64) ([]*Track, error) {
	return l.GetTracksByArtistContext(context.Background(), artistID)
}

// GetTracksByAlbumRelease gets track files using an artist ID.
func (l *Lidarr) GetTracksByArtistContext(ctx context.Context, artistID int64) ([]*Track, error) {
	req := starr.Request{URI: bpTrack, Query: make(url.Values)}
	req.Query.Add("artistId", starr.Str(artistID))

	var output []*Track
	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetTracksByAlbumRelease gets track files using an album release ID.
func (l *Lidarr) GetTracksByAlbumRelease(albumID int64) ([]*Track, error) {
	return l.GetTracksByAlbumContext(context.Background(), albumID)
}

// GetTracksByAlbumReleaseContext gets track files using an album release ID.
func (l *Lidarr) GetTracksByAlbumReleaseContext(ctx context.Context, albumReleaseID int64) ([]*Track, error) {
	req := starr.Request{URI: bpTrack, Query: make(url.Values)}
	req.Query.Add("albumReleaseId", starr.Str(albumReleaseID))

	var output []*Track
	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}
