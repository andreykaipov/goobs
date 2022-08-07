// This file has been automatically generated. Don't edit it.

package sceneitems

import typedefs "github.com/andreykaipov/goobs/api/typedefs"

/*
SetSceneItemTransformParams represents the params body for the "SetSceneItemTransform" request.
Sets the transform and crop info of a scene item.
*/
type SetSceneItemTransformParams struct {
	// Numeric ID of the scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`

	// Scene item transform info
	SceneItemTransform *typedefs.SceneItemTransform `json:"sceneItemTransform,omitempty"`

	// Name of the scene the item is in
	SceneName string `json:"sceneName,omitempty"`
}

/*
SetSceneItemTransformResponse represents the response body for the "SetSceneItemTransform" request.
Sets the transform and crop info of a scene item.
*/
type SetSceneItemTransformResponse struct{}

// SetSceneItemTransform sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetSceneItemTransform(params *SetSceneItemTransformParams) (*SetSceneItemTransformResponse, error) {
	resp, err := c.SendRequest("SetSceneItemTransform", params)
	if err != nil {
		return nil, err
	}
	return resp.(*SetSceneItemTransformResponse), nil
}
