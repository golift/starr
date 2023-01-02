package readarr_test

import (
	"net/http"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"golift.io/starr"
	"golift.io/starr/readarr"
	"golift.io/starr/starrtest"
)

const notificationResponseBody = `{
    "onGrab": false,
    "onReleaseImport": false,
    "onUpgrade": true,
    "onRename": false,
    "onAuthorDelete": false,
    "onBookDelete": false,
    "onBookFileDelete": false,
    "onBookFileDeleteForUpgrade": false,
    "onHealthIssue": false,
    "onDownloadFailure": false,
    "onImportFailure": false,
    "onBookRetag": false,
    "supportsOnGrab": true,
    "supportsOnReleaseImport": true,
    "supportsOnUpgrade": true,
    "supportsOnRename": true,
    "supportsOnAuthorDelete": true,
    "supportsOnBookDelete": true,
    "supportsOnBookFileDelete": true,
    "supportsOnBookFileDeleteForUpgrade": true,
    "supportsOnHealthIssue": true,
    "includeHealthWarnings": false,
    "supportsOnDownloadFailure": false,
    "supportsOnImportFailure": false,
    "supportsOnBookRetag": true,
	"name": "Test",
	"fields": [
	  {
		"order": 0,
		"name": "path",
		"label": "Path",
		"value": "/scripts/readarr.sh",
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
	"infoLink": "https://wiki.servarr.com/readarr/supported#customscript",
	"message": {
	  "message": "Testing will execute the script with the EventType set to Test",
	  "type": "warning"
	},
	"tags": [],
	"id": 3
  }`

const addNotification = `{"onUpgrade":true,"name":"Test","implementation":"CustomScript","configContract":` +
	`"CustomScriptSettings","fields":[{"name":"path","value":"/scripts/readarr.sh"}]}`

const updateNotification = `{"onUpgrade":true,"id":3,"name":"Test","implementation":"CustomScript","configContract":` +
	`"CustomScriptSettings","fields":[{"name":"path","value":"/scripts/readarr.sh"}]}`

func TestGetNotifications(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:            "200",
			ExpectedPath:    path.Join("/", starr.API, readarr.APIver, "notification"),
			ExpectedRequest: "",
			ExpectedMethod:  "GET",
			ResponseStatus:  200,
			ResponseBody:    "[" + notificationResponseBody + "]",
			WithRequest:     nil,
			WithResponse: []*readarr.NotificationOutput{
				{
					OnUpgrade:                          true,
					SupportsOnGrab:                     true,
					SupportsOnReleaseImport:            true,
					SupportsOnUpgrade:                  true,
					SupportsOnRename:                   true,
					SupportsOnAuthorDelete:             true,
					SupportsOnBookDelete:               true,
					SupportsOnBookFileDelete:           true,
					SupportsOnBookFileDeleteForUpgrade: true,
					SupportsOnDownloadFailure:          false,
					SupportsOnImportFailure:            false,
					SupportsOnBookRetag:                true,
					SupportsOnHealthIssue:              true,
					ID:                                 3,
					Name:                               "Test",
					ImplementationName:                 "Custom Script",
					Implementation:                     "CustomScript",
					ConfigContract:                     "CustomScriptSettings",
					InfoLink:                           "https://wiki.servarr.com/readarr/supported#customscript",
					Tags:                               []int{},
					Fields: []*starr.FieldOutput{
						{
							Order:    0,
							Name:     "path",
							Label:    "Path",
							Value:    "/scripts/readarr.sh",
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
			ExpectedPath:   path.Join("/", starr.API, readarr.APIver, "notification"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:   ([]*readarr.NotificationOutput)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := readarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
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
			ExpectedPath:    path.Join("/", starr.API, readarr.APIver, "notification", "1"),
			ExpectedRequest: "",
			ExpectedMethod:  "GET",
			ResponseStatus:  200,
			ResponseBody:    notificationResponseBody,
			WithRequest:     nil,
			WithResponse: &readarr.NotificationOutput{
				OnUpgrade:                          true,
				SupportsOnGrab:                     true,
				SupportsOnReleaseImport:            true,
				SupportsOnUpgrade:                  true,
				SupportsOnRename:                   true,
				SupportsOnAuthorDelete:             true,
				SupportsOnBookDelete:               true,
				SupportsOnBookFileDelete:           true,
				SupportsOnBookFileDeleteForUpgrade: true,
				SupportsOnDownloadFailure:          false,
				SupportsOnImportFailure:            false,
				SupportsOnBookRetag:                true,
				SupportsOnHealthIssue:              true,
				ID:                                 3,
				Name:                               "Test",
				ImplementationName:                 "Custom Script",
				Implementation:                     "CustomScript",
				ConfigContract:                     "CustomScriptSettings",
				InfoLink:                           "https://wiki.servarr.com/readarr/supported#customscript",
				Tags:                               []int{},
				Fields: []*starr.FieldOutput{
					{
						Order:    0,
						Name:     "path",
						Label:    "Path",
						Value:    "/scripts/readarr.sh",
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
			ExpectedPath:   path.Join("/", starr.API, readarr.APIver, "notification", "1"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:   (*readarr.NotificationOutput)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := readarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
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
			ExpectedPath:   path.Join("/", starr.API, readarr.APIver, "notification"),
			ExpectedMethod: "POST",
			ResponseStatus: 200,
			WithRequest: &readarr.NotificationInput{
				OnUpgrade:      true,
				Name:           "Test",
				Implementation: "CustomScript",
				ConfigContract: "CustomScriptSettings",
				Fields: []*starr.FieldInput{
					{
						Name:  "path",
						Value: "/scripts/readarr.sh",
					},
				},
			},
			ExpectedRequest: addNotification + "\n",
			ResponseBody:    notificationResponseBody,
			WithResponse: &readarr.NotificationOutput{
				OnUpgrade:                          true,
				SupportsOnGrab:                     true,
				SupportsOnReleaseImport:            true,
				SupportsOnUpgrade:                  true,
				SupportsOnRename:                   true,
				SupportsOnAuthorDelete:             true,
				SupportsOnBookDelete:               true,
				SupportsOnBookFileDelete:           true,
				SupportsOnBookFileDeleteForUpgrade: true,
				SupportsOnDownloadFailure:          false,
				SupportsOnImportFailure:            false,
				SupportsOnBookRetag:                true,
				SupportsOnHealthIssue:              true,
				ID:                                 3,
				Name:                               "Test",
				ImplementationName:                 "Custom Script",
				Implementation:                     "CustomScript",
				ConfigContract:                     "CustomScriptSettings",
				InfoLink:                           "https://wiki.servarr.com/readarr/supported#customscript",
				Tags:                               []int{},
				Fields: []*starr.FieldOutput{
					{
						Order:    0,
						Name:     "path",
						Label:    "Path",
						Value:    "/scripts/readarr.sh",
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
			ExpectedPath:   path.Join("/", starr.API, readarr.APIver, "notification"),
			ExpectedMethod: "POST",
			ResponseStatus: 404,
			WithRequest: &readarr.NotificationInput{
				OnUpgrade:      true,
				Name:           "Test",
				Implementation: "CustomScript",
				ConfigContract: "CustomScriptSettings",
				Fields: []*starr.FieldInput{
					{
						Name:  "path",
						Value: "/scripts/readarr.sh",
					},
				},
			},
			ExpectedRequest: addNotification + "\n",
			ResponseBody:    `{"message": "NotFound"}`,
			WithError:       &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:    (*readarr.NotificationOutput)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := readarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.AddNotification(test.WithRequest.(*readarr.NotificationInput))
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
			ExpectedPath:   path.Join("/", starr.API, readarr.APIver, "notification", "3"),
			ExpectedMethod: "PUT",
			ResponseStatus: 200,
			WithRequest: &readarr.NotificationInput{
				OnUpgrade:      true,
				ID:             3,
				Name:           "Test",
				Implementation: "CustomScript",
				ConfigContract: "CustomScriptSettings",
				Fields: []*starr.FieldInput{
					{
						Name:  "path",
						Value: "/scripts/readarr.sh",
					},
				},
			},
			ExpectedRequest: updateNotification + "\n",
			ResponseBody:    notificationResponseBody,
			WithResponse: &readarr.NotificationOutput{
				OnUpgrade:                          true,
				SupportsOnGrab:                     true,
				SupportsOnReleaseImport:            true,
				SupportsOnUpgrade:                  true,
				SupportsOnRename:                   true,
				SupportsOnAuthorDelete:             true,
				SupportsOnBookDelete:               true,
				SupportsOnBookFileDelete:           true,
				SupportsOnBookFileDeleteForUpgrade: true,
				SupportsOnDownloadFailure:          false,
				SupportsOnImportFailure:            false,
				SupportsOnBookRetag:                true,
				SupportsOnHealthIssue:              true,
				ID:                                 3,
				Name:                               "Test",
				ImplementationName:                 "Custom Script",
				Implementation:                     "CustomScript",
				ConfigContract:                     "CustomScriptSettings",
				InfoLink:                           "https://wiki.servarr.com/readarr/supported#customscript",
				Tags:                               []int{},
				Fields: []*starr.FieldOutput{
					{
						Order:    0,
						Name:     "path",
						Label:    "Path",
						Value:    "/scripts/readarr.sh",
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
			ExpectedPath:   path.Join("/", starr.API, readarr.APIver, "notification", "3"),
			ExpectedMethod: "PUT",
			ResponseStatus: 404,
			WithRequest: &readarr.NotificationInput{
				OnUpgrade:      true,
				ID:             3,
				Name:           "Test",
				Implementation: "CustomScript",
				ConfigContract: "CustomScriptSettings",
				Fields: []*starr.FieldInput{
					{
						Name:  "path",
						Value: "/scripts/readarr.sh",
					},
				},
			},
			ExpectedRequest: updateNotification + "\n",
			ResponseBody:    `{"message": "NotFound"}`,
			WithError:       &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:    (*readarr.NotificationOutput)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := readarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.UpdateNotification(test.WithRequest.(*readarr.NotificationInput))
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
			ExpectedPath:   path.Join("/", starr.API, readarr.APIver, "notification", "2"),
			ExpectedMethod: "DELETE",
			WithRequest:    int64(2),
			ResponseStatus: 200,
			ResponseBody:   "{}",
			WithError:      nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, readarr.APIver, "notification", "2"),
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
			client := readarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			err := client.DeleteNotification(test.WithRequest.(int64))
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
		})
	}
}
