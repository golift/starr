package radarr_test

import (
	"net/http"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golift.io/starr"
	"golift.io/starr/radarr"
	"golift.io/starr/starrtest"
)

const downloadClientConfigBody = `{
    "downloadClientWorkingFolders": "_UNPACK_|_FAILED_",
    "enableCompletedDownloadHandling": true,
	"checkForFinishedDownloadInterval": 1,
	"autoRedownloadFailed": false,
    "id": 1
}`

func TestGetDownloadClientConfig(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:            "200",
			ExpectedPath:    path.Join("/", starr.API, radarr.APIver, "config", "downloadClient"),
			ExpectedRequest: "",
			ExpectedMethod:  "GET",
			ResponseStatus:  200,
			ResponseBody:    downloadClientConfigBody,
			WithRequest:     nil,
			WithResponse: &radarr.DownloadClientConfig{
				EnableCompletedDownloadHandling:  true,
				AutoRedownloadFailed:             false,
				CheckForFinishedDownloadInterval: 1,
				ID:                               1,
				DownloadClientWorkingFolders:     "_UNPACK_|_FAILED_",
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "config", "downloadClient"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:   (*radarr.DownloadClientConfig)(nil),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := radarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetDownloadClientConfig()
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestUpdateDownloadClientConfig(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "202",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "config", "downloadClient", "1"),
			ExpectedMethod: "PUT",
			ResponseStatus: 202,
			WithRequest: &radarr.DownloadClientConfig{
				EnableCompletedDownloadHandling:  true,
				AutoRedownloadFailed:             false,
				CheckForFinishedDownloadInterval: 1,
				ID:                               1,
				DownloadClientWorkingFolders:     "_UNPACK_|_FAILED_",
			},
			ExpectedRequest: `{"enableCompletedDownloadHandling":true,"autoRedownloadFailed":false,` +
				`"checkForFinishedDownloadInterval":1,"id":1,"downloadClientWorkingFolders":"_UNPACK_|_FAILED_"}` + "\n",
			ResponseBody: downloadClientConfigBody,
			WithResponse: &radarr.DownloadClientConfig{
				EnableCompletedDownloadHandling:  true,
				AutoRedownloadFailed:             false,
				CheckForFinishedDownloadInterval: 1,
				ID:                               1,
				DownloadClientWorkingFolders:     "_UNPACK_|_FAILED_",
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "config", "downloadClient", "1"),
			ExpectedMethod: "PUT",
			WithRequest: &radarr.DownloadClientConfig{
				EnableCompletedDownloadHandling:  true,
				AutoRedownloadFailed:             false,
				CheckForFinishedDownloadInterval: 1,
				ID:                               1,
				DownloadClientWorkingFolders:     "_UNPACK_|_FAILED_",
			},
			ExpectedRequest: `{"enableCompletedDownloadHandling":true,"autoRedownloadFailed":false,` +
				`"checkForFinishedDownloadInterval":1,"id":1,"downloadClientWorkingFolders":"_UNPACK_|_FAILED_"}` + "\n",
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:   (*radarr.DownloadClientConfig)(nil),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := radarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.UpdateDownloadClientConfig(test.WithRequest.(*radarr.DownloadClientConfig))
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, output, test.WithResponse, "response is not the same as expected")
		})
	}
}
