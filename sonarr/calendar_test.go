package sonarr_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golift.io/starr"
	"golift.io/starr/sonarr"
	"golift.io/starr/starrtest"
)

var testCalendarJSON = `{
	"seriesId": 1,
	"tvdbId": 8484175,
	"episodeFileId": 0,
	"seasonNumber": 0,
	"episodeNumber": 1,
	"title": "Dead Zone",
	"airDate": "1989-07-15",
	"airDateUtc": "1989-07-14T15:00:00Z",
	"overview": "...",
	"hasFile": false,
	"monitored": false,
	"unverifiedSceneNumbering": false,
	"series": {
	  "title": "Dragon Ball Z",
	  "sortTitle": "dragon ball z",
	  "status": "ended",
	  "ended": true,
	  "overview": "...",
	  "network": "Fuji TV",
	  "images": [
		{
		  "coverType": "fanart",
		  "url": "https://artworks.thetvdb.com/banners/fanart/original/81472-26.jpg"
		}
	  ],
	  "seasons": [
		{
		  "seasonNumber": 0,
		  "monitored": false
		}
	  ],
	  "year": 1989,
	  "path": "/tv/[DBNL] Dragon Ball Z Digitally Remastered",
	  "qualityProfileId": 1,
	  "languageProfileId": 1,
	  "seasonFolder": true,
	  "monitored": false,
	  "useSceneNumbering": true,
	  "runtime": 25,
	  "tvdbId": 81472,
	  "tvRageId": 3373,
	  "tvMazeId": 2103,
	  "firstAired": "1989-04-26T00:00:00Z",
	  "seriesType": "standard",
	  "cleanTitle": "dragonballz",
	  "imdbId": "tt0214341",
	  "titleSlug": "dragon-ball-z",
	  "certification": "TV-PG",
	  "genres": [
		"Action",
		"Adventure",
		"Animation",
		"Anime",
		"Fantasy",
		"Science Fiction",
		"Thriller"
	  ],
	  "tags": [],
	  "added": "2018-04-19T08:46:03.367045Z",
	  "ratings": {
		"votes": 4659,
		"value": 8.7
	  },
	  "id": 1
	},
	"images": [],
	"id": 1
  }
  `

// This matches the json above.
var testCalendarStruct = sonarr.Episode{
	ID:                       1,
	SeriesID:                 1,
	TvdbID:                   8484175,
	EpisodeFileID:            0,
	SeasonNumber:             0,
	EpisodeNumber:            1,
	Title:                    "Dead Zone",
	AirDate:                  "1989-07-15",
	AirDateUtc:               time.Date(1989, 7, 14, 15, 0, 0, 0, time.UTC),
	Overview:                 "...",
	HasFile:                  false,
	Monitored:                false,
	UnverifiedSceneNumbering: false,
	Series: &sonarr.Series{
		Ended:             true,
		Monitored:         false,
		SeasonFolder:      true,
		UseSceneNumbering: true,
		Runtime:           25,
		Year:              1989,
		ID:                1,
		LanguageProfileID: 1,
		QualityProfileID:  1,
		TvdbID:            81472,
		TvMazeID:          2103,
		TvRageID:          3373,
		AirTime:           "",
		Certification:     "TV-PG",
		CleanTitle:        "dragonballz",
		ImdbID:            "tt0214341",
		Network:           "Fuji TV",
		Overview:          "...",
		Path:              "/tv/[DBNL] Dragon Ball Z Digitally Remastered",
		SeriesType:        "standard",
		SortTitle:         "dragon ball z",
		Status:            "ended",
		Title:             "Dragon Ball Z",
		TitleSlug:         "dragon-ball-z",
		Added:             time.Date(2018, 4, 19, 8, 46, 3, 367045000, time.UTC),
		FirstAired:        time.Date(1989, 4, 26, 0, 0, 0, 0, time.UTC),
		Ratings: &starr.Ratings{
			Votes: 4659,
			Value: 8.7,
		},
		Tags:            []int{},
		Genres:          []string{"Action", "Adventure", "Animation", "Anime", "Fantasy", "Science Fiction", "Thriller"},
		AlternateTitles: []*sonarr.AlternateTitle(nil),
		Seasons:         []*sonarr.Season{{SeasonNumber: 0, Monitored: false}},
		Images: []*starr.Image{
			{CoverType: "fanart", URL: "https://artworks.thetvdb.com/banners/fanart/original/81472-26.jpg"},
		},
	},
	Images: []*starr.Image{},
}

func TestGetCalendar(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name: "200",
			ExpectedPath: "/api/v3/calendar" +
				"?end=2020-02-20T04%3A20%3A20.000Z" +
				"&includeEpisodeFile=false" +
				"&includeEpisodeImages=false" +
				"&includeSeries=true" +
				"&start=2020-02-20T04%3A20%3A20.000Z" +
				"&unmonitored=true",
			ResponseStatus: http.StatusOK,
			ResponseBody:   `[` + testCalendarJSON + `]`,
			WithRequest: sonarr.Calendar{
				Start:                time.Unix(1582172420, 0),
				End:                  time.Unix(1582172420, 0),
				Unmonitored:          true,
				IncludeSeries:        true,
				IncludeEpisodeFile:   false,
				IncludeEpisodeImages: false,
			},
			WithError:      nil,
			ExpectedMethod: http.MethodGet,
			WithResponse:   []*sonarr.Episode{&testCalendarStruct},
		},
		{
			Name: "404",
			ExpectedPath: "/api/v3/calendar" +
				"?end=2020-02-20T04%3A20%3A20.000Z" +
				"&includeEpisodeFile=false" +
				"&includeEpisodeImages=false" +
				"&includeSeries=false" +
				"&start=2020-02-20T04%3A20%3A20.000Z" +
				"&unmonitored=true",
			ResponseStatus: http.StatusNotFound,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      &starr.ReqError{Code: http.StatusNotFound},
			ExpectedMethod: http.MethodGet,
			WithRequest: sonarr.Calendar{
				Start:       time.Unix(1582172420, 0),
				End:         time.Unix(1582172420, 0),
				Unmonitored: true,
			},
			WithResponse: []*sonarr.Episode(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetCalendar(test.WithRequest.(sonarr.Calendar))
			assert.ErrorIs(t, err, test.WithError, "the wrong error was returned")
			assert.EqualValues(t, test.WithResponse, output, "make sure ResponseBody and WithResponse are a match")
		})
	}
}

func TestGetCalendarID(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   "/api/v3/calendar/1",
			ResponseStatus: http.StatusOK,
			ResponseBody:   testCalendarJSON,
			WithRequest:    int64(1),
			WithError:      nil,
			ExpectedMethod: http.MethodGet,
			WithResponse:   &testCalendarStruct,
		},
		{
			Name:           "404",
			ExpectedPath:   "/api/v3/calendar/1",
			ResponseStatus: http.StatusNotFound,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      &starr.ReqError{Code: http.StatusNotFound},
			ExpectedMethod: http.MethodGet,
			WithRequest:    int64(1),
			WithResponse:   (*sonarr.Episode)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetCalendarID(test.WithRequest.(int64))
			assert.ErrorIs(t, err, test.WithError, "the wrong error was returned")
			assert.EqualValues(t, test.WithResponse, output, "make sure ResponseBody and WithResponse are a match")
		})
	}
}
