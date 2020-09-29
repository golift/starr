package starr

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

// ReadarQueue is the /api/v1/queue endpoint.
type ReadarQueue struct {
	Page          int          `json:"page"`
	PageSize      int          `json:"pageSize"`
	SortKey       string       `json:"sortKey"`
	SortDirection string       `json:"sortDirection"`
	TotalRecords  int          `json:"totalRecords"`
	Records       []BookRecord `json:"records"`
}

// BookRecord is a book from the queue API path.
type BookRecord struct {
	AuthorID                int64           `json:"authorId"`
	BookID                  int64           `json:"bookId"`
	Quality                 BookQuality     `json:"quality"`
	Size                    float64         `json:"size"`
	Title                   string          `json:"title"`
	Sizeleft                float64         `json:"sizeleft"`
	Timeleft                string          `json:"timeleft"`
	EstimatedCompletionTime time.Time       `json:"estimatedCompletionTime"`
	Status                  string          `json:"status"`
	TrackedDownloadStatus   string          `json:"trackedDownloadStatus,omitempty"`
	TrackedDownloadState    string          `json:"trackedDownloadState,omitempty"`
	StatusMessages          []StatusMessage `json:"statusMessages,omitempty"`
	DownloadID              string          `json:"downloadId,omitempty"`
	Protocol                string          `json:"protocol"`
	DownloadClient          string          `json:"downloadClient,omitempty"`
	Indexer                 string          `json:"indexer"`
	OutputPath              string          `json:"outputPath,omitempty"`
	DownloadForced          bool            `json:"downloadForced"`
	ID                      int64           `json:"id"`
}

// BookQuality is attached to each BookRecord.
type BookQuality struct {
	Quality struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"quality"`
	Revision struct {
		Version  int64 `json:"version"`
		Real     int64 `json:"real"`
		IsRepack bool  `json:"isRepack"`
	} `json:"revision"`
}

// ReadarrQueue returns the Readarr Queue (processing, but not yet imported).
func (c *Config) ReadarrQueue(maxRecords int) ([]*ReadarQueue, error) {
	var queue []*ReadarQueue

	if maxRecords < 1 {
		maxRecords = 1
	}

	params := make(url.Values)
	params["pageSize"] = []string{strconv.Itoa(maxRecords)}

	rawJSON, err := c.Req("v1/queue", params)
	if err != nil {
		return nil, fmt.Errorf("c.Req(queue): %w", err)
	}

	if err = json.Unmarshal(rawJSON, &queue); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(response): %w", err)
	}

	return queue, nil
}
