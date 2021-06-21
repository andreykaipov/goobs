// This file has been automatically generated. Don't edit it.

package typedefs

// SceneItemTransform represents the complex type for SceneItemTransform.
type SceneItemTransform struct {
	// The bounding box of the object (source, scene item, etc).
	Bounds *Bounds `json:"bounds,omitempty"`

	// The crop specification for the object (source, scene item, etc).
	Crop *Crop `json:"crop,omitempty"`

	// List of children (if this item is a group)
	GroupChildren []SceneItemTransform `json:"groupChildren,omitempty"`

	// Scene item height (base source height multiplied by the vertical scaling factor)
	Height float64 `json:"height,omitempty"`

	// If the scene item is locked in position.
	Locked bool `json:"locked"`

	// Name of the item's parent (if this item belongs to a group)
	ParentGroupName string `json:"parentGroupName,omitempty"`

	// The position of the object (source, scene item, etc).
	Position *Position `json:"position,omitempty"`

	// The clockwise rotation of the scene item in degrees around the point of alignment.
	Rotation float64 `json:"rotation,omitempty"`

	// The scaling specification for the object (source, scene item, etc).
	Scale *Scale `json:"scale,omitempty"`

	// Base source (without scaling) of the source
	SourceHeight int `json:"sourceHeight,omitempty"`

	// Base width (without scaling) of the source
	SourceWidth int `json:"sourceWidth,omitempty"`

	// If the scene item is visible.
	Visible bool `json:"visible"`

	// Scene item width (base source width multiplied by the horizontal scaling factor)
	Width float64 `json:"width,omitempty"`
}
