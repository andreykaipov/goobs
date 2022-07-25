// This file has been automatically generated. Don't edit it.

package scenes

import requests "github.com/andreykaipov/goobs/api/requests"

/*
CreateSceneParams represents the params body for the "CreateScene" request.
Creates a new scene in OBS.
*/
type CreateSceneParams struct {
	requests.ParamsBasic

	// Name for the new scene
	SceneName string `json:"sceneName,omitempty"`
}

// GetSelfName just returns "CreateScene".
func (o *CreateSceneParams) GetSelfName() string {
	return "CreateScene"
}

/*
CreateSceneResponse represents the response body for the "CreateScene" request.
Creates a new scene in OBS.
*/
type CreateSceneResponse struct {
	requests.ResponseBasic
}

// CreateScene sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) CreateScene(params *CreateSceneParams) (*CreateSceneResponse, error) {
	data := &CreateSceneResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
