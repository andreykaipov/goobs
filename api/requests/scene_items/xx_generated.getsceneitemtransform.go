// This file has been automatically generated. Don't edit it.

package sceneitems

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetSceneItemTransformParams represents the params body for the "GetSceneItemTransform" request.
Gets the transform and crop info of a scene item.

Scenes and Groups
*/
type GetSceneItemTransformParams struct {
	requests.ParamsBasic

	// Numeric ID of the scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`

	// Name of the scene the item is in
	SceneName string `json:"sceneName,omitempty"`
}

// GetSelfName just returns "GetSceneItemTransform".
func (o *GetSceneItemTransformParams) GetSelfName() string {
	return "GetSceneItemTransform"
}

/*
GetSceneItemTransformResponse represents the response body for the "GetSceneItemTransform" request.
Gets the transform and crop info of a scene item.

Scenes and Groups
*/
type GetSceneItemTransformResponse struct {
	requests.ResponseBasic

	// Object containing scene item transform info
	SceneItemTransform interface{} `json:"sceneItemTransform,omitempty"`
}

// GetSceneItemTransform sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetSceneItemTransform(params *GetSceneItemTransformParams) (*GetSceneItemTransformResponse, error) {
	data := &GetSceneItemTransformResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
