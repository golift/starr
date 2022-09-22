package sonarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"

	"golift.io/starr"
)

const bpTag = APIver + "/tag"

// GetTags returns all configured tags.
func (s *Sonarr) GetTags() ([]*starr.Tag, error) {
	return s.GetTagsContext(context.Background())
}

// GetTagsContext returns all configured tags.
func (s *Sonarr) GetTagsContext(ctx context.Context) ([]*starr.Tag, error) {
	var output []*starr.Tag

	req := starr.Request{URI: bpTag}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetTag returns a single tag.
func (s *Sonarr) GetTag(tagID int) (*starr.Tag, error) {
	return s.GetTagContext(context.Background(), tagID)
}

// GetTagContext returns a single tag.
func (s *Sonarr) GetTagContext(ctx context.Context, tagID int) (*starr.Tag, error) {
	var output starr.Tag

	req := starr.Request{URI: path.Join(bpTag, fmt.Sprint(tagID))}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// AddTag creates a tag.
func (s *Sonarr) AddTag(tag *starr.Tag) (*starr.Tag, error) {
	return s.AddTagContext(context.Background(), tag)
}

// AddTagContext creates a tag.
func (s *Sonarr) AddTagContext(ctx context.Context, tag *starr.Tag) (*starr.Tag, error) {
	var output starr.Tag

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(tag); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpTag, err)
	}

	req := starr.Request{URI: bpTag, Body: &body}
	if err := s.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateTag updates the tag.
func (s *Sonarr) UpdateTag(tag *starr.Tag) (*starr.Tag, error) {
	return s.UpdateTagContext(context.Background(), tag)
}

// UpdateTagContext updates the tag.
func (s *Sonarr) UpdateTagContext(ctx context.Context, tag *starr.Tag) (*starr.Tag, error) {
	var output starr.Tag

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(tag); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpTag, err)
	}

	req := starr.Request{URI: path.Join(bpTag, fmt.Sprint(tag.ID)), Body: &body}
	if err := s.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteTag removes a single tag.
func (s *Sonarr) DeleteTag(tagID int) error {
	return s.DeleteTagContext(context.Background(), tagID)
}

// DeleteTagContext removes a single tag.
func (s *Sonarr) DeleteTagContext(ctx context.Context, tagID int) error {
	req := starr.Request{URI: path.Join(bpTag, fmt.Sprint(tagID))}
	if err := s.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
