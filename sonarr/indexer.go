package sonarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"path"

	"golift.io/starr"
)

const bpIndexer = APIver + "/indexer"

// IndexerInput is the input for a new or updated indexer.
type IndexerInput struct {
	EnableAutomaticSearch   bool                `json:"enableAutomaticSearch"`
	EnableInteractiveSearch bool                `json:"enableInteractiveSearch"`
	EnableRss               bool                `json:"enableRss"`
	DownloadClientID        int64               `json:"downloadClientId"`
	Priority                int64               `json:"priority"`
	ID                      int64               `json:"id,omitempty"`
	ConfigContract          string              `json:"configContract"`
	Implementation          string              `json:"implementation"`
	Name                    string              `json:"name"`
	Protocol                string              `json:"protocol"`
	Tags                    []int               `json:"tags"`
	Fields                  []*starr.FieldInput `json:"fields"`
}

// IndexerOutput is the output from the indexer methods.
type IndexerOutput struct {
	EnableAutomaticSearch   bool                 `json:"enableAutomaticSearch"`
	EnableInteractiveSearch bool                 `json:"enableInteractiveSearch"`
	EnableRss               bool                 `json:"enableRss"`
	SupportsRss             bool                 `json:"supportsRss"`
	SupportsSearch          bool                 `json:"supportsSearch"`
	DownloadClientID        int64                `json:"downloadClientId"`
	Priority                int64                `json:"priority"`
	ID                      int64                `json:"id,omitempty"`
	ConfigContract          string               `json:"configContract"`
	Implementation          string               `json:"implementation"`
	ImplementationName      string               `json:"implementationName"`
	InfoLink                string               `json:"infoLink"`
	Name                    string               `json:"name"`
	Protocol                string               `json:"protocol"`
	Tags                    []int                `json:"tags"`
	Fields                  []*starr.FieldOutput `json:"fields"`
}

// GetIndexers returns all configured indexers.
func (s *Sonarr) GetIndexers() ([]*IndexerOutput, error) {
	return s.GetIndexersContext(context.Background())
}

// GetIndexersContext returns all configured indexers.
func (s *Sonarr) GetIndexersContext(ctx context.Context) ([]*IndexerOutput, error) {
	var output []*IndexerOutput

	req := starr.Request{URI: bpIndexer}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetIndexer returns a single indexer.
func (s *Sonarr) GetIndexer(indexerID int64) (*IndexerOutput, error) {
	return s.GetIndexerContext(context.Background(), indexerID)
}

// GetIndGetIndexerContextexer returns a single indexer.
func (s *Sonarr) GetIndexerContext(ctx context.Context, indexerID int64) (*IndexerOutput, error) {
	var output IndexerOutput

	req := starr.Request{URI: path.Join(bpIndexer, fmt.Sprint(indexerID))}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// TestIndexer tests an indexer.
func (s *Sonarr) TestIndexer(indexer *IndexerInput) error {
	return s.TestIndexerContext(context.Background(), indexer)
}

// TestIndexerContext tests an indexer.
func (s *Sonarr) TestIndexerContext(ctx context.Context, indexer *IndexerInput) error {
	var output interface{} // any ok

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(indexer); err != nil {
		return fmt.Errorf("json.Marshal(%s): %w", bpIndexer, err)
	}

	req := starr.Request{URI: path.Join(bpIndexer, "test"), Body: &body}
	if err := s.PostInto(ctx, req, &output); err != nil {
		return fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return nil
}

// AddIndexer creates a indexer.
func (s *Sonarr) AddIndexer(indexer *IndexerInput) (*IndexerOutput, error) {
	return s.AddIndexerContext(context.Background(), indexer)
}

// AddIndexerContext creates a indexer.
func (s *Sonarr) AddIndexerContext(ctx context.Context, indexer *IndexerInput) (*IndexerOutput, error) {
	var output IndexerOutput

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(indexer); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpIndexer, err)
	}

	req := starr.Request{URI: bpIndexer, Body: &body}
	if err := s.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateIndexer updates the indexer.
func (s *Sonarr) UpdateIndexer(indexer *IndexerInput, force bool) (*IndexerOutput, error) {
	return s.UpdateIndexerContext(context.Background(), indexer, force)
}

// UpdateIndexerContext updates the indexer.
func (s *Sonarr) UpdateIndexerContext(
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
	if err := s.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteIndexer removes a single indexer.
func (s *Sonarr) DeleteIndexer(indexerID int64) error {
	return s.DeleteIndexerContext(context.Background(), indexerID)
}

// DeleteIndexerContext removes a single indexer.
func (s *Sonarr) DeleteIndexerContext(ctx context.Context, indexerID int64) error {
	req := starr.Request{URI: path.Join(bpIndexer, fmt.Sprint(indexerID))}
	if err := s.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
