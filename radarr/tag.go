package radarr

import (
	"context"
	"encoding/json"
	"fmt"

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

// UpdateTag updates the label for a tag.
func (r *Radarr) UpdateTagContext(ctx context.Context, tagID int, label string) (int, error) {
	body, err := json.Marshal(&starr.Tag{Label: label, ID: tagID})
	if err != nil {
		return 0, fmt.Errorf("json.Marshal(tag): %w", err)
	}

	var tag starr.Tag
	if err = r.PutInto(ctx, "v3/tag", nil, body, &tag); err != nil {
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
	body, err := json.Marshal(&starr.Tag{Label: label, ID: 0})
	if err != nil {
		return 0, fmt.Errorf("json.Marshal(tag): %w", err)
	}

	var tag starr.Tag
	if err = r.PostInto(ctx, "v3/tag", nil, body, &tag); err != nil {
		return tag.ID, fmt.Errorf("api.Post(tag): %w", err)
	}

	return tag.ID, nil
}
