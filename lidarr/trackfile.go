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
	Duration        *TrackDuration   `json:"duration"`
	Quality         *starr.Quality   `json:"quality"`
	MediaInfo       *AudioMediaInfo  `json:"mediaInfo"`
	TrackNumbers    []int            `json:"trackNumbers"`
	ReleaseGroup    string           `json:"releaseGroup"`
	ReleaseHash     string           `json:"releaseHash"`
}

// TrackDuration is part of AudioTags.
type TrackDuration struct {
	Ticks             int64 `json:"ticks"`
	Days              int64 `json:"days"`
	Hours             int64 `json:"hours"`
	Milliseconds      int64 `json:"milliseconds"`
	Minutes           int64 `json:"minutes"`
	Seconds           int64 `json:"seconds"`
	TotalDays         int64 `json:"totalDays"`
	TotalHours        int64 `json:"totalHours"`
	TotalMilliseconds int64 `json:"totalMilliseconds"`
	TotalMinutes      int64 `json:"totalMinutes"`
	TotalSeconds      int64 `json:"totalSeconds"`
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

	params := url.Values{"artistId": []string{fmt.Sprint(artistID)}}
	if err := l.GetInto(ctx, bpTrackFile, params, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", bpTrackFile, err)
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

	params := url.Values{"albumId": []string{fmt.Sprint(albumID)}}
	if err := l.GetInto(ctx, bpTrackFile, params, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", bpTrackFile, err)
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

	params := url.Values{"trackFileIds": make([]string, len(trackFileIDs))}
	for idx, fileID := range trackFileIDs {
		params["trackFileIds"][idx] = fmt.Sprint(fileID)
	}

	if err := l.GetInto(ctx, bpTrackFile, params, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", bpTrackFile, err)
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
		return nil, fmt.Errorf("json.Marshal(trackFile): %w", err)
	}

	uri := path.Join(bpTrackFile, fmt.Sprint(trackFile.ID))
	if err := l.PutInto(ctx, uri, nil, &body, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", uri, err)
	}

	return &output, nil
}

// DeleteTrackFile deletes a track file.
func (l *Lidarr) DeleteTrackFile(trackFileID int64) error {
	return l.DeleteTrackFileContext(context.Background(), trackFileID)
}

// DeleteTrackFileContext deletes a track file.
func (l *Lidarr) DeleteTrackFileContext(ctx context.Context, trackFileID int64) error {
	uri := path.Join(bpTrackFile, fmt.Sprint(trackFileID))
	if err := l.DeleteAny(ctx, uri, nil); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", uri, err)
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
	if err := json.NewEncoder(&body).Encode(postData); err != nil {
		return fmt.Errorf("json.Marshal(trackFileIDs): %w", err)
	}

	uri := path.Join(bpTrackFile, "bulk")
	if err := l.DeleteAny(ctx, uri, &starr.Params{Reader: &body}); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", uri, err)
	}

	return nil
}
