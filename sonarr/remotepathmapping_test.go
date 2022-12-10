package sonarr_test

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"golift.io/starr"
	"golift.io/starr/sonarr"
	"golift.io/starr/starrtest"
)

const (
	remotePathMapping = `{
		"host": "transmission",
		"remotePath": "/remote/",
		"localPath": "/local/",
		"id": 2
	}`
)

func TestGetRemotePathMappings(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "remotePathMapping"),
			ExpectedMethod: "GET",
			ResponseStatus: 200,
			ResponseBody:   `[` + remotePathMapping + `]`,
			WithResponse: []*starr.RemotePathMapping{
				{
					Host:       "transmission",
					RemotePath: "/remote/",
					LocalPath:  "/local/",
					ID:         2,
				},
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "remotePathMapping"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      starr.ErrInvalidStatusCode,
			WithResponse:   []*starr.RemotePathMapping(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetRemotePathMappings()
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestGetRemotePathMapping(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "remotePathMapping", "1"),
			ExpectedMethod: "GET",
			ResponseStatus: 200,
			WithRequest:    int64(1),
			ResponseBody:   remotePathMapping,
			WithResponse: &starr.RemotePathMapping{
				Host:       "transmission",
				RemotePath: "/remote/",
				LocalPath:  "/local/",
				ID:         2,
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "remotePathMapping", "1"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			WithRequest:    int64(1),
			ResponseBody:   `{"message": "NotFound"}`,
			WithResponse:   (*starr.RemotePathMapping)(nil),
			WithError:      starr.ErrInvalidStatusCode,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetRemotePathMapping(test.WithRequest.(int64))
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestAddRemotePathMapping(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "201",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "remotePathMapping"),
			ExpectedMethod: "POST",
			ResponseStatus: 201,
			WithRequest: &starr.RemotePathMapping{
				Host:       "transmission",
				RemotePath: "/remote/",
				LocalPath:  "/local/",
			},
			ExpectedRequest: `{"host":"transmission","remotePath":"/remote/","localPath":"/local/"}` + "\n",
			ResponseBody:    remotePathMapping,
			WithResponse: &starr.RemotePathMapping{
				Host:       "transmission",
				RemotePath: "/remote/",
				LocalPath:  "/local/",
				ID:         2,
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "remotePathMapping"),
			ExpectedMethod: "POST",
			ResponseStatus: 404,
			WithRequest: &starr.RemotePathMapping{
				Host:       "transmission",
				RemotePath: "/remote/",
				LocalPath:  "/local/",
			},
			ExpectedRequest: `{"host":"transmission","remotePath":"/remote/","localPath":"/local/"}` + "\n",
			ResponseBody:    `{"message": "NotFound"}`,
			WithError:       starr.ErrInvalidStatusCode,
			WithResponse:    (*starr.RemotePathMapping)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.AddRemotePathMapping(test.WithRequest.(*starr.RemotePathMapping))
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestUpdateRemotePathMapping(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "201",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "remotePathMapping", "2"),
			ExpectedMethod: "PUT",
			ResponseStatus: 201,
			WithRequest: &starr.RemotePathMapping{
				Host:       "transmission",
				RemotePath: "/remote/",
				LocalPath:  "/local/",
				ID:         2,
			},
			ExpectedRequest: `{"id":2,"host":"transmission","remotePath":"/remote/","localPath":"/local/"}` + "\n",
			ResponseBody:    remotePathMapping,
			WithResponse: &starr.RemotePathMapping{
				Host:       "transmission",
				RemotePath: "/remote/",
				LocalPath:  "/local/",
				ID:         2,
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "remotePathMapping", "2"),
			ExpectedMethod: "PUT",
			ResponseStatus: 404,
			WithRequest: &starr.RemotePathMapping{
				Host:       "transmission",
				RemotePath: "/remote/",
				LocalPath:  "/local/",
				ID:         2,
			},
			ExpectedRequest: `{"id":2,"host":"transmission","remotePath":"/remote/","localPath":"/local/"}` + "\n",
			ResponseBody:    `{"message": "NotFound"}`,
			WithError:       starr.ErrInvalidStatusCode,
			WithResponse:    (*starr.RemotePathMapping)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.UpdateRemotePathMapping(test.WithRequest.(*starr.RemotePathMapping))
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestDeleteRemotePathMapping(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "remotePathMapping", "2"),
			ExpectedMethod: "DELETE",
			WithRequest:    int64(2),
			ResponseStatus: 200,
			ResponseBody:   "{}",
			WithError:      nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "remotePathMapping", "2"),
			ExpectedMethod: "DELETE",
			WithRequest:    int64(2),
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      starr.ErrInvalidStatusCode,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			err := client.DeleteRemotePathMapping(test.WithRequest.(int64))
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
		})
	}
}
