package starrshared

import "golift.io/starr"

// AutoTagging is the auto tagging API resource shared by Sonarr, Lidarr, Readarr, and Radarr.
type AutoTagging struct {
	ID                      int                         `json:"id,omitempty"`
	Name                    string                      `json:"name,omitempty"`
	RemoveTagsAutomatically bool                        `json:"removeTagsAutomatically"`
	Tags                    []int                       `json:"tags,omitempty"`
	Specifications          []*AutoTaggingSpecification `json:"specifications,omitempty"`
}

// AutoTaggingSpecification is one rule inside an AutoTagging definition.
type AutoTaggingSpecification struct {
	ID                 int                 `json:"id,omitempty"`
	Name               string              `json:"name,omitempty"`
	Implementation     string              `json:"implementation,omitempty"`
	ImplementationName string              `json:"implementationName,omitempty"`
	Negate             bool                `json:"negate"`
	Required           bool                `json:"required"`
	Fields             []*starr.FieldInput `json:"fields,omitempty"`
}
