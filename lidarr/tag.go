package lidarr

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
func (l *Lidarr) GetTags() ([]*starr.Tag, error) {
	return l.GetTagsContext(context.Background())
}

func (l *Lidarr) GetTagsContext(ctx context.Context) ([]*starr.Tag, error) {
	var output []*starr.Tag

	if _, err := l.GetInto(ctx, bpTag, nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(tag): %w", err)
	}

	return output, nil
}

// GetTag returns a single tag.
func (l *Lidarr) GetTag(tagID int) (*starr.Tag, error) {
	return l.GetTagContext(context.Background(), tagID)
}

func (l *Lidarr) GetTagContext(ctx context.Context, tagID int) (*starr.Tag, error) {
	var output *starr.Tag

	uri := path.Join(bpTag, strconv.Itoa(tagID))
	if _, err := l.GetInto(ctx, uri, nil, &output); err != nil {
		return nil, fmt.Errorf("api.Get(tag): %w", err)
	}

	return output, nil
}

// AddTag creates a tag.
func (l *Lidarr) AddTag(tag *starr.Tag) (*starr.Tag, error) {
	return l.AddTagContext(context.Background(), tag)
}

func (l *Lidarr) AddTagContext(ctx context.Context, tag *starr.Tag) (*starr.Tag, error) {
	var output starr.Tag

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(tag); err != nil {
		return nil, fmt.Errorf("json.Marshal(tag): %w", err)
	}

	if _, err := l.PostInto(ctx, bpTag, nil, &body, &output); err != nil {
		return nil, fmt.Errorf("api.Post(tag): %w", err)
	}

	return &output, nil
}

// UpdateTag updates the tag.
func (l *Lidarr) UpdateTag(tag *starr.Tag) (*starr.Tag, error) {
	return l.UpdateTagContext(context.Background(), tag)
}

func (l *Lidarr) UpdateTagContext(ctx context.Context, tag *starr.Tag) (*starr.Tag, error) {
	var output starr.Tag

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(tag); err != nil {
		return nil, fmt.Errorf("json.Marshal(tag): %w", err)
	}

	uri := path.Join(bpTag, strconv.Itoa(tag.ID))
	if _, err := l.PutInto(ctx, uri, nil, &body, &output); err != nil {
		return nil, fmt.Errorf("api.Put(tag): %w", err)
	}

	return &output, nil
}

// DeleteTag removes a single tag.
func (l *Lidarr) DeleteTag(tagID int) error {
	return l.DeleteTagContext(context.Background(), tagID)
}

func (l *Lidarr) DeleteTagContext(ctx context.Context, tagID int) error {
	uri := path.Join(bpTag, strconv.Itoa(tagID))
	if _, err := l.Delete(ctx, uri, nil); err != nil {
		return fmt.Errorf("api.Delete(tag): %w", err)
	}

	return nil
}
