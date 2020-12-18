package starr

//go:generate mockgen -destination=mocks/apier.go golift.io/starr APIer

import "net/url"

// APIer is used by the sub packages to allow mocking the http methods in tests.
type APIer interface {
	Get(path string, params url.Values) ([]byte, error)
	Post(path string, params url.Values, body []byte) ([]byte, error)
	GetInto(path string, params url.Values, v interface{}) error
	PostInto(path string, params url.Values, body []byte, v interface{}) error
}

var _ APIer = (*Config)(nil)
