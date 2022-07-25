// This file has been automatically generated. Don't edit it.

package scenes

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetCurrentProgramSceneParams represents the params body for the "SetCurrentProgramScene" request.
Sets the current program scene.
*/
type SetCurrentProgramSceneParams struct {
	requests.ParamsBasic

	// Scene to set as the current program scene
	SceneName string `json:"sceneName,omitempty"`
}

// GetSelfName just returns "SetCurrentProgramScene".
func (o *SetCurrentProgramSceneParams) GetSelfName() string {
	return "SetCurrentProgramScene"
}

/*
SetCurrentProgramSceneResponse represents the response body for the "SetCurrentProgramScene" request.
Sets the current program scene.
*/
type SetCurrentProgramSceneResponse struct {
	requests.ResponseBasic
}

// SetCurrentProgramScene sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetCurrentProgramScene(params *SetCurrentProgramSceneParams) (*SetCurrentProgramSceneResponse, error) {
	data := &SetCurrentProgramSceneResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
