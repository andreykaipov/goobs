// This file has been automatically generated. Don't edit it.

package scenes

import requests "github.com/andreykaipov/goobs/api/requests"

/*
RemoveSceneParams represents the params body for the "RemoveScene" request.
Removes a scene from OBS.
*/
type RemoveSceneParams struct {
	requests.ParamsBasic

	// Name of the scene to remove
	SceneName string `json:"sceneName,omitempty"`
}

// GetSelfName just returns "RemoveScene".
func (o *RemoveSceneParams) GetSelfName() string {
	return "RemoveScene"
}

/*
RemoveSceneResponse represents the response body for the "RemoveScene" request.
Removes a scene from OBS.
*/
type RemoveSceneResponse struct {
	requests.ResponseBasic
}

// RemoveScene sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) RemoveScene(params *RemoveSceneParams) (*RemoveSceneResponse, error) {
	data := &RemoveSceneResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
