package readarr

import (
	"context"
	"encoding/json"
	"fmt"

	"golift.io/starr"
)

// GetTags returns all the tags.
func (r *Readarr) GetTags() ([]*starr.Tag, error) {
	return r.GetTagsContext(context.Background())
}

func (r *Readarr) GetTagsContext(ctx context.Context) ([]*starr.Tag, error) {
	var tags []*starr.Tag

	err := r.GetInto(ctx, "v1/tag", nil, &tags)
	if err != nil {
		return nil, fmt.Errorf("api.Get(tag): %w", err)
	}

	return tags, nil
}

// UpdateTag updates the label for a tag.
func (r *Readarr) UpdateTag(tagID int, label string) (int, error) {
	return r.UpdateTagContext(context.Background(), tagID, label)
}

func (r *Readarr) UpdateTagContext(ctx context.Context, tagID int, label string) (int, error) {
	body, err := json.Marshal(&starr.Tag{Label: label, ID: tagID})
	if err != nil {
		return 0, fmt.Errorf("json.Marshal(tag): %w", err)
	}

	var tag starr.Tag
	if err = r.PutInto(ctx, "v1/tag", nil, body, &tag); err != nil {
		return tag.ID, fmt.Errorf("api.Put(tag): %w", err)
	}

	return tag.ID, nil
}

// AddTag adds a tag or returns the ID for an existing tag.
func (r *Readarr) AddTag(label string) (int, error) {
	return r.AddTagContext(context.Background(), label)
}

func (r *Readarr) AddTagContext(ctx context.Context, label string) (int, error) {
	body, err := json.Marshal(&starr.Tag{Label: label, ID: 0})
	if err != nil {
		return 0, fmt.Errorf("json.Marshal(tag): %w", err)
	}

	var tag starr.Tag
	if err = r.PostInto(ctx, "v1/tag", nil, body, &tag); err != nil {
		return tag.ID, fmt.Errorf("api.Post(tag): %w", err)
	}

	return tag.ID, nil
}
