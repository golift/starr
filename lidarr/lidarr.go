package lidarr

import (
	"fmt"
	"net/url"
	"strconv"
)

// GetQualityDefinition returns the Quality Definitions.
func (l *Lidarr) GetQualityDefinition() ([]*QualityDefinition, error) {
	var definition []*QualityDefinition
	if err := l.GetInto("v1/qualitydefinition", nil, &definition); err != nil {
		return nil, fmt.Errorf("api.Get(qualitydefinition): %w", err)
	}

	return definition, nil
}

// GetQualityProfiles returns the quality profiles.
func (l *Lidarr) GetQualityProfiles() ([]*QualityProfile, error) {
	var profiles []*QualityProfile
	if err := l.GetInto("v1/qualityprofile", nil, &profiles); err != nil {
		return nil, fmt.Errorf("api.Get(qualityprofile): %w", err)
	}

	return profiles, nil
}

// GetRootFolders returns all configured root folders.
func (l *Lidarr) GetRootFolders() ([]*RootFolder, error) {
	var folders []*RootFolder
	if err := l.GetInto("v1/rootFolder", nil, &folders); err != nil {
		return nil, fmt.Errorf("api.Get(rootFolder): %w", err)
	}

	return folders, nil
}

// GetQueue returns the Lidarr Queue.
func (l *Lidarr) GetQueue(maxRecords int) (*Queue, error) {
	if maxRecords < 1 {
		maxRecords = 1
	}

	params := make(url.Values)
	params.Set("sortKey", "timeleft")
	params.Set("sortDir", "asc")
	params.Set("pageSize", strconv.Itoa(maxRecords))

	var queue *Queue
	if err := l.GetInto("v1/queue", params, queue); err != nil {
		return nil, fmt.Errorf("api.Get(queue): %w", err)
	}

	return queue, nil
}

// GetSystemStatus returns system status.
func (l *Lidarr) GetSystemStatus() (*SystemStatus, error) {
	var status *SystemStatus
	if err := l.GetInto("v1/system/status", nil, status); err != nil {
		return status, fmt.Errorf("api.Get(system/status): %w", err)
	}

	return status, nil
}

/* unknown structure/input data format
type Album struct{}

func (l *Lidarr) GetAlbum(albumID string) ([]*Album, error) {
	var albums []*Album

	params := make(url.Values)

	if albumID != "" {
		params.Add("albumID", albumID)
	}

	rawJSON, err := l.config.Req("v1/system/status", nil)
	if err != nil {
		return nil, fmt.Errorf("api.Get(status): %w", err)
	}

	if err = json.Unmarshal(rawJSON, &albums); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(response): %w", err)
	}

	return albums, nil
}
*/
