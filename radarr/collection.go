package radarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"path"

	"golift.io/starr"
)

const bpCollection = APIver + "/collection"

// Collection is the /api/v3/collection resource.
type Collection struct {
	ID                  int            `json:"id,omitempty"`
	Title               string         `json:"title,omitempty"`
	SortTitle           string         `json:"sortTitle,omitempty"`
	TmdbID              int64          `json:"tmdbId,omitempty"`
	Images              []*starr.Image `json:"images,omitempty"`
	Overview            string         `json:"overview,omitempty"`
	Monitored           bool           `json:"monitored"`
	RootFolderPath      string         `json:"rootFolderPath,omitempty"`
	QualityProfileID    int64          `json:"qualityProfileId,omitempty"`
	SearchOnAdd         bool           `json:"searchOnAdd"`
	MinimumAvailability Availability   `json:"minimumAvailability,omitempty"`
	Movies              []*Movie       `json:"movies,omitempty"`
	MissingMovies       int            `json:"missingMovies,omitempty"`
	Tags                []int          `json:"tags,omitempty"`
}

// CollectionUpdate is the body for PUT /collection (bulk update).
type CollectionUpdate struct {
	CollectionIDs       []int           `json:"collectionIds"`
	Monitored           *bool           `json:"monitored,omitempty"`
	QualityProfileID    *int            `json:"qualityProfileId,omitempty"`
	RootFolderPath      string          `json:"rootFolderPath,omitempty"`
	SearchOnAdd         *bool           `json:"searchOnAdd,omitempty"`
	MinimumAvailability Availability    `json:"minimumAvailability,omitempty"`
	Tags                []int           `json:"tags,omitempty"`
	ApplyTags           starr.ApplyTags `json:"applyTags,omitempty"`
}

// GetCollections returns collections, optionally filtered by TMDb id.
func (r *Radarr) GetCollections(tmdbID int64) ([]*Collection, error) {
	return r.GetCollectionsContext(context.Background(), tmdbID)
}

// GetCollectionsContext returns collections, optionally filtered by TMDb id.
func (r *Radarr) GetCollectionsContext(ctx context.Context, tmdbID int64) ([]*Collection, error) {
	params := make(url.Values)
	if tmdbID != 0 {
		params.Set("tmdbId", starr.Str(tmdbID))
	}

	var output []*Collection

	req := starr.Request{URI: bpCollection, Query: params}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetCollection returns a single collection by database id.
func (r *Radarr) GetCollection(collectionID int64) (*Collection, error) {
	return r.GetCollectionContext(context.Background(), collectionID)
}

// GetCollectionContext returns a single collection by database id.
func (r *Radarr) GetCollectionContext(ctx context.Context, collectionID int64) (*Collection, error) {
	var output Collection

	req := starr.Request{URI: path.Join(bpCollection, starr.Str(collectionID))}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateCollections applies a bulk update to collections.
func (r *Radarr) UpdateCollections(update *CollectionUpdate) ([]*Collection, error) {
	return r.UpdateCollectionsContext(context.Background(), update)
}

// UpdateCollectionsContext applies a bulk update to collections.
func (r *Radarr) UpdateCollectionsContext(ctx context.Context, update *CollectionUpdate) ([]*Collection, error) {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(update); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpCollection, err)
	}

	var output []*Collection

	req := starr.Request{URI: bpCollection, Body: &body}
	if err := r.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return output, nil
}

// UpdateCollection updates a single collection.
func (r *Radarr) UpdateCollection(collection *Collection) (*Collection, error) {
	return r.UpdateCollectionContext(context.Background(), collection)
}

// UpdateCollectionContext updates a single collection.
func (r *Radarr) UpdateCollectionContext(ctx context.Context, collection *Collection) (*Collection, error) {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(collection); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpCollection, err)
	}

	var output Collection

	req := starr.Request{URI: path.Join(bpCollection, starr.Str(int64(collection.ID))), Body: &body}
	if err := r.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}
