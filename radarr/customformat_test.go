package radarr_test

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"golift.io/starr"
	"golift.io/starr/radarr"
	"golift.io/starr/starrtest"
)

const customFormatResponseBody = `{
    "id": 1,
    "name": "test",
    "includeCustomFormatWhenRenaming": false,
    "specifications": [
        {
            "name": "Surround Sound",
            "implementation": "ReleaseTitleSpecification",
            "implementationName": "Release Title",
            "infoLink": "https://wiki.servarr.com/radarr/settings#custom-formats-2",
            "negate": false,
            "required": false,
            "fields": [
                {
                    "order": 0,
                    "name": "value",
                    "label": "Regular Expression",
                    "helpText": "Custom Format RegEx is Case Insensitive",
                    "value": "DTS.?(HD|ES|X(?!\\D))|TRUEHD|ATMOS|DD(\\+|P).?([5-9])|EAC3.?([5-9])",
                    "type": "textbox",
                    "advanced": false
                }
            ]
        },
        {
            "name": "Arabic",
            "implementation": "LanguageSpecification",
            "implementationName": "Language",
            "infoLink": "https://wiki.servarr.com/radarr/settings#custom-formats-2",
            "negate": false,
            "required": false,
            "fields": [
                {
                    "order": 0,
                    "name": "value",
                    "label": "Language",
                    "value": 31,
                    "type": "select",
                    "advanced": false,
                    "selectOptions": [
                        {
                            "value": 0,
                            "name": "Unknown",
                            "order": 0,
                            "dividerAfter": true
                        },
                        {
                            "value": 31,
                            "name": "Arabic",
                            "order": 0,
                            "dividerAfter": false
                        }
                    ]
                }
            ]
        }
    ]
}`

const addCustomFormat = `{"name":"test","includeCustomFormatWhenRenaming":false,"specifications":` +
	`[{"name":"Surround Sound","implementation":"ReleaseTitleSpecification","negate":false,"required":false,"fields":` +
	`[{"name":"value","value":"DTS.?(HD|ES|X(?!\\D))|TRUEHD|ATMOS|DD(\\+|P).?([5-9])|EAC3.?([5-9])"}]},{"name":"Arabic",` +
	`"implementation":"LanguageSpecification","negate":false,"required":false,"fields":[{"name":"value","value":31}]}]}`

const updateCustomFormat = `{"id":1,"name":"test","includeCustomFormatWhenRenaming":false,"specifications":` +
	`[{"name":"Surround Sound","implementation":"ReleaseTitleSpecification","negate":false,"required":false,"fields":` +
	`[{"name":"value","value":"DTS.?(HD|ES|X(?!\\D))|TRUEHD|ATMOS|DD(\\+|P).?([5-9])|EAC3.?([5-9])"}]},{"name":"Arabic",` +
	`"implementation":"LanguageSpecification","negate":false,"required":false,"fields":[{"name":"value","value":31}]}]}`

func TestGetCustomFormats(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:            "200",
			ExpectedPath:    path.Join("/", starr.API, radarr.APIver, "customFormat"),
			ExpectedRequest: "",
			ExpectedMethod:  "GET",
			ResponseStatus:  200,
			ResponseBody:    "[" + customFormatResponseBody + "]",
			WithRequest:     nil,
			WithResponse: []*radarr.CustomFormatOutput{
				{
					ID:                    1,
					Name:                  "test",
					IncludeCFWhenRenaming: false,
					Specifications: []*radarr.CustomFormatOutputSpec{
						{
							Name:               "Surround Sound",
							Implementation:     "ReleaseTitleSpecification",
							ImplementationName: "Release Title",
							InfoLink:           "https://wiki.servarr.com/radarr/settings#custom-formats-2",
							Negate:             false,
							Required:           false,
							Fields: []*starr.FieldOutput{
								{
									Order:    0,
									Name:     "value",
									Label:    "Regular Expression",
									HelpText: "Custom Format RegEx is Case Insensitive",
									Value:    "DTS.?(HD|ES|X(?!\\D))|TRUEHD|ATMOS|DD(\\+|P).?([5-9])|EAC3.?([5-9])",
									Type:     "textbox",
									Advanced: false,
								},
							},
						},
						{
							Name:               "Arabic",
							Implementation:     "LanguageSpecification",
							ImplementationName: "Language",
							InfoLink:           "https://wiki.servarr.com/radarr/settings#custom-formats-2",
							Negate:             false,
							Required:           false,
							Fields: []*starr.FieldOutput{
								{
									Order: 0,
									Name:  "value",
									Label: "Language",
									// float because of unmarshal.
									Value:    float64(31),
									Type:     "select",
									Advanced: false,
									SelectOptions: []*starr.SelectOption{
										{
											Value:        0,
											Name:         "Unknown",
											Order:        0,
											DividerAfter: true,
										},
										{
											Value:        31,
											Name:         "Arabic",
											Order:        0,
											DividerAfter: false,
										},
									},
								},
							},
						},
					},
				},
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "customFormat"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      starr.ErrInvalidStatusCode,
			WithResponse:   ([]*radarr.CustomFormatOutput)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := radarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetCustomFormats()
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestGetCustomFormat(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:            "200",
			ExpectedPath:    path.Join("/", starr.API, radarr.APIver, "customFormat", "1"),
			ExpectedRequest: "",
			ExpectedMethod:  "GET",
			ResponseStatus:  200,
			ResponseBody:    customFormatResponseBody,
			WithRequest:     nil,
			WithResponse: &radarr.CustomFormatOutput{
				ID:                    1,
				Name:                  "test",
				IncludeCFWhenRenaming: false,
				Specifications: []*radarr.CustomFormatOutputSpec{
					{
						Name:               "Surround Sound",
						Implementation:     "ReleaseTitleSpecification",
						ImplementationName: "Release Title",
						InfoLink:           "https://wiki.servarr.com/radarr/settings#custom-formats-2",
						Negate:             false,
						Required:           false,
						Fields: []*starr.FieldOutput{
							{
								Order:    0,
								Name:     "value",
								Label:    "Regular Expression",
								HelpText: "Custom Format RegEx is Case Insensitive",
								Value:    "DTS.?(HD|ES|X(?!\\D))|TRUEHD|ATMOS|DD(\\+|P).?([5-9])|EAC3.?([5-9])",
								Type:     "textbox",
								Advanced: false,
							},
						},
					},
					{
						Name:               "Arabic",
						Implementation:     "LanguageSpecification",
						ImplementationName: "Language",
						InfoLink:           "https://wiki.servarr.com/radarr/settings#custom-formats-2",
						Negate:             false,
						Required:           false,
						Fields: []*starr.FieldOutput{
							{
								Order: 0,
								Name:  "value",
								Label: "Language",
								// float because of unmarshal.
								Value:    float64(31),
								Type:     "select",
								Advanced: false,
								SelectOptions: []*starr.SelectOption{
									{
										Value:        0,
										Name:         "Unknown",
										Order:        0,
										DividerAfter: true,
									},
									{
										Value:        31,
										Name:         "Arabic",
										Order:        0,
										DividerAfter: false,
									},
								},
							},
						},
					},
				},
			},
			WithError: nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "customFormat", "1"),
			ExpectedMethod: "GET",
			ResponseStatus: 404,
			ResponseBody:   `{"message": "NotFound"}`,
			WithError:      starr.ErrInvalidStatusCode,
			WithResponse:   (*radarr.CustomFormatOutput)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := radarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.GetCustomFormat(1)
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestAddCustomFormat(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "customFormat"),
			ExpectedMethod: "POST",
			ResponseStatus: 200,
			WithRequest: &radarr.CustomFormatInput{
				IncludeCFWhenRenaming: false,
				Name:                  "test",
				Specifications: []*radarr.CustomFormatInputSpec{
					{
						Name:           "Surround Sound",
						Implementation: "ReleaseTitleSpecification",
						Negate:         false,
						Required:       false,
						Fields: []*starr.FieldInput{
							{
								Name:  "value",
								Value: "DTS.?(HD|ES|X(?!\\D))|TRUEHD|ATMOS|DD(\\+|P).?([5-9])|EAC3.?([5-9])",
							},
						},
					},
					{
						Implementation: "LanguageSpecification",
						Negate:         false,
						Required:       false,
						Fields: []*starr.FieldInput{
							{
								Name:  "value",
								Value: 31,
							},
						},
						Name: "Arabic",
					},
				},
			},
			ExpectedRequest: addCustomFormat + "\n",
			ResponseBody:    customFormatResponseBody,
			WithResponse: &radarr.CustomFormatOutput{
				ID:                    1,
				Name:                  "test",
				IncludeCFWhenRenaming: false,
				Specifications: []*radarr.CustomFormatOutputSpec{
					{
						Name:               "Surround Sound",
						Implementation:     "ReleaseTitleSpecification",
						ImplementationName: "Release Title",
						InfoLink:           "https://wiki.servarr.com/radarr/settings#custom-formats-2",
						Negate:             false,
						Required:           false,
						Fields: []*starr.FieldOutput{
							{
								Order:    0,
								Name:     "value",
								Label:    "Regular Expression",
								HelpText: "Custom Format RegEx is Case Insensitive",
								Value:    "DTS.?(HD|ES|X(?!\\D))|TRUEHD|ATMOS|DD(\\+|P).?([5-9])|EAC3.?([5-9])",
								Type:     "textbox",
								Advanced: false,
							},
						},
					},
					{
						Name:               "Arabic",
						Implementation:     "LanguageSpecification",
						ImplementationName: "Language",
						InfoLink:           "https://wiki.servarr.com/radarr/settings#custom-formats-2",
						Negate:             false,
						Required:           false,
						Fields: []*starr.FieldOutput{
							{
								Order: 0,
								Name:  "value",
								Label: "Language",
								// float because of unmarshal.
								Value:    float64(31),
								Type:     "select",
								Advanced: false,
								SelectOptions: []*starr.SelectOption{
									{
										Value:        0,
										Name:         "Unknown",
										Order:        0,
										DividerAfter: true,
									},
									{
										Value:        31,
										Name:         "Arabic",
										Order:        0,
										DividerAfter: false,
									},
								},
							},
						},
					},
				},
			},
			WithError: nil,
		},
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "customFormat"),
			ExpectedMethod: "POST",
			ResponseStatus: 404,
			WithRequest: &radarr.CustomFormatInput{
				IncludeCFWhenRenaming: false,
				Name:                  "test",
				Specifications: []*radarr.CustomFormatInputSpec{
					{
						Name:           "Surround Sound",
						Implementation: "ReleaseTitleSpecification",
						Negate:         false,
						Required:       false,
						Fields: []*starr.FieldInput{
							{
								Name:  "value",
								Value: "DTS.?(HD|ES|X(?!\\D))|TRUEHD|ATMOS|DD(\\+|P).?([5-9])|EAC3.?([5-9])",
							},
						},
					},
					{
						Implementation: "LanguageSpecification",
						Negate:         false,
						Required:       false,
						Fields: []*starr.FieldInput{
							{
								Name:  "value",
								Value: 31,
							},
						},
						Name: "Arabic",
					},
				},
			},
			ExpectedRequest: addCustomFormat + "\n",
			ResponseBody:    `{"message": "NotFound"}`,
			WithError:       starr.ErrInvalidStatusCode,
			WithResponse:    (*radarr.CustomFormatOutput)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := radarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.AddCustomFormat(test.WithRequest.(*radarr.CustomFormatInput))
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestUpdateCustomFormat(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "customFormat", "1"),
			ExpectedMethod: "PUT",
			ResponseStatus: 200,
			WithRequest: &radarr.CustomFormatInput{
				ID:                    1,
				IncludeCFWhenRenaming: false,
				Name:                  "test",
				Specifications: []*radarr.CustomFormatInputSpec{
					{
						Name:           "Surround Sound",
						Implementation: "ReleaseTitleSpecification",
						Negate:         false,
						Required:       false,
						Fields: []*starr.FieldInput{
							{
								Name:  "value",
								Value: "DTS.?(HD|ES|X(?!\\D))|TRUEHD|ATMOS|DD(\\+|P).?([5-9])|EAC3.?([5-9])",
							},
						},
					},
					{
						Implementation: "LanguageSpecification",
						Negate:         false,
						Required:       false,
						Fields: []*starr.FieldInput{
							{
								Name:  "value",
								Value: 31,
							},
						},
						Name: "Arabic",
					},
				},
			},
			ExpectedRequest: updateCustomFormat + "\n",
			ResponseBody:    customFormatResponseBody,
			WithResponse: &radarr.CustomFormatOutput{
				ID:                    1,
				Name:                  "test",
				IncludeCFWhenRenaming: false,
				Specifications: []*radarr.CustomFormatOutputSpec{
					{
						Name:               "Surround Sound",
						Implementation:     "ReleaseTitleSpecification",
						ImplementationName: "Release Title",
						InfoLink:           "https://wiki.servarr.com/radarr/settings#custom-formats-2",
						Negate:             false,
						Required:           false,
						Fields: []*starr.FieldOutput{
							{
								Order:    0,
								Name:     "value",
								Label:    "Regular Expression",
								HelpText: "Custom Format RegEx is Case Insensitive",
								Value:    "DTS.?(HD|ES|X(?!\\D))|TRUEHD|ATMOS|DD(\\+|P).?([5-9])|EAC3.?([5-9])",
								Type:     "textbox",
								Advanced: false,
							},
						},
					},
					{
						Name:               "Arabic",
						Implementation:     "LanguageSpecification",
						ImplementationName: "Language",
						InfoLink:           "https://wiki.servarr.com/radarr/settings#custom-formats-2",
						Negate:             false,
						Required:           false,
						Fields: []*starr.FieldOutput{
							{
								Order: 0,
								Name:  "value",
								Label: "Language",
								// float because of unmarshal.
								Value:    float64(31),
								Type:     "select",
								Advanced: false,
								SelectOptions: []*starr.SelectOption{
									{
										Value:        0,
										Name:         "Unknown",
										Order:        0,
										DividerAfter: true,
									},
									{
										Value:        31,
										Name:         "Arabic",
										Order:        0,
										DividerAfter: false,
									},
								},
							},
						},
					},
				},
			},
			WithError: nil,
		},
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "customFormat", "1"),
			ExpectedMethod: "PUT",
			ResponseStatus: 404,
			WithRequest: &radarr.CustomFormatInput{
				ID:                    1,
				IncludeCFWhenRenaming: false,
				Name:                  "test",
				Specifications: []*radarr.CustomFormatInputSpec{
					{
						Name:           "Surround Sound",
						Implementation: "ReleaseTitleSpecification",
						Negate:         false,
						Required:       false,
						Fields: []*starr.FieldInput{
							{
								Name:  "value",
								Value: "DTS.?(HD|ES|X(?!\\D))|TRUEHD|ATMOS|DD(\\+|P).?([5-9])|EAC3.?([5-9])",
							},
						},
					},
					{
						Implementation: "LanguageSpecification",
						Negate:         false,
						Required:       false,
						Fields: []*starr.FieldInput{
							{
								Name:  "value",
								Value: 31,
							},
						},
						Name: "Arabic",
					},
				},
			},
			ExpectedRequest: updateCustomFormat + "\n",
			ResponseBody:    `{"message": "NotFound"}`,
			WithError:       starr.ErrInvalidStatusCode,
			WithResponse:    (*radarr.CustomFormatOutput)(nil),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			mockServer := test.GetMockServer(t)
			client := radarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			output, err := client.UpdateCustomFormat(test.WithRequest.(*radarr.CustomFormatInput))
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
			assert.EqualValues(t, test.WithResponse, output, "response is not the same as expected")
		})
	}
}

func TestDeleteCustomFormat(t *testing.T) {
	t.Parallel()

	tests := []*starrtest.MockData{
		{
			Name:           "200",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "customFormat", "2"),
			ExpectedMethod: "DELETE",
			WithRequest:    int64(2),
			ResponseStatus: 200,
			ResponseBody:   "{}",
			WithError:      nil,
		},
		{
			Name:           "404",
			ExpectedPath:   path.Join("/", starr.API, radarr.APIver, "customFormat", "2"),
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
			client := radarr.New(starr.New("mockAPIkey", mockServer.URL, 0))
			err := client.DeleteCustomFormat(test.WithRequest.(int64))
			assert.ErrorIs(t, err, test.WithError, "error is not the same as expected")
		})
	}
}
