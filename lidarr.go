package starr

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

// LidarQueue is the /api/v1/queue endpoint.
type LidarQueue struct {
	Page          int             `json:"page"`
	PageSize      int             `json:"pageSize"`
	SortKey       string          `json:"sortKey"`
	SortDirection string          `json:"sortDirection"`
	TotalRecords  int             `json:"totalRecords"`
	Records       []*LidarrRecord `json:"records"`
}

// LidarrRecord represents the records returns by the /api/v1/queue endpoint.
type LidarrRecord struct {
	ArtistID int64 `json:"artistId"`
	AlbumID  int64 `json:"albumId"`
	Quality  struct {
		Quality struct {
			ID   int64  `json:"id"`
			Name string `json:"name"`
		} `json:"quality"`
		Revision struct {
			Version  int  `json:"version"`
			Real     int  `json:"real"`
			IsRepack bool `json:"isRepack"`
		} `json:"revision"`
	} `json:"quality"`
	Size                    float64       `json:"size"`
	Title                   string        `json:"title"`
	Sizeleft                float64       `json:"sizeleft"`
	Timeleft                string        `json:"timeleft"`
	EstimatedCompletionTime time.Time     `json:"estimatedCompletionTime"`
	Status                  string        `json:"status"`
	TrackedDownloadStatus   string        `json:"trackedDownloadStatus"`
	StatusMessages          []interface{} `json:"statusMessages"`
	DownloadID              string        `json:"downloadId"`
	Protocol                string        `json:"protocol"`
	DownloadClient          string        `json:"downloadClient"`
	Indexer                 string        `json:"indexer"`
	OutputPath              string        `json:"outputPath"`
	DownloadForced          bool          `json:"downloadForced"`
	ID                      int64         `json:"id"`
}

// LidarrQueue returns the Lidarr Queue
func (c *Config) LidarrQueue(maxRecords int) ([]*LidarrRecord, error) {
	var queue *LidarQueue

	if maxRecords < 1 {
		maxRecords = 1
	}

	params := make(url.Values)

	params.Set("sortKey", "timeleft")
	params.Set("sortDir", "asc")
	params.Set("pageSize", strconv.Itoa(maxRecords))

	rawJSON, err := c.Req("v1/queue", params)
	if err != nil {
		return nil, fmt.Errorf("c.Req(queue): %v", err)
	}

	if err = json.Unmarshal(rawJSON, &queue); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(response): %v", err)
	}

	return queue.Records, nil
}
