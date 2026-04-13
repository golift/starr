package lidarr

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
func (l *Lidarr) GetTags() ([]*starr.Tag, error) {
	return l.GetTagsContext(context.Background())
}

// GetTagsContext returns all configured tags.
func (l *Lidarr) GetTagsContext(ctx context.Context) ([]*starr.Tag, error) {
	var output []*starr.Tag

	req := starr.Request{URI: bpTag}
	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetTag returns a single tag.
func (l *Lidarr) GetTag(tagID int) (*starr.Tag, error) {
	return l.GetTagContext(context.Background(), tagID)
}

// GetTagContext returns a single tag.
func (l *Lidarr) GetTagContext(ctx context.Context, tagID int) (*starr.Tag, error) {
	var output starr.Tag

	req := starr.Request{URI: path.Join(bpTag, starr.Str(tagID))}
	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// AddTag creates a tag.
func (l *Lidarr) AddTag(tag *starr.Tag) (*starr.Tag, error) {
	return l.AddTagContext(context.Background(), tag)
}

// AddTagContext creates a tag.
func (l *Lidarr) AddTagContext(ctx context.Context, tag *starr.Tag) (*starr.Tag, error) {
	var output starr.Tag

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(tag); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpTag, err)
	}

	req := starr.Request{URI: bpTag, Body: &body}
	if err := l.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateTag updates a tag.
func (l *Lidarr) UpdateTag(tag *starr.Tag) (*starr.Tag, error) {
	return l.UpdateTagContext(context.Background(), tag)
}

// UpdateTagContext updates a tag.
func (l *Lidarr) UpdateTagContext(ctx context.Context, tag *starr.Tag) (*starr.Tag, error) {
	var output starr.Tag

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(tag); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpTag, err)
	}

	req := starr.Request{URI: path.Join(bpTag, starr.Str(tag.ID)), Body: &body}
	if err := l.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteTag removes a single tag.
func (l *Lidarr) DeleteTag(tagID int) error {
	return l.DeleteTagContext(context.Background(), tagID)
}

// DeleteTagContext removes a single tag.
func (l *Lidarr) DeleteTagContext(ctx context.Context, tagID int) error {
	req := starr.Request{URI: path.Join(bpTag, starr.Str(tagID))}
	if err := l.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}

// TagDetails is the /api/v1/tag/detail resource.
type TagDetails struct {
	ID                int    `json:"id"`
	Label             string `json:"label,omitempty"`
	DelayProfileIDs   []int  `json:"delayProfileIds,omitempty"`
	ImportListIDs     []int  `json:"importListIds,omitempty"`
	NotificationIDs   []int  `json:"notificationIds,omitempty"`
	IndexerIDs        []int  `json:"indexerIds,omitempty"`
	DownloadClientIDs []int  `json:"downloadClientIds,omitempty"`
	AutoTagIDs        []int  `json:"autoTagIds,omitempty"`
	ArtistIDs         []int  `json:"artistIds,omitempty"`
	IndexerProxyIDs   []int  `json:"indexerProxyIds,omitempty"`
}

// GetTagDetails returns tag usage details for all tags.
func (l *Lidarr) GetTagDetails() ([]*TagDetails, error) {
	return l.GetTagDetailsContext(context.Background())
}

// GetTagDetailsContext returns tag usage details for all tags.
func (l *Lidarr) GetTagDetailsContext(ctx context.Context) ([]*TagDetails, error) {
	var output []*TagDetails

	req := starr.Request{URI: path.Join(bpTag, "detail")}
	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetTagDetail returns tag usage details for a single tag.
func (l *Lidarr) GetTagDetail(tagID int) (*TagDetails, error) {
	return l.GetTagDetailContext(context.Background(), tagID)
}

// GetTagDetailContext returns tag usage details for a single tag.
func (l *Lidarr) GetTagDetailContext(ctx context.Context, tagID int) (*TagDetails, error) {
	var output TagDetails

	req := starr.Request{URI: path.Join(bpTag, "detail", starr.Str(tagID))}
	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}
