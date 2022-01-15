package lidarr

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"

	"golift.io/starr"
)

// Lidarr contains all the methods to interact with a Lidarr server.
type Lidarr struct {
	starr.APIer
}

// New returns a Lidarr object used to interact with the Lidarr API.
func New(config *starr.Config) *Lidarr {
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

	return &Lidarr{APIer: config}
}

// GetQueue returns a single page from the Lidarr Queue (processing, but not yet imported).
// WARNING: 12/30/2021 - this method changed.
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
		curr, err := l.GetQueuePageContext(ctx, &starr.Req{PageSize: perPage, Page: page})
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
func (l *Lidarr) GetQueuePage(params *starr.Req) (*Queue, error) {
	return l.GetQueuePageContext(context.Background(), params)
}

// GetQueuePageContext returns a single page from the Lidarr Queue.
// The page size and number is configurable with the input request parameters.
func (l *Lidarr) GetQueuePageContext(ctx context.Context, params *starr.Req) (*Queue, error) {
	var queue Queue

	params.CheckSet("sortKey", "timeleft")
	params.CheckSet("includeUnknownArtistItems", "true")

	err := l.GetInto(ctx, "v1/queue", params.Params(), &queue)
	if err != nil {
		return nil, fmt.Errorf("api.Get(queue): %w", err)
	}

	return &queue, nil
}

// GetQualityDefinition returns the Quality Definitions.
func (l *Lidarr) GetQualityDefinition() ([]*QualityDefinition, error) {
	return l.GetQualityDefinitionContext(context.Background())
}

// GetQualityDefinitionContext returns the Quality Definitions.
func (l *Lidarr) GetQualityDefinitionContext(ctx context.Context) ([]*QualityDefinition, error) {
	var definition []*QualityDefinition

	err := l.GetInto(ctx, "v1/qualitydefinition", nil, &definition)
	if err != nil {
		return nil, fmt.Errorf("api.Get(qualitydefinition): %w", err)
	}

	return definition, nil
}

// GetRootFolders returns all configured root folders.
func (l *Lidarr) GetRootFolders() ([]*RootFolder, error) {
	return l.GetRootFoldersContext(context.Background())
}

// GetRootFoldersContext returns all configured root folders.
func (l *Lidarr) GetRootFoldersContext(ctx context.Context) ([]*RootFolder, error) {
	var folders []*RootFolder

	err := l.GetInto(ctx, "v1/rootFolder", nil, &folders)
	if err != nil {
		return nil, fmt.Errorf("api.Get(rootFolder): %w", err)
	}

	return folders, nil
}

// GetMetadataProfiles returns the metadata profiles.
func (l *Lidarr) GetMetadataProfiles() ([]*MetadataProfile, error) {
	return l.GetMetadataProfilesContext(context.Background())
}

// GetMetadataProfilesContext returns the metadata profiles.
func (l *Lidarr) GetMetadataProfilesContext(ctx context.Context) ([]*MetadataProfile, error) {
	var profiles []*MetadataProfile

	err := l.GetInto(ctx, "v1/metadataprofile", nil, &profiles)
	if err != nil {
		return nil, fmt.Errorf("api.Get(metadataprofile): %w", err)
	}

	return profiles, nil
}
