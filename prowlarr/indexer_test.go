package prowlarr_test

import (
	"net/http"
	"path"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golift.io/starr"
	"golift.io/starr/prowlarr"
	"golift.io/starr/starrtest"
)

const (
	indexerResponseBody = `{
    "indexerUrls": [
        "https://torrentapi.org"
    ],
    "legacyUrls": [],
    "definitionName": "Rarbg",
    "description": "RARBG is a Public torrent site for MOVIES / TV / GENERAL",
    "language": "en-US",
    "encoding": "Unicode (UTF-8)",
    "enable": false,
    "redirect": false,
    "supportsRss": true,
    "supportsSearch": true,
    "supportsRedirect": false,
    "appProfileId": 0,
    "protocol": "torrent",
    "privacy": "public",
    "capabilities": {
        "limitsMax": 100,
        "limitsDefault": 100,
        "categories": [
            {
                "id": 2000,
                "name": "Movies",
                "subCategories": [
                    {
                        "id": 2030,
                        "name": "Movies/SD",
                        "subCategories": []
                    }
                ]
            }
        ],
        "supportsRawSearch": false,
        "searchParams": [
            "q"
        ],
        "tvSearchParams": [
            "q",
            "season",
            "ep",
            "imdbId",
            "tvdbId"
        ],
        "movieSearchParams": [
            "q",
            "imdbId",
            "tmdbId"
        ],
        "musicSearchParams": [
            "q"
        ],
        "bookSearchParams": []
    },
    "priority": 25,
    "added": "2019-06-04T01:00:00Z",
    "sortName": "nyaa",
    "name": "Nyaa",
    "fields": [
        {
            "order": 0,
            "name": "baseUrl",
            "label": "Base Url",
            "helpText": "Select which baseurl Prowlarr will use for requests to the site",
            "value": "http://nyaa.si",
            "type": "select",
            "advanced": false,
            "selectOptionsProviderAction": "getUrls"
        },
        {
            "order": 1,
            "name": "rankedOnly",
            "label": "Ranked Only",
            "helpText": "Only include ranked results.",
            "value": false,
            "type": "checkbox",
            "advanced": false
        }
    ],
    "implementationName": "Rarbg",
    "implementation": "Rarbg",
    "configContract": "RarbgSettings",
    "infoLink": "https://wiki.servarr.com/prowlarr/supported-indexers#rarbg",
    "tags": [],
    "id": 2
}`
	addIndexer = `{"enable":false,"redirect":false,"priority":25,"appProfileId":2,` +
		`"configContract":"RarbgSettings","implementation":"Rarbg","name":"Nyaa","protocol":"torrent",` +
		`"fields":[{"name":"baseUrl","value":"http://nyaa.si"},{"name":"animeStandardFormatSearch","value":false}]}`
	updateIndexer = `{"enable":false,"redirect":false,"priority":25,"id":2,"appProfileId":2,` +
		`"configContract":"RarbgSettings","implementation":"Rarbg","name":"Nyaa","protocol":"torrent",` +
		`"fields":[{"name":"baseUrl","value":"http://nyaa.si"},{"name":"animeStandardFormatSearch","value":false}]}`
)

var (
	loc, _          = time.LoadLocation("")
	date            = time.Date(2019, 6, 4, 1, 0, 0, 0, loc)
	indexerResponse = prowlarr.IndexerOutput{
		IndexerUrls:      []string{"https://torrentapi.org"},
		LegacyUrls:       []string{},
		DefinitionName:   "Rarbg",
		Description:      "RARBG is a Public torrent site for MOVIES / TV / GENERAL",
		Language:         "en-US",
		Encoding:         "Unicode (UTF-8)",
		Enable:           false,
		Redirect:         false,
		SupportsRss:      true,
		SupportsSearch:   true,
		SupportsRedirect: false,
		AppProfileID:     0,
		Protocol:         "torrent",
		Privacy:          "public",
		Capabilities: &prowlarr.Capabilities{
			LimitsMax:     100,
			LimitsDefault: 100,
			Categories: []*prowlarr.Categories{
				{
					ID:   2000,
					Name: "Movies",
					SubCategories: []*prowlarr.Categories{
						{
							ID:            2030,
							Name:          "Movies/SD",
							SubCategories: []*prowlarr.Categories{},
						},
					},
				},
			},
			SupportsRawSearch: false,
			SearchParams:      []string{"q"},
			TvSearchParams: []string{
				"q",
				"season",
				"ep",
				"imdbId",
				"tvdbId",
			},
			MovieSearchParams: []string{
				"q",
				"imdbId",
				"tmdbId",
			},
			MusicSearchParams: []string{"q"},
			BookSearchParams:  []string{},
		},
		Priority: 25,
		Added:    date,
		SortName: "nyaa",
		Name:     "Nyaa",
		Fields: []*starr.FieldOutput{
			{
				Order:                       0,
				Name:                        "baseUrl",
				Label:                       "Base Url",
				HelpText:                    "Select which baseurl Prowlarr will use for requests to the site",
				Value:                       "http://nyaa.si",
				Type:                        "select",
				Advanced:                    false,
				SelectOptionsProviderAction: "getUrls",
			},
			{
				Order:    1,
				Name:     "rankedOnly",
				Label:    "Ranked Only",
				HelpText: "Only include ranked results.",
				Value:    false,
				Type:     "checkbox",
				Advanced: false,
			},
		},
		ImplementationName: "Rarbg",
		Implementation:     "Rarbg",
		ConfigContract:     "RarbgSettings",
		InfoLink:           "https://wiki.servarr.com/prowlarr/supported-indexers#rarbg",
		Tags:               []int{},
		ID:                 2,
	}
)

func TestGetIndexers(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:            "200",
			ExpectedPath:    path.Join("/", starr.API, prowlarr.APIver, "indexer"),
			ExpectedRequest: "",
			ExpectedMethod:  "GET",
			ResponseStatus:  200,
			ResponseBody:    "[" + indexerResponseBody + "]",
			WithRequest:     nil,
			WithResponse: []*prowlarr.IndexerOutput{
				&indexerResponse,
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, prowlarr.APIver, "indexer"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:   ([]*prowlarr.IndexerOutput)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := prowlarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
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
			ExpectedPath:    path.Join("/", starr.API, prowlarr.APIver, "indexer", "2"),
			ExpectedRequest: "",
			ExpectedMethod:  "GET",
			ResponseStatus:  200,
			ResponseBody:    indexerResponseBody,
			WithRequest:     int64(2),
			WithResponse:    &indexerResponse,
			WithError:       nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, prowlarr.APIver, "indexer", "2"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithRequest:    int64(2),
			WithError:      &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:   (*prowlarr.IndexerOutput)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := prowlarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetIndexer(test.WithRequest.(int64))
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
			ExpectedPath:   path.Join("/", starr.API, prowlarr.APIver, "indexer?forceSave=true"),
			ExpectedMethod: "POST",
			ResponseStatus: 200,
			WithRequest: &prowlarr.IndexerInput{
				Enable:       false,
				Protocol:     "torrent",
				AppProfileID: 2,
				Priority:     25,
				Name:         "Nyaa",
				Fields: []*starr.FieldInput{
					{
						Name:  "baseUrl",
						Value: "http://nyaa.si",
					},
					{
						Name:  "animeStandardFormatSearch",
						Value: false,
					},
				},
				Implementation: "Rarbg",
				ConfigContract: "RarbgSettings",
			},
			ExpectedRequest: addIndexer + "\n",
			ResponseBody:    indexerResponseBody,
			WithResponse:    &indexerResponse,
			WithError:       nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, prowlarr.APIver, "indexer?forceSave=true"),
			ExpectedMethod: "POST",
			ResponseStatus: 404,
			WithRequest: &prowlarr.IndexerInput{
				Enable:       false,
				Protocol:     "torrent",
				AppProfileID: 2,
				Priority:     25,
				Name:         "Nyaa",
				Fields: []*starr.FieldInput{
					{
						Name:  "baseUrl",
						Value: "http://nyaa.si",
					},
					{
						Name:  "animeStandardFormatSearch",
						Value: false,
					},
				},
				Implementation: "Rarbg",
				ConfigContract: "RarbgSettings",
			},
			ExpectedRequest: addIndexer + "\n",
			ResponseBody:    `{"message": "NotFound"}`,
			WithError:       &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:    (*prowlarr.IndexerOutput)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := prowlarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.AddIndexer(test.WithRequest.(*prowlarr.IndexerInput))
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
			ExpectedPath:   path.Join("/", starr.API, prowlarr.APIver, "indexer", "2?forceSave=false"),
			ExpectedMethod: "PUT",
			ResponseStatus: 200,
			WithRequest: &prowlarr.IndexerInput{
				Enable:       false,
				Protocol:     "torrent",
				AppProfileID: 2,
				Priority:     25,
				Name:         "Nyaa",
				Fields: []*starr.FieldInput{
					{
						Name:  "baseUrl",
						Value: "http://nyaa.si",
					},
					{
						Name:  "animeStandardFormatSearch",
						Value: false,
					},
				},
				Implementation: "Rarbg",
				ConfigContract: "RarbgSettings",
				ID:             2,
			},
			ExpectedRequest: updateIndexer + "\n",
			ResponseBody:    indexerResponseBody,
			WithResponse:    &indexerResponse,
			WithError:       nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, prowlarr.APIver, "indexer", "2?forceSave=false"),
			ExpectedMethod: "PUT",
			ResponseStatus: 404,
			WithRequest: &prowlarr.IndexerInput{
				Enable:       false,
				Protocol:     "torrent",
				AppProfileID: 2,
				Priority:     25,
				Name:         "Nyaa",
				Fields: []*starr.FieldInput{
					{
						Name:  "baseUrl",
						Value: "http://nyaa.si",
					},
					{
						Name:  "animeStandardFormatSearch",
						Value: false,
					},
				},
				Implementation: "Rarbg",
				ConfigContract: "RarbgSettings",
				ID:             2,
			},
			ExpectedRequest: updateIndexer + "\n",
			ResponseBody:    `{"message": "NotFound"}`,
			WithError:       &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:    (*prowlarr.IndexerOutput)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := prowlarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.UpdateIndexer(test.WithRequest.(*prowlarr.IndexerInput), false)
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
			ExpectedPath:   path.Join("/", starr.API, prowlarr.APIver, "indexer", "2"),
			ExpectedMethod: "DELETE",
			WithRequest:    int64(2),
			ResponseStatus: 200,
			ResponseBody:   "{}",
			WithError:      nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, prowlarr.APIver, "indexer", "2"),
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
			client := prowlarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			err := client.DeleteIndexer(test.WithRequest.(int64))
			require.ErrorIs(t, err, test.WithError, "error is not the same as expected")
		})
	}
}
