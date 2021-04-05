package radarr

// Radarr v3 structs

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"

	"golift.io/starr"
)

// GetSystemStatus returns system status.
func (r *Radarr) GetSystemStatus() (*SystemStatus, error) {
	var status SystemStatus

	err := r.GetInto("v3/system/status", nil, &status)
	if err != nil {
		return nil, fmt.Errorf("api.Get(system/status): %w", err)
	}

	return &status, nil
}

// GetTags returns all the tags.
func (r *Radarr) GetTags() ([]*starr.Tag, error) {
	var tags []*starr.Tag

	err := r.GetInto("v3/tag", nil, &tags)
	if err != nil {
		return nil, fmt.Errorf("api.Get(tag): %w", err)
	}

	return tags, nil
}

// UpdateTag updates the label for a tag.
func (r *Radarr) UpdateTag(tagID int, label string) (int, error) {
	body, err := json.Marshal(&starr.Tag{Label: label, ID: tagID})
	if err != nil {
		return 0, fmt.Errorf("json.Marshal(tag): %w", err)
	}

	var tag starr.Tag
	if err = r.PutInto("v3/tag", nil, body, &tag); err != nil {
		return tag.ID, fmt.Errorf("api.Put(tag): %w", err)
	}

	return tag.ID, nil
}

// AddTag adds a tag or returns the ID for an existing tag.
func (r *Radarr) AddTag(label string) (int, error) {
	body, err := json.Marshal(&starr.Tag{Label: label, ID: 0})
	if err != nil {
		return 0, fmt.Errorf("json.Marshal(tag): %w", err)
	}

	var tag starr.Tag
	if err = r.PostInto("v3/tag", nil, body, &tag); err != nil {
		return tag.ID, fmt.Errorf("api.Post(tag): %w", err)
	}

	return tag.ID, nil
}

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

func (r *Radarr) GetExclusions() ([]*Exclusion, error) {
	var exclusions []*Exclusion

	err := r.GetInto("v3/exclusions", nil, &exclusions)
	if err != nil {
		return nil, fmt.Errorf("api.Get(exclusions): %w", err)
	}

	return exclusions, nil
}

var ErrRequestErr = fmt.Errorf("request error")

func (r *Radarr) DeleteExclusions(ids []int64) error {
	var errs string

	for _, id := range ids {
		_, err := r.Delete("v3/exclusions/"+strconv.FormatInt(id, 10), nil)
		if err != nil {
			errs += err.Error() + " "
		}
	}

	if errs != "" {
		return fmt.Errorf("%w: %s", ErrRequestErr, errs)
	}

	return nil
}

func (r *Radarr) AddExclusions(exclusions []*Exclusion) error {
	for i := range exclusions {
		exclusions[i].ID = 0
	}

	body, err := json.Marshal(exclusions)
	if err != nil {
		return fmt.Errorf("json.Marshal(movie): %w", err)
	}

	_, err = r.Post("v3/exclusions/bulk", nil, body)
	if err != nil {
		return fmt.Errorf("api.Post(exclusions): %w", err)
	}

	return nil
}

// GetCustomFormats returns all configured Custom Formats.
func (r *Radarr) GetCustomFormats() ([]*CustomFormat, error) {
	var cf []*CustomFormat
	if err := r.GetInto("v3/customFormat", nil, &cf); err != nil {
		return nil, fmt.Errorf("api.Get(customFormat): %w", err)
	}

	return cf, nil
}

// AddCustomFormat creates a new custom format and returns the response (with ID).
func (r *Radarr) AddCustomFormat(cf *CustomFormat) (*CustomFormat, error) {
	if cf == nil {
		return nil, nil
	}

	cf.ID = 0 // ID must be zero when adding.

	body, err := json.Marshal(cf)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal(customFormat): %w", err)
	}

	var output CustomFormat
	if err := r.PostInto("v3/customFormat", nil, body, &output); err != nil {
		return nil, fmt.Errorf("api.Post(customFormat): %w", err)
	}

	return &output, nil
}

// UpdateCustomFormat updates an existing custom format and returns the response.
func (r *Radarr) UpdateCustomFormat(cf *CustomFormat, cfID int) (*CustomFormat, error) {
	if cfID == 0 {
		cfID = cf.ID
	}

	body, err := json.Marshal(cf)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal(customFormat): %w", err)
	}

	var output CustomFormat
	if err := r.PutInto("v3/customFormat/"+strconv.Itoa(cfID), nil, body, &output); err != nil {
		return nil, fmt.Errorf("api.Put(customFormat): %w", err)
	}

	return &output, nil
}
