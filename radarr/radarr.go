package radarr

// Radarr v3 structs

import (
	"encoding/json"
	"fmt"
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
func (r *Radarr) GetHistory(maxRecords, page int) (*History, error) {
	if maxRecords < 1 {
		maxRecords = 10
	}

	if page < 1 {
		page = 1
	}

	params := make(url.Values)
	params.Set("sortKey", "date")
	params.Set("sortDir", "asc")
	params.Set("pageSize", strconv.Itoa(maxRecords))
	params.Set("page", strconv.Itoa(page))

	var history History

	err := r.GetInto("v3/history", params, &history)
	if err != nil {
		return nil, fmt.Errorf("api.Get(history): %w", err)
	}

	return &history, nil
}

// GetQueue returns the Radarr Queue (processing, but not yet imported).
func (r *Radarr) GetQueue(maxRecords, page int) (*Queue, error) {
	if maxRecords < 1 {
		maxRecords = 10
	}

	if page < 1 {
		page = 1
	}

	params := make(url.Values)
	params.Set("sortKey", "timeleft")
	params.Set("sortDir", "asc")
	params.Set("pageSize", strconv.Itoa(maxRecords))
	params.Set("page", strconv.Itoa(page))
	params.Set("includeUnknownMovieItems", "true")

	var queue Queue

	err := r.GetInto("v3/queue", params, &queue)
	if err != nil {
		return nil, fmt.Errorf("api.Get(queue): %w", err)
	}

	return &queue, nil
}

// GetMovie grabs a movie from the queue, or all movies if tmdbId is 0.
func (r *Radarr) GetMovie(tmdbID int64) ([]*Movie, error) {
	params := make(url.Values)
	if tmdbID != 0 {
		params.Set("tmdbId", strconv.FormatInt(tmdbID, starr.Base10))
	}

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

	err := r.GetInto("v3/movie/"+strconv.FormatInt(movieID, starr.Base10), nil, &movie)
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

// AddQualityProfile updates a quality profile in place.
func (r *Radarr) AddQualityProfile(profile *QualityProfile) (int64, error) {
	post, err := json.Marshal(profile)
	if err != nil {
		return 0, fmt.Errorf("json.Marshal(profile): %w", err)
	}

	var output QualityProfile

	err = r.PostInto("v3/qualityProfile", nil, post, &output)
	if err != nil {
		return 0, fmt.Errorf("api.Post(qualityProfile): %w", err)
	}

	return output.ID, nil
}

// UpdateQualityProfile updates a quality profile in place.
func (r *Radarr) UpdateQualityProfile(profile *QualityProfile) error {
	put, err := json.Marshal(profile)
	if err != nil {
		return fmt.Errorf("json.Marshal(profile): %w", err)
	}

	_, err = r.Put("v3/qualityProfile/"+strconv.FormatInt(profile.ID, starr.Base10), nil, put)
	if err != nil {
		return fmt.Errorf("api.Put(qualityProfile): %w", err)
	}

	return nil
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

	_, err = r.Put("v3/movie/"+strconv.FormatInt(movieID, starr.Base10), params, put)
	if err != nil {
		return fmt.Errorf("api.Put(movie): %w", err)
	}

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

// GetExclusions returns all configured exclusions from Radarr.
func (r *Radarr) GetExclusions() ([]*Exclusion, error) {
	var exclusions []*Exclusion

	err := r.GetInto("v3/exclusions", nil, &exclusions)
	if err != nil {
		return nil, fmt.Errorf("api.Get(exclusions): %w", err)
	}

	return exclusions, nil
}

// ErrRequestError is returned when bad input is provided.
var ErrRequestError = fmt.Errorf("request error")

// DeleteExclusions removes exclusions from Radarr.
func (r *Radarr) DeleteExclusions(ids []int64) error {
	var errs string

	for _, id := range ids {
		_, err := r.Delete("v3/exclusions/"+strconv.FormatInt(id, starr.Base10), nil)
		if err != nil {
			errs += err.Error() + " "
		}
	}

	if errs != "" {
		return fmt.Errorf("%w: %s", ErrRequestError, errs)
	}

	return nil
}

// AddExclusions adds an exclusion to Radarr.
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

// GetImportLists returns all import lists.
func (r *Radarr) GetImportLists() ([]*ImportList, error) {
	var il []*ImportList
	if err := r.GetInto("v3/importlist", nil, &il); err != nil {
		return nil, fmt.Errorf("api.Get(importlist): %w", err)
	}

	return il, nil
}

// CreateImportList creates an import list in Radarr.
func (r *Radarr) CreateImportList(il *ImportList) (*ImportList, error) {
	il.ID = 0

	body, err := json.Marshal(il)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal(importlist): %w", err)
	}

	var output ImportList
	if err := r.PostInto("v3/importlist", nil, body, &output); err != nil {
		return nil, fmt.Errorf("api.Post(importlist): %w", err)
	}

	return &output, nil
}

// DeleteImportList removes an import list from Radarr.
func (r *Radarr) DeleteImportList(ids []int64) error {
	var errs string

	for _, id := range ids {
		_, err := r.Delete("v3/importlist/"+strconv.FormatInt(id, starr.Base10), nil)
		if err != nil {
			errs += fmt.Errorf("api.Delete(importlist): %w", err).Error() + " "
		}
	}

	if errs != "" {
		return fmt.Errorf("%w: %s", ErrRequestError, errs)
	}

	return nil
}

// UpdateImportList updates an existing import list and returns the response.
func (r *Radarr) UpdateImportList(il *ImportList) (*ImportList, error) {
	body, err := json.Marshal(il)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal(importlist): %w", err)
	}

	var output ImportList
	if err := r.PutInto("v3/importlist/"+strconv.FormatInt(il.ID, starr.Base10), nil, body, &output); err != nil {
		return nil, fmt.Errorf("api.Put(importlist): %w", err)
	}

	return &output, nil
}

// GetCommands returns all available Radarr commands.
func (r *Radarr) GetCommands() ([]*CommandResponse, error) {
	var output []*CommandResponse

	if err := r.GetInto("v3/command", nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(command): %w", err)
	}

	return output, nil
}

// SendCommand sends a command to Radarr.
func (r *Radarr) SendCommand(cmd *CommandRequest) (*CommandResponse, error) {
	if cmd == nil || cmd.Name == "" {
		return nil, nil
	}

	body, err := json.Marshal(cmd)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal(cmd): %w", err)
	}

	var output CommandResponse

	if err := r.PostInto("v3/command", nil, body, &output); err != nil {
		return nil, fmt.Errorf("api.Post(command): %w", err)
	}

	return &output, nil
}

// Lookup will search for movies matching the specified search term.
func (r *Radarr) Lookup(term string) ([]Movie, error) {
	if term == "" {
		return nil, nil
	}

	var out []Movie

	params := make(url.Values)
	params.Set("term", url.QueryEscape(term))

	err := r.GetInto("v3/movie/lookup", params, &out)
	if err != nil {
		return nil, fmt.Errorf("failed to lookup movie: %w", err)
	}

	return out, nil
}
