package sonarr

import (
	"context"
	"encoding/json"
	"fmt"

	"golift.io/starr"
)

// GetTags returns all the tags.
func (s *Sonarr) GetTags() ([]*starr.Tag, error) {
	return s.GetTagsContext(context.Background())
}

func (s *Sonarr) GetTagsContext(ctx context.Context) ([]*starr.Tag, error) {
	var tags []*starr.Tag

	err := s.GetInto(ctx, "v3/tag", nil, &tags)
	if err != nil {
		return nil, fmt.Errorf("api.Get(tag): %w", err)
	}

	return tags, nil
}

// UpdateTag updates the label for a tag.
func (s *Sonarr) UpdateTag(tagID int, label string) (int, error) {
	return s.UpdateTagContext(context.Background(), tagID, label)
}

func (s *Sonarr) UpdateTagContext(ctx context.Context, tagID int, label string) (int, error) {
	body, err := json.Marshal(&starr.Tag{Label: label, ID: tagID})
	if err != nil {
		return 0, fmt.Errorf("json.Marshal(tag): %w", err)
	}

	var tag starr.Tag
	if err = s.PutInto(ctx, "v3/tag", nil, body, &tag); err != nil {
		return tag.ID, fmt.Errorf("api.Put(tag): %w", err)
	}

	return tag.ID, nil
}

// AddTag adds a tag or returns the ID for an existing tag.
func (s *Sonarr) AddTag(label string) (int, error) {
	return s.AddTagContext(context.Background(), label)
}

func (s *Sonarr) AddTagContext(ctx context.Context, label string) (int, error) {
	body, err := json.Marshal(&starr.Tag{Label: label, ID: 0})
	if err != nil {
		return 0, fmt.Errorf("json.Marshal(tag): %w", err)
	}

	var tag starr.Tag
	if err = s.PostInto(ctx, "v3/tag", nil, body, &tag); err != nil {
		return tag.ID, fmt.Errorf("api.Post(tag): %w", err)
	}

	return tag.ID, nil
}
