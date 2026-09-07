package lidarr_test

import (
	"net/http"
	"net/url"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golift.io/starr"
	"golift.io/starr/lidarr"
	"golift.io/starr/starrtest"
)

func TestGetTracksByAlbumRelease(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name: "200",
			ExpectedPath: path.Join("/", starr.API, lidarr.APIver, "track") + "?" + url.Values{
				"albumReleaseId": {"9"},
			}.Encode(),
			ExpectedMethod: "GET",
			WithRequest:    int64(9),
			ResponseStatus: http.StatusOK,
			ResponseBody:   `[{"id":1,"albumId":2,"title":"Track"}]`,
			WithResponse: []*lidarr.Track{{
				ID:      1,
				AlbumID: 2,
				Title:   "Track",
			}},
			WithError: nil,
		},
		{
			Name: "404",
			ExpectedPath: path.Join("/", starr.API, lidarr.APIver, "track") + "?" + url.Values{
				"albumReleaseId": {"9"},
			}.Encode(),
			ExpectedMethod: "GET",
			WithRequest:    int64(9),
			ResponseStatus: http.StatusNotFound,
			ResponseBody:   `{"message": "NotFound"}`,
			WithResponse:   []*lidarr.Track(nil),
			WithError:      &starr.ReqError{Code: http.StatusNotFound},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := lidarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetTracksByAlbumRelease(test.WithRequest.(int64))
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}
