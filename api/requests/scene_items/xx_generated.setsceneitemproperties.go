// This file has been automatically generated. Don't edit it.

package sceneitems

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetSceneItemPropertiesParams represents the params body for the "SetSceneItemProperties" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetSceneItemProperties.
*/
type SetSceneItemPropertiesParams struct {
	requests.Params

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

// GetRequestType returns the RequestType of SetSceneItemPropertiesParams
func (o *SetSceneItemPropertiesParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of SetSceneItemPropertiesParams
func (o *SetSceneItemPropertiesParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on SetSceneItemPropertiesParams
func (o *SetSceneItemPropertiesParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
SetSceneItemPropertiesResponse represents the response body for the "SetSceneItemProperties" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetSceneItemProperties.
*/
type SetSceneItemPropertiesResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of SetSceneItemPropertiesResponse
func (o *SetSceneItemPropertiesResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of SetSceneItemPropertiesResponse
func (o *SetSceneItemPropertiesResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of SetSceneItemPropertiesResponse
func (o *SetSceneItemPropertiesResponse) GetError() string {
	return o.Error
}

// SetSceneItemProperties sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetSceneItemProperties(
	params *SetSceneItemPropertiesParams,
) (*SetSceneItemPropertiesResponse, error) {
	params.RequestType = "SetSceneItemProperties"
	data := &SetSceneItemPropertiesResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
