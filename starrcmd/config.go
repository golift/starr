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
	// ErrNoEventFound is returned if an event type is not found.
	// This should only happen when testing and you forget a variable.
	ErrNoEventFound = fmt.Errorf("no eventType environment variable found")
)

// DateFormat matches the date output from all five starr apps. Hopefully it doesn't change!
const DateFormat = "1/2/2006 3:04:05 PM"

// Event is a custom type to hold our EventType.
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

// CmdEvent holds the current event type and the app that triggered it.
// Get one of these by calling New().
type CmdEvent struct {
	App  starr.App
	Type Event
}

// New returns the current Event and Application it's from, or an error if the type doesn't exist.
// When running from a Starr App Custom Script this should not return an error.
func New() (*CmdEvent, error) {
	for _, cmdEvent := range []*CmdEvent{
		{starr.Radarr, Event(os.Getenv("radarr_eventtype"))},
		{starr.Sonarr, Event(os.Getenv("sonarr_eventtype"))},
		{starr.Lidarr, Event(os.Getenv("lidarr_eventtype"))},
		{starr.Readarr, Event(os.Getenv("readarr_eventtype"))},
		{starr.Prowlarr, Event(os.Getenv("prowlarr_eventtype"))},
	} {
		if cmdEvent.Type != "" {
			return cmdEvent, nil
		}
	}

	return nil, ErrNoEventFound
}

// NewMust returns a command event without returning an error. It will panic if the event does not exist.
// When running from a Starr App Custom Script this should not panic.
func NewMust() *CmdEvent {
	cmdEvent, _ := New()
	if cmdEvent == nil {
		panic(ErrNoEventFound)
	}

	return cmdEvent
}

// NewMustNoPanic returns a command event and allows your code to handle the problem.
// When running from a Starr App Custom Script this should always return a proper event.
func NewMustNoPanic() *CmdEvent {
	cmdEvent, _ := New()
	if cmdEvent == nil {
		return &CmdEvent{}
	}

	return cmdEvent
}
