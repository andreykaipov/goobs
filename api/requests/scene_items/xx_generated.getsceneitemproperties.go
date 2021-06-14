// This file has been automatically generated. Don't edit it.

package sceneitems

import (
	requests "github.com/andreykaipov/goobs/api/requests"
	typedefs "github.com/andreykaipov/goobs/api/typedefs"
)

/*
GetSceneItemPropertiesParams represents the params body for the "GetSceneItemProperties" request.
Gets the scene specific properties of the specified source item.
Coordinates are relative to the item's parent (the scene or group it belongs to).

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#GetSceneItemProperties.
*/
type GetSceneItemPropertiesParams struct {
	requests.ParamsBasic

	Item struct {
		// Scene Item ID (if the `item` field is an object)
		Id int `json:"id"`

		// Scene Item name (if the `item` field is an object)
		Name string `json:"name"`
	} `json:"item"`

	// Name of the scene the scene item belongs to. Defaults to the current scene.
	SceneName string `json:"scene-name"`
}

// GetSelfName just returns "GetSceneItemProperties".
func (o *GetSceneItemPropertiesParams) GetSelfName() string {
	return "GetSceneItemProperties"
}

/*
GetSceneItemPropertiesResponse represents the response body for the "GetSceneItemProperties" request.
Gets the scene specific properties of the specified source item.
Coordinates are relative to the item's parent (the scene or group it belongs to).

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#GetSceneItemProperties.
*/
type GetSceneItemPropertiesResponse struct {
	requests.ResponseBasic

	// The point on the source that the item is manipulated from. The sum of 1=Left or 2=Right, and
	// 4=Top or 8=Bottom, or omit to center on that axis.
	Alignment int `json:"alignment"`

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
		// The number of pixels cropped off the bottom of the source before scaling.
		Bottom int `json:"bottom"`

		// The number of pixels cropped off the left of the source before scaling.
		Left int `json:"left"`

		// The number of pixels cropped off the right of the source before scaling.
		Right int `json:"right"`

		// The number of pixels cropped off the top of the source before scaling.
		Top int `json:"top"`
	} `json:"crop"`

	// List of children (if this item is a group)
	GroupChildren []typedefs.SceneItemTransform `json:"groupChildren"`

	// Scene item height (base source height multiplied by the vertical scaling factor)
	Height float64 `json:"height"`

	// Scene Item ID.
	ItemId int `json:"itemId"`

	// If the source's transform is locked.
	Locked bool `json:"locked"`

	// If the source is muted.
	Muted bool `json:"muted"`

	// Scene Item name.
	Name string `json:"name"`

	// Name of the item's parent (if this item belongs to a group)
	ParentGroupName string `json:"parentGroupName"`

	Position struct {
		// The point on the source that the item is manipulated from.
		Alignment float64 `json:"alignment"`

		// The x position of the source from the left.
		X float64 `json:"x"`

		// The y position of the source from the top.
		Y float64 `json:"y"`
	} `json:"position"`

	// The clockwise rotation of the item in degrees around the point of alignment.
	Rotation float64 `json:"rotation"`

	Scale struct {
		// The x-scale factor of the source.
		X float64 `json:"x"`

		// The y-scale factor of the source.
		Y float64 `json:"y"`
	} `json:"scale"`

	// Base source (without scaling) of the source
	SourceHeight int `json:"sourceHeight"`

	// Base width (without scaling) of the source
	SourceWidth int `json:"sourceWidth"`

	// If the source is visible.
	Visible bool `json:"visible"`

	// Scene item width (base source width multiplied by the horizontal scaling factor)
	Width float64 `json:"width"`
}

// GetSceneItemProperties sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetSceneItemProperties(
	params *GetSceneItemPropertiesParams,
) (*GetSceneItemPropertiesResponse, error) {
	data := &GetSceneItemPropertiesResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
