package lidarr

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"
)

// GetQualityDefinition returns the Quality Definitions.
func (l *Lidarr) GetQualityDefinition() ([]*QualityDefinition, error) {
	var definition []*QualityDefinition

	err := l.GetInto("v1/qualitydefinition", nil, &definition)
	if err != nil {
		return nil, fmt.Errorf("api.Get(qualitydefinition): %w", err)
	}

	return definition, nil
}

// GetQualityProfiles returns the quality profiles.
func (l *Lidarr) GetQualityProfiles() ([]*QualityProfile, error) {
	var profiles []*QualityProfile

	err := l.GetInto("v1/qualityprofile", nil, &profiles)
	if err != nil {
		return nil, fmt.Errorf("api.Get(qualityprofile): %w", err)
	}

	return profiles, nil
}

// GetRootFolders returns all configured root folders.
func (l *Lidarr) GetRootFolders() ([]*RootFolder, error) {
	var folders []*RootFolder

	err := l.GetInto("v1/rootFolder", nil, &folders)
	if err != nil {
		return nil, fmt.Errorf("api.Get(rootFolder): %w", err)
	}

	return folders, nil
}

// GetMetadataProfiles returns the metadata profiles.
func (l *Lidarr) GetMetadataProfiles() ([]*MetadataProfile, error) {
	var profiles []*MetadataProfile

	err := l.GetInto("v1/metadataprofile", nil, &profiles)
	if err != nil {
		return nil, fmt.Errorf("api.Get(metadataprofile): %w", err)
	}

	return profiles, nil
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

	var queue Queue

	err := l.GetInto("v1/queue", params, &queue)
	if err != nil {
		return nil, fmt.Errorf("api.Get(queue): %w", err)
	}

	return &queue, nil
}

// GetSystemStatus returns system status.
func (l *Lidarr) GetSystemStatus() (*SystemStatus, error) {
	var status SystemStatus

	err := l.GetInto("v1/system/status", nil, &status)
	if err != nil {
		return nil, fmt.Errorf("api.Get(system/status): %w", err)
	}

	return &status, nil
}

// GetArtist returns an artist or all artists.
func (l *Lidarr) GetArtist(mbID string) ([]*Artist, error) {
	params := make(url.Values)

	if mbID != "" {
		params.Add("mbId", mbID)
	}

	var artist []*Artist

	err := l.GetInto("v1/artist", params, &artist)
	if err != nil {
		return artist, fmt.Errorf("api.Get(artist): %w", err)
	}

	return artist, nil
}

// AddArtist adds a new artist to Lidarr, and probably does not yet work.
func (l *Lidarr) AddArtist(artist *AddArtistInput) (*AddArtistOutput, error) {
	body, err := json.Marshal(artist)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal(album): %w", err)
	}

	params := make(url.Values)
	params.Add("moveFiles", "true")

	var output AddArtistOutput

	err = l.PostInto("v1/artist", params, body, &output)
	if err != nil {
		return nil, fmt.Errorf("api.Post(artist): %w", err)
	}

	return &output, nil
}

// GetAlbum returns an album or all albums if mbID is 0.
// mbID is the music brainz UUID for a "release-group".
func (l *Lidarr) GetAlbum(mbID string) ([]*Album, error) {
	params := make(url.Values)

	if mbID != "" {
		params.Add("ForeignAlbumId", mbID)
	}

	var albums []*Album

	err := l.GetInto("v1/album", params, &albums)
	if err != nil {
		return nil, fmt.Errorf("api.Get(album): %w", err)
	}

	return albums, nil
}

// GetAlbumByID returns an album by DB ID.
func (l *Lidarr) GetAlbumByID(albumID int64) (*Album, error) {
	var album Album

	err := l.GetInto("v1/album"+strconv.FormatInt(albumID, 10), nil, &album)
	if err != nil {
		return nil, fmt.Errorf("api.Get(album): %w", err)
	}

	return &album, nil
}

// UpdateAlbum updates an album in place.
func (l *Lidarr) UpdateAlbum(albumID int64, album *Album) error {
	put, err := json.Marshal(album)
	if err != nil {
		return fmt.Errorf("json.Marshal(album): %w", err)
	}

	params := make(url.Values)
	params.Add("moveFiles", "true")

	b, err := l.Put("v1/album/"+strconv.FormatInt(albumID, 10), params, put)
	if err != nil {
		return fmt.Errorf("api.Put(album): %w", err)
	}

	log.Println("SHOW THIS TO CAPTAIN plz:", string(b))

	return nil
}

// AddAlbum adds a new album to Lidarr, and probably does not yet work.
func (l *Lidarr) AddAlbum(album *AddAlbumInput) (*AddAlbumOutput, error) {
	if album.Releases == nil {
		album.Releases = make([]*AddAlbumInputRelease, 0)
	}

	body, err := json.Marshal(album)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal(album): %w", err)
	}

	params := make(url.Values)
	params.Add("moveFiles", "true")

	var output AddAlbumOutput

	err = l.PostInto("v1/album", params, body, &output)
	if err != nil {
		return nil, fmt.Errorf("api.Post(album): %w", err)
	}

	return &output, nil
}
