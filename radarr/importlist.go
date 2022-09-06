package radarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"
	"strconv"

	"golift.io/starr"
)

// ImportList represents the api/v3/importlist endpoint.
type ImportList struct {
	ID                  int64    `json:"id"`
	Name                string   `json:"name"`
	Enabled             bool     `json:"enabled"`
	EnableAuto          bool     `json:"enableAuto"`
	ShouldMonitor       bool     `json:"shouldMonitor"`
	SearchOnAdd         bool     `json:"searchOnAdd"`
	RootFolderPath      string   `json:"rootFolderPath"`
	QualityProfileID    int64    `json:"qualityProfileId"`
	MinimumAvailability string   `json:"minimumAvailability"`
	ListType            string   `json:"listType"`
	ListOrder           int64    `json:"listOrder"`
	Fields              []*Field `json:"fields"`
	ImplementationName  string   `json:"implementationName"`
	Implementation      string   `json:"implementation"`
	ConfigContract      string   `json:"configContract"`
	InfoLink            string   `json:"infoLink"`
	Tags                []int    `json:"tags"`
}

// Field is currently only part of ImportList.
type Field struct {
	Name          string          `json:"name"`
	Value         interface{}     `json:"value"` // sometimes number, sometimes string. 'Type' may tell you.
	Label         string          `json:"label"`
	HelpText      string          `json:"helpText"`
	Type          string          `json:"type"`
	Order         int64           `json:"order"`
	Advanced      bool            `json:"advanced"`
	SelectOptions []*SelectOption `json:"selectOptions,omitempty"`
}

// SelectOption is part of a Field from an ImportList.
type SelectOption struct {
	Value        int    `json:"value"`
	Name         string `json:"name"`
	Order        int    `json:"order"`
	DividerAfter bool   `json:"dividerAfter"`
}

// GetImportLists returns all import lists.
func (r *Radarr) GetImportLists() ([]*ImportList, error) {
	return r.GetImportListsContext(context.Background())
}

// GetImportListsContext returns all import lists.
func (r *Radarr) GetImportListsContext(ctx context.Context) ([]*ImportList, error) {
	var output []*ImportList
	if err := r.GetInto(ctx, "v3/importlist", nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(importlist): %w", err)
	}

	return output, nil
}

// CreateImportList creates an import list in Radarr.
func (r *Radarr) CreateImportList(il *ImportList) (*ImportList, error) {
	return r.CreateImportListContext(context.Background(), il)
}

// CreateImportListContext creates an import list in Radarr.
func (r *Radarr) CreateImportListContext(ctx context.Context, list *ImportList) (*ImportList, error) {
	list.ID = 0

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(list); err != nil {
		return nil, fmt.Errorf("json.Marshal(list): %w", err)
	}

	var output ImportList
	if err := r.PostInto(ctx, "v3/importlist", nil, &body, &output); err != nil {
		return nil, fmt.Errorf("api.Post(importlist): %w", err)
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
		req := &starr.Request{URI: path.Join("v3/importlist/", fmt.Sprint(id))}
		if err := r.DeleteAny(ctx, req); err != nil {
			errs += fmt.Errorf("api.Delete(importlist): %w", err).Error() + " "
		}
	}

	if errs != "" {
		return fmt.Errorf("%w: %s", starr.ErrRequestError, errs)
	}

	return nil
}

// UpdateImportList updates an existing import list and returns the response.
func (r *Radarr) UpdateImportList(list *ImportList) (*ImportList, error) {
	return r.UpdateImportListContext(context.Background(), list)
}

// UpdateImportListContext updates an existing import list and returns the response.
func (r *Radarr) UpdateImportListContext(ctx context.Context, list *ImportList) (*ImportList, error) {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(list); err != nil {
		return nil, fmt.Errorf("json.Marshal(list): %w", err)
	}

	var output ImportList

	err := r.PutInto(ctx, "v3/importlist/"+strconv.FormatInt(list.ID, starr.Base10), nil, &body, &output)
	if err != nil {
		return nil, fmt.Errorf("api.Put(importlist): %w", err)
	}

	return &output, nil
}
