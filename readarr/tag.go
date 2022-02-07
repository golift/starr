package readarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strconv"

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
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(&starr.Tag{Label: label, ID: tagID}); err != nil {
		return 0, fmt.Errorf("json.Marshal(tag): %w", err)
	}

	var tag starr.Tag
	if err := r.PutInto(ctx, "v1/tag/"+strconv.Itoa(tagID), nil, &body, &tag); err != nil {
		return tag.ID, fmt.Errorf("api.Put(tag): %w", err)
	}

	return tag.ID, nil
}

// AddTag adds a tag or returns the ID for an existing tag.
func (r *Readarr) AddTag(label string) (int, error) {
	return r.AddTagContext(context.Background(), label)
}

func (r *Readarr) AddTagContext(ctx context.Context, label string) (int, error) {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(&starr.Tag{Label: label}); err != nil {
		return 0, fmt.Errorf("json.Marshal(tag): %w", err)
	}

	var tag starr.Tag
	if err := r.PostInto(ctx, "v1/tag", nil, &body, &tag); err != nil {
		return tag.ID, fmt.Errorf("api.Post(tag): %w", err)
	}

	return tag.ID, nil
}

// GetTag returns a single tag.
func (r *Readarr) GetTag(tagID int) (*starr.Tag, error) {
	return r.GetTagContext(context.Background(), tagID)
}

func (r *Readarr) GetTagContext(ctx context.Context, tagID int) (*starr.Tag, error) {
	var tag *starr.Tag

	err := r.GetInto(ctx, "v1/tag/"+strconv.Itoa(tagID), nil, &tag)
	if err != nil {
		return nil, fmt.Errorf("api.Get(tag): %w", err)
	}

	return tag, nil
}

// DeleteTag removes a single tag.
func (r *Readarr) DeleteTag(tagID int) error {
	return r.DeleteTagContext(context.Background(), tagID)
}

func (r *Readarr) DeleteTagContext(ctx context.Context, tagID int) error {
	_, err := r.Delete(ctx, "v1/tag/"+strconv.Itoa(tagID), nil)
	if err != nil {
		return fmt.Errorf("api.Delete(tag): %w", err)
	}

	return nil
}
