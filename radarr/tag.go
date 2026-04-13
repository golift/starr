package radarr

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
func (r *Radarr) GetTags() ([]*starr.Tag, error) {
	return r.GetTagsContext(context.Background())
}

// GetTagsContext returns all configured tags.
func (r *Radarr) GetTagsContext(ctx context.Context) ([]*starr.Tag, error) {
	var output []*starr.Tag

	req := starr.Request{URI: bpTag}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetTag returns a single tag.
func (r *Radarr) GetTag(tagID int) (*starr.Tag, error) {
	return r.GetTagContext(context.Background(), tagID)
}

// GetTagContext returns a single tag.
func (r *Radarr) GetTagContext(ctx context.Context, tagID int) (*starr.Tag, error) {
	var output starr.Tag

	req := starr.Request{URI: path.Join(bpTag, starr.Str(tagID))}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// AddTag creates a tag.
func (r *Radarr) AddTag(tag *starr.Tag) (*starr.Tag, error) {
	return r.AddTagContext(context.Background(), tag)
}

// AddTagContext creates a tag.
func (r *Radarr) AddTagContext(ctx context.Context, tag *starr.Tag) (*starr.Tag, error) {
	var output starr.Tag

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(tag); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpTag, err)
	}

	req := starr.Request{URI: bpTag, Body: &body}
	if err := r.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateTag updates a tag.
func (r *Radarr) UpdateTag(tag *starr.Tag) (*starr.Tag, error) {
	return r.UpdateTagContext(context.Background(), tag)
}

// UpdateTagContext updates a tag.
func (r *Radarr) UpdateTagContext(ctx context.Context, tag *starr.Tag) (*starr.Tag, error) {
	var output starr.Tag

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(tag); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpTag, err)
	}

	req := starr.Request{URI: path.Join(bpTag, starr.Str(tag.ID)), Body: &body}
	if err := r.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteTag removes a single tag.
func (r *Radarr) DeleteTag(tagID int) error {
	return r.DeleteTagContext(context.Background(), tagID)
}

// DeleteTagContext removes a single tag.
func (r *Radarr) DeleteTagContext(ctx context.Context, tagID int) error {
	req := starr.Request{URI: path.Join(bpTag, starr.Str(tagID))}
	if err := r.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}

// TagDetails is the /api/v3/tag/detail resource.
type TagDetails struct {
	ID                int    `json:"id"`
	Label             string `json:"label,omitempty"`
	DelayProfileIDs   []int  `json:"delayProfileIds,omitempty"`
	ImportListIDs     []int  `json:"importListIds,omitempty"`
	NotificationIDs   []int  `json:"notificationIds,omitempty"`
	IndexerIDs        []int  `json:"indexerIds,omitempty"`
	DownloadClientIDs []int  `json:"downloadClientIds,omitempty"`
	AutoTagIDs        []int  `json:"autoTagIds,omitempty"`
	MovieIDs          []int  `json:"movieIds,omitempty"`
	IndexerProxyIDs   []int  `json:"indexerProxyIds,omitempty"`
}

// GetTagDetails returns tag usage details for all tags.
func (r *Radarr) GetTagDetails() ([]*TagDetails, error) {
	return r.GetTagDetailsContext(context.Background())
}

// GetTagDetailsContext returns tag usage details for all tags.
func (r *Radarr) GetTagDetailsContext(ctx context.Context) ([]*TagDetails, error) {
	var output []*TagDetails

	req := starr.Request{URI: path.Join(bpTag, "detail")}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetTagDetail returns tag usage details for a single tag.
func (r *Radarr) GetTagDetail(tagID int) (*TagDetails, error) {
	return r.GetTagDetailContext(context.Background(), tagID)
}

// GetTagDetailContext returns tag usage details for a single tag.
func (r *Radarr) GetTagDetailContext(ctx context.Context, tagID int) (*TagDetails, error) {
	var output TagDetails

	req := starr.Request{URI: path.Join(bpTag, "detail", starr.Str(tagID))}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}
