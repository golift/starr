package lidarr_test

import (
	"net/http"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"golift.io/starr"
	"golift.io/starr/lidarr"
	"golift.io/starr/starrtest"
)

const notificationResponseBody = `{
    "onGrab": false,
    "onReleaseImport": false,
    "onUpgrade": true,
    "onRename": false,
    "onHealthIssue": false,
    "onDownloadFailure": false,
    "onImportFailure": false,
    "onTrackRetag": false,
    "onApplicationUpdate": false,
    "supportsOnGrab": true,
    "supportsOnReleaseImport": true,
    "supportsOnUpgrade": true,
    "supportsOnRename": true,
    "supportsOnHealthIssue": true,
    "includeHealthWarnings": false,
    "supportsOnDownloadFailure": false,
    "supportsOnImportFailure": false,
    "supportsOnTrackRetag": true,
    "supportsOnApplicationUpdate": true,
	"name": "Test",
	"fields": [
	  {
		"order": 0,
		"name": "path",
		"label": "Path",
		"value": "/scripts/lidarr.sh",
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
	"infoLink": "https://wiki.servarr.com/lidarr/supported#customscript",
	"message": {
	  "message": "Testing will execute the script with the EventType set to Test",
	  "type": "warning"
	},
	"tags": [],
	"id": 3
  }`

const addNotification = `{"onUpgrade":true,"name":"Test","implementation":"CustomScript","configContract":` +
	`"CustomScriptSettings","fields":[{"name":"path","value":"/scripts/lidarr.sh"}]}`

const updateNotification = `{"onUpgrade":true,"id":3,"name":"Test","implementation":"CustomScript","configContract":` +
	`"CustomScriptSettings","fields":[{"name":"path","value":"/scripts/lidarr.sh"}]}`

func TestGetNotifications(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:            "200",
			ExpectedPath:    path.Join("/", starr.API, lidarr.APIver, "notification"),
			ExpectedRequest: "",
			ExpectedMethod:  "GET",
			ResponseStatus:  200,
			ResponseBody:    "[" + notificationResponseBody + "]",
			WithRequest:     nil,
			WithResponse: []*lidarr.NotificationOutput{
				{
					OnUpgrade:                   true,
					SupportsOnGrab:              true,
					SupportsOnReleaseImport:     true,
					SupportsOnUpgrade:           true,
					SupportsOnRename:            true,
					SupportsOnApplicationUpdate: true,
					SupportsOnTrackRetag:        true,
					SupportsOnDownloadFailure:   false,
					SupportsOnImportFailure:     false,
					SupportsOnHealthIssue:       true,
					ID:                          3,
					Name:                        "Test",
					ImplementationName:          "Custom Script",
					Implementation:              "CustomScript",
					ConfigContract:              "CustomScriptSettings",
					InfoLink:                    "https://wiki.servarr.com/lidarr/supported#customscript",
					Tags:                        []int{},
					Fields: []*starr.FieldOutput{
						{
							Order:    0,
							Name:     "path",
							Label:    "Path",
							Value:    "/scripts/lidarr.sh",
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
			ExpectedPath:   path.Join("/", starr.API, lidarr.APIver, "notification"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:   ([]*lidarr.NotificationOutput)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := lidarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetNotifications()
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestGetNotification(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:            "200",
			ExpectedPath:    path.Join("/", starr.API, lidarr.APIver, "notification", "1"),
			ExpectedRequest: "",
			ExpectedMethod:  "GET",
			ResponseStatus:  200,
			ResponseBody:    notificationResponseBody,
			WithRequest:     nil,
			WithResponse: &lidarr.NotificationOutput{
				OnUpgrade:                   true,
				SupportsOnGrab:              true,
				SupportsOnReleaseImport:     true,
				SupportsOnUpgrade:           true,
				SupportsOnRename:            true,
				SupportsOnApplicationUpdate: true,
				SupportsOnTrackRetag:        true,
				SupportsOnDownloadFailure:   false,
				SupportsOnImportFailure:     false,
				SupportsOnHealthIssue:       true,
				ID:                          3,
				Name:                        "Test",
				ImplementationName:          "Custom Script",
				Implementation:              "CustomScript",
				ConfigContract:              "CustomScriptSettings",
				InfoLink:                    "https://wiki.servarr.com/lidarr/supported#customscript",
				Tags:                        []int{},
				Fields: []*starr.FieldOutput{
					{
						Order:    0,
						Name:     "path",
						Label:    "Path",
						Value:    "/scripts/lidarr.sh",
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
			ExpectedPath:   path.Join("/", starr.API, lidarr.APIver, "notification", "1"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:   (*lidarr.NotificationOutput)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := lidarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetNotification(1)
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestAddNotification(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, lidarr.APIver, "notification"),
			ExpectedMethod: "POST",
			ResponseStatus: 200,
			WithRequest: &lidarr.NotificationInput{
				OnUpgrade:      true,
				Name:           "Test",
				Implementation: "CustomScript",
				ConfigContract: "CustomScriptSettings",
				Fields: []*starr.FieldInput{
					{
						Name:  "path",
						Value: "/scripts/lidarr.sh",
					},
				},
			},
			ExpectedRequest: addNotification + "\n",
			ResponseBody:    notificationResponseBody,
			WithResponse: &lidarr.NotificationOutput{
				OnUpgrade:                   true,
				SupportsOnGrab:              true,
				SupportsOnReleaseImport:     true,
				SupportsOnUpgrade:           true,
				SupportsOnRename:            true,
				SupportsOnApplicationUpdate: true,
				SupportsOnTrackRetag:        true,
				SupportsOnDownloadFailure:   false,
				SupportsOnImportFailure:     false,
				SupportsOnHealthIssue:       true,
				ID:                          3,
				Name:                        "Test",
				ImplementationName:          "Custom Script",
				Implementation:              "CustomScript",
				ConfigContract:              "CustomScriptSettings",
				InfoLink:                    "https://wiki.servarr.com/lidarr/supported#customscript",
				Tags:                        []int{},
				Fields: []*starr.FieldOutput{
					{
						Order:    0,
						Name:     "path",
						Label:    "Path",
						Value:    "/scripts/lidarr.sh",
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
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, lidarr.APIver, "notification"),
			ExpectedMethod: "POST",
			ResponseStatus: 404,
			WithRequest: &lidarr.NotificationInput{
				OnUpgrade:      true,
				Name:           "Test",
				Implementation: "CustomScript",
				ConfigContract: "CustomScriptSettings",
				Fields: []*starr.FieldInput{
					{
						Name:  "path",
						Value: "/scripts/lidarr.sh",
					},
				},
			},
			ExpectedRequest: addNotification + "\n",
			ResponseBody:    `{"message": "NotFound"}`,
			WithError:       &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:    (*lidarr.NotificationOutput)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := lidarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.AddNotification(test.WithRequest.(*lidarr.NotificationInput))
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestUpdateNotification(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, lidarr.APIver, "notification", "3"),
			ExpectedMethod: "PUT",
			ResponseStatus: 200,
			WithRequest: &lidarr.NotificationInput{
				OnUpgrade:      true,
				ID:             3,
				Name:           "Test",
				Implementation: "CustomScript",
				ConfigContract: "CustomScriptSettings",
				Fields: []*starr.FieldInput{
					{
						Name:  "path",
						Value: "/scripts/lidarr.sh",
					},
				},
			},
			ExpectedRequest: updateNotification + "\n",
			ResponseBody:    notificationResponseBody,
			WithResponse: &lidarr.NotificationOutput{
				OnUpgrade:                   true,
				SupportsOnGrab:              true,
				SupportsOnReleaseImport:     true,
				SupportsOnUpgrade:           true,
				SupportsOnRename:            true,
				SupportsOnApplicationUpdate: true,
				SupportsOnTrackRetag:        true,
				SupportsOnDownloadFailure:   false,
				SupportsOnImportFailure:     false,
				SupportsOnHealthIssue:       true,
				ID:                          3,
				Name:                        "Test",
				ImplementationName:          "Custom Script",
				Implementation:              "CustomScript",
				ConfigContract:              "CustomScriptSettings",
				InfoLink:                    "https://wiki.servarr.com/lidarr/supported#customscript",
				Tags:                        []int{},
				Fields: []*starr.FieldOutput{
					{
						Order:    0,
						Name:     "path",
						Label:    "Path",
						Value:    "/scripts/lidarr.sh",
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
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, lidarr.APIver, "notification", "3"),
			ExpectedMethod: "PUT",
			ResponseStatus: 404,
			WithRequest: &lidarr.NotificationInput{
				OnUpgrade:      true,
				ID:             3,
				Name:           "Test",
				Implementation: "CustomScript",
				ConfigContract: "CustomScriptSettings",
				Fields: []*starr.FieldInput{
					{
						Name:  "path",
						Value: "/scripts/lidarr.sh",
					},
				},
			},
			ExpectedRequest: updateNotification + "\n",
			ResponseBody:    `{"message": "NotFound"}`,
			WithError:       &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:    (*lidarr.NotificationOutput)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := lidarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.UpdateNotification(test.WithRequest.(*lidarr.NotificationInput))
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestDeleteNotification(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, lidarr.APIver, "notification", "2"),
			ExpectedMethod: "DELETE",
			WithRequest:    int64(2),
			ResponseStatus: 200,
			ResponseBody:   "{}",
			WithError:      nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, lidarr.APIver, "notification", "2"),
			ExpectedMethod: "DELETE",
			WithRequest:    int64(2),
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      &starr.ReqError{Code: http.StatusNotFound},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := lidarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			err := client.DeleteNotification(test.WithRequest.(int64))
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
		})
	}
}
