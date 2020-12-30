package radarr

// Radarr v3 structs

import (
	"encoding/json"
	"fmt"
	"log"
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

	err := r.GetInto("v3/history", params, &history)
	if err != nil {
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

	err := r.GetInto("v3/queue", params, &queue)
	if err != nil {
		return nil, fmt.Errorf("api.Get(queue): %w", err)
	}

	return queue, nil
}

// GetMovie grabs a movie from the queue, or all movies if tmdbId is 0.
func (r *Radarr) GetMovie(tmdbID int64) ([]*Movie, error) {
	params := make(url.Values)
	params.Set("tmdbId", strconv.FormatInt(tmdbID, 10))

	var movie []*Movie

	err := r.GetInto("v3/movie", params, &movie)
	if err != nil {
		return nil, fmt.Errorf("api.Get(movie): %w", err)
	}

	return movie, nil
}

// GetMovieByID grabs a movie from the database by DB [movie] ID.
func (r *Radarr) GetMovieByID(movieID int64) (*Movie, error) {
	var movie Movie

	err := r.GetInto("v3/movie/"+strconv.FormatInt(movieID, 10), nil, &movie)
	if err != nil {
		return nil, fmt.Errorf("api.Get(movie): %w", err)
	}

	return &movie, nil
}

// GetQualityProfiles returns all configured quality profiles.
func (r *Radarr) GetQualityProfiles() ([]*QualityProfile, error) {
	var profiles []*QualityProfile

	err := r.GetInto("v3/qualityProfile", nil, &profiles)
	if err != nil {
		return nil, fmt.Errorf("api.Get(qualityProfile): %w", err)
	}

	return profiles, nil
}

// GetRootFolders returns all configured root folders.
func (r *Radarr) GetRootFolders() ([]*RootFolder, error) {
	var folders []*RootFolder

	err := r.GetInto("v3/rootFolder", nil, &folders)
	if err != nil {
		return nil, fmt.Errorf("api.Get(rootFolder): %w", err)
	}

	return folders, nil
}

// UpdateMovie sends a PUT request to update a movie in place.
func (r *Radarr) UpdateMovie(movieID int64, movie *Movie) error {
	put, err := json.Marshal(movie)
	if err != nil {
		return fmt.Errorf("json.Marshal(movie): %w", err)
	}

	params := make(url.Values)
	params.Add("moveFiles", "true")

	b, err := r.Put("v3/movie/"+strconv.FormatInt(movieID, 10), params, put)
	if err != nil {
		return fmt.Errorf("api.Put(movie): %w", err)
	}

	log.Println("SHOW THIS TO CAPTAIN plz:", string(b))

	return nil
}

// AddMovie adds a movie to the queue.
func (r *Radarr) AddMovie(movie *AddMovieInput) (*AddMovieOutput, error) {
	body, err := json.Marshal(movie)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal(movie): %w", err)
	}

	params := make(url.Values)
	params.Add("moveFiles", "true")

	var output AddMovieOutput
	if err := r.PostInto("v3/movie", params, body, &output); err != nil {
		return nil, fmt.Errorf("api.Post(movie): %w", err)
	}

	return &output, nil
}
