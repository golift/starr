package sonarr_test

import (
	"net/http"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"golift.io/starr"
	"golift.io/starr/sonarr"
	"golift.io/starr/starrtest"
)

const importListResponseBody = `{
	"enableAutomaticAdd": false,
	"shouldMonitor": "all",
	"rootFolderPath": "/config",
	"qualityProfileId": 1,
	"seriesType": "standard",
	"seasonFolder": true,
	"listType": "plex",
	"listOrder": 1,
	"name": "PlexImport",
	"fields": [
	  {
		"order": 0,
		"name": "accessToken",
		"label": "Access Token",
		"type": "textbox",
		"value": "test",
		"advanced": false,
		"hidden": "hidden"
	  },
	  {
		"order": 1,
		"name": "signIn",
		"label": "Authenticate with Plex.tv",
		"value": "startOAuth",
		"type": "oAuth",
		"advanced": false
	  }
	],
	"implementationName": "Plex Watchlist",
	"implementation": "PlexImport",
	"configContract": "PlexListSettings",
	"infoLink": "https://wiki.servarr.com/sonarr/supported#pleximport",
	"tags": [],
	"id": 4
  }`

const addImportList = `{"enableAutomaticAdd":false,"seasonFolder":true,"listOrder":0,"qualityProfileId":1,"configContract":"PlexListSettings","implementation":"PlexImport","name":"PlexImport","rootFolderPath":"/config","seriesType":"standard","shouldMonitor":"all","fields":[{"name":"accessToken","value":"test"}]}`

const updateImportList = `{"enableAutomaticAdd":false,"seasonFolder":true,"listOrder":0,"qualityProfileId":1,"id":4,"configContract":"PlexListSettings","implementation":"PlexImport","name":"PlexImport","rootFolderPath":"/config","seriesType":"standard","shouldMonitor":"all","fields":[{"name":"accessToken","value":"test"}]}`

func TestGetImportLists(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:            "200",
			ExpectedPath:    path.Join("/", starr.API, sonarr.APIver, "importList"),
			ExpectedRequest: "",
			ExpectedMethod:  "GET",
			ResponseStatus:  200,
			ResponseBody:    "[" + importListResponseBody + "]",
			WithRequest:     nil,
			WithResponse: []*sonarr.ImportListOutput{
				{
					EnableAutomaticAdd: false,
					ShouldMonitor:      "all",
					RootFolderPath:     "/config",
					QualityProfileID:   1,
					SeriesType:         "standard",
					SeasonFolder:       true,
					ListType:           "plex",
					ListOrder:          1,
					Name:               "PlexImport",
					Fields: []*starr.FieldOutput{
						{
							Order:    0,
							Name:     "accessToken",
							Label:    "Access Token",
							Type:     "textbox",
							Value:    "test",
							Advanced: false,
							Hidden:   "hidden",
						},
						{
							Order:    1,
							Name:     "signIn",
							Label:    "Authenticate with Plex.tv",
							Value:    "startOAuth",
							Type:     "oAuth",
							Advanced: false,
						},
					},
					ImplementationName: "Plex Watchlist",
					Implementation:     "PlexImport",
					ConfigContract:     "PlexListSettings",
					InfoLink:           "https://wiki.servarr.com/sonarr/supported#pleximport",
					Tags:               []int{},
					ID:                 4,
				},
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "importList"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:   ([]*sonarr.ImportListOutput)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetImportLists()
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestGetImportList(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:            "200",
			ExpectedPath:    path.Join("/", starr.API, sonarr.APIver, "importList", "1"),
			ExpectedRequest: "",
			ExpectedMethod:  "GET",
			ResponseStatus:  200,
			ResponseBody:    importListResponseBody,
			WithRequest:     nil,
			WithResponse: &sonarr.ImportListOutput{
				EnableAutomaticAdd: false,
				ShouldMonitor:      "all",
				RootFolderPath:     "/config",
				QualityProfileID:   1,
				SeriesType:         "standard",
				SeasonFolder:       true,
				ListType:           "plex",
				ListOrder:          1,
				Name:               "PlexImport",
				Fields: []*starr.FieldOutput{
					{
						Order:    0,
						Name:     "accessToken",
						Label:    "Access Token",
						Type:     "textbox",
						Value:    "test",
						Advanced: false,
						Hidden:   "hidden",
					},
					{
						Order:    1,
						Name:     "signIn",
						Label:    "Authenticate with Plex.tv",
						Value:    "startOAuth",
						Type:     "oAuth",
						Advanced: false,
					},
				},
				ImplementationName: "Plex Watchlist",
				Implementation:     "PlexImport",
				ConfigContract:     "PlexListSettings",
				InfoLink:           "https://wiki.servarr.com/sonarr/supported#pleximport",
				Tags:               []int{},
				ID:                 4,
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "importList", "1"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:   (*sonarr.ImportListOutput)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetImportList(1)
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestAddImportList(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "importList?forceSave=true"),
			ExpectedMethod: "POST",
			ResponseStatus: 200,
			WithRequest: &sonarr.ImportListInput{
				EnableAutomaticAdd: false,
				ShouldMonitor:      "all",
				RootFolderPath:     "/config",
				QualityProfileID:   1,
				SeriesType:         "standard",
				SeasonFolder:       true,
				Name:               "PlexImport",
				Fields: []*starr.FieldInput{
					{
						Name:  "accessToken",
						Value: "test",
					},
				},
				Implementation: "PlexImport",
				ConfigContract: "PlexListSettings",
				Tags:           []int{},
			},
			ExpectedRequest: addImportList + "\n",
			ResponseBody:    importListResponseBody,
			WithResponse: &sonarr.ImportListOutput{
				EnableAutomaticAdd: false,
				ShouldMonitor:      "all",
				RootFolderPath:     "/config",
				QualityProfileID:   1,
				SeriesType:         "standard",
				SeasonFolder:       true,
				ListType:           "plex",
				ListOrder:          1,
				Name:               "PlexImport",
				Fields: []*starr.FieldOutput{
					{
						Order:    0,
						Name:     "accessToken",
						Label:    "Access Token",
						Type:     "textbox",
						Value:    "test",
						Advanced: false,
						Hidden:   "hidden",
					},
					{
						Order:    1,
						Name:     "signIn",
						Label:    "Authenticate with Plex.tv",
						Value:    "startOAuth",
						Type:     "oAuth",
						Advanced: false,
					},
				},
				ImplementationName: "Plex Watchlist",
				Implementation:     "PlexImport",
				ConfigContract:     "PlexListSettings",
				InfoLink:           "https://wiki.servarr.com/sonarr/supported#pleximport",
				Tags:               []int{},
				ID:                 4,
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "importList?forceSave=true"),
			ExpectedMethod: "POST",
			ResponseStatus: 404,
			WithRequest: &sonarr.ImportListInput{
				EnableAutomaticAdd: false,
				ShouldMonitor:      "all",
				RootFolderPath:     "/config",
				QualityProfileID:   1,
				SeriesType:         "standard",
				SeasonFolder:       true,
				Name:               "PlexImport",
				Fields: []*starr.FieldInput{
					{
						Name:  "accessToken",
						Value: "test",
					},
				},
				Implementation: "PlexImport",
				ConfigContract: "PlexListSettings",
				Tags:           []int{},
			},
			ExpectedRequest: addImportList + "\n",
			ResponseBody:    `{"message": "NotFound"}`,
			WithError:       &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:    (*sonarr.ImportListOutput)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.AddImportList(test.WithRequest.(*sonarr.ImportListInput))
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestUpdateImportList(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "importList", "4?forceSave=false"),
			ExpectedMethod: "PUT",
			ResponseStatus: 200,
			WithRequest: &sonarr.ImportListInput{
				EnableAutomaticAdd: false,
				ShouldMonitor:      "all",
				RootFolderPath:     "/config",
				QualityProfileID:   1,
				SeriesType:         "standard",
				SeasonFolder:       true,
				Name:               "PlexImport",
				Fields: []*starr.FieldInput{
					{
						Name:  "accessToken",
						Value: "test",
					},
				},
				Implementation: "PlexImport",
				ConfigContract: "PlexListSettings",
				Tags:           []int{},
				ID:             4,
			},
			ExpectedRequest: updateImportList + "\n",
			ResponseBody:    importListResponseBody,
			WithResponse: &sonarr.ImportListOutput{
				EnableAutomaticAdd: false,
				ShouldMonitor:      "all",
				RootFolderPath:     "/config",
				QualityProfileID:   1,
				SeriesType:         "standard",
				SeasonFolder:       true,
				ListType:           "plex",
				ListOrder:          1,
				Name:               "PlexImport",
				Fields: []*starr.FieldOutput{
					{
						Order:    0,
						Name:     "accessToken",
						Label:    "Access Token",
						Type:     "textbox",
						Value:    "test",
						Advanced: false,
						Hidden:   "hidden",
					},
					{
						Order:    1,
						Name:     "signIn",
						Label:    "Authenticate with Plex.tv",
						Value:    "startOAuth",
						Type:     "oAuth",
						Advanced: false,
					},
				},
				ImplementationName: "Plex Watchlist",
				Implementation:     "PlexImport",
				ConfigContract:     "PlexListSettings",
				InfoLink:           "https://wiki.servarr.com/sonarr/supported#pleximport",
				Tags:               []int{},
				ID:                 4,
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "importList", "4?forceSave=false"),
			ExpectedMethod: "PUT",
			ResponseStatus: 404,
			WithRequest: &sonarr.ImportListInput{
				EnableAutomaticAdd: false,
				ShouldMonitor:      "all",
				RootFolderPath:     "/config",
				QualityProfileID:   1,
				SeriesType:         "standard",
				SeasonFolder:       true,
				Name:               "PlexImport",
				Fields: []*starr.FieldInput{
					{
						Name:  "accessToken",
						Value: "test",
					},
				},
				Implementation: "PlexImport",
				ConfigContract: "PlexListSettings",
				Tags:           []int{},
				ID:             4,
			},
			ExpectedRequest: updateImportList + "\n",
			ResponseBody:    `{"message": "NotFound"}`,
			WithError:       &starr.ReqError{Code: http.StatusNotFound},
			WithResponse:    (*sonarr.ImportListOutput)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.UpdateImportList(test.WithRequest.(*sonarr.ImportListInput), false)
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestDeleteImportList(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "importList", "2"),
			ExpectedMethod: "DELETE",
			WithRequest:    int64(2),
			ResponseStatus: 200,
			ResponseBody:   "{}",
			WithError:      nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, sonarr.APIver, "importList", "2"),
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
			client := sonarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			err := client.DeleteImportList(test.WithRequest.(int64))
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
		})
	}
}
