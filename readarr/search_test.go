package readarr_test

import (
	"net/http"
	"path"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golift.io/starr"
	"golift.io/starr/readarr"
	"golift.io/starr/starrtest"
)

var testSearchJSON = `{
		"foreignId": "123456789",
		"book": {
			"title": "Book",
			"authorTitle": "lastname, firstname Book",
			"seriesTitle": "",
			"disambiguation": "",
			"overview": "book overview",
			"authorId": 123,
			"foreignBookId": "987654",
			"foreignEditionId": "123123",
			"titleSlug": "321321",
			"monitored": true,
			"anyEditionOk": true,
			"ratings": {
				"votes": 123,
				"value": 4.0,
				"popularity": 123.0
			},
			"releaseDate": "1970-01-01T01:00:00Z",
			"pageCount": 123,
			"genres": [
				"Fiction",
				"Classics",
				"Novels"
			],
			"author": {
				"authorMetadataId": 123,
				"status": "continuing",
				"ended": false,
				"authorName": "Firstname Lastname",
				"authorNameLastFirst": "Lastname, Firstname",
				"foreignAuthorId": "1234",
				"titleSlug": "1234",
				"overview": "Firstname Lastname is an author.",
				"links": [
					{
						"url": "https://www.example.com/author",
						"name": "Example"
					}
				],
				"images": [
					{
						"url": "https://www.example.com/author/image",
						"coverType": "poster",
						"extension": ".jpg"
					}
				],
				"path": "/books/Firstname Lastname",
				"qualityProfileId": 1,
				"metadataProfileId": 2,
				"monitored": true,
				"monitorNewItems": "none",
				"folder": "Firstname Lastname",
				"genres": [],
				"cleanName": "firstnamelastname",
				"sortName": "firstname lastname",
				"sortNameLastFirst": "lastname, firstname",
				"tags": [],
				"added": "2025-01-01T01:00:00Z",
				"ratings": {
					"votes": 123,
					"value": 4.0,
					"popularity": 123.0
				},
				"statistics": {
					"bookFileCount": 0,
					"bookCount": 0,
					"availableBookCount": 0,
					"totalBookCount": 0,
					"sizeOnDisk": 0,
					"percentOfBooks": 0
				},
				"id": 123
			},
			"images": [
				{
					"url": "/MediaCover/Books/123/cover.jpg?lastWrite=636992634250000000",
					"coverType": "cover",
					"extension": ".jpg",
					"remoteUrl": "https://www.example.com/book/cover"
				}
			],
			"links": [
				{
					"url": "https://www.example.com/book/edition",
					"name": "Example Editions"
				},
				{
					"url": "https://www.example.com/book/book",
					"name": "Example Book"
				}
			],
			"added": "2025-01-01T01:01:00Z",
			"remoteCover": "https://www.example.com/book/remotecover",
			"editions": [
				{
					"bookId": 123,
					"foreignEditionId": "123",
					"titleSlug": "123",
					"isbn13": "123456789012",
					"asin": "123456789X",
					"title": "Book",
					"language": "eng",
					"overview": "book overview",
					"format": "Hardcover",
					"isEbook": false,
					"disambiguation": "",
					"publisher": "publisher",
					"pageCount": 123,
					"releaseDate": "1970-01-01T01:00:00Z",
					"images": [
						{
							"url": "/MediaCover/Books/123/cover.jpg?lastWrite=636992634250000000",
							"coverType": "cover",
							"extension": ".jpg",
							"remoteUrl": "https://example.com/book/edition/image"
						}
					],
					"links": [
						{
							"url": "https://example.com/book/edition/info",
							"name": "Example Book"
						}
					],
					"ratings": {
						"votes": 123,
						"value": 4.0,
						"popularity": 123.0
					},
					"monitored": true,
					"manualAdd": false,
					"grabbed": false,
					"id": 123
				}
			],
			"grabbed": false,
			"id": 123
		},
		"id": 2
}`

var testEditionImage = starr.Image{
	URL:       "/MediaCover/Books/123/cover.jpg?lastWrite=636992634250000000",
	CoverType: "cover",
	Extension: ".jpg",
	RemoteURL: "https://example.com/book/edition/image",
}

var testBookImage = starr.Image{
	URL:       "/MediaCover/Books/123/cover.jpg?lastWrite=636992634250000000",
	CoverType: "cover",
	Extension: ".jpg",
	RemoteURL: "https://www.example.com/book/cover",
}

var testAuthorImage = starr.Image{
	URL:       "https://www.example.com/author/image",
	CoverType: "poster",
	Extension: ".jpg",
}

var testAuthorLink = starr.Link{
	URL:  "https://www.example.com/author",
	Name: "Example",
}

var testEditionLink = starr.Link{
	URL:  "https://example.com/book/edition/info",
	Name: "Example Book",
}

var testBookLink1 = starr.Link{
	URL:  "https://www.example.com/book/edition",
	Name: "Example Editions",
}

var testBookLink2 = starr.Link{
	URL:  "https://www.example.com/book/book",
	Name: "Example Book",
}

var testBookRating = starr.Ratings{
	Votes:      123,
	Value:      4.0,
	Popularity: 123.0,
}

var testAuthorRating = starr.Ratings{
	Votes:      123,
	Value:      4.0,
	Popularity: 123.0,
}

var testEditionRating = testBookRating

var testEdition = readarr.Edition{
	BookID:           123,
	ForeignEditionID: "123",
	TitleSlug:        "123",
	Isbn13:           "123456789012",
	Asin:             "123456789X",
	Title:            "Book",
	Overview:         "book overview",
	Format:           "Hardcover",
	IsEbook:          false,
	Publisher:        "publisher",
	PageCount:        123,
	ReleaseDate:      time.Date(1970, 1, 1, 1, 0, 0, 0, time.UTC),
	Images:           []*starr.Image{&testEditionImage},
	Links:            []*starr.Link{&testEditionLink},
	Ratings:          &testEditionRating,
	Monitored:        true,
	ManualAdd:        false,
	ID:               123,
}

var testStatistics = readarr.Statistics{
	BookFileCount:      0,
	BookCount:          0,
	AvailableBookCount: 0,
	TotalBookCount:     0,
	SizeOnDisk:         0,
	PercentOfBooks:     0,
}

var testAuthor = readarr.Author{
	AuthorMetadataID:    123,
	Status:              "continuing",
	Ended:               false,
	AuthorName:          "Firstname Lastname",
	AuthorNameLastFirst: "Lastname, Firstname",
	ForeignAuthorID:     "1234",
	TitleSlug:           "1234",
	Overview:            "Firstname Lastname is an author.",
	Links:               []*starr.Link{&testAuthorLink},
	Images:              []*starr.Image{&testAuthorImage},
	Path:                "/books/Firstname Lastname",
	QualityProfileID:    1,
	MetadataProfileID:   2,
	Monitored:           true,
	MonitorNewItems:     "none",
	Genres:              []string{},
	CleanName:           "firstnamelastname",
	SortName:            "firstname lastname",
	SortNameLastFirst:   "lastname, firstname",
	Tags:                []int{},
	Added:               time.Date(2025, 1, 1, 1, 0, 0, 0, time.UTC),
	Ratings:             &testAuthorRating,
	Statistics:          &testStatistics,
	ID:                  123,
}

var testSearchStruct = readarr.SearchResult{
	ForeignID: "123456789",
	ID:        2,
	Book: &readarr.Book{
		Title:          "Book",
		AuthorTitle:    "lastname, firstname Book",
		SeriesTitle:    "",
		Disambiguation: "",
		Overview:       "book overview",
		AuthorID:       123,
		ForeignBookID:  "987654",
		TitleSlug:      "321321",
		Monitored:      true,
		AnyEditionOk:   true,
		Ratings:        &testBookRating,
		ReleaseDate:    time.Date(1970, 1, 1, 1, 0, 0, 0, time.UTC),
		PageCount:      123,
		Genres: []string{
			"Fiction",
			"Classics",
			"Novels",
		},
		Author:      &testAuthor,
		Images:      []*starr.Image{&testBookImage},
		Links:       []*starr.Link{&testBookLink1, &testBookLink2},
		Added:       time.Date(2025, 1, 1, 1, 1, 0, 0, time.UTC),
		RemoteCover: "https://www.example.com/book/remotecover",
		Editions:    []*readarr.Edition{&testEdition},
		Grabbed:     false,
		ID:          123,
	},
	Author: nil,
}

func TestSearch(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, readarr.APIver, "search?term=book"),
			ResponseStatus: http.StatusOK,
			ResponseBody:   `[` + testSearchJSON + `]`,
			WithRequest:    "book",
			WithError:      nil,
			ExpectedMethod: http.MethodGet,
			WithResponse:   []*readarr.SearchResult{&testSearchStruct},
		},

		{
			Name:           "noresults",
			ExpectedPath:   path.Join("/", starr.API, readarr.APIver, "search?term=somestringthatisntabooktitle"),
			ResponseStatus: http.StatusOK,
			ResponseBody:   `[]`,
			WithRequest:    "somestringthatisntabooktitle",
			WithError:      nil,
			ExpectedMethod: http.MethodGet,
			WithResponse:   []*readarr.SearchResult{},
		},

		{
			Name:           "emptyterm",
			ExpectedPath:   path.Join("/", starr.API, readarr.APIver, "search?term="),
			ResponseStatus: http.StatusOK,
			ResponseBody:   `[]`,
			WithRequest:    "",
			WithError:      nil,
			ExpectedMethod: http.MethodGet,
			WithResponse:   []*readarr.SearchResult(nil),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := readarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.Search(test.WithRequest.(string))
			require.ErrorIs(t, err, test.WithError, "the wrong error was returned")
			assert.EqualValues(t, test.WithResponse, output, "make sure ResponseBody and WithResponse are a match")
		})
	}
}
