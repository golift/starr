package lidarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"path"
	"time"

	"golift.io/starr"
)

const bpAlbum = APIver + "/album"

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
	Genres         []string         `json:"genres"`
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
	req := starr.Request{Query: make(url.Values), URI: bpAlbum}
	if mbID != "" {
		req.Query.Add("ForeignAlbumId", mbID)
	}

	var output []*Album

	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetAlbumByID returns an album by DB ID.
func (l *Lidarr) GetAlbumByID(albumID int64) (*Album, error) {
	return l.GetAlbumByIDContext(context.Background(), albumID)
}

// GetAlbumByIDContext returns an album by DB ID.
func (l *Lidarr) GetAlbumByIDContext(ctx context.Context, albumID int64) (*Album, error) {
	var output Album

	req := starr.Request{URI: path.Join(bpAlbum, starr.Itoa(albumID))}
	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateAlbum updates an album in place; the output of this is currently unknown!!!!
func (l *Lidarr) UpdateAlbum(albumID int64, album *Album, moveFiles bool) (*Album, error) {
	return l.UpdateAlbumContext(context.Background(), albumID, album, moveFiles)
}

// UpdateAlbumContext updates an album in place; the output of this is currently unknown!!!!
func (l *Lidarr) UpdateAlbumContext(ctx context.Context, albumID int64, album *Album, moveFiles bool) (*Album, error) {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(album); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpAlbum, err)
	}

	var output Album

	req := starr.Request{
		URI:   path.Join(bpAlbum, starr.Itoa(albumID)),
		Query: make(url.Values),
		Body:  &body,
	}
	req.Query.Add("moveFiles", starr.Itoa(moveFiles))

	if err := l.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
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

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(album); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpAlbum, err)
	}

	req := starr.Request{
		URI:   bpAlbum,
		Query: make(url.Values),
		Body:  &body,
	}

	var output Album
	if err := l.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
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

	req := starr.Request{
		URI:   path.Join(bpAlbum, "lookup"),
		Query: make(url.Values),
	}
	req.Query.Set("term", term)

	err := l.GetInto(ctx, req, &output)
	if err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// DeleteAlbum removes an album from the database.
// Setting deleteFiles true will delete all content for the album.
func (l *Lidarr) DeleteAlbum(albumID int64, deleteFiles, addImportExclusion bool) error {
	return l.DeleteAlbumContext(context.Background(), albumID, deleteFiles, addImportExclusion)
}

// DeleteAlbumContext removes an album from the database.
// Setting deleteFiles true will delete all content for the album.
func (l *Lidarr) DeleteAlbumContext(ctx context.Context, albumID int64, deleteFiles, addImportExclusion bool) error {
	req := starr.Request{URI: path.Join(bpAlbum, starr.Itoa(albumID)), Query: make(url.Values)}
	req.Query.Set("deleteFiles", starr.Itoa(deleteFiles))
	req.Query.Set("addImportListExclusion", starr.Itoa(addImportExclusion))

	if err := l.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
