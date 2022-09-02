package radarr

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
func (r *Radarr) GetTags() ([]*starr.Tag, error) {
	return r.GetTagsContext(context.Background())
}

func (r *Radarr) GetTagsContext(ctx context.Context) ([]*starr.Tag, error) {
	var output []*starr.Tag

	if err := r.GetInto(ctx, bpTag, nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(tag): %w", err)
	}

	return output, nil
}

// GetTag returns a single tag.
func (r *Radarr) GetTag(tagID int) (*starr.Tag, error) {
	return r.GetTagContext(context.Background(), tagID)
}

func (r *Radarr) GetTagContext(ctx context.Context, tagID int) (*starr.Tag, error) {
	var output *starr.Tag

	uri := path.Join(bpTag, strconv.Itoa(tagID))
	if err := r.GetInto(ctx, uri, nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(tag): %w", err)
	}

	return output, nil
}

// AddTag creates a tag.
func (r *Radarr) AddTag(tag *starr.Tag) (*starr.Tag, error) {
	return r.AddTagContext(context.Background(), tag)
}

func (r *Radarr) AddTagContext(ctx context.Context, tag *starr.Tag) (*starr.Tag, error) {
	var output starr.Tag

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(tag); err != nil {
		return nil, fmt.Errorf("json.Marshal(tag): %w", err)
	}

	if err := r.PostInto(ctx, bpTag, nil, &body, &output); err != nil {
		return nil, fmt.Errorf("api.Post(tag): %w", err)
	}

	return &output, nil
}

// UpdateTag updates the tag.
func (r *Radarr) UpdateTag(tag *starr.Tag) (*starr.Tag, error) {
	return r.UpdateTagContext(context.Background(), tag)
}

func (r *Radarr) UpdateTagContext(ctx context.Context, tag *starr.Tag) (*starr.Tag, error) {
	var output starr.Tag

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(tag); err != nil {
		return nil, fmt.Errorf("json.Marshal(tag): %w", err)
	}

	uri := path.Join(bpTag, strconv.Itoa(tag.ID))
	if err := r.PutInto(ctx, uri, nil, &body, &output); err != nil {
		return nil, fmt.Errorf("api.Put(tag): %w", err)
	}

	return &output, nil
}

// DeleteTag removes a single tag.
func (r *Radarr) DeleteTag(tagID int) error {
	return r.DeleteTagContext(context.Background(), tagID)
}

func (r *Radarr) DeleteTagContext(ctx context.Context, tagID int) error {
	var output interface{}

	uri := path.Join(bpTag, strconv.Itoa(tagID))
	if err := r.DeleteInto(ctx, uri, nil, &output); err != nil {
		return fmt.Errorf("api.Delete(tag): %w", err)
	}

	return nil
}
