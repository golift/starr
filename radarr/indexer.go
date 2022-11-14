package radarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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
	DownloadClientID        int64                `json:"DownloadClientID"`
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
func (r *Radarr) GetIndexers() ([]*IndexerOutput, error) {
	return r.GetIndexersContext(context.Background())
}

// GetIndexersContext returns all configured indexers.
func (r *Radarr) GetIndexersContext(ctx context.Context) ([]*IndexerOutput, error) {
	var output []*IndexerOutput

	req := starr.Request{URI: bpIndexer}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetIndexer returns a single indexer.
func (r *Radarr) GetIndexer(indexerID int64) (*IndexerOutput, error) {
	return r.GetIndexerContext(context.Background(), indexerID)
}

// GetIndGetIndexerContextexer returns a single indexer.
func (r *Radarr) GetIndexerContext(ctx context.Context, indexerID int64) (*IndexerOutput, error) {
	var output IndexerOutput

	req := starr.Request{URI: path.Join(bpIndexer, fmt.Sprint(indexerID))}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// AddIndexer creates a indexer.
func (r *Radarr) AddIndexer(indexer *IndexerInput) (*IndexerOutput, error) {
	return r.AddIndexerContext(context.Background(), indexer)
}

// AddIndexerContext creates a indexer.
func (r *Radarr) AddIndexerContext(ctx context.Context, indexer *IndexerInput) (*IndexerOutput, error) {
	var output IndexerOutput

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(indexer); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpIndexer, err)
	}

	req := starr.Request{URI: bpIndexer, Body: &body}
	if err := r.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateIndexer updates the indexer.
func (r *Radarr) UpdateIndexer(indexer *IndexerInput) (*IndexerOutput, error) {
	return r.UpdateIndexerContext(context.Background(), indexer)
}

// UpdateIndexerContext updates the indexer.
func (r *Radarr) UpdateIndexerContext(ctx context.Context, indexer *IndexerInput) (*IndexerOutput, error) {
	var output IndexerOutput

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(indexer); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpIndexer, err)
	}

	req := starr.Request{URI: path.Join(bpIndexer, fmt.Sprint(indexer.ID)), Body: &body}
	if err := r.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteIndexer removes a single indexer.
func (r *Radarr) DeleteIndexer(indexerID int64) error {
	return r.DeleteIndexerContext(context.Background(), indexerID)
}

// DeleteIndexerContext removes a single indexer.
func (r *Radarr) DeleteIndexerContext(ctx context.Context, indexerID int64) error {
	req := starr.Request{URI: path.Join(bpIndexer, fmt.Sprint(indexerID))}
	if err := r.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
