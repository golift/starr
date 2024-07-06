package lidarr

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
	QualityProfileID      int64               `json:"qualityProfileId,omitempty"`
	MetadataProfileID     int64               `json:"metadataProfileId,omitempty"`
	ConfigContract        string              `json:"configContract,omitempty"`
	Implementation        string              `json:"implementation,omitempty"`
	ListType              string              `json:"listType,omitempty"`
	MonitorNewItems       string              `json:"monitorNewItems,omitempty"`
	Name                  string              `json:"name,omitempty"`
	RootFolderPath        string              `json:"rootFolderPath,omitempty"`
	ShouldMonitor         string              `json:"shouldMonitor,omitempty"`
	Tags                  []int               `json:"tags,omitempty"`
	Fields                []*starr.FieldInput `json:"fields,omitempty"`
}

// ImportListOutput is the output from the import list methods.
type ImportListOutput struct {
	EnableAutomaticAdd    bool                 `json:"enableAutomaticAdd"`
	ShouldMonitorExisting bool                 `json:"shouldMonitorExisting"`
	ShouldSearch          bool                 `json:"shouldSearch"`
	ListOrder             int                  `json:"listOrder"`
	ID                    int64                `json:"id"`
	QualityProfileID      int64                `json:"qualityProfileId"`
	MetadataProfileID     int64                `json:"metadataProfileId"`
	ShouldMonitor         string               `json:"shouldMonitor"`
	RootFolderPath        string               `json:"rootFolderPath"`
	MonitorNewItems       string               `json:"monitorNewItems"`
	ListType              string               `json:"listType"`
	Name                  string               `json:"name"`
	ImplementationName    string               `json:"implementationName"`
	Implementation        string               `json:"implementation"`
	ConfigContract        string               `json:"configContract"`
	InfoLink              string               `json:"infoLink"`
	Tags                  []int                `json:"tags"`
	Fields                []*starr.FieldOutput `json:"fields"`
	Message               struct {
		Message string `json:"message"` // this is a weird place for a message
		Type    string `json:"type"`
	} `json:"message"`
}

// GetImportLists returns all configured import lists.
func (l *Lidarr) GetImportLists() ([]*ImportListOutput, error) {
	return l.GetImportListsContext(context.Background())
}

// GetImportListsContext returns all configured import lists.
func (l *Lidarr) GetImportListsContext(ctx context.Context) ([]*ImportListOutput, error) {
	var output []*ImportListOutput

	req := starr.Request{URI: bpImportList}
	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetImportList returns a single import list.
func (l *Lidarr) GetImportList(importListID int64) (*ImportListOutput, error) {
	return l.GetImportListContext(context.Background(), importListID)
}

// GetIndGetImportListContextexer returns a single import list.
func (l *Lidarr) GetImportListContext(ctx context.Context, importListID int64) (*ImportListOutput, error) {
	var output ImportListOutput

	req := starr.Request{URI: path.Join(bpImportList, starr.Itoa(importListID))}
	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// AddImportList creates an import list without testing it.
func (l *Lidarr) AddImportList(importList *ImportListInput) (*ImportListOutput, error) {
	return l.AddImportListContext(context.Background(), importList)
}

// AddImportListContext creates an import list without testing it.
func (l *Lidarr) AddImportListContext(ctx context.Context, importList *ImportListInput) (*ImportListOutput, error) {
	var (
		output ImportListOutput
		body   bytes.Buffer
	)

	importList.ID = 0
	if err := json.NewEncoder(&body).Encode(importList); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpImportList, err)
	}

	req := starr.Request{URI: bpImportList, Body: &body, Query: url.Values{"forceSave": []string{"true"}}}
	if err := l.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// TestImportList tests an import list.
func (l *Lidarr) TestImportList(list *ImportListInput) error {
	return l.TestImportListContextt(context.Background(), list)
}

// TestImportListContextt tests an import list.
func (l *Lidarr) TestImportListContextt(ctx context.Context, list *ImportListInput) error {
	var output interface{}

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(list); err != nil {
		return fmt.Errorf("json.Marshal(%s): %w", bpImportList, err)
	}

	req := starr.Request{URI: path.Join(bpImportList, "test"), Body: &body}
	if err := l.PostInto(ctx, req, &output); err != nil {
		return fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return nil
}

// UpdateImportList updates the import list.
func (l *Lidarr) UpdateImportList(importList *ImportListInput, force bool) (*ImportListOutput, error) {
	return l.UpdateImportListContext(context.Background(), importList, force)
}

// UpdateImportListContext updates the import list.
func (l *Lidarr) UpdateImportListContext(
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
		URI:   path.Join(bpImportList, starr.Itoa(importList.ID)),
		Body:  &body,
		Query: url.Values{"forceSave": []string{starr.Itoa(force)}},
	}
	if err := l.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteImportList removes a single import list.
func (l *Lidarr) DeleteImportList(importListID int64) error {
	return l.DeleteImportListContext(context.Background(), importListID)
}

// DeleteImportListContext removes a single import list.
func (l *Lidarr) DeleteImportListContext(ctx context.Context, importListID int64) error {
	req := starr.Request{URI: path.Join(bpImportList, starr.Itoa(importListID))}
	if err := l.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
