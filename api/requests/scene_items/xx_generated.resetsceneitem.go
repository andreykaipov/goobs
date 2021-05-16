// This file has been automatically generated. Don't edit it.

package sceneitems

import requests "github.com/andreykaipov/goobs/api/requests"

/*
ResetSceneItemParams represents the params body for the "ResetSceneItem" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ResetSceneItem.
*/
type ResetSceneItemParams struct {
	requests.ParamsBasic

	// Name of the source item.
	Item string `json:"item"`

	// Name of the scene the source belongs to. Defaults to the current scene.
	SceneName string `json:"scene-name"`
}

// Name just returns "ResetSceneItem".
func (o *ResetSceneItemParams) Name() string {
	return "ResetSceneItem"
}

/*
ResetSceneItemResponse represents the response body for the "ResetSceneItem" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ResetSceneItem.
*/
type ResetSceneItemResponse struct {
	requests.ResponseBasic
}

// ResetSceneItem sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) ResetSceneItem(params *ResetSceneItemParams) (*ResetSceneItemResponse, error) {
	data := &ResetSceneItemResponse{}
	if err := c.WriteMessage(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
