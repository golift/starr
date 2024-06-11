package starr_test

import (
	"net/http"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golift.io/starr"
)

func TestSetAPIPath(t *testing.T) {
	t.Parallel()

	api := path.Join("/", starr.API)

	// These must all return the same value...
	assert.Equal(t, api+"/v1/test", starr.SetAPIPath("v1/test"))     // no slashes.
	assert.Equal(t, api+"/v1/test", starr.SetAPIPath("v1/test/"))    // trailing slash.
	assert.Equal(t, api+"/v1/test", starr.SetAPIPath("/v1/test"))    // leading slash.
	assert.Equal(t, api+"/v1/test", starr.SetAPIPath("/v1/test/"))   // both slashes.
	assert.Equal(t, api+"/v1/test", starr.SetAPIPath("api/v1/test")) // ...and repeat.
	assert.Equal(t, api+"/v1/test", starr.SetAPIPath("api/v1/test/"))
	assert.Equal(t, api+"/v1/test", starr.SetAPIPath("/api/v1/test"))
	assert.Equal(t, api+"/v1/test", starr.SetAPIPath("/api/v1/test/"))

	// These must all return the same value...
	assert.Equal(t, api+"/v1/test/another/level", starr.SetAPIPath("v1/test/another/level"))
	assert.Equal(t, api+"/v1/test/another/level", starr.SetAPIPath("v1/test/another/level/"))
	assert.Equal(t, api+"/v1/test/another/level", starr.SetAPIPath("/v1/test/another/level"))
	assert.Equal(t, api+"/v1/test/another/level", starr.SetAPIPath("/v1/test/another/level/"))
	assert.Equal(t, api+"/v1/test/another/level", starr.SetAPIPath("api/v1/test/another/level"))
	assert.Equal(t, api+"/v1/test/another/level", starr.SetAPIPath("api/v1/test/another/level/"))
	assert.Equal(t, api+"/v1/test/another/level", starr.SetAPIPath("/api/v1/test/another/level"))
	assert.Equal(t, api+"/v1/test/another/level", starr.SetAPIPath("/api/v1/test/another/level/"))
}

func TestReqError(t *testing.T) {
	t.Parallel()

	err := &starr.ReqError{Code: http.StatusForbidden}
	require.ErrorIs(t, err, starr.ErrInvalidStatusCode)
	assert.Equal(t, "invalid status code, 403 >= 300", err.Error())

	err.Body = []byte("Some Body")
	assert.Equal(t, "invalid status code, 403 >= 300, Some Body", err.Error())

	err.Msg = "Some message"
	assert.Equal(t, "invalid status code, 403 >= 300, Some message", err.Error())

	err.Name = "Varname"
	assert.Equal(t, "invalid status code, 403 >= 300, Varname: Some message", err.Error())
}
