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
			WithRequest: &radarr.BulkEdit{
				MovieIDs:    []int64{7, 3},
				Monitored:   starr.True(),
				DeleteFiles: starr.False(),
			},
			ExpectedRequest: `{"movieIds":[7,3],"monitored":true,"deleteFiles":false}` + "\n",
			ExpectedMethod:  http.MethodPut,
			WithResponse:    []*radarr.Movie{{ID: 7, Monitored: true}, {ID: 3, Monitored: true}},
		},
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "movie", "editor"),
			ResponseStatus: http.StatusOK,
			ResponseBody: `[{"id":17,"minimumAvailability":"tba","tags":[44,55,66]},` +
				`{"id":13,"minimumAvailability":"tba","tags":[44,55,66]}]`,
			WithError: nil,
			WithRequest: &radarr.BulkEdit{
				MovieIDs:            []int64{17, 13},
				Tags:                []int{44, 55, 66},
				ApplyTags:           starr.TagsAdd.Ptr(),
				MinimumAvailability: radarr.AvailabilityToBeAnnounced.Ptr(),
			},
			ExpectedRequest: `{"movieIds":[17,13],"minimumAvailability":"tba","tags":[44,55,66],"applyTags":"add"}` + "\n",
			ExpectedMethod:  http.MethodPut,
			WithResponse: []*radarr.Movie{
				{ID: 17, MinimumAvailability: radarr.AvailabilityToBeAnnounced, Tags: []int{44, 55, 66}},
				{ID: 13, MinimumAvailability: radarr.AvailabilityToBeAnnounced, Tags: []int{44, 55, 66}},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := radarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.EditMovies(test.WithRequest.(*radarr.BulkEdit))
			assert.ErrorIs(t, err, test.WithError, "the wrong error was returned")
			assert.EqualValues(t, test.WithResponse, output, "make sure ResponseBody and WithResponse are a match")
		})
	}
}

func TestDeleteMovies(t *testing.T) {
	t.Parallel()

	tests := []*starr.TestMockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "movie", "editor"),
			ResponseStatus: http.StatusOK,
			WithError:      nil,
			WithRequest: &radarr.BulkEdit{
				MovieIDs:    []int64{7, 3},
				Monitored:   starr.False(),
				DeleteFiles: starr.True(),
			},
			ExpectedRequest: `{"movieIds":[7,3],"monitored":false,"deleteFiles":true}` + "\n",
			ExpectedMethod:  http.MethodDelete,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := radarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			err := client.DeleteMovies(test.WithRequest.(*radarr.BulkEdit))
			assert.ErrorIs(t, err, test.WithError, "the wrong error was returned")
		})
	}
}
