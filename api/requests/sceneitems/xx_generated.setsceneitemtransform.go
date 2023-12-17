// This file has been automatically generated. Don't edit it.

package sceneitems

import typedefs "github.com/andreykaipov/goobs/api/typedefs"

// Represents the request body for the SetSceneItemTransform request.
type SetSceneItemTransformParams struct {
	// Numeric ID of the scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`

	// Object containing scene item transform info to update
	SceneItemTransform *typedefs.SceneItemTransform `json:"sceneItemTransform,omitempty"`

	// Name of the scene the item is in
	SceneName string `json:"sceneName,omitempty"`
}

func NewSetSceneItemTransformParams() *SetSceneItemTransformParams {
	return &SetSceneItemTransformParams{}
}
func (o *SetSceneItemTransformParams) WithSceneItemId(x float64) *SetSceneItemTransformParams {
	o.SceneItemId = x
	return o
}

func (o *SetSceneItemTransformParams) WithSceneItemTransform(
	x *typedefs.SceneItemTransform,
) *SetSceneItemTransformParams {
	o.SceneItemTransform = x
	return o
}
func (o *SetSceneItemTransformParams) WithSceneName(x string) *SetSceneItemTransformParams {
	o.SceneName = x
	return o
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
