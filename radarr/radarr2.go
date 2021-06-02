package radarr

import (
	"fmt"
	"net/url"
	"time"

	"golift.io/starr"
)

/* This is all deprecated and will be removed in the future. Switch to v3. */

// GetQueueV2 returns the Radarr Queue (processing, but not yet imported).
func (r *Radarr) GetQueueV2() ([]*QueueV2, error) {
	params := make(url.Values)
	params.Set("sort_by", "timeleft")
	params.Set("order", "asc")

	var queue []*QueueV2
	if err := r.GetInto("queue", params, &queue); err != nil {
		return nil, fmt.Errorf("api.Get(queue): %w", err)
	}

	return queue, nil
}

// QueueV2 is the /api/queue endpoint.
type QueueV2 struct {
	ID                      int64                  `json:"id"`
	Size                    float64                `json:"size"`
	Sizeleft                float64                `json:"sizeleft"`
	EstimatedCompletionTime time.Time              `json:"estimatedCompletionTime"`
	Title                   string                 `json:"title"`
	Timeleft                string                 `json:"timeleft"`
	Status                  string                 `json:"status"`
	TrackedDownloadStatus   string                 `json:"trackedDownloadStatus"`
	DownloadID              string                 `json:"downloadId"`
	Protocol                string                 `json:"protocol"`
	Movie                   *QueueMovieV2          `json:"movie"`
	Quality                 *starr.Quality         `json:"quality"`
	StatusMessages          []*starr.StatusMessage `json:"statusMessages"`
}

// QueueMovie is part of a Movie in the Queue.
type QueueMovieV2 struct {
	Downloaded            bool           `json:"downloaded"`
	HasFile               bool           `json:"hasFile"`
	Monitored             bool           `json:"monitored"`
	IsAvailable           bool           `json:"isAvailable"`
	SecondaryYearSourceID int            `json:"secondaryYearSourceId"`
	Year                  int            `json:"year"`
	ProfileID             int64          `json:"profileId"`
	Runtime               int            `json:"runtime"`
	QualityProfileID      int64          `json:"qualityProfileId"`
	ID                    int64          `json:"id"`
	TmdbID                int64          `json:"tmdbId"`
	SizeOnDisk            int64          `json:"sizeOnDisk"`
	InCinemas             time.Time      `json:"inCinemas"`
	PhysicalRelease       time.Time      `json:"physicalRelease"`
	LastInfoSync          time.Time      `json:"lastInfoSync"`
	Added                 time.Time      `json:"added"`
	Title                 string         `json:"title"`
	SortTitle             string         `json:"sortTitle"`
	Status                string         `json:"status"`
	Overview              string         `json:"overview"`
	Website               string         `json:"website"`
	YouTubeTrailerID      string         `json:"youTubeTrailerId"`
	Studio                string         `json:"studio"`
	Path                  string         `json:"path"`
	PathState             string         `json:"pathState"`
	MinimumAvailability   string         `json:"minimumAvailability"`
	FolderName            string         `json:"folderName"`
	CleanTitle            string         `json:"cleanTitle"`
	ImdbID                string         `json:"imdbId"`
	TitleSlug             string         `json:"titleSlug"`
	Genres                []string       `json:"genres"`
	Tags                  []int          `json:"tags"`
	Images                []*starr.Image `json:"images"`
	Ratings               *starr.Ratings `json:"ratings"`
}
