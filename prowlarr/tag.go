package prowlarr

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
func (p *Prowlarr) GetTags() ([]*starr.Tag, error) {
	return p.GetTagsContext(context.Background())
}

// GetTagsContext returns all configured tags.
func (p *Prowlarr) GetTagsContext(ctx context.Context) ([]*starr.Tag, error) {
	var output []*starr.Tag

	req := starr.Request{URI: bpTag}
	if err := p.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetTag returns a single tag.
func (p *Prowlarr) GetTag(tagID int) (*starr.Tag, error) {
	return p.GetTagContext(context.Background(), tagID)
}

// GetTagContext returns a single tag.
func (p *Prowlarr) GetTagContext(ctx context.Context, tagID int) (*starr.Tag, error) {
	var output starr.Tag

	req := starr.Request{URI: path.Join(bpTag, fmt.Sprint(tagID))}
	if err := p.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// AddTag creates a tag.
func (p *Prowlarr) AddTag(tag *starr.Tag) (*starr.Tag, error) {
	return p.AddTagContext(context.Background(), tag)
}

// AddTagContext creates a tag.
func (p *Prowlarr) AddTagContext(ctx context.Context, tag *starr.Tag) (*starr.Tag, error) {
	var output starr.Tag

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(tag); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpTag, err)
	}

	req := starr.Request{URI: bpTag, Body: &body}
	if err := p.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateTag updates a tag.
func (p *Prowlarr) UpdateTag(tag *starr.Tag) (*starr.Tag, error) {
	return p.UpdateTagContext(context.Background(), tag)
}

// UpdateTagContext updates a tag.
func (p *Prowlarr) UpdateTagContext(ctx context.Context, tag *starr.Tag) (*starr.Tag, error) {
	var output starr.Tag

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(tag); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpTag, err)
	}

	req := starr.Request{URI: path.Join(bpTag, fmt.Sprint(tag.ID)), Body: &body}
	if err := p.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteTag removes a single tag.
func (p *Prowlarr) DeleteTag(tagID int) error {
	return p.DeleteTagContext(context.Background(), tagID)
}

func (p *Prowlarr) DeleteTagContext(ctx context.Context, tagID int) error {
	req := starr.Request{URI: path.Join(bpTag, fmt.Sprint(tagID))}
	if err := p.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
