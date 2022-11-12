package prowlarr_test

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"golift.io/starr"
	"golift.io/starr/prowlarr"
)

const notificationResponseBody = `{
    "onHealthIssue": false,
    "onApplicationUpdate": true,
    "supportsOnHealthIssue": true,
    "includeHealthWarnings": false,
    "supportsOnApplicationUpdate": true,
	"name": "Test",
	"fields": [
	  {
		"order": 0,
		"name": "path",
		"label": "Path",
		"value": "/scripts/prowlarr.sh",
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
	"infoLink": "https://wiki.servarr.com/prowlarr/supported#customscript",
	"message": {
	  "message": "Testing will execute the script with the EventType set to Test",
	  "type": "warning"
	},
	"tags": [],
	"id": 3
  }`

const addNotification = `{"onApplicationUpdate":true,"name":"Test","implementation":"CustomScript","configContract":` +
	`"CustomScriptSettings","fields":[{"name":"path","value":"/scripts/prowlarr.sh"}]}`

const updateNotification = `{"onApplicationUpdate":true,"id":3,"name":"Test","implementation":"CustomScript",` +
	`"configContract":"CustomScriptSettings","fields":[{"name":"path","value":"/scripts/prowlarr.sh"}]}`

func TestGetNotifications(t *testing.T) {
	t.Parallel()

	tests := []*starr.TestMockData{
		{
			Name:            "200",
			ExpectedPath:    path.Join("/", starr.API, prowlarr.APIver, "notification"),
			ExpectedRequest: "",
			ExpectedMethod:  "GET",
			ResponseStatus:  200,
			ResponseBody:    "[" + notificationResponseBody + "]",
			WithRequest:     nil,
			WithResponse: []*prowlarr.NotificationOutput{
				{
					OnApplicationUpdate:         true,
					SupportsOnHealthIssue:       true,
					SupportsOnApplicationUpdate: true,
					ID:                          3,
					Name:                        "Test",
					ImplementationName:          "Custom Script",
					Implementation:              "CustomScript",
					ConfigContract:              "CustomScriptSettings",
					InfoLink:                    "https://wiki.servarr.com/prowlarr/supported#customscript",
					Tags:                        []int{},
					Fields: []*starr.FieldOutput{
						{
							Order:    0,
							Name:     "path",
							Label:    "Path",
							Value:    "/scripts/prowlarr.sh",
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
			ExpectedPath:   path.Join("/", starr.API, prowlarr.APIver, "notification"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      starr.ErrInvalidStatusCode,
			WithResponse:   ([]*prowlarr.NotificationOutput)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := prowlarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetNotifications()
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestGetNotification(t *testing.T) {
	t.Parallel()

	tests := []*starr.TestMockData{
		{
			Name:            "200",
			ExpectedPath:    path.Join("/", starr.API, prowlarr.APIver, "notification", "1"),
			ExpectedRequest: "",
			ExpectedMethod:  "GET",
			ResponseStatus:  200,
			ResponseBody:    notificationResponseBody,
			WithRequest:     nil,
			WithResponse: &prowlarr.NotificationOutput{
				OnApplicationUpdate:         true,
				SupportsOnHealthIssue:       true,
				SupportsOnApplicationUpdate: true,
				ID:                          3,
				Name:                        "Test",
				ImplementationName:          "Custom Script",
				Implementation:              "CustomScript",
				ConfigContract:              "CustomScriptSettings",
				InfoLink:                    "https://wiki.servarr.com/prowlarr/supported#customscript",
				Tags:                        []int{},
				Fields: []*starr.FieldOutput{
					{
						Order:    0,
						Name:     "path",
						Label:    "Path",
						Value:    "/scripts/prowlarr.sh",
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
			ExpectedPath:   path.Join("/", starr.API, prowlarr.APIver, "notification", "1"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      starr.ErrInvalidStatusCode,
			WithResponse:   (*prowlarr.NotificationOutput)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := prowlarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetNotification(1)
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestAddNotification(t *testing.T) {
	t.Parallel()

	tests := []*starr.TestMockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, prowlarr.APIver, "notification"),
			ExpectedMethod: "POST",
			ResponseStatus: 200,
			WithRequest: &prowlarr.NotificationInput{
				OnApplicationUpdate: true,
				Name:                "Test",
				Implementation:      "CustomScript",
				ConfigContract:      "CustomScriptSettings",
				Fields: []*starr.FieldInput{
					{
						Name:  "path",
						Value: "/scripts/prowlarr.sh",
					},
				},
			},
			ExpectedRequest: addNotification + "\n",
			ResponseBody:    notificationResponseBody,
			WithResponse: &prowlarr.NotificationOutput{
				OnApplicationUpdate:         true,
				SupportsOnHealthIssue:       true,
				SupportsOnApplicationUpdate: true,
				ID:                          3,
				Name:                        "Test",
				ImplementationName:          "Custom Script",
				Implementation:              "CustomScript",
				ConfigContract:              "CustomScriptSettings",
				InfoLink:                    "https://wiki.servarr.com/prowlarr/supported#customscript",
				Tags:                        []int{},
				Fields: []*starr.FieldOutput{
					{
						Order:    0,
						Name:     "path",
						Label:    "Path",
						Value:    "/scripts/prowlarr.sh",
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
			ExpectedPath:   path.Join("/", starr.API, prowlarr.APIver, "notification"),
			ExpectedMethod: "POST",
			ResponseStatus: 404,
			WithRequest: &prowlarr.NotificationInput{
				OnApplicationUpdate: true,
				Name:                "Test",
				Implementation:      "CustomScript",
				ConfigContract:      "CustomScriptSettings",
				Fields: []*starr.FieldInput{
					{
						Name:  "path",
						Value: "/scripts/prowlarr.sh",
					},
				},
			},
			ExpectedRequest: addNotification + "\n",
			ResponseBody:    `{"message": "NotFound"}`,
			WithError:       starr.ErrInvalidStatusCode,
			WithResponse:    (*prowlarr.NotificationOutput)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := prowlarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.AddNotification(test.WithRequest.(*prowlarr.NotificationInput))
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestUpdateNotification(t *testing.T) {
	t.Parallel()

	tests := []*starr.TestMockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, prowlarr.APIver, "notification", "3"),
			ExpectedMethod: "PUT",
			ResponseStatus: 200,
			WithRequest: &prowlarr.NotificationInput{
				OnApplicationUpdate: true,
				ID:                  3,
				Name:                "Test",
				Implementation:      "CustomScript",
				ConfigContract:      "CustomScriptSettings",
				Fields: []*starr.FieldInput{
					{
						Name:  "path",
						Value: "/scripts/prowlarr.sh",
					},
				},
			},
			ExpectedRequest: updateNotification + "\n",
			ResponseBody:    notificationResponseBody,
			WithResponse: &prowlarr.NotificationOutput{
				OnApplicationUpdate:         true,
				SupportsOnHealthIssue:       true,
				SupportsOnApplicationUpdate: true,
				ID:                          3,
				Name:                        "Test",
				ImplementationName:          "Custom Script",
				Implementation:              "CustomScript",
				ConfigContract:              "CustomScriptSettings",
				InfoLink:                    "https://wiki.servarr.com/prowlarr/supported#customscript",
				Tags:                        []int{},
				Fields: []*starr.FieldOutput{
					{
						Order:    0,
						Name:     "path",
						Label:    "Path",
						Value:    "/scripts/prowlarr.sh",
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
			ExpectedPath:   path.Join("/", starr.API, prowlarr.APIver, "notification", "3"),
			ExpectedMethod: "PUT",
			ResponseStatus: 404,
			WithRequest: &prowlarr.NotificationInput{
				OnApplicationUpdate: true,
				ID:                  3,
				Name:                "Test",
				Implementation:      "CustomScript",
				ConfigContract:      "CustomScriptSettings",
				Fields: []*starr.FieldInput{
					{
						Name:  "path",
						Value: "/scripts/prowlarr.sh",
					},
				},
			},
			ExpectedRequest: updateNotification + "\n",
			ResponseBody:    `{"message": "NotFound"}`,
			WithError:       starr.ErrInvalidStatusCode,
			WithResponse:    (*prowlarr.NotificationOutput)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := prowlarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.UpdateNotification(test.WithRequest.(*prowlarr.NotificationInput))
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestDeleteNotification(t *testing.T) {
	t.Parallel()

	tests := []*starr.TestMockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, prowlarr.APIver, "notification", "2"),
			ExpectedMethod: "DELETE",
			WithRequest:    int64(2),
			ResponseStatus: 200,
			ResponseBody:   "{}",
			WithError:      nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, prowlarr.APIver, "notification", "2"),
			ExpectedMethod: "DELETE",
			WithRequest:    int64(2),
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
			client := prowlarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			err := client.DeleteNotification(test.WithRequest.(int64))
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
		})
	}
}
