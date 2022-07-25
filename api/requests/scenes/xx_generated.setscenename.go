// This file has been automatically generated. Don't edit it.

package scenes

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetSceneNameParams represents the params body for the "SetSceneName" request.
Sets the name of a scene (rename).
*/
type SetSceneNameParams struct {
	requests.ParamsBasic

	// New name for the scene
	NewSceneName string `json:"newSceneName,omitempty"`

	// Name of the scene to be renamed
	SceneName string `json:"sceneName,omitempty"`
}

// GetSelfName just returns "SetSceneName".
func (o *SetSceneNameParams) GetSelfName() string {
	return "SetSceneName"
}

/*
SetSceneNameResponse represents the response body for the "SetSceneName" request.
Sets the name of a scene (rename).
*/
type SetSceneNameResponse struct {
	requests.ResponseBasic
}

// SetSceneName sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetSceneName(params *SetSceneNameParams) (*SetSceneNameResponse, error) {
	data := &SetSceneNameResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
