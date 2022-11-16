package readarr_test

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"golift.io/starr"
	"golift.io/starr/readarr"
)

const downloadClientConfigBody = `{
	"downloadClientWorkingFolders": "_UNPACK_|_FAILED_",
	"enableCompletedDownloadHandling": true,
	"removeCompletedDownloads": false,
	"autoRedownloadFailed": false,
	"removeFailedDownloads": false,
	"id": 1
  }`

const updateDownloadClientConfig = `{"enableCompletedDownloadHandling":true,"autoRedownloadFailed":false,"id":1,` +
	`"downloadClientWorkingFolders":"_UNPACK_|_FAILED_","removeCompletedDownloads":false,"removeFailedDownloads":false}`

func TestGetDownloadClientConfig(t *testing.T) {
	t.Parallel()

	tests := []*starr.TestMockData{
		{
			Name:            "200",
			ExpectedPath:    path.Join("/", starr.API, readarr.APIver, "config", "downloadClient"),
			ExpectedRequest: "",
			ExpectedMethod:  "GET",
			ResponseStatus:  200,
			ResponseBody:    downloadClientConfigBody,
			WithRequest:     nil,
			WithResponse: &readarr.DownloadClientConfig{
				EnableCompletedDownloadHandling: true,
				AutoRedownloadFailed:            false,
				RemoveCompletedDownloads:        false,
				RemoveFailedDownloads:           false,
				ID:                              1,
				DownloadClientWorkingFolders:    "_UNPACK_|_FAILED_",
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, readarr.APIver, "config", "downloadClient"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      starr.ErrInvalidStatusCode,
			WithResponse:   (*readarr.DownloadClientConfig)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := readarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetDownloadClientConfig()
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestUpdateDownloadClientConfig(t *testing.T) {
	t.Parallel()

	tests := []*starr.TestMockData{
		{
			Name:           "202",
			ExpectedPath:   path.Join("/", starr.API, readarr.APIver, "config", "downloadClient", "1"),
			ExpectedMethod: "PUT",
			ResponseStatus: 202,
			WithRequest: &readarr.DownloadClientConfig{
				EnableCompletedDownloadHandling: true,
				AutoRedownloadFailed:            false,
				RemoveCompletedDownloads:        false,
				RemoveFailedDownloads:           false,
				ID:                              1,
				DownloadClientWorkingFolders:    "_UNPACK_|_FAILED_",
			},
			ExpectedRequest: updateDownloadClientConfig + "\n",
			ResponseBody:    downloadClientConfigBody,
			WithResponse: &readarr.DownloadClientConfig{
				EnableCompletedDownloadHandling: true,
				AutoRedownloadFailed:            false,
				RemoveCompletedDownloads:        false,
				RemoveFailedDownloads:           false,
				ID:                              1,
				DownloadClientWorkingFolders:    "_UNPACK_|_FAILED_",
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, readarr.APIver, "config", "downloadClient", "1"),
			ExpectedMethod: "PUT",
			WithRequest: &readarr.DownloadClientConfig{
				EnableCompletedDownloadHandling: true,
				AutoRedownloadFailed:            false,
				RemoveCompletedDownloads:        false,
				RemoveFailedDownloads:           false,
				ID:                              1,
				DownloadClientWorkingFolders:    "_UNPACK_|_FAILED_",
			},
			ExpectedRequest: updateDownloadClientConfig + "\n",
			ResponseStatus:  404,
			ResponseBody:    `{"message": "NotFound"}`,
			WithError:       starr.ErrInvalidStatusCode,
			WithResponse:    (*readarr.DownloadClientConfig)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := readarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.UpdateDownloadClientConfig(test.WithRequest.(*readarr.DownloadClientConfig))
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, output, test.WithResponse, "response is not the same as expected")
		})
	}
}
