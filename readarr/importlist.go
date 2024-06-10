package readarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"path"

	"golift.io/starr"
)

const bpImportList = APIver + "/importlist"

// ImportListInput is the input for a new or updated import list.
type ImportListInput struct {
	EnableAutomaticAdd    bool                `json:"enableAutomaticAdd"`
	ShouldMonitorExisting bool                `json:"shouldMonitorExisting"`
	ShouldSearch          bool                `json:"shouldSearch"`
	ListOrder             int                 `json:"listOrder"`
	ID                    int64               `json:"id,omitempty"` // for update not add.
	MetadataProfileID     int64               `json:"metadataProfileId,omitempty"`
	QualityProfileID      int64               `json:"qualityProfileId,omitempty"`
	ListType              string              `json:"listType,omitempty"`
	ConfigContract        string              `json:"configContract,omitempty"`
	Implementation        string              `json:"implementation,omitempty"`
	Name                  string              `json:"name,omitempty"`
	RootFolderPath        string              `json:"rootFolderPath,omitempty"`
	ShouldMonitor         string              `json:"shouldMonitor,omitempty"`
	MonitorNewItems       string              `json:"monitorNewItems,omitempty"`
	Tags                  []int               `json:"tags,omitempty"`
	Fields                []*starr.FieldInput `json:"fields,omitempty"`
}

// ImportListOutput is the output from the import list methods.
type ImportListOutput struct {
	EnableAutomaticAdd    bool                 `json:"enableAutomaticAdd"`
	ShouldMonitorExisting bool                 `json:"shouldMonitorExisting"`
	ShouldSearch          bool                 `json:"shouldSearch"`
	ID                    int64                `json:"id"`
	ListOrder             int64                `json:"listOrder"`
	MetadataProfileID     int64                `json:"metadataProfileId"`
	QualityProfileID      int64                `json:"qualityProfileId"`
	ConfigContract        string               `json:"configContract"`
	Implementation        string               `json:"implementation"`
	ImplementationName    string               `json:"implementationName"`
	InfoLink              string               `json:"infoLink"`
	ListType              string               `json:"listType"`
	MonitorNewItems       string               `json:"monitorNewItems"`
	Name                  string               `json:"name"`
	RootFolderPath        string               `json:"rootFolderPath"`
	ShouldMonitor         string               `json:"shouldMonitor"`
	Tags                  []int                `json:"tags"`
	Fields                []*starr.FieldOutput `json:"fields"`
}

// GetImportLists returns all configured import lists.
func (r *Readarr) GetImportLists() ([]*ImportListOutput, error) {
	return r.GetImportListsContext(context.Background())
}

// GetImportListsContext returns all configured import lists.
func (r *Readarr) GetImportListsContext(ctx context.Context) ([]*ImportListOutput, error) {
	var output []*ImportListOutput

	req := starr.Request{URI: bpImportList}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetImportList returns a single import list.
func (r *Readarr) GetImportList(importListID int64) (*ImportListOutput, error) {
	return r.GetImportListContext(context.Background(), importListID)
}

// GetIndGetImportListContextexer returns a single import list.
func (r *Readarr) GetImportListContext(ctx context.Context, importListID int64) (*ImportListOutput, error) {
	var output ImportListOutput

	req := starr.Request{URI: path.Join(bpImportList, fmt.Sprint(importListID))}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// AddImportList creates an import list without testing it.
func (r *Readarr) AddImportList(importList *ImportListInput) (*ImportListOutput, error) {
	return r.AddImportListContext(context.Background(), importList)
}

// AddImportListContext creates an import list without testing it.
func (r *Readarr) AddImportListContext(ctx context.Context, importList *ImportListInput) (*ImportListOutput, error) {
	var output ImportListOutput

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(importList); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpImportList, err)
	}

	req := starr.Request{URI: bpImportList, Body: &body, Query: url.Values{"forceSave": []string{"true"}}}
	if err := r.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// TestImportList tests an import list.
func (r *Readarr) TestImportList(list *ImportListInput) error {
	return r.TestImportListContextt(context.Background(), list)
}

// TestImportListContextt tests an import list.
func (r *Readarr) TestImportListContextt(ctx context.Context, list *ImportListInput) error {
	var output interface{}

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(list); err != nil {
		return fmt.Errorf("json.Marshal(%s): %w", bpImportList, err)
	}

	req := starr.Request{URI: path.Join(bpImportList, "test"), Body: &body}
	if err := r.PostInto(ctx, req, &output); err != nil {
		return fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return nil
}

// UpdateImportList updates the import list.
func (r *Readarr) UpdateImportList(importList *ImportListInput, force bool) (*ImportListOutput, error) {
	return r.UpdateImportListContext(context.Background(), importList, force)
}

// UpdateImportListContext updates the import list.
func (r *Readarr) UpdateImportListContext(
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
	if err := r.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteImportList removes a single import list.
func (r *Readarr) DeleteImportList(importListID int64) error {
	return r.DeleteImportListContext(context.Background(), importListID)
}

// DeleteImportListContext removes a single import list.
func (r *Readarr) DeleteImportListContext(ctx context.Context, importListID int64) error {
	req := starr.Request{URI: path.Join(bpImportList, fmt.Sprint(importListID))}
	if err := r.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
