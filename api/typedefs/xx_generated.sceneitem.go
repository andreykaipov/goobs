// This file has been automatically generated. Don't edit it.

package typedefs

/*
SceneItem represents the complex type for SceneItem.

Generated from https://github.com/Palakis/obs-websocket/blob/4.7.0/docs/generated/protocol.md#SceneItem.
*/
type SceneItem struct {
	Cx float64 `json:"cx"`

	Cy float64 `json:"cy"`

	// List of children (if this item is a group)
	GroupChildren []SceneItem `json:"groupChildren"`

	// Scene item ID
	Id int `json:"id"`

	// Whether or not this Scene Item is locked and can't be moved around
	Locked bool `json:"locked"`

	// The name of this Scene Item.
	Name string `json:"name"`

	// Name of the item's parent (if this item belongs to a group)
	ParentGroupName string `json:"parentGroupName"`

	// Whether or not this Scene Item is set to "visible".
	Render bool `json:"render"`

	SourceCx float64 `json:"source_cx"`

	SourceCy float64 `json:"source_cy"`

	// Source type. Value is one of the following: "input", "filter", "transition", "scene" or
	// "unknown"
	Type string `json:"type"`

	Volume float64 `json:"volume"`

	X float64 `json:"x"`

	Y float64 `json:"y"`
}
