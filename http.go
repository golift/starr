package starr

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"reflect"
	"strings"
)

// API is the beginning of every API path.
const API = "api"

/* The methods in this file provide assumption-ridden HTTP calls for Starr apps. */

// Request contains the GET and/or POST values for an HTTP request.
type Request struct {
	URI   string     // Required: path portion of the URL.
	Query url.Values // GET parameters work for any request type.
	Body  io.Reader  // Used in PUT, POST, DELETE. Not for GET.
}

// ReqError is returned when a Starr app returns an invalid status code.
type ReqError struct {
	Code int
	Body []byte
	Msg  string
	Name string
	Err  error // sub error, often nil, or not useful.
	http.Header
}

// String turns a request into a string. Usually used in error messages.
func (r *Request) String() string {
	return r.URI
}

// Req makes an authenticated request to a starr application and returns the response.
// Do not forget to read and close the response Body if there is no error.
func (c *Config) Req(ctx context.Context, method string, req Request) (*http.Response, error) {
	return c.req(ctx, method, req)
}

// api is an internal function to call an api path.
func (c *Config) api(ctx context.Context, method string, req Request) (*http.Response, error) {
	req.URI = SetAPIPath(req.URI)
	return c.req(ctx, method, req)
}

// req is our abstraction method for calling a starr application.
func (c *Config) req(ctx context.Context, method string, req Request) (*http.Response, error) {
	if c.Client == nil { // we must have an http client.
		return nil, ErrNilClient
	}

	httpReq, err := http.NewRequestWithContext(ctx, method, strings.TrimSuffix(c.URL, "/")+req.URI, req.Body)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequestWithContext(%s): %w", req.URI, err)
	}

	c.SetHeaders(httpReq)

	if req.Query != nil {
		httpReq.URL.RawQuery = req.Query.Encode()
	}

	resp, err := c.Client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("httpClient.Do(req): %w", err)
	}

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return nil, parseNon200(resp)
	}

	return resp, nil
}

// parseNon200 attempts to extract an error message from a non-200 response.
func parseNon200(resp *http.Response) *ReqError {
	defer resp.Body.Close()

	response := &ReqError{Code: resp.StatusCode, Header: resp.Header}
	if response.Body, response.Err = io.ReadAll(resp.Body); response.Err != nil {
		return response
	}

	var msg struct {
		Msg string `json:"message"`
	}

	if response.Err = json.Unmarshal(response.Body, &msg); response.Err == nil && msg.Msg != "" {
		response.Msg = msg.Msg
		return response
	}

	type propError struct {
		Msg  string `json:"errorMessage"`
		Name string `json:"propertyName"`
	}

	var errMsg propError

	if response.Err = json.Unmarshal(response.Body, &errMsg); response.Err == nil && errMsg.Msg != "" {
		response.Name, response.Msg = errMsg.Name, errMsg.Msg
		return response
	}

	// Sometimes we get a list of errors. This grabs the first one.
	var errMsg2 []propError

	if response.Err = json.Unmarshal(response.Body, &errMsg2); response.Err == nil && len(errMsg2) > 0 {
		response.Name, response.Msg = errMsg2[0].Name, errMsg2[0].Msg
		return response
	}

	return response
}

// closeResp should be used to close requests that don't require a response body.
func closeResp(resp *http.Response) {
	if resp != nil && resp.Body != nil {
		_, _ = io.ReadAll(resp.Body)
		resp.Body.Close()
	}
}

// SetHeaders sets all our request headers based on method and other data.
func (c *Config) SetHeaders(req *http.Request) {
	// This app allows http auth, in addition to api key (nginx proxy).
	if auth := c.HTTPUser + ":" + c.HTTPPass; auth != ":" {
		req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(auth)))
	}

	if req.Body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	if req.Method == http.MethodPost && strings.HasSuffix(req.URL.RequestURI(), "/login") {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else {
		req.Header.Set("Accept", "application/json")
	}

	req.Header.Set("User-Agent", "go-starr: https://"+reflect.TypeOf(Config{}).PkgPath()) //nolint:exhaustivestruct
	req.Header.Set("X-API-Key", c.APIKey)
}

// SetAPIPath makes sure the path starts with /api.
func SetAPIPath(uriPath string) string {
	if strings.HasPrefix(uriPath, API+"/") ||
		strings.HasPrefix(uriPath, path.Join("/", API)+"/") {
		return path.Join("/", uriPath)
	}

	return path.Join("/", API, uriPath)
}

// Error returns the formatted error message for an invalid status code.
func (r *ReqError) Error() string {
	const (
		prefix  = "invalid status code"
		maxBody = 400 // arbitrary.
	)

	msg := fmt.Sprintf("%s, %d < %d", prefix, r.Code, http.StatusOK)
	if r.Code >= http.StatusMultipleChoices { // 300
		msg = fmt.Sprintf("%s, %d >= %d", prefix, r.Code, http.StatusMultipleChoices)
	}

	switch body := string(r.Body); {
	case r.Name != "":
		return fmt.Sprintf("%s, %s: %s", msg, r.Name, r.Msg)
	case r.Msg != "":
		return fmt.Sprintf("%s, %s", msg, r.Msg)
	case len(body) > maxBody:
		return fmt.Sprintf("%s, %s", msg, body[:maxBody])
	case len(body) != 0:
		return fmt.Sprintf("%s, %s", msg, body)
	default:
		return msg
	}
}

// Is provides a custom error match facility.
func (r *ReqError) Is(tgt error) bool {
	target, ok := tgt.(*ReqError)
	return ok && (r.Code == target.Code || target.Code == -1)
}
