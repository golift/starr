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

func TestGetTags(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "tag"),
			ExpectedMethod: "GET",
			ResponseStatus: 200,
			ResponseBody:   `[{"label": "amzn","id": 1},{"label": "netflix","id": 2}]`,
			WithResponse: []*starr.Tag{
				{
					Label: "amzn",
					ID:    1,
				},
				{
					Label: "netflix",
					ID:    2,
				},
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "tag"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:   []*starr.Tag(nil),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetTags()
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestGetTag(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "tag", "1"),
			ExpectedMethod: "GET",
			ResponseStatus: 200,
			WithRequest:    1,
			ResponseBody:   `{"label": "amzn","id": 1}`,
			WithResponse: &starr.Tag{
				Label: "amzn",
				ID:    1,
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "tag", "1"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			WithRequest:    1,
			ResponseBody:   `{"message": "NotFound"}`,
			WithResponse:   (*starr.Tag)(nil),
			WithError:      &starr.ReqError{Code: http.StatusNotFound},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetTag(test.WithRequest.(int))
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestAddTag(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "tag"),
			ExpectedMethod: "POST",
			ResponseStatus: 200,
			WithRequest: &starr.Tag{
				Label: "amzn",
			},
			ExpectedRequest: `{"label":"amzn"}` + "\n",
			ResponseBody:    `{"label": "amzn","id": 1}`,
			WithResponse: &starr.Tag{
				Label: "amzn",
				ID:    1,
			},
			WithError: nil,
		},
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "tag"),
			ExpectedMethod: "POST",
			ResponseStatus: 404,
			WithRequest: &starr.Tag{
				Label: "amzn",
			},
			ExpectedRequest: `{"label":"amzn"}` + "\n",
			ResponseBody:    `{"message": "NotFound"}`,
			WithError:       &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:    (*starr.Tag)(nil),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.AddTag(test.WithRequest.(*starr.Tag))
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestUpdateTag(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "tag", "1"),
			ExpectedMethod: "PUT",
			ResponseStatus: 200,
			WithRequest: &starr.Tag{
				ID:    1,
				Label: "amzn",
			},
			ExpectedRequest: `{"id":1,"label":"amzn"}` + "\n",
			ResponseBody:    `{"id": 1,"label": "amzn"}`,
			WithResponse: &starr.Tag{
				ID:    1,
				Label: "amzn",
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "tag", "1"),
			ExpectedMethod: "PUT",
			ResponseStatus: 404,
			WithRequest: &starr.Tag{
				ID:    1,
				Label: "amzn",
			},
			ExpectedRequest: `{"id":1,"label":"amzn"}` + "\n",
			ResponseBody:    `{"message": "NotFound"}`,
			WithError:       &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:    (*starr.Tag)(nil),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.UpdateTag(test.WithRequest.(*starr.Tag))
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestDeleteTag(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "tag", "1"),
			ExpectedMethod: "DELETE",
			WithRequest:    1,
			ResponseStatus: 200,
			ResponseBody:   "{}",
			WithError:      nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "tag", "1"),
			ExpectedMethod: "DELETE",
			WithRequest:    1,
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      &starr.ReqError{Code: http.StatusNotFound},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			err := client.DeleteTag(test.WithRequest.(int))
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
		})
	}
}
