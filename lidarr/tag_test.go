package lidarr_test

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"golift.io/starr"
	"golift.io/starr/lidarr"
)

func TestGetTags(t *testing.T) {
	t.Parallel()

	tests := []*starr.TestMockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, lidarr.APIver, "tag"),
			ExpectedMethod: "GET",
			ResponseStatus: 200,
			ResponseBody:   `[{"label": "flac","id": 1},{"label": "mp3","id": 2}]`,
			WithResponse: []*starr.Tag{
				{
					Label: "flac",
					ID:    1,
				},
				{
					Label: "mp3",
					ID:    2,
				},
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, lidarr.APIver, "tag"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      starr.ErrInvalidStatusCode,
			WithResponse:   []*starr.Tag(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := lidarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetTags()
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestGetTag(t *testing.T) {
	t.Parallel()

	tests := []*starr.TestMockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, lidarr.APIver, "tag/1"),
			ExpectedMethod: "GET",
			ResponseStatus: 200,
			WithRequest:    1,
			ResponseBody:   `{"label": "flac","id": 1}`,
			WithResponse: &starr.Tag{
				Label: "flac",
				ID:    1,
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, lidarr.APIver, "tag/1"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			WithRequest:    1,
			ResponseBody:   `{"message": "NotFound"}`,
			WithResponse:   (*starr.Tag)(nil),
			WithError:      starr.ErrInvalidStatusCode,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := lidarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetTag(test.WithRequest.(int))
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestAddTag(t *testing.T) {
	t.Parallel()

	tests := []*starr.TestMockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, lidarr.APIver, "tag"),
			ExpectedMethod: "POST",
			ResponseStatus: 200,
			WithRequest: &starr.Tag{
				Label: "flac",
			},
			ExpectedRequest: `{"label":"flac"}` + "\n",
			ResponseBody:    `{"label": "flac","id": 1}`,
			WithResponse: &starr.Tag{
				Label: "flac",
				ID:    1,
			},
			WithError: nil,
		},
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, lidarr.APIver, "tag"),
			ExpectedMethod: "POST",
			ResponseStatus: 404,
			WithRequest: &starr.Tag{
				Label: "flac",
			},
			ExpectedRequest: `{"label":"flac"}` + "\n",
			ResponseBody:    `{"message": "NotFound"}`,
			WithError:       starr.ErrInvalidStatusCode,
			WithResponse:    (*starr.Tag)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := lidarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.AddTag(test.WithRequest.(*starr.Tag))
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestUpdateTag(t *testing.T) {
	t.Parallel()

	tests := []*starr.TestMockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, lidarr.APIver, "tag/1"),
			ExpectedMethod: "PUT",
			ResponseStatus: 200,
			WithRequest: &starr.Tag{
				ID:    1,
				Label: "flac",
			},
			ExpectedRequest: `{"id":1,"label":"flac"}` + "\n",
			ResponseBody:    `{"id": 1,"label": "flac"}`,
			WithResponse: &starr.Tag{
				ID:    1,
				Label: "flac",
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, lidarr.APIver, "tag/1"),
			ExpectedMethod: "PUT",
			ResponseStatus: 404,
			WithRequest: &starr.Tag{
				ID:    1,
				Label: "flac",
			},
			ExpectedRequest: `{"id":1,"label":"flac"}` + "\n",
			ResponseBody:    `{"message": "NotFound"}`,
			WithError:       starr.ErrInvalidStatusCode,
			WithResponse:    (*starr.Tag)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := lidarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.UpdateTag(test.WithRequest.(*starr.Tag))
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestDeleteTag(t *testing.T) {
	t.Parallel()

	tests := []*starr.TestMockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, lidarr.APIver, "tag/1"),
			ExpectedMethod: "DELETE",
			WithRequest:    1,
			ResponseStatus: 200,
			ResponseBody:   "{}",
			WithError:      nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, lidarr.APIver, "tag/1"),
			ExpectedMethod: "DELETE",
			WithRequest:    1,
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
			client := lidarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			err := client.DeleteTag(test.WithRequest.(int))
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
		})
	}
}
