// Package starrtest provides methods that are shared by all the tests in the other sub packages.
package starrtest

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockData allows generic testing of http inputs and outputs.
// This is used by the submodule tests.
type MockData struct {
	// Caller's request.
	WithRequest interface{}
	// Caller's response.
	WithResponse interface{}
	// Caller's response.
	WithError error
	// A name for the test.
	Name string
	// The path expected in the request ie. /api/v1/thing
	ExpectedPath string
	// The request body (json) expected from the caller.
	ExpectedRequest string
	// The request method (GET/POST) expected from the caller.
	ExpectedMethod string
	// The (json) response body returned to caller.
	ResponseBody string
	// This is the status that gets returned the caller.
	ResponseStatus int
}

const (
	// Error body for 401 response.
	BodyUnauthorized = `{"error": "Unauthorized"}`
	// Error body for 404 response.
	BodyNotFound = `{"message": "NotFound"}`
	// Error body for 405 response.
	BodyMethodNotAllowed = `{"message": "MethodNotAllowed"}`
)

// GetMockServer is used in all the submodule http tests.
func (test *MockData) GetMockServer(t *testing.T) *httptest.Server {
	t.Helper()

	return httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		assert.Equal(t, test.ExpectedPath, req.URL.String(),
			"test.ExpectedPath does not match the actual path")
		writer.WriteHeader(test.ResponseStatus)
		assert.Equal(t, test.ExpectedMethod, req.Method,
			"test.ExpectedMethod does not match the actual method")

		body, err := io.ReadAll(req.Body)
		assert.NoError(t, err)
		assert.Equal(t, test.ExpectedRequest, string(body),
			"test.ExpectedRequest does not match body for actual request")

		_, err = writer.Write([]byte(test.ResponseBody))
		assert.NoError(t, err)
	}))
}
