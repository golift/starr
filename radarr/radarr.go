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
	params := make(url.Values)
	params.Set("sortKey", "date")
	params.Set("sortDir", "asc")
	params.Set("page", "1")
	params.Set("pageSize", "0")

	var history History
	if err := r.GetInto("v3/history", params, &history); err != nil {
		return nil, fmt.Errorf("api.Get(history): %w", err)
	}

	return history.Records, nil
}

// GetQueue returns the Radarr Queue (processing, but not yet imported).
func (r *Radarr) GetQueue() ([]*Queue, error) {
	params := make(url.Values)
	params.Set("sort_by", "timeleft")
	params.Set("order", "asc")

	var queue []*Queue
	if err := r.GetInto("v3/queue", params, &queue); err != nil {
		return nil, fmt.Errorf("api.Get(queue): %w", err)
	}

	return queue, nil
}

// GetMovie grabs a movie from the queue, or all movies if tmdbId is 0.
func (r *Radarr) GetMovie(tmdbID int) ([]*Movie, error) {
	params := make(url.Values)
	params.Set("tmdbId", strconv.Itoa(tmdbID))

	var movie []*Movie
	if err := r.GetInto("v3/movie", params, &movie); err != nil {
		return nil, fmt.Errorf("api.Get(movie): %w", err)
	}

	return movie, nil
}

// GetQualityProfiles returns all configured quality profiles.
func (r *Radarr) GetQualityProfiles() ([]*QualityProfile, error) {
	var profiles []*QualityProfile
	if err := r.GetInto("v3/qualityProfile", nil, &profiles); err != nil {
		return nil, fmt.Errorf("api.Get(qualityProfile): %w", err)
	}

	return profiles, nil
}

// RootFolders returns all configured root folders.
func (r *Radarr) GetRootFolders() ([]*RootFolder, error) {
	var folders []*RootFolder
	if err := r.GetInto("v3/rootFolder", nil, &folders); err != nil {
		return nil, fmt.Errorf("api.Get(rootFolder): %w", err)
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

	var added *AddMovieOutput
	if err := r.PostInto("v3/movie", params, body, added); err != nil {
		return nil, fmt.Errorf("api.Post(movie): %w", err)
	}

	return added, nil
}
