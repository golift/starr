# Starr

[![GoDoc](https://godoc.org/golift.io/starr/svc?status.svg)](https://pkg.go.dev/golift.io/starr)
[![Go Report Card](https://goreportcard.com/badge/golift.io/starr)](https://goreportcard.com/report/golift.io/rotatorr)
[![MIT License](http://img.shields.io/:license-mit-blue.svg)](https://github.com/golift/starr/blob/master/LICENSE)
[![travis](https://travis-ci.org/golift/starr.svg?branch=main "Travis Tests")](https://travis-ci.org/golift/starr)
[![discord](https://badgen.net/badge/icon/Discord?color=0011ff&label&icon=https://simpleicons.now.sh/discord/eee "GoLift Discord")](https://golift.io/discord)

### Another way to say `*arr`

 **Go library to interact with APIs in all the Starr apps.**

-   [Lidarr](http://lidarr.audio)
-   [Sonarr](http://sonarr.tv)
-   [Radarr](http://radarr.video)
-   [Readarr](http://readarr.com)

This library is slowly updated as new methods are needed or requested. If you have
specific needs this library doesn't currently meet, but should or could, please
[let me know](https://github.com/golift/starr/issues/new)!

This library is currently in use by:

-   [Unpackerr](https://github.com/davidnewhall/unpackerr/)
-   [DiscordNotifier-Client](https://github.com/Go-Lift-TV/discordnotifier-client/)

# Usage

Get it:
```shell
go get -u golift.io/starr
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
