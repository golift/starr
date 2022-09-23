package radarr_test

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"golift.io/starr"
	"golift.io/starr/radarr"
)

const mediaManagementBody = `{
  "autoUnmonitorPreviouslyDownloadedMovies": false,
  "recycleBin": "",
  "recycleBinCleanupDays": 7,
  "downloadPropersAndRepacks": "preferAndUpgrade",
  "createEmptyMovieFolders": false,
  "deleteEmptyFolders": false,
  "fileDate": "none",
  "rescanAfterRefresh": "always",
  "autoRenameFolders": false,
  "pathsDefaultStatic": false,
  "setPermissionsLinux": false,
  "chmodFolder": "755",
  "chownGroup": "",
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
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "config", "mediaManagement"),
			ExpectedMethod: "GET",
			ResponseStatus: 200,
			ResponseBody:   mediaManagementBody,
			WithResponse: &radarr.MediaManagement{
				AutoRenameFolders:                       false,
				AutoUnmonitorPreviouslyDownloadedMovies: false,
				CopyUsingHardlinks:                      true,
				CreateEmptyMovieFolders:                 false,
				DeleteEmptyFolders:                      false,
				EnableMediaInfo:                         true,
				ImportExtraFiles:                        false,
				PathsDefaultStatic:                      false,
				SetPermissionsLinux:                     false,
				SkipFreeSpaceCheckWhenImporting:         false,
				ID:                                      1,
				MinimumFreeSpaceWhenImporting:           100,
				RecycleBinCleanupDays:                   7,
				ChmodFolder:                             "755",
				ChownGroup:                              "",
				DownloadPropersAndRepacks:               "preferAndUpgrade",
				ExtraFileExtensions:                     "srt",
				FileDate:                                "none",
				RecycleBin:                              "",
				RescanAfterRefresh:                      "always",
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "config", "mediaManagement"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      starr.ErrInvalidStatusCode,
			WithResponse:   (*radarr.MediaManagement)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := radarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
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
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "config", "mediaManagement"),
			ExpectedMethod: "PUT",
			ResponseStatus: 202,
			WithRequest: &radarr.MediaManagement{
				EnableMediaInfo: true,
			},
			ExpectedRequest: `{"enableMediaInfo":true}` + "\n",
			ResponseBody:    mediaManagementBody,
			WithResponse: &radarr.MediaManagement{
				AutoRenameFolders:                       false,
				AutoUnmonitorPreviouslyDownloadedMovies: false,
				CopyUsingHardlinks:                      true,
				CreateEmptyMovieFolders:                 false,
				DeleteEmptyFolders:                      false,
				EnableMediaInfo:                         true,
				ImportExtraFiles:                        false,
				PathsDefaultStatic:                      false,
				SetPermissionsLinux:                     false,
				SkipFreeSpaceCheckWhenImporting:         false,
				ID:                                      1,
				MinimumFreeSpaceWhenImporting:           100,
				RecycleBinCleanupDays:                   7,
				ChmodFolder:                             "755",
				ChownGroup:                              "",
				DownloadPropersAndRepacks:               "preferAndUpgrade",
				ExtraFileExtensions:                     "srt",
				FileDate:                                "none",
				RecycleBin:                              "",
				RescanAfterRefresh:                      "always",
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "config", "mediaManagement"),
			ExpectedMethod: "PUT",
			WithRequest: &radarr.MediaManagement{
				EnableMediaInfo: true,
			},
			ExpectedRequest: `{"enableMediaInfo":true}` + "\n",
			ResponseStatus:  404,
			ResponseBody:    `{"message": "NotFound"}`,
			WithError:       starr.ErrInvalidStatusCode,
			WithResponse:    (*radarr.MediaManagement)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := radarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.UpdateMediaManagement(test.WithRequest.(*radarr.MediaManagement))
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}
