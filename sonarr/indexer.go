package sonarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"
	"strconv"

	"golift.io/starr"
)

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

const bpIndexer = APIver + "/indexer"

// GetIndexers returns all configured indexerss.
func (s *Sonarr) GetIndexers() ([]*IndexerOutput, error) {
	return s.GetIndexersContext(context.Background())
}

func (s *Sonarr) GetIndexersContext(ctx context.Context) ([]*IndexerOutput, error) {
	var output []*IndexerOutput

	if _, err := s.GetInto(ctx, bpIndexer, nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(qualityProfile): %w", err)
	}

	return output, nil
}

// GetIndexer returns a single indexer.
func (s *Sonarr) GetIndexer(indexerID int) (*IndexerOutput, error) {
	return s.GetIndexerContext(context.Background(), indexerID)
}

func (s *Sonarr) GetIndexerContext(ctx context.Context, indexerID int) (*IndexerOutput, error) {
	var output *IndexerOutput

	uri := path.Join(bpIndexer, strconv.Itoa(indexerID))
	if _, err := s.GetInto(ctx, uri, nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(indexer): %w", err)
	}

	return output, nil
}

// AddIndexer creates a indexer.
func (s *Sonarr) AddIndexer(indexer *IndexerInput) (*IndexerOutput, error) {
	return s.AddIndexerContext(context.Background(), indexer)
}

func (s *Sonarr) AddIndexerContext(ctx context.Context, indexer *IndexerInput) (*IndexerOutput, error) {
	var output *IndexerOutput

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(indexer); err != nil {
		return nil, fmt.Errorf("json.Marshal(indexer): %w", err)
	}

	if _, err := s.PostInto(ctx, bpIndexer, nil, &body, &output); err != nil {
		return nil, fmt.Errorf("api.Post(indexer): %w", err)
	}

	return output, nil
}

// UpdateIndexer updates the indexer.
func (s *Sonarr) UpdateIndexer(indexer *IndexerInput) (*IndexerOutput, error) {
	return s.UpdateIndexerContext(context.Background(), indexer)
}

func (s *Sonarr) UpdateIndexerContext(ctx context.Context, indexer *IndexerInput) (*IndexerOutput, error) {
	var output IndexerOutput

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(indexer); err != nil {
		return nil, fmt.Errorf("json.Marshal(Indexer): %w", err)
	}

	uri := path.Join(bpIndexer, strconv.Itoa(int(indexer.ID)))
	if _, err := s.PutInto(ctx, uri, nil, &body, &output); err != nil {
		return nil, fmt.Errorf("api.Put(Indexer): %w", err)
	}

	return &output, nil
}

// DeleteIndexer removes a single indexer.
func (s *Sonarr) DeleteIndexer(indexerID int) error {
	return s.DeleteIndexerContext(context.Background(), indexerID)
}

func (s *Sonarr) DeleteIndexerContext(ctx context.Context, indexerID int) error {
	uri := path.Join(bpIndexer, strconv.Itoa(indexerID))
	if _, err := s.Delete(ctx, uri, nil); err != nil {
		return fmt.Errorf("api.Delete(Indexer): %w", err)
	}

	return nil
}
