// This file has been automatically generated. Don't edit it.

package studiomode

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetPreviewSceneParams represents the params body for the "SetPreviewScene" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetPreviewScene.
*/
type SetPreviewSceneParams struct {
	requests.ParamsBasic

	// The name of the scene to preview.
	SceneName string `json:"scene-name"`
}

// Name just returns "SetPreviewScene".
func (o *SetPreviewSceneParams) Name() string {
	return "SetPreviewScene"
}

/*
SetPreviewSceneResponse represents the response body for the "SetPreviewScene" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetPreviewScene.
*/
type SetPreviewSceneResponse struct {
	requests.ResponseBasic
}

// SetPreviewScene sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetPreviewScene(params *SetPreviewSceneParams) (*SetPreviewSceneResponse, error) {
	data := &SetPreviewSceneResponse{}
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
