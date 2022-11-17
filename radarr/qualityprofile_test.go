package radarr_test

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"golift.io/starr"
	"golift.io/starr/radarr"
)

const (
	qualityProfileResponse = `{
		"name": "test",
		"upgradeAllowed": false,
		"cutoff": 1003,
		"items": [
		  {
			"name": "WEB 2160p",
			"items": [
				{
					"quality": {
					  "id": 18,
					  "name": "WEBDL-2160p",
					  "source": "webdl",
					  "resolution": 2160,
					  "modifier": "none"
					},
					"allowed": true
				  },
				  {
					"quality": {
					  "id": 17,
					  "name": "WEBRip-2160p",
					  "source": "webrip",
					  "resolution": 2160,
					  "modifier": "none"
					},
					"allowed": true
				}
			],
			"allowed": true,
			"id": 1003
		  }
		],
		"minFormatScore": 0,
		"cutoffFormatScore": 0,
		"formatItems": [],
		"language": {
		  "id": 1,
		  "name": "English"
		},
		"id": 7
	  }`

	addQualityProfileRequest = `{"name":"test","upgradeAllowed":false,"cutoff":1003,"items":[{"name":"WEB 2160p",` +
		`"id":1003,"items":[{"quality":{"id":18,"name":"WEBDL-2160p","source":"webdl","resolution":2160,"modifier":"none"},` +
		`"allowed":true},{"quality":{"id":17,"name":"WEBRip-2160p","source":"webrip","resolution":2160,"modifier":"none"},` +
		`"allowed":true}],"allowed":true}],"minFormatScore":0,"cutoffFormatScore":0,"formatItems":null,` +
		`"language":{"id":1,"name":"English"}}` + "\n"
	updateQualityProfileRequest = `{"id":7,"name":"test","upgradeAllowed":false,"cutoff":1003,"items":` +
		`[{"name":"WEB 2160p","id":1003,"items":[{"quality":{"id":18,"name":"WEBDL-2160p","source":"webdl",` +
		`"resolution":2160,"modifier":"none"},"allowed":true},{"quality":{"id":17,"name":"WEBRip-2160p","source":"webrip",` +
		`"resolution":2160,"modifier":"none"},"allowed":true}],"allowed":true}],"minFormatScore":0,` +
		`"cutoffFormatScore":0,"formatItems":null,"language":{"id":1,"name":"English"}}` + "\n"
)

func TestGetQualityProfiles(t *testing.T) {
	t.Parallel()

	tests := []*starr.TestMockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "qualityProfile"),
			ExpectedMethod: "GET",
			ResponseStatus: 200,
			ResponseBody:   `[` + qualityProfileResponse + `]`,
			WithResponse: []*radarr.QualityProfile{
				{
					ID:             7,
					Name:           "test",
					UpgradeAllowed: false,
					Cutoff:         1003,
					FormatItems:    []*starr.FormatItem{},
					Qualities: []*starr.Quality{
						{
							Name: "WEB 2160p",
							ID:   1003,
							Items: []*starr.Quality{
								{
									Allowed: true,
									Quality: &starr.BaseQuality{
										ID:         18,
										Name:       "WEBDL-2160p",
										Source:     "webdl",
										Resolution: 2160,
										Modifier:   "none",
									},
								},
								{
									Allowed: true,
									Quality: &starr.BaseQuality{
										ID:         17,
										Name:       "WEBRip-2160p",
										Source:     "webrip",
										Resolution: 2160,
										Modifier:   "none",
									},
								},
							},
							Allowed: true,
						},
					},
					MinFormatScore:    0,
					CutoffFormatScore: 0,
					Language: &starr.Value{
						ID:   1,
						Name: "English",
					},
				},
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "qualityProfile"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      starr.ErrInvalidStatusCode,
			WithResponse:   []*radarr.QualityProfile(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := radarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetQualityProfiles()
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestGetQualityProfile(t *testing.T) {
	t.Parallel()

	tests := []*starr.TestMockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "qualityProfile", "7"),
			ExpectedMethod: "GET",
			ResponseStatus: 200,
			WithRequest:    int64(7),
			ResponseBody:   qualityProfileResponse,
			WithResponse: &radarr.QualityProfile{
				ID:             7,
				Name:           "test",
				UpgradeAllowed: false,
				Cutoff:         1003,
				FormatItems:    []*starr.FormatItem{},
				Qualities: []*starr.Quality{
					{
						Name: "WEB 2160p",
						ID:   1003,
						Items: []*starr.Quality{
							{
								Allowed: true,
								Quality: &starr.BaseQuality{
									ID:         18,
									Name:       "WEBDL-2160p",
									Source:     "webdl",
									Resolution: 2160,
									Modifier:   "none",
								},
							},
							{
								Allowed: true,
								Quality: &starr.BaseQuality{
									ID:         17,
									Name:       "WEBRip-2160p",
									Source:     "webrip",
									Resolution: 2160,
									Modifier:   "none",
								},
							},
						},
						Allowed: true,
					},
				},
				MinFormatScore:    0,
				CutoffFormatScore: 0,
				Language: &starr.Value{
					ID:   1,
					Name: "English",
				},
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "qualityProfile", "1"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			WithRequest:    int64(1),
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      starr.ErrInvalidStatusCode,
			WithResponse:   (*radarr.QualityProfile)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := radarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetQualityProfile(test.WithRequest.(int64))
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestAddQualityProfile(t *testing.T) {
	t.Parallel()

	tests := []*starr.TestMockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "qualityProfile"),
			ExpectedMethod: "POST",
			ResponseStatus: 200,
			WithRequest: &radarr.QualityProfile{
				Name:   "test",
				Cutoff: 1003,
				Qualities: []*starr.Quality{
					{
						Name: "WEB 2160p",
						ID:   1003,
						Items: []*starr.Quality{
							{
								Allowed: true,
								Quality: &starr.BaseQuality{
									ID:         18,
									Name:       "WEBDL-2160p",
									Source:     "webdl",
									Resolution: 2160,
									Modifier:   "none",
								},
							},
							{
								Allowed: true,
								Quality: &starr.BaseQuality{
									ID:         17,
									Name:       "WEBRip-2160p",
									Source:     "webrip",
									Resolution: 2160,
									Modifier:   "none",
								},
							},
						},
						Allowed: true,
					},
				},
				MinFormatScore:    0,
				CutoffFormatScore: 0,
				Language: &starr.Value{
					ID:   1,
					Name: "English",
				},
			},
			ExpectedRequest: addQualityProfileRequest,
			ResponseBody:    qualityProfileResponse,
			WithResponse: &radarr.QualityProfile{
				ID:             7,
				Name:           "test",
				UpgradeAllowed: false,
				Cutoff:         1003,
				FormatItems:    []*starr.FormatItem{},
				Qualities: []*starr.Quality{
					{
						Name: "WEB 2160p",
						ID:   1003,
						Items: []*starr.Quality{
							{
								Allowed: true,
								Quality: &starr.BaseQuality{
									ID:         18,
									Name:       "WEBDL-2160p",
									Source:     "webdl",
									Resolution: 2160,
									Modifier:   "none",
								},
							},
							{
								Allowed: true,
								Quality: &starr.BaseQuality{
									ID:         17,
									Name:       "WEBRip-2160p",
									Source:     "webrip",
									Resolution: 2160,
									Modifier:   "none",
								},
							},
						},
						Allowed: true,
					},
				},
				MinFormatScore:    0,
				CutoffFormatScore: 0,
				Language: &starr.Value{
					ID:   1,
					Name: "English",
				},
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "qualityProfile"),
			ExpectedMethod: "POST",
			WithRequest: &radarr.QualityProfile{
				Name:   "test",
				Cutoff: 1003,
				Qualities: []*starr.Quality{
					{
						Name: "WEB 2160p",
						ID:   1003,
						Items: []*starr.Quality{
							{
								Allowed: true,
								Quality: &starr.BaseQuality{
									ID:         18,
									Name:       "WEBDL-2160p",
									Source:     "webdl",
									Resolution: 2160,
									Modifier:   "none",
								},
							},
							{
								Allowed: true,
								Quality: &starr.BaseQuality{
									ID:         17,
									Name:       "WEBRip-2160p",
									Source:     "webrip",
									Resolution: 2160,
									Modifier:   "none",
								},
							},
						},
						Allowed: true,
					},
				},
				MinFormatScore:    0,
				CutoffFormatScore: 0,
				Language: &starr.Value{
					ID:   1,
					Name: "English",
				},
			},
			ExpectedRequest: addQualityProfileRequest,
			ResponseStatus:  404,
			ResponseBody:    `{"message": "NotFound"}`,
			WithError:       starr.ErrInvalidStatusCode,
			WithResponse:    (*radarr.QualityProfile)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := radarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.AddQualityProfile(test.WithRequest.(*radarr.QualityProfile))
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestUpdateQualityProfile(t *testing.T) {
	t.Parallel()

	tests := []*starr.TestMockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "qualityProfile", "7"),
			ExpectedMethod: "PUT",
			ResponseStatus: 200,
			WithRequest: &radarr.QualityProfile{
				Name:   "test",
				Cutoff: 1003,
				Qualities: []*starr.Quality{
					{
						Name: "WEB 2160p",
						ID:   1003,
						Items: []*starr.Quality{
							{
								Allowed: true,
								Quality: &starr.BaseQuality{
									ID:         18,
									Name:       "WEBDL-2160p",
									Source:     "webdl",
									Resolution: 2160,
									Modifier:   "none",
								},
							},
							{
								Allowed: true,
								Quality: &starr.BaseQuality{
									ID:         17,
									Name:       "WEBRip-2160p",
									Source:     "webrip",
									Resolution: 2160,
									Modifier:   "none",
								},
							},
						},
						Allowed: true,
					},
				},
				MinFormatScore:    0,
				CutoffFormatScore: 0,
				Language: &starr.Value{
					ID:   1,
					Name: "English",
				},
				ID: 7,
			},
			ExpectedRequest: updateQualityProfileRequest,
			ResponseBody:    qualityProfileResponse,
			WithResponse: &radarr.QualityProfile{
				ID:             7,
				Name:           "test",
				UpgradeAllowed: false,
				Cutoff:         1003,
				FormatItems:    []*starr.FormatItem{},
				Qualities: []*starr.Quality{
					{
						Name: "WEB 2160p",
						ID:   1003,
						Items: []*starr.Quality{
							{
								Allowed: true,
								Quality: &starr.BaseQuality{
									ID:         18,
									Name:       "WEBDL-2160p",
									Source:     "webdl",
									Resolution: 2160,
									Modifier:   "none",
								},
							},
							{
								Allowed: true,
								Quality: &starr.BaseQuality{
									ID:         17,
									Name:       "WEBRip-2160p",
									Source:     "webrip",
									Resolution: 2160,
									Modifier:   "none",
								},
							},
						},
						Allowed: true,
					},
				},
				MinFormatScore:    0,
				CutoffFormatScore: 0,
				Language: &starr.Value{
					ID:   1,
					Name: "English",
				},
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "qualityProfile", "7"),
			ExpectedMethod: "PUT",
			WithRequest: &radarr.QualityProfile{
				Name:   "test",
				Cutoff: 1003,
				Qualities: []*starr.Quality{
					{
						Name: "WEB 2160p",
						ID:   1003,
						Items: []*starr.Quality{
							{
								Allowed: true,
								Quality: &starr.BaseQuality{
									ID:         18,
									Name:       "WEBDL-2160p",
									Source:     "webdl",
									Resolution: 2160,
									Modifier:   "none",
								},
							},
							{
								Allowed: true,
								Quality: &starr.BaseQuality{
									ID:         17,
									Name:       "WEBRip-2160p",
									Source:     "webrip",
									Resolution: 2160,
									Modifier:   "none",
								},
							},
						},
						Allowed: true,
					},
				},
				MinFormatScore:    0,
				CutoffFormatScore: 0,
				Language: &starr.Value{
					ID:   1,
					Name: "English",
				},
				ID: 7,
			},
			ExpectedRequest: updateQualityProfileRequest,
			ResponseStatus:  404,
			ResponseBody:    `{"message": "NotFound"}`,
			WithError:       starr.ErrInvalidStatusCode,
			WithResponse:    (*radarr.QualityProfile)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := radarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.UpdateQualityProfile(test.WithRequest.(*radarr.QualityProfile))
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestDeleteQualityProfile(t *testing.T) {
	t.Parallel()

	tests := []*starr.TestMockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "qualityProfile", "10"),
			ExpectedMethod: "DELETE",
			ResponseStatus: 200,
			WithRequest:    int64(10),
			ResponseBody:   "{}",
			WithError:      nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "qualityProfile", "10"),
			ExpectedMethod: "DELETE",
			WithRequest:    int64(10),
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      starr.ErrInvalidStatusCode,
			WithResponse:   (*radarr.QualityProfile)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := radarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			err := client.DeleteQualityProfile(test.WithRequest.(int64))
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
		})
	}
}
