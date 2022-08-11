// This file has been automatically generated. Don't edit it.

package sceneitems

import typedefs "github.com/andreykaipov/goobs/api/typedefs"

// Represents the request body for the GetSceneItemTransform request.
type GetSceneItemTransformParams struct {
	// Numeric ID of the scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`

	// Name of the scene the item is in
	SceneName string `json:"sceneName,omitempty"`
}

// Returns the associated request.
func (o *GetSceneItemTransformParams) GetRequestName() string {
	return "GetSceneItemTransform"
}

// Represents the response body for the GetSceneItemTransform request.
type GetSceneItemTransformResponse struct {
	// Scene item transform info
	SceneItemTransform *typedefs.SceneItemTransform `json:"sceneItemTransform,omitempty"`
}

/*
Gets the transform and crop info of a scene item.

Scenes and Groups
*/
func (c *Client) GetSceneItemTransform(params *GetSceneItemTransformParams) (*GetSceneItemTransformResponse, error) {
	data := &GetSceneItemTransformResponse{}
	return data, c.SendRequest(params, data)
}
