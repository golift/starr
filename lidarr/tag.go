package lidarr

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"golift.io/starr"
)

// GetTags returns all the tags.
func (l *Lidarr) GetTags() ([]*starr.Tag, error) {
	return l.GetTagsContext(context.Background())
}

// GetTagsContext returns all the tags.
func (l *Lidarr) GetTagsContext(ctx context.Context) ([]*starr.Tag, error) {
	var tags []*starr.Tag

	err := l.GetInto(ctx, "v1/tag", nil, &tags)
	if err != nil {
		return nil, fmt.Errorf("api.Get(tag): %w", err)
	}

	return tags, nil
}

// AddTag adds a tag or returns the ID for an existing tag.
func (l *Lidarr) AddTag(label string) (int, error) {
	return l.AddTagContext(context.Background(), label)
}

// AddTagContext adds a tag or returns the ID for an existing tag.
func (l *Lidarr) AddTagContext(ctx context.Context, label string) (int, error) {
	body, err := json.Marshal(&starr.Tag{Label: label, ID: 0})
	if err != nil {
		return 0, fmt.Errorf("json.Marshal(tag): %w", err)
	}

	var tag starr.Tag
	if err = l.PostInto(ctx, "v1/tag", nil, body, &tag); err != nil {
		return tag.ID, fmt.Errorf("api.Post(tag): %w", err)
	}

	return tag.ID, nil
}

// UpdateTag updates the label for a tag.
func (l *Lidarr) UpdateTag(tagID int, label string) (int, error) {
	return l.UpdateTagContext(context.Background(), tagID, label)
}

// UpdateTagContext updates the label for a tag.
func (l *Lidarr) UpdateTagContext(ctx context.Context, tagID int, label string) (int, error) {
	body, err := json.Marshal(&starr.Tag{Label: label, ID: tagID})
	if err != nil {
		return 0, fmt.Errorf("json.Marshal(tag): %w", err)
	}

	var tag starr.Tag
	if err = l.PutInto(ctx, "v1/tag/"+strconv.Itoa(tagID), nil, body, &tag); err != nil {
		return tag.ID, fmt.Errorf("api.Put(tag): %w", err)
	}

	return tag.ID, nil
}

// GetTag returns a single tag.
func (l *Lidarr) GetTag(tagID int) (*starr.Tag, error) {
	return l.GetTagContext(context.Background(), tagID)
}

func (l *Lidarr) GetTagContext(ctx context.Context, tagID int) (*starr.Tag, error) {
	var tag *starr.Tag

	err := l.GetInto(ctx, "v1/tag/"+strconv.Itoa(tagID), nil, &tag)
	if err != nil {
		return nil, fmt.Errorf("api.Get(tag): %w", err)
	}

	return tag, nil
}

// DeleteTag removes a single tag.
func (l *Lidarr) DeleteTag(tagID int) error {
	return l.DeleteTagContext(context.Background(), tagID)
}

func (l *Lidarr) DeleteTagContext(ctx context.Context, tagID int) error {
	_, err := l.Delete(ctx, "v1/tag/"+strconv.Itoa(tagID), nil)
	if err != nil {
		return fmt.Errorf("api.Delete(tag): %w", err)
	}

	return nil
}
