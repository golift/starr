package lidarr_test

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golift.io/starr/lidarr"
	"golift.io/starr/mocks"
)

// testGetReady is used in almost every test.
func testGetReady(t *testing.T) (*mocks.MockAPIer, *lidarr.Lidarr, *assert.Assertions) {
	t.Helper()
	ctrl := gomock.NewController(t)
	mock := mocks.NewMockAPIer(ctrl)

	return mock, &lidarr.Lidarr{APIer: mock}, assert.New(t)
}
