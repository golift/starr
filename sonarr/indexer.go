package sonarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"
	"sort"
	"strconv"

	"golift.io/starr"
)

type Indexer struct {
	EnableAutomaticSearch   bool           `json:"enableAutomaticSearch"`
	EnableInteractiveSearch bool           `json:"enableInteractiveSearch"`
	EnableRss               bool           `json:"enableRss"`
	SupportsRss             bool           `json:"supportsRss"`
	SupportsSearch          bool           `json:"supportsSearch"`
	Priority                int64          `json:"priority"`
	ID                      int64          `json:"id,omitempty"`
	ConfigContract          string         `json:"configContract"`
	Implementation          string         `json:"implementation"`
	ImplementationName      string         `json:"implementationName"`
	InfoLink                string         `json:"infoLink"`
	Name                    string         `json:"name"`
	Protocol                string         `json:"protocol"`
	Tags                    []*starr.Value `json:"tags"`
	Fields                  []*starr.Field `json:"fields"`
}

const bpIndexer = APIver + "/indexer"

var (
	IndexerFieldsString   = []string{"additionalParameters", "apiKey", "apiPath", "baseUrl", "captchaToken", "cookie", "passkey", "username"}
	IndexerFieldsBool     = []string{"allowZeroSize", "rankedOnly"}
	IndexerFieldsInt      = []string{"delay", "minimumSeeders", "seedCriteria.seasonPackSeedTime", "seedCriteria.seedTime"}
	IndexerFieldsFloat    = []string{"seedCriteria.seedRatio"}
	IndexerFieldsIntSlice = []string{"animeCategories", "categories"}
)

// GetIndexers returns all configured indexers.
func (s *Sonarr) GetIndexers() ([]*Indexer, error) {
	return s.GetIndexersContext(context.Background())
}

func (s *Sonarr) GetIndexersContext(ctx context.Context) ([]*Indexer, error) {
	var output []*Indexer

	if _, err := s.GetInto(ctx, bpIndexer, nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(indexer): %w", err)
	}

	for i := range output {
		output[i].Fields, _ = correctIndexerFieldsValue(output[i].Fields)
	}
	return output, nil
}

// GetIndexer returns a single indexer.
func (s *Sonarr) GetIndexer(indexerID int) (*Indexer, error) {
	return s.GetIndexerContext(context.Background(), indexerID)
}

func (s *Sonarr) GetIndexerContext(ctx context.Context, indexerID int) (*Indexer, error) {
	var output *Indexer

	uri := path.Join(bpIndexer, strconv.Itoa(indexerID))
	if _, err := s.GetInto(ctx, uri, nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(indexer): %w", err)
	}

	output.Fields, _ = correctIndexerFieldsValue(output.Fields)
	return output, nil
}

// AddIndexer creates a indexer.
func (s *Sonarr) AddIndexer(indexer *Indexer) (*Indexer, error) {
	return s.AddIndexerContext(context.Background(), indexer)
}

func (s *Sonarr) AddIndexerContext(ctx context.Context, indexer *Indexer) (*Indexer, error) {
	var output Indexer

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(indexer); err != nil {
		return nil, fmt.Errorf("json.Marshal(indexer): %w", err)
	}

	if _, err := s.PostInto(ctx, bpIndexer, nil, &body, &output); err != nil {
		return nil, fmt.Errorf("api.Post(indexer): %w", err)
	}

	output.Fields, _ = correctIndexerFieldsValue(output.Fields)
	return &output, nil
}

// UpdateIndexer updates the indexer.
func (s *Sonarr) UpdateIndexer(indexer *Indexer) (*Indexer, error) {
	return s.UpdateIndexerContext(context.Background(), indexer)
}

func (s *Sonarr) UpdateIndexerContext(ctx context.Context, indexer *Indexer) (*Indexer, error) {
	var output Indexer

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(indexer); err != nil {
		return nil, fmt.Errorf("json.Marshal(indexer): %w", err)
	}

	uri := path.Join(bpIndexer, strconv.Itoa(int(indexer.ID)))
	if _, err := s.PutInto(ctx, uri, nil, &body, &output); err != nil {
		return nil, fmt.Errorf("api.Put(indexer): %w", err)
	}

	output.Fields, _ = correctIndexerFieldsValue(output.Fields)
	return &output, nil
}

// DeleteIndexer removes a single indexer.
func (s *Sonarr) DeleteIndexer(indexerID int) error {
	return s.DeleteIndexerContext(context.Background(), indexerID)
}

func (s *Sonarr) DeleteIndexerContext(ctx context.Context, indexerID int) error {
	uri := path.Join(bpIndexer, strconv.Itoa(indexerID))
	if _, err := s.Delete(ctx, uri, nil); err != nil {
		return fmt.Errorf("api.Delete(indexer): %w", err)
	}

	return nil
}

// TestIndexer test the Indexer connectivity.
func (s *Sonarr) TestIndexer(indexer *Indexer) (*Indexer, error) {
	return s.AddIndexerContext(context.Background(), indexer)
}

func (s *Sonarr) TestIndexerContext(ctx context.Context, indexer *Indexer) (*Indexer, error) {
	var output Indexer

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(indexer); err != nil {
		return nil, fmt.Errorf("json.Marshal(indexer): %w", err)
	}

	uri := path.Join(bpIndexer, "test")
	if _, err := s.PostInto(ctx, uri, nil, &body, &output); err != nil {
		return nil, fmt.Errorf("api.Post(indexerTest): %w", err)
	}

	return &output, nil
}

func correctIndexerFieldsValue(fields []*starr.Field) ([]*starr.Field, error) {
	output := make([]*starr.Field, len(fields))
	for i, f := range fields {
		output[i].Name = f.Name
		// check for string paramenters
		if v, a := f.Value.(string); (sort.SearchStrings(IndexerFieldsString, f.Name) != len(IndexerFieldsString)) && !a {
			output[i].Value = v
			continue
		}
		// check for int parameters
		if v, a := f.Value.(int64); (sort.SearchStrings(IndexerFieldsInt, f.Name) != len(IndexerFieldsInt)) && !a {
			output[i].Value = v
			continue
		}
		// check for bool parameters
		if v, a := f.Value.(bool); (sort.SearchStrings(IndexerFieldsBool, f.Name) != len(IndexerFieldsBool)) && !a {
			output[i].Value = v
			continue
		}
		// check for int slice parameters
		if j := sort.SearchStrings(IndexerFieldsIntSlice, f.Name); j != len(IndexerFieldsIntSlice) {
			slice := make([]int64, len(f.Value.([]interface{})))
			var assert bool
			for k, v := range f.Value.([]interface{}) {
				slice[k], assert = v.(int64)
				if !assert {
					return nil, fmt.Errorf("parameter %s is not of expected type", f.Name)
				}
			}
			output[i].Value = slice
			continue
		}
		// check for float parameters
		if v, a := f.Value.(float64); (sort.SearchStrings(IndexerFieldsFloat, f.Name) != len(IndexerFieldsFloat)) && !a {
			output[i].Value = v
			continue
		}
		return nil, fmt.Errorf("parameter %s is not of expected type", f.Name)
	}
	return output, nil
}
