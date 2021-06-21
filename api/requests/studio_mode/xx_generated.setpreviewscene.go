// This file has been automatically generated. Don't edit it.

package studiomode

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetPreviewSceneParams represents the params body for the "SetPreviewScene" request.
Set the active preview scene.
Will return an `error` if Studio Mode is not enabled.
Since 4.1.0.
*/
type SetPreviewSceneParams struct {
	requests.ParamsBasic

	// The name of the scene to preview.
	SceneName string `json:"scene-name,omitempty"`
}

// GetSelfName just returns "SetPreviewScene".
func (o *SetPreviewSceneParams) GetSelfName() string {
	return "SetPreviewScene"
}

/*
SetPreviewSceneResponse represents the response body for the "SetPreviewScene" request.
Set the active preview scene.
Will return an `error` if Studio Mode is not enabled.
Since v4.1.0.
*/
type SetPreviewSceneResponse struct {
	requests.ResponseBasic
}

// SetPreviewScene sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetPreviewScene(params *SetPreviewSceneParams) (*SetPreviewSceneResponse, error) {
	data := &SetPreviewSceneResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
