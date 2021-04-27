// This file has been automatically generated. Don't edit it.

package sceneitems

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetSceneItemPropertiesParams represents the params body for the "GetSceneItemProperties" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSceneItemProperties.
*/
type GetSceneItemPropertiesParams struct {
	requests.Params

	// The name of the source.
	Item string `json:"item"`

	// the name of the scene that the source item belongs to. Defaults to the current scene.
	SceneName string `json:"scene-name"`
}

// GetRequestType returns the RequestType of GetSceneItemPropertiesParams
func (o *GetSceneItemPropertiesParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of GetSceneItemPropertiesParams
func (o *GetSceneItemPropertiesParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on GetSceneItemPropertiesParams
func (o *GetSceneItemPropertiesParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetSceneItemPropertiesResponse represents the response body for the "GetSceneItemProperties" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSceneItemProperties.
*/
type GetSceneItemPropertiesResponse struct {
	requests.Response

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

	// If the source is visible.
	Visible bool `json:"visible"`
}

// GetMessageID returns the MessageID of GetSceneItemPropertiesResponse
func (o *GetSceneItemPropertiesResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of GetSceneItemPropertiesResponse
func (o *GetSceneItemPropertiesResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of GetSceneItemPropertiesResponse
func (o *GetSceneItemPropertiesResponse) GetError() string {
	return o.Error
}

// GetSceneItemProperties sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetSceneItemProperties(
	params *GetSceneItemPropertiesParams,
) (*GetSceneItemPropertiesResponse, error) {
	params.RequestType = "GetSceneItemProperties"
	data := &GetSceneItemPropertiesResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
