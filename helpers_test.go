package starr_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golift.io/starr"
)

//nolint:testifylint // we want to test each one and not fail on an error.
func TestNone(t *testing.T) {
	t.Parallel()
	assert.ErrorIs(t, starr.None(starr.ErrNilClient), starr.ErrNilClient)
	assert.ErrorIs(t, starr.None("string", starr.ErrNilClient), starr.ErrNilClient)
	assert.ErrorIs(t, starr.None(uint(1), starr.ErrNilClient), starr.ErrNilClient)
	assert.ErrorIs(t, starr.None("string", uint(1), starr.ErrNilClient), starr.ErrNilClient)
	assert.ErrorIs(t, starr.None(1.0, "string", starr.ErrNilClient), starr.ErrNilClient)
	assert.NoError(t, starr.None(1.0, "string"))
	assert.NoError(t, starr.None("string"))
	assert.NoError(t, starr.None(1.0))
	assert.NoError(t, starr.None())
}
