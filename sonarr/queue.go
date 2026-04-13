package sonarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
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
	HasPostImportCategory   bool                   `json:"downloadClientHasPostImportCategory"`
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
	Protocol                starr.Protocol         `json:"protocol"`
	DownloadClient          string                 `json:"downloadClient"`
	Indexer                 string                 `json:"indexer"`
	OutputPath              string                 `json:"outputPath"`
	ErrorMessage            string                 `json:"errorMessage"`
}

// QueueStatus is the aggregate queue status from /api/v3/queue/status.
type QueueStatus struct {
	ID              int  `json:"id,omitempty"`
	TotalCount      int  `json:"totalCount,omitempty"`
	Count           int  `json:"count,omitempty"`
	UnknownCount    int  `json:"unknownCount,omitempty"`
	Errors          bool `json:"errors"`
	Warnings        bool `json:"warnings"`
	UnknownErrors   bool `json:"unknownErrors"`
	UnknownWarnings bool `json:"unknownWarnings"`
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

// GetQueueContext returns a single page from the Sonarr Queue (processing, but not yet imported).
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
	req := starr.Request{URI: path.Join(bpQueue, starr.Str(queueID)), Query: opts.Values()}
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

	var output any // any ok

	req := starr.Request{URI: path.Join(bpQueue, "grab", "bulk"), Body: &body}
	if err := s.PostInto(ctx, req, &output); err != nil {
		return fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return nil
}

// GetQueueDetails returns the raw JSON array from /api/v3/queue/details.
// Unmarshal into your own structs if you need typed fields beyond the main queue list.
func (s *Sonarr) GetQueueDetails(query url.Values) ([]byte, error) {
	return s.GetQueueDetailsContext(context.Background(), query)
}

// GetQueueDetailsContext returns the raw JSON array from /api/v3/queue/details.
func (s *Sonarr) GetQueueDetailsContext(ctx context.Context, query url.Values) ([]byte, error) {
	uri := starr.SetAPIPath(path.Join(bpQueue, "details"))
	req := starr.Request{URI: uri, Query: query}

	resp, err := s.Get(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body from %s: %w", uri, err)
	}

	return body, nil
}

// GetQueueStatus returns aggregate queue status.
func (s *Sonarr) GetQueueStatus() (*QueueStatus, error) {
	return s.GetQueueStatusContext(context.Background())
}

// GetQueueStatusContext returns aggregate queue status.
func (s *Sonarr) GetQueueStatusContext(ctx context.Context) (*QueueStatus, error) {
	var output QueueStatus

	req := starr.Request{URI: path.Join(bpQueue, "status")}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteQueueBulk removes multiple queue items.
func (s *Sonarr) DeleteQueueBulk(ids []int64, opts *starr.QueueDeleteOpts) error {
	return s.DeleteQueueBulkContext(context.Background(), ids, opts)
}

// DeleteQueueBulkContext removes multiple queue items.
func (s *Sonarr) DeleteQueueBulkContext(ctx context.Context, ids []int64, opts *starr.QueueDeleteOpts) error {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(struct {
		IDs []int64 `json:"ids"`
	}{IDs: ids}); err != nil {
		return fmt.Errorf("json.Marshal(%s): %w", bpQueue, err)
	}

	var params url.Values
	if opts != nil {
		params = opts.Values()
	} else {
		params = (&starr.QueueDeleteOpts{}).Values()
	}

	uri := starr.SetAPIPath(path.Join(bpQueue, "bulk"))
	req := starr.Request{URI: uri, Body: &body, Query: params}

	resp, err := s.Delete(ctx, req)
	if err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}
	defer resp.Body.Close()

	_, _ = io.ReadAll(resp.Body)

	return nil
}

// QueueGrabOne tells Sonarr to grab a single delayed queue item.
func (s *Sonarr) QueueGrabOne(queueID int64) error {
	return s.QueueGrabOneContext(context.Background(), queueID)
}

// QueueGrabOneContext tells Sonarr to grab a single delayed queue item.
func (s *Sonarr) QueueGrabOneContext(ctx context.Context, queueID int64) error {
	var output any

	req := starr.Request{URI: path.Join(bpQueue, "grab", starr.Str(queueID))}
	if err := s.PostInto(ctx, req, &output); err != nil {
		return fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return nil
}
