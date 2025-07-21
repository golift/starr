package sonarr_test

import (
	"net/http"
	"path"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golift.io/starr"
	"golift.io/starr/sonarr"
	"golift.io/starr/starrtest"
)

const (
	episode = `{
		"seriesId": 53,
		"tvdbId": 8444132,
		"episodeFileId": 3260,
		"seasonNumber": 1,
		"episodeNumber": 1,
		"title": "When You're Lost in the Darkness",
		"airDate": "2023-01-15",
		"airDateUtc": "2023-01-16T02:00:00Z",
		"overview": "Twenty years after a fungal outbreak ravages the planet, survivors Joel and Tess are tasked with a mission that could change everything.",
		"episodeFile": {
			"seriesId": 53,
			"seasonNumber": 1,
			"relativePath": "Season 01/The Last of Us (2023) - S01E01 - When You're Lost in the Darkness [AMZN WEBDL-1080p Proper][EAC3 Atmos 5.1][h264]-FLUX.mkv",
			"path": "/media/tv/The Last of Us (2023) [imdb-tt3581920] [tvdbid-392256]/Season 01/The Last of Us (2023) - S01E01 - When You're Lost in the Darkness [AMZN WEBDL-1080p Proper][EAC3 Atmos 5.1][h264]-FLUX.mkv",
			"size": 6002269811,
			"dateAdded": "2023-03-14T03:34:28Z",
			"sceneName": "The.Last.Of.Us.S01E01.When.Your.Lost.in.the.Darkness.REPACK.1080p.AMZN.WEB-DL.DDP5.1.Atmos.H.264-FLUX",
			"releaseGroup": "FLUX",
			"languages": [
				{
					"id": 1,
					"name": "English"
				}
			],
			"quality": {
				"quality": {
					"id": 3,
					"name": "WEBDL-1080p",
					"source": "web",
					"resolution": 1080
				},
				"revision": {
					"version": 2,
					"real": 0,
					"isRepack": true
				}
			},
			"customFormats": [
				{
					"id": 26,
					"name": "AMZN"
				},
				{
					"id": 46,
					"name": "WEB Tier 01"
				},
				{
					"id": 49,
					"name": "Repack/Proper"
				}
			],
			"mediaInfo": {
				"audioBitrate": 768000,
				"audioChannels": 5.1,
				"audioCodec": "EAC3 Atmos",
				"audioLanguages": "eng",
				"audioStreamCount": 1,
				"videoBitDepth": 8,
				"videoBitrate": 0,
				"videoCodec": "h264",
				"videoFps": 23.976,
				"videoDynamicRange": "",
				"videoDynamicRangeType": "",
				"resolution": "1920x1080",
				"runTime": "1:20:56",
				"scanType": "Progressive",
				"subtitles": "eng/eng/spa/spa"
			},
			"qualityCutoffNotMet": false,
			"id": 3260
		},
		"hasFile": true,
		"monitored": true,
		"unverifiedSceneNumbering": false,
		"series": {
			"title": "The Last of Us",
			"sortTitle": "last of us",
			"status": "continuing",
			"ended": false,
			"overview": "After a global pandemic destroys civilization, a hardened survivor takes charge of a 14-year-old girl who may be humanity’s last hope.",
			"network": "HBO",
			"airTime": "21:00",
			"images": [
				{
					"coverType": "banner",
					"remoteUrl": "https://artworks.thetvdb.com/banners/v4/series/392256/banners/63cded40a22f4.jpg"
				},
				{
					"coverType": "poster",
					"remoteUrl": "https://artworks.thetvdb.com/banners/v4/series/392256/posters/6362e8b41ca10.jpg"
				},
				{
					"coverType": "fanart",
					"remoteUrl": "https://artworks.thetvdb.com/banners/v4/series/392256/backgrounds/637e513564daf.jpg"
				},
				{
					"coverType": "unknown",
					"remoteUrl": "https://artworks.thetvdb.com/banners/v4/series/392256/clearlogo/636712bfd63b2.png"
				}
			],
			"originalLanguage": {
				"id": 1,
				"name": "English"
			},
			"seasons": [
				{
					"seasonNumber": 0,
					"monitored": false
				},
				{
					"seasonNumber": 1,
					"monitored": true
				}
			],
			"year": 2023,
			"path": "/media/tv/The Last of Us (2023) [imdb-tt3581920] [tvdbid-392256]",
			"qualityProfileId": 7,
			"seasonFolder": true,
			"monitored": true,
			"useSceneNumbering": false,
			"runtime": 58,
			"tvdbId": 392256,
			"tvRageId": 0,
			"tvMazeId": 46562,
			"firstAired": "2023-01-15T00:00:00Z",
			"seriesType": "standard",
			"cleanTitle": "thelastus",
			"imdbId": "tt3581920",
			"titleSlug": "the-last-of-us",
			"certification": "TV-MA",
			"genres": [
				"Action",
				"Adventure",
				"Drama",
				"Horror",
				"Science Fiction"
			],
			"tags": [],
			"added": "2023-01-17T21:14:56Z",
			"ratings": {
				"votes": 0,
				"value": 0
			},
			"languageProfileId": 1,
			"id": 53
		},
		"images": [
			{
				"coverType": "screenshot",
				"remoteUrl": "https://artworks.thetvdb.com/banners/v4/episode/8444132/screencap/638bf7ef8ed12.jpg"
			}
		],
		"grabbed": false,
		"id": 4186
	}`
)

func TestGetEpisodeByID(t *testing.T) {
	t.Parallel()

	loc, _ := time.LoadLocation("")
	airDate := time.Date(2023, 1, 16, 2, 0, 0, 0, loc)
	addedDate := time.Date(2023, 1, 17, 21, 14, 56, 0, loc)
	firstAiredDate := time.Date(2023, 1, 15, 0, 0, 0, 0, loc)

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "episode", "4186"),
			ExpectedMethod: "GET",
			ResponseStatus: 200,
			ResponseBody:   episode,
			WithRequest:    4186,
			WithResponse: &sonarr.Episode{
				ID:            4186,
				SeriesID:      53,
				TvdbID:        8444132,
				EpisodeFileID: 3260,
				SeasonNumber:  1,
				EpisodeNumber: 1,
				AirDate:       "2023-01-15",
				AirDateUtc:    airDate,
				Title:         "When You're Lost in the Darkness",
				Overview:      "Twenty years after a fungal outbreak ravages the planet, survivors Joel and Tess are tasked with a mission that could change everything.",
				HasFile:       true,
				Monitored:     true,
				Images: []*starr.Image{
					{
						CoverType: "screenshot",
						RemoteURL: "https://artworks.thetvdb.com/banners/v4/episode/8444132/screencap/638bf7ef8ed12.jpg",
					},
				},
				Series: &sonarr.Series{
					ID:                53,
					Ended:             false,
					Monitored:         true,
					SeasonFolder:      true,
					Runtime:           58,
					Year:              2023,
					LanguageProfileID: 1,
					QualityProfileID:  7,
					TvdbID:            392256,
					TvMazeID:          46562,
					TvRageID:          0,
					AirTime:           "21:00",
					Certification:     "TV-MA",
					CleanTitle:        "thelastus",
					ImdbID:            "tt3581920",
					Network:           "HBO",
					Overview:          "After a global pandemic destroys civilization, a hardened survivor takes charge of a 14-year-old girl who may be humanity’s last hope.",
					Path:              "/media/tv/The Last of Us (2023) [imdb-tt3581920] [tvdbid-392256]",
					SeriesType:        "standard",
					SortTitle:         "last of us",
					Status:            "continuing",
					Title:             "The Last of Us",
					TitleSlug:         "the-last-of-us",
					Added:             addedDate,
					FirstAired:        firstAiredDate,
					Ratings:           &starr.Ratings{},
					Tags:              []int{},
					Genres: []string{
						"Action",
						"Adventure",
						"Drama",
						"Horror",
						"Science Fiction",
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
					Images: []*starr.Image{
						{
							CoverType: "banner",
							RemoteURL: "https://artworks.thetvdb.com/banners/v4/series/392256/banners/63cded40a22f4.jpg",
						},
						{
							CoverType: "poster",
							RemoteURL: "https://artworks.thetvdb.com/banners/v4/series/392256/posters/6362e8b41ca10.jpg",
						},
						{
							CoverType: "fanart",
							RemoteURL: "https://artworks.thetvdb.com/banners/v4/series/392256/backgrounds/637e513564daf.jpg",
						},
						{
							CoverType: "unknown",
							RemoteURL: "https://artworks.thetvdb.com/banners/v4/series/392256/clearlogo/636712bfd63b2.png",
						},
					},
				},
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "episode", "4186"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			ResponseBody:   starrtest.BodyNotFound,
			WithRequest:    4186,
			WithError:      &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:   (*sonarr.Episode)(nil),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetEpisodeByID(int64(test.WithRequest.(int)))
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, output, test.WithResponse, "response is not the same as expected")
		})
	}
}
