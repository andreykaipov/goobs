// This file has been automatically generated. Don't edit it.

package typedefs

// SceneItem represents the complex type for SceneItem.
type SceneItem struct {
	// The point on the source that the item is manipulated from. The sum of 1=Left or 2=Right, and 4=Top or 8=Bottom,
	// or omit to center on that axis.
	Alignment float64 `json:"alignment,omitempty"`

	Cx float64 `json:"cx,omitempty"`

	Cy float64 `json:"cy,omitempty"`

	// List of children (if this item is a group)
	GroupChildren []SceneItem `json:"groupChildren,omitempty"`

	// Scene item ID
	Id int `json:"id,omitempty"`

	// Whether or not this Scene Item is locked and can't be moved around
	Locked bool `json:"locked"`

	// Whether or not this Scene Item is muted.
	Muted bool `json:"muted"`

	// The name of this Scene Item.
	Name string `json:"name,omitempty"`

	// Name of the item's parent (if this item belongs to a group)
	ParentGroupName string `json:"parentGroupName,omitempty"`

	// Whether or not this Scene Item is set to "visible".
	Render bool `json:"render"`

	SourceCx float64 `json:"source_cx,omitempty"`

	SourceCy float64 `json:"source_cy,omitempty"`

	// Source type. Value is one of the following: "input", "filter", "transition", "scene" or "unknown"
	Type string `json:"type,omitempty"`

	Volume float64 `json:"volume,omitempty"`

	X float64 `json:"x,omitempty"`

	Y float64 `json:"y,omitempty"`
}
