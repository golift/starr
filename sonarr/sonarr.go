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
	if err := s.GetInto("v3/system/status", nil, status); err != nil {
		return status, fmt.Errorf("api.Get(system/status): %w", err)
	}

	return status, nil
}

// GetLanguageProfiles returns all configured language profiles.
func (s *Sonarr) GetLanguageProfiles() ([]*LanguageProfile, error) {
	var profiles []*LanguageProfile
	if err := s.GetInto("v3/languageprofile", nil, &profiles); err != nil {
		return nil, fmt.Errorf("api.Get(languageprofile): %w", err)
	}

	return profiles, nil
}

// GetQualityProfiles returns all configured quality profiles.
func (s *Sonarr) GetQualityProfiles() ([]*QualityProfile, error) {
	var profiles []*QualityProfile
	if err := s.GetInto("v3/profile", nil, &profiles); err != nil {
		return nil, fmt.Errorf("api.Get(profile): %w", err)
	}

	return profiles, nil
}

// RootFolders returns all configured root folders.
func (s *Sonarr) GetRootFolders() ([]*RootFolder, error) {
	var folders []*RootFolder
	if err := s.GetInto("v3/rootfolder", nil, &folders); err != nil {
		return nil, fmt.Errorf("api.Get(rootfolder): %w", err)
	}

	return folders, nil
}

// GetSeriesLookup searches for a series using a search term or a tvdbid.
// Provide a search term or a tvdbid. If you provide both, tvdbID is used.
func (s *Sonarr) GetSeriesLookup(term string, tvdbID int) ([]*SeriesLookup, error) {
	params := make(url.Values)

	if tvdbID > 0 {
		params.Add("term", "tvdbid:"+strconv.Itoa(tvdbID))
	} else {
		params.Add("term", term)
	}

	var series []*SeriesLookup
	if err := s.GetInto("v3/series/lookup", params, &series); err != nil {
		return nil, fmt.Errorf("api.Get(series/lookup): %w", err)
	}

	return series, nil
}

// GetSeries locates and returns a series by tvdbID. If tvdbID is 0, returns all series.
func (s *Sonarr) GetSeries(tvdbID int) ([]*Series, error) {
	params := make(url.Values)

	if tvdbID != 0 {
		params.Add("tvdbId", strconv.Itoa(tvdbID))
	}

	var series []*Series
	if err := s.GetInto("v3/series", params, &series); err != nil {
		return nil, fmt.Errorf("api.Get(series): %w", err)
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

	added := &AddSeriesOutput{}
	if err = s.PostInto("v3/series", params, body, added); err != nil {
		return nil, fmt.Errorf("api.Post(series): %w", err)
	}

	return added, nil
}
