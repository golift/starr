package readarr

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
func (r *Readarr) GetTags() ([]*starr.Tag, error) {
	return r.GetTagsContext(context.Background())
}

// GetTagsContext returns all configured tags.
func (r *Readarr) GetTagsContext(ctx context.Context) ([]*starr.Tag, error) {
	var output []*starr.Tag

	req := starr.Request{URI: bpTag}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", req, err)
	}

	return output, nil
}

// GetTag returns a single tag.
func (r *Readarr) GetTag(tagID int) (*starr.Tag, error) {
	return r.GetTagContext(context.Background(), tagID)
}

// GetTagContext returns a single tag.
func (r *Readarr) GetTagContext(ctx context.Context, tagID int) (*starr.Tag, error) {
	var output starr.Tag

	req := starr.Request{URI: path.Join(bpTag, fmt.Sprint(tagID))}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", req, err)
	}

	return &output, nil
}

// AddTag creates a tag.
func (r *Readarr) AddTag(tag *starr.Tag) (*starr.Tag, error) {
	return r.AddTagContext(context.Background(), tag)
}

// AddTagContext creates a tag.
func (r *Readarr) AddTagContext(ctx context.Context, tag *starr.Tag) (*starr.Tag, error) {
	var output starr.Tag

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(tag); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpTag, err)
	}

	req := starr.Request{URI: bpTag, Body: &body}
	if err := r.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", req, err)
	}

	return &output, nil
}

// UpdateTag updates a tag.
func (r *Readarr) UpdateTag(tag *starr.Tag) (*starr.Tag, error) {
	return r.UpdateTagContext(context.Background(), tag)
}

// UpdateTagContext updates a tag.
func (r *Readarr) UpdateTagContext(ctx context.Context, tag *starr.Tag) (*starr.Tag, error) {
	var output starr.Tag

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(tag); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpTag, err)
	}

	req := starr.Request{URI: path.Join(bpTag, fmt.Sprint(tag.ID)), Body: &body}
	if err := r.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", req, err)
	}

	return &output, nil
}

// DeleteTag removes a single tag.
func (r *Readarr) DeleteTag(tagID int) error {
	return r.DeleteTagContext(context.Background(), tagID)
}

// DeleteTagContext removes a single tag.
func (r *Readarr) DeleteTagContext(ctx context.Context, tagID int) error {
	req := starr.Request{URI: path.Join(bpTag, fmt.Sprint(tagID))}
	if err := r.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", req, err)
	}

	return nil
}
