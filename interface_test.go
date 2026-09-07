package starr_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golift.io/starr"
)

const initializeJS = `
window.sonarr = {
    version: "1.2.3.4",
    release: "1.2.3",
    apiKey: "abcdefg",
    apiRoot: "/dev/",
    instanceName: "Sonarr",
    theme: "dark",
};
`

// The initialize.js file lives at the root of the web server, so a Config URL without
// a trailing slash must not swallow the /initialize.js path, like New() + Get() allow.
func TestGetInitializeJS(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		urlSuffix string
	}{
		{"url without trailing slash", ""},
		{"url with trailing slash", "/"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			// Scoped to this subtest: the handler runs in another goroutine, and the
			// subtests run in parallel, so it must not be shared.
			var requestPath string

			mockServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
				requestPath = req.URL.Path
				_, _ = writer.Write([]byte(initializeJS))
			}))
			defer mockServer.Close()

			config := starr.New("mockAPIkey", mockServer.URL+test.urlSuffix, 0)
			output, err := config.GetInitializeJS(context.Background())
			require.NoError(t, err)
			assert.Equal(t, "/initialize.js", requestPath, "wrong path requested")
			require.NotNil(t, output)
			assert.Equal(t, "/dev/", output.APIRoot)
			assert.Equal(t, "abcdefg", output.APIKey)
			assert.Equal(t, "1.2.3.4", output.Version)
			assert.Equal(t, "Sonarr", output.InstanceName)
			assert.Equal(t, "sonarr", output.App)
		})
	}
}

func TestGetInitializeJSErrorStatus(t *testing.T) {
	t.Parallel()

	mockServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, _ *http.Request) {
		http.Error(writer, "NotFound", http.StatusNotFound)
	}))
	defer mockServer.Close()

	config := starr.New("mockAPIkey", mockServer.URL, 0)
	output, err := config.GetInitializeJS(context.Background())
	require.Nil(t, output)

	reqErr := new(starr.ReqError)
	require.ErrorAs(t, err, &reqErr)
	assert.Equal(t, http.StatusNotFound, reqErr.Code)
}
