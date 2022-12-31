package lidarr_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golift.io/starr"
	"golift.io/starr/lidarr"
	"golift.io/starr/starrtest"
)

var testCalendarJSON = `{
  "title": "Mount Westmore",
  "disambiguation": "",
  "overview": "",
  "artistId": 163,
  "foreignAlbumId": "95b28969-3252-45a7-9e1b-b2d8f59eee45",
  "monitored": true,
  "anyReleaseOk": true,
  "profileId": 3,
  "duration": 3897000,
  "albumType": "Album",
  "secondaryTypes": [],
  "mediumCount": 1,
  "ratings": {
    "votes": 0,
    "value": 0
  },
  "releaseDate": "2022-12-09T00:00:00Z",
  "releases": [
    {
      "id": 16428,
      "albumId": 3722,
      "foreignReleaseId": "c2db5e7b-9225-4d01-83bf-b6a4e8c36c02",
      "title": "Mount Westmore",
      "status": "Official",
      "duration": 3897000,
      "trackCount": 16,
      "media": [
        {
          "mediumNumber": 1,
          "mediumName": "",
          "mediumFormat": "Digital Media"
        }
      ],
      "mediumCount": 1,
      "disambiguation": "",
      "country": [
        "[Worldwide]"
      ],
      "label": [
        "Mount Westmore, LLC"
      ],
      "format": "Digital Media",
      "monitored": true
    }
  ],
  "genres": [],
  "media": [
    {
      "mediumNumber": 1,
      "mediumName": "",
      "mediumFormat": "Digital Media"
    }
  ],
  "images": [
    {
      "url": "/lidarr/MediaCover/Albums/3722/cover.jpg?lastWrite=638062138360000000",
      "coverType": "cover",
      "extension": ".jpg",
      "remoteUrl": "https://imagecache.lidarr.audio/v1/caa/c2db5e7b-9225-4d01-83bf-b6a4e8c36c02/34308803649-1200.jpg"
    }
  ],
  "links": [],
  "statistics": {
    "trackFileCount": 0,
    "trackCount": 16,
    "totalTrackCount": 16,
    "sizeOnDisk": 0,
    "percentOfTracks": 0
  },
  "grabbed": false,
  "id": 3722
}`

// This matches the json above.
var testCalendarStruct = lidarr.Album{
	ID:             3722,
	Title:          "Mount Westmore",
	ArtistID:       163,
	ForeignAlbumID: "95b28969-3252-45a7-9e1b-b2d8f59eee45",
	ProfileID:      3,
	SecondaryTypes: []interface{}{},
	Duration:       3897000,
	AlbumType:      "Album",
	MediumCount:    1,
	Links:          []*starr.Link{},
	Ratings:        &starr.Ratings{Votes: 0, Value: 0},
	ReleaseDate:    time.Date(2022, time.December, 9, 0, 0, 0, 0, time.UTC),
	Media: []*lidarr.Media{{
		MediumNumber: 1,
		MediumName:   "",
		MediumFormat: "Digital Media",
	}},
	Releases: []*lidarr.Release{{
		ID:               16428,
		AlbumID:          3722,
		ForeignReleaseID: "c2db5e7b-9225-4d01-83bf-b6a4e8c36c02",
		Title:            "Mount Westmore",
		Status:           "Official",
		Duration:         3897000,
		TrackCount:       16,
		Media: []*lidarr.Media{{
			MediumNumber: 1,
			MediumName:   "",
			MediumFormat: "Digital Media",
		}},
		MediumCount:    1,
		Disambiguation: "",
		Country:        []string{"[Worldwide]"},
		Label:          []string{"Mount Westmore, LLC"},
		Format:         "Digital Media",
		Monitored:      true,
	}},
	Genres: []string{},
	Images: []*starr.Image{
		{
			URL:       "/lidarr/MediaCover/Albums/3722/cover.jpg?lastWrite=638062138360000000",
			CoverType: "cover",
			Extension: ".jpg",
			RemoteURL: "https://imagecache.lidarr.audio/v1/caa/c2db5e7b-9225-4d01-83bf-b6a4e8c36c02/34308803649-1200.jpg",
		},
	},
	Statistics: &lidarr.Statistics{
		TrackFileCount:  0,
		TrackCount:      16,
		TotalTrackCount: 16,
		SizeOnDisk:      0,
		PercentOfTracks: 0,
	},
	Monitored:    true,
	AnyReleaseOk: true,
	Grabbed:      false,
}

func TestGetCalendar(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name: "200",
			ExpectedPath: "/api/v1/calendar" +
				"?end=2020-02-20T04%3A20%3A20.000Z" +
				"&includeArtist=false" +
				"&start=2020-02-20T04%3A20%3A20.000Z" +
				"&unmonitored=true",
			ResponseStatus: http.StatusOK,
			ResponseBody:   `[` + testCalendarJSON + `]`,
			WithRequest: lidarr.Calendar{
				Start:         time.Unix(1582172420, 0),
				End:           time.Unix(1582172420, 0),
				Unmonitored:   true,
				IncludeArtist: false,
			},
			WithError:      nil,
			ExpectedMethod: http.MethodGet,
			WithResponse:   []*lidarr.Album{&testCalendarStruct},
		},
		{
			Name: "404",
			ExpectedPath: "/api/v1/calendar" +
				"?end=2020-02-20T04%3A20%3A20.000Z" +
				"&includeArtist=false" +
				"&start=2020-02-20T04%3A20%3A20.000Z" +
				"&unmonitored=true",
			ResponseStatus: http.StatusNotFound,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      starr.ErrInvalidStatusCode,
			ExpectedMethod: http.MethodGet,
			WithRequest: lidarr.Calendar{
				Start:       time.Unix(1582172420, 0),
				End:         time.Unix(1582172420, 0),
				Unmonitored: true,
			},
			WithResponse: []*lidarr.Album(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := lidarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetCalendar(test.WithRequest.(lidarr.Calendar))
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
			WithError:      starr.ErrInvalidStatusCode,
			ExpectedMethod: http.MethodGet,
			WithRequest:    int64(1),
			WithResponse:   (*lidarr.Album)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := lidarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetCalendarID(test.WithRequest.(int64))
			assert.ErrorIs(t, err, test.WithError, "the wrong error was returned")
			assert.EqualValues(t, test.WithResponse, output, "make sure ResponseBody and WithResponse are a match")
		})
	}
}
