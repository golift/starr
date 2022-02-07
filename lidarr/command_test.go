package lidarr_test

import (
	"context"
	"net/url"
	"testing"

	gomock "github.com/golang/mock/gomock"

	// override this name to make all tests look the same. ez.
	apparr "golift.io/starr/lidarr"
)

func TestGetCommands(t *testing.T) {
	t.Parallel()
	mock, app, assert := testGetReady(t)

	// Setup an expectation, return values and some test code for the APIer call (GetInto).
	mock.EXPECT().GetInto(gomock.Any(), apparr.APIver+"/command", nil, gomock.Any()).Return(nil).Do(
		// This is a fake starr.GetInto() func. This is used to mock and validate data in this method call.
		// The last argument is normally an interface{};
		// using the correct type here causes a panic if the funcion is (somehow) wrong.
		func(ctx context.Context, path string, params url.Values, output *[]*apparr.CommandResponse) {
			// This may change, but for now there are no params needed to get commands.
			assert.Nil(params, "params passed to GetInto must be nil")
			// Add something to the provided interface to make sure it comes out right.
			*output = append(*output, &apparr.CommandResponse{ID: 1, Name: "mine"})
		})

	// Now that the mock is ready, run the test.
	output, err := app.GetCommands()
	// Verify the output from the test.
	assert.Nil(err, "no error must be returned")
	assert.NotNil(output, "output must not be returned nil")
	assert.Len(output, 1, "wrong length returned by the output")
	assert.EqualValues(output[0].ID, 1, "wrong ID returned")
}
