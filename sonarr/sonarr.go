package sonarr

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"

	"golift.io/starr"
)

// GetQueue returns a single page from the Sonarr Queue (processing, but not yet imported).
// WARNING: 12/30/2021 - this method changed.
// If you need control over the page, use sonarr.GetQueuePage().
// This function simply returns the number of queue records desired,
// up to the number of records present in the application.
// It grabs records in (paginated) batches of perPage, and concatenates
// them into one list.  Passing zero for records will return all of them.
func (s *Sonarr) GetQueue(records, perPage int) (*Queue, error) { //nolint:dupl
	queue := &Queue{Records: []*QueueRecord{}}
	perPage = starr.SetPerPage(records, perPage)

	for page := 1; ; page++ {
		curr, err := s.GetQueuePage(&starr.Req{PageSize: perPage, Page: page})
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

// GetQueuePage returns a single page from the Sonarr Queue.
// The page size and number is configurable with the input request parameters.
func (s *Sonarr) GetQueuePage(params *starr.Req) (*Queue, error) {
	var queue Queue

	params.CheckSet("sortKey", "timeleft")
	params.CheckSet("includeUnknownSeriesItems", "true")

	err := s.GetInto("v3/queue", params.Params(), &queue)
	if err != nil {
		return nil, fmt.Errorf("api.Get(queue): %w", err)
	}

	return &queue, nil
}

// GetTags returns all the tags.
func (s *Sonarr) GetTags() ([]*starr.Tag, error) {
	var tags []*starr.Tag

	err := s.GetInto("v3/tag", nil, &tags)
	if err != nil {
		return nil, fmt.Errorf("api.Get(tag): %w", err)
	}

	return tags, nil
}

// UpdateTag updates the label for a tag.
func (s *Sonarr) UpdateTag(tagID int, label string) (int, error) {
	body, err := json.Marshal(&starr.Tag{Label: label, ID: tagID})
	if err != nil {
		return 0, fmt.Errorf("json.Marshal(tag): %w", err)
	}

	var tag starr.Tag
	if err = s.PutInto("v3/tag", nil, body, &tag); err != nil {
		return tag.ID, fmt.Errorf("api.Put(tag): %w", err)
	}

	return tag.ID, nil
}

// AddTag adds a tag or returns the ID for an existing tag.
func (s *Sonarr) AddTag(label string) (int, error) {
	body, err := json.Marshal(&starr.Tag{Label: label, ID: 0})
	if err != nil {
		return 0, fmt.Errorf("json.Marshal(tag): %w", err)
	}

	var tag starr.Tag
	if err = s.PostInto("v3/tag", nil, body, &tag); err != nil {
		return tag.ID, fmt.Errorf("api.Post(tag): %w", err)
	}

	return tag.ID, nil
}

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

// AddQualityProfile updates a quality profile in place.
func (s *Sonarr) AddQualityProfile(profile *QualityProfile) (int64, error) {
	post, err := json.Marshal(profile)
	if err != nil {
		return 0, fmt.Errorf("json.Marshal(profile): %w", err)
	}

	var output QualityProfile

	err = s.PostInto("v3/qualityProfile", nil, post, &output)
	if err != nil {
		return 0, fmt.Errorf("api.Post(qualityProfile): %w", err)
	}

	return output.ID, nil
}

// UpdateQualityProfile updates a quality profile in place.
func (s *Sonarr) UpdateQualityProfile(profile *QualityProfile) error {
	put, err := json.Marshal(profile)
	if err != nil {
		return fmt.Errorf("json.Marshal(profile): %w", err)
	}

	_, err = s.Put("v3/qualityProfile/"+strconv.FormatInt(profile.ID, starr.Base10), nil, put)
	if err != nil {
		return fmt.Errorf("api.Put(qualityProfile): %w", err)
	}

	return nil
}

// GetReleaseProfiles returns all configured release profiles.
func (s *Sonarr) GetReleaseProfiles() ([]*ReleaseProfile, error) {
	var profiles []*ReleaseProfile

	err := s.GetInto("v3/releaseProfile", nil, &profiles)
	if err != nil {
		return nil, fmt.Errorf("api.Get(releaseProfile): %w", err)
	}

	return profiles, nil
}

// AddReleaseProfile updates a release profile in place.
func (s *Sonarr) AddReleaseProfile(profile *ReleaseProfile) (int64, error) {
	post, err := json.Marshal(profile)
	if err != nil {
		return 0, fmt.Errorf("json.Marshal(profile): %w", err)
	}

	var output ReleaseProfile

	err = s.PostInto("v3/releaseProfile", nil, post, &output)
	if err != nil {
		return 0, fmt.Errorf("api.Post(releaseProfile): %w", err)
	}

	return output.ID, nil
}

// UpdateReleaseProfile updates a release profile in place.
func (s *Sonarr) UpdateReleaseProfile(profile *ReleaseProfile) error {
	put, err := json.Marshal(profile)
	if err != nil {
		return fmt.Errorf("json.Marshal(profile): %w", err)
	}

	_, err = s.Put("v3/releaseProfile/"+strconv.FormatInt(profile.ID, starr.Base10), nil, put)
	if err != nil {
		return fmt.Errorf("api.Put(releaseProfile): %w", err)
	}

	return nil
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
		params.Add("term", "tvdbid:"+strconv.FormatInt(tvdbID, starr.Base10))
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
		params.Add("tvdbId", strconv.FormatInt(tvdbID, starr.Base10))
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

	err := s.GetInto("v3/series/"+strconv.FormatInt(seriesID, starr.Base10), nil, &series)
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

// GetCommands returns all available Sonarr commands.
// These can be used with SendCommand.
func (s *Sonarr) GetCommands() ([]*CommandResponse, error) {
	var output []*CommandResponse

	if err := s.GetInto("v3/command", nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(command): %w", err)
	}

	return output, nil
}

// SendCommand sends a command to Sonarr.
func (s *Sonarr) SendCommand(cmd *CommandRequest) (*CommandResponse, error) {
	var output CommandResponse

	if cmd == nil || cmd.Name == "" {
		return &output, nil
	}

	body, err := json.Marshal(cmd)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal(cmd): %w", err)
	}

	if err := s.PostInto("v3/command", nil, body, &output); err != nil {
		return nil, fmt.Errorf("api.Post(command): %w", err)
	}

	return &output, nil
}

// GetCommandStatus returns the status of an already started command.
func (s *Sonarr) GetCommandStatus(commandID int64) (*CommandResponse, error) {
	var output CommandResponse

	if commandID == 0 {
		return &output, nil
	}

	err := s.GetInto("v3/command/"+strconv.FormatInt(commandID, starr.Base10), nil, &output)
	if err != nil {
		return nil, fmt.Errorf("api.Post(command): %w", err)
	}

	return &output, nil
}

// GetSeriesEpisodes returns all episodes for a series by series ID.
// You can get series IDs from GetAllSeries() and GetSeries().
func (s *Sonarr) GetSeriesEpisodes(seriesID int64) ([]*Episode, error) {
	var output []*Episode

	params := make(url.Values)
	params.Add("seriesId", strconv.FormatInt(seriesID, starr.Base10))

	err := s.GetInto("v3/episode?seriesId", params, &output)
	if err != nil {
		return nil, fmt.Errorf("api.Get(episode): %w", err)
	}

	return output, nil
}

// MonitorEpisode sends a request to monitor (true) or unmonitor (false) a list of episodes by ID.
// You can get episode IDs from GetSeriesEpisodes().
func (s *Sonarr) MonitorEpisode(episodeIDs []int64, monitor bool) ([]*Episode, error) {
	var (
		input, _ = json.Marshal(&struct {
			E []int64 `json:"episodeIds"`
			M bool    `json:"monitored"`
		}{E: episodeIDs, M: monitor})
		output []*Episode
	)

	if err := s.PutInto("v3/episode/monitor", nil, input, &output); err != nil {
		return nil, fmt.Errorf("api.Put(episode/monitor): %w", err)
	}

	return output, nil
}

// GetHistory returns the Sonarr History (grabs/failures/completed).
// WARNING: 12/30/2021 - this method changed.
// If you need control over the page, use sonarr.GetHistoryPage().
// This function simply returns the number of history records desired,
// up to the number of records present in the application.
// It grabs records in (paginated) batches of perPage, and concatenates
// them into one list.  Passing zero for records will return all of them.
func (s *Sonarr) GetHistory(records, perPage int) (*History, error) { //nolint:dupl
	hist := &History{Records: []*HistoryRecord{}}
	perPage = starr.SetPerPage(records, perPage)

	for page := 1; ; page++ {
		curr, err := s.GetHistoryPage(&starr.Req{PageSize: perPage, Page: page})
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

// GetHistoryPage returns a single page from the Sonarr History (grabs/failures/completed).
// The page size and number is configurable with the input request parameters.
func (s *Sonarr) GetHistoryPage(params *starr.Req) (*History, error) {
	var history History

	err := s.GetInto("v3/history", params.Params(), &history)
	if err != nil {
		return nil, fmt.Errorf("api.Get(history): %w", err)
	}

	return &history, nil
}

// Fail marks the given history item as failed by id.
func (s *Sonarr) Fail(historyID int64) error {
	if historyID < 1 {
		return fmt.Errorf("%w: invalid history ID: %d", starr.ErrRequestError, historyID)
	}

	params := make(url.Values)

	_, err := s.Get("v3/history/failed/"+strconv.FormatInt(historyID, starr.Base10), params)
	if err != nil {
		return fmt.Errorf("api.Get(history/failed): %w", err)
	}

	return nil
}

// Lookup will search for series matching the specified search term.
// Searches for new shows on TheTVDB.com utilizing sonarr.tv's caching and augmentation proxy.
func (s *Sonarr) Lookup(term string) ([]*Series, error) {
	var output []*Series

	if term == "" {
		return output, nil
	}

	params := make(url.Values)
	params.Set("term", term)

	err := s.GetInto("v3/series/lookup", params, &output)
	if err != nil {
		return nil, fmt.Errorf("api.Get(series/lookup): %w", err)
	}

	return output, nil
}

// GetBackupFiles returns all available Sonarr backup files.
// Use GetBody to download a file using BackupFile.Path.
func (s *Sonarr) GetBackupFiles() ([]*starr.BackupFile, error) {
	var output []*starr.BackupFile

	if err := s.GetInto("v3/system/backup", nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(system/backup): %w", err)
	}

	return output, nil
}
