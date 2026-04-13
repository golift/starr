package starrshared

import "encoding/json"

// CustomFilter is the /customfilter resource (UI saved filters).
type CustomFilter struct {
	ID      int               `json:"id,omitempty"`
	Type    string            `json:"type,omitempty"`
	Label   string            `json:"label,omitempty"`
	Filters []json.RawMessage `json:"filters,omitempty"`
}
