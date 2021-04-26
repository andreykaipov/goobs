// This file has been automatically generated. Don't edit it.

package sceneitems

// GetSceneItemPropertiesParams contains the request body for the
// [GetSceneItemProperties](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSceneItemProperties)
// request.
type GetSceneItemPropertiesParams struct {
	// The name of the source.
	Item string `json:"item"`

	// the name of the scene that the source item belongs to. Defaults to the current scene.
	SceneName string `json:"scene-name"`
}

// GetSceneItemPropertiesResponse contains the request body for the
// [GetSceneItemProperties](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSceneItemProperties)
// request.
type GetSceneItemPropertiesResponse struct {
	Bounds struct {
		// Alignment of the bounding box.
		Alignment int `json:"alignment"`

		// Type of bounding box.
		Type string `json:"type"`

		// Width of the bounding box.
		X float64 `json:"x"`

		// Height of the bounding box.
		Y float64 `json:"y"`
	} `json:"bounds"`

	Crop struct {
		// The number of pixels cropped off the bottom of the source before scaling.
		Bottom int `json:"bottom"`

		// The number of pixels cropped off the left of the source before scaling.
		Left int `json:"left"`

		// The number of pixels cropped off the right of the source before scaling.
		Right int `json:"right"`

		// The number of pixels cropped off the top of the source before scaling.
		Top int `json:"top"`
	} `json:"crop"`

	// The name of the source.
	Name string `json:"name"`

	Position struct {
		// The point on the source that the item is manipulated from.
		Alignment int `json:"alignment"`

		// The x position of the source from the left.
		X int `json:"x"`

		// The y position of the source from the top.
		Y int `json:"y"`
	} `json:"position"`

	// The clockwise rotation of the item in degrees around the point of alignment.
	Rotation float64 `json:"rotation"`

	Scale struct {
		// The x-scale factor of the source.
		X float64 `json:"x"`

		// The y-scale factor of the source.
		Y float64 `json:"y"`
	} `json:"scale"`

	// If the source is visible.
	Visible bool `json:"visible"`
}

// SetSceneItemPropertiesParams contains the request body for the
// [SetSceneItemProperties](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetSceneItemProperties)
// request.
type SetSceneItemPropertiesParams struct {
	Bounds struct {
		// The new alignment of the bounding box. (0-2, 4-6, 8-10)
		Alignment int `json:"alignment"`

		// The new bounds type of the source.
		Type string `json:"type"`

		// The new width of the bounding box.
		X float64 `json:"x"`

		// The new height of the bounding box.
		Y float64 `json:"y"`
	} `json:"bounds"`

	Crop struct {
		// The new amount of pixels cropped off the bottom of the source before scaling.
		Bottom int `json:"bottom"`

		// The new amount of pixels cropped off the left of the source before scaling.
		Left int `json:"left"`

		// The new amount of pixels cropped off the right of the source before scaling.
		Right int `json:"right"`

		// The new amount of pixels cropped off the top of the source before scaling.
		Top int `json:"top"`
	} `json:"crop"`

	// The name of the source.
	Item string `json:"item"`

	Position struct {
		// The new alignment of the source.
		Alignment int `json:"alignment"`

		// The new x position of the source.
		X int `json:"x"`

		// The new y position of the source.
		Y int `json:"y"`
	} `json:"position"`

	// The new clockwise rotation of the item in degrees.
	Rotation float64 `json:"rotation"`

	Scale struct {
		// The new x scale of the item.
		X float64 `json:"x"`

		// The new y scale of the item.
		Y float64 `json:"y"`
	} `json:"scale"`

	// the name of the scene that the source item belongs to. Defaults to the current scene.
	SceneName string `json:"scene-name"`

	// The new visibility of the source. 'true' shows source, 'false' hides source.
	Visible bool `json:"visible"`
}

// SetSceneItemPropertiesResponse contains the request body for the
// [SetSceneItemProperties](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetSceneItemProperties)
// request.
type SetSceneItemPropertiesResponse struct{}

// ResetSceneItemParams contains the request body for the
// [ResetSceneItem](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ResetSceneItem)
// request.
type ResetSceneItemParams struct {
	// Name of the source item.
	Item string `json:"item"`

	// Name of the scene the source belongs to. Defaults to the current scene.
	SceneName string `json:"scene-name"`
}

// ResetSceneItemResponse contains the request body for the
// [ResetSceneItem](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ResetSceneItem)
// request.
type ResetSceneItemResponse struct{}

// DeleteSceneItemParams contains the request body for the
// [DeleteSceneItem](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#DeleteSceneItem)
// request.
type DeleteSceneItemParams struct {
	Item struct {
		// id of the scene item.
		Id int `json:"id"`

		// name of the scene item (prefer `id`, including both is acceptable).
		Name string `json:"name"`
	} `json:"item"`

	// Name of the scene the source belongs to. Defaults to the current scene.
	Scene string `json:"scene"`
}

// DeleteSceneItemResponse contains the request body for the
// [DeleteSceneItem](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#DeleteSceneItem)
// request.
type DeleteSceneItemResponse struct{}

// DuplicateSceneItemParams contains the request body for the
// [DuplicateSceneItem](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#DuplicateSceneItem)
// request.
type DuplicateSceneItemParams struct {
	// Name of the scene to copy the item from. Defaults to the current scene.
	FromScene string `json:"fromScene"`

	Item struct {
		// id of the scene item.
		Id int `json:"id"`

		// name of the scene item (prefer `id`, including both is acceptable).
		Name string `json:"name"`
	} `json:"item"`

	// Name of the scene to create the item in. Defaults to the current scene.
	ToScene string `json:"toScene"`
}

// DuplicateSceneItemResponse contains the request body for the
// [DuplicateSceneItem](https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#DuplicateSceneItem)
// request.
type DuplicateSceneItemResponse struct {
	Item struct {
		// New item ID
		Id int `json:"id"`

		// New item name
		Name string `json:"name"`
	} `json:"item"`

	// Name of the scene where the new item was created
	Scene string `json:"scene"`
}
