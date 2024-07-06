package prowlarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"path"
	"time"

	"golift.io/starr"
)

const bpIndexer = APIver + "/indexer"

// IndexerInput is the input for a new or updated indexer.
type IndexerInput struct {
	Enable         bool                `json:"enable"`
	Redirect       bool                `json:"redirect"`
	Priority       int64               `json:"priority"`
	ID             int64               `json:"id,omitempty"`
	AppProfileID   int64               `json:"appProfileId"`
	ConfigContract string              `json:"configContract"`
	Implementation string              `json:"implementation"`
	Name           string              `json:"name"`
	Protocol       starr.Protocol      `json:"protocol"`
	Tags           []int               `json:"tags,omitempty"`
	Fields         []*starr.FieldInput `json:"fields"`
}

// IndexerOutput is the output from the indexer methods.
type IndexerOutput struct {
	Enable             bool                 `json:"enable"`
	Redirect           bool                 `json:"redirect"`
	SupportsRss        bool                 `json:"supportsRss"`
	SupportsSearch     bool                 `json:"supportsSearch"`
	SupportsRedirect   bool                 `json:"supportsRedirect"`
	AppProfileID       int64                `json:"appProfileId"`
	ID                 int64                `json:"id,omitempty"`
	Priority           int64                `json:"priority"`
	SortName           string               `json:"sortName"`
	Name               string               `json:"name"`
	Protocol           starr.Protocol       `json:"protocol"`
	Privacy            string               `json:"privacy"`
	DefinitionName     string               `json:"definitionName"`
	Description        string               `json:"description"`
	Language           string               `json:"language"`
	Encoding           string               `json:"encoding,omitempty"`
	ImplementationName string               `json:"implementationName"`
	Implementation     string               `json:"implementation"`
	ConfigContract     string               `json:"configContract"`
	InfoLink           string               `json:"infoLink"`
	Added              time.Time            `json:"added"`
	Capabilities       *Capabilities        `json:"capabilities,omitempty"`
	Tags               []int                `json:"tags"`
	IndexerUrls        []string             `json:"indexerUrls"`
	LegacyUrls         []string             `json:"legacyUrls"`
	Fields             []*starr.FieldOutput `json:"fields"`
}

// Capabilities is part of IndexerOutput.
type Capabilities struct {
	SupportsRawSearch bool          `json:"supportsRawSearch"`
	LimitsMax         int64         `json:"limitsMax"`
	LimitsDefault     int64         `json:"limitsDefault"`
	SearchParams      []string      `json:"searchParams"`
	TvSearchParams    []string      `json:"tvSearchParams"`
	MovieSearchParams []string      `json:"movieSearchParams"`
	MusicSearchParams []string      `json:"musicSearchParams"`
	BookSearchParams  []string      `json:"bookSearchParams"`
	Categories        []*Categories `json:"categories"`
}

// Categories is part of Capabilities.
type Categories struct {
	ID            int64         `json:"id"`
	Name          string        `json:"name"`
	SubCategories []*Categories `json:"subCategories"`
}

// GetIndexers returns all configured indexers.
func (p *Prowlarr) GetIndexers() ([]*IndexerOutput, error) {
	return p.GetIndexersContext(context.Background())
}

// GetIndexersContext returns all configured indexers.
func (p *Prowlarr) GetIndexersContext(ctx context.Context) ([]*IndexerOutput, error) {
	var output []*IndexerOutput

	req := starr.Request{URI: bpIndexer}
	if err := p.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// TestIndexer tests an indexer.
func (p *Prowlarr) TestIndexer(indexer *IndexerInput) error {
	return p.TestIndexerContext(context.Background(), indexer)
}

// TestIndexerContext tests an indexer.
func (p *Prowlarr) TestIndexerContext(ctx context.Context, indexer *IndexerInput) error {
	var output interface{}

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(indexer); err != nil {
		return fmt.Errorf("json.Marshal(%s): %w", bpIndexer, err)
	}

	req := starr.Request{URI: path.Join(bpIndexer, "test"), Body: &body}
	if err := p.PostInto(ctx, req, &output); err != nil {
		return fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return nil
}

// GetIndexer returns a single indexer.
func (p *Prowlarr) GetIndexer(indexerID int64) (*IndexerOutput, error) {
	return p.GetIndexerContext(context.Background(), indexerID)
}

// GetIndexerContext returns a single indexer.
func (p *Prowlarr) GetIndexerContext(ctx context.Context, indexerID int64) (*IndexerOutput, error) {
	var output IndexerOutput

	req := starr.Request{URI: path.Join(bpIndexer, fmt.Sprint(indexerID))}
	if err := p.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// AddIndexer creates an indexer without testing it.
func (p *Prowlarr) AddIndexer(indexer *IndexerInput) (*IndexerOutput, error) {
	return p.AddIndexerContext(context.Background(), indexer)
}

// AddIndexerContext creates an indexer without testing it.
func (p *Prowlarr) AddIndexerContext(ctx context.Context, indexer *IndexerInput) (*IndexerOutput, error) {
	var (
		output IndexerOutput
		body   bytes.Buffer
	)

	indexer.ID = 0
	if err := json.NewEncoder(&body).Encode(indexer); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpIndexer, err)
	}

	req := starr.Request{URI: bpIndexer, Body: &body, Query: url.Values{"forceSave": []string{"true"}}}
	if err := p.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateIndexer updates the indexer.
func (p *Prowlarr) UpdateIndexer(indexer *IndexerInput, force bool) (*IndexerOutput, error) {
	return p.UpdateIndexerContext(context.Background(), indexer, force)
}

// UpdateIndexerContext updates the indexer.
func (p *Prowlarr) UpdateIndexerContext(
	ctx context.Context,
	indexer *IndexerInput,
	force bool,
) (*IndexerOutput, error) {
	var output IndexerOutput

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(indexer); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpIndexer, err)
	}

	req := starr.Request{
		URI:   path.Join(bpIndexer, fmt.Sprint(indexer.ID)),
		Body:  &body,
		Query: url.Values{"forceSave": []string{fmt.Sprint(force)}},
	}
	if err := p.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteIndexer removes a single indexer.
func (p *Prowlarr) DeleteIndexer(indexerID int64) error {
	return p.DeleteIndexerContext(context.Background(), indexerID)
}

// DeleteIndexerContext removes a single indexer.
func (p *Prowlarr) DeleteIndexerContext(ctx context.Context, indexerID int64) error {
	req := starr.Request{URI: path.Join(bpIndexer, fmt.Sprint(indexerID))}
	if err := p.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
