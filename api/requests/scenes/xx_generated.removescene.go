// This file has been automatically generated. Don't edit it.

package scenes

// Represents the request body for the RemoveScene request.
type RemoveSceneParams struct {
	// Name of the scene to remove
	SceneName *string `json:"sceneName,omitempty"`

	// UUID of the scene to remove
	SceneUuid *string `json:"sceneUuid,omitempty"`
}

func NewRemoveSceneParams() *RemoveSceneParams {
	return &RemoveSceneParams{}
}
func (o *RemoveSceneParams) WithSceneName(x string) *RemoveSceneParams {
	o.SceneName = &x
	return o
}
func (o *RemoveSceneParams) WithSceneUuid(x string) *RemoveSceneParams {
	o.SceneUuid = &x
	return o
}

// Returns the associated request.
func (o *RemoveSceneParams) GetRequestName() string {
	return "RemoveScene"
}

// Represents the response body for the RemoveScene request.
type RemoveSceneResponse struct {
	_response
}

// Removes a scene from OBS.
func (c *Client) RemoveScene(paramss ...*RemoveSceneParams) (*RemoveSceneResponse, error) {
	if len(paramss) == 0 {
		paramss = []*RemoveSceneParams{{}}
	}
	params := paramss[0]
	data := &RemoveSceneResponse{}
	return data, c.client.SendRequest(params, data)
}
