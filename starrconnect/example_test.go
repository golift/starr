package starrconnect_test

import (
	"fmt"
	"net/http"
	"time"

	"golift.io/starr/starrconnect"
)

func ExampleSonarrHandler() {
	handler := &starrconnect.SonarrHandler{
		OnGrab: func(grab *starrconnect.SonarrGrab) error {
			fmt.Println("grab", grab.Series.Title)
			// Put your logic here.
			return nil
		},
		OnDownload: func(download *starrconnect.SonarrDownload) error {
			fmt.Println("download", download.Series.Title)
			// Put your logic here.
			return nil
		},
	}

	_ = (&http.Server{
		Handler:           handler,
		ReadHeaderTimeout: 10 * time.Second,
	})

	fmt.Println("ok")
	// Output: ok
}

func ExampleParseRadarr() {
	body := []byte(`{"eventType":"MovieAdded","instanceName":"Radarr","applicationUrl":"http://localhost:7878","movie":{"id":1,"title":"Example","year":2020,"tmdbId":1,"imdbId":"","overview":"","genres":[],"images":[],"tags":[]},"addMethod":"manual"}`)

	envelope, err := starrconnect.ParseRadarr(body)
	if err != nil {
		panic(err)
	}

	added, err := envelope.GetMovieAdded()
	if err != nil {
		panic(err)
	}

	fmt.Println(added.Movie.Title)
	// Output: Example
}
