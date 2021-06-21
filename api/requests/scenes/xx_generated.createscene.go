// This file has been automatically generated. Don't edit it.

package scenes

import requests "github.com/andreykaipov/goobs/api/requests"

/*
CreateSceneParams represents the params body for the "CreateScene" request.
Create a new scene scene.
Since 4.9.0.
*/
type CreateSceneParams struct {
	requests.ParamsBasic

	// Name of the scene to create.
	SceneName string `json:"sceneName,omitempty"`
}

// GetSelfName just returns "CreateScene".
func (o *CreateSceneParams) GetSelfName() string {
	return "CreateScene"
}

/*
CreateSceneResponse represents the response body for the "CreateScene" request.
Create a new scene scene.
Since v4.9.0.
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
