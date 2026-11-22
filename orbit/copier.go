// Package orbit provides functions to modify data structures among the various starr libraries.
// These functions cannot live in the starr library without causing an import cycle.
// These are wrappers around the starr library and other sub modules.
package orbit

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"golift.io/starr/lidarr"
	"golift.io/starr/prowlarr"
	"golift.io/starr/radarr"
	"golift.io/starr/readarr"
	"golift.io/starr/sonarr"
)

var ErrNotPtr = errors.New("must provide a pointer to a non-nil value")

// Copy is an easy way to copy one data structure to another.
func Copy(src, dst any) error {
	if src == nil || reflect.TypeOf(src).Kind() != reflect.Ptr {
		return fmt.Errorf("copy source: %w", ErrNotPtr)
	} else if dst == nil || reflect.TypeOf(dst).Kind() != reflect.Ptr {
		return fmt.Errorf("copy destination: %w", ErrNotPtr)
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(src); err != nil {
		return fmt.Errorf("encoding: %w", err)
	}

	if err := json.NewDecoder(&buf).Decode(dst); err != nil {
		return fmt.Errorf("decoding: %w", err)
	}

	return nil
}

// IndexerInput represents all possible Indexer inputs.
type IndexerInput interface {
	lidarr.IndexerInput | prowlarr.IndexerInput | radarr.IndexerInput |
		readarr.IndexerInput | sonarr.IndexerInput
}

// IndexerOutput represents all possible Indexer outputs.
type IndexerOutput interface {
	lidarr.IndexerOutput | prowlarr.IndexerOutput | radarr.IndexerOutput |
		readarr.IndexerOutput | sonarr.IndexerOutput
}

// CopyIndexers copies a slice of indexers from one type to another, so you may copy them among instances.
// The destination must be a pointer to a slice, so it can be updated in place.
// The destination slice may be empty but the pointer to it must not be nil.
func CopyIndexers[S IndexerInput | IndexerOutput, D IndexerInput](src []*S, dst *[]*D, keepTags bool) ([]*D, error) {
	if dst == nil {
		return nil, ErrNotPtr
	}

	var err error

	for idx, indexer := range src {
		if len(*dst)-1 >= idx { // The destination slice location exists, so update it in place.
			_, err = CopyIndexer(indexer, (*dst)[idx], keepTags)
		} else { // The destination slice is shorter than the source, so append to it.
			newIndexer := new(D)
			newIndexer, err = CopyIndexer(indexer, newIndexer, keepTags)
			*dst = append(*dst, newIndexer) // This happens before checking the error.
		}

		if err != nil {
			break
		}
	}

	return *dst, err
}

// CopyIndexer copies an indexer from one type to another, so you may copy them among instances.
func CopyIndexer[S IndexerInput | IndexerOutput, D IndexerInput](src *S, dst *D, keepTags bool) (*D, error) {
	if err := Copy(src, dst); err != nil {
		return dst, err
	}

	element := reflect.ValueOf(dst).Elem()
	zeroField(element.FieldByName("ID"), true)
	zeroField(element.FieldByName("Tags"), !keepTags)

	return dst, nil
}

func zeroField(field reflect.Value, really bool) {
	if really && field.CanSet() {
		field.SetZero()
	}
}
