// This file has been automatically generated. Don't edit it.

package sceneitems

import (
	requests "github.com/andreykaipov/goobs/api/requests"
	typedefs "github.com/andreykaipov/goobs/api/typedefs"
)

/*
SetSceneItemPropertiesParams represents the params body for the "SetSceneItemProperties" request.
Sets the scene specific properties of a source. Unspecified properties will remain unchanged.
Coordinates are relative to the item's parent (the scene or group it belongs to).
Since 4.3.0.
*/
type SetSceneItemPropertiesParams struct {
	requests.ParamsBasic

	// The bounding box of the object (source, scene item, etc).
	Bounds *typedefs.Bounds `json:"bounds,omitempty"`

	// The crop specification for the object (source, scene item, etc).
	Crop *typedefs.Crop `json:"crop,omitempty"`

	// The item specification for this object.
	Item *typedefs.Item `json:"item,omitempty"`

	// The new locked status of the source. 'true' keeps it in its current position, 'false' allows movement.
	Locked bool `json:"locked"`

	// The position of the object (source, scene item, etc).
	Position *typedefs.Position `json:"position,omitempty"`

	// The new clockwise rotation of the item in degrees.
	Rotation float64 `json:"rotation,omitempty"`

	// The scaling specification for the object (source, scene item, etc).
	Scale *typedefs.Scale `json:"scale,omitempty"`

	// Name of the scene the source item belongs to. Defaults to the current scene.
	SceneName string `json:"scene-name,omitempty"`

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
