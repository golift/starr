package sonarr_test

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golift.io/starr/mocks"
	"golift.io/starr/sonarr"
)

// testGetReady is used in almost every test.
func testGetReady(t *testing.T) (*mocks.MockAPIer, *sonarr.Sonarr, *assert.Assertions) {
	t.Helper()
	ctrl := gomock.NewController(t)
	mock := mocks.NewMockAPIer(ctrl)

	return mock, &sonarr.Sonarr{APIer: mock}, assert.New(t)
}
