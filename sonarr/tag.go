package sonarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"
	"strconv"

	"golift.io/starr"
)

const bpTag = APIver + "/tag"

// GetTags returns all configured tags.
func (s *Sonarr) GetTags() ([]*starr.Tag, error) {
	return s.GetTagsContext(context.Background())
}

func (s *Sonarr) GetTagsContext(ctx context.Context) ([]*starr.Tag, error) {
	var output []*starr.Tag

	if _, err := s.GetInto(ctx, bpTag, nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(tag): %w", err)
	}

	return output, nil
}

// GetTag returns a single tag.
func (s *Sonarr) GetTag(tagID int) (*starr.Tag, error) {
	return s.GetTagContext(context.Background(), tagID)
}

func (s *Sonarr) GetTagContext(ctx context.Context, tagID int) (*starr.Tag, error) {
	var output *starr.Tag

	uri := path.Join(bpTag, strconv.Itoa(tagID))
	if _, err := s.GetInto(ctx, uri, nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(tag): %w", err)
	}

	return output, nil
}

// AddTag creates a tag.
func (s *Sonarr) AddTag(tag *starr.Tag) (*starr.Tag, error) {
	return s.AddTagContext(context.Background(), tag)
}

func (s *Sonarr) AddTagContext(ctx context.Context, tag *starr.Tag) (*starr.Tag, error) {
	var output starr.Tag

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(tag); err != nil {
		return nil, fmt.Errorf("json.Marshal(tag): %w", err)
	}

	if _, err := s.PostInto(ctx, bpTag, nil, &body, &output); err != nil {
		return nil, fmt.Errorf("api.Post(tag): %w", err)
	}

	return &output, nil
}

// UpdateTag updates the tag.
func (s *Sonarr) UpdateTag(tag *starr.Tag) (*starr.Tag, error) {
	return s.UpdateTagContext(context.Background(), tag)
}

func (s *Sonarr) UpdateTagContext(ctx context.Context, tag *starr.Tag) (*starr.Tag, error) {
	var output starr.Tag

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(tag); err != nil {
		return nil, fmt.Errorf("json.Marshal(tag): %w", err)
	}

	uri := path.Join(bpTag, strconv.Itoa(tag.ID))
	if _, err := s.PutInto(ctx, uri, nil, &body, &output); err != nil {
		return nil, fmt.Errorf("api.Put(tag): %w", err)
	}

	return &output, nil
}

// DeleteTag removes a single tag.
func (s *Sonarr) DeleteTag(tagID int) error {
	return s.DeleteTagContext(context.Background(), tagID)
}

func (s *Sonarr) DeleteTagContext(ctx context.Context, tagID int) error {
	uri := path.Join(bpTag, strconv.Itoa(tagID))
	if _, err := s.Delete(ctx, uri, nil); err != nil {
		return fmt.Errorf("api.Delete(tag): %w", err)
	}

	return nil
}
