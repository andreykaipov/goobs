// This file has been automatically generated. Don't edit it.

package typedefs

/*
SceneItemTransform represents the complex type for SceneItemTransform.

Generated from https://github.com/Palakis/obs-websocket/blob/4.7.0/docs/generated/protocol.md#SceneItemTransform.
*/
type SceneItemTransform struct {
	Bounds struct {
		// Alignment of the bounding box.
		Alignment int `json:"alignment"`

		// Type of bounding box. Can be "OBS_BOUNDS_STRETCH", "OBS_BOUNDS_SCALE_INNER",
		// "OBS_BOUNDS_SCALE_OUTER", "OBS_BOUNDS_SCALE_TO_WIDTH", "OBS_BOUNDS_SCALE_TO_HEIGHT",
		// "OBS_BOUNDS_MAX_ONLY" or "OBS_BOUNDS_NONE".
		Type string `json:"type"`

		// Width of the bounding box.
		X float64 `json:"x"`

		// Height of the bounding box.
		Y float64 `json:"y"`
	} `json:"bounds"`

	Crop struct {
		// The number of pixels cropped off the bottom of the scene item before scaling.
		Bottom int `json:"bottom"`

		// The number of pixels cropped off the left of the scene item before scaling.
		Left int `json:"left"`

		// The number of pixels cropped off the right of the scene item before scaling.
		Right int `json:"right"`

		// The number of pixels cropped off the top of the scene item before scaling.
		Top int `json:"top"`
	} `json:"crop"`

	// List of children (if this item is a group)
	GroupChildren []SceneItemTransform `json:"groupChildren"`

	// Scene item height (base source height multiplied by the vertical scaling factor)
	Height float64 `json:"height"`

	// If the scene item is locked in position.
	Locked bool `json:"locked"`

	// Name of the item's parent (if this item belongs to a group)
	ParentGroupName string `json:"parentGroupName"`

	Position struct {
		// The point on the scene item that the item is manipulated from.
		Alignment float64 `json:"alignment"`

		// The x position of the scene item from the left.
		X float64 `json:"x"`

		// The y position of the scene item from the top.
		Y float64 `json:"y"`
	} `json:"position"`

	// The clockwise rotation of the scene item in degrees around the point of alignment.
	Rotation float64 `json:"rotation"`

	Scale struct {
		// The x-scale factor of the scene item.
		X float64 `json:"x"`

		// The y-scale factor of the scene item.
		Y float64 `json:"y"`
	} `json:"scale"`

	// Base source (without scaling) of the source
	SourceHeight int `json:"sourceHeight"`

	// Base width (without scaling) of the source
	SourceWidth int `json:"sourceWidth"`

	// If the scene item is visible.
	Visible bool `json:"visible"`

	// Scene item width (base source width multiplied by the horizontal scaling factor)
	Width float64 `json:"width"`
}
