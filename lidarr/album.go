package lidarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"golift.io/starr"
)

// Album is the /api/v1/album endpoint.
type Album struct {
	ID             int64            `json:"id,omitempty"`
	Title          string           `json:"title"`
	Disambiguation string           `json:"disambiguation"`
	Overview       string           `json:"overview"`
	ArtistID       int64            `json:"artistId"`
	ForeignAlbumID string           `json:"foreignAlbumId"`
	ProfileID      int64            `json:"profileId"`
	Duration       int              `json:"duration"`
	AlbumType      string           `json:"albumType"`
	SecondaryTypes []interface{}    `json:"secondaryTypes"`
	MediumCount    int              `json:"mediumCount"`
	Ratings        *starr.Ratings   `json:"ratings"`
	ReleaseDate    time.Time        `json:"releaseDate"`
	Releases       []*Release       `json:"releases"`
	Genres         []interface{}    `json:"genres"`
	Media          []*Media         `json:"media"`
	Artist         *Artist          `json:"artist"`
	Links          []*starr.Link    `json:"links"`
	Images         []*starr.Image   `json:"images"`
	Statistics     *Statistics      `json:"statistics"`
	RemoteCover    string           `json:"remoteCover,omitempty"`
	AddOptions     *AlbumAddOptions `json:"addOptions,omitempty"`
	Monitored      bool             `json:"monitored"`
	AnyReleaseOk   bool             `json:"anyReleaseOk"`
	Grabbed        bool             `json:"grabbed"`
}

// Release is part of an Album.
type Release struct {
	ID               int64    `json:"id"`
	AlbumID          int64    `json:"albumId"`
	ForeignReleaseID string   `json:"foreignReleaseId"`
	Title            string   `json:"title"`
	Status           string   `json:"status"`
	Duration         int      `json:"duration"`
	TrackCount       int      `json:"trackCount"`
	Media            []*Media `json:"media"`
	MediumCount      int      `json:"mediumCount"`
	Disambiguation   string   `json:"disambiguation"`
	Country          []string `json:"country"`
	Label            []string `json:"label"`
	Format           string   `json:"format"`
	Monitored        bool     `json:"monitored"`
}

// Media is part of an Album.
type Media struct {
	MediumNumber int64  `json:"mediumNumber"`
	MediumName   string `json:"mediumName"`
	MediumFormat string `json:"mediumFormat"`
}

// ArtistAddOptions is part of an artist and an album.
type ArtistAddOptions struct {
	Monitor                string `json:"monitor,omitempty"`
	Monitored              bool   `json:"monitored,omitempty"`
	SearchForMissingAlbums bool   `json:"searchForMissingAlbums,omitempty"`
}

// AddAlbumInput is currently unknown.
type AddAlbumInput struct {
	ForeignAlbumID string                  `json:"foreignAlbumId"`
	Monitored      bool                    `json:"monitored"`
	Releases       []*AddAlbumInputRelease `json:"releases"`
	AddOptions     *AlbumAddOptions        `json:"addOptions"`
	Artist         *Artist                 `json:"artist"`
}

// AddAlbumInputRelease is part of AddAlbumInput.
type AddAlbumInputRelease struct {
	ForeignReleaseID string   `json:"foreignReleaseId"`
	Title            string   `json:"title"`
	Media            []*Media `json:"media"`
	Monitored        bool     `json:"monitored"`
}

// AlbumAddOptions is part of an Album.
type AlbumAddOptions struct {
	SearchForNewAlbum bool `json:"searchForNewAlbum,omitempty"`
}

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

	err := l.GetInto(ctx, "v1/album", params, &albums)
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

	err := l.GetInto(ctx, "v1/album/"+strconv.FormatInt(albumID, starr.Base10), nil, &album)
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

	err := l.PutInto(ctx, "v1/album/"+strconv.FormatInt(albumID, starr.Base10), params, &body, &output)
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
	if err := l.PostInto(ctx, "v1/album", params, &body, &output); err != nil {
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

	err := l.GetInto(ctx, "v1/album/lookup", params, &output)
	if err != nil {
		return nil, fmt.Errorf("api.Get(album/lookup): %w", err)
	}

	return output, nil
}
