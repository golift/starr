package radarr_test

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"golift.io/starr"
	"golift.io/starr/radarr"
)

const namingBody = `{
	"renameMovies": true,
	"replaceIllegalCharacters": true,
	"colonReplacementFormat": "delete",
	"standardMovieFormat": "{Movie.Title}.{Release.Year}.{Quality.Title}",
	"movieFolderFormat": "{Movie Title} ({Release Year})",
	"includeQuality": true,
	"replaceSpaces": true,
	"id": 1
  }`

func TestGetNaming(t *testing.T) {
	t.Parallel()

	tests := []*starr.TestMockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "config", "naming"),
			ExpectedMethod: "GET",
			ResponseStatus: 200,
			ResponseBody:   namingBody,
			WithResponse: &radarr.Naming{
				ID:                       1,
				ReplaceIllegalCharacters: true,
				IncludeQuality:           true,
				ReplaceSpaces:            true,
				RenameMovies:             true,
				ColonReplacementFormat:   "delete",
				StandardMovieFormat:      "{Movie.Title}.{Release.Year}.{Quality.Title}",
				MovieFolderFormat:        "{Movie Title} ({Release Year})",
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "config", "naming"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      starr.ErrInvalidStatusCode,
			WithResponse:   (*radarr.Naming)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := radarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetNaming()
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestUpdateNaming(t *testing.T) {
	t.Parallel()

	tests := []*starr.TestMockData{
		{
			Name:           "202",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "config", "naming"),
			ExpectedMethod: "PUT",
			ResponseStatus: 202,
			WithRequest: &radarr.Naming{
				ID:                       1,
				ReplaceIllegalCharacters: true,
			},
			ExpectedRequest: `{"replaceIllegalCharacters":true,"id":1,"standardMovieFormat":"","movieFolderFormat":""}` + "\n",
			ResponseBody:    namingBody,
			WithResponse: &radarr.Naming{
				ID:                       1,
				ReplaceIllegalCharacters: true,
				IncludeQuality:           true,
				ReplaceSpaces:            true,
				RenameMovies:             true,
				ColonReplacementFormat:   "delete",
				StandardMovieFormat:      "{Movie.Title}.{Release.Year}.{Quality.Title}",
				MovieFolderFormat:        "{Movie Title} ({Release Year})",
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "config", "naming"),
			ExpectedMethod: "PUT",
			WithRequest: &radarr.Naming{
				ID:                       1,
				ReplaceIllegalCharacters: true,
			},
			ExpectedRequest: `{"replaceIllegalCharacters":true,"id":1,"standardMovieFormat":"","movieFolderFormat":""}` + "\n",
			ResponseStatus:  404,
			ResponseBody:    `{"message": "NotFound"}`,
			WithError:       starr.ErrInvalidStatusCode,
			WithResponse:    (*radarr.Naming)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := radarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.UpdateNaming(test.WithRequest.(*radarr.Naming))
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "test.WithResponse does not match the actual response")
		})
	}
}
