package sonarr_test

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"golift.io/starr"
	"golift.io/starr/sonarr"
)

const notificationResponseBody = `{
	"onGrab": false,
	"onDownload": true,
	"onUpgrade": true,
	"onRename": false,
	"onSeriesDelete": false,
	"onEpisodeFileDelete": false,
	"onEpisodeFileDeleteForUpgrade": true,
	"onHealthIssue": false,
	"onApplicationUpdate": false,
	"supportsOnGrab": true,
	"supportsOnDownload": true,
	"supportsOnUpgrade": true,
	"supportsOnRename": true,
	"supportsOnSeriesDelete": true,
	"supportsOnEpisodeFileDelete": true,
	"supportsOnEpisodeFileDeleteForUpgrade": true,
	"supportsOnHealthIssue": true,
	"supportsOnApplicationUpdate": true,
	"includeHealthWarnings": false,
	"name": "Test",
	"fields": [
	  {
		"order": 0,
		"name": "path",
		"label": "Path",
		"value": "/scripts/sonarr.sh",
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
	"infoLink": "https://wiki.servarr.com/sonarr/supported#customscript",
	"message": {
	  "message": "Testing will execute the script with the EventType set to Test",
	  "type": "warning"
	},
	"tags": [],
	"id": 3
  }`

const addNotification = `{"onGrab":false,"onDownload":true,"onUpgrade":true,"onRename":false,` +
	`"onSeriesDelete":false,"onEpisodeFileDelete":true,"onEpisodeFileDeleteForUpgrade":false,` +
	`"onHealthIssue":false,"onApplicationUpdate":false,"includeHealthWarnings":false,"name":"Test",` +
	`"implementation":"CustomScript","configContract":"CustomScriptSettings","tags":null,` +
	`"fields":[{"name":"path","value":"/scripts/sonarr.sh"},{"name":"apiPath","value":"/api"}]}`

const updateNotification = `{"onGrab":false,"onDownload":true,"onUpgrade":true,"onRename":false,` +
	`"onSeriesDelete":false,"onEpisodeFileDelete":true,"onEpisodeFileDeleteForUpgrade":false,` +
	`"onHealthIssue":false,"onApplicationUpdate":false,"includeHealthWarnings":false,"id":3,"name":"Test",` +
	`"implementation":"CustomScript","configContract":"CustomScriptSettings","tags":null,` +
	`"fields":[{"name":"path","value":"/scripts/sonarr.sh"},{"name":"apiPath","value":"/api"}]}`

func TestGetNotifications(t *testing.T) {
	t.Parallel()

	tests := []*starr.TestMockData{
		{
			Name:            "200",
			ExpectedPath:    path.Join("/", starr.API, sonarr.APIver, "notification"),
			ExpectedRequest: "",
			ExpectedMethod:  "GET",
			ResponseStatus:  200,
			ResponseBody:    "[" + notificationResponseBody + "]",
			WithRequest:     nil,
			WithResponse: []*sonarr.NotificationOutput{
				{
					OnGrab:                                false,
					OnDownload:                            true,
					OnUpgrade:                             true,
					OnRename:                              false,
					OnSeriesDelete:                        false,
					OnEpisodeFileDelete:                   false,
					OnEpisodeFileDeleteForUpgrade:         true,
					OnHealthIssue:                         false,
					OnApplicationUpdate:                   false,
					SupportsOnGrab:                        true,
					SupportsOnDownload:                    true,
					SupportsOnUpgrade:                     true,
					SupportsOnRename:                      true,
					SupportsOnSeriesDelete:                true,
					SupportsOnEpisodeFileDelete:           true,
					SupportsOnEpisodeFileDeleteForUpgrade: true,
					SupportsOnHealthIssue:                 true,
					SupportsOnApplicationUpdate:           true,
					IncludeHealthWarnings:                 false,
					ID:                                    3,
					Name:                                  "Test",
					ImplementationName:                    "Custom Script",
					Implementation:                        "CustomScript",
					ConfigContract:                        "CustomScriptSettings",
					InfoLink:                              "https://wiki.servarr.com/sonarr/supported#customscript",
					Tags:                                  []int{},
					Fields: []*starr.FieldOutput{
						{
							Order:    0,
							Name:     "path",
							Label:    "Path",
							Value:    "/scripts/sonarr.sh",
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
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "notification"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      starr.ErrInvalidStatusCode,
			WithResponse:   ([]*sonarr.NotificationOutput)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
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
			ExpectedPath:    path.Join("/", starr.API, sonarr.APIver, "notification", "1"),
			ExpectedRequest: "",
			ExpectedMethod:  "GET",
			ResponseStatus:  200,
			ResponseBody:    notificationResponseBody,
			WithRequest:     nil,
			WithResponse: &sonarr.NotificationOutput{
				OnGrab:                                false,
				OnDownload:                            true,
				OnUpgrade:                             true,
				OnRename:                              false,
				OnSeriesDelete:                        false,
				OnEpisodeFileDelete:                   false,
				OnEpisodeFileDeleteForUpgrade:         true,
				OnHealthIssue:                         false,
				OnApplicationUpdate:                   false,
				SupportsOnGrab:                        true,
				SupportsOnDownload:                    true,
				SupportsOnUpgrade:                     true,
				SupportsOnRename:                      true,
				SupportsOnSeriesDelete:                true,
				SupportsOnEpisodeFileDelete:           true,
				SupportsOnEpisodeFileDeleteForUpgrade: true,
				SupportsOnHealthIssue:                 true,
				SupportsOnApplicationUpdate:           true,
				IncludeHealthWarnings:                 false,
				ID:                                    3,
				Name:                                  "Test",
				ImplementationName:                    "Custom Script",
				Implementation:                        "CustomScript",
				ConfigContract:                        "CustomScriptSettings",
				InfoLink:                              "https://wiki.servarr.com/sonarr/supported#customscript",
				Tags:                                  []int{},
				Fields: []*starr.FieldOutput{
					{
						Order:    0,
						Name:     "path",
						Label:    "Path",
						Value:    "/scripts/sonarr.sh",
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
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "notification", "1"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      starr.ErrInvalidStatusCode,
			WithResponse:   (*sonarr.NotificationOutput)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
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
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "notification"),
			ExpectedMethod: "POST",
			ResponseStatus: 200,
			WithRequest: &sonarr.NotificationInput{
				OnGrab:                        false,
				OnDownload:                    true,
				OnUpgrade:                     true,
				OnRename:                      false,
				OnSeriesDelete:                false,
				OnEpisodeFileDelete:           true,
				OnEpisodeFileDeleteForUpgrade: false,
				OnHealthIssue:                 false,
				OnApplicationUpdate:           false,
				IncludeHealthWarnings:         false,
				Name:                          "Test",
				Implementation:                "CustomScript",
				ConfigContract:                "CustomScriptSettings",
				Fields: []*starr.FieldInput{
					{
						Name:  "path",
						Value: "/scripts/sonarr.sh",
					},
					{
						Name:  "apiPath",
						Value: "/api",
					},
				},
			},
			ExpectedRequest: addNotification + "\n",
			ResponseBody:    notificationResponseBody,
			WithResponse: &sonarr.NotificationOutput{
				OnGrab:                                false,
				OnDownload:                            true,
				OnUpgrade:                             true,
				OnRename:                              false,
				OnSeriesDelete:                        false,
				OnEpisodeFileDelete:                   false,
				OnEpisodeFileDeleteForUpgrade:         true,
				OnHealthIssue:                         false,
				OnApplicationUpdate:                   false,
				SupportsOnGrab:                        true,
				SupportsOnDownload:                    true,
				SupportsOnUpgrade:                     true,
				SupportsOnRename:                      true,
				SupportsOnSeriesDelete:                true,
				SupportsOnEpisodeFileDelete:           true,
				SupportsOnEpisodeFileDeleteForUpgrade: true,
				SupportsOnHealthIssue:                 true,
				SupportsOnApplicationUpdate:           true,
				IncludeHealthWarnings:                 false,
				ID:                                    3,
				Name:                                  "Test",
				ImplementationName:                    "Custom Script",
				Implementation:                        "CustomScript",
				ConfigContract:                        "CustomScriptSettings",
				InfoLink:                              "https://wiki.servarr.com/sonarr/supported#customscript",
				Tags:                                  []int{},
				Fields: []*starr.FieldOutput{
					{
						Order:    0,
						Name:     "path",
						Label:    "Path",
						Value:    "/scripts/sonarr.sh",
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
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "notification"),
			ExpectedMethod: "POST",
			ResponseStatus: 404,
			WithRequest: &sonarr.NotificationInput{
				OnGrab:                        false,
				OnDownload:                    true,
				OnUpgrade:                     true,
				OnRename:                      false,
				OnSeriesDelete:                false,
				OnEpisodeFileDelete:           true,
				OnEpisodeFileDeleteForUpgrade: false,
				OnHealthIssue:                 false,
				OnApplicationUpdate:           false,
				IncludeHealthWarnings:         false,
				Name:                          "Test",
				Implementation:                "CustomScript",
				ConfigContract:                "CustomScriptSettings",
				Fields: []*starr.FieldInput{
					{
						Name:  "path",
						Value: "/scripts/sonarr.sh",
					},
					{
						Name:  "apiPath",
						Value: "/api",
					},
				},
			},
			ExpectedRequest: addNotification + "\n",
			ResponseBody:    `{"message": "NotFound"}`,
			WithError:       starr.ErrInvalidStatusCode,
			WithResponse:    (*sonarr.NotificationOutput)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.AddNotification(test.WithRequest.(*sonarr.NotificationInput))
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
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "notification", "3"),
			ExpectedMethod: "PUT",
			ResponseStatus: 200,
			WithRequest: &sonarr.NotificationInput{
				OnGrab:                        false,
				OnDownload:                    true,
				OnUpgrade:                     true,
				OnRename:                      false,
				OnSeriesDelete:                false,
				OnEpisodeFileDelete:           true,
				OnEpisodeFileDeleteForUpgrade: false,
				OnHealthIssue:                 false,
				OnApplicationUpdate:           false,
				IncludeHealthWarnings:         false,
				ID:                            3,
				Name:                          "Test",
				Implementation:                "CustomScript",
				ConfigContract:                "CustomScriptSettings",
				Fields: []*starr.FieldInput{
					{
						Name:  "path",
						Value: "/scripts/sonarr.sh",
					},
					{
						Name:  "apiPath",
						Value: "/api",
					},
				},
			},
			ExpectedRequest: updateNotification + "\n",
			ResponseBody:    notificationResponseBody,
			WithResponse: &sonarr.NotificationOutput{
				OnGrab:                                false,
				OnDownload:                            true,
				OnUpgrade:                             true,
				OnRename:                              false,
				OnSeriesDelete:                        false,
				OnEpisodeFileDelete:                   false,
				OnEpisodeFileDeleteForUpgrade:         true,
				OnHealthIssue:                         false,
				OnApplicationUpdate:                   false,
				SupportsOnGrab:                        true,
				SupportsOnDownload:                    true,
				SupportsOnUpgrade:                     true,
				SupportsOnRename:                      true,
				SupportsOnSeriesDelete:                true,
				SupportsOnEpisodeFileDelete:           true,
				SupportsOnEpisodeFileDeleteForUpgrade: true,
				SupportsOnHealthIssue:                 true,
				SupportsOnApplicationUpdate:           true,
				IncludeHealthWarnings:                 false,
				ID:                                    3,
				Name:                                  "Test",
				ImplementationName:                    "Custom Script",
				Implementation:                        "CustomScript",
				ConfigContract:                        "CustomScriptSettings",
				InfoLink:                              "https://wiki.servarr.com/sonarr/supported#customscript",
				Tags:                                  []int{},
				Fields: []*starr.FieldOutput{
					{
						Order:    0,
						Name:     "path",
						Label:    "Path",
						Value:    "/scripts/sonarr.sh",
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
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "notification", "3"),
			ExpectedMethod: "PUT",
			ResponseStatus: 404,
			WithRequest: &sonarr.NotificationInput{
				OnGrab:                        false,
				OnDownload:                    true,
				OnUpgrade:                     true,
				OnRename:                      false,
				OnSeriesDelete:                false,
				OnEpisodeFileDelete:           true,
				OnEpisodeFileDeleteForUpgrade: false,
				OnHealthIssue:                 false,
				OnApplicationUpdate:           false,
				IncludeHealthWarnings:         false,
				ID:                            3,
				Name:                          "Test",
				Implementation:                "CustomScript",
				ConfigContract:                "CustomScriptSettings",
				Fields: []*starr.FieldInput{
					{
						Name:  "path",
						Value: "/scripts/sonarr.sh",
					},
					{
						Name:  "apiPath",
						Value: "/api",
					},
				},
			},
			ExpectedRequest: updateNotification + "\n",
			ResponseBody:    `{"message": "NotFound"}`,
			WithError:       starr.ErrInvalidStatusCode,
			WithResponse:    (*sonarr.NotificationOutput)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.UpdateNotification(test.WithRequest.(*sonarr.NotificationInput))
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
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "notification", "2"),
			ExpectedMethod: "DELETE",
			WithRequest:    int64(2),
			ResponseStatus: 200,
			ResponseBody:   "{}",
			WithError:      nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "notification", "2"),
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
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			err := client.DeleteNotification(test.WithRequest.(int64))
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
		})
	}
}
