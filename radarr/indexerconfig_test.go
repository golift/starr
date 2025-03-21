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

const indexerConfigBody = `{
	"minimumAge": 0,
	"maximumSize": 0,
	"retention": 0,
	"rssSyncInterval": 60,
	"preferIndexerFlags": false,
	"availabilityDelay": 0,
	"allowHardcodedSubs": false,
	"whitelistedHardcodedSubs": "",
	"id": 1
  }`

const indexerRequest = `{"whitelistedHardcodedSubs":"","id":1,"maximumSize":0,"minimumAge":0,"retention":0,` +
	`"rssSyncInterval":60,"availabilityDelay":0,"preferIndexerFlags":false,"allowHardcodedSubs":false}` + "\n"

func TestGetIndexerConfig(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "config", "indexer"),
			ExpectedMethod: "GET",
			ResponseStatus: 200,
			ResponseBody:   indexerConfigBody,
			WithResponse: &radarr.IndexerConfig{
				WhitelistedHardcodedSubs: "",
				ID:                       1,
				MaximumSize:              0,
				MinimumAge:               0,
				Retention:                0,
				RssSyncInterval:          60,
				AvailabilityDelay:        0,
				PreferIndexerFlags:       false,
				AllowHardcodedSubs:       false,
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "config", "indexer"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:   (*radarr.IndexerConfig)(nil),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := radarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
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
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "config", "indexer", "1"),
			ExpectedMethod: "PUT",
			ResponseStatus: 202,
			WithRequest: &radarr.IndexerConfig{
				WhitelistedHardcodedSubs: "",
				ID:                       1,
				MaximumSize:              0,
				MinimumAge:               0,
				Retention:                0,
				RssSyncInterval:          60,
				AvailabilityDelay:        0,
				PreferIndexerFlags:       false,
				AllowHardcodedSubs:       false,
			},
			ExpectedRequest: indexerRequest,
			ResponseBody:    indexerConfigBody,
			WithResponse: &radarr.IndexerConfig{
				WhitelistedHardcodedSubs: "",
				ID:                       1,
				MaximumSize:              0,
				MinimumAge:               0,
				Retention:                0,
				RssSyncInterval:          60,
				AvailabilityDelay:        0,
				PreferIndexerFlags:       false,
				AllowHardcodedSubs:       false,
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "config", "indexer", "1"),
			ExpectedMethod: "PUT",
			WithRequest: &radarr.IndexerConfig{
				ID:                 1,
				MaximumSize:        0,
				MinimumAge:         0,
				Retention:          0,
				RssSyncInterval:    60,
				AvailabilityDelay:  0,
				PreferIndexerFlags: false,
				AllowHardcodedSubs: false,
			},
			ExpectedRequest: indexerRequest,
			ResponseStatus:  404,
			ResponseBody:    `{"message": "NotFound"}`,
			WithError:       &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:    (*radarr.IndexerConfig)(nil),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := radarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.UpdateIndexerConfig(test.WithRequest.(*radarr.IndexerConfig))
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, output, test.WithResponse, "response is not the same as expected")
		})
	}
}
