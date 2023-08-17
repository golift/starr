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

const bpTrackFile = APIver + "/trackfile"

// TrackFile represents the data sent to and returned from the trackfile endpoint.
type TrackFile struct {
	ID            int64          `json:"id"`
	ArtistID      int64          `json:"artistId"`
	AlbumID       int64          `json:"albumId"`
	Path          string         `json:"path"`
	Size          int64          `json:"size"`
	DateAdded     time.Time      `json:"dateAdded"`
	Quality       *starr.Quality `json:"quality"`
	QualityWeight int            `json:"qualityWeight"`
	MediaInfo     MediaInfo      `json:"mediaInfo"`
	CutoffNotMet  bool           `json:"qualityCutoffNotMet"`
	AudioTags     *AudioTags     `json:"audioTags"`
}

// MediaInfo is part of a TrackFile.
type MediaInfo struct {
	ID              int64  `json:"id"`
	AudioChannels   int    `json:"audioChannels"`
	AudioBitRate    string `json:"audioBitRate"`
	AudioCodec      string `json:"audioCodec"`
	AudioBits       string `json:"audioBits"`
	AudioSampleRate string `json:"audioSampleRate"`
}

/* I've never seen audio tags in the wild. */

// AudioTags is (optionally) part of a TrackFile.
type AudioTags struct {
	Title           string           `json:"title"`
	CleanTitle      string           `json:"cleanTitle"`
	ArtistTitle     string           `json:"artistTitle"`
	AlbumTitle      string           `json:"albumTitle"`
	ArtistTitleInfo *ArtistTitleInfo `json:"artistTitleInfo"`
	ArtistMBID      string           `json:"artistMBId"`
	AlbumMBID       string           `json:"albumMBId"`
	ReleaseMBID     string           `json:"releaseMBId"`
	RecordingMBID   string           `json:"recordingMBId"`
	TrackMBID       string           `json:"trackMBId"`
	DiscNumber      int              `json:"discNumber"`
	DiscCount       int              `json:"discCount"`
	Country         *AudioCountry    `json:"country"`
	Year            int              `json:"year"`
	Label           string           `json:"label"`
	CatalogNumber   string           `json:"catalogNumber"`
	Disambiguation  string           `json:"disambiguation"`
	Duration        *starr.TimeSpan  `json:"duration"`
	Quality         *starr.Quality   `json:"quality"`
	MediaInfo       *AudioMediaInfo  `json:"mediaInfo"`
	TrackNumbers    []int            `json:"trackNumbers"`
	ReleaseGroup    string           `json:"releaseGroup"`
	ReleaseHash     string           `json:"releaseHash"`
}

// AudioMediaInfo is part of AudioTags.
type AudioMediaInfo struct {
	AudioFormat     string `json:"audioFormat"`
	AudioBitrate    int64  `json:"audioBitrate"`
	AudioChannels   int    `json:"audioChannels"`
	AudioBits       int    `json:"audioBits"`
	AudioSampleRate int    `json:"audioSampleRate"`
}

// AudioCountry is part of AudioTags.
type AudioCountry struct {
	TwoLetterCode string `json:"twoLetterCode"`
	Name          string `json:"name"`
}

// ArtistTitleInfo is part of AudioTags.
type ArtistTitleInfo struct {
	Title            string `json:"title"`
	TitleWithoutYear string `json:"titleWithoutYear"`
	Year             int    `json:"year"`
}

// GetTrackFilesForArtist returns the track files for an artist.
func (l *Lidarr) GetTrackFilesForArtist(artistID int64) ([]*TrackFile, error) {
	return l.GetTrackFilesForArtistContext(context.Background(), artistID)
}

// GetTrackFilesForArtistContext returns the track files for an artist.
func (l *Lidarr) GetTrackFilesForArtistContext(ctx context.Context, artistID int64) ([]*TrackFile, error) {
	var output []*TrackFile

	req := starr.Request{URI: bpTrackFile, Query: make(url.Values)}
	req.Query.Add("artistId", fmt.Sprint(artistID))

	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetTrackFilesForAlbum returns the track files for an album.
func (l *Lidarr) GetTrackFilesForAlbum(albumID int64) ([]*TrackFile, error) {
	return l.GetTrackFilesForAlbumContext(context.Background(), albumID)
}

// GetTrackFilesForAlbumContext returns the track files for an album.
func (l *Lidarr) GetTrackFilesForAlbumContext(ctx context.Context, albumID int64) ([]*TrackFile, error) {
	var output []*TrackFile

	req := starr.Request{URI: bpTrackFile, Query: make(url.Values)}
	req.Query.Add("albumId", fmt.Sprint(albumID))

	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetTrackFiles returns the requested track files by ID.
func (l *Lidarr) GetTrackFiles(trackFileIDs []int64) ([]*TrackFile, error) {
	return l.GetTrackFilesContext(context.Background(), trackFileIDs)
}

// GetTrackFilesContext returns the requested track files by their IDs.
func (l *Lidarr) GetTrackFilesContext(ctx context.Context, trackFileIDs []int64) ([]*TrackFile, error) {
	var output []*TrackFile

	if len(trackFileIDs) == 0 {
		return output, nil
	}

	req := starr.Request{
		URI:   bpTrackFile,
		Query: url.Values{"trackFileIds": make([]string, len(trackFileIDs))},
	}

	for idx, fileID := range trackFileIDs {
		req.Query["trackFileIds"][idx] = fmt.Sprint(fileID)
	}

	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// UpdateTrackFile updates a track file.
func (l *Lidarr) UpdateTrackFile(trackFile *TrackFile) (*TrackFile, error) {
	return l.UpdateTrackFileContext(context.Background(), trackFile)
}

// UpdateTrackFileContext updates a track file.
func (l *Lidarr) UpdateTrackFileContext(ctx context.Context, trackFile *TrackFile) (*TrackFile, error) {
	var output TrackFile

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(trackFile); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpTrackFile, err)
	}

	req := starr.Request{URI: path.Join(bpTrackFile, fmt.Sprint(trackFile.ID)), Body: &body}
	if err := l.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteTrackFile deletes a track file.
func (l *Lidarr) DeleteTrackFile(trackFileID int64) error {
	return l.DeleteTrackFileContext(context.Background(), trackFileID)
}

// DeleteTrackFileContext deletes a track file.
func (l *Lidarr) DeleteTrackFileContext(ctx context.Context, trackFileID int64) error {
	req := starr.Request{URI: path.Join(bpTrackFile, fmt.Sprint(trackFileID))}
	if err := l.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}

// DeleteTrackFiles bulk deletes track files by their IDs.
func (l *Lidarr) DeleteTrackFiles(trackFileIDs []int64) error {
	return l.DeleteTrackFilesContext(context.Background(), trackFileIDs)
}

// DeleteTrackFilesContext bulk deletes track files by their IDs.
func (l *Lidarr) DeleteTrackFilesContext(ctx context.Context, trackFileIDs []int64) error {
	postData := struct {
		T []int64 `json:"trackFileIDs"`
	}{trackFileIDs}

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(&postData); err != nil {
		return fmt.Errorf("json.Marshal(%s): %w", bpTrackFile, err)
	}

	req := starr.Request{URI: path.Join(bpTrackFile, "bulk"), Body: &body}
	if err := l.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
