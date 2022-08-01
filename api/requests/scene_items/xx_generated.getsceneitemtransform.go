// This file has been automatically generated. Don't edit it.

package sceneitems

/*
GetSceneItemTransformParams represents the params body for the "GetSceneItemTransform" request.
Gets the transform and crop info of a scene item.

Scenes and Groups
*/
type GetSceneItemTransformParams struct {
	// Numeric ID of the scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`

	// Name of the scene the item is in
	SceneName string `json:"sceneName,omitempty"`
}

/*
GetSceneItemTransformResponse represents the response body for the "GetSceneItemTransform" request.
Gets the transform and crop info of a scene item.

Scenes and Groups
*/
type GetSceneItemTransformResponse struct {
	// Object containing scene item transform info
	SceneItemTransform interface{} `json:"sceneItemTransform,omitempty"`
}

// GetSceneItemTransform sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetSceneItemTransform(params *GetSceneItemTransformParams) (*GetSceneItemTransformResponse, error) {
	resp, err := c.SendRequest("GetSceneItemTransform", params)
	if err != nil {
		return nil, err
	}
	return resp.(*GetSceneItemTransformResponse), nil
}
