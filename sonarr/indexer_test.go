package sonarr_test

import (
	"net/http"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golift.io/starr"
	"golift.io/starr/sonarr"
	"golift.io/starr/starrtest"
)

const indexerResponseBody = `{
	"enableRss": true,
	"enableAutomaticSearch": true,
	"enableInteractiveSearch": true,
	"supportsRss": true,
	"supportsSearch": true,
	"protocol": "usenet",
	"priority": 25,
	"downloadClientId": 0,
	"name": "NZBgeek",
	"fields": [
	  {
		"order": 0,
		"name": "baseUrl",
		"label": "URL",
		"value": "https://api.nzbgeek.info",
		"type": "textbox",
		"advanced": false
	  },
	  {
		"order": 1,
		"name": "apiPath",
		"label": "API Path",
		"helpText": "Path to the api, usually /api",
		"value": "/api",
		"type": "textbox",
		"advanced": true
	  }
	],
	"implementationName": "Newznab",
	"implementation": "Newznab",
	"configContract": "NewznabSettings",
	"infoLink": "https://wiki.servarr.com/sonarr/supported#newznab",
	"tags": [],
	"id": 1
  }`

const addIndexer = `{"enableAutomaticSearch":true,"enableInteractiveSearch":true,"enableRss":true,` +
	`"downloadClientId":0,"priority":25,"configContract":"NewznabSettings","implementation":"Newznab"` +
	`,"name":"NZBgeek","protocol":"usenet","tags":[],` +
	`"fields":[{"name":"baseUrl","value":"https://api.nzbgeek.info"},{"name":"apiPath","value":"/api"}]}`

const updateIndexer = `{"enableAutomaticSearch":true,"enableInteractiveSearch":true,"enableRss":true,` +
	`"downloadClientId":0,"priority":25,"id":1,"configContract":"NewznabSettings","implementation":"Newznab",` +
	`"name":"NZBgeek","protocol":"usenet","tags":[],` +
	`"fields":[{"name":"baseUrl","value":"https://api.nzbgeek.info"},{"name":"apiPath","value":"/api"}]}`

func TestGetIndexers(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:            "200",
			ExpectedPath:    path.Join("/", starr.API, sonarr.APIver, "indexer"),
			ExpectedRequest: "",
			ExpectedMethod:  "GET",
			ResponseStatus:  200,
			ResponseBody:    "[" + indexerResponseBody + "]",
			WithRequest:     nil,
			WithResponse: []*sonarr.IndexerOutput{
				{
					EnableAutomaticSearch:   true,
					EnableInteractiveSearch: true,
					EnableRss:               true,
					SupportsRss:             true,
					SupportsSearch:          true,
					Priority:                25,
					ID:                      1,
					ConfigContract:          "NewznabSettings",
					Implementation:          "Newznab",
					ImplementationName:      "Newznab",
					InfoLink:                "https://wiki.servarr.com/sonarr/supported#newznab",
					Name:                    "NZBgeek",
					Protocol:                "usenet",
					Fields: []*starr.FieldOutput{
						{
							Order:    0,
							Name:     "baseUrl",
							Label:    "URL",
							Value:    "https://api.nzbgeek.info",
							Type:     "textbox",
							Advanced: false,
						},
						{
							Order:    1,
							Name:     "apiPath",
							Label:    "API Path",
							HelpText: "Path to the api, usually /api",
							Value:    "/api",
							Type:     "textbox",
							Advanced: true,
						},
					},
					Tags: []int{},
				},
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "indexer"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:   ([]*sonarr.IndexerOutput)(nil),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetIndexers()
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestGetIndexer(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:            "200",
			ExpectedPath:    path.Join("/", starr.API, sonarr.APIver, "indexer", "1"),
			ExpectedRequest: "",
			ExpectedMethod:  "GET",
			ResponseStatus:  200,
			ResponseBody:    indexerResponseBody,
			WithRequest:     nil,
			WithResponse: &sonarr.IndexerOutput{
				EnableAutomaticSearch:   true,
				EnableInteractiveSearch: true,
				EnableRss:               true,
				SupportsRss:             true,
				SupportsSearch:          true,
				Priority:                25,
				ID:                      1,
				ConfigContract:          "NewznabSettings",
				Implementation:          "Newznab",
				ImplementationName:      "Newznab",
				InfoLink:                "https://wiki.servarr.com/sonarr/supported#newznab",
				Name:                    "NZBgeek",
				Protocol:                "usenet",
				Fields: []*starr.FieldOutput{
					{
						Order:    0,
						Name:     "baseUrl",
						Label:    "URL",
						Value:    "https://api.nzbgeek.info",
						Type:     "textbox",
						Advanced: false,
					},
					{
						Order:    1,
						Name:     "apiPath",
						Label:    "API Path",
						HelpText: "Path to the api, usually /api",
						Value:    "/api",
						Type:     "textbox",
						Advanced: true,
					},
				},
				Tags: []int{},
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "indexer", "1"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:   (*sonarr.IndexerOutput)(nil),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetIndexer(1)
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestAddIndexer(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "indexer?forceSave=true"),
			ExpectedMethod: "POST",
			ResponseStatus: 200,
			WithRequest: &sonarr.IndexerInput{
				EnableAutomaticSearch:   true,
				EnableInteractiveSearch: true,
				EnableRss:               true,
				DownloadClientID:        0,
				Priority:                25,
				ConfigContract:          "NewznabSettings",
				Implementation:          "Newznab",
				Name:                    "NZBgeek",
				Protocol:                "usenet",
				Tags:                    []int{},
				Fields: []*starr.FieldInput{
					{
						Name:  "baseUrl",
						Value: "https://api.nzbgeek.info",
					},
					{
						Name:  "apiPath",
						Value: "/api",
					},
				},
			},
			ExpectedRequest: addIndexer + "\n",
			ResponseBody:    indexerResponseBody,
			WithResponse: &sonarr.IndexerOutput{
				EnableAutomaticSearch:   true,
				EnableInteractiveSearch: true,
				EnableRss:               true,
				SupportsRss:             true,
				SupportsSearch:          true,
				Priority:                25,
				ID:                      1,
				ConfigContract:          "NewznabSettings",
				Implementation:          "Newznab",
				ImplementationName:      "Newznab",
				InfoLink:                "https://wiki.servarr.com/sonarr/supported#newznab",
				Name:                    "NZBgeek",
				Protocol:                "usenet",
				Fields: []*starr.FieldOutput{
					{
						Order:    0,
						Name:     "baseUrl",
						Label:    "URL",
						Value:    "https://api.nzbgeek.info",
						Type:     "textbox",
						Advanced: false,
					},
					{
						Order:    1,
						Name:     "apiPath",
						Label:    "API Path",
						HelpText: "Path to the api, usually /api",
						Value:    "/api",
						Type:     "textbox",
						Advanced: true,
					},
				},
				Tags: []int{},
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "indexer?forceSave=true"),
			ExpectedMethod: "POST",
			ResponseStatus: 404,
			WithRequest: &sonarr.IndexerInput{
				EnableAutomaticSearch:   true,
				EnableInteractiveSearch: true,
				EnableRss:               true,
				DownloadClientID:        0,
				Priority:                25,
				ConfigContract:          "NewznabSettings",
				Implementation:          "Newznab",
				Name:                    "NZBgeek",
				Protocol:                "usenet",
				Tags:                    []int{},
				Fields: []*starr.FieldInput{
					{
						Name:  "baseUrl",
						Value: "https://api.nzbgeek.info",
					},
					{
						Name:  "apiPath",
						Value: "/api",
					},
				},
			},
			ExpectedRequest: addIndexer + "\n",
			ResponseBody:    `{"message": "NotFound"}`,
			WithError:       &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:    (*sonarr.IndexerOutput)(nil),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.AddIndexer(test.WithRequest.(*sonarr.IndexerInput))
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestUpdateIndexer(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "indexer", "1?forceSave=false"),
			ExpectedMethod: "PUT",
			ResponseStatus: 200,
			WithRequest: &sonarr.IndexerInput{
				EnableAutomaticSearch:   true,
				EnableInteractiveSearch: true,
				EnableRss:               true,
				DownloadClientID:        0,
				Priority:                25,
				ConfigContract:          "NewznabSettings",
				Implementation:          "Newznab",
				Name:                    "NZBgeek",
				Protocol:                "usenet",
				Tags:                    []int{},
				Fields: []*starr.FieldInput{
					{
						Name:  "baseUrl",
						Value: "https://api.nzbgeek.info",
					},
					{
						Name:  "apiPath",
						Value: "/api",
					},
				},
				ID: 1,
			},
			ExpectedRequest: updateIndexer + "\n",
			ResponseBody:    indexerResponseBody,
			WithResponse: &sonarr.IndexerOutput{
				EnableAutomaticSearch:   true,
				EnableInteractiveSearch: true,
				EnableRss:               true,
				SupportsRss:             true,
				SupportsSearch:          true,
				Priority:                25,
				ID:                      1,
				ConfigContract:          "NewznabSettings",
				Implementation:          "Newznab",
				ImplementationName:      "Newznab",
				InfoLink:                "https://wiki.servarr.com/sonarr/supported#newznab",
				Name:                    "NZBgeek",
				Protocol:                "usenet",
				Fields: []*starr.FieldOutput{
					{
						Order:    0,
						Name:     "baseUrl",
						Label:    "URL",
						Value:    "https://api.nzbgeek.info",
						Type:     "textbox",
						Advanced: false,
					},
					{
						Order:    1,
						Name:     "apiPath",
						Label:    "API Path",
						HelpText: "Path to the api, usually /api",
						Value:    "/api",
						Type:     "textbox",
						Advanced: true,
					},
				},
				Tags: []int{},
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "indexer", "1?forceSave=false"),
			ExpectedMethod: "PUT",
			ResponseStatus: 404,
			WithRequest: &sonarr.IndexerInput{
				EnableAutomaticSearch:   true,
				EnableInteractiveSearch: true,
				EnableRss:               true,
				DownloadClientID:        0,
				Priority:                25,
				ConfigContract:          "NewznabSettings",
				Implementation:          "Newznab",
				Name:                    "NZBgeek",
				Protocol:                "usenet",
				Tags:                    []int{},
				Fields: []*starr.FieldInput{
					{
						Name:  "baseUrl",
						Value: "https://api.nzbgeek.info",
					},
					{
						Name:  "apiPath",
						Value: "/api",
					},
				},
				ID: 1,
			},
			ExpectedRequest: updateIndexer + "\n",
			ResponseBody:    `{"message": "NotFound"}`,
			WithError:       &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:    (*sonarr.IndexerOutput)(nil),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.UpdateIndexer(test.WithRequest.(*sonarr.IndexerInput), false)
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestDeleteIndexer(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "indexer", "2"),
			ExpectedMethod: "DELETE",
			WithRequest:    int64(2),
			ResponseStatus: 200,
			ResponseBody:   "{}",
			WithError:      nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "indexer", "2"),
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
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			err := client.DeleteIndexer(test.WithRequest.(int64))
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
		})
	}
}
