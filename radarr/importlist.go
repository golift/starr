package radarr

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

// ImportList represents the api/v3/importlist endpoint.
type ImportListInput struct {
	EnableAuto          bool                `json:"enableAuto"`
	Enabled             bool                `json:"enabled"`
	SearchOnAdd         bool                `json:"searchOnAdd"`
	ListOrder           int                 `json:"listOrder"`
	ID                  int64               `json:"id,omitempty"`
	QualityProfileID    int64               `json:"qualityProfileId,omitempty"`
	ConfigContract      string              `json:"configContract,omitempty"`
	Implementation      string              `json:"implementation,omitempty"`
	ImplementationName  string              `json:"implementationName,omitempty"`
	InfoLink            string              `json:"infoLink,omitempty"`
	ListType            string              `json:"listType,omitempty"`
	Monitor             string              `json:"monitor,omitempty"`
	Name                string              `json:"name,omitempty"`
	RootFolderPath      string              `json:"rootFolderPath,omitempty"`
	MinimumAvailability Availability        `json:"minimumAvailability,omitempty"`
	Tags                []int               `json:"tags,omitempty"`
	Fields              []*starr.FieldInput `json:"fields,omitempty"`
}

// ImportList represents the api/v3/importlist endpoint.
type ImportListOutput struct {
	EnableAuto          bool                 `json:"enableAuto"`
	Enabled             bool                 `json:"enabled"`
	SearchOnAdd         bool                 `json:"searchOnAdd"`
	ID                  int64                `json:"id"`
	ListOrder           int64                `json:"listOrder"`
	QualityProfileID    int64                `json:"qualityProfileId"`
	ConfigContract      string               `json:"configContract"`
	Implementation      string               `json:"implementation"`
	ImplementationName  string               `json:"implementationName"`
	InfoLink            string               `json:"infoLink"`
	Monitor             string               `json:"monitor"`
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

// AddImportList creates an import list in Radarr without testing it.
func (r *Radarr) AddImportList(list *ImportListInput) (*ImportListOutput, error) {
	return r.AddImportListContext(context.Background(), list)
}

// AddImportListContext creates an import list in Radarr without testing it.
func (r *Radarr) AddImportListContext(ctx context.Context, list *ImportListInput) (*ImportListOutput, error) {
	var (
		output ImportListOutput
		body   bytes.Buffer
	)

	list.ID = 0
	if err := json.NewEncoder(&body).Encode(list); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpImportList, err)
	}

	req := starr.Request{URI: bpImportList, Body: &body, Query: url.Values{"forceSave": []string{"true"}}}
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
		req := starr.Request{URI: path.Join(bpImportList, starr.Itoa(id))}
		if err := r.DeleteAny(ctx, req); err != nil {
			errs += fmt.Errorf("api.Delete(%s): %w", &req, err).Error() + " "
		}
	}

	if errs != "" {
		return fmt.Errorf("%w: %s", starr.ErrRequestError, errs)
	}

	return nil
}

// TestImportList tests an import list.
func (r *Radarr) TestImportList(list *ImportListInput) error {
	return r.TestImportListContextt(context.Background(), list)
}

// TestImportListContextt tests an import list.
func (r *Radarr) TestImportListContextt(ctx context.Context, list *ImportListInput) error {
	var output interface{} // any ok

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

// UpdateImportList updates an existing import list and returns the response.
func (r *Radarr) UpdateImportList(list *ImportListInput, force bool) (*ImportListOutput, error) {
	return r.UpdateImportListContext(context.Background(), list, force)
}

// UpdateImportListContext updates an existing import list and returns the response.
func (r *Radarr) UpdateImportListContext(
	ctx context.Context,
	importList *ImportListInput,
	force bool,
) (*ImportListOutput, error) {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(importList); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpImportList, err)
	}

	var output ImportListOutput

	req := starr.Request{
		URI:   path.Join(bpImportList, starr.Itoa(importList.ID)),
		Body:  &body,
		Query: url.Values{"forceSave": []string{starr.Itoa(force)}},
	}
	if err := r.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}
