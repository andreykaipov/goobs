// This file has been automatically generated. Don't edit it.

package inputs

import requests "github.com/andreykaipov/goobs/api/requests"

/*
CreateInputParams represents the params body for the "CreateInput" request.
Creates a new input, adding it as a scene item to the specified scene.
*/
type CreateInputParams struct {
	requests.ParamsBasic

	// The kind of input to be created
	InputKind string `json:"inputKind,omitempty"`

	// Name of the new input to created
	InputName string `json:"inputName,omitempty"`

	// Settings object to initialize the input with
	InputSettings interface{} `json:"inputSettings,omitempty"`

	// Whether to set the created scene item to enabled or disabled
	SceneItemEnabled bool `json:"sceneItemEnabled,omitempty"`

	// Name of the scene to add the input to as a scene item
	SceneName string `json:"sceneName,omitempty"`
}

// GetSelfName just returns "CreateInput".
func (o *CreateInputParams) GetSelfName() string {
	return "CreateInput"
}

/*
CreateInputResponse represents the response body for the "CreateInput" request.
Creates a new input, adding it as a scene item to the specified scene.
*/
type CreateInputResponse struct {
	requests.ResponseBasic

	// ID of the newly created scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`
}

// CreateInput sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) CreateInput(params *CreateInputParams) (*CreateInputResponse, error) {
	data := &CreateInputResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
