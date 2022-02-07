package radarr_test

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golift.io/starr/mocks"
	"golift.io/starr/radarr"
)

// testGetReady is used in almost every test.
func testGetReady(t *testing.T) (*mocks.MockAPIer, *radarr.Radarr, *assert.Assertions) {
	t.Helper()
	ctrl := gomock.NewController(t)
	mock := mocks.NewMockAPIer(ctrl)

	return mock, &radarr.Radarr{APIer: mock}, assert.New(t)
}
