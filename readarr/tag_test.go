package readarr_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"golift.io/starr"
	"golift.io/starr/readarr"
)

func TestGetTags(t *testing.T) {
	t.Parallel()

	tests := []struct {
		responseStatus   int
		name             string
		expectedPath     string
		responseBody     string
		withError        error
		expectedResponse []*starr.Tag
	}{
		{
			name:           "200",
			expectedPath:   "/api/v1/tag",
			responseStatus: 200,
			responseBody:   "[{\"label\": \"amzn\",\"id\": 1},{\"label\": \"epub\",\"id\": 2}]",
			expectedResponse: []*starr.Tag{
				{
					Label: "amzn",
					ID:    1,
				},
				{
					Label: "epub",
					ID:    2,
				},
			},
			withError: nil,
		},
		{
			name:             "404",
			expectedPath:     "/api/v1/tag",
			responseStatus:   404,
			responseBody:     `{"message": "NotFound"}`,
			withError:        starr.ErrInvalidStatusCode,
			expectedResponse: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, test.expectedPath, r.URL.String())
				w.WriteHeader(test.responseStatus)
				_, err := w.Write([]byte(test.responseBody))
				assert.NoError(t, err)
			}))
			client := readarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetTags()
			assert.ErrorIs(t, err, test.withError, "error is not the same as expected")
			assert.EqualValues(t, output, test.expectedResponse, "response is not the same as expected")
		})
	}
}

func TestGetTag(t *testing.T) {
	t.Parallel()

	tests := []struct {
		responseStatus   int
		tagID            int
		name             string
		expectedPath     string
		responseBody     string
		withError        error
		expectedResponse *starr.Tag
	}{
		{
			name:           "200",
			tagID:          1,
			expectedPath:   "/api/v1/tag/1",
			responseStatus: 200,
			responseBody:   "{\"label\": \"amzn\",\"id\": 1}",
			expectedResponse: &starr.Tag{
				Label: "amzn",
				ID:    1,
			},
			withError: nil,
		},
		{
			name:             "404",
			tagID:            1,
			expectedPath:     "/api/v1/tag/1",
			responseStatus:   404,
			responseBody:     `{"message": "NotFound"}`,
			withError:        starr.ErrInvalidStatusCode,
			expectedResponse: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, test.expectedPath, r.URL.String())
				w.WriteHeader(test.responseStatus)
				_, err := w.Write([]byte(test.responseBody))
				assert.NoError(t, err)
			}))
			client := readarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetTag(test.tagID)
			assert.ErrorIs(t, err, test.withError, "error is not the same as expected")
			assert.EqualValues(t, output, test.expectedResponse, "response is not the same as expected")
		})
	}
}

func TestAddTag(t *testing.T) {
	t.Parallel()

	tests := []struct {
		responseStatus   int
		name             string
		expectedPath     string
		responseBody     string
		withError        error
		tag              *starr.Tag
		expectedResponse *starr.Tag
	}{
		{
			name:           "200",
			expectedPath:   "/api/v1/tag",
			responseStatus: 200,
			tag: &starr.Tag{
				Label: "amzn",
			},
			responseBody: "{\"label\": \"amzn\",\"id\": 1}",
			expectedResponse: &starr.Tag{
				Label: "amzn",
				ID:    1,
			},
			withError: nil,
		},
		{
			name:             "404",
			expectedPath:     "/api/v1/tag",
			responseStatus:   404,
			responseBody:     `{"message": "NotFound"}`,
			withError:        starr.ErrInvalidStatusCode,
			expectedResponse: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, test.expectedPath, r.URL.String())
				w.WriteHeader(test.responseStatus)
				_, err := w.Write([]byte(test.responseBody))
				assert.NoError(t, err)
			}))
			client := readarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.AddTag(test.tag)
			assert.ErrorIs(t, err, test.withError, "error is not the same as expected")
			assert.EqualValues(t, output, test.expectedResponse, "response is not the same as expected")
		})
	}
}

func TestUpdateTag(t *testing.T) {
	t.Parallel()

	tests := []struct {
		responseStatus   int
		name             string
		expectedPath     string
		responseBody     string
		withError        error
		tag              *starr.Tag
		expectedResponse *starr.Tag
	}{
		{
			name:           "200",
			expectedPath:   "/api/v1/tag/1",
			responseStatus: 200,
			tag: &starr.Tag{
				Label: "amzn",
				ID:    1,
			},
			responseBody: "{\"label\": \"amzn\",\"id\": 1}",
			expectedResponse: &starr.Tag{
				Label: "amzn",
				ID:    1,
			},
			withError: nil,
		},
		{
			name:         "404",
			expectedPath: "/api/v1/tag/1",
			tag: &starr.Tag{
				Label: "amzn",
				ID:    1,
			},
			responseStatus:   404,
			responseBody:     `{"message": "NotFound"}`,
			withError:        starr.ErrInvalidStatusCode,
			expectedResponse: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, test.expectedPath, r.URL.String())
				w.WriteHeader(test.responseStatus)
				_, err := w.Write([]byte(test.responseBody))
				assert.NoError(t, err)
			}))
			client := readarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.UpdateTag(test.tag)
			assert.ErrorIs(t, err, test.withError, "error is not the same as expected")
			assert.EqualValues(t, output, test.expectedResponse, "response is not the same as expected")
		})
	}
}

func TestDeleteTag(t *testing.T) {
	t.Parallel()

	tests := []struct {
		responseStatus int
		tagID          int
		name           string
		expectedPath   string
		responseBody   string
		withError      error
	}{
		{
			name:           "200",
			tagID:          1,
			expectedPath:   "/api/v1/tag/1",
			responseStatus: 200,
			responseBody:   "{}",
			withError:      nil,
		},
		{
			name:           "404",
			tagID:          1,
			expectedPath:   "/api/v1/tag/1",
			responseStatus: 404,
			responseBody:   `{"message": "NotFound"}`,
			withError:      starr.ErrInvalidStatusCode,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, test.expectedPath, r.URL.String())
				w.WriteHeader(test.responseStatus)
				_, err := w.Write([]byte(test.responseBody))
				assert.NoError(t, err)
			}))
			client := readarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			err := client.DeleteTag(test.tagID)
			assert.ErrorIs(t, err, test.withError, "error is not the same as expected")
		})
	}
}
