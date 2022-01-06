package sonarr

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"

	"golift.io/starr"
)

// GetAllSeries returns all configured series.
// This may not deal well with pagination atm.
func (s *Sonarr) GetAllSeries() ([]*Series, error) {
	return s.GetSeries(0)
}

// GetSeries locates and returns a series by tvdbID. If tvdbID is 0, returns all series.
func (s *Sonarr) GetSeries(tvdbID int64) ([]*Series, error) {
	params := make(url.Values)

	if tvdbID != 0 {
		params.Add("tvdbId", strconv.FormatInt(tvdbID, starr.Base10))
	}

	var series []*Series

	err := s.GetInto("v3/series", params, &series)
	if err != nil {
		return nil, fmt.Errorf("api.Get(series): %w", err)
	}

	return series, nil
}

// UpdateSeries updates a series in place.
func (s *Sonarr) UpdateSeries(seriesID int64, series *Series) error {
	put, err := json.Marshal(series)
	if err != nil {
		return fmt.Errorf("json.Marshal(series): %w", err)
	}

	params := make(url.Values)
	params.Add("moveFiles", "true")

	b, err := s.Put("v3/series/"+strconv.FormatInt(seriesID, starr.Base10), params, put)
	if err != nil {
		return fmt.Errorf("api.Put(series): %w", err)
	}

	log.Println("SHOW THIS TO CAPTAIN plz:", string(b))

	return nil
}

// AddSeries adds a new series to Sonarr.
func (s *Sonarr) AddSeries(series *AddSeriesInput) (*AddSeriesOutput, error) {
	body, err := json.Marshal(series)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal(series): %w", err)
	}

	params := make(url.Values)
	params.Add("moveFiles", "true")

	var output AddSeriesOutput
	if err = s.PostInto("v3/series", params, body, &output); err != nil {
		return nil, fmt.Errorf("api.Post(series): %w", err)
	}

	return &output, nil
}

// GetSeriesByID locates and returns a series by DB [series] ID.
func (s *Sonarr) GetSeriesByID(seriesID int64) (*Series, error) {
	var series Series

	err := s.GetInto("v3/series/"+strconv.FormatInt(seriesID, starr.Base10), nil, &series)
	if err != nil {
		return nil, fmt.Errorf("api.Get(series): %w", err)
	}

	return &series, nil
}

// GetSeriesLookup searches for a series [in Servarr] using a search term or a tvdbid.
// Provide a search term or a tvdbid. If you provide both, tvdbID is used.
func (s *Sonarr) GetSeriesLookup(term string, tvdbID int64) ([]*Series, error) {
	params := make(url.Values)

	if tvdbID > 0 {
		params.Add("term", "tvdbid:"+strconv.FormatInt(tvdbID, starr.Base10))
	} else {
		params.Add("term", term)
	}

	var series []*Series

	err := s.GetInto("v3/series/lookup", params, &series)
	if err != nil {
		return nil, fmt.Errorf("api.Get(series/lookup): %w", err)
	}

	return series, nil
}

// Lookup will search for series matching the specified search term.
// Searches for new shows on TheTVDB.com utilizing sonarr.tv's caching and augmentation proxy.
func (s *Sonarr) Lookup(term string) ([]*Series, error) {
	return s.GetSeriesLookup(term, 0)
}
