package radarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"golift.io/starr"
)

const bpRelease = APIver + "/release"

// Release is the output from the Radarr release endpoint.
type Release struct {
	ID                  int64          `json:"id"`
	GUID                string         `json:"guid"`
	Quality             *starr.Quality `json:"quality"`
	CustomFormats       []any          `json:"customFormats"`
	CustomFormatScore   int64          `json:"customFormatScore"`
	QualityWeight       int64          `json:"qualityWeight"`
	Age                 int64          `json:"age"`
	AgeHours            float64        `json:"ageHours"`
	AgeMinutes          float64        `json:"ageMinutes"`
	Size                int64          `json:"size"`
	IndexerID           int64          `json:"indexerId"`
	Indexer             string         `json:"indexer"`
	ReleaseGroup        string         `json:"releaseGroup"`
	ReleaseHash         string         `json:"releaseHash"`
	Title               string         `json:"title"`
	SceneSource         bool           `json:"sceneSource"`
	MovieTitles         []string       `json:"movieTitles"`
	Languages           []*starr.Value `json:"languages"`
	MappedMovieID       int64          `json:"mappedMovieId"`
	Approved            bool           `json:"approved"`
	TemporarilyRejected bool           `json:"temporarilyRejected"`
	Rejected            bool           `json:"rejected"`
	TmdbID              int64          `json:"tmdbId"`
	ImdbID              int64          `json:"imdbId"`
	Rejections          []string       `json:"rejections"`
	PublishDate         time.Time      `json:"publishDate"`
	CommentURL          string         `json:"commentUrl"`
	DownloadURL         string         `json:"downloadUrl"`
	InfoURL             string         `json:"infoUrl"`
	DownloadAllowed     bool           `json:"downloadAllowed"`
	ReleaseWeight       int64          `json:"releaseWeight"`
	Edition             string         `json:"edition"`
	MagnetURL           string         `json:"magnetUrl"`
	InfoHash            string         `json:"infoHash"`
	SubGroup            string         `json:"subGroup"`
	Seeders             int            `json:"seeders"`
	Leechers            int            `json:"leechers"`
	Protocol            starr.Protocol `json:"protocol"`
	IndexerFlags        []string       `json:"indexerFlags,omitempty"`
	MovieID             int64          `json:"movieId"`
	DownloadClientID    int64          `json:"downloadClientId"`
	DownloadClient      string         `json:"downloadClient"`
	ShouldOverride      bool           `json:"shouldOverride"`
}

// SearchRelease searches for and returns a list releases available for download.
func (r *Radarr) SearchRelease(movieID int64) ([]*Release, error) {
	return r.SearchReleaseContext(context.Background(), movieID)
}

// SearchReleaseContext searches for and returns a list releases available for download.
func (r *Radarr) SearchReleaseContext(ctx context.Context, movieID int64) ([]*Release, error) {
	req := starr.Request{URI: bpRelease, Query: make(url.Values)}
	req.Query.Set("movieId", starr.Str(movieID))

	var output []*Release
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// Grab attempts to download a release by GUID.
func (r *Radarr) Grab(guid string, indexerID int64) (*Release, error) {
	return r.GrabReleaseContext(context.Background(), &Release{IndexerID: indexerID, GUID: guid})
}

// GrabContext attempts to download a release by GUID.
func (r *Radarr) GrabContext(ctx context.Context, guid string, indexerID int64) (*Release, error) {
	return r.GrabReleaseContext(ctx, &Release{IndexerID: indexerID, GUID: guid})
}

// GrabRelease attempts to download a release for a movie from a search.
// Pass the release for the item from the SearchRelease output for the release you wish to download.
// If the release.MovieID is 0 then release.MappedMovieID is used. Both may be 0, and that's OK unless
// release.ShouldOverride is true. If release.ShouldOverride is true, then Languages, MovieID and Quality
// must be present in the release.
func (r *Radarr) GrabRelease(release *Release) (*Release, error) {
	return r.GrabReleaseContext(context.Background(), release)
}

// GrabReleaseContext attempts to download a release for a movie from a search.
// Pass the release for the item from the SearchRelease output for the release you wish to download.
// If the release.MovieID is 0 then release.MappedMovieID is used. Both may be 0, and that's OK unless
// release.ShouldOverride is true. If release.ShouldOverride is true, then Languages, MovieID and Quality
// must be present in the release.
func (r *Radarr) GrabReleaseContext(ctx context.Context, release *Release) (*Release, error) {
	grab := struct { // These are the required fields on the Radarr POST /release endpoint.
		GUID     string         `json:"guid"`
		Indexer  int64          `json:"indexerId"`
		Override bool           `json:"shouldOverride"`
		Language []*starr.Value `json:"languages,omitempty"`
		MovieID  int64          `json:"movieId,omitempty"`
		Quality  *starr.Quality `json:"quality,omitempty"`
	}{
		GUID:     release.GUID,
		Indexer:  release.IndexerID,
		Override: release.ShouldOverride,
		Language: release.Languages,
		MovieID:  release.MovieID,
		Quality:  release.Quality,
	}

	if grab.MovieID == 0 { // Best effort?
		grab.MovieID = release.MappedMovieID
	}

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(&grab); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpRelease, err)
	}

	var output Release

	req := starr.Request{URI: bpRelease, Body: &body}
	if err := r.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}
