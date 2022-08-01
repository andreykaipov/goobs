// This file has been automatically generated. Don't edit it.

package scenes

/*
SetSceneSceneTransitionOverrideParams represents the params body for the "SetSceneSceneTransitionOverride" request.
Gets the scene transition overridden for a scene.
*/
type SetSceneSceneTransitionOverrideParams struct {
	// Name of the scene
	SceneName string `json:"sceneName,omitempty"`

	// Duration to use for any overridden transition. Specify `null` to remove
	TransitionDuration float64 `json:"transitionDuration,omitempty"`

	// Name of the scene transition to use as override. Specify `null` to remove
	TransitionName string `json:"transitionName,omitempty"`
}

/*
SetSceneSceneTransitionOverrideResponse represents the response body for the "SetSceneSceneTransitionOverride" request.
Gets the scene transition overridden for a scene.
*/
type SetSceneSceneTransitionOverrideResponse struct{}

// SetSceneSceneTransitionOverride sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetSceneSceneTransitionOverride(
	params *SetSceneSceneTransitionOverrideParams,
) (*SetSceneSceneTransitionOverrideResponse, error) {
	resp, err := c.SendRequest("SetSceneSceneTransitionOverride", params)
	if err != nil {
		return nil, err
	}
	return resp.(*SetSceneSceneTransitionOverrideResponse), nil
}
