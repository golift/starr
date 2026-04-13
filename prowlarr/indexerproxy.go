package prowlarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"path"

	"golift.io/starr"
)

const bpIndexerProxy = APIver + "/indexerproxy"

// IndexerProxyInput is used to create or update an indexer proxy.
type IndexerProxyInput struct {
	ID             int64               `json:"id,omitempty"`
	Name           string              `json:"name,omitempty"`
	Implementation string              `json:"implementation,omitempty"`
	ConfigContract string              `json:"configContract,omitempty"`
	Fields         []*starr.FieldInput `json:"fields,omitempty"`
}

// IndexerProxyOutput is returned from indexer proxy endpoints.
type IndexerProxyOutput struct {
	ID                 int64                `json:"id,omitempty"`
	Name               string               `json:"name,omitempty"`
	Implementation     string               `json:"implementation,omitempty"`
	ImplementationName string               `json:"implementationName,omitempty"`
	ConfigContract     string               `json:"configContract,omitempty"`
	Fields             []*starr.FieldOutput `json:"fields,omitempty"`
}

// GetIndexerProxies returns all indexer proxies.
func (p *Prowlarr) GetIndexerProxies() ([]*IndexerProxyOutput, error) {
	return p.GetIndexerProxiesContext(context.Background())
}

// GetIndexerProxiesContext returns all indexer proxies.
func (p *Prowlarr) GetIndexerProxiesContext(ctx context.Context) ([]*IndexerProxyOutput, error) {
	var output []*IndexerProxyOutput

	req := starr.Request{URI: bpIndexerProxy}
	if err := p.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetIndexerProxy returns a single indexer proxy.
func (p *Prowlarr) GetIndexerProxy(id int64) (*IndexerProxyOutput, error) {
	return p.GetIndexerProxyContext(context.Background(), id)
}

// GetIndexerProxyContext returns a single indexer proxy.
func (p *Prowlarr) GetIndexerProxyContext(ctx context.Context, id int64) (*IndexerProxyOutput, error) {
	var output IndexerProxyOutput

	req := starr.Request{URI: path.Join(bpIndexerProxy, starr.Str(id))}
	if err := p.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// GetIndexerProxySchema returns indexer proxy templates.
func (p *Prowlarr) GetIndexerProxySchema() ([]*IndexerProxyOutput, error) {
	return p.GetIndexerProxySchemaContext(context.Background())
}

// GetIndexerProxySchemaContext returns indexer proxy templates.
func (p *Prowlarr) GetIndexerProxySchemaContext(ctx context.Context) ([]*IndexerProxyOutput, error) {
	var output []*IndexerProxyOutput

	req := starr.Request{URI: path.Join(bpIndexerProxy, "schema")}
	if err := p.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// AddIndexerProxy creates an indexer proxy.
func (p *Prowlarr) AddIndexerProxy(proxy *IndexerProxyInput, forceSave bool) (*IndexerProxyOutput, error) {
	return p.AddIndexerProxyContext(context.Background(), proxy, forceSave)
}

// AddIndexerProxyContext creates an indexer proxy.
func (p *Prowlarr) AddIndexerProxyContext(
	ctx context.Context,
	proxy *IndexerProxyInput,
	forceSave bool,
) (*IndexerProxyOutput, error) {
	var output IndexerProxyOutput

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(proxy); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpIndexerProxy, err)
	}

	req := starr.Request{
		URI:   bpIndexerProxy,
		Body:  &body,
		Query: url.Values{"forceSave": []string{starr.Str(forceSave)}},
	}
	if err := p.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateIndexerProxy updates an indexer proxy.
func (p *Prowlarr) UpdateIndexerProxy(proxy *IndexerProxyInput, forceSave bool) (*IndexerProxyOutput, error) {
	return p.UpdateIndexerProxyContext(context.Background(), proxy, forceSave)
}

// UpdateIndexerProxyContext updates an indexer proxy.
func (p *Prowlarr) UpdateIndexerProxyContext(
	ctx context.Context,
	proxy *IndexerProxyInput,
	forceSave bool,
) (*IndexerProxyOutput, error) {
	var output IndexerProxyOutput

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(proxy); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpIndexerProxy, err)
	}

	req := starr.Request{
		URI:   path.Join(bpIndexerProxy, starr.Str(proxy.ID)),
		Body:  &body,
		Query: url.Values{"forceSave": []string{starr.Str(forceSave)}},
	}
	if err := p.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteIndexerProxy removes an indexer proxy.
func (p *Prowlarr) DeleteIndexerProxy(id int64) error {
	return p.DeleteIndexerProxyContext(context.Background(), id)
}

// DeleteIndexerProxyContext removes an indexer proxy.
func (p *Prowlarr) DeleteIndexerProxyContext(ctx context.Context, id int64) error {
	req := starr.Request{URI: path.Join(bpIndexerProxy, starr.Str(id))}
	if err := p.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}

// TestIndexerProxy tests indexer proxy settings.
func (p *Prowlarr) TestIndexerProxy(proxy *IndexerProxyInput, forceTest bool) error {
	return p.TestIndexerProxyContext(context.Background(), proxy, forceTest)
}

// TestIndexerProxyContext tests indexer proxy settings.
func (p *Prowlarr) TestIndexerProxyContext(ctx context.Context, proxy *IndexerProxyInput, forceTest bool) error {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(proxy); err != nil {
		return fmt.Errorf("json.Marshal(%s): %w", path.Join(bpIndexerProxy, "test"), err)
	}

	var output any

	req := starr.Request{
		URI:   path.Join(bpIndexerProxy, "test"),
		Body:  &body,
		Query: url.Values{"forceTest": []string{starr.Str(forceTest)}},
	}
	if err := p.PostInto(ctx, req, &output); err != nil {
		return fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return nil
}
