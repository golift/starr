package sonarr_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golift.io/starr"
	"golift.io/starr/sonarr"
)

//nolint:funlen
func TestGetCommands(t *testing.T) {
	t.Parallel()

	somedate := time.Now().Add(-36 * time.Hour).Round(time.Millisecond).UTC()
	datejson, _ := somedate.MarshalJSON()

	tests := []struct {
		responseStatus   int
		name             string
		expectedPath     string
		responseBody     string
		withError        error
		expectedMethod   string
		expectedResponse []*sonarr.CommandResponse
	}{
		{
			name:           "200",
			expectedPath:   path.Join("/", starr.API, sonarr.APIver, "command"),
			responseStatus: http.StatusOK,
			responseBody: `[{"id":1234,"name":"SomeCommand","commandName":"SomeCommandName","message":` +
				`"Command Message","priority":"testalert","status":"statusalert","queued":` + string(datejson) +
				`,"started":` + string(datejson) + `,"ended":` + string(datejson) +
				`,"stateChangeTime":` + string(datejson) + `,"lastExecutionTime":` + string(datejson) +
				`,"duration":"woofun","trigger":"someTrigger","sendUpdatesToClient":true,"updateScheduledTask":true` +
				`,"body": {"mapstring": "mapinterface"}` +
				`}]`,
			withError:      nil,
			expectedMethod: "GET",
			expectedResponse: []*sonarr.CommandResponse{{
				ID:                  1234,
				Name:                "SomeCommand",
				CommandName:         "SomeCommandName",
				Message:             "Command Message",
				Priority:            "testalert",
				Status:              "statusalert",
				Queued:              somedate,
				Started:             somedate,
				Ended:               somedate,
				StateChangeTime:     somedate,
				LastExecutionTime:   somedate,
				Duration:            "woofun",
				Trigger:             "someTrigger",
				SendUpdatesToClient: true,
				UpdateScheduledTask: true,
				Body:                map[string]interface{}{"mapstring": "mapinterface"},
			}},
		},
		{
			name:             "404",
			expectedPath:     path.Join("/", starr.API, sonarr.APIver, "command"),
			responseStatus:   http.StatusNotFound,
			responseBody:     `{"message": "NotFound"}`,
			withError:        starr.ErrInvalidStatusCode,
			expectedMethod:   "GET",
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
				assert.Equal(t, r.Method, test.expectedMethod)
			}))
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetCommands()
			assert.ErrorIs(t, err, test.withError, "error is not the same as expected")
			assert.EqualValues(t, test.expectedResponse, output, "response is not the same as expected")
		})
	}
}

//nolint:funlen
func TestSendCommand(t *testing.T) {
	t.Parallel()

	somedate := time.Now().Add(-36 * time.Hour).Round(time.Millisecond).UTC()
	datejson, _ := somedate.MarshalJSON()

	tests := []struct {
		responseStatus   int
		name             string
		expectedPath     string
		responseBody     string
		withError        error
		withRequest      *sonarr.CommandRequest
		expectedRequest  string
		expectedMethod   string
		expectedResponse *sonarr.CommandResponse
	}{
		{
			name:           "200",
			expectedPath:   path.Join("/", starr.API, sonarr.APIver, "command"),
			responseStatus: http.StatusOK,
			responseBody: `{"id":1234,"name":"SomeCommand","commandName":"SomeCommandName","message":` +
				`"Command Message","priority":"testalert","status":"statusalert","queued":` + string(datejson) +
				`,"started":` + string(datejson) + `,"ended":` + string(datejson) +
				`,"stateChangeTime":` + string(datejson) + `,"lastExecutionTime":` + string(datejson) +
				`,"duration":"woofun","trigger":"someTrigger","sendUpdatesToClient":true,"updateScheduledTask":true` +
				`,"body": {"mapstring": "mapinterface"}` +
				`}`,
			withError: nil,
			withRequest: &sonarr.CommandRequest{
				Name:      "SomeCommand",
				SeriesIDs: []int64{1, 3, 7},
			},
			expectedRequest: `{"name":"SomeCommand","seriesIds":[1,3,7]}` + "\n",
			expectedMethod:  "POST",
			expectedResponse: &sonarr.CommandResponse{
				ID:                  1234,
				Name:                "SomeCommand",
				CommandName:         "SomeCommandName",
				Message:             "Command Message",
				Priority:            "testalert",
				Status:              "statusalert",
				Queued:              somedate,
				Started:             somedate,
				Ended:               somedate,
				StateChangeTime:     somedate,
				LastExecutionTime:   somedate,
				Duration:            "woofun",
				Trigger:             "someTrigger",
				SendUpdatesToClient: true,
				UpdateScheduledTask: true,
				Body:                map[string]interface{}{"mapstring": "mapinterface"},
			},
		},
		{
			name:             "404",
			expectedPath:     path.Join("/", starr.API, sonarr.APIver, "command"),
			responseStatus:   http.StatusNotFound,
			responseBody:     `{"message": "NotFound"}`,
			withError:        starr.ErrInvalidStatusCode,
			expectedMethod:   "POST",
			expectedResponse: nil,
			withRequest:      &sonarr.CommandRequest{Name: "Something"},
			expectedRequest:  `{"name":"Something"}` + "\n",
		},
		{
			name:             "noname", // no name provided? returns empty (non-nil) response.
			withRequest:      &sonarr.CommandRequest{Name: ""},
			expectedResponse: &sonarr.CommandResponse{},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				assert.Equal(t, test.expectedPath, req.URL.String())
				w.WriteHeader(test.responseStatus)

				_, err := w.Write([]byte(test.responseBody))
				assert.NoError(t, err)
				assert.Equal(t, req.Method, test.expectedMethod)

				body, err := ioutil.ReadAll(req.Body)
				assert.NoError(t, err)
				assert.Equal(t, test.expectedRequest, string(body))
			}))

			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.SendCommand(test.withRequest)
			assert.ErrorIs(t, err, test.withError, "error is not the same as expected")
			assert.EqualValues(t, test.expectedResponse, output, "response is not the same as expected")
		})
	}
}

//nolint:funlen
func TestGetCommandStatus(t *testing.T) {
	t.Parallel()

	somedate := time.Now().Add(-36 * time.Hour).Round(time.Millisecond).UTC()
	datejson, _ := somedate.MarshalJSON()

	tests := []struct {
		responseStatus   int
		name             string
		expectedPath     string
		responseBody     string
		withError        error
		expectedMethod   string
		withCommandID    int64
		expectedResponse *sonarr.CommandResponse
	}{
		{
			name:           "200",
			expectedPath:   path.Join("/", starr.API, sonarr.APIver, "command", "146"),
			responseStatus: http.StatusOK,
			responseBody: `{"id":1234,"name":"SomeCommand","commandName":"SomeCommandName","message":` +
				`"Command Message","priority":"testalert","status":"statusalert","queued":` + string(datejson) +
				`,"started":` + string(datejson) + `,"ended":` + string(datejson) +
				`,"stateChangeTime":` + string(datejson) + `,"lastExecutionTime":` + string(datejson) +
				`,"duration":"woofun","trigger":"someTrigger","sendUpdatesToClient":true,"updateScheduledTask":true` +
				`,"body": {"mapstring": "mapinterface"}}`,
			withError:      nil,
			expectedMethod: "GET",
			withCommandID:  146,
			expectedResponse: &sonarr.CommandResponse{
				ID:                  1234,
				Name:                "SomeCommand",
				CommandName:         "SomeCommandName",
				Message:             "Command Message",
				Priority:            "testalert",
				Status:              "statusalert",
				Queued:              somedate,
				Started:             somedate,
				Ended:               somedate,
				StateChangeTime:     somedate,
				LastExecutionTime:   somedate,
				Duration:            "woofun",
				Trigger:             "someTrigger",
				SendUpdatesToClient: true,
				UpdateScheduledTask: true,
				Body:                map[string]interface{}{"mapstring": "mapinterface"},
			},
		},
		{
			name:             "404",
			expectedPath:     path.Join("/", starr.API, sonarr.APIver, "command", "123"),
			responseStatus:   http.StatusNotFound,
			withCommandID:    123,
			responseBody:     `{"message": "NotFound"}`,
			withError:        starr.ErrInvalidStatusCode,
			expectedMethod:   "GET",
			expectedResponse: nil,
		},
		{
			name:             "command0", // command zero returns empty (non-nil) response.
			withCommandID:    0,
			withError:        nil,
			expectedMethod:   "GET",
			expectedResponse: &sonarr.CommandResponse{},
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
				assert.Equal(t, r.Method, test.expectedMethod)
			}))
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetCommandStatus(test.withCommandID)
			assert.ErrorIs(t, err, test.withError, "error is not the same as expected")
			assert.EqualValues(t, test.expectedResponse, output, "response is not the same as expected")
		})
	}
}
