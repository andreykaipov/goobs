// This file has been automatically generated. Don't edit it.

package sceneitems

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetSceneItemPropertiesParams represents the params body for the "SetSceneItemProperties" request.
Sets the scene specific properties of a source. Unspecified properties will remain unchanged.
Coordinates are relative to the item's parent (the scene or group it belongs to).
Since 4.3.0.
*/
type SetSceneItemPropertiesParams struct {
	requests.ParamsBasic

	Bounds struct {
		// The new alignment of the bounding box. (0-2, 4-6, 8-10)
		Alignment int `json:"alignment"`

		// The new bounds type of the source. Can be "OBS_BOUNDS_STRETCH", "OBS_BOUNDS_SCALE_INNER",
		// "OBS_BOUNDS_SCALE_OUTER", "OBS_BOUNDS_SCALE_TO_WIDTH", "OBS_BOUNDS_SCALE_TO_HEIGHT", "OBS_BOUNDS_MAX_ONLY" or
		// "OBS_BOUNDS_NONE".
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

	Item struct {
		// Scene Item ID (if the `item` field is an object)
		Id int `json:"id"`

		// Scene Item name (if the `item` field is an object)
		Name string `json:"name"`
	} `json:"item"`

	// The new locked status of the source. 'true' keeps it in its current position, 'false' allows movement.
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
		// The new scale filter of the source. Can be "OBS_SCALE_DISABLE", "OBS_SCALE_POINT", "OBS_SCALE_BICUBIC",
		// "OBS_SCALE_BILINEAR", "OBS_SCALE_LANCZOS" or "OBS_SCALE_AREA".
		Filter string `json:"filter"`

		// The new x scale of the item.
		X float64 `json:"x"`

		// The new y scale of the item.
		Y float64 `json:"y"`
	} `json:"scale"`

	// Name of the scene the source item belongs to. Defaults to the current scene.
	SceneName string `json:"scene-name"`

	// The new visibility of the source. 'true' shows source, 'false' hides source.
	Visible bool `json:"visible"`
}

// GetSelfName just returns "SetSceneItemProperties".
func (o *SetSceneItemPropertiesParams) GetSelfName() string {
	return "SetSceneItemProperties"
}

/*
SetSceneItemPropertiesResponse represents the response body for the "SetSceneItemProperties" request.
Sets the scene specific properties of a source. Unspecified properties will remain unchanged.
Coordinates are relative to the item's parent (the scene or group it belongs to).
Since v4.3.0.
*/
type SetSceneItemPropertiesResponse struct {
	requests.ResponseBasic
}

// SetSceneItemProperties sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetSceneItemProperties(params *SetSceneItemPropertiesParams) (*SetSceneItemPropertiesResponse, error) {
	data := &SetSceneItemPropertiesResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
