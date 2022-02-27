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

// GetAllSeries returns all configured series.
// This may not deal well with pagination atm.
func (s *Sonarr) GetAllSeries() ([]*Series, error) {
	return s.GetAllSeriesContext(context.Background())
}

func (s *Sonarr) GetAllSeriesContext(ctx context.Context) ([]*Series, error) {
	return s.GetSeriesContext(ctx, 0)
}

// GetSeries locates and returns a series by tvdbID. If tvdbID is 0, returns all series.
func (s *Sonarr) GetSeries(tvdbID int64) ([]*Series, error) {
	return s.GetSeriesContext(context.Background(), tvdbID)
}

func (s *Sonarr) GetSeriesContext(ctx context.Context, tvdbID int64) ([]*Series, error) {
	params := make(url.Values)

	if tvdbID != 0 {
		params.Add("tvdbId", strconv.FormatInt(tvdbID, starr.Base10))
	}

	var series []*Series

	_, err := s.GetInto(ctx, "v3/series", params, &series)
	if err != nil {
		return nil, fmt.Errorf("api.Get(series): %w", err)
	}

	return series, nil
}

// UpdateSeries updates a series in place.
func (s *Sonarr) UpdateSeries(seriesID int64, series *Series) error {
	return s.UpdateSeriesContext(context.Background(), seriesID, series)
}

func (s *Sonarr) UpdateSeriesContext(ctx context.Context, seriesID int64, series *Series) error {
	params := make(url.Values)
	params.Add("moveFiles", "true")

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(series); err != nil {
		return fmt.Errorf("json.Marshal(series): %w", err)
	}

	_, err := s.Put(ctx, "v3/series/"+strconv.FormatInt(seriesID, starr.Base10), params, &body)
	if err != nil {
		return fmt.Errorf("api.Put(series): %w", err)
	}

	return nil
}

// AddSeries adds a new series to Sonarr.
func (s *Sonarr) AddSeries(series *AddSeriesInput) (*AddSeriesOutput, error) {
	return s.AddSeriesContext(context.Background(), series)
}

func (s *Sonarr) AddSeriesContext(ctx context.Context, series *AddSeriesInput) (*AddSeriesOutput, error) {
	params := make(url.Values)
	params.Add("moveFiles", "true")

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(series); err != nil {
		return nil, fmt.Errorf("json.Marshal(series): %w", err)
	}

	var output AddSeriesOutput
	if _, err := s.PostInto(ctx, "v3/series", params, &body, &output); err != nil {
		return nil, fmt.Errorf("api.Post(series): %w", err)
	}

	return &output, nil
}

// GetSeriesByID locates and returns a series by DB [series] ID.
func (s *Sonarr) GetSeriesByID(seriesID int64) (*Series, error) {
	return s.GetSeriesByIDContext(context.Background(), seriesID)
}

func (s *Sonarr) GetSeriesByIDContext(ctx context.Context, seriesID int64) (*Series, error) {
	var series Series

	_, err := s.GetInto(ctx, "v3/series/"+strconv.FormatInt(seriesID, starr.Base10), nil, &series)
	if err != nil {
		return nil, fmt.Errorf("api.Get(series): %w", err)
	}

	return &series, nil
}

// GetSeriesLookup searches for a series [in Servarr] using a search term or a tvdbid.
// Provide a search term or a tvdbid. If you provide both, tvdbID is used.
func (s *Sonarr) GetSeriesLookup(term string, tvdbID int64) ([]*Series, error) {
	return s.GetSeriesLookupContext(context.Background(), term, tvdbID)
}

func (s *Sonarr) GetSeriesLookupContext(ctx context.Context, term string, tvdbID int64) ([]*Series, error) {
	params := make(url.Values)

	if tvdbID > 0 {
		params.Add("term", "tvdbid:"+strconv.FormatInt(tvdbID, starr.Base10))
	} else {
		params.Add("term", term)
	}

	var series []*Series

	_, err := s.GetInto(ctx, "v3/series/lookup", params, &series)
	if err != nil {
		return nil, fmt.Errorf("api.Get(series/lookup): %w", err)
	}

	return series, nil
}

// Lookup will search for series matching the specified search term.
// Searches for new shows on TheTVDB.com utilizing sonarr.tv's caching and augmentation proxy.
func (s *Sonarr) Lookup(term string) ([]*Series, error) {
	return s.LookupContext(context.Background(), term)
}

func (s *Sonarr) LookupContext(ctx context.Context, term string) ([]*Series, error) {
	return s.GetSeriesLookupContext(ctx, term, 0)
}
