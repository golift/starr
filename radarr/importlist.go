package radarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"

	"golift.io/starr"
)

const bpImportList = APIver + "/importlist"

// ImportList represents the api/v3/importlist endpoint.
type ImportListInput struct {
	EnableAuto          bool                `json:"enableAuto"`
	Enabled             bool                `json:"enabled"`
	SearchOnAdd         bool                `json:"searchOnAdd"`
	ShouldMonitor       bool                `json:"shouldMonitor"`
	ID                  int64               `json:"id"`
	QualityProfileID    int64               `json:"qualityProfileId"`
	ConfigContract      string              `json:"configContract"`
	Implementation      string              `json:"implementation"`
	Name                string              `json:"name"`
	RootFolderPath      string              `json:"rootFolderPath"`
	MinimumAvailability Availability        `json:"minimumAvailability"`
	Tags                []int               `json:"tags"`
	Fields              []*starr.FieldInput `json:"fields"`
}

// ImportList represents the api/v3/importlist endpoint.
type ImportListOutput struct {
	EnableAuto          bool                 `json:"enableAuto"`
	Enabled             bool                 `json:"enabled"`
	SearchOnAdd         bool                 `json:"searchOnAdd"`
	ShouldMonitor       bool                 `json:"shouldMonitor"`
	ID                  int64                `json:"id"`
	ListOrder           int64                `json:"listOrder"`
	QualityProfileID    int64                `json:"qualityProfileId"`
	ConfigContract      string               `json:"configContract"`
	Implementation      string               `json:"implementation"`
	ImplementationName  string               `json:"implementationName"`
	InfoLink            string               `json:"infoLink"`
	ListType            string               `json:"listType"`
	Name                string               `json:"name"`
	RootFolderPath      string               `json:"rootFolderPath"`
	MinimumAvailability Availability         `json:"minimumAvailability"`
	Tags                []int                `json:"tags"`
	Fields              []*starr.FieldOutput `json:"fields"`
}

// GetImportLists returns all import lists.
func (r *Radarr) GetImportLists() ([]*ImportListOutput, error) {
	return r.GetImportListsContext(context.Background())
}

// GetImportListsContext returns all import lists.
func (r *Radarr) GetImportListsContext(ctx context.Context) ([]*ImportListOutput, error) {
	var output []*ImportListOutput

	req := starr.Request{URI: bpImportList}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// CreateImportList creates an import list in Radarr.
func (r *Radarr) CreateImportList(list *ImportListInput) (*ImportListOutput, error) {
	return r.CreateImportListContext(context.Background(), list)
}

// CreateImportListContext creates an import list in Radarr.
func (r *Radarr) CreateImportListContext(ctx context.Context, list *ImportListInput) (*ImportListOutput, error) {
	list.ID = 0

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(list); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpImportList, err)
	}

	var output ImportListOutput

	req := starr.Request{URI: bpImportList, Body: &body}
	if err := r.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteImportList removes an import list from Radarr.
func (r *Radarr) DeleteImportList(ids []int64) error {
	return r.DeleteImportListContext(context.Background(), ids)
}

// DeleteImportListContext removes an import list from Radarr.
func (r *Radarr) DeleteImportListContext(ctx context.Context, ids []int64) error {
	var errs string

	for _, id := range ids {
		req := starr.Request{URI: path.Join(bpImportList, fmt.Sprint(id))}
		if err := r.DeleteAny(ctx, req); err != nil {
			errs += fmt.Errorf("api.Delete(%s): %w", &req, err).Error() + " "
		}
	}

	if errs != "" {
		return fmt.Errorf("%w: %s", starr.ErrRequestError, errs)
	}

	return nil
}

// UpdateImportList updates an existing import list and returns the response.
func (r *Radarr) UpdateImportList(list *ImportListInput) (*ImportListOutput, error) {
	return r.UpdateImportListContext(context.Background(), list)
}

// UpdateImportListContext updates an existing import list and returns the response.
func (r *Radarr) UpdateImportListContext(ctx context.Context, list *ImportListInput) (*ImportListOutput, error) {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(list); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpImportList, err)
	}

	var output ImportListOutput

	req := starr.Request{URI: path.Join(bpImportList, fmt.Sprint(list.ID)), Body: &body}
	if err := r.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}
