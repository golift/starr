package sonarr_test

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"golift.io/starr"
	"golift.io/starr/sonarr"
)

const mediaManagementBody = `{
	"autoUnmonitorPreviouslyDownloadedEpisodes": false,
	"recycleBin": "",
	"recycleBinCleanupDays": 7,
	"downloadPropersAndRepacks": "preferAndUpgrade",
	"createEmptySeriesFolders": false,
	"deleteEmptyFolders": false,
	"fileDate": "none",
	"rescanAfterRefresh": "always",
	"setPermissionsLinux": false,
	"chmodFolder": "755",
	"chownGroup": "",
	"episodeTitleRequired": "always",
	"skipFreeSpaceCheckWhenImporting": false,
	"minimumFreeSpaceWhenImporting": 100,
	"copyUsingHardlinks": true,
	"importExtraFiles": false,
	"extraFileExtensions": "srt",
	"enableMediaInfo": true,
	"id": 1
  }`

func TestGetMediaManagement(t *testing.T) {
	t.Parallel()

	tests := []*starr.TestMockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "config", "mediaManagement"),
			ExpectedMethod: "GET",
			ResponseStatus: 200,
			ResponseBody:   mediaManagementBody,
			WithResponse: &sonarr.MediaManagement{
				ID: 1,
				AutoUnmonitorPreviouslyDownloadedEpisodes: false,
				RecycleBin:                      "",
				RecycleBinCleanupDays:           7,
				DownloadPropersAndRepacks:       "preferAndUpgrade",
				CreateEmptySeriesFolders:        false,
				DeleteEmptyFolders:              false,
				FileDate:                        "none",
				RescanAfterRefresh:              "always",
				SetPermissionsLinux:             false,
				ChmodFolder:                     "755",
				ChownGroup:                      "",
				EpisodeTitleRequired:            "always",
				SkipFreeSpaceCheckWhenImporting: false,
				MinimumFreeSpaceWhenImporting:   100,
				CopyUsingHardlinks:              true,
				ImportExtraFiles:                false,
				ExtraFileExtensions:             "srt",
				EnableMediaInfo:                 true,
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "config", "mediaManagement"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      starr.ErrInvalidStatusCode,
			WithResponse:   (*sonarr.MediaManagement)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetMediaManagement()
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestUpdateMediaManagement(t *testing.T) {
	t.Parallel()

	tests := []*starr.TestMockData{
		{
			Name:           "202",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "config", "mediaManagement"),
			ExpectedMethod: "PUT",
			ResponseStatus: 202,
			WithRequest: &sonarr.MediaManagement{
				EnableMediaInfo: true,
			},
			ExpectedRequest: `{"enableMediaInfo":true}` + "\n",
			ResponseBody:    mediaManagementBody,
			WithResponse: &sonarr.MediaManagement{
				ID: 1,
				AutoUnmonitorPreviouslyDownloadedEpisodes: false,
				RecycleBin:                      "",
				RecycleBinCleanupDays:           7,
				DownloadPropersAndRepacks:       "preferAndUpgrade",
				CreateEmptySeriesFolders:        false,
				DeleteEmptyFolders:              false,
				FileDate:                        "none",
				RescanAfterRefresh:              "always",
				SetPermissionsLinux:             false,
				ChmodFolder:                     "755",
				ChownGroup:                      "",
				EpisodeTitleRequired:            "always",
				SkipFreeSpaceCheckWhenImporting: false,
				MinimumFreeSpaceWhenImporting:   100,
				CopyUsingHardlinks:              true,
				ImportExtraFiles:                false,
				ExtraFileExtensions:             "srt",
				EnableMediaInfo:                 true,
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "config", "mediaManagement"),
			ExpectedMethod: "PUT",
			WithRequest: &sonarr.MediaManagement{
				EnableMediaInfo: true,
			},
			ExpectedRequest: `{"enableMediaInfo":true}` + "\n",
			ResponseStatus:  404,
			ResponseBody:    `{"message": "NotFound"}`,
			WithError:       starr.ErrInvalidStatusCode,
			WithResponse:    (*sonarr.MediaManagement)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.UpdateMediaManagement(test.WithRequest.(*sonarr.MediaManagement))
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}
