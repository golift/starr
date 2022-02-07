package sonarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strconv"

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
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(&starr.Tag{Label: label, ID: tagID}); err != nil {
		return 0, fmt.Errorf("json.Marshal(tag): %w", err)
	}

	var tag starr.Tag
	if err := s.PutInto(ctx, "v3/tag/"+strconv.Itoa(tagID), nil, &body, &tag); err != nil {
		return tag.ID, fmt.Errorf("api.Put(tag): %w", err)
	}

	return tag.ID, nil
}

// AddTag adds a tag or returns the ID for an existing tag.
func (s *Sonarr) AddTag(label string) (int, error) {
	return s.AddTagContext(context.Background(), label)
}

func (s *Sonarr) AddTagContext(ctx context.Context, label string) (int, error) {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(&starr.Tag{Label: label}); err != nil {
		return 0, fmt.Errorf("json.Marshal(tag): %w", err)
	}

	var tag starr.Tag
	if err := s.PostInto(ctx, "v3/tag", nil, &body, &tag); err != nil {
		return tag.ID, fmt.Errorf("api.Post(tag): %w", err)
	}

	return tag.ID, nil
}

// GetTag returns a single tag.
func (s *Sonarr) GetTag(tagID int) (*starr.Tag, error) {
	return s.GetTagContext(context.Background(), tagID)
}

func (s *Sonarr) GetTagContext(ctx context.Context, tagID int) (*starr.Tag, error) {
	var tag *starr.Tag

	err := s.GetInto(ctx, "v3/tag/"+strconv.Itoa(tagID), nil, &tag)
	if err != nil {
		return nil, fmt.Errorf("api.Get(tag): %w", err)
	}

	return tag, nil
}

// DeleteTag removes a single tag.
func (s *Sonarr) DeleteTag(tagID int) error {
	return s.DeleteTagContext(context.Background(), tagID)
}

func (s *Sonarr) DeleteTagContext(ctx context.Context, tagID int) error {
	_, err := s.Delete(ctx, "v3/tag/"+strconv.Itoa(tagID), nil)
	if err != nil {
		return fmt.Errorf("api.Delete(tag): %w", err)
	}

	return nil
}
