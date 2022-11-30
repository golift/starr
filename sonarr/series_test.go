package sonarr_test

import (
	"path"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golift.io/starr"
	"golift.io/starr/sonarr"
)

const (
	firstSeries = `{
	"title": "Breaking Bad",
	"alternateTitles": [],
	"sortTitle": "breaking bad",
	"status": "ended",
	"ended": true,
	"overview": "Cool teacher becomes drug dealer",
	"previousAiring": "2019-06-04T01:00:00Z",
	"network": "Netflix",
	"airTime": "21:00",
	"images": [
		{
			"coverType": "banner",
			"url": "/MediaCover/1/banner.jpg?lastWrite=637829401993017870",
			"remoteUrl": "https://artworks.thetvdb.com/banners/graphical/81189-g21.jpg"
		},
		{
			"coverType": "poster",
			"url": "/MediaCover/1/poster.jpg?lastWrite=637829401994817870",
			"remoteUrl": "https://artworks.thetvdb.com/banners/posters/81189-10.jpg"
		},
		{
			"coverType": "fanart",
			"url": "/MediaCover/1/fanart.jpg?lastWrite=637829401996517870",
			"remoteUrl": "https://artworks.thetvdb.com/banners/fanart/original/81189-21.jpg"
		}
	],
	"seasons": [
		{
			"seasonNumber": 0,
			"monitored": false,
			"statistics": {
				"episodeFileCount": 0,
				"episodeCount": 0,
				"totalEpisodeCount": 19,
				"sizeOnDisk": 0,
				"percentOfEpisodes": 0.0
			}
		},
		{
			"seasonNumber": 1,
			"monitored": true,
			"statistics": {
				"previousAiring": "2019-06-04T01:00:00Z",
				"episodeFileCount": 0,
				"episodeCount": 7,
				"totalEpisodeCount": 7,
				"sizeOnDisk": 0,
				"percentOfEpisodes": 0.0
			}
		},
		{
			"seasonNumber": 2,
			"monitored": true,
			"statistics": {
				"previousAiring": "2019-06-04T01:00:00Z",
				"episodeFileCount": 0,
				"episodeCount": 13,
				"totalEpisodeCount": 13,
				"sizeOnDisk": 0,
				"percentOfEpisodes": 0.0
			}
		},
		{
			"seasonNumber": 3,
			"monitored": true,
			"statistics": {
				"previousAiring": "2019-06-04T01:00:00Z",
				"episodeFileCount": 0,
				"episodeCount": 13,
				"totalEpisodeCount": 13,
				"sizeOnDisk": 0,
				"percentOfEpisodes": 0.0
			}
		},
		{
			"seasonNumber": 4,
			"monitored": true,
			"statistics": {
				"previousAiring": "2019-06-04T01:00:00Z",
				"episodeFileCount": 0,
				"episodeCount": 13,
				"totalEpisodeCount": 13,
				"sizeOnDisk": 0,
				"percentOfEpisodes": 0.0
			}
		},
		{
			"seasonNumber": 5,
			"monitored": true,
			"statistics": {
				"previousAiring": "2019-06-04T01:00:00Z",
				"episodeFileCount": 0,
				"episodeCount": 16,
				"totalEpisodeCount": 16,
				"sizeOnDisk": 0,
				"percentOfEpisodes": 0.0
			}
		}
	],
	"year": 2008,
	"path": "/series/Breaking Bad",
	"qualityProfileId": 1,
	"languageProfileId": 1,
	"seasonFolder": true,
	"monitored": true,
	"useSceneNumbering": false,
	"runtime": 47,
	"tvdbId": 81189,
	"tvRageId": 18164,
	"tvMazeId": 169,
	"firstAired": "2019-06-04T01:00:00Z",
	"seriesType": "standard",
	"cleanTitle": "breakingbad",
	"imdbId": "tt0903747",
	"titleSlug": "breaking-bad",
	"rootFolderPath": "/series/",
	"certification": "TV-MA",
	"genres": [
		"Crime",
		"Drama",
		"Suspense",
		"Thriller"
	],
	"tags": [
		11
	],
	"added": "2019-06-04T01:00:00Z",
	"ratings": {
		"votes": 31714,
		"value": 9.4
	},
	"statistics": {
		"seasonCount": 5,
		"episodeFileCount": 0,
		"episodeCount": 62,
		"totalEpisodeCount": 81,
		"sizeOnDisk": 0,
		"percentOfEpisodes": 0.0
	},
	"id": 1
}`
	secondSeries = `{
	"title": "Chernobyl",
	"alternateTitles": [],
	"sortTitle": "chernobyl",
	"status": "ended",
	"ended": true,
	"overview": "A lot of energy wasted",
	"previousAiring": "2019-06-04T01:00:00Z",
	"network": "HBO",
	"airTime": "21:00",
	"images": [
		{
			"coverType": "banner",
			"url": "/MediaCover/2/banner.jpg?lastWrite=637829402715717870",
			"remoteUrl": "https://artworks.thetvdb.com/banners/graphical/5cc9f74c2ddd3.jpg"
		},
		{
			"coverType": "poster",
			"url": "/MediaCover/2/poster.jpg?lastWrite=637829402718117870",
			"remoteUrl": "https://artworks.thetvdb.com/banners/posters/5cc12861c93e4.jpg"
		},
		{
			"coverType": "fanart",
			"url": "/MediaCover/2/fanart.jpg?lastWrite=637829402721317870",
			"remoteUrl": "https://artworks.thetvdb.com/banners/series/360893/backgrounds/62017319.jpg"
		}
	],
	"seasons": [
		{
			"seasonNumber": 0,
			"monitored": false,
			"statistics": {
				"episodeFileCount": 0,
				"episodeCount": 0,
				"totalEpisodeCount": 14,
				"sizeOnDisk": 0,
				"percentOfEpisodes": 0.0
			}
		},
		{
			"seasonNumber": 1,
			"monitored": true,
			"statistics": {
				"previousAiring": "2019-06-04T01:00:00Z",
				"episodeFileCount": 0,
				"episodeCount": 5,
				"totalEpisodeCount": 5,
				"sizeOnDisk": 0,
				"percentOfEpisodes": 0.0
			}
		}
	],
	"year": 2019,
	"path": "/series/Chernobyl",
	"qualityProfileId": 1,
	"languageProfileId": 1,
	"seasonFolder": true,
	"monitored": true,
	"useSceneNumbering": false,
	"runtime": 65,
	"tvdbId": 360893,
	"tvRageId": 0,
	"tvMazeId": 30770,
	"firstAired": "2019-06-04T01:00:00Z",
	"seriesType": "standard",
	"cleanTitle": "chernobyl",
	"imdbId": "tt7366338",
	"titleSlug": "chernobyl",
	"rootFolderPath": "/series/",
	"certification": "TV-MA",
	"genres": [
		"Drama",
		"History",
		"Mini-Series",
		"Thriller"
	],
	"tags": [
		11
	],
	"added": "2019-06-04T01:00:00Z",
	"ratings": {
		"votes": 83,
		"value": 8.7
	},
	"statistics": {
		"seasonCount": 1,
		"episodeFileCount": 0,
		"episodeCount": 5,
		"totalEpisodeCount": 19,
		"sizeOnDisk": 0,
		"percentOfEpisodes": 0.0
	},
	"id": 2
}`
	listSeries = `[` + firstSeries + `,` + secondSeries + `]`
	addSeries  = `{"monitored":true,"seasonFolder":true,"languageProfileId":1,"qualityProfileId":1,` +
		`"tvdbId":360893,"path":"/series/Chernobyl","title":"Chernobyl","titleSlug":"chernobyl",` +
		`"rootFolderPath":"/series/","tags":[11],` +
		`"seasons":[{"monitored":false,"seasonNumber":0},{"monitored":true,"seasonNumber":1}],` +
		`"images":[{"coverType":"banner","remoteUrl":"https://artworks.thetvdb.com/banners/graphical/5cc9f74c2ddd3.jpg"},` +
		`{"coverType":"poster","remoteUrl":"https://artworks.thetvdb.com/banners/posters/5cc12861c93e4.jpg"},` +
		`{"coverType":"fanart","remoteUrl":"https://artworks.thetvdb.com/banners/series/360893/backgrounds/62017319.jpg"}],` +
		`"addOptions":{"searchForMissingEpisodes":true}}` +
		"\n"
	updateSeries = `{"monitored":true,"seasonFolder":true,"id":1,"languageProfileId":1,"qualityProfileId":1,` +
		`"tvdbId":360893,"path":"/series/Chernobyl","title":"Chernobyl","rootFolderPath":` +
		`"/series/","tags":[11],"seasons":[{"monitored":false,"seasonNumber":0},{"monitored":true,"seasonNumber":1}]}` +
		"\n"
)

func TestGetAllSeries(t *testing.T) {
	t.Parallel()

	loc, _ := time.LoadLocation("")
	date := time.Date(2019, 6, 4, 1, 0, 0, 0, loc)

	tests := []*starr.TestMockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "series"),
			ExpectedMethod: "GET",
			ResponseStatus: 200,
			ResponseBody:   listSeries,
			WithResponse: []*sonarr.Series{
				{
					Title:           "Breaking Bad",
					AlternateTitles: []*sonarr.AlternateTitle{},
					SortTitle:       "breaking bad",
					Status:          "ended",
					Ended:           true,
					Overview:        "Cool teacher becomes drug dealer",
					PreviousAiring:  date,
					Network:         "Netflix",
					AirTime:         "21:00",
					Images: []*starr.Image{
						{
							CoverType: "banner",
							URL:       "/MediaCover/1/banner.jpg?lastWrite=637829401993017870",
							RemoteURL: "https://artworks.thetvdb.com/banners/graphical/81189-g21.jpg",
						},
						{
							CoverType: "poster",
							URL:       "/MediaCover/1/poster.jpg?lastWrite=637829401994817870",
							RemoteURL: "https://artworks.thetvdb.com/banners/posters/81189-10.jpg",
						},
						{
							CoverType: "fanart",
							URL:       "/MediaCover/1/fanart.jpg?lastWrite=637829401996517870",
							RemoteURL: "https://artworks.thetvdb.com/banners/fanart/original/81189-21.jpg",
						},
					},
					Seasons: []*sonarr.Season{
						{
							SeasonNumber: 0,
							Monitored:    false,
							Statistics: &sonarr.Statistics{
								EpisodeFileCount:  0,
								EpisodeCount:      0,
								TotalEpisodeCount: 19,
								SizeOnDisk:        0,
								PercentOfEpisodes: 0.0,
							},
						},
						{
							SeasonNumber: 1,
							Monitored:    true,
							Statistics: &sonarr.Statistics{
								PreviousAiring:    date,
								EpisodeFileCount:  0,
								EpisodeCount:      7,
								TotalEpisodeCount: 7,
								SizeOnDisk:        0,
								PercentOfEpisodes: 0.0,
							},
						},
						{
							SeasonNumber: 2,
							Monitored:    true,
							Statistics: &sonarr.Statistics{
								PreviousAiring:    date,
								EpisodeFileCount:  0,
								EpisodeCount:      13,
								TotalEpisodeCount: 13,
								SizeOnDisk:        0,
								PercentOfEpisodes: 0.0,
							},
						},
						{
							SeasonNumber: 3,
							Monitored:    true,
							Statistics: &sonarr.Statistics{
								PreviousAiring:    date,
								EpisodeFileCount:  0,
								EpisodeCount:      13,
								TotalEpisodeCount: 13,
								SizeOnDisk:        0,
								PercentOfEpisodes: 0.0,
							},
						},
						{
							SeasonNumber: 4,
							Monitored:    true,
							Statistics: &sonarr.Statistics{
								PreviousAiring:    date,
								EpisodeFileCount:  0,
								EpisodeCount:      13,
								TotalEpisodeCount: 13,
								SizeOnDisk:        0,
								PercentOfEpisodes: 0.0,
							},
						},
						{
							SeasonNumber: 5,
							Monitored:    true,
							Statistics: &sonarr.Statistics{
								PreviousAiring:    date,
								EpisodeFileCount:  0,
								EpisodeCount:      16,
								TotalEpisodeCount: 16,
								SizeOnDisk:        0,
								PercentOfEpisodes: 0.0,
							},
						},
					},
					Year:              2008,
					Path:              "/series/Breaking Bad",
					QualityProfileID:  1,
					LanguageProfileID: 1,
					SeasonFolder:      true,
					Monitored:         true,
					UseSceneNumbering: false,
					Runtime:           47,
					TvdbID:            81189,
					TvRageID:          18164,
					TvMazeID:          169,
					FirstAired:        date,
					SeriesType:        "standard",
					CleanTitle:        "breakingbad",
					ImdbID:            "tt0903747",
					TitleSlug:         "breaking-bad",
					RootFolderPath:    "/series/",
					Certification:     "TV-MA",
					Genres: []string{
						"Crime",
						"Drama",
						"Suspense",
						"Thriller",
					},
					Tags: []int{
						11,
					},
					Added: date,
					Ratings: &starr.Ratings{
						Votes: 31714,
						Value: 9.4,
					},
					Statistics: &sonarr.Statistics{
						SeasonCount:       5,
						EpisodeFileCount:  0,
						EpisodeCount:      62,
						TotalEpisodeCount: 81,
						SizeOnDisk:        0,
						PercentOfEpisodes: 0.0,
					},
					ID: 1,
				},
				{
					Title:           "Chernobyl",
					AlternateTitles: []*sonarr.AlternateTitle{},
					SortTitle:       "chernobyl",
					Status:          "ended",
					Ended:           true,
					Overview:        "A lot of energy wasted",
					PreviousAiring:  date,
					Network:         "HBO",
					AirTime:         "21:00",
					Images: []*starr.Image{
						{
							CoverType: "banner",
							URL:       "/MediaCover/2/banner.jpg?lastWrite=637829402715717870",
							RemoteURL: "https://artworks.thetvdb.com/banners/graphical/5cc9f74c2ddd3.jpg",
						},
						{
							CoverType: "poster",
							URL:       "/MediaCover/2/poster.jpg?lastWrite=637829402718117870",
							RemoteURL: "https://artworks.thetvdb.com/banners/posters/5cc12861c93e4.jpg",
						},
						{
							CoverType: "fanart",
							URL:       "/MediaCover/2/fanart.jpg?lastWrite=637829402721317870",
							RemoteURL: "https://artworks.thetvdb.com/banners/series/360893/backgrounds/62017319.jpg",
						},
					},
					Seasons: []*sonarr.Season{
						{
							SeasonNumber: 0,
							Monitored:    false,
							Statistics: &sonarr.Statistics{
								EpisodeFileCount:  0,
								EpisodeCount:      0,
								TotalEpisodeCount: 14,
								SizeOnDisk:        0,
								PercentOfEpisodes: 0.0,
							},
						},
						{
							SeasonNumber: 1,
							Monitored:    true,
							Statistics: &sonarr.Statistics{
								PreviousAiring:    date,
								EpisodeFileCount:  0,
								EpisodeCount:      5,
								TotalEpisodeCount: 5,
								SizeOnDisk:        0,
								PercentOfEpisodes: 0.0,
							},
						},
					},
					Year:              2019,
					Path:              "/series/Chernobyl",
					QualityProfileID:  1,
					LanguageProfileID: 1,
					SeasonFolder:      true,
					Monitored:         true,
					UseSceneNumbering: false,
					Runtime:           65,
					TvdbID:            360893,
					TvRageID:          0,
					TvMazeID:          30770,
					FirstAired:        date,
					SeriesType:        "standard",
					CleanTitle:        "chernobyl",
					ImdbID:            "tt7366338",
					TitleSlug:         "chernobyl",
					RootFolderPath:    "/series/",
					Certification:     "TV-MA",
					Genres: []string{
						"Drama",
						"History",
						"Mini-Series",
						"Thriller",
					},
					Tags: []int{
						11,
					},
					Added: date,
					Ratings: &starr.Ratings{
						Votes: 83,
						Value: 8.7,
					},
					Statistics: &sonarr.Statistics{
						SeasonCount:       1,
						EpisodeFileCount:  0,
						EpisodeCount:      5,
						TotalEpisodeCount: 19,
						SizeOnDisk:        0,
						PercentOfEpisodes: 0.0,
					},
					ID: 2,
				},
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "series"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			ResponseBody:   starr.BodyNotFound,
			WithError:      starr.ErrInvalidStatusCode,
			WithResponse:   []*sonarr.Series(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetAllSeries()
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, output, test.WithResponse, "response is not the same as expected")
		})
	}
}

func TestGetSeries(t *testing.T) {
	t.Parallel()

	loc, _ := time.LoadLocation("")
	date := time.Date(2019, 6, 4, 1, 0, 0, 0, loc)

	tests := []*starr.TestMockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "series?tvdbId=360893"),
			ExpectedMethod: "GET",
			ResponseStatus: 200,
			ResponseBody:   "[" + secondSeries + "]",
			WithRequest:    360893,
			WithResponse: []*sonarr.Series{
				{
					Title:           "Chernobyl",
					AlternateTitles: []*sonarr.AlternateTitle{},
					SortTitle:       "chernobyl",
					Status:          "ended",
					Ended:           true,
					Overview:        "A lot of energy wasted",
					PreviousAiring:  date,
					Network:         "HBO",
					AirTime:         "21:00",
					Images: []*starr.Image{
						{
							CoverType: "banner",
							URL:       "/MediaCover/2/banner.jpg?lastWrite=637829402715717870",
							RemoteURL: "https://artworks.thetvdb.com/banners/graphical/5cc9f74c2ddd3.jpg",
						},
						{
							CoverType: "poster",
							URL:       "/MediaCover/2/poster.jpg?lastWrite=637829402718117870",
							RemoteURL: "https://artworks.thetvdb.com/banners/posters/5cc12861c93e4.jpg",
						},
						{
							CoverType: "fanart",
							URL:       "/MediaCover/2/fanart.jpg?lastWrite=637829402721317870",
							RemoteURL: "https://artworks.thetvdb.com/banners/series/360893/backgrounds/62017319.jpg",
						},
					},
					Seasons: []*sonarr.Season{
						{
							SeasonNumber: 0,
							Monitored:    false,
							Statistics: &sonarr.Statistics{
								EpisodeFileCount:  0,
								EpisodeCount:      0,
								TotalEpisodeCount: 14,
								SizeOnDisk:        0,
								PercentOfEpisodes: 0.0,
							},
						},
						{
							SeasonNumber: 1,
							Monitored:    true,
							Statistics: &sonarr.Statistics{
								PreviousAiring:    date,
								EpisodeFileCount:  0,
								EpisodeCount:      5,
								TotalEpisodeCount: 5,
								SizeOnDisk:        0,
								PercentOfEpisodes: 0.0,
							},
						},
					},
					Year:              2019,
					Path:              "/series/Chernobyl",
					QualityProfileID:  1,
					LanguageProfileID: 1,
					SeasonFolder:      true,
					Monitored:         true,
					UseSceneNumbering: false,
					Runtime:           65,
					TvdbID:            360893,
					TvRageID:          0,
					TvMazeID:          30770,
					FirstAired:        date,
					SeriesType:        "standard",
					CleanTitle:        "chernobyl",
					ImdbID:            "tt7366338",
					TitleSlug:         "chernobyl",
					RootFolderPath:    "/series/",
					Certification:     "TV-MA",
					Genres: []string{
						"Drama",
						"History",
						"Mini-Series",
						"Thriller",
					},
					Tags: []int{
						11,
					},
					Added: date,
					Ratings: &starr.Ratings{
						Votes: 83,
						Value: 8.7,
					},
					Statistics: &sonarr.Statistics{
						SeasonCount:       1,
						EpisodeFileCount:  0,
						EpisodeCount:      5,
						TotalEpisodeCount: 19,
						SizeOnDisk:        0,
						PercentOfEpisodes: 0.0,
					},
					ID: 2,
				},
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "series?tvdbId=360893"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			ResponseBody:   starr.BodyNotFound,
			WithRequest:    360893,
			WithError:      starr.ErrInvalidStatusCode,
			WithResponse:   []*sonarr.Series(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetSeries(int64(test.WithRequest.(int)))
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, output, test.WithResponse, "response is not the same as expected")
		})
	}
}

func TestGetSeriesByID(t *testing.T) {
	t.Parallel()

	loc, _ := time.LoadLocation("")
	date := time.Date(2019, 6, 4, 1, 0, 0, 0, loc)

	tests := []*starr.TestMockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "series", "2"),
			ExpectedMethod: "GET",
			ResponseStatus: 200,
			ResponseBody:   secondSeries,
			WithRequest:    2,
			WithResponse: &sonarr.Series{
				Title:           "Chernobyl",
				AlternateTitles: []*sonarr.AlternateTitle{},
				SortTitle:       "chernobyl",
				Status:          "ended",
				Ended:           true,
				Overview:        "A lot of energy wasted",
				PreviousAiring:  date,
				Network:         "HBO",
				AirTime:         "21:00",
				Images: []*starr.Image{
					{
						CoverType: "banner",
						URL:       "/MediaCover/2/banner.jpg?lastWrite=637829402715717870",
						RemoteURL: "https://artworks.thetvdb.com/banners/graphical/5cc9f74c2ddd3.jpg",
					},
					{
						CoverType: "poster",
						URL:       "/MediaCover/2/poster.jpg?lastWrite=637829402718117870",
						RemoteURL: "https://artworks.thetvdb.com/banners/posters/5cc12861c93e4.jpg",
					},
					{
						CoverType: "fanart",
						URL:       "/MediaCover/2/fanart.jpg?lastWrite=637829402721317870",
						RemoteURL: "https://artworks.thetvdb.com/banners/series/360893/backgrounds/62017319.jpg",
					},
				},
				Seasons: []*sonarr.Season{
					{
						SeasonNumber: 0,
						Monitored:    false,
						Statistics: &sonarr.Statistics{
							EpisodeFileCount:  0,
							EpisodeCount:      0,
							TotalEpisodeCount: 14,
							SizeOnDisk:        0,
							PercentOfEpisodes: 0.0,
						},
					},
					{
						SeasonNumber: 1,
						Monitored:    true,
						Statistics: &sonarr.Statistics{
							PreviousAiring:    date,
							EpisodeFileCount:  0,
							EpisodeCount:      5,
							TotalEpisodeCount: 5,
							SizeOnDisk:        0,
							PercentOfEpisodes: 0.0,
						},
					},
				},
				Year:              2019,
				Path:              "/series/Chernobyl",
				QualityProfileID:  1,
				LanguageProfileID: 1,
				SeasonFolder:      true,
				Monitored:         true,
				UseSceneNumbering: false,
				Runtime:           65,
				TvdbID:            360893,
				TvRageID:          0,
				TvMazeID:          30770,
				FirstAired:        date,
				SeriesType:        "standard",
				CleanTitle:        "chernobyl",
				ImdbID:            "tt7366338",
				TitleSlug:         "chernobyl",
				RootFolderPath:    "/series/",
				Certification:     "TV-MA",
				Genres: []string{
					"Drama",
					"History",
					"Mini-Series",
					"Thriller",
				},
				Tags: []int{
					11,
				},
				Added: date,
				Ratings: &starr.Ratings{
					Votes: 83,
					Value: 8.7,
				},
				Statistics: &sonarr.Statistics{
					SeasonCount:       1,
					EpisodeFileCount:  0,
					EpisodeCount:      5,
					TotalEpisodeCount: 19,
					SizeOnDisk:        0,
					PercentOfEpisodes: 0.0,
				},
				ID: 2,
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "series", "2"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			ResponseBody:   starr.BodyNotFound,
			WithRequest:    2,
			WithError:      starr.ErrInvalidStatusCode,
			WithResponse:   (*sonarr.Series)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetSeriesByID(int64(test.WithRequest.(int)))
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, output, test.WithResponse, "response is not the same as expected")
		})
	}
}

func TestAddSeries(t *testing.T) {
	t.Parallel()

	loc, _ := time.LoadLocation("")
	date := time.Date(2019, 6, 4, 1, 0, 0, 0, loc)

	tests := []*starr.TestMockData{
		{
			Name:            "200",
			ExpectedPath:    path.Join("/", starr.API, sonarr.APIver, "series?moveFiles=true"),
			ExpectedMethod:  "POST",
			ExpectedRequest: addSeries,
			ResponseStatus:  200,
			WithRequest: &sonarr.AddSeriesInput{
				Title: "Chernobyl",
				Images: []*starr.Image{
					{
						CoverType: "banner",
						RemoteURL: "https://artworks.thetvdb.com/banners/graphical/5cc9f74c2ddd3.jpg",
					},
					{
						CoverType: "poster",
						RemoteURL: "https://artworks.thetvdb.com/banners/posters/5cc12861c93e4.jpg",
					},
					{
						CoverType: "fanart",
						RemoteURL: "https://artworks.thetvdb.com/banners/series/360893/backgrounds/62017319.jpg",
					},
				},
				Seasons: []*sonarr.Season{
					{
						SeasonNumber: 0,
						Monitored:    false,
					},
					{
						SeasonNumber: 1,
						Monitored:    true,
					},
				},
				Path:              "/series/Chernobyl",
				QualityProfileID:  1,
				LanguageProfileID: 1,
				SeasonFolder:      true,
				Monitored:         true,
				TvdbID:            360893,
				TitleSlug:         "chernobyl",
				RootFolderPath:    "/series/",
				Tags: []int{
					11,
				},
				AddOptions: &sonarr.AddSeriesOptions{
					SearchForMissingEpisodes: true,
				},
			},
			ResponseBody: secondSeries,
			WithResponse: &sonarr.Series{
				Title:           "Chernobyl",
				AlternateTitles: []*sonarr.AlternateTitle{},
				SortTitle:       "chernobyl",
				Status:          "ended",
				Ended:           true,
				Overview:        "A lot of energy wasted",
				PreviousAiring:  date,
				Network:         "HBO",
				AirTime:         "21:00",
				Images: []*starr.Image{
					{
						CoverType: "banner",
						URL:       "/MediaCover/2/banner.jpg?lastWrite=637829402715717870",
						RemoteURL: "https://artworks.thetvdb.com/banners/graphical/5cc9f74c2ddd3.jpg",
					},
					{
						CoverType: "poster",
						URL:       "/MediaCover/2/poster.jpg?lastWrite=637829402718117870",
						RemoteURL: "https://artworks.thetvdb.com/banners/posters/5cc12861c93e4.jpg",
					},
					{
						CoverType: "fanart",
						URL:       "/MediaCover/2/fanart.jpg?lastWrite=637829402721317870",
						RemoteURL: "https://artworks.thetvdb.com/banners/series/360893/backgrounds/62017319.jpg",
					},
				},
				Seasons: []*sonarr.Season{
					{
						SeasonNumber: 0,
						Monitored:    false,
						Statistics: &sonarr.Statistics{
							EpisodeFileCount:  0,
							EpisodeCount:      0,
							TotalEpisodeCount: 14,
							SizeOnDisk:        0,
							PercentOfEpisodes: 0.0,
						},
					},
					{
						SeasonNumber: 1,
						Monitored:    true,
						Statistics: &sonarr.Statistics{
							PreviousAiring:    date,
							EpisodeFileCount:  0,
							EpisodeCount:      5,
							TotalEpisodeCount: 5,
							SizeOnDisk:        0,
							PercentOfEpisodes: 0.0,
						},
					},
				},
				Year:              2019,
				Path:              "/series/Chernobyl",
				QualityProfileID:  1,
				LanguageProfileID: 1,
				SeasonFolder:      true,
				Monitored:         true,
				UseSceneNumbering: false,
				Runtime:           65,
				TvdbID:            360893,
				TvRageID:          0,
				TvMazeID:          30770,
				FirstAired:        date,
				SeriesType:        "standard",
				CleanTitle:        "chernobyl",
				ImdbID:            "tt7366338",
				TitleSlug:         "chernobyl",
				RootFolderPath:    "/series/",
				Certification:     "TV-MA",
				Genres: []string{
					"Drama",
					"History",
					"Mini-Series",
					"Thriller",
				},
				Tags: []int{
					11,
				},
				Added: date,
				Ratings: &starr.Ratings{
					Votes: 83,
					Value: 8.7,
				},
				Statistics: &sonarr.Statistics{
					SeasonCount:       1,
					EpisodeFileCount:  0,
					EpisodeCount:      5,
					TotalEpisodeCount: 19,
					SizeOnDisk:        0,
					PercentOfEpisodes: 0.0,
				},
				ID: 2,
			},
			WithError: nil,
		},
		{
			Name:            "404",
			ExpectedPath:    path.Join("/", starr.API, sonarr.APIver, "series?moveFiles=true"),
			ExpectedMethod:  "POST",
			ExpectedRequest: addSeries,
			ResponseStatus:  404,
			WithRequest: &sonarr.AddSeriesInput{
				Title: "Chernobyl",
				Images: []*starr.Image{
					{
						CoverType: "banner",
						RemoteURL: "https://artworks.thetvdb.com/banners/graphical/5cc9f74c2ddd3.jpg",
					},
					{
						CoverType: "poster",
						RemoteURL: "https://artworks.thetvdb.com/banners/posters/5cc12861c93e4.jpg",
					},
					{
						CoverType: "fanart",
						RemoteURL: "https://artworks.thetvdb.com/banners/series/360893/backgrounds/62017319.jpg",
					},
				},
				Seasons: []*sonarr.Season{
					{
						SeasonNumber: 0,
						Monitored:    false,
					},
					{
						SeasonNumber: 1,
						Monitored:    true,
					},
				},
				Path:              "/series/Chernobyl",
				QualityProfileID:  1,
				LanguageProfileID: 1,
				SeasonFolder:      true,
				Monitored:         true,
				TvdbID:            360893,
				TitleSlug:         "chernobyl",
				RootFolderPath:    "/series/",
				Tags: []int{
					11,
				},
				AddOptions: &sonarr.AddSeriesOptions{
					SearchForMissingEpisodes: true,
				},
			},
			ResponseBody: starr.BodyNotFound,
			WithError:    starr.ErrInvalidStatusCode,
			WithResponse: (*sonarr.Series)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.AddSeries(test.WithRequest.(*sonarr.AddSeriesInput))
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, output, test.WithResponse, "response is not the same as expected")
		})
	}
}

func TestUpdateSeries(t *testing.T) {
	t.Parallel()

	loc, _ := time.LoadLocation("")
	date := time.Date(2019, 6, 4, 1, 0, 0, 0, loc)

	tests := []*starr.TestMockData{
		{
			Name:            "200",
			ExpectedPath:    path.Join("/", starr.API, sonarr.APIver, "series/1?moveFiles=true"),
			ExpectedMethod:  "PUT",
			ExpectedRequest: updateSeries,
			ResponseStatus:  200,
			WithRequest: &sonarr.AddSeriesInput{
				Title: "Chernobyl",
				Seasons: []*sonarr.Season{
					{
						SeasonNumber: 0,
						Monitored:    false,
					},
					{
						SeasonNumber: 1,
						Monitored:    true,
					},
				},
				Path:              "/series/Chernobyl",
				QualityProfileID:  1,
				LanguageProfileID: 1,
				SeasonFolder:      true,
				Monitored:         true,
				TvdbID:            360893,
				RootFolderPath:    "/series/",
				Tags: []int{
					11,
				},
				ID: 1,
			},
			ResponseBody: secondSeries,
			WithResponse: &sonarr.Series{
				Title:           "Chernobyl",
				AlternateTitles: []*sonarr.AlternateTitle{},
				SortTitle:       "chernobyl",
				Status:          "ended",
				Ended:           true,
				Overview:        "A lot of energy wasted",
				PreviousAiring:  date,
				Network:         "HBO",
				AirTime:         "21:00",
				Images: []*starr.Image{
					{
						CoverType: "banner",
						URL:       "/MediaCover/2/banner.jpg?lastWrite=637829402715717870",
						RemoteURL: "https://artworks.thetvdb.com/banners/graphical/5cc9f74c2ddd3.jpg",
					},
					{
						CoverType: "poster",
						URL:       "/MediaCover/2/poster.jpg?lastWrite=637829402718117870",
						RemoteURL: "https://artworks.thetvdb.com/banners/posters/5cc12861c93e4.jpg",
					},
					{
						CoverType: "fanart",
						URL:       "/MediaCover/2/fanart.jpg?lastWrite=637829402721317870",
						RemoteURL: "https://artworks.thetvdb.com/banners/series/360893/backgrounds/62017319.jpg",
					},
				},
				Seasons: []*sonarr.Season{
					{
						SeasonNumber: 0,
						Monitored:    false,
						Statistics: &sonarr.Statistics{
							EpisodeFileCount:  0,
							EpisodeCount:      0,
							TotalEpisodeCount: 14,
							SizeOnDisk:        0,
							PercentOfEpisodes: 0.0,
						},
					},
					{
						SeasonNumber: 1,
						Monitored:    true,
						Statistics: &sonarr.Statistics{
							PreviousAiring:    date,
							EpisodeFileCount:  0,
							EpisodeCount:      5,
							TotalEpisodeCount: 5,
							SizeOnDisk:        0,
							PercentOfEpisodes: 0.0,
						},
					},
				},
				Year:              2019,
				Path:              "/series/Chernobyl",
				QualityProfileID:  1,
				LanguageProfileID: 1,
				SeasonFolder:      true,
				Monitored:         true,
				UseSceneNumbering: false,
				Runtime:           65,
				TvdbID:            360893,
				TvRageID:          0,
				TvMazeID:          30770,
				FirstAired:        date,
				SeriesType:        "standard",
				CleanTitle:        "chernobyl",
				ImdbID:            "tt7366338",
				TitleSlug:         "chernobyl",
				RootFolderPath:    "/series/",
				Certification:     "TV-MA",
				Genres: []string{
					"Drama",
					"History",
					"Mini-Series",
					"Thriller",
				},
				Tags: []int{
					11,
				},
				Added: date,
				Ratings: &starr.Ratings{
					Votes: 83,
					Value: 8.7,
				},
				Statistics: &sonarr.Statistics{
					SeasonCount:       1,
					EpisodeFileCount:  0,
					EpisodeCount:      5,
					TotalEpisodeCount: 19,
					SizeOnDisk:        0,
					PercentOfEpisodes: 0.0,
				},
				ID: 2,
			},
			WithError: nil,
		},
		{
			Name:            "404",
			ExpectedPath:    path.Join("/", starr.API, sonarr.APIver, "series/1?moveFiles=true"),
			ExpectedMethod:  "PUT",
			ExpectedRequest: updateSeries,
			ResponseStatus:  404,
			WithRequest: &sonarr.AddSeriesInput{
				Title: "Chernobyl",
				Seasons: []*sonarr.Season{
					{
						SeasonNumber: 0,
						Monitored:    false,
					},
					{
						SeasonNumber: 1,
						Monitored:    true,
					},
				},
				Path:              "/series/Chernobyl",
				QualityProfileID:  1,
				LanguageProfileID: 1,
				SeasonFolder:      true,
				Monitored:         true,
				TvdbID:            360893,
				RootFolderPath:    "/series/",
				Tags: []int{
					11,
				},
				ID: 1,
			},
			ResponseBody: starr.BodyNotFound,
			WithError:    starr.ErrInvalidStatusCode,
			WithResponse: (*sonarr.Series)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.UpdateSeries(test.WithRequest.(*sonarr.AddSeriesInput), false)
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, output, test.WithResponse, "response is not the same as expected")
		})
	}
}

func TestDeleteSeries(t *testing.T) {
	t.Parallel()

	tests := []*starr.TestMockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "series/2?addImportListExclusion=false&deleteFiles=true"),
			ExpectedMethod: "DELETE",
			ResponseStatus: 200,
			ResponseBody:   "{}",
			WithRequest:    2,
			WithError:      nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "series/2?addImportListExclusion=false&deleteFiles=true"),
			ExpectedMethod: "DELETE",
			ResponseStatus: 404,
			ResponseBody:   starr.BodyNotFound,
			WithRequest:    2,
			WithError:      starr.ErrInvalidStatusCode,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			err := client.DeleteSeriesDefault(test.WithRequest.(int))
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
		})
	}
}

func TestLookupName(t *testing.T) {
	t.Parallel()

	loc, _ := time.LoadLocation("")
	date := time.Date(2019, 6, 4, 1, 0, 0, 0, loc)

	tests := []*starr.TestMockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "series/lookup?term=Chernobyl+Breaking+Bad"),
			ExpectedMethod: "GET",
			WithRequest:    "Chernobyl Breaking Bad",
			ResponseStatus: 200,
			ResponseBody:   listSeries,
			WithResponse: []*sonarr.Series{
				{
					Title:           "Breaking Bad",
					AlternateTitles: []*sonarr.AlternateTitle{},
					SortTitle:       "breaking bad",
					Status:          "ended",
					Ended:           true,
					Overview:        "Cool teacher becomes drug dealer",
					PreviousAiring:  date,
					Network:         "Netflix",
					AirTime:         "21:00",
					Images: []*starr.Image{
						{
							CoverType: "banner",
							URL:       "/MediaCover/1/banner.jpg?lastWrite=637829401993017870",
							RemoteURL: "https://artworks.thetvdb.com/banners/graphical/81189-g21.jpg",
						},
						{
							CoverType: "poster",
							URL:       "/MediaCover/1/poster.jpg?lastWrite=637829401994817870",
							RemoteURL: "https://artworks.thetvdb.com/banners/posters/81189-10.jpg",
						},
						{
							CoverType: "fanart",
							URL:       "/MediaCover/1/fanart.jpg?lastWrite=637829401996517870",
							RemoteURL: "https://artworks.thetvdb.com/banners/fanart/original/81189-21.jpg",
						},
					},
					Seasons: []*sonarr.Season{
						{
							SeasonNumber: 0,
							Monitored:    false,
							Statistics: &sonarr.Statistics{
								EpisodeFileCount:  0,
								EpisodeCount:      0,
								TotalEpisodeCount: 19,
								SizeOnDisk:        0,
								PercentOfEpisodes: 0.0,
							},
						},
						{
							SeasonNumber: 1,
							Monitored:    true,
							Statistics: &sonarr.Statistics{
								PreviousAiring:    date,
								EpisodeFileCount:  0,
								EpisodeCount:      7,
								TotalEpisodeCount: 7,
								SizeOnDisk:        0,
								PercentOfEpisodes: 0.0,
							},
						},
						{
							SeasonNumber: 2,
							Monitored:    true,
							Statistics: &sonarr.Statistics{
								PreviousAiring:    date,
								EpisodeFileCount:  0,
								EpisodeCount:      13,
								TotalEpisodeCount: 13,
								SizeOnDisk:        0,
								PercentOfEpisodes: 0.0,
							},
						},
						{
							SeasonNumber: 3,
							Monitored:    true,
							Statistics: &sonarr.Statistics{
								PreviousAiring:    date,
								EpisodeFileCount:  0,
								EpisodeCount:      13,
								TotalEpisodeCount: 13,
								SizeOnDisk:        0,
								PercentOfEpisodes: 0.0,
							},
						},
						{
							SeasonNumber: 4,
							Monitored:    true,
							Statistics: &sonarr.Statistics{
								PreviousAiring:    date,
								EpisodeFileCount:  0,
								EpisodeCount:      13,
								TotalEpisodeCount: 13,
								SizeOnDisk:        0,
								PercentOfEpisodes: 0.0,
							},
						},
						{
							SeasonNumber: 5,
							Monitored:    true,
							Statistics: &sonarr.Statistics{
								PreviousAiring:    date,
								EpisodeFileCount:  0,
								EpisodeCount:      16,
								TotalEpisodeCount: 16,
								SizeOnDisk:        0,
								PercentOfEpisodes: 0.0,
							},
						},
					},
					Year:              2008,
					Path:              "/series/Breaking Bad",
					QualityProfileID:  1,
					LanguageProfileID: 1,
					SeasonFolder:      true,
					Monitored:         true,
					UseSceneNumbering: false,
					Runtime:           47,
					TvdbID:            81189,
					TvRageID:          18164,
					TvMazeID:          169,
					FirstAired:        date,
					SeriesType:        "standard",
					CleanTitle:        "breakingbad",
					ImdbID:            "tt0903747",
					TitleSlug:         "breaking-bad",
					RootFolderPath:    "/series/",
					Certification:     "TV-MA",
					Genres: []string{
						"Crime",
						"Drama",
						"Suspense",
						"Thriller",
					},
					Tags: []int{
						11,
					},
					Added: date,
					Ratings: &starr.Ratings{
						Votes: 31714,
						Value: 9.4,
					},
					Statistics: &sonarr.Statistics{
						SeasonCount:       5,
						EpisodeFileCount:  0,
						EpisodeCount:      62,
						TotalEpisodeCount: 81,
						SizeOnDisk:        0,
						PercentOfEpisodes: 0.0,
					},
					ID: 1,
				},
				{
					Title:           "Chernobyl",
					AlternateTitles: []*sonarr.AlternateTitle{},
					SortTitle:       "chernobyl",
					Status:          "ended",
					Ended:           true,
					Overview:        "A lot of energy wasted",
					PreviousAiring:  date,
					Network:         "HBO",
					AirTime:         "21:00",
					Images: []*starr.Image{
						{
							CoverType: "banner",
							URL:       "/MediaCover/2/banner.jpg?lastWrite=637829402715717870",
							RemoteURL: "https://artworks.thetvdb.com/banners/graphical/5cc9f74c2ddd3.jpg",
						},
						{
							CoverType: "poster",
							URL:       "/MediaCover/2/poster.jpg?lastWrite=637829402718117870",
							RemoteURL: "https://artworks.thetvdb.com/banners/posters/5cc12861c93e4.jpg",
						},
						{
							CoverType: "fanart",
							URL:       "/MediaCover/2/fanart.jpg?lastWrite=637829402721317870",
							RemoteURL: "https://artworks.thetvdb.com/banners/series/360893/backgrounds/62017319.jpg",
						},
					},
					Seasons: []*sonarr.Season{
						{
							SeasonNumber: 0,
							Monitored:    false,
							Statistics: &sonarr.Statistics{
								EpisodeFileCount:  0,
								EpisodeCount:      0,
								TotalEpisodeCount: 14,
								SizeOnDisk:        0,
								PercentOfEpisodes: 0.0,
							},
						},
						{
							SeasonNumber: 1,
							Monitored:    true,
							Statistics: &sonarr.Statistics{
								PreviousAiring:    date,
								EpisodeFileCount:  0,
								EpisodeCount:      5,
								TotalEpisodeCount: 5,
								SizeOnDisk:        0,
								PercentOfEpisodes: 0.0,
							},
						},
					},
					Year:              2019,
					Path:              "/series/Chernobyl",
					QualityProfileID:  1,
					LanguageProfileID: 1,
					SeasonFolder:      true,
					Monitored:         true,
					UseSceneNumbering: false,
					Runtime:           65,
					TvdbID:            360893,
					TvRageID:          0,
					TvMazeID:          30770,
					FirstAired:        date,
					SeriesType:        "standard",
					CleanTitle:        "chernobyl",
					ImdbID:            "tt7366338",
					TitleSlug:         "chernobyl",
					RootFolderPath:    "/series/",
					Certification:     "TV-MA",
					Genres: []string{
						"Drama",
						"History",
						"Mini-Series",
						"Thriller",
					},
					Tags: []int{
						11,
					},
					Added: date,
					Ratings: &starr.Ratings{
						Votes: 83,
						Value: 8.7,
					},
					Statistics: &sonarr.Statistics{
						SeasonCount:       1,
						EpisodeFileCount:  0,
						EpisodeCount:      5,
						TotalEpisodeCount: 19,
						SizeOnDisk:        0,
						PercentOfEpisodes: 0.0,
					},
					ID: 2,
				},
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "series/lookup?term=Chernobyl+Breaking+Bad"),
			ExpectedMethod: "GET",
			WithRequest:    "Chernobyl Breaking Bad",
			ResponseStatus: 404,
			ResponseBody:   starr.BodyNotFound,
			WithError:      starr.ErrInvalidStatusCode,
			WithResponse:   []*sonarr.Series(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.Lookup(test.WithRequest.(string))
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, output, test.WithResponse, "response is not the same as expected")
		})
	}
}

func TestLookupID(t *testing.T) {
	t.Parallel()

	loc, _ := time.LoadLocation("")
	date := time.Date(2019, 6, 4, 1, 0, 0, 0, loc)

	tests := []*starr.TestMockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "series/lookup?term=tvdbid%3A360893"),
			ExpectedMethod: "GET",
			WithRequest:    360893,
			ResponseStatus: 200,
			ResponseBody:   "[" + secondSeries + "]",
			WithResponse: []*sonarr.Series{
				{
					Title:           "Chernobyl",
					AlternateTitles: []*sonarr.AlternateTitle{},
					SortTitle:       "chernobyl",
					Status:          "ended",
					Ended:           true,
					Overview:        "A lot of energy wasted",
					PreviousAiring:  date,
					Network:         "HBO",
					AirTime:         "21:00",
					Images: []*starr.Image{
						{
							CoverType: "banner",
							URL:       "/MediaCover/2/banner.jpg?lastWrite=637829402715717870",
							RemoteURL: "https://artworks.thetvdb.com/banners/graphical/5cc9f74c2ddd3.jpg",
						},
						{
							CoverType: "poster",
							URL:       "/MediaCover/2/poster.jpg?lastWrite=637829402718117870",
							RemoteURL: "https://artworks.thetvdb.com/banners/posters/5cc12861c93e4.jpg",
						},
						{
							CoverType: "fanart",
							URL:       "/MediaCover/2/fanart.jpg?lastWrite=637829402721317870",
							RemoteURL: "https://artworks.thetvdb.com/banners/series/360893/backgrounds/62017319.jpg",
						},
					},
					Seasons: []*sonarr.Season{
						{
							SeasonNumber: 0,
							Monitored:    false,
							Statistics: &sonarr.Statistics{
								EpisodeFileCount:  0,
								EpisodeCount:      0,
								TotalEpisodeCount: 14,
								SizeOnDisk:        0,
								PercentOfEpisodes: 0.0,
							},
						},
						{
							SeasonNumber: 1,
							Monitored:    true,
							Statistics: &sonarr.Statistics{
								PreviousAiring:    date,
								EpisodeFileCount:  0,
								EpisodeCount:      5,
								TotalEpisodeCount: 5,
								SizeOnDisk:        0,
								PercentOfEpisodes: 0.0,
							},
						},
					},
					Year:              2019,
					Path:              "/series/Chernobyl",
					QualityProfileID:  1,
					LanguageProfileID: 1,
					SeasonFolder:      true,
					Monitored:         true,
					UseSceneNumbering: false,
					Runtime:           65,
					TvdbID:            360893,
					TvRageID:          0,
					TvMazeID:          30770,
					FirstAired:        date,
					SeriesType:        "standard",
					CleanTitle:        "chernobyl",
					ImdbID:            "tt7366338",
					TitleSlug:         "chernobyl",
					RootFolderPath:    "/series/",
					Certification:     "TV-MA",
					Genres: []string{
						"Drama",
						"History",
						"Mini-Series",
						"Thriller",
					},
					Tags: []int{
						11,
					},
					Added: date,
					Ratings: &starr.Ratings{
						Votes: 83,
						Value: 8.7,
					},
					Statistics: &sonarr.Statistics{
						SeasonCount:       1,
						EpisodeFileCount:  0,
						EpisodeCount:      5,
						TotalEpisodeCount: 19,
						SizeOnDisk:        0,
						PercentOfEpisodes: 0.0,
					},
					ID: 2,
				},
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "series/lookup?term=tvdbid%3A360893"),
			ExpectedMethod: "GET",
			WithRequest:    360893,
			ResponseStatus: 404,
			ResponseBody:   starr.BodyNotFound,
			WithError:      starr.ErrInvalidStatusCode,
			WithResponse:   []*sonarr.Series(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetSeriesLookup("", int64(test.WithRequest.(int)))
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, output, test.WithResponse, "response is not the same as expected")
		})
	}
}
