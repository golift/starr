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

const (
	firstDelayProfile = `{
		"enableUsenet": true,
		"enableTorrent": true,
		"preferredProtocol": "usenet",
		"usenetDelay": 0,
		"torrentDelay": 0,
		"bypassIfHighestQuality": true,
		"order": 2147483647,
		"tags": [],
		"id": 1
	}`
	secondDelayProfile = `{
		"enableUsenet": false,
		"enableTorrent": true,
		"preferredProtocol": "torrent",
		"usenetDelay": 0,
		"torrentDelay": 0,
		"bypassIfHighestQuality": false,
		"order": 1,
		"tags": [11],
		"id": 10
	}`
	delayProfileRequest = `{"enableTorrent":true,"order":1,"tags":[11],"preferredProtocol":"torrent"}` + "\n"
)

func TestGetDelayProfiles(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "delayProfile"),
			ExpectedMethod: "GET",
			ResponseStatus: 200,
			ResponseBody:   `[` + firstDelayProfile + `,` + secondDelayProfile + `]`,
			WithResponse: []*sonarr.DelayProfile{
				{
					EnableUsenet:           true,
					EnableTorrent:          true,
					PreferredProtocol:      "usenet",
					UsenetDelay:            0,
					TorrentDelay:           0,
					BypassIfHighestQuality: true,
					Order:                  2147483647,
					Tags:                   []int{},
					ID:                     1,
				},
				{
					EnableUsenet:           false,
					EnableTorrent:          true,
					PreferredProtocol:      "torrent",
					UsenetDelay:            0,
					TorrentDelay:           0,
					BypassIfHighestQuality: false,
					Order:                  1,
					Tags:                   []int{11},
					ID:                     10,
				},
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "delayProfile"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:   []*sonarr.DelayProfile(nil),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetDelayProfiles()
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestGetDelayProfile(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "delayProfile/1"),
			ExpectedMethod: "GET",
			ResponseStatus: 200,
			WithRequest:    int64(1),
			ResponseBody:   firstDelayProfile,
			WithResponse: &sonarr.DelayProfile{
				EnableUsenet:           true,
				EnableTorrent:          true,
				PreferredProtocol:      "usenet",
				UsenetDelay:            0,
				TorrentDelay:           0,
				BypassIfHighestQuality: true,
				Order:                  2147483647,
				Tags:                   []int{},
				ID:                     1,
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "delayProfile", "1"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			WithRequest:    int64(1),
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:   (*sonarr.DelayProfile)(nil),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetDelayProfile(test.WithRequest.(int64))
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestAddDelayProfile(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "delayProfile"),
			ExpectedMethod: "POST",
			ResponseStatus: 200,
			WithRequest: &sonarr.DelayProfile{
				EnableUsenet:           false,
				EnableTorrent:          true,
				PreferredProtocol:      "torrent",
				UsenetDelay:            0,
				TorrentDelay:           0,
				BypassIfHighestQuality: false,
				Order:                  1,
				Tags:                   []int{11},
			},
			ExpectedRequest: delayProfileRequest,
			ResponseBody:    secondDelayProfile,
			WithResponse: &sonarr.DelayProfile{
				EnableUsenet:           false,
				EnableTorrent:          true,
				PreferredProtocol:      "torrent",
				UsenetDelay:            0,
				TorrentDelay:           0,
				BypassIfHighestQuality: false,
				Order:                  1,
				Tags:                   []int{11},
				ID:                     10,
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "delayProfile"),
			ExpectedMethod: "POST",
			WithRequest: &sonarr.DelayProfile{
				EnableUsenet:           false,
				EnableTorrent:          true,
				PreferredProtocol:      "torrent",
				UsenetDelay:            0,
				TorrentDelay:           0,
				BypassIfHighestQuality: false,
				Order:                  1,
				Tags:                   []int{11},
			},
			ExpectedRequest: delayProfileRequest,
			ResponseStatus:  404,
			ResponseBody:    `{"message": "NotFound"}`,
			WithError:       &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:    (*sonarr.DelayProfile)(nil),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.AddDelayProfile(test.WithRequest.(*sonarr.DelayProfile))
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestUpdateDelayProfile(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "delayProfile", "10"),
			ExpectedMethod: "PUT",
			ResponseStatus: 200,
			WithRequest: &sonarr.DelayProfile{
				EnableTorrent: true,
				ID:            10,
				Tags:          []int{11},
			},
			ExpectedRequest: `{"enableTorrent":true,"id":10,"tags":[11]}` + "\n",
			ResponseBody:    secondDelayProfile,
			WithResponse: &sonarr.DelayProfile{
				EnableUsenet:           false,
				EnableTorrent:          true,
				PreferredProtocol:      "torrent",
				UsenetDelay:            0,
				TorrentDelay:           0,
				BypassIfHighestQuality: false,
				Order:                  1,
				Tags:                   []int{11},
				ID:                     10,
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "delayProfile", "10"),
			ExpectedMethod: "PUT",
			WithRequest: &sonarr.DelayProfile{
				EnableTorrent: true,
				ID:            10,
				Tags:          []int{11},
			},
			ExpectedRequest: `{"enableTorrent":true,"id":10,"tags":[11]}` + "\n",
			ResponseStatus:  404,
			ResponseBody:    `{"message": "NotFound"}`,
			WithError:       &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:    (*sonarr.DelayProfile)(nil),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.UpdateDelayProfile(test.WithRequest.(*sonarr.DelayProfile))
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestDeleteDelayProfile(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "delayProfile", "10"),
			ExpectedMethod: "DELETE",
			ResponseStatus: 200,
			WithRequest:    int64(10),
			ResponseBody:   "{}",
			WithError:      nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "delayProfile", "10"),
			ExpectedMethod: "DELETE",
			WithRequest:    int64(10),
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:   (*sonarr.DelayProfile)(nil),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			err := client.DeleteDelayProfile(test.WithRequest.(int64))
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
		})
	}
}
