package sonarr

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"
)

// GetSystemStatus returns system status.
func (s *Sonarr) GetSystemStatus() (*SystemStatus, error) {
	var status SystemStatus

	err := s.GetInto("v3/system/status", nil, &status)
	if err != nil {
		return nil, fmt.Errorf("api.Get(system/status): %w", err)
	}

	return &status, nil
}

// GetLanguageProfiles returns all configured language profiles.
func (s *Sonarr) GetLanguageProfiles() ([]*LanguageProfile, error) {
	var profiles []*LanguageProfile

	err := s.GetInto("v3/languageprofile", nil, &profiles)
	if err != nil {
		return nil, fmt.Errorf("api.Get(languageprofile): %w", err)
	}

	return profiles, nil
}

// GetQualityProfiles returns all configured quality profiles.
func (s *Sonarr) GetQualityProfiles() ([]*QualityProfile, error) {
	var profiles []*QualityProfile

	err := s.GetInto("v3/qualityprofile", nil, &profiles)
	if err != nil {
		return nil, fmt.Errorf("api.Get(qualityprofile): %w", err)
	}

	return profiles, nil
}

// GetRootFolders returns all configured root folders.
func (s *Sonarr) GetRootFolders() ([]*RootFolder, error) {
	var folders []*RootFolder

	err := s.GetInto("v3/rootfolder", nil, &folders)
	if err != nil {
		return nil, fmt.Errorf("api.Get(rootfolder): %w", err)
	}

	return folders, nil
}

// GetSeriesLookup searches for a series [in Servarr] using a search term or a tvdbid.
// Provide a search term or a tvdbid. If you provide both, tvdbID is used.
func (s *Sonarr) GetSeriesLookup(term string, tvdbID int64) ([]*SeriesLookup, error) {
	params := make(url.Values)

	if tvdbID > 0 {
		params.Add("term", "tvdbid:"+strconv.FormatInt(tvdbID, 10))
	} else {
		params.Add("term", term)
	}

	var series []*SeriesLookup

	err := s.GetInto("v3/series/lookup", params, &series)
	if err != nil {
		return nil, fmt.Errorf("api.Get(series/lookup): %w", err)
	}

	return series, nil
}

// GetSeries locates and returns a series by tvdbID. If tvdbID is 0, returns all series.
func (s *Sonarr) GetSeries(tvdbID int64) ([]*Series, error) {
	params := make(url.Values)

	if tvdbID != 0 {
		params.Add("tvdbId", strconv.FormatInt(tvdbID, 10))
	}

	var series []*Series

	err := s.GetInto("v3/series", params, &series)
	if err != nil {
		return nil, fmt.Errorf("api.Get(series): %w", err)
	}

	return series, nil
}

// GetSeriesByID locates and returns a series by DB [series] ID.
func (s *Sonarr) GetSeriesByID(seriesID int64) (*Series, error) {
	var series Series

	err := s.GetInto("v3/series/"+strconv.FormatInt(seriesID, 10), nil, &series)
	if err != nil {
		return nil, fmt.Errorf("api.Get(series): %w", err)
	}

	return &series, nil
}

// GetAllSeries returns all configured series.
// This may not deal well with pagination atm.
func (s *Sonarr) GetAllSeries() ([]*Series, error) {
	return s.GetSeries(0)
}

// UpdateSeries updates a series to in place.
func (s *Sonarr) UpdateSeries(seriesID int64, series *Series) error {
	put, err := json.Marshal(series)
	if err != nil {
		return fmt.Errorf("json.Marshal(series): %w", err)
	}

	params := make(url.Values)
	params.Add("moveFiles", "true")

	b, err := s.Put("v3/series/"+strconv.FormatInt(seriesID, 10), params, put)
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
