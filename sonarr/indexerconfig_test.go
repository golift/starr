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

const indexerConfigBody = `{
	"minimumAge": 0,
	"retention": 0,
	"maximumSize": 0,
	"rssSyncInterval": 20,
	"id": 1
}`

func TestGetIndexerConfig(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "config", "indexer"),
			ExpectedMethod: "GET",
			ResponseStatus: 200,
			ResponseBody:   indexerConfigBody,
			WithResponse: &sonarr.IndexerConfig{
				ID:              1,
				MaximumSize:     0,
				MinimumAge:      0,
				Retention:       0,
				RssSyncInterval: 20,
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "config", "indexer"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:   (*sonarr.IndexerConfig)(nil),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetIndexerConfig()
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestUpdateIndexerConfig(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "202",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "config", "indexer", "1"),
			ExpectedMethod: "PUT",
			ResponseStatus: 202,
			WithRequest: &sonarr.IndexerConfig{
				ID:              1,
				MaximumSize:     0,
				MinimumAge:      0,
				Retention:       0,
				RssSyncInterval: 20,
			},
			ExpectedRequest: `{"id":1,"maximumSize":0,"minimumAge":0,"retention":0,"rssSyncInterval":20}` + "\n",
			ResponseBody:    indexerConfigBody,
			WithResponse: &sonarr.IndexerConfig{
				ID:              1,
				MaximumSize:     0,
				MinimumAge:      0,
				Retention:       0,
				RssSyncInterval: 20,
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "config", "indexer", "1"),
			ExpectedMethod: "PUT",
			WithRequest: &sonarr.IndexerConfig{
				ID:              1,
				MaximumSize:     0,
				MinimumAge:      0,
				Retention:       0,
				RssSyncInterval: 20,
			},
			ExpectedRequest: `{"id":1,"maximumSize":0,"minimumAge":0,"retention":0,"rssSyncInterval":20}` + "\n",
			ResponseStatus:  404,
			ResponseBody:    `{"message": "NotFound"}`,
			WithError:       &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:    (*sonarr.IndexerConfig)(nil),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.UpdateIndexerConfig(test.WithRequest.(*sonarr.IndexerConfig))
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, output, test.WithResponse, "response is not the same as expected")
		})
	}
}
