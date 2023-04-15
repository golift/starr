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

const bpImportList = APIver + "/importList"

// ImportListInput is the input for a new or updated import list.
type ImportListInput struct {
	EnableAutomaticAdd bool                `json:"enableAutomaticAdd"`
	SeasonFolder       bool                `json:"seasonFolder"`
	ListOrder          int                 `json:"listOrder"`
	QualityProfileID   int64               `json:"qualityProfileId"`
	ID                 int64               `json:"id,omitempty"` // update only.
	ConfigContract     string              `json:"configContract"`
	Implementation     string              `json:"implementation"`
	ImplementationName string              `json:"implementationName"`
	InfoLink           string              `json:"infoLink"`
	ListType           string              `json:"listType"`
	MinRefreshInterval string              `json:"minRefreshInterval"`
	Name               string              `json:"name"`
	RootFolderPath     string              `json:"rootFolderPath"`
	SeriesType         string              `json:"seriesType"`
	ShouldMonitor      string              `json:"shouldMonitor"`
	Tags               []int               `json:"tags"`
	Fields             []*starr.FieldInput `json:"fields"`
}

// ImportListOutput is the output from the import list methods.
type ImportListOutput struct {
	EnableAutomaticAdd bool                 `json:"enableAutomaticAdd"`
	SeasonFolder       bool                 `json:"seasonFolder"`
	QualityProfileID   int64                `json:"qualityProfileId"`
	ListOrder          int64                `json:"listOrder"`
	ID                 int64                `json:"id"`
	ConfigContract     string               `json:"configContract"`
	Implementation     string               `json:"implementation"`
	ImplementationName string               `json:"implementationName"`
	InfoLink           string               `json:"infoLink"`
	ListType           string               `json:"listType"`
	MinRefreshInterval string               `json:"minRefreshInterval"`
	Name               string               `json:"name"`
	RootFolderPath     string               `json:"rootFolderPath"`
	SeriesType         string               `json:"seriesType"`
	ShouldMonitor      string               `json:"shouldMonitor"`
	Tags               []int                `json:"tags"`
	Fields             []*starr.FieldOutput `json:"fields"`
}

// GetImportLists returns all configured import lists.
func (s *Sonarr) GetImportLists() ([]*ImportListOutput, error) {
	return s.GetImportListsContext(context.Background())
}

// GetImportListsContext returns all configured import lists.
func (s *Sonarr) GetImportListsContext(ctx context.Context) ([]*ImportListOutput, error) {
	var output []*ImportListOutput

	req := starr.Request{URI: bpImportList}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetImportList returns a single import list.
func (s *Sonarr) GetImportList(importListID int64) (*ImportListOutput, error) {
	return s.GetImportListContext(context.Background(), importListID)
}

// GetIndGetImportListContextexer returns a single import list.
func (s *Sonarr) GetImportListContext(ctx context.Context, importListID int64) (*ImportListOutput, error) {
	var output ImportListOutput

	req := starr.Request{URI: path.Join(bpImportList, fmt.Sprint(importListID))}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// AddImportList creates a import list.
func (s *Sonarr) AddImportList(importList *ImportListInput) (*ImportListOutput, error) {
	return s.AddImportListContext(context.Background(), importList)
}

// AddImportListContext creates a import list.
func (s *Sonarr) AddImportListContext(ctx context.Context, importList *ImportListInput) (*ImportListOutput, error) {
	var output ImportListOutput

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(importList); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpImportList, err)
	}

	req := starr.Request{URI: bpImportList, Body: &body}
	if err := s.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateImportList updates the import list.
func (s *Sonarr) UpdateImportList(importList *ImportListInput, force bool) (*ImportListOutput, error) {
	return s.UpdateImportListContext(context.Background(), importList, force)
}

// UpdateImportListContext updates the import list.
func (s *Sonarr) UpdateImportListContext(
	ctx context.Context,
	importList *ImportListInput,
	force bool,
) (*ImportListOutput, error) {
	var output ImportListOutput

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(importList); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpImportList, err)
	}

	req := starr.Request{
		URI:   path.Join(bpImportList, fmt.Sprint(importList.ID)),
		Body:  &body,
		Query: url.Values{"forceSave": []string{fmt.Sprint(force)}},
	}
	if err := s.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteImportList removes a single import list.
func (s *Sonarr) DeleteImportList(importListID int64) error {
	return s.DeleteImportListContext(context.Background(), importListID)
}

// DeleteImportListContext removes a single import list.
func (s *Sonarr) DeleteImportListContext(ctx context.Context, importListID int64) error {
	req := starr.Request{URI: path.Join(bpImportList, fmt.Sprint(importListID))}
	if err := s.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
