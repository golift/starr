package lidarr

import (
	"context"
	"fmt"
	"path"
	"time"

	"golift.io/starr"
)

const bpQueue = APIver + "/queue"

// Queue is the /api/v1/queue endpoint.
type Queue struct {
	Page          int            `json:"page"`
	PageSize      int            `json:"pageSize"`
	SortKey       string         `json:"sortKey"`
	SortDirection string         `json:"sortDirection"`
	TotalRecords  int            `json:"totalRecords"`
	Records       []*QueueRecord `json:"records"`
}

// QueueRecord represents the records returns by the /api/v1/queue endpoint.
type QueueRecord struct {
	ArtistID                int64                  `json:"artistId"`
	AlbumID                 int64                  `json:"albumId"`
	Quality                 *starr.Quality         `json:"quality"`
	Size                    float64                `json:"size"`
	Title                   string                 `json:"title"`
	Sizeleft                float64                `json:"sizeleft"`
	Timeleft                string                 `json:"timeleft"`
	EstimatedCompletionTime time.Time              `json:"estimatedCompletionTime"`
	Status                  string                 `json:"status"`
	TrackedDownloadStatus   string                 `json:"trackedDownloadStatus"`
	StatusMessages          []*starr.StatusMessage `json:"statusMessages"`
	DownloadID              string                 `json:"downloadId"`
	Protocol                string                 `json:"protocol"`
	DownloadClient          string                 `json:"downloadClient"`
	Indexer                 string                 `json:"indexer"`
	OutputPath              string                 `json:"outputPath"`
	DownloadForced          bool                   `json:"downloadForced"`
	ID                      int64                  `json:"id"`
	ErrorMessage            string                 `json:"errorMessage"`
}

// GetQueue returns a single page from the Lidarr Queue (processing, but not yet imported).
// If you need control over the page, use lidarr.GetQueuePage().
// This function simply returns the number of queue records desired,
// up to the number of records present in the application.
// It grabs records in (paginated) batches of perPage, and concatenates
// them into one list.  Passing zero for records will return all of them.
func (l *Lidarr) GetQueue(records, perPage int) (*Queue, error) {
	return l.GetQueueContext(context.Background(), records, perPage)
}

// GetQueueContext returns a single page from the Lidarr Queue (processing, but not yet imported).
func (l *Lidarr) GetQueueContext(ctx context.Context, records, perPage int) (*Queue, error) {
	queue := &Queue{Records: []*QueueRecord{}}
	perPage = starr.SetPerPage(records, perPage)

	for page := 1; ; page++ {
		curr, err := l.GetQueuePageContext(ctx, &starr.PageReq{PageSize: perPage, Page: page})
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

// GetQueuePage returns a single page from the Lidarr Queue.
// The page size and number is configurable with the input request parameters.
func (l *Lidarr) GetQueuePage(params *starr.PageReq) (*Queue, error) {
	return l.GetQueuePageContext(context.Background(), params)
}

// GetQueuePageContext returns a single page from the Lidarr Queue.
// The page size and number is configurable with the input request parameters.
func (l *Lidarr) GetQueuePageContext(ctx context.Context, params *starr.PageReq) (*Queue, error) {
	var output Queue

	params.CheckSet("sortKey", "timeleft")
	params.CheckSet("includeUnknownArtistItems", "true")

	req := starr.Request{URI: bpQueue, Query: params.Params()}
	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteQueue deletes an item from the Activity Queue.
func (l *Lidarr) DeleteQueue(queueID int64, opts *starr.QueueDeleteOpts) error {
	return l.DeleteQueueContext(context.Background(), queueID, opts)
}

// DeleteQueueContext deletes an item from the Activity Queue.
func (l *Lidarr) DeleteQueueContext(ctx context.Context, queueID int64, opts *starr.QueueDeleteOpts) error {
	req := starr.Request{URI: path.Join(bpQueue, fmt.Sprint(queueID)), Query: opts.Values()}
	if err := l.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
