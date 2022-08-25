package sonarr_test

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"golift.io/starr"
	"golift.io/starr/sonarr"
)

const indexerBody = `{
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

func TestGet(t *testing.T) {
	t.Parallel()

	tests := []*starr.TestMockData{
		{
			Name:            "200",
			ExpectedPath:    path.Join("/", starr.API, sonarr.APIver, "indexer/1"),
			ExpectedRequest: "",
			ExpectedMethod:  "GET",
			ResponseStatus:  200,
			ResponseBody:    indexerBody,
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
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "indexer/1"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      starr.ErrInvalidStatusCode,
			WithResponse:   (*sonarr.IndexerOutput)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetIndexer(1)
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}
