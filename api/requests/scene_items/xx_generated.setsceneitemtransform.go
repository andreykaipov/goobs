// This file has been automatically generated. Don't edit it.

package sceneitems

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetSceneItemTransformParams represents the params body for the "SetSceneItemTransform" request.
Sets the transform and crop info of a scene item.
*/
type SetSceneItemTransformParams struct {
	requests.ParamsBasic

	// Numeric ID of the scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`

	// Object containing scene item transform info to update
	SceneItemTransform interface{} `json:"sceneItemTransform,omitempty"`

	// Name of the scene the item is in
	SceneName string `json:"sceneName,omitempty"`
}

// GetSelfName just returns "SetSceneItemTransform".
func (o *SetSceneItemTransformParams) GetSelfName() string {
	return "SetSceneItemTransform"
}

/*
SetSceneItemTransformResponse represents the response body for the "SetSceneItemTransform" request.
Sets the transform and crop info of a scene item.
*/
type SetSceneItemTransformResponse struct {
	requests.ResponseBasic
}

// SetSceneItemTransform sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetSceneItemTransform(params *SetSceneItemTransformParams) (*SetSceneItemTransformResponse, error) {
	data := &SetSceneItemTransformResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
