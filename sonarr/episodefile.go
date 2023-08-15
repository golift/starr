package sonarr

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

const bpEpisodeFile = APIver + "/episodeFile"

// EpisodeFile is the output from the /api/v3/episodeFile endpoint.
type EpisodeFile struct {
	ID                   int64                 `json:"id"`
	SeriesID             int64                 `json:"seriesId"`
	SeasonNumber         int                   `json:"seasonNumber"`
	RelativePath         string                `json:"relativePath"`
	Path                 string                `json:"path"`
	Size                 int64                 `json:"size"`
	DateAdded            time.Time             `json:"dateAdded"`
	SceneName            string                `json:"sceneName"`
	ReleaseGroup         string                `json:"releaseGroup"`
	Language             *starr.Value          `json:"language"`
	Quality              *starr.Quality        `json:"quality"`
	MediaInfo            *MediaInfo            `json:"mediaInfo"`
	QualityCutoffNotMet  bool                  `json:"qualityCutoffNotMet"`
	LanguageCutoffNotMet bool                  `json:"languageCutoffNotMet"`
	CustomFormats        []*CustomFormatOutput `json:"customFormats"` // v4 only
}

// MediaInfo is part of an EpisodeFile.
type MediaInfo struct {
	AudioBitrate     int            `json:"audioBitrate"`
	AudioChannels    float64        `json:"audioChannels"`
	AudioCodec       string         `json:"audioCodec"`
	AudioLanguages   string         `json:"audioLanguages"`
	AudioStreamCount int            `json:"audioStreamCount"`
	VideoBitDepth    int            `json:"videoBitDepth"`
	VideoBitrate     int            `json:"videoBitrate"`
	VideoCodec       string         `json:"videoCodec"`
	VideoFPS         float64        `json:"videoFps"`
	Resolution       string         `json:"resolution"`
	RunTime          starr.PlayTime `json:"runTime"`
	ScanType         string         `json:"scanType"`
	Subtitles        string         `json:"subtitles"`
}

// GetEpisodeFiles returns information about episode files by episode file IDs.
func (s *Sonarr) GetEpisodeFiles(episodeFileIDs ...int64) ([]*EpisodeFile, error) {
	return s.GetEpisodeFilesContext(context.Background(), episodeFileIDs...)
}

// GetEpisodeFilesContext returns information about episode files by episode file IDs.
func (s *Sonarr) GetEpisodeFilesContext(ctx context.Context, episodeFileIDs ...int64) ([]*EpisodeFile, error) {
	var ids string
	for _, efID := range episodeFileIDs {
		ids += fmt.Sprintf("%d,", efID) // the extra comma is ok.
	}

	req := starr.Request{URI: bpEpisodeFile, Query: make(url.Values)}
	req.Query.Add("episodeFileIds", ids)

	var output []*EpisodeFile
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetSeriesEpisodeFile returns information about all episode files in a series.
func (s *Sonarr) GetSeriesEpisodeFiles(seriesID int64) ([]*EpisodeFile, error) {
	return s.GetSeriesEpisodeFilesContext(context.Background(), seriesID)
}

// GetSeriesEpisodeFilesContext returns information about episode files by episode file ID.
func (s *Sonarr) GetSeriesEpisodeFilesContext(ctx context.Context, seriesID int64) ([]*EpisodeFile, error) {
	var output []*EpisodeFile

	req := starr.Request{URI: bpEpisodeFile, Query: make(url.Values)}
	req.Query.Add("seriesId", fmt.Sprint(seriesID))

	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// UpdateEpisodeFile updates an episode file's quality. Use GetQualityProfiles() to find the available IDs.
func (s *Sonarr) UpdateEpisodeFileQuality(episodeFileID, qualityID int64) (*EpisodeFile, error) {
	return s.UpdateEpisodeFileQualityContext(context.Background(), episodeFileID, qualityID)
}

// UpdateEpisodeFileQualityContext updates an episode file, and takes a context.
func (s *Sonarr) UpdateEpisodeFileQualityContext(
	ctx context.Context,
	episodeFileID int64,
	qualityID int64,
) (*EpisodeFile, error) {
	var body bytes.Buffer

	err := json.NewEncoder(&body).Encode(&EpisodeFile{
		ID:      episodeFileID,
		Quality: &starr.Quality{Quality: &starr.BaseQuality{ID: qualityID}},
	})
	if err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpEpisodeFile, err)
	}

	var output EpisodeFile

	req := starr.Request{URI: path.Join(bpEpisodeFile, fmt.Sprint(episodeFileID)), Body: &body}
	if err := s.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteEpisodeFile deletes an episode file.
func (s *Sonarr) DeleteEpisodeFile(episodeFileID int64) error {
	return s.DeleteEpisodeFileContext(context.Background(), episodeFileID)
}

// DeleteEpisodeFileContext deletes an episode file, and takes a context.
func (s *Sonarr) DeleteEpisodeFileContext(ctx context.Context, episodeFileID int64) error {
	req := starr.Request{URI: path.Join(bpEpisodeFile, fmt.Sprint(episodeFileID))}
	if err := s.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
