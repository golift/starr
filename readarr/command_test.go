package readarr_test

import (
	"net/http"
	"path"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golift.io/starr"
	"golift.io/starr/readarr"
	"golift.io/starr/starrtest"
)

func TestGetCommands(t *testing.T) {
	t.Parallel()

	somedate := time.Now().Add(-36 * time.Hour).Round(time.Millisecond).UTC()
	datejson, _ := somedate.MarshalJSON()
	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, readarr.APIver, "command"),
			ResponseStatus: http.StatusOK,
			ResponseBody: `[{"id":1234,"name":"SomeCommand","commandName":"SomeCommandName","message":` +
				`"Command Message","priority":"testalert","status":"statusalert","queued":` + string(datejson) +
				`,"started":` + string(datejson) + `,"ended":` + string(datejson) +
				`,"stateChangeTime":` + string(datejson) + `,"lastExecutionTime":` + string(datejson) +
				`,"duration":"woofun","trigger":"someTrigger","sendUpdatesToClient":true,"updateScheduledTask":true` +
				`,"body": {"mapstring": "mapinterface"}` +
				`}]`,
			WithError:      nil,
			ExpectedMethod: "GET",
			WithResponse: []*readarr.CommandResponse{{
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
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, readarr.APIver, "command"),
			ResponseStatus: http.StatusNotFound,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      starr.ErrInvalidStatusCode,
			ExpectedMethod: "GET",
			WithResponse:   []*readarr.CommandResponse(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := readarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetCommands()
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestSendCommand(t *testing.T) {
	t.Parallel()

	somedate := time.Now().Add(-36 * time.Hour).Round(time.Millisecond).UTC()
	datejson, _ := somedate.MarshalJSON()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, readarr.APIver, "command"),
			ResponseStatus: http.StatusOK,
			ResponseBody: `{"id":1234,"name":"SomeCommand","commandName":"SomeCommandName","message":` +
				`"Command Message","priority":"testalert","status":"statusalert","queued":` + string(datejson) +
				`,"started":` + string(datejson) + `,"ended":` + string(datejson) +
				`,"stateChangeTime":` + string(datejson) + `,"lastExecutionTime":` + string(datejson) +
				`,"duration":"woofun","trigger":"someTrigger","sendUpdatesToClient":true,"updateScheduledTask":true` +
				`,"body": {"mapstring": "mapinterface"}` +
				`}`,
			WithError: nil,
			WithRequest: &readarr.CommandRequest{
				Name:    "SomeCommand",
				BookIDs: []int64{1, 3, 7},
			},
			ExpectedRequest: `{"name":"SomeCommand","bookIds":[1,3,7]}` + "\n",
			ExpectedMethod:  "POST",
			WithResponse: &readarr.CommandResponse{
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
			Name:            "404",
			ExpectedPath:    path.Join("/", starr.API, readarr.APIver, "command"),
			ResponseStatus:  http.StatusNotFound,
			ResponseBody:    `{"message": "NotFound"}`,
			WithError:       starr.ErrInvalidStatusCode,
			ExpectedMethod:  "POST",
			WithResponse:    (*readarr.CommandResponse)(nil),
			WithRequest:     &readarr.CommandRequest{Name: "Something"},
			ExpectedRequest: `{"name":"Something"}` + "\n",
		},
		{
			Name:         "noname", // no name provided? returns empty (non-nil) response.
			WithRequest:  &readarr.CommandRequest{Name: ""},
			WithResponse: &readarr.CommandResponse{},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := readarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.SendCommand(test.WithRequest.(*readarr.CommandRequest))
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}
