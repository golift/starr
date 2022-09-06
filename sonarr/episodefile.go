package sonarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"golift.io/starr"
)

// GetEpisodeFiles returns information about episode files by episode file IDs.
func (s *Sonarr) GetEpisodeFiles(episodeFileIDs ...int64) ([]*EpisodeFile, error) {
	return s.GetEpisodeFilesContext(context.Background(), episodeFileIDs...)
}

// GetEpisodeFilesContext returns information about episode files by episode file IDs.
func (s *Sonarr) GetEpisodeFilesContext(ctx context.Context, episodeFileIDs ...int64) ([]*EpisodeFile, error) {
	var (
		output []*EpisodeFile
		ids    string
	)

	for _, efID := range episodeFileIDs {
		ids += strconv.FormatInt(efID, starr.Base10) + "," // the extra comma is ok.
	}

	params := make(url.Values)
	params.Add("episodeFileIds", ids)

	err := s.GetInto(ctx, "v3/episodeFile", params, &output)
	if err != nil {
		return nil, fmt.Errorf("api.Get(episodeFile): %w", err)
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

	params := make(url.Values)
	params.Add("seriesId", strconv.FormatInt(seriesID, starr.Base10))

	err := s.GetInto(ctx, "v3/episodeFile", params, &output)
	if err != nil {
		return nil, fmt.Errorf("api.Get(episodeFile): %w", err)
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
	if err := json.NewEncoder(&body).Encode(&EpisodeFile{
		ID:      episodeFileID,
		Quality: &starr.Quality{Quality: &starr.BaseQuality{ID: qualityID}},
	}); err != nil {
		return nil, fmt.Errorf("json.Marshal(episodeFileID): %w", err)
	}

	var output EpisodeFile

	err := s.PutInto(ctx, "v3/episodeFile/"+strconv.FormatInt(episodeFileID, starr.Base10), nil, &body, &output)
	if err != nil {
		return nil, fmt.Errorf("api.Put(episodeFile): %w", err)
	}

	return &output, nil
}

// DeleteEpisodeFile deletes an episode file.
func (s *Sonarr) DeleteEpisodeFile(episodeFileID int64) error {
	return s.DeleteEpisodeFileContext(context.Background(), episodeFileID)
}

// DeleteEpisodeFileContext deletes an episode file, and takes a context.
func (s *Sonarr) DeleteEpisodeFileContext(ctx context.Context, episodeFileID int64) error {
	req := &starr.Request{URI: "v3/episodeFile/" + fmt.Sprint(episodeFileID)}
	if err := s.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", req.URI, err)
	}

	return nil
}
