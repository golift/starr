package radarr

import (
	"encoding/json"
	"fmt"
	"strconv"

	"golift.io/starr"
)

// GetImportLists returns all import lists.
func (r *Radarr) GetImportLists() ([]*ImportList, error) {
	var output []*ImportList
	if err := r.GetInto("v3/importlist", nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(importlist): %w", err)
	}

	return output, nil
}

// CreateImportList creates an import list in Radarr.
func (r *Radarr) CreateImportList(il *ImportList) (*ImportList, error) {
	il.ID = 0

	body, err := json.Marshal(il)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal(importlist): %w", err)
	}

	var output ImportList
	if err := r.PostInto("v3/importlist", nil, body, &output); err != nil {
		return nil, fmt.Errorf("api.Post(importlist): %w", err)
	}

	return &output, nil
}

// DeleteImportList removes an import list from Radarr.
func (r *Radarr) DeleteImportList(ids []int64) error {
	var errs string

	for _, id := range ids {
		_, err := r.Delete("v3/importlist/"+strconv.FormatInt(id, starr.Base10), nil)
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
	body, err := json.Marshal(list)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal(importlist): %w", err)
	}

	var output ImportList

	err = r.PutInto("v3/importlist/"+strconv.FormatInt(list.ID, starr.Base10), nil, body, &output)
	if err != nil {
		return nil, fmt.Errorf("api.Put(importlist): %w", err)
	}

	return &output, nil
}
