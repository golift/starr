package sonarr_test

import (
	"net/http"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"golift.io/starr"
	"golift.io/starr/sonarr"
	"golift.io/starr/starrtest"
)

const namingBody = `{
	"colonReplacementFormat": 0,
	"renameEpisodes": false,
	"replaceIllegalCharacters": true,
	"multiEpisodeStyle": 0,
	"standardEpisodeFormat": "{Series Title} - S{season:00}E{episode:00} - {Episode Title} {Quality Full}",
	"dailyEpisodeFormat": "{Series Title} - {Air-Date} - {Episode Title} {Quality Full}",
	"animeEpisodeFormat": "{Series Title} - S{season:00}E{episode:00} - {Episode Title} {Quality Full}",
	"seriesFolderFormat": "{Series Title}",
	"seasonFolderFormat": "Season {season}",
	"specialsFolderFormat": "Specials",
	"id": 1
}`

func TestGetNaming(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "config", "naming"),
			ExpectedMethod: "GET",
			ResponseStatus: 200,
			ResponseBody:   namingBody,
			WithResponse: &sonarr.Naming{
				ID:                       1,
				RenameEpisodes:           false,
				ReplaceIllegalCharacters: true,
				MultiEpisodeStyle:        0,
				StandardEpisodeFormat:    "{Series Title} - S{season:00}E{episode:00} - {Episode Title} {Quality Full}",
				DailyEpisodeFormat:       "{Series Title} - {Air-Date} - {Episode Title} {Quality Full}",
				AnimeEpisodeFormat:       "{Series Title} - S{season:00}E{episode:00} - {Episode Title} {Quality Full}",
				SeriesFolderFormat:       "{Series Title}",
				SeasonFolderFormat:       "Season {season}",
				SpecialsFolderFormat:     "Specials",
				ColonReplacementFormat:   sonarr.ColonDelete,
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "config", "naming"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:   (*sonarr.Naming)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetNaming()
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestUpdateNaming(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "202",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "config", "naming"),
			ExpectedMethod: "PUT",
			ResponseStatus: 202,
			WithRequest: &sonarr.Naming{
				ReplaceIllegalCharacters: true,
			},
			ExpectedRequest: `{"replaceIllegalCharacters":true,"id":1}` + "\n",
			ResponseBody:    namingBody,
			WithResponse: &sonarr.Naming{
				ID:                       1,
				RenameEpisodes:           false,
				ReplaceIllegalCharacters: true,
				MultiEpisodeStyle:        0,
				StandardEpisodeFormat:    "{Series Title} - S{season:00}E{episode:00} - {Episode Title} {Quality Full}",
				DailyEpisodeFormat:       "{Series Title} - {Air-Date} - {Episode Title} {Quality Full}",
				AnimeEpisodeFormat:       "{Series Title} - S{season:00}E{episode:00} - {Episode Title} {Quality Full}",
				SeriesFolderFormat:       "{Series Title}",
				SeasonFolderFormat:       "Season {season}",
				SpecialsFolderFormat:     "Specials",
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "config", "naming"),
			ExpectedMethod: "PUT",
			WithRequest: &sonarr.Naming{
				ReplaceIllegalCharacters: true,
			},
			ExpectedRequest: `{"replaceIllegalCharacters":true,"id":1}` + "\n",
			ResponseStatus:  404,
			ResponseBody:    `{"message": "NotFound"}`,
			WithError:       &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:    (*sonarr.Naming)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.UpdateNaming(test.WithRequest.(*sonarr.Naming))
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}
