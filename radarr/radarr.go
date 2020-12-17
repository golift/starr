package radarr

// Radarr v3 structs

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// GetHistory returns the Radarr History (grabs/failures/completed).
func (r *Radarr) GetHistory() ([]*Record, error) {
	var history History

	params := make(url.Values)

	params.Set("sortKey", "date")
	params.Set("sortDir", "asc")
	params.Set("page", "1")
	params.Set("pageSize", "0")

	rawJSON, err := r.config.Req("v3/history", params)
	if err != nil {
		return nil, fmt.Errorf("c.Req(queue): %w", err)
	}

	if err = json.Unmarshal(rawJSON, &history); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(response): %w", err)
	}

	return history.Records, nil
}

// GetQueue returns the Radarr Queue (processing, but not yet imported).
func (r *Radarr) GetQueue() ([]*Queue, error) {
	var queue []*Queue

	params := make(url.Values)

	params.Set("sort_by", "timeleft")
	params.Set("order", "asc")

	rawJSON, err := r.config.Req("v3/queue", params)
	if err != nil {
		return nil, fmt.Errorf("c.Req(queue): %w", err)
	}

	if err = json.Unmarshal(rawJSON, &queue); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(response): %w", err)
	}

	return queue, nil
}

// GetMovie grabs a movie from the queue, or all movies if tmdbId is 0.
func (r *Radarr) GetMovie(tmdbID int) ([]*Movie, error) {
	var movie []*Movie

	params := make(url.Values)

	params.Set("tmdbId", strconv.Itoa(tmdbID))

	rawJSON, err := r.config.Req("v3/movie", params)
	if err != nil {
		return nil, fmt.Errorf("c.Req(movie): %w", err)
	}

	if err = json.Unmarshal(rawJSON, &movie); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(response): %w", err)
	}

	return movie, nil
}

// GetQualityProfiles returns all configured quality profiles.
func (r *Radarr) GetQualityProfiles() ([]*QualityProfile, error) {
	var profiles []*QualityProfile

	rawJSON, err := r.config.Req("v3/qualityProfile", nil)
	if err != nil {
		return nil, fmt.Errorf("c.Req(qualityProfile): %w", err)
	}

	if err = json.Unmarshal(rawJSON, &profiles); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(response): %w", err)
	}

	return profiles, nil
}

// RootFolders returns all configured root folders.
func (r *Radarr) GetRootFolders() ([]*RootFolder, error) {
	var folders []*RootFolder

	rawJSON, err := r.config.Req("v3/rootFolder", nil)
	if err != nil {
		return nil, fmt.Errorf("c.Req(rootFolder): %w", err)
	}

	if err = json.Unmarshal(rawJSON, &folders); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(response): %w", err)
	}

	return folders, nil
}

// AddMovie adds a movie to the queue.
func (r *Radarr) AddMovie(movie *AddMovieInput) (*AddMovieOutput, error) {
	body, err := json.Marshal(movie)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal(movie): %w", err)
	}

	params := make(url.Values)
	params.Add("moveFiles", "true")

	rawJSON, err := r.config.Req("v3/movie", params, body...)
	if err != nil {
		return nil, fmt.Errorf("c.Req(movie): %w", err)
	}

	var addedMovie *AddMovieOutput

	if err = json.Unmarshal(rawJSON, &addedMovie); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(response): %w", err)
	}

	return addedMovie, nil
}
