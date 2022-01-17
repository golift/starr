package radarr

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"

	"golift.io/starr"
)

// Radarr contains all the methods to interact with a Radarr server.
type Radarr struct {
	starr.APIer
}

// Filter values are integers. Given names for ease of discovery.
//nolint:lll
// https://github.com/Radarr/Radarr/blob/2bca1a71a2ed5130ea642343cb76250f3bf5bc4e/src/NzbDrone.Core/History/History.cs#L33-L44
const (
	FilterUnknown starr.Filtering = iota
	FilterGrabbed
	_ // 2 is unused
	FilterDownloadFolderImported
	FilterDownloadFailed
	_ // 5 is unused. FilterDeleted
	FilterFileDeleted
	_ // FilterFolderImported // not used yet, 1/17/2022
	FilterRenamed
	FilterIgnored
)

// New returns a Radarr object used to interact with the Radarr API.
func New(config *starr.Config) *Radarr {
	if config.Client == nil {
		//nolint:exhaustivestruct,gosec
		config.Client = &http.Client{
			Timeout: config.Timeout.Duration,
			CheckRedirect: func(r *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: !config.ValidSSL},
			},
		}
	}

	if config.Debugf == nil {
		config.Debugf = func(string, ...interface{}) {}
	}

	return &Radarr{APIer: config}
}

// GetQueue returns a single page from the Radarr Queue (processing, but not yet imported).
// WARNING: 12/30/2021 - this method changed. The second argument no longer
// controls which page is returned, but instead adjusts the pagination size.
// If you need control over the page, use radarr.GetQueuePage().
// This function simply returns the number of queue records desired,
// up to the number of records present in the application.
// It grabs records in (paginated) batches of perPage, and concatenates
// them into one list.  Passing zero for records will return all of them.
func (r *Radarr) GetQueue(records, perPage int) (*Queue, error) {
	return r.GetQueueContext(context.Background(), records, perPage)
}

// GetQueueContext returns a single page from the Radarr Queue (processing, but not yet imported).
func (r *Radarr) GetQueueContext(ctx context.Context, records, perPage int) (*Queue, error) {
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

// GetQueuePage returns a single page from the Radarr Queue.
// The page size and number is configurable with the input request parameters.
func (r *Radarr) GetQueuePage(params *starr.Req) (*Queue, error) {
	return r.GetQueuePageContext(context.Background(), params)
}

// GetQueuePage returns a single page from the Radarr Queue.
// The page size and number is configurable with the input request parameters.
func (r *Radarr) GetQueuePageContext(ctx context.Context, params *starr.Req) (*Queue, error) {
	var queue Queue

	params.CheckSet("sortKey", "timeleft")
	params.CheckSet("includeUnknownMovieItems", "true")

	err := r.GetInto(ctx, "v3/queue", params.Params(), &queue)
	if err != nil {
		return nil, fmt.Errorf("api.Get(queue): %w", err)
	}

	return &queue, nil
}

// GetRootFolders returns all configured root folders.
func (r *Radarr) GetRootFolders() ([]*RootFolder, error) {
	return r.GetRootFoldersContext(context.Background())
}

// GetRootFoldersContext returns all configured root folders.
func (r *Radarr) GetRootFoldersContext(ctx context.Context) ([]*RootFolder, error) {
	var folders []*RootFolder

	err := r.GetInto(ctx, "v3/rootFolder", nil, &folders)
	if err != nil {
		return nil, fmt.Errorf("api.Get(rootFolder): %w", err)
	}

	return folders, nil
}
