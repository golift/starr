package readarr

import (
	"context"
	"fmt"
	"time"

	"golift.io/starr"
)

// Queue is the /api/v1/queue endpoint.
type Queue struct {
	Page          int            `json:"page"`
	PageSize      int            `json:"pageSize"`
	SortKey       string         `json:"sortKey"`
	SortDirection string         `json:"sortDirection"`
	TotalRecords  int            `json:"totalRecords"`
	Records       []*QueueRecord `json:"records"`
}

// QueueRecord is a book from the queue API path.
type QueueRecord struct {
	AuthorID                int64                  `json:"authorId"`
	BookID                  int64                  `json:"bookId"`
	Quality                 *starr.Quality         `json:"quality"`
	Size                    float64                `json:"size"`
	Title                   string                 `json:"title"`
	Sizeleft                float64                `json:"sizeleft"`
	Timeleft                string                 `json:"timeleft"`
	EstimatedCompletionTime time.Time              `json:"estimatedCompletionTime"`
	Status                  string                 `json:"status"`
	TrackedDownloadStatus   string                 `json:"trackedDownloadStatus,omitempty"`
	TrackedDownloadState    string                 `json:"trackedDownloadState,omitempty"`
	StatusMessages          []*starr.StatusMessage `json:"statusMessages,omitempty"`
	DownloadID              string                 `json:"downloadId,omitempty"`
	Protocol                string                 `json:"protocol"`
	DownloadClient          string                 `json:"downloadClient,omitempty"`
	Indexer                 string                 `json:"indexer"`
	OutputPath              string                 `json:"outputPath,omitempty"`
	DownloadForced          bool                   `json:"downloadForced"`
	ID                      int64                  `json:"id"`
	ErrorMessage            string                 `json:"errorMessage"`
}

// GetQueue returns a single page from the Readarr Queue (processing, but not yet imported).
// WARNING: 12/30/2021 - this method changed.
// If you need control over the page, use readarr.GetQueuePage().
// This function simply returns the number of queue records desired,
// up to the number of records present in the application.
// It grabs records in (paginated) batches of perPage, and concatenates
// them into one list.  Passing zero for records will return all of them.
func (r *Readarr) GetQueue(records, perPage int) (*Queue, error) {
	return r.GetQueueContext(context.Background(), records, perPage)
}

func (r *Readarr) GetQueueContext(ctx context.Context, records, perPage int) (*Queue, error) {
	queue := &Queue{Records: []*QueueRecord{}}
	perPage = starr.SetPerPage(records, perPage)

	for page := 1; ; page++ {
		curr, err := r.GetQueuePageContext(ctx, &starr.Req{PageSize: perPage, Page: page})
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

// GetQueuePage returns a single page from the Readarr Queue.
// The page size and number is configurable with the input request parameters.
func (r *Readarr) GetQueuePage(params *starr.Req) (*Queue, error) {
	return r.GetQueuePageContext(context.Background(), params)
}

func (r *Readarr) GetQueuePageContext(ctx context.Context, params *starr.Req) (*Queue, error) {
	var queue Queue

	params.CheckSet("sortKey", "timeleft")
	params.CheckSet("includeUnknownAuthorItems", "true")

	_, err := r.GetInto(ctx, "v1/queue", params.Params(), &queue)
	if err != nil {
		return nil, fmt.Errorf("api.Get(queue): %w", err)
	}

	return &queue, nil
}
