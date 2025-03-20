package radarr_test

import (
	"net/http"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golift.io/starr"
	"golift.io/starr/radarr"
	"golift.io/starr/starrtest"
)

const notificationResponseBody = `{
    "onGrab": false,
    "onDownload": true,
    "onUpgrade": false,
    "onRename": false,
    "onMovieAdded": false,
    "onMovieDelete": false,
    "onMovieFileDelete": false,
    "onMovieFileDeleteForUpgrade": false,
    "onHealthIssue": false,
    "onApplicationUpdate": false,
    "supportsOnGrab": true,
    "supportsOnDownload": true,
    "supportsOnUpgrade": true,
    "supportsOnRename": true,
    "supportsOnMovieAdded": true,
    "supportsOnMovieDelete": true,
    "supportsOnMovieFileDelete": true,
    "supportsOnMovieFileDeleteForUpgrade": true,
    "supportsOnHealthIssue": true,
    "supportsOnApplicationUpdate": true,
    "includeHealthWarnings": false,
	"name": "Test",
	"fields": [
	  {
		"order": 0,
		"name": "path",
		"label": "Path",
		"value": "/scripts/radarr.sh",
		"type": "filePath",
		"advanced": false
	  },
	  {
		"order": 1,
		"name": "arguments",
		"label": "Arguments",
		"helpText": "Arguments to pass to the script",
		"type": "textbox",
		"advanced": false,
		"hidden": "hiddenIfNotSet"
	  }
	],
	"implementationName": "Custom Script",
	"implementation": "CustomScript",
	"configContract": "CustomScriptSettings",
	"infoLink": "https://wiki.servarr.com/radarr/supported#customscript",
	"message": {
	  "message": "Testing will execute the script with the EventType set to Test",
	  "type": "warning"
	},
	"tags": [],
	"id": 3
  }`

const addNotification = `{"onDownload":true,"name":"Test","implementation":"CustomScript","configContract":` +
	`"CustomScriptSettings","fields":[{"name":"path","value":"/scripts/radarr.sh"}]}`

const updateNotification = `{"onDownload":true,"id":3,"name":"Test","implementation":"CustomScript","configContract":` +
	`"CustomScriptSettings","fields":[{"name":"path","value":"/scripts/radarr.sh"}]}`

func TestGetNotifications(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:            "200",
			ExpectedPath:    path.Join("/", starr.API, radarr.APIver, "notification"),
			ExpectedRequest: "",
			ExpectedMethod:  "GET",
			ResponseStatus:  200,
			ResponseBody:    "[" + notificationResponseBody + "]",
			WithRequest:     nil,
			WithResponse: []*radarr.NotificationOutput{
				{
					OnDownload:                          true,
					SupportsOnGrab:                      true,
					SupportsOnDownload:                  true,
					SupportsOnUpgrade:                   true,
					SupportsOnRename:                    true,
					SupportsOnMovieAdded:                true,
					SupportsOnMovieDelete:               true,
					SupportsOnMovieFileDelete:           true,
					SupportsOnMovieFileDeleteForUpgrade: true,
					SupportsOnHealthIssue:               true,
					SupportsOnApplicationUpdate:         true,
					ID:                                  3,
					Name:                                "Test",
					ImplementationName:                  "Custom Script",
					Implementation:                      "CustomScript",
					ConfigContract:                      "CustomScriptSettings",
					InfoLink:                            "https://wiki.servarr.com/radarr/supported#customscript",
					Tags:                                []int{},
					Fields: []*starr.FieldOutput{
						{
							Order:    0,
							Name:     "path",
							Label:    "Path",
							Value:    "/scripts/radarr.sh",
							Type:     "filePath",
							Advanced: false,
						},
						{
							Order:    1,
							Name:     "arguments",
							Label:    "Arguments",
							HelpText: "Arguments to pass to the script",
							Hidden:   "hiddenIfNotSet",
							Type:     "textbox",
							Advanced: false,
						},
					},
				},
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "notification"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:   ([]*radarr.NotificationOutput)(nil),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := radarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetNotifications()
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestGetNotification(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:            "200",
			ExpectedPath:    path.Join("/", starr.API, radarr.APIver, "notification", "1"),
			ExpectedRequest: "",
			ExpectedMethod:  "GET",
			ResponseStatus:  200,
			ResponseBody:    notificationResponseBody,
			WithRequest:     nil,
			WithResponse: &radarr.NotificationOutput{
				OnDownload:                          true,
				SupportsOnGrab:                      true,
				SupportsOnDownload:                  true,
				SupportsOnUpgrade:                   true,
				SupportsOnRename:                    true,
				SupportsOnMovieAdded:                true,
				SupportsOnMovieDelete:               true,
				SupportsOnMovieFileDelete:           true,
				SupportsOnMovieFileDeleteForUpgrade: true,
				SupportsOnHealthIssue:               true,
				SupportsOnApplicationUpdate:         true,
				ID:                                  3,
				Name:                                "Test",
				ImplementationName:                  "Custom Script",
				Implementation:                      "CustomScript",
				ConfigContract:                      "CustomScriptSettings",
				InfoLink:                            "https://wiki.servarr.com/radarr/supported#customscript",
				Tags:                                []int{},
				Fields: []*starr.FieldOutput{
					{
						Order:    0,
						Name:     "path",
						Label:    "Path",
						Value:    "/scripts/radarr.sh",
						Type:     "filePath",
						Advanced: false,
					},
					{
						Order:    1,
						Name:     "arguments",
						Label:    "Arguments",
						HelpText: "Arguments to pass to the script",
						Hidden:   "hiddenIfNotSet",
						Type:     "textbox",
						Advanced: false,
					},
				},
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "notification", "1"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:   (*radarr.NotificationOutput)(nil),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := radarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetNotification(1)
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestAddNotification(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "notification"),
			ExpectedMethod: "POST",
			ResponseStatus: 200,
			WithRequest: &radarr.NotificationInput{
				OnDownload:     true,
				Name:           "Test",
				Implementation: "CustomScript",
				ConfigContract: "CustomScriptSettings",
				Fields: []*starr.FieldInput{
					{
						Name:  "path",
						Value: "/scripts/radarr.sh",
					},
				},
			},
			ExpectedRequest: addNotification + "\n",
			ResponseBody:    notificationResponseBody,
			WithResponse: &radarr.NotificationOutput{
				OnDownload:                          true,
				SupportsOnGrab:                      true,
				SupportsOnDownload:                  true,
				SupportsOnUpgrade:                   true,
				SupportsOnRename:                    true,
				SupportsOnMovieAdded:                true,
				SupportsOnMovieDelete:               true,
				SupportsOnMovieFileDelete:           true,
				SupportsOnMovieFileDeleteForUpgrade: true,
				SupportsOnHealthIssue:               true,
				SupportsOnApplicationUpdate:         true,
				ID:                                  3,
				Name:                                "Test",
				ImplementationName:                  "Custom Script",
				Implementation:                      "CustomScript",
				ConfigContract:                      "CustomScriptSettings",
				InfoLink:                            "https://wiki.servarr.com/radarr/supported#customscript",
				Tags:                                []int{},
				Fields: []*starr.FieldOutput{
					{
						Order:    0,
						Name:     "path",
						Label:    "Path",
						Value:    "/scripts/radarr.sh",
						Type:     "filePath",
						Advanced: false,
					},
					{
						Order:    1,
						Name:     "arguments",
						Label:    "Arguments",
						HelpText: "Arguments to pass to the script",
						Hidden:   "hiddenIfNotSet",
						Type:     "textbox",
						Advanced: false,
					},
				},
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "notification"),
			ExpectedMethod: "POST",
			ResponseStatus: 404,
			WithRequest: &radarr.NotificationInput{
				OnDownload:     true,
				Name:           "Test",
				Implementation: "CustomScript",
				ConfigContract: "CustomScriptSettings",
				Fields: []*starr.FieldInput{
					{
						Name:  "path",
						Value: "/scripts/radarr.sh",
					},
				},
			},
			ExpectedRequest: addNotification + "\n",
			ResponseBody:    `{"message": "NotFound"}`,
			WithError:       &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:    (*radarr.NotificationOutput)(nil),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := radarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.AddNotification(test.WithRequest.(*radarr.NotificationInput))
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestUpdateNotification(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "notification", "3"),
			ExpectedMethod: "PUT",
			ResponseStatus: 200,
			WithRequest: &radarr.NotificationInput{
				OnDownload:     true,
				ID:             3,
				Name:           "Test",
				Implementation: "CustomScript",
				ConfigContract: "CustomScriptSettings",
				Fields: []*starr.FieldInput{
					{
						Name:  "path",
						Value: "/scripts/radarr.sh",
					},
				},
			},
			ExpectedRequest: updateNotification + "\n",
			ResponseBody:    notificationResponseBody,
			WithResponse: &radarr.NotificationOutput{
				OnDownload:                          true,
				SupportsOnGrab:                      true,
				SupportsOnDownload:                  true,
				SupportsOnUpgrade:                   true,
				SupportsOnRename:                    true,
				SupportsOnMovieAdded:                true,
				SupportsOnMovieDelete:               true,
				SupportsOnMovieFileDelete:           true,
				SupportsOnMovieFileDeleteForUpgrade: true,
				SupportsOnHealthIssue:               true,
				SupportsOnApplicationUpdate:         true,
				ID:                                  3,
				Name:                                "Test",
				ImplementationName:                  "Custom Script",
				Implementation:                      "CustomScript",
				ConfigContract:                      "CustomScriptSettings",
				InfoLink:                            "https://wiki.servarr.com/radarr/supported#customscript",
				Tags:                                []int{},
				Fields: []*starr.FieldOutput{
					{
						Order:    0,
						Name:     "path",
						Label:    "Path",
						Value:    "/scripts/radarr.sh",
						Type:     "filePath",
						Advanced: false,
					},
					{
						Order:    1,
						Name:     "arguments",
						Label:    "Arguments",
						HelpText: "Arguments to pass to the script",
						Hidden:   "hiddenIfNotSet",
						Type:     "textbox",
						Advanced: false,
					},
				},
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "notification", "3"),
			ExpectedMethod: "PUT",
			ResponseStatus: 404,
			WithRequest: &radarr.NotificationInput{
				OnDownload:     true,
				ID:             3,
				Name:           "Test",
				Implementation: "CustomScript",
				ConfigContract: "CustomScriptSettings",
				Fields: []*starr.FieldInput{
					{
						Name:  "path",
						Value: "/scripts/radarr.sh",
					},
				},
			},
			ExpectedRequest: updateNotification + "\n",
			ResponseBody:    `{"message": "NotFound"}`,
			WithError:       &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:    (*radarr.NotificationOutput)(nil),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := radarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.UpdateNotification(test.WithRequest.(*radarr.NotificationInput))
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestDeleteNotification(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "notification", "2"),
			ExpectedMethod: "DELETE",
			WithRequest:    int64(2),
			ResponseStatus: 200,
			ResponseBody:   "{}",
			WithError:      nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "notification", "2"),
			ExpectedMethod: "DELETE",
			WithRequest:    int64(2),
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      &starr.ReqError{Code: http.StatusNotFound},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := radarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			err := client.DeleteNotification(test.WithRequest.(int64))
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
		})
	}
}
