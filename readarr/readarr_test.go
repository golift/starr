package readarr_test

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golift.io/starr/mocks"
	"golift.io/starr/readarr"
)

// testGetReady is used in almost every test.
func testGetReady(t *testing.T) (*mocks.MockAPIer, *readarr.Readarr, *assert.Assertions) {
	t.Helper()
	ctrl := gomock.NewController(t)
	mock := mocks.NewMockAPIer(ctrl)

	return mock, &readarr.Readarr{APIer: mock}, assert.New(t)
}
