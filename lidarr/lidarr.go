package lidarr

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// GetQualityDefinition returns the Quality Definitions.
func (l *Lidarr) GetQualityDefinition() ([]*QualityDefinition, error) {
	var definition []*QualityDefinition

	rawJSON, err := l.config.Req("v1/qualitydefinition", nil)
	if err != nil {
		return nil, fmt.Errorf("c.Req(qualitydefinition): %w", err)
	}

	if err = json.Unmarshal(rawJSON, &definition); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(response): %w", err)
	}

	return definition, nil
}

// GetQualityProfiles returns the quality profiles.
func (l *Lidarr) GetQualityProfiles() ([]*QualityProfile, error) {
	var profiles []*QualityProfile

	rawJSON, err := l.config.Req("v1/qualityprofile", nil)
	if err != nil {
		return nil, fmt.Errorf("c.Req(qualityprofile): %w", err)
	}

	if err = json.Unmarshal(rawJSON, &profiles); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(response): %w", err)
	}

	return profiles, nil
}

// GetRootFolders returns all configured root folders.
func (l *Lidarr) GetRootFolders() ([]*RootFolder, error) {
	var folders []*RootFolder

	rawJSON, err := l.config.Req("v1/rootFolder", nil)
	if err != nil {
		return nil, fmt.Errorf("c.Req(rootFolder): %w", err)
	}

	if err = json.Unmarshal(rawJSON, &folders); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(response): %w", err)
	}

	return folders, nil
}

// GetQueue returns the Lidarr Queue.
func (l *Lidarr) GetQueue(maxRecords int) (*Queue, error) {
	var queue *Queue

	if maxRecords < 1 {
		maxRecords = 1
	}

	params := make(url.Values)

	params.Set("sortKey", "timeleft")
	params.Set("sortDir", "asc")
	params.Set("pageSize", strconv.Itoa(maxRecords))

	rawJSON, err := l.config.Req("v1/queue", params)
	if err != nil {
		return nil, fmt.Errorf("c.Req(queue): %w", err)
	}

	if err = json.Unmarshal(rawJSON, &queue); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(response): %w", err)
	}

	return queue, nil
}

// GetSystemStatus returns system status.
func (l *Lidarr) GetSystemStatus() (*SystemStatus, error) {
	var status *SystemStatus

	rawJSON, err := l.config.Req("v1/system/status", nil)
	if err != nil {
		return status, fmt.Errorf("c.Req(status): %w", err)
	}

	if err = json.Unmarshal(rawJSON, &status); err != nil {
		return status, fmt.Errorf("json.Unmarshal(response): %w", err)
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
		return nil, fmt.Errorf("c.Req(status): %w", err)
	}

	if err = json.Unmarshal(rawJSON, &albums); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(response): %w", err)
	}

	return albums, nil
}
*/
