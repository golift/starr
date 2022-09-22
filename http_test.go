package starr_test

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
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
