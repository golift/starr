package radarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"golift.io/starr"
)

// GetImportLists returns all import lists.
func (r *Radarr) GetImportLists() ([]*ImportList, error) {
	return r.GetImportListsContext(context.Background())
}

// GetImportListsContext returns all import lists.
func (r *Radarr) GetImportListsContext(ctx context.Context) ([]*ImportList, error) {
	var output []*ImportList
	if _, err := r.GetInto(ctx, "v3/importlist", nil, &output); err != nil {
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
	if _, err := r.PostInto(ctx, "v3/importlist", nil, &body, &output); err != nil {
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
		_, err := r.Delete(ctx, "v3/importlist/"+strconv.FormatInt(id, starr.Base10), nil)
		if err != nil {
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

	_, err := r.PutInto(ctx, "v3/importlist/"+strconv.FormatInt(list.ID, starr.Base10), nil, &body, &output)
	if err != nil {
		return nil, fmt.Errorf("api.Put(importlist): %w", err)
	}

	return &output, nil
}
