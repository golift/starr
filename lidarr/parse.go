package lidarr

import (
	"context"
	"fmt"
	"net/url"

	"golift.io/starr"
)

const bpParse = APIver + "/parse"

// ParsedAlbumInfo is returned when a release title parses as an album.
type ParsedAlbumInfo struct {
	ReleaseTitle     string         `json:"releaseTitle,omitempty"`
	AlbumTitle       string         `json:"albumTitle,omitempty"`
	ArtistName       string         `json:"artistName,omitempty"`
	AlbumType        string         `json:"albumType,omitempty"`
	Quality          *starr.Quality `json:"quality,omitempty"`
	ReleaseDate      string         `json:"releaseDate,omitempty"`
	Discography      bool           `json:"discography,omitempty"`
	DiscographyStart int            `json:"discographyStart,omitempty"`
}

// ParseOutput is returned from GET /api/v1/parse.
type ParseOutput struct {
	ID                int64                 `json:"id"`
	Title             string                `json:"title,omitempty"`
	ParsedAlbumInfo   *ParsedAlbumInfo      `json:"parsedAlbumInfo,omitempty"`
	Artist            *Artist               `json:"artist,omitempty"`
	Albums            []*Album              `json:"albums,omitempty"`
	CustomFormats     []*CustomFormatOutput `json:"customFormats,omitempty"`
	CustomFormatScore int64                 `json:"customFormatScore"`
}

// Parse resolves a release title into parsed album metadata.
func (l *Lidarr) Parse(title string) (*ParseOutput, error) {
	return l.ParseContext(context.Background(), title)
}

// ParseContext resolves a release title into parsed album metadata.
func (l *Lidarr) ParseContext(ctx context.Context, title string) (*ParseOutput, error) {
	var output *ParseOutput

	req := starr.Request{URI: bpParse, Query: make(url.Values)}
	req.Query.Set("title", title)

	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}
