// This file has been automatically generated. Don't edit it.

package sceneitems

import typedefs "github.com/andreykaipov/goobs/api/typedefs"

// Represents the request body for the SetSceneItemTransform request.
type SetSceneItemTransformParams struct {
	// Numeric ID of the scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`

	// Scene item transform info
	SceneItemTransform *typedefs.SceneItemTransform `json:"sceneItemTransform,omitempty"`

	// Name of the scene the item is in
	SceneName string `json:"sceneName,omitempty"`
}

// Returns the associated request.
func (o *SetSceneItemTransformParams) GetRequestName() string {
	return "SetSceneItemTransform"
}

// Represents the response body for the SetSceneItemTransform request.
type SetSceneItemTransformResponse struct{}

// Sets the transform and crop info of a scene item.
func (c *Client) SetSceneItemTransform(params *SetSceneItemTransformParams) (*SetSceneItemTransformResponse, error) {
	data := &SetSceneItemTransformResponse{}
	return data, c.SendRequest(params, data)
}
