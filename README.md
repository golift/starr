# Starr

[![GoDoc](https://godoc.org/golift.io/starr/svc?status.svg)](https://pkg.go.dev/golift.io/starr)
[![Go Report Card](https://goreportcard.com/badge/golift.io/starr)](https://goreportcard.com/report/golift.io/starr)
[![MIT License](http://img.shields.io/:license-mit-blue.svg)](https://github.com/golift/starr/blob/main/LICENSE)
[![discord](https://badgen.net/badge/icon/Discord?color=0011ff&label&icon=https://simpleicons.now.sh/discord/eee "GoLift Discord")](https://golift.io/discord)

## The correct way to say `*arr`.

 **Go library to interact with APIs in all the Starr apps.**

-   [Lidarr](http://lidarr.audio) ([over 80 methods](https://pkg.go.dev/golift.io/starr@main/lidarr))
-   [Prowlarr](https://prowlarr.com) ([over 20 methods](https://pkg.go.dev/golift.io/starr@main/prowlarr))
-   [Radarr](http://radarr.video) ([over 100 methods](https://pkg.go.dev/golift.io/starr@main/radarr))
-   [Readarr](http://readarr.com) ([over 70 methods](https://pkg.go.dev/golift.io/starr@main/readarr))
-   [Sonarr](http://sonarr.tv) ([over 100 methods](https://pkg.go.dev/golift.io/starr@main/sonarr))

[Custom Scripts support](https://wiki.servarr.com/radarr/custom-scripts) is also included.
[Check out the types and methods](https://pkg.go.dev/golift.io/starr@main/starrcmd) to get that data.

## One 🌟 To Rule Them All

This library is slowly updated as new methods are needed or requested. If you have
specific needs this library doesn't currently meet, but should or could, please
[let us know](https://github.com/golift/starr/issues/new)!

This library is currently in use by:

-   [Toolbarr](https://github.com/Notifiarr/toolbarr/) (all of it)
-   [Unpackerr](https://github.com/Unpackerr/unpackerr/) (queue only)
-   [Notifiarr](https://github.com/Notifiarr/notifiarr/) (a lot of it)
-   [Checkrr](https://github.com/aetaric/checkrr/)

# Usage

Get it:
```shell
go get golift.io/starr
```

Use it:
```go
import "golift.io/starr"
```

## Example

```go
package main

import (
	"fmt"

	"golift.io/starr"
	"golift.io/starr/lidarr"
)

func main() {
	// Get a starr.Config that can plug into any Starr app.
	// starr.New(apiKey, appURL string, timeout time.Duration)
	c := starr.New("abc1234ahsuyka123jh12", "http://localhost:8686", 0)
	// Lets make a lidarr server with the default starr Config.
	l := lidarr.New(c)

	// In addition to GetSystemStatus, you have things like:
	// * l.GetAlbum(albumID int)
	// * l.GetQualityDefinition()
	// * l.GetQualityProfiles()
	// * l.GetRootFolders()
	// * l.GetQueue(maxRecords int)
	// * l.GetAlbum(albumUUID string)
	// * l.GetArtist(artistUUID string)
	status, err := l.GetSystemStatus()
	if err != nil {
		panic(err)
	}

	fmt.Println(status)
}
```
