package sonarr

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

const bpRelease = APIver + "/release"

// Release is the output from the Sonarr release endpoint.
type Release struct {
	ID                           int64                 `json:"id"`
	GUID                         string                `json:"guid"`
	Quality                      starr.Quality         `json:"quality"`
	QualityWeight                int64                 `json:"qualityWeight"`
	Age                          int64                 `json:"age"`
	AgeHours                     int                   `json:"ageHours"`
	AgeMinutes                   int                   `json:"ageMinutes"`
	Size                         int                   `json:"size"`
	IndexerID                    int64                 `json:"indexerId"`
	Indexer                      string                `json:"indexer"`
	ReleaseGroup                 string                `json:"releaseGroup"`
	SubGroup                     string                `json:"subGroup"`
	ReleaseHash                  string                `json:"releaseHash"`
	Title                        string                `json:"title"`
	FullSeason                   bool                  `json:"fullSeason"`
	SceneSource                  bool                  `json:"sceneSource"`
	SeasonNumber                 int                   `json:"seasonNumber"`
	Languages                    []*starr.Value        `json:"languages"`
	LanguageWeight               int64                 `json:"languageWeight"`
	AirDate                      string                `json:"airDate"`
	SeriesTitle                  string                `json:"seriesTitle"`
	EpisodeNumbers               []int                 `json:"episodeNumbers"`
	AbsoluteEpisodeNumbers       []int                 `json:"absoluteEpisodeNumbers"`
	MappedSeasonNumber           int                   `json:"mappedSeasonNumber"`
	MappedEpisodeNumbers         []int                 `json:"mappedEpisodeNumbers"`
	MappedAbsoluteEpisodeNumbers []int                 `json:"mappedAbsoluteEpisodeNumbers"`
	MappedSeriesID               int64                 `json:"mappedSeriesId"`
	MappedEpisodeInfo            []*ReleaseEpisodeInfo `json:"mappedEpisodeInfo"`
	Approved                     bool                  `json:"approved"`
	TemporarilyRejected          bool                  `json:"temporarilyRejected"`
	Rejected                     bool                  `json:"rejected"`
	TvdbID                       int64                 `json:"tvdbId"`
	TvRageID                     int                   `json:"tvRageId"`
	Rejections                   []string              `json:"rejections"`
	PublishDate                  time.Time             `json:"publishDate"`
	CommentURL                   string                `json:"commentUrl"`
	DownloadURL                  string                `json:"downloadUrl"`
	InfoURL                      string                `json:"infoUrl"`
	EpisodeRequested             bool                  `json:"episodeRequested"`
	DownloadAllowed              bool                  `json:"downloadAllowed"`
	ReleaseWeight                int64                 `json:"releaseWeight"`
	CustomFormats                []*CustomFormatOutput `json:"customFormats"`
	CustomFormatScore            int64                 `json:"customFormatScore"`
	SceneMapping                 ReleaseSceneMapping   `json:"sceneMapping"`
	MagnetURL                    string                `json:"magnetUrl"`
	InfoHash                     string                `json:"infoHash"`
	Seeders                      int                   `json:"seeders"`
	Leechers                     int                   `json:"leechers"`
	Protocol                     starr.Protocol        `json:"protocol"`
	IndexerFlags                 int64                 `json:"indexerFlags"`
	IsDaily                      bool                  `json:"isDaily"`
	IsAbsoluteNumbering          bool                  `json:"isAbsoluteNumbering"`
	IsPossibleSpecialEpisode     bool                  `json:"isPossibleSpecialEpisode"`
	Special                      bool                  `json:"special"`
	SeriesID                     int64                 `json:"seriesId"`
	EpisodeID                    int64                 `json:"episodeId"`
	EpisodeIDs                   []int64               `json:"episodeIds"`
	DownloadClientID             int64                 `json:"downloadClientId"`
	DownloadClient               string                `json:"downloadClient"`
	ShouldOverride               bool                  `json:"shouldOverride"`
}

// ReleaseSceneMapping is part of a release.
type ReleaseSceneMapping struct {
	Title             string `json:"title"`
	SeasonNumber      int    `json:"seasonNumber"`
	SceneSeasonNumber int    `json:"sceneSeasonNumber"`
	SceneOrigin       string `json:"sceneOrigin"`
	Comment           string `json:"comment"`
}

// ReleaseEpisodeInfo is part of a release.
type ReleaseEpisodeInfo struct {
	ID                    int64  `json:"id"`
	SeasonNumber          int    `json:"seasonNumber"`
	EpisodeNumber         int    `json:"episodeNumber"`
	AbsoluteEpisodeNumber int    `json:"absoluteEpisodeNumber"`
	Title                 string `json:"title"`
}

// SearchRelease is the input needed to search for releases through Sonarr.
type SearchRelease struct {
	SeriesID     int64 `json:"seriesId"`
	EpisodeID    int64 `json:"episodeId"`
	SeasonNumber int   `json:"seasonNumber"`
}

// SearchRelease searches for and returns a list releases available for download.
func (s *Sonarr) SearchRelease(input *SearchRelease) ([]*Release, error) {
	return s.SearchReleaseContext(context.Background(), input)
}

// SearchReleaseContext searches for and returns a list releases available for download.
func (s *Sonarr) SearchReleaseContext(ctx context.Context, input *SearchRelease) ([]*Release, error) {
	req := starr.Request{URI: bpRelease, Query: make(url.Values)}
	req.Query.Set("seriesId", starr.Itoa(input.SeriesID))
	req.Query.Set("episodeId", starr.Itoa(input.EpisodeID))
	req.Query.Set("seasonNumber", strconv.Itoa(input.SeasonNumber))

	var output []*Release
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// Grab is the output from the GrabRelease method.
type Grab struct {
	GUID                     string         `json:"guid"`
	QualityWeight            int64          `json:"qualityWeight"`
	Age                      int64          `json:"age"`
	AgeHours                 int            `json:"ageHours"`
	AgeMinutes               int            `json:"ageMinutes"`
	Size                     int            `json:"size"`
	IndexerID                int64          `json:"indexerId"`
	FullSeason               bool           `json:"fullSeason"`
	SceneSource              bool           `json:"sceneSource"`
	SeasonNumber             int            `json:"seasonNumber"`
	LanguageWeight           int64          `json:"languageWeight"`
	Approved                 bool           `json:"approved"`
	TemporarilyRejected      bool           `json:"temporarilyRejected"`
	Rejected                 bool           `json:"rejected"`
	TvdbID                   int64          `json:"tvdbId"`
	TvRageID                 int64          `json:"tvRageId"`
	PublishDate              time.Time      `json:"publishDate"`
	EpisodeRequested         bool           `json:"episodeRequested"`
	DownloadAllowed          bool           `json:"downloadAllowed"`
	ReleaseWeight            int64          `json:"releaseWeight"`
	CustomFormatScore        int64          `json:"customFormatScore"`
	Protocol                 starr.Protocol `json:"protocol"`
	IndexerFlags             int64          `json:"indexerFlags"`
	IsDaily                  bool           `json:"isDaily"`
	IsAbsoluteNumbering      bool           `json:"isAbsoluteNumbering"`
	IsPossibleSpecialEpisode bool           `json:"isPossibleSpecialEpisode"`
	Special                  bool           `json:"special"`
}

// GrabRelease adds a release and attempts to download it.
// Pass the release for the item from the SearchRelease output.
func (s *Sonarr) GrabRelease(release *Release) (*Grab, error) {
	return s.GrabReleaseContext(context.Background(), release)
}

// GrabReleaseContext adds a release and attempts to download it.
// Pass the release for the item from the SearchRelease output.
func (s *Sonarr) GrabReleaseContext(ctx context.Context, release *Release) (*Grab, error) {
	var grab = struct { // We only use/need the guid and indexerID from the release.
		G string `json:"guid"`
		I int64  `json:"indexerId"`
	}{G: release.GUID, I: release.IndexerID}

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(&grab); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpRelease, err)
	}

	var output Grab

	req := starr.Request{URI: bpRelease, Body: &body}
	if err := s.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}
