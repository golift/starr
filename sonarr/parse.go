package sonarr

import (
	"context"
	"fmt"
	"net/url"

	"golift.io/starr"
)

const bpParse = APIver + "/parse"

// ParseInput is the input for the Sonarr parse endpoint. Must provide either Title or Path.
type ParseInput struct {
	Title string
	Path  string
}

// SeriesTitleInfo has only been seen in the parse endpoint so far.
type SeriesTitleInfo struct {
	Year             int    `json:"year"`
	Title            string `json:"title"`
	TitleWithoutYear string `json:"titleWithoutYear"`
}

// ParsedEpisodeInfo is provided in ParseOutput when an item was properly parsed.
type ParsedEpisodeInfo struct {
	EpisodeNumbers                []int            `json:"episodeNumbers"`
	AbsoluteEpisodeNumbers        []int            `json:"absoluteEpisodeNumbers"`
	SpecialAbsoluteEpisodeNumbers []interface{}    `json:"specialAbsoluteEpisodeNumbers"`
	Languages                     []*starr.Value   `json:"languages"`
	SeasonNumber                  int              `json:"seasonNumber"`
	SeasonPart                    int64            `json:"seasonPart"`
	FullSeason                    bool             `json:"fullSeason"`
	IsPartialSeason               bool             `json:"isPartialSeason"`
	IsMultiSeason                 bool             `json:"isMultiSeason"`
	IsSeasonExtra                 bool             `json:"isSeasonExtra"`
	IsSplitEpisode                bool             `json:"isSplitEpisode"`
	IsMiniSeries                  bool             `json:"isMiniSeries"`
	Special                       bool             `json:"special"`
	IsDaily                       bool             `json:"isDaily"`
	IsAbsoluteNumbering           bool             `json:"isAbsoluteNumbering"`
	IsPossibleSpecialEpisode      bool             `json:"isPossibleSpecialEpisode"`
	IsPossibleSceneSeasonSpecial  bool             `json:"isPossibleSceneSeasonSpecial"`
	ReleaseTitle                  string           `json:"releaseTitle"`
	SeriesTitle                   string           `json:"seriesTitle"`
	ReleaseGroup                  string           `json:"releaseGroup"`
	ReleaseHash                   string           `json:"releaseHash"`
	ReleaseTokens                 string           `json:"releaseTokens"`
	ReleaseType                   string           `json:"releaseType"`
	SeriesTitleInfo               *SeriesTitleInfo `json:"seriesTitleInfo"`
	Quality                       *starr.Quality   `json:"quality"`
}

// ParseOutput is what you get from the parse endpoint when you provide a parsable path or title.
type ParseOutput struct {
	Episodes          []*Episode            `json:"episodes"`
	Languages         []*starr.Value        `json:"languages"`
	CustomFormats     []*CustomFormatOutput `json:"customFormats"`
	CustomFormatScore int64                 `json:"customFormatScore"`
	ID                int64                 `json:"id"`
	Title             string                `json:"title"`
	// You need to check this for nil before accessing it.
	// If the parse failed, this won't exist, and you won't get an error.
	ParsedEpisodeInfo *ParsedEpisodeInfo `json:"parsedEpisodeInfo"`
}

// Parse a title or path into episode info.
func (s *Sonarr) Parse(input *ParseInput) (*ParseOutput, error) {
	return s.ParseContext(context.Background(), input)
}

// ParseContext parses a title or path into episode info.
func (s *Sonarr) ParseContext(ctx context.Context, input *ParseInput) (*ParseOutput, error) {
	var output *ParseOutput

	req := starr.Request{URI: bpParse, Query: make(url.Values)}
	req.Query.Set("title", input.Title)
	req.Query.Set("path", input.Path)

	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}
