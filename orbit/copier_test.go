package orbit_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golift.io/starr"
	"golift.io/starr/orbit"
	"golift.io/starr/prowlarr"
	"golift.io/starr/sonarr"
)

func copyData(t *testing.T) (*prowlarr.IndexerOutput, *sonarr.IndexerInput) {
	t.Helper()

	return &prowlarr.IndexerOutput{
			ID:             2,
			Priority:       3,
			Name:           "yes",
			Protocol:       "usenet",
			Implementation: "core",
			ConfigContract: "hancock",
			Tags:           []int{1, 2, 5},
			Fields: []*starr.FieldOutput{
				{Name: "One", Value: "one"},
				{Name: "Two", Value: 2.0},
				{Name: "Three", Value: uint(3)},
				{Name: "Five", Value: 5},
			},
		},
		&sonarr.IndexerInput{
			// These are not part of the used input, so set them before copying.
			EnableAutomaticSearch:   true,
			EnableInteractiveSearch: true,
			EnableRss:               true,
			DownloadClientID:        15,
		}
}

func TestCopyIndexers(t *testing.T) {
	t.Parallel()
	src1, dst1 := copyData(t)
	src2, dst2 := copyData(t)
	src3, dst3 := copyData(t)
	src4, _ := copyData(t)
	src5, _ := copyData(t)
	// We test for these.
	src1.Priority = 1
	src2.Priority = 2
	src3.Priority = 3
	src4.Priority = 4
	src5.Priority = 5
	// Make two lists.
	srcs := append([]*prowlarr.IndexerOutput{}, src1, src2, src3, src4, src5)
	dsts := append([]*sonarr.IndexerInput{}, dst1, dst2, dst3) // Short by 2.
	// Copy the lists.
	dsts2, err := orbit.CopyIndexers(srcs, &dsts, true)
	require.NoError(t, err)
	// Make sure both outputs have a length matching the input.
	assert.Len(t, dsts, len(srcs))
	assert.Len(t, dsts2, len(srcs))
	// Test that values got copied.
	for idx, src := range srcs {
		assert.Zero(t, dsts[idx].ID)
		assert.Equal(t, src.Priority, dsts[idx].Priority)
		assert.Equal(t, src.Tags, dsts[idx].Tags)
	}
}

// TestCopyIndexersNilDest test a nil destination pointer and slice.
func TestCopyIndexersNilDest(t *testing.T) {
	t.Parallel()
	src1, _ := copyData(t)
	src2, _ := copyData(t)
	// Make two lists.
	srcs := append([]*prowlarr.IndexerOutput{}, src1, src2)
	dsts := new([]*sonarr.IndexerInput) // Super empty.
	*dsts = nil                         // Nil the slice.
	// Copy the lists.
	dsts2, err := orbit.CopyIndexers(srcs, dsts, false)
	require.NoError(t, err)
	// Make sure both outputs have a length matching the input.
	assert.Len(t, *dsts, len(srcs))
	assert.Len(t, dsts2, len(srcs))
	// Test that tags got removed.
	for idx, src := range srcs {
		assert.Zero(t, (*dsts)[idx].ID)
		assert.Equal(t, src.Priority, (*dsts)[idx].Priority)
		assert.NotEqual(t, src.Tags, (*dsts)[idx].Tags)
	}

	// Make an error.
	dsts = nil // This is a no-no.
	require.ErrorIs(t, starr.None(orbit.CopyIndexers(srcs, dsts, false)), orbit.ErrNotPtr)
}

func TestCopyIndexer(t *testing.T) {
	t.Parallel()

	src, dst := copyData(t)
	// Verify everything copies over.
	require.NoError(t, starr.None(orbit.CopyIndexer(src, dst, true)))
	assert.Equal(t, src.Fields[0].Value, dst.Fields[0].Value)
	assert.Equal(t, src.Fields[1].Value, dst.Fields[1].Value)
	assert.EqualValues(t, src.Fields[2].Value, dst.Fields[2].Value)
	assert.EqualValues(t, src.Fields[3].Value, dst.Fields[3].Value)
	assert.Equal(t, src.Fields[0].Name, dst.Fields[0].Name)
	assert.Equal(t, src.Fields[1].Name, dst.Fields[1].Name)
	assert.Equal(t, src.Fields[2].Name, dst.Fields[2].Name)
	assert.Equal(t, src.Fields[3].Name, dst.Fields[3].Name)
	assert.Zero(t, dst.ID)
	assert.Equal(t, src.Priority, dst.Priority)
	assert.Equal(t, src.Name, dst.Name)
	assert.Equal(t, src.Protocol, dst.Protocol)
	assert.Equal(t, src.Implementation, dst.Implementation)
	assert.Equal(t, src.ConfigContract, dst.ConfigContract)
	assert.Equal(t, src.Tags[0], dst.Tags[0])
	assert.Equal(t, src.Tags[1], dst.Tags[1])
	assert.Equal(t, src.Tags[2], dst.Tags[2])
	// Check passed in values.
	assert.Equal(t, int64(15), dst.DownloadClientID)
	assert.True(t, dst.EnableAutomaticSearch)
	assert.True(t, dst.EnableInteractiveSearch)
	assert.True(t, dst.EnableRss)
	// Make sure tags get depleted.
	starr.Must(orbit.CopyIndexer(src, dst, false))
	assert.Zero(t, dst.Tags)
}

func TestCopy(t *testing.T) {
	t.Parallel()

	broken := struct{}{}
	good := &prowlarr.IndexerOutput{}

	require.ErrorIs(t, orbit.Copy(broken, good), orbit.ErrNotPtr)
	require.ErrorIs(t, orbit.Copy(good, broken), orbit.ErrNotPtr)
}
