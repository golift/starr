//nolint:gochecknoglobals,gochecknoinits,nlreturn
package starrcmd

import (
	"fmt"
	"os"

	"golift.io/starr"
)

/* Notes to future developers of this module:
- Counters should be int.
- IDs should be int64.
- Sizes should be int64 (bytes).
- Avoid uint* int8, int16, int32, float32, or add parsers for them. See golift.io/cnfg.
- Avoid external modules for env parsing; those require custom types.
- Some slices are allowed, add more when needed. See parseSlices().
  - Slices must have a split character. ,, or ,| (usually).
  - Missing the split character will cause a panic() during parsing.
  - Add tests for all methods and data types to catch panics before release.
- The time.Time format is hard coded. If new formats arise, find a way to fix it?
- No time.Duration types exist, but we can write a parser for those if they arise.
*/

var (
	// ErrInvalidEvent is returned if you invoke a procedure for the wrong event.
	ErrInvalidEvent = fmt.Errorf("incorrect event type requested")
	// EventType is set automatically  by init() but can be overridden in your code if needed.
	EventType Event
	// Application is set automatically by init() but can be overridden in your code if needed.
	Application starr.App
)

const DateFormat = "1/2/2006 3:04:05 PM"

// Event is a hard type to hold our EventType.
type Event string

// This list of constants represents all available and existing Event Types for all five Starr apps.
// Lidarr is complete; 1/30/2022.
// Prowlarr is complete; 1/30/2022.
// Radarr is complete; 1/30/2022.
// Readarr is complete; 1/30/2022.
// Sonarr is complete; 1/30/2022.
const (
	EventTest              Event = "Test"              // All Apps, useless
	EventHealthIssue       Event = "HealthIssue"       // All Apps
	EventApplicationUpdate Event = "ApplicationUpdate" // All Apps
	EventGrab              Event = "Grab"              // All Apps except Prowlarr
	EventRename            Event = "Rename"            // All Apps except Prowlarr
	EventDownload          Event = "Download"          // All Apps except Prowlarr/Lidarr
	EventTrackRetag        Event = "TrackRetag"        // Lidarr & Readarr
	EventAlbumDownload     Event = "AlbumDownload"     // Lidarr
	EventMovieFileDelete   Event = "MovieFileDelete"   // Radarr
	EventMovieDelete       Event = "MovieDelete"       // Radarr
	EventBookDelete        Event = "BookDelete"        // Readarr
	EventAuthorDelete      Event = "AuthorDelete"      // Readarr
	EventBookFileDelete    Event = "BookFileDelete"    // Readarr
	EventSeriesDelete      Event = "SeriesDelete"      // Sonarr
	EventEpisodeFileDelete Event = "EpisodeFileDelete" // Sonarr
)

func init() {
	for _, starrApp := range []struct {
		event string
		app   starr.App
	}{
		{os.Getenv("radarr_eventtype"), starr.Radarr},
		{os.Getenv("sonarr_eventtype"), starr.Sonarr},
		{os.Getenv("lidarr_eventtype"), starr.Lidarr},
		{os.Getenv("readarr_eventtype"), starr.Readarr},
		{os.Getenv("prowlarr_eventtype"), starr.Prowlarr},
	} {
		if starrApp.event != "" {
			EventType, Application = Event(starrApp.event), starrApp.app
			return
		}
	}
}
