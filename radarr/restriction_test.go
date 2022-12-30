package radarr_test

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"golift.io/starr"
	"golift.io/starr/radarr"
	"golift.io/starr/starrtest"
)

const restrictionBody = `{
	"required": "test1",
	"ignored": "test2",
	"tags": [],
	"id": 2
  }`

func TestGetRestrictions(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "restriction"),
			ExpectedMethod: "GET",
			ResponseStatus: 200,
			ResponseBody:   "[" + restrictionBody + "]",
			WithResponse: []*radarr.Restriction{
				{
					Tags:     []int{},
					Required: "test1",
					Ignored:  "test2",
					ID:       2,
				},
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "restriction"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      starr.ErrInvalidStatusCode,
			WithResponse:   []*radarr.Restriction(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := radarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetRestrictions()
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestGetRestriction(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "restriction", "2"),
			ExpectedMethod: "GET",
			ResponseStatus: 200,
			WithRequest:    int64(2),
			ResponseBody:   restrictionBody,
			WithResponse: &radarr.Restriction{
				Tags:     []int{},
				Required: "test1",
				Ignored:  "test2",
				ID:       2,
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "restriction", "2"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			WithRequest:    int64(2),
			ResponseBody:   `{"message": "NotFound"}`,
			WithResponse:   (*radarr.Restriction)(nil),
			WithError:      starr.ErrInvalidStatusCode,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := radarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetRestriction(test.WithRequest.(int64))
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestAddRestriction(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "restriction"),
			ExpectedMethod: "POST",
			ResponseStatus: 200,
			WithRequest: &radarr.Restriction{
				Required: "test1",
				Ignored:  "test2",
			},
			ExpectedRequest: `{"required":"test1","ignored":"test2"}` + "\n",
			ResponseBody:    restrictionBody,
			WithResponse: &radarr.Restriction{
				Tags:     []int{},
				Required: "test1",
				Ignored:  "test2",
				ID:       2,
			},
			WithError: nil,
		},
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "restriction"),
			ExpectedMethod: "POST",
			ResponseStatus: 404,
			WithRequest: &radarr.Restriction{
				Required: "test1",
				Ignored:  "test2",
			},
			ExpectedRequest: `{"required":"test1","ignored":"test2"}` + "\n",
			ResponseBody:    `{"message": "NotFound"}`,
			WithError:       starr.ErrInvalidStatusCode,
			WithResponse:    (*radarr.Restriction)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := radarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.AddRestriction(test.WithRequest.(*radarr.Restriction))
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestUpdateRestriction(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "restriction", "2"),
			ExpectedMethod: "PUT",
			ResponseStatus: 200,
			WithRequest: &radarr.Restriction{
				Required: "test1",
				Ignored:  "test2",
				ID:       2,
			},
			ExpectedRequest: `{"required":"test1","ignored":"test2","id":2}` + "\n",
			ResponseBody:    restrictionBody,
			WithResponse: &radarr.Restriction{
				Tags:     []int{},
				Required: "test1",
				Ignored:  "test2",
				ID:       2,
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "restriction", "2"),
			ExpectedMethod: "PUT",
			ResponseStatus: 404,
			WithRequest: &radarr.Restriction{
				Required: "test1",
				Ignored:  "test2",
				ID:       2,
			},
			ExpectedRequest: `{"required":"test1","ignored":"test2","id":2}` + "\n",
			ResponseBody:    `{"message": "NotFound"}`,
			WithError:       starr.ErrInvalidStatusCode,
			WithResponse:    (*radarr.Restriction)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := radarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.UpdateRestriction(test.WithRequest.(*radarr.Restriction))
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestDeleteRestriction(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "restriction", "1"),
			ExpectedMethod: "DELETE",
			WithRequest:    int64(1),
			ResponseStatus: 200,
			ResponseBody:   "{}",
			WithError:      nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "restriction", "1"),
			ExpectedMethod: "DELETE",
			WithRequest:    int64(1),
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
			client := radarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			err := client.DeleteRestriction(test.WithRequest.(int64))
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
		})
	}
}
