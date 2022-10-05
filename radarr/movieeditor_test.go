package radarr_test

import (
	"net/http"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"golift.io/starr"
	"golift.io/starr/radarr"
)

func TestEditMovies(t *testing.T) {
	t.Parallel()

	tests := []*starr.TestMockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "movie", "editor"),
			ResponseStatus: http.StatusOK,
			ResponseBody:   `[{"id": 7, "monitored": true},{"id": 3, "monitored": true}]`,
			WithError:      nil,
			WithRequest: &radarr.EditMovies{
				MovieIDs:    []int64{7, 3},
				Monitored:   starr.True(),
				DeleteFiles: starr.False(),
			},
			ExpectedRequest: `{"movieIds":[7,3],"monitored":true,"deleteFiles":false}` + "\n",
			ExpectedMethod:  http.MethodPut,
			WithResponse:    []*radarr.Movie{{ID: 7, Monitored: true}, {ID: 3, Monitored: true}},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := radarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.EditMovies(test.WithRequest.(*radarr.EditMovies))
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}
