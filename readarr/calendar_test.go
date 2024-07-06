package readarr_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golift.io/starr"
	"golift.io/starr/readarr"
	"golift.io/starr/starrtest"
)

var testCalendarJSON = `{
	"title": "The Invitation",
	"authorTitle": "foley, lucy The Invitation",
	"seriesTitle": "",
	"disambiguation": "",
	"authorId": 1,
	"foreignBookId": "46232124",
	"titleSlug": "46232124",
	"monitored": true,
	"anyEditionOk": true,
	"ratings": {
	  "votes": 4797,
	  "value": 3.67,
	  "popularity": 17604.989999999998
	},
	"releaseDate": "2016-08-02T07:00:00Z",
	"pageCount": 432,
	"genres": [
	  "historical-fiction",
	  "fiction",
	  "romance",
	  "italy"
	],
	"images": [
	  {
		"url": "/readarr/MediaCover/Books/1/cover.jpg?lastWrite=637735196543077695",
		"coverType": "cover",
		"extension": ".jpg"
	  }
	],
	"links": [
	  {
		"url": "https://www.goodreads.com/work/editions/46232124",
		"name": "Goodreads Editions"
	  },
	  {
		"url": "https://www.goodreads.com/book/show/28118525-the-invitation",
		"name": "Goodreads Book"
	  }
	],
	"statistics": {
	  "bookFileCount": 0,
	  "bookCount": 1,
	  "totalBookCount": 1,
	  "sizeOnDisk": 0,
	  "percentOfBooks": 0
	},
	"added": "2020-09-29T04:47:05Z",
	"grabbed": false,
	"id": 1
  }`

// This matches the json above.
var testCalendarStruct = readarr.Book{
	Added:          time.Date(2020, time.September, 29, 4, 47, 5, 0, time.UTC),
	AnyEditionOk:   true,
	Author:         nil,
	AuthorID:       1,
	AuthorTitle:    "foley, lucy The Invitation",
	Disambiguation: "",
	ForeignBookID:  "46232124",
	Genres:         []string{"historical-fiction", "fiction", "romance", "italy"},
	Grabbed:        false,
	ID:             1,
	Images: []*starr.Image{{
		URL:       "/readarr/MediaCover/Books/1/cover.jpg?lastWrite=637735196543077695",
		CoverType: "cover",
		Extension: ".jpg",
	}},
	Links: []*starr.Link{{
		URL:  "https://www.goodreads.com/work/editions/46232124",
		Name: "Goodreads Editions",
	}, {
		URL:  "https://www.goodreads.com/book/show/28118525-the-invitation",
		Name: "Goodreads Book",
	}},
	Monitored: true,
	PageCount: 432,
	Ratings: &starr.Ratings{
		Votes:      4797,
		Value:      3.67,
		Popularity: 17604.989999999998,
	},
	ReleaseDate: time.Date(2016, time.August, 2, 7, 0, 0, 0, time.UTC),
	SeriesTitle: "",
	Statistics: &readarr.Statistics{
		BookFileCount:      0,
		BookCount:          1,
		AvailableBookCount: 0,
		TotalBookCount:     1,
		SizeOnDisk:         0,
		PercentOfBooks:     0,
	},
	Title:     "The Invitation",
	TitleSlug: "46232124",
}

func TestGetCalendar(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name: "200",
			ExpectedPath: "/api/v1/calendar" +
				"?end=2020-02-20T04%3A20%3A20.000Z" +
				"&includeAuthor=false" +
				"&start=2020-02-20T04%3A20%3A20.000Z" +
				"&unmonitored=true",
			ResponseStatus: http.StatusOK,
			ResponseBody:   `[` + testCalendarJSON + `]`,
			WithRequest: readarr.Calendar{
				Start:         time.Unix(1582172420, 0),
				End:           time.Unix(1582172420, 0),
				Unmonitored:   true,
				IncludeAuthor: false,
			},
			WithError:      nil,
			ExpectedMethod: http.MethodGet,
			WithResponse:   []*readarr.Book{&testCalendarStruct},
		},
		{
			Name: "404",
			ExpectedPath: "/api/v1/calendar" +
				"?end=2020-02-20T04%3A20%3A20.000Z" +
				"&includeAuthor=false" +
				"&start=2020-02-20T04%3A20%3A20.000Z" +
				"&unmonitored=true",
			ResponseStatus: http.StatusNotFound,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      &starr.ReqError{Code: http.StatusNotFound},
			ExpectedMethod: http.MethodGet,
			WithRequest: readarr.Calendar{
				Start:       time.Unix(1582172420, 0),
				End:         time.Unix(1582172420, 0),
				Unmonitored: true,
			},
			WithResponse: []*readarr.Book(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := readarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetCalendar(test.WithRequest.(readarr.Calendar))
			require.ErrorIs(t, err, test.WithError, "the wrong error was returned")
			assert.EqualValues(t, test.WithResponse, output, "make sure ResponseBody and WithResponse are a match")
		})
	}
}

func TestGetCalendarID(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   "/api/v1/calendar/1",
			ResponseStatus: http.StatusOK,
			ResponseBody:   testCalendarJSON,
			WithRequest:    int64(1),
			WithError:      nil,
			ExpectedMethod: http.MethodGet,
			WithResponse:   &testCalendarStruct,
		},
		{
			Name:           "404",
			ExpectedPath:   "/api/v1/calendar/1",
			ResponseStatus: http.StatusNotFound,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      &starr.ReqError{Code: http.StatusNotFound},
			ExpectedMethod: http.MethodGet,
			WithRequest:    int64(1),
			WithResponse:   (*readarr.Book)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := readarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetCalendarID(test.WithRequest.(int64))
			require.ErrorIs(t, err, test.WithError, "the wrong error was returned")
			assert.EqualValues(t, test.WithResponse, output, "make sure ResponseBody and WithResponse are a match")
		})
	}
}
