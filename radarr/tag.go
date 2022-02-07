package radarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"golift.io/starr"
)

// GetTags returns all the tags.
func (r *Radarr) GetTags() ([]*starr.Tag, error) {
	return r.GetTagsContext(context.Background())
}

// GetTagsContext returns all the tags.
func (r *Radarr) GetTagsContext(ctx context.Context) ([]*starr.Tag, error) {
	var tags []*starr.Tag

	err := r.GetInto(ctx, "v3/tag", nil, &tags)
	if err != nil {
		return nil, fmt.Errorf("api.Get(tag): %w", err)
	}

	return tags, nil
}

// UpdateTag updates the label for a tag.
func (r *Radarr) UpdateTag(tagID int, label string) (int, error) {
	return r.UpdateTagContext(context.Background(), tagID, label)
}

// UpdateTagContext updates the label for a tag.
func (r *Radarr) UpdateTagContext(ctx context.Context, tagID int, label string) (int, error) {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(&starr.Tag{Label: label, ID: tagID}); err != nil {
		return 0, fmt.Errorf("json.Marshal(tag): %w", err)
	}

	var tag starr.Tag
	if err := r.PutInto(ctx, "v3/tag/"+strconv.Itoa(tagID), nil, &body, &tag); err != nil {
		return tag.ID, fmt.Errorf("api.Put(tag): %w", err)
	}

	return tag.ID, nil
}

// AddTag adds a tag or returns the ID for an existing tag.
func (r *Radarr) AddTag(label string) (int, error) {
	return r.AddTagContext(context.Background(), label)
}

// AddTagContext adds a tag or returns the ID for an existing tag.
func (r *Radarr) AddTagContext(ctx context.Context, label string) (int, error) {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(&starr.Tag{Label: label}); err != nil {
		return 0, fmt.Errorf("json.Marshal(tag): %w", err)
	}

	var tag starr.Tag
	if err := r.PostInto(ctx, "v3/tag", nil, &body, &tag); err != nil {
		return tag.ID, fmt.Errorf("api.Post(tag): %w", err)
	}

	return tag.ID, nil
}

// GetTag returns a single tag.
func (r *Radarr) GetTag(tagID int) (*starr.Tag, error) {
	return r.GetTagContext(context.Background(), tagID)
}

func (r *Radarr) GetTagContext(ctx context.Context, tagID int) (*starr.Tag, error) {
	var tag *starr.Tag

	err := r.GetInto(ctx, "v3/tag"+strconv.Itoa(tagID), nil, &tag)
	if err != nil {
		return nil, fmt.Errorf("api.Get(tag): %w", err)
	}

	return tag, nil
}

// DeleteTag removes a single tag.
func (r *Radarr) DeleteTag(tagID int) error {
	return r.DeleteTagContext(context.Background(), tagID)
}

func (r *Radarr) DeleteTagContext(ctx context.Context, tagID int) error {
	_, err := r.Delete(ctx, "v3/tag"+strconv.Itoa(tagID), nil)
	if err != nil {
		return fmt.Errorf("api.Get(tag): %w", err)
	}

	return nil
}
