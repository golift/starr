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
	EnableAutomaticSearch   bool           `json:"enableAutomaticSearch"`
	EnableInteractiveSearch bool           `json:"enableInteractiveSearch"`
	EnableRss               bool           `json:"enableRss"`
	Priority                int64          `json:"priority"`
	ID                      int64          `json:"id,omitempty"`
	ConfigContract          string         `json:"configContract"`
	Implementation          string         `json:"implementation"`
	Name                    string         `json:"name"`
	Protocol                string         `json:"protocol"`
	Fields                  IndexerFields  `json:"fields"`
	Tags                    []*starr.Value `json:"tags"`
}

type IndexerOutput struct {
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
	Fields                  IndexerFields  `json:"fields"`
	Tags                    []*starr.Value `json:"tags"`
}

type IndexerFields struct {
	AllowZeroSize *bool `json:"allowZeroSize,omitempty"`
	// RankedOnly           bool    `json:"rankedOnly,omitempty"`
	// Delay                int64   `json:"delay,omitempty"`
	// MinimumSeeders       int64   `json:"minimumSeeders,omitempty"`
	// SeasonPackSeedTime   int64   `json:"seasonPackSeedTime,omitempty"`
	// SeedTime             int64   `json:"seedTime,omitempty"`
	// SeedRatio            float64 `json:"seedRatio,omitempty"`
	// AdditionalParameters string  `json:"additionalParameters,omitempty"`
	ApiKey  *string `json:"apiKey,omitempty"`
	ApiPath *string `json:"apiPath,omitempty"`
	// BaseUrl              string  `json:"baseUrl,omitempty"`
	// CaptchaToken         string  `json:"captchaToken,omitempty"`
	// Cookie               string  `json:"cookie,omitempty"`
	// Passkey              string  `json:"passkey,omitempty"`
	// Username             string  `json:"username,omitempty"`
	// AnimeCategories      []int64 `json:"animeCategories,omitempty"`
	// Categories           []int64 `json:"Categories,omitempty"`
}

type indexerAPI struct {
	EnableAutomaticSearch   bool           `json:"enableAutomaticSearch"`
	EnableInteractiveSearch bool           `json:"enableInteractiveSearch"`
	EnableRss               bool           `json:"enableRss"`
	SupportsRss             bool           `json:"supportsRss,omitempty"`
	SupportsSearch          bool           `json:"supportsSearch,omitempty"`
	Priority                int64          `json:"priority"`
	ID                      int64          `json:"id,omitempty"`
	ConfigContract          string         `json:"configContract"`
	Implementation          string         `json:"implementation"`
	ImplementationName      string         `json:"implementationName,omitempty"`
	InfoLink                string         `json:"infoLink,omitempty"`
	Name                    string         `json:"name"`
	Protocol                string         `json:"protocol"`
	Tags                    []*starr.Value `json:"tags"`
	Fields                  []*starr.Field `json:"fields"`
}

const bpIndexer = APIver + "/indexer"

// var (
// 	IndexerFieldsString   = []string{"additionalParameters", "apiKey", "apiPath", "baseUrl", "captchaToken", "cookie", "passkey", "username"}
// 	IndexerFieldsBool     = []string{"allowZeroSize", "rankedOnly"}
// 	IndexerFieldsInt      = []string{"delay", "minimumSeeders", "seedCriteria.seasonPackSeedTime", "seedCriteria.seedTime"}
// 	IndexerFieldsFloat    = []string{"seedCriteria.seedRatio"}
// 	IndexerFieldsIntSlice = []string{"animeCategories", "categories"}
// )

// GetIndexer returns a single indexer.
func (s *Sonarr) GetIndexer(indexerID int) (*IndexerOutput, error) {
	return s.GetIndexerContext(context.Background(), indexerID)
}

func (s *Sonarr) GetIndexerContext(ctx context.Context, indexerID int) (*IndexerOutput, error) {
	var response *indexerAPI

	uri := path.Join(bpIndexer, strconv.Itoa(indexerID))
	if _, err := s.GetInto(ctx, uri, nil, &response); err != nil {
		return nil, fmt.Errorf("api.Get(indexer): %w", err)
	}

	output, err := getIndexerOutput(response)
	if err != nil {
		return nil, err
	}

	return output, nil
}

// AddIndexer creates a indexer.
func (s *Sonarr) AddIndexer(indexer *IndexerInput) (*IndexerOutput, error) {
	return s.AddIndexerContext(context.Background(), indexer)
}

func (s *Sonarr) AddIndexerContext(ctx context.Context, indexer *IndexerInput) (*IndexerOutput, error) {
	request, err := setIndexerInput(indexer)
	if err != nil {
		return nil, err
	}

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(request); err != nil {
		return nil, fmt.Errorf("json.Marshal(indexer): %w", err)
	}

	var response *indexerAPI
	if _, err := s.PostInto(ctx, bpIndexer, nil, &body, &response); err != nil {
		return nil, fmt.Errorf("api.Post(indexer): %w", err)
	}

	output, err := getIndexerOutput(response)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func getIndexerOutput(indexer *indexerAPI) (*IndexerOutput, error) {
	fields, err := getIndexerFields(indexer.Fields)
	if err != nil {
		return nil, err
	}
	return &IndexerOutput{
		EnableAutomaticSearch:   indexer.EnableAutomaticSearch,
		EnableInteractiveSearch: indexer.EnableInteractiveSearch,
		EnableRss:               indexer.EnableRss,
		SupportsRss:             indexer.SupportsRss,
		SupportsSearch:          indexer.SupportsSearch,
		Priority:                indexer.Priority,
		ID:                      indexer.ID,
		ConfigContract:          indexer.ConfigContract,
		Implementation:          indexer.Implementation,
		ImplementationName:      indexer.ImplementationName,
		InfoLink:                indexer.InfoLink,
		Name:                    indexer.Name,
		Protocol:                indexer.Protocol,
		Fields:                  *fields,
		Tags:                    indexer.Tags,
	}, nil
}

func getIndexerFields(fields []*starr.Field) (*IndexerFields, error) {
	var output *IndexerFields
	var err error
	for _, f := range fields {
		switch f.Name {
		case "apiKey":
			*output.ApiKey, err = f.GetFieldValueString()
			if err != nil {
				return nil, err
			}
		case "apiPath":
			*output.ApiPath, err = f.GetFieldValueString()
			if err != nil {
				return nil, err
			}
		case "allowZeroSize":
			*output.AllowZeroSize, err = f.GetFieldValueBool()
			if err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("Field %s is not a known indexer field", f.Name)
		}
	}

	return output, nil
}

func setIndexerInput(indexer *IndexerInput) (*indexerAPI, error) {
	fields := setIndexerFields(&indexer.Fields)
	return &indexerAPI{
		EnableAutomaticSearch:   indexer.EnableAutomaticSearch,
		EnableInteractiveSearch: indexer.EnableInteractiveSearch,
		EnableRss:               indexer.EnableRss,
		Priority:                indexer.Priority,
		ID:                      indexer.ID,
		ConfigContract:          indexer.ConfigContract,
		Implementation:          indexer.Implementation,
		Name:                    indexer.Name,
		Protocol:                indexer.Protocol,
		Fields:                  fields,
		Tags:                    indexer.Tags,
	}, nil
}

func setIndexerFields(fields *IndexerFields) []*starr.Field {
	var output []*starr.Field
	if fields.ApiKey != nil {
		output = append(output, &starr.Field{
			Name:  "apiKey",
			Value: *fields.ApiKey,
		})
	}
	if fields.ApiPath != nil {
		output = append(output, &starr.Field{
			Name:  "apiPath",
			Value: *fields.ApiPath,
		})
	}
	if fields.AllowZeroSize != nil {
		output = append(output, &starr.Field{
			Name:  "allowZeroSize",
			Value: *fields.AllowZeroSize,
		})
	}
	return output
}
