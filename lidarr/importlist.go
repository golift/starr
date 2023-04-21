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
	QualityProfileID      int64               `json:"qualityProfileId"`
	MetadataProfileID     int64               `json:"metadataProfileId"`
	ConfigContract        string              `json:"configContract"`
	Implementation        string              `json:"implementation"`
	ListType              string              `json:"listType"`
	MonitorNewItems       string              `json:"monitorNewItems"`
	Name                  string              `json:"name"`
	RootFolderPath        string              `json:"rootFolderPath"`
	ShouldMonitor         string              `json:"shouldMonitor"`
	Tags                  []int               `json:"tags"`
	Fields                []*starr.FieldInput `json:"fields"`
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

	req := starr.Request{URI: path.Join(bpImportList, fmt.Sprint(importListID))}
	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// AddImportList creates a import list.
func (l *Lidarr) AddImportList(importList *ImportListInput) (*ImportListOutput, error) {
	return l.AddImportListContext(context.Background(), importList)
}

// AddImportListContext creates a import list.
func (l *Lidarr) AddImportListContext(ctx context.Context, importList *ImportListInput) (*ImportListOutput, error) {
	var output ImportListOutput

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(importList); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpImportList, err)
	}

	req := starr.Request{URI: bpImportList, Body: &body}
	if err := l.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
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
		URI:   path.Join(bpImportList, fmt.Sprint(importList.ID)),
		Body:  &body,
		Query: url.Values{"forceSave": []string{fmt.Sprint(force)}},
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
	req := starr.Request{URI: path.Join(bpImportList, fmt.Sprint(importListID))}
	if err := l.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
