//nolint:exhaustive,testableexamples
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
func Example() {
	// Do your own configuration input here.
	cmd, err := starrcmd.New()
	if err != nil {
		panic(err)
	}

	switch cmd.App {
	case starr.Radarr:
		DoRadarr(cmd)
	case starr.Sonarr:
		DoSonarr(cmd)
	case starr.Readarr:
		DoReadarr(cmd)
	case starr.Lidarr:
		DoLidarr(cmd)
	case starr.Prowlarr:
		DoProwlarr(cmd)
	}
}

// DoRadarr handles any Radarr event.
func DoRadarr(cmd *starrcmd.CmdEvent) { //nolint:cyclop
	fmt.Println("Processing Radarr Event: ", cmd.Type)

	switch cmd.Type {
	case starrcmd.EventGrab:
		grab, err := cmd.GetRadarrGrab()
		if err != nil {
			panic(err)
		}

		fmt.Println(grab.Title)
	case starrcmd.EventApplicationUpdate:
		update, err := cmd.GetRadarrApplicationUpdate()
		if err != nil {
			panic(err)
		}

		fmt.Println(update.Message)
	case starrcmd.EventDownload:
		download, err := cmd.GetRadarrDownload()
		if err != nil {
			panic(err)
		}

		fmt.Println(download.Title)
	case starrcmd.EventHealthIssue:
		health, err := cmd.GetRadarrHealthIssue()
		if err != nil {
			panic(err)
		}

		fmt.Println(health.IssueType, health.Message)
	case starrcmd.EventMovieFileDelete:
		movie, err := cmd.GetRadarrMovieFileDelete()
		if err != nil {
			panic(err)
		}

		fmt.Println(movie.Title, movie.Path)
	case starrcmd.EventTest:
		// nothing, it's useless
	default:
		fmt.Println("Ignored Radarr Event: ", cmd.Type)
	}
}

/* The following procedures are just more examples. They're left empty on purpose. */

// DoSonarr handles any Sonarr event.
func DoSonarr(command *starrcmd.CmdEvent) {
	fmt.Println("Processing Sonarr Event: ", command.Type)

	switch command.Type {
	case starrcmd.EventGrab:
	case starrcmd.EventApplicationUpdate:
	case starrcmd.EventDownload:
	case starrcmd.EventHealthIssue:
	default:
		fmt.Println("Ignored Sonarr Event: ", command.Type)
	}
}

// DoLidarr handles any Lidarr event.
func DoLidarr(command *starrcmd.CmdEvent) {
	fmt.Println("Processing Lidarr Event: ", command.Type)

	switch command.Type {
	case starrcmd.EventGrab:
	case starrcmd.EventApplicationUpdate:
	case starrcmd.EventDownload:
	case starrcmd.EventHealthIssue:
	default:
		fmt.Println("Ignored Lidarr Event: ", command.Type)
	}
}

// DoReadarr handles any Readarr event.
func DoReadarr(command *starrcmd.CmdEvent) {
	fmt.Println("Processing Readarr Event: ", command.Type)

	switch command.Type {
	case starrcmd.EventGrab:
	case starrcmd.EventApplicationUpdate:
	case starrcmd.EventDownload:
	case starrcmd.EventHealthIssue:
	default:
		fmt.Println("Ignored Readarr Event: ", command.Type)
	}
}

// DoProwlarr handles any Prowlarr event.
func DoProwlarr(command *starrcmd.CmdEvent) {
	fmt.Println("Processing Prowlarr Event: ", command.Type)

	switch command.Type {
	case starrcmd.EventApplicationUpdate:
	case starrcmd.EventHealthIssue:
	default:
		fmt.Println("Ignored Prowlarr Event: ", command.Type)
	}
}
