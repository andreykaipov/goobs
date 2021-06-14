// This file has been automatically generated. Don't edit it.

package sceneitems

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetSceneItemPropertiesParams represents the params body for the "GetSceneItemProperties" request.
Gets the scene specific properties of the specified source item.
Coordinates are relative to the item's parent (the scene or group it belongs to).

Generated from https://github.com/Palakis/obs-websocket/blob/4.7.0/docs/generated/protocol.md#GetSceneItemProperties.
*/
type GetSceneItemPropertiesParams struct {
	requests.ParamsBasic

	// The name of the source.
	Item string `json:"item"`

	// the name of the scene that the source item belongs to. Defaults to the current scene.
	SceneName string `json:"scene-name"`
}

// Name just returns "GetSceneItemProperties".
func (o *GetSceneItemPropertiesParams) Name() string {
	return "GetSceneItemProperties"
}

/*
GetSceneItemPropertiesResponse represents the response body for the "GetSceneItemProperties" request.
Gets the scene specific properties of the specified source item.
Coordinates are relative to the item's parent (the scene or group it belongs to).

Generated from https://github.com/Palakis/obs-websocket/blob/4.7.0/docs/generated/protocol.md#GetSceneItemProperties.
*/
type GetSceneItemPropertiesResponse struct {
	requests.ResponseBasic

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

	// Scene item height (base source height multiplied by the vertical scaling factor)
	Height float64 `json:"height"`

	// If the source's transform is locked.
	Locked bool `json:"locked"`

	// The name of the source.
	Name string `json:"name"`

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
