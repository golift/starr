package radarr_test

/*
import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golift.io/starr"
	"golift.io/starr/radarr"
)

var testMovieJSON = `{
	  "title": "Beast",
	  "originalTitle": "Beast",
	  "originalLanguage": {
		"id": 1,
		"name": "English"
	  },
	  "alternateTitles": [
		{
		  "sourceType": "tmdb",
		  "movieMetadataId": 2671,
		  "title": "Zvērs",
		  "sourceId": 0,
		  "votes": 0,
		  "voteCount": 0,
		  "language": {
			"id": 1,
			"name": "English"
		  },
		  "id": 18219
		}
	  ],
	  "secondaryYearSourceId": 0,
	  "sortTitle": "beast",
	  "sizeOnDisk": 1796921629,
	  "status": "released",
	  "overview": "...",
	  "inCinemas": "2022-08-11T00:00:00Z",
	  "physicalRelease": "2022-10-11T00:00:00Z",
	  "digitalRelease": "2022-09-09T00:00:00Z",
	  "images": [
		{
		  "coverType": "fanart",
		  "url": "https://image.tmdb.org/t/p/original/8TUb2U9GN3PonbXAQ1FBcJ4XeXu.jpg"
		}
	  ],
	  "website": "https://www.beastmovie.com",
	  "year": 2022,
	  "hasFile": true,
	  "youTubeTrailerId": "oQMc7Sq36mI",
	  "studio": "Universal Pictures",
	  "path": "/movies/Beast (2022)",
	  "qualityProfileId": 4,
	  "monitored": true,
	  "minimumAvailability": "announced",
	  "isAvailable": true,
	  "folderName": "/movies/Beast (2022)",
	  "runtime": 93,
	  "cleanTitle": "beast",
	  "imdbId": "tt13223398",
	  "tmdbId": 760741,
	  "titleSlug": "760741",
	  "certification": "R",
	  "genres": [
		"Thriller"
	  ],
	  "tags": [],
	  "added": "2022-08-30T08:27:15Z",
	  "ratings": {
		"rottenTomatoes": {
		  "votes": 0,
		  "value": 69,
		  "type": "user"
		}
	  },
	  "movieFile": {},
	  "popularity": 2240.269,
	  "id": 2295
	}`

// This matches the json above.
var testMovieStruct = radarr.Movie{
	ID:    2295,
	Title: "Beast",
	OriginalLanguage: &starr.Value{
		ID:   1,
		Name: "English",
	},
	AlternateTitles: []*radarr.AlternativeTitle{{
		SourceType:      "tmdb",
		MovieMetadataID: 2671,
		Title:           "Zvērs",
		SourceID:        0,
		Votes:           0,
		VoteCount:       0,
		Language: &starr.Value{
			ID:   1,
			Name: "English",
		},
		ID: 18219,
	}},
	Path:             "/movies/Beast (2022)",
	QualityProfileID: 4,
	TmdbID:           760741,
	OriginalTitle:    "Beast",
	Popularity:       2240.269,

	SecondaryYearSourceID: 0,
	SortTitle:             "beast",
	SizeOnDisk:            1796921629,
	Status:                "released",
	Overview:              "...",
	InCinemas:             time.Date(2022, 8, 11, 0, 0, 0, 0, time.UTC),
	PhysicalRelease:       time.Date(2022, 10, 11, 0, 0, 0, 0, time.UTC),
	DigitalRelease:        time.Date(2022, 9, 9, 0, 0, 0, 0, time.UTC),
	Images: []*starr.Image{{
		CoverType: "fanart",
		URL:       "https://image.tmdb.org/t/p/original/8TUb2U9GN3PonbXAQ1FBcJ4XeXu.jpg",
	}},
	Website:             "https://www.beastmovie.com",
	Year:                2022,
	HasFile:             true,
	YouTubeTrailerID:    "oQMc7Sq36mI",
	Studio:              "Universal Pictures",
	Monitored:           true,
	MinimumAvailability: radarr.AvailabilityAnnounced,
	IsAvailable:         true,
	FolderName:          "/movies/Beast (2022)",
	Runtime:             93,
	CleanTitle:          "beast",
	ImdbID:              "tt13223398",
	TitleSlug:           "760741",
	Certification:       "R",
	Genres:              []string{"Thriller"},
	Tags:                []int{},
	Added:               time.Date(2022, 8, 30, 8, 27, 15, 0, time.UTC),
	Ratings:             map[string]starr.Ratings{"rottenTomatoes": {Votes: 0, Value: 69, Type: "user"}},
	MovieFile:           &radarr.MovieFile{}, // this could get tested..
}

func TestGetCalendar(t *testing.T) {
	t.Parallel()

	tests := []*starr.TestMockData{
		{
			Name: "200",
			ExpectedPath: "/api/v3/calendar" +
				"?end=2020-02-20T04%3A20%3A20.000Z" +
				"&start=2020-02-20T04%3A20%3A20.000Z" +
				"&unmonitored=true",
			ResponseStatus: http.StatusOK,
			ResponseBody:   `[` + testMovieJSON + `]`,
			WithRequest: radarr.Calendar{
				Start:       time.Unix(1582172420, 0),
				End:         time.Unix(1582172420, 0),
				Unmonitored: starr.True(),
			},
			WithError:      nil,
			ExpectedMethod: http.MethodGet,
			WithResponse:   []*radarr.Movie{&testMovieStruct},
		},
		{
			Name: "404",
			ExpectedPath: "/api/v3/calendar" +
				"?end=2020-02-20T04%3A20%3A20.000Z" +
				"&start=2020-02-20T04%3A20%3A20.000Z" +
				"&unmonitored=true",
			ResponseStatus: http.StatusNotFound,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      starr.ErrInvalidStatusCode,
			ExpectedMethod: http.MethodGet,
			WithRequest: radarr.Calendar{
				Start:       time.Unix(1582172420, 0),
				End:         time.Unix(1582172420, 0),
				Unmonitored: starr.True(),
			},
			WithResponse: []*radarr.Movie(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := radarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetCalendar(test.WithRequest.(radarr.Calendar))
			assert.ErrorIs(t, err, test.WithError, "the wrong error was returned")
			assert.EqualValues(t, test.WithResponse, output, "make sure ResponseBody and WithResponse are a match")
		})
	}
}
/**/
