package sonarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"
	"time"

	"golift.io/starr"
)

const bpQueue = APIver + "/queue"

// Queue is the /api/v3/queue endpoint.
type Queue struct {
	Page          int            `json:"page"`
	PageSize      int            `json:"pageSize"`
	SortKey       string         `json:"sortKey"`
	SortDirection string         `json:"sortDirection"`
	TotalRecords  int            `json:"totalRecords"`
	Records       []*QueueRecord `json:"records"`
}

// QueueRecord is part of Queue.
type QueueRecord struct {
	ID                      int64                  `json:"id"`
	SeriesID                int64                  `json:"seriesId"`
	EpisodeID               int64                  `json:"episodeId"`
	Language                *starr.Value           `json:"language"`
	Quality                 *starr.Quality         `json:"quality"`
	Size                    float64                `json:"size"`
	Title                   string                 `json:"title"`
	Sizeleft                float64                `json:"sizeleft"`
	Timeleft                string                 `json:"timeleft"`
	EstimatedCompletionTime time.Time              `json:"estimatedCompletionTime"`
	Status                  string                 `json:"status"`
	TrackedDownloadStatus   string                 `json:"trackedDownloadStatus"`
	TrackedDownloadState    string                 `json:"trackedDownloadState"`
	StatusMessages          []*starr.StatusMessage `json:"statusMessages"`
	DownloadID              string                 `json:"downloadId"`
	Protocol                string                 `json:"protocol"`
	DownloadClient          string                 `json:"downloadClient"`
	Indexer                 string                 `json:"indexer"`
	OutputPath              string                 `json:"outputPath"`
	ErrorMessage            string                 `json:"errorMessage"`
}

// GetQueue returns a single page from the Sonarr Queue (processing, but not yet imported).
// If you need control over the page, use sonarr.GetQueuePage().
// This function simply returns the number of queue records desired,
// up to the number of records present in the application.
// It grabs records in (paginated) batches of perPage, and concatenates
// them into one list.  Passing zero for records will return all of them.
func (s *Sonarr) GetQueue(records, perPage int) (*Queue, error) {
	return s.GetQueueContext(context.Background(), records, perPage)
}

// GetQueue returns a single page from the Sonarr Queue (processing, but not yet imported).
// If you need control over the page, use sonarr.GetQueuePageContext().
func (s *Sonarr) GetQueueContext(ctx context.Context, records, perPage int) (*Queue, error) {
	queue := &Queue{Records: []*QueueRecord{}}
	perPage = starr.SetPerPage(records, perPage)

	for page := 1; ; page++ {
		curr, err := s.GetQueuePageContext(ctx, &starr.PageReq{PageSize: perPage, Page: page})
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
func (s *Sonarr) GetQueuePage(params *starr.PageReq) (*Queue, error) {
	return s.GetQueuePageContext(context.Background(), params)
}

// GetQueuePageContext returns a single page from the Sonarr Queue.
// The page size and number is configurable with the input request parameters.
func (s *Sonarr) GetQueuePageContext(ctx context.Context, params *starr.PageReq) (*Queue, error) {
	var output Queue

	params.CheckSet("sortKey", "timeleft")
	params.CheckSet("includeUnknownSeriesItems", "true")

	req := starr.Request{URI: bpQueue, Query: params.Params()}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteQueue deletes an item from the Activity Queue.
func (s *Sonarr) DeleteQueue(queueID int64, opts *starr.QueueDeleteOpts) error {
	return s.DeleteQueueContext(context.Background(), queueID, opts)
}

// DeleteQueueContext deletes an item from the Activity Queue.
func (s *Sonarr) DeleteQueueContext(ctx context.Context, queueID int64, opts *starr.QueueDeleteOpts) error {
	req := starr.Request{URI: path.Join(bpQueue, fmt.Sprint(queueID)), Query: opts.Values()}
	if err := s.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}

// QueueGrab tells the app to grab an item that's in queue.
// Most often used on items with a delay set from a delay profile.
func (s *Sonarr) QueueGrab(ids ...int64) error {
	return s.QueueGrabContext(context.Background(), ids...)
}

// QueueGrabContext tells the app to grab an item that's in queue, probably set to a delay.
// Most often used on items with a delay set from a delay profile.
func (s *Sonarr) QueueGrabContext(ctx context.Context, ids ...int64) error {
	idList := struct {
		IDs []int64 `json:"ids"`
	}{IDs: ids}

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(idList); err != nil {
		return fmt.Errorf("json.Marshal(%s): %w", bpQueue, err)
	}

	var output interface{} // any ok

	req := starr.Request{URI: path.Join(bpQueue, "grab", "bulk"), Body: &body}
	if err := s.PostInto(ctx, req, &output); err != nil {
		return fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return nil
}
