// This file has been automatically generated. Don't edit it.

package sceneitems

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetSceneItemPropertiesParams represents the params body for the "SetSceneItemProperties" request.
Sets the scene specific properties of a source. Unspecified properties will remain unchanged.
Coordinates are relative to the item's parent (the scene or group it belongs to).

Generated from https://github.com/Palakis/obs-websocket/blob/4.6.1/docs/generated/protocol.md#SetSceneItemProperties.
*/
type SetSceneItemPropertiesParams struct {
	requests.ParamsBasic

	Bounds struct {
		// The new alignment of the bounding box. (0-2, 4-6, 8-10)
		Alignment int `json:"alignment"`

		// The new bounds type of the source. Can be "OBS_BOUNDS_STRETCH", "OBS_BOUNDS_SCALE_INNER",
		// "OBS_BOUNDS_SCALE_OUTER", "OBS_BOUNDS_SCALE_TO_WIDTH", "OBS_BOUNDS_SCALE_TO_HEIGHT",
		// "OBS_BOUNDS_MAX_ONLY" or "OBS_BOUNDS_NONE".
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

	// The new locked status of the source. 'true' keeps it in its current position, 'false' allows
	// movement.
	Locked bool `json:"locked"`

	Position struct {
		// The new alignment of the source.
		Alignment float64 `json:"alignment"`

		// The new x position of the source.
		X float64 `json:"x"`

		// The new y position of the source.
		Y float64 `json:"y"`
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

// Name just returns "SetSceneItemProperties".
func (o *SetSceneItemPropertiesParams) Name() string {
	return "SetSceneItemProperties"
}

/*
SetSceneItemPropertiesResponse represents the response body for the "SetSceneItemProperties" request.
Sets the scene specific properties of a source. Unspecified properties will remain unchanged.
Coordinates are relative to the item's parent (the scene or group it belongs to).

Generated from https://github.com/Palakis/obs-websocket/blob/4.6.1/docs/generated/protocol.md#SetSceneItemProperties.
*/
type SetSceneItemPropertiesResponse struct {
	requests.ResponseBasic
}

// SetSceneItemProperties sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetSceneItemProperties(
	params *SetSceneItemPropertiesParams,
) (*SetSceneItemPropertiesResponse, error) {
	data := &SetSceneItemPropertiesResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
