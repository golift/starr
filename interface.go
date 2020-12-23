package starr

//go:generate mockgen -destination=mocks/apier.go golift.io/starr APIer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// APIer is used by the sub packages to allow mocking the http methods in tests.
type APIer interface {
	Get(path string, params url.Values) ([]byte, error)
	Post(path string, params url.Values, body []byte) ([]byte, error)
	Put(path string, params url.Values, body []byte) ([]byte, error)
	Delete(path string, params url.Values) ([]byte, error)
	GetInto(path string, params url.Values, v interface{}) error
	PostInto(path string, params url.Values, body []byte, v interface{}) error
	PutInto(path string, params url.Values, body []byte, v interface{}) error
	DeleteInto(path string, params url.Values, v interface{}) error
}

// Config must satify the APIer struct.
var _ APIer = (*Config)(nil)

// Get makes a GET http request and returns the body.
func (c *Config) Get(path string, params url.Values) ([]byte, error) {
	return c.req(path, http.MethodGet, params, nil)
}

// Put makes a PUT http request and returns the body.
func (c *Config) Put(path string, params url.Values, body []byte) ([]byte, error) {
	return c.req(path, http.MethodPut, params, bytes.NewBuffer(body))
}

// Post makes a POST http request and returns the body.
func (c *Config) Post(path string, params url.Values, body []byte) ([]byte, error) {
	return c.req(path, http.MethodPost, params, bytes.NewBuffer(body))
}

// Get makes a DELETE http request and returns the body.
func (c *Config) Delete(path string, params url.Values) ([]byte, error) {
	return c.req(path, http.MethodDelete, params, nil)
}

// GetInto performs an HTTP GET against an API path and unmarshals the payload into the provided pointer interface.
func (c *Config) GetInto(path string, params url.Values, v interface{}) error {
	data, err := c.Get(path, params)

	return unmarshal(v, data, err)
}

// PutInto performs an HTTP PUT against an API path and unmarshals the payload into the provided pointer interface.
func (c *Config) PutInto(path string, params url.Values, body []byte, v interface{}) error {
	data, err := c.Put(path, params, body)

	return unmarshal(v, data, err)
}

// PostInto performs an HTTP POST against an API path and unmarshals the payload into the provided pointer interface.
func (c *Config) PostInto(path string, params url.Values, body []byte, v interface{}) error {
	data, err := c.Post(path, params, body)

	return unmarshal(v, data, err)
}

// DeleteInto performs an HTTP DELETE against an API path and unmarshals the payload into a pointer interface.
func (c *Config) DeleteInto(path string, params url.Values, v interface{}) error {
	data, err := c.Delete(path, params)

	return unmarshal(v, data, err)
}

// unmarshal is an extra procedure to check an error and unmarshal the payload.
// This allows the methods above to have all their logic abstracted.
func unmarshal(v interface{}, data []byte, err error) error {
	if err != nil {
		return err
	} else if v == nil {
		return fmt.Errorf("this is a code bug: %w", ErrNilInterface)
	} else if err = json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("json parse error: %w", err)
	}

	return nil
}
