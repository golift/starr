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
// WARNING: 12/30/2021 - this method changed. The second argument no longer
// controls which page is returned, but instead adjusts the pagination size.
// If you need control over the page, use radarr.GetHistoryPage().
// This function simply returns the number of history records desired,
// up to the number of records present in the application.
// It grabs records in (paginated) batches of perPage, and concatenates
// them into one list.  Passing zero for records will return all of them.
func (r *Radarr) GetHistory(records, perPage int) (*History, error) { //nolint:dupl
	hist := &History{Records: []*HistoryRecord{}}
	perPage = starr.SetPerPage(records, perPage)

	for page := 1; ; page++ {
		curr, err := r.GetHistoryPage(&starr.Req{PageSize: perPage, Page: page})
		if err != nil {
			return nil, err
		}

		hist.Records = append(hist.Records, curr.Records...)
		if len(hist.Records) >= curr.TotalRecords ||
			(len(hist.Records) >= records && records != 0) ||
			len(curr.Records) == 0 {
			hist.PageSize = curr.TotalRecords
			hist.TotalRecords = curr.TotalRecords
			hist.SortDirection = curr.SortDirection
			hist.SortKey = curr.SortKey

			break
		}

		perPage = starr.AdjustPerPage(records, curr.TotalRecords, len(hist.Records), perPage)
	}

	return hist, nil
}

// GetHistoryPage returns a single page from the Radarr History (grabs/failures/completed).
// The page size and number is configurable with the input request parameters.
func (r *Radarr) GetHistoryPage(params *starr.Req) (*History, error) {
	var history History

	err := r.GetInto("v3/history", params.Params(), &history)
	if err != nil {
		return nil, fmt.Errorf("api.Get(history): %w", err)
	}

	return &history, nil
}

// GetQueue returns a single page from the Radarr Queue (processing, but not yet imported).
// WARNING: 12/30/2021 - this method changed. The second argument no longer
// controls which page is returned, but instead adjusts the pagination size.
// If you need control over the page, use radarr.GetQueuePage().
// This function simply returns the number of queue records desired,
// up to the number of records present in the application.
// It grabs records in (paginated) batches of perPage, and concatenates
// them into one list.  Passing zero for records will return all of them.
func (r *Radarr) GetQueue(records, perPage int) (*Queue, error) { //nolint:dupl
	queue := &Queue{Records: []*QueueRecord{}}
	perPage = starr.SetPerPage(records, perPage)

	for page := 1; ; page++ {
		curr, err := r.GetQueuePage(&starr.Req{PageSize: perPage, Page: page})
		if err != nil {
			return nil, err
		}

		queue.Records = append(queue.Records, curr.Records...)
		if len(queue.Records) >= curr.TotalRecords ||
			(len(queue.Records) >= records && records != 0) ||
			len(curr.Records) == 0 {
			queue.PageSize = curr.TotalRecords
			queue.TotalRecords = curr.TotalRecords
			queue.SortDirection = curr.SortDirection
			queue.SortKey = curr.SortKey

			break
		}

		perPage = starr.AdjustPerPage(records, curr.TotalRecords, len(queue.Records), perPage)
	}

	return queue, nil
}

// GetQueuePage returns a single page from the Radarr Queue.
// The page size and number is configurable with the input request parameters.
func (r *Radarr) GetQueuePage(params *starr.Req) (*Queue, error) {
	var queue Queue

	params.CheckSet("sortKey", "timeleft")
	params.CheckSet("includeUnknownMovieItems", "true")

	err := r.GetInto("v3/queue", params.Params(), &queue)
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
	var output []*CustomFormat
	if err := r.GetInto("v3/customFormat", nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(customFormat): %w", err)
	}

	return output, nil
}

// AddCustomFormat creates a new custom format and returns the response (with ID).
func (r *Radarr) AddCustomFormat(format *CustomFormat) (*CustomFormat, error) {
	var output CustomFormat

	if format == nil {
		return &output, nil
	}

	format.ID = 0 // ID must be zero when adding.

	body, err := json.Marshal(format)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal(customFormat): %w", err)
	}

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
	var output []*ImportList
	if err := r.GetInto("v3/importlist", nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(importlist): %w", err)
	}

	return output, nil
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
func (r *Radarr) UpdateImportList(list *ImportList) (*ImportList, error) {
	body, err := json.Marshal(list)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal(importlist): %w", err)
	}

	var output ImportList

	err = r.PutInto("v3/importlist/"+strconv.FormatInt(list.ID, starr.Base10), nil, body, &output)
	if err != nil {
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
	var output CommandResponse

	if cmd == nil || cmd.Name == "" {
		return &output, nil
	}

	body, err := json.Marshal(cmd)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal(cmd): %w", err)
	}

	if err := r.PostInto("v3/command", nil, body, &output); err != nil {
		return nil, fmt.Errorf("api.Post(command): %w", err)
	}

	return &output, nil
}

// Lookup will search for movies matching the specified search term.
func (r *Radarr) Lookup(term string) ([]*Movie, error) {
	var output []*Movie

	if term == "" {
		return output, nil
	}

	params := make(url.Values)
	params.Set("term", term)

	err := r.GetInto("v3/movie/lookup", params, &output)
	if err != nil {
		return nil, fmt.Errorf("api.Get(movie/lookup): %w", err)
	}

	return output, nil
}

// GetBackupFiles returns all available Radarr backup files.
// Use GetBody to download a file using BackupFile.Path.
func (r *Radarr) GetBackupFiles() ([]*starr.BackupFile, error) {
	var output []*starr.BackupFile

	if err := r.GetInto("v3/system/backup", nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(system/backup): %w", err)
	}

	return output, nil
}
