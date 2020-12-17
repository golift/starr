package sonarr

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// GetSystemStatus returns system status.
func (s *Sonarr) GetSystemStatus() (*SystemStatus, error) {
	var status *SystemStatus

	rawJSON, err := s.config.Req("v3/system/status", nil)
	if err != nil {
		return status, fmt.Errorf("c.Req(system/status): %w", err)
	}

	if err = json.Unmarshal(rawJSON, &status); err != nil {
		return status, fmt.Errorf("json.Unmarshal(response): %w", err)
	}

	return status, nil
}

// GetLanguageProfiles returns all configured language profiles.
func (s *Sonarr) GetLanguageProfiles() ([]*LanguageProfile, error) {
	var profiles []*LanguageProfile

	rawJSON, err := s.config.Req("v3/languageprofile", nil)
	if err != nil {
		return nil, fmt.Errorf("c.Req(languageprofile): %w", err)
	}

	if err = json.Unmarshal(rawJSON, &profiles); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(response): %w", err)
	}

	return profiles, nil
}

// GetQualityProfiles returns all configured quality profiles.
func (s *Sonarr) GetQualityProfiles() ([]*QualityProfile, error) {
	var profiles []*QualityProfile

	rawJSON, err := s.config.Req("v3/profile", nil)
	if err != nil {
		return nil, fmt.Errorf("c.Req(profile): %w", err)
	}

	if err = json.Unmarshal(rawJSON, &profiles); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(response): %w", err)
	}

	return profiles, nil
}

// RootFolders returns all configured root folders.
func (s *Sonarr) GetRootFolders() ([]*RootFolder, error) {
	var folders []*RootFolder

	rawJSON, err := s.config.Req("v3/rootfolder", nil)
	if err != nil {
		return nil, fmt.Errorf("c.Req(rootfolder): %w", err)
	}

	if err = json.Unmarshal(rawJSON, &folders); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(response): %w", err)
	}

	return folders, nil
}

// GetSeriesLookup searches for a series using a search term or a tvdbid.
// Provide a search term or a tvdbid. If you provide both, tvdbID is used.
func (s *Sonarr) GetSeriesLookup(term string, tvdbID int) ([]*SeriesLookup, error) {
	var series []*SeriesLookup

	params := make(url.Values)

	if tvdbID > 0 {
		params.Add("term", "tvdbid:"+strconv.Itoa(tvdbID))
	} else {
		params.Add("term", term)
	}

	rawJSON, err := s.config.Req("v3/series/lookup", params)
	if err != nil {
		return nil, fmt.Errorf("c.Req(series/lookup): %w", err)
	}

	if err = json.Unmarshal(rawJSON, &series); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(response): %w", err)
	}

	return series, nil
}

// GetSeries locates and returns a series by tvdbID. If tvdbID is 0, returns all series.
func (s *Sonarr) GetSeries(tvdbID int) ([]*Series, error) {
	var series []*Series

	params := make(url.Values)

	if tvdbID != 0 {
		params.Add("tvdbId", strconv.Itoa(tvdbID))
	}

	rawJSON, err := s.config.Req("v3/series", params)
	if err != nil {
		return nil, fmt.Errorf("c.Req(series): %w", err)
	}

	if err = json.Unmarshal(rawJSON, &series); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(response): %w", err)
	}

	return series, nil
}

// GetAllSeries returns all configured series.
// This may not deal well with pagination atm.
func (s *Sonarr) GetAllSeries() ([]*Series, error) {
	return s.GetSeries(0)
}

// AddSeries adds a new series to Sonarr.
func (s *Sonarr) AddSeries(series *AddSeriesInput) (*AddSeriesOutput, error) {
	body, err := json.Marshal(series)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal(series): %w", err)
	}

	params := make(url.Values)
	params.Add("moveFiles", "true")

	rawJSON, err := s.config.Req("v3/series", params, body...)
	if err != nil {
		return nil, fmt.Errorf("c.Req(series): %w", err)
	}

	var added *AddSeriesOutput

	if err = json.Unmarshal(rawJSON, &added); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(response): %w", err)
	}

	return added, nil
}
