package sonarr

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"golift.io/starr"
)

// ErrUnsupported is returned when the selected Sonarr API has no such route.
var ErrUnsupported = errors.New("sonarr api version does not support this call")

// APIVersion returns the API version this client sends, v3 or v5.
func (s *Sonarr) APIVersion() string {
	if s == nil || s.apiVer == "" {
		return APIver
	}

	return s.apiVer
}

// WithAPI returns a client that shares this server config and uses version.
// version must be v3 or v5. Any other value selects v3.
func (s *Sonarr) WithAPI(version string) *Sonarr {
	if version != APIv5 {
		version = APIver
	}

	return &Sonarr{APIer: s.APIer, apiVer: version}
}

// NewV5 returns a Sonarr client that calls /api/v5.
func NewV5(config *starr.Config) *Sonarr {
	return New(config).WithAPI(APIv5)
}

// ForServer probes GET /api/v3/system/status and returns a v5 client when the
// major version is 5 or higher. A failed probe, or any older version, returns v3.
// The probe error is returned with that v3 client.
func (s *Sonarr) ForServer(ctx context.Context) (*Sonarr, error) {
	probe := s.WithAPI(APIver)

	status, err := probe.GetSystemStatusContext(ctx)
	if err != nil {
		return probe, err
	}

	if majorVersion(status.Version) >= 5 { //nolint:mnd // Sonarr 5 is the version boundary.
		return s.WithAPI(APIv5), nil
	}

	return probe, nil
}

func majorVersion(version string) int {
	version = strings.TrimSpace(version)
	version = strings.TrimPrefix(version, "v")
	version = strings.TrimPrefix(version, "V")
	head, _, _ := strings.Cut(version, ".")

	major, err := strconv.Atoi(head)
	if err != nil {
		return 0
	}

	return major
}

// prepare rewrites a v3 path when this client is speaking to Sonarr 5.
// v3 is a no-op, including /ping and paths that already include /api.
func (s *Sonarr) prepare(req starr.Request) starr.Request {
	req.URI = s.rewriteURI(req.URI)

	return req
}

func (s *Sonarr) rewriteURI(uri string) string {
	if s.APIVersion() != APIv5 {
		return uri
	}

	uri = strings.ReplaceAll(uri, APIver+"/", APIv5+"/")
	uri = strings.ReplaceAll(uri, "config/host", "settings/general")
	uri = strings.ReplaceAll(uri, "config/", "settings/")
	uri = strings.ReplaceAll(uri, "/notification", "/connection")
	uri = strings.ReplaceAll(uri, "/qualityDefinition/update", "/qualityDefinition")

	return uri
}

// Get rewrites the URI, then performs the GET.
//
//nolint:wrapcheck // Callers wrap this the same way they wrapped the embedded client.
func (s *Sonarr) Get(ctx context.Context, req starr.Request) (*http.Response, error) {
	return s.APIer.Get(ctx, s.prepare(req))
}

// Post rewrites the URI, then performs the POST.
//
//nolint:wrapcheck // Callers wrap this the same way they wrapped the embedded client.
func (s *Sonarr) Post(ctx context.Context, req starr.Request) (*http.Response, error) {
	return s.APIer.Post(ctx, s.prepare(req))
}

// Put rewrites the URI, then performs the PUT.
//
//nolint:wrapcheck // Callers wrap this the same way they wrapped the embedded client.
func (s *Sonarr) Put(ctx context.Context, req starr.Request) (*http.Response, error) {
	return s.APIer.Put(ctx, s.prepare(req))
}

// Delete rewrites the URI, then performs the DELETE.
//
//nolint:wrapcheck // Callers wrap this the same way they wrapped the embedded client.
func (s *Sonarr) Delete(ctx context.Context, req starr.Request) (*http.Response, error) {
	return s.APIer.Delete(ctx, s.prepare(req))
}

// GetInto rewrites the URI, then performs the GET and decodes JSON.
//
//nolint:wrapcheck // Callers wrap this the same way they wrapped the embedded client.
func (s *Sonarr) GetInto(ctx context.Context, req starr.Request, output any) error {
	return s.APIer.GetInto(ctx, s.prepare(req), output)
}

// PostInto rewrites the URI, then performs the POST and decodes JSON.
//
//nolint:wrapcheck // Callers wrap this the same way they wrapped the embedded client.
func (s *Sonarr) PostInto(ctx context.Context, req starr.Request, output any) error {
	return s.APIer.PostInto(ctx, s.prepare(req), output)
}

// PutInto rewrites the URI, then performs the PUT and decodes JSON.
//
//nolint:wrapcheck // Callers wrap this the same way they wrapped the embedded client.
func (s *Sonarr) PutInto(ctx context.Context, req starr.Request, output any) error {
	return s.APIer.PutInto(ctx, s.prepare(req), output)
}

// DeleteAny rewrites the URI, then performs the DELETE.
//
//nolint:wrapcheck // Callers wrap this the same way they wrapped the embedded client.
func (s *Sonarr) DeleteAny(ctx context.Context, req starr.Request) error {
	return s.APIer.DeleteAny(ctx, s.prepare(req))
}
