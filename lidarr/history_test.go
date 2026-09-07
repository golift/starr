package lidarr_test

import (
	"net/http"
	"path"
	"testing"

	"github.com/stretchr/testify/require"
	"golift.io/starr"
	"golift.io/starr/lidarr"
	"golift.io/starr/starrtest"
)

func TestFail(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:            "200",
			ExpectedPath:    path.Join("/", starr.API, lidarr.APIver, "history", "failed", "102"),
			ExpectedMethod:  "POST",
			ExpectedRequest: "",
			WithRequest:     int64(102),
			ResponseStatus:  http.StatusOK,
			ResponseBody:    "{}",
			WithError:       nil,
		},
		{
			Name:            "404",
			ExpectedPath:    path.Join("/", starr.API, lidarr.APIver, "history", "failed", "102"),
			ExpectedMethod:  "POST",
			ExpectedRequest: "",
			WithRequest:     int64(102),
			ResponseStatus:  http.StatusNotFound,
			ResponseBody:    `{"message": "NotFound"}`,
			WithError:       &starr.ReqError{Code: http.StatusNotFound},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := lidarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			err := client.Fail(test.WithRequest.(int64))
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
		})
	}

	t.Run("invalidID", func(t *testing.T) {
		t.Parallel()
		client := lidarr.New(starr.New("mockAPIkey", "http://127.0.0.1", 0))
		err := client.Fail(0)
		require.ErrorIs(t, err, starr.ErrRequestError, "invalid history ID should not hit the API")
	})
}
