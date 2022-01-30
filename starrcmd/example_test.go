//nolint:exhaustive
package starrcmd_test

import (
	"fmt"

	"golift.io/starr"
	"golift.io/starr/starrcmd"
)

/* This example assumes you want to handle custom script command hooks for all applications.
   Trim it to what you need; this is just an example.
*/

// This is an example main() function that uses the golft.io/starr/starrcmd module.
func ExampleMain() {
	// Do your own configuration input here.
	// c := &Config{stuff}
	switch starrcmd.Application {
	case starr.Radarr:
		DoRadarr()
	case starr.Sonarr:
		DoSonarr()
	case starr.Readarr:
		DoReadarr()
	case starr.Lidarr:
		DoLidarr()
	case starr.Prowlarr:
		DoProwlarr()
	}
}

// DoRadarr handles any Radarr event.
func DoRadarr() { //nolint:cyclop
	fmt.Println("Processing Radarr Event: ", starrcmd.EventType)

	switch starrcmd.EventType {
	case starrcmd.EventGrab:
		grab, err := starrcmd.GetRadarrGrab()
		if err != nil {
			panic(err)
		}

		fmt.Println(grab.Title)
	case starrcmd.EventApplicationUpdate:
		update, err := starrcmd.GetRadarrApplicationUpdate()
		if err != nil {
			panic(err)
		}

		fmt.Println(update.Message)
	case starrcmd.EventDownload:
		download, err := starrcmd.GetRadarrDownload()
		if err != nil {
			panic(err)
		}

		fmt.Println(download.Title)
	case starrcmd.EventHealthIssue:
		health, err := starrcmd.GetRadarrHealthIssue()
		if err != nil {
			panic(err)
		}

		fmt.Println(health.IssueType, health.Message)
	case starrcmd.EventMovieFileDelete:
		movie, err := starrcmd.GetRadarrMovieFileDelete()
		if err != nil {
			panic(err)
		}

		fmt.Println(movie.Title, movie.Path)
	case starrcmd.EventTest:
		// nothing, it's useless
	default:
		fmt.Println("Ignored Radarr Event: ", starrcmd.EventType)
	}
}

/* The following procedures are just more examples. They're left empty on purpose. */

// DoSonarr handles any Sonarr event.
func DoSonarr() {
	fmt.Println("Processing Sonarr Event: ", starrcmd.EventType)

	switch starrcmd.EventType {
	case starrcmd.EventGrab:
	case starrcmd.EventApplicationUpdate:
	case starrcmd.EventDownload:
	case starrcmd.EventHealthIssue:
	default:
		fmt.Println("Ignored Sonarr Event: ", starrcmd.EventType)
	}
}

// DoLidarr handles any Lidarr event.
func DoLidarr() {
	fmt.Println("Processing Lidarr Event: ", starrcmd.EventType)

	switch starrcmd.EventType {
	case starrcmd.EventGrab:
	case starrcmd.EventApplicationUpdate:
	case starrcmd.EventDownload:
	case starrcmd.EventHealthIssue:
	default:
		fmt.Println("Ignored Lidarr Event: ", starrcmd.EventType)
	}
}

// DoReadarr handles any Readarr event.
func DoReadarr() {
	fmt.Println("Processing Readarr Event: ", starrcmd.EventType)

	switch starrcmd.EventType {
	case starrcmd.EventGrab:
	case starrcmd.EventApplicationUpdate:
	case starrcmd.EventDownload:
	case starrcmd.EventHealthIssue:
	default:
		fmt.Println("Ignored Readarr Event: ", starrcmd.EventType)
	}
}

// DoProwlarr handles any Prowlarr event.
func DoProwlarr() {
	fmt.Println("Processing Prowlarr Event: ", starrcmd.EventType)

	switch starrcmd.EventType {
	case starrcmd.EventApplicationUpdate:
	case starrcmd.EventHealthIssue:
	default:
		fmt.Println("Ignored Prowlarr Event: ", starrcmd.EventType)
	}
}
