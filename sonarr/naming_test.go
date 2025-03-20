package sonarr_test

import (
	"net/http"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	"customcolonReplacementFormat": "",
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
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetNaming()
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
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
				RenameEpisodes:               true,
				ReplaceIllegalCharacters:     true,
				ColonReplacementFormat:       5,
				ID:                           4,
				MultiEpisodeStyle:            2,
				DailyEpisodeFormat:           "a",
				AnimeEpisodeFormat:           "b",
				SeriesFolderFormat:           "c",
				SeasonFolderFormat:           "d",
				SpecialsFolderFormat:         "e",
				StandardEpisodeFormat:        "f",
				CustomColonReplacementFormat: "g",
			},
			ExpectedRequest: `{"renameEpisodes":true,"replaceIllegalCharacters":true,"colonReplacementFormat":5,` +
				`"id":1,"multiEpisodeStyle":2,"dailyEpisodeFormat":"a","animeEpisodeFormat":"b","seriesFolderFormat":"c",` +
				`"seasonFolderFormat":"d","specialsFolderFormat":"e","standardEpisodeFormat":"f","customColonReplacementFormat":"g"}` + "\n",
			ResponseBody: namingBody,
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
			ExpectedRequest: `{"renameEpisodes":false,"replaceIllegalCharacters":true,"colonReplacementFormat":0,` +
				`"id":1,"multiEpisodeStyle":0,"dailyEpisodeFormat":"","animeEpisodeFormat":"","seriesFolderFormat":"",` +
				`"seasonFolderFormat":"","specialsFolderFormat":"","standardEpisodeFormat":"","customColonReplacementFormat":""}` + "\n",
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:   (*sonarr.Naming)(nil),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.UpdateNaming(test.WithRequest.(*sonarr.Naming))
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}
