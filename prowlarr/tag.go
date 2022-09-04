package prowlarr

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
func (p *Prowlarr) GetTags() ([]*starr.Tag, error) {
	return p.GetTagsContext(context.Background())
}

func (p *Prowlarr) GetTagsContext(ctx context.Context) ([]*starr.Tag, error) {
	var output []*starr.Tag

	if err := p.GetInto(ctx, bpTag, nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(tag): %w", err)
	}

	return output, nil
}

// GetTag returns a single tag.
func (p *Prowlarr) GetTag(tagID int) (*starr.Tag, error) {
	return p.GetTagContext(context.Background(), tagID)
}

func (p *Prowlarr) GetTagContext(ctx context.Context, tagID int) (*starr.Tag, error) {
	var output *starr.Tag

	uri := path.Join(bpTag, strconv.Itoa(tagID))
	if err := p.GetInto(ctx, uri, nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(tag): %w", err)
	}

	return output, nil
}

// AddTag creates a tag.
func (p *Prowlarr) AddTag(tag *starr.Tag) (*starr.Tag, error) {
	return p.AddTagContext(context.Background(), tag)
}

func (p *Prowlarr) AddTagContext(ctx context.Context, tag *starr.Tag) (*starr.Tag, error) {
	var output starr.Tag

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(tag); err != nil {
		return nil, fmt.Errorf("json.Marshal(tag): %w", err)
	}

	if err := p.PostInto(ctx, bpTag, nil, &body, &output); err != nil {
		return nil, fmt.Errorf("api.Post(tag): %w", err)
	}

	return &output, nil
}

// UpdateTag updates the tag.
func (p *Prowlarr) UpdateTag(tag *starr.Tag) (*starr.Tag, error) {
	return p.UpdateTagContext(context.Background(), tag)
}

func (p *Prowlarr) UpdateTagContext(ctx context.Context, tag *starr.Tag) (*starr.Tag, error) {
	var output starr.Tag

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(tag); err != nil {
		return nil, fmt.Errorf("json.Marshal(tag): %w", err)
	}

	uri := path.Join(bpTag, strconv.Itoa(tag.ID))
	if err := p.PutInto(ctx, uri, nil, &body, &output); err != nil {
		return nil, fmt.Errorf("api.Put(tag): %w", err)
	}

	return &output, nil
}

// DeleteTag removes a single tag.
func (p *Prowlarr) DeleteTag(tagID int) error {
	return p.DeleteTagContext(context.Background(), tagID)
}

func (p *Prowlarr) DeleteTagContext(ctx context.Context, tagID int) error {
	uri := path.Join(bpTag, strconv.Itoa(tagID))
	if err := p.DeleteAny(ctx, uri, nil); err != nil {
		return fmt.Errorf("api.Delete(tag): %w", err)
	}

	return nil
}
