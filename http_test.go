package starr_test

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"golift.io/starr"
)

func TestSetPath(t *testing.T) {
	t.Parallel()

	api := path.Join("/", starr.API)
	cnfg := starr.Config{URL: "http://short.zz"}

	// These must all return the same value...
	assert.Equal(t, cnfg.URL+api+"/v1/test", cnfg.SetPath("v1/test"))     // no slashes.
	assert.Equal(t, cnfg.URL+api+"/v1/test", cnfg.SetPath("v1/test/"))    // trailing slash.
	assert.Equal(t, cnfg.URL+api+"/v1/test", cnfg.SetPath("/v1/test"))    // leading slash.
	assert.Equal(t, cnfg.URL+api+"/v1/test", cnfg.SetPath("/v1/test/"))   // both slashes.
	assert.Equal(t, cnfg.URL+api+"/v1/test", cnfg.SetPath("api/v1/test")) // ...and repeat.
	assert.Equal(t, cnfg.URL+api+"/v1/test", cnfg.SetPath("api/v1/test/"))
	assert.Equal(t, cnfg.URL+api+"/v1/test", cnfg.SetPath("/api/v1/test"))
	assert.Equal(t, cnfg.URL+api+"/v1/test", cnfg.SetPath("/api/v1/test/"))

	// These must all return the same value...
	assert.Equal(t, cnfg.URL+api+"/v1/test/another/level", cnfg.SetPath("v1/test/another/level"))
	assert.Equal(t, cnfg.URL+api+"/v1/test/another/level", cnfg.SetPath("v1/test/another/level/"))
	assert.Equal(t, cnfg.URL+api+"/v1/test/another/level", cnfg.SetPath("/v1/test/another/level"))
	assert.Equal(t, cnfg.URL+api+"/v1/test/another/level", cnfg.SetPath("/v1/test/another/level/"))
	assert.Equal(t, cnfg.URL+api+"/v1/test/another/level", cnfg.SetPath("api/v1/test/another/level"))
	assert.Equal(t, cnfg.URL+api+"/v1/test/another/level", cnfg.SetPath("api/v1/test/another/level/"))
	assert.Equal(t, cnfg.URL+api+"/v1/test/another/level", cnfg.SetPath("/api/v1/test/another/level"))
	assert.Equal(t, cnfg.URL+api+"/v1/test/another/level", cnfg.SetPath("/api/v1/test/another/level/"))
}
