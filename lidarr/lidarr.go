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
	if err := l.a.GetInto("v1/qualitydefinition", nil, &definition); err != nil {
		return nil, fmt.Errorf("api.Get(qualitydefinition): %w", err)
	}

	return definition, nil
}

// GetQualityProfiles returns the quality profiles.
func (l *Lidarr) GetQualityProfiles() ([]*QualityProfile, error) {
	var profiles []*QualityProfile
	if err := l.a.GetInto("v1/qualityprofile", nil, &profiles); err != nil {
		return nil, fmt.Errorf("api.Get(qualityprofile): %w", err)
	}

	return profiles, nil
}

// GetRootFolders returns all configured root folders.
func (l *Lidarr) GetRootFolders() ([]*RootFolder, error) {
	var folders []*RootFolder
	if err := l.a.GetInto("v1/rootFolder", nil, &folders); err != nil {
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
	if err := l.a.GetInto("v1/queue", params, queue); err != nil {
		return nil, fmt.Errorf("api.Get(queue): %w", err)
	}

	return queue, nil
}

// GetSystemStatus returns system status.
func (l *Lidarr) GetSystemStatus() (*SystemStatus, error) {
	var status *SystemStatus
	if err := l.a.GetInto("v1/system/status", nil, status); err != nil {
		return status, fmt.Errorf("api.Get(system/status): %w", err)
	}

	return status, nil
}

// GetArtist returns an artist or all artists.
// mbID is the music brainz UUID for an "artist".
func (l *Lidarr) GetArtist(artistUUID string) ([]*Artist, error) {
	params := make(url.Values)

	if artistUUID != "" {
		params.Add("mbId", artistUUID)
	}

	var artist []*Artist
	if err := l.a.GetInto("v1/artist", params, &artist); err != nil {
		return artist, fmt.Errorf("api.Get(artist): %w", err)
	}

	return artist, nil
}

// GetAlbum returns an album or all albums if mbID is 0.
// mbID is the music brainz UUID for a "release-group".
func (l *Lidarr) GetAlbum(albumUUID string) ([]*Album, error) {
	params := make(url.Values)

	if albumUUID != "" {
		params.Add("ForeignAlbumId", albumUUID)
	}

	albums := []*Album{}
	if err := l.a.GetInto("v1/album", params, &albums); err != nil {
		return nil, fmt.Errorf("api.Get(album): %w", err)
	}

	return albums, nil
}

// AddAlbum adds a new album to Lidarr, and probably does not yet work.
func (l *Lidarr) AddAlbum(album *AddAlbumInput) (*AddAlbumOutput, error) {
	body, err := json.Marshal(album)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal(album): %w", err)
	}

	params := make(url.Values)
	params.Add("moveFiles", "true")

	var added AddAlbumOutput
	if err := l.a.PostInto("v1/album", params, body, &added); err != nil {
		return nil, fmt.Errorf("api.Post(album): %w", err)
	}

	return &added, nil
}
