// This file has been automatically generated. Don't edit it.

package scenes

/*
GetSceneSceneTransitionOverrideParams represents the params body for the "GetSceneSceneTransitionOverride" request.
Gets the scene transition overridden for a scene.
*/
type GetSceneSceneTransitionOverrideParams struct {
	// Name of the scene
	SceneName string `json:"sceneName,omitempty"`
}

/*
GetSceneSceneTransitionOverrideResponse represents the response body for the "GetSceneSceneTransitionOverride" request.
Gets the scene transition overridden for a scene.
*/
type GetSceneSceneTransitionOverrideResponse struct {
	// Duration of the overridden scene transition, else `null`
	TransitionDuration float64 `json:"transitionDuration,omitempty"`

	// Name of the overridden scene transition, else `null`
	TransitionName string `json:"transitionName,omitempty"`
}

// GetSceneSceneTransitionOverride sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetSceneSceneTransitionOverride(
	params *GetSceneSceneTransitionOverrideParams,
) (*GetSceneSceneTransitionOverrideResponse, error) {
	resp, err := c.SendRequest("GetSceneSceneTransitionOverride", params)
	if err != nil {
		return nil, err
	}
	return resp.(*GetSceneSceneTransitionOverrideResponse), nil
}
