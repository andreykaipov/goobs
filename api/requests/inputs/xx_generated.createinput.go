// This file has been automatically generated. Don't edit it.

package inputs

// Represents the request body for the CreateInput request.
type CreateInputParams struct {
	// The kind of input to be created
	InputKind string `json:"inputKind,omitempty"`

	// Name of the new input to created
	InputName string `json:"inputName,omitempty"`

	// Settings object to initialize the input with
	InputSettings map[string]interface{} `json:"inputSettings,omitempty"`

	// Whether to set the created scene item to enabled or disabled
	SceneItemEnabled *bool `json:"sceneItemEnabled,omitempty"`

	// Name of the scene to add the input to as a scene item
	SceneName string `json:"sceneName,omitempty"`
}

// Returns the associated request.
func (o *CreateInputParams) GetRequestName() string {
	return "CreateInput"
}

// Represents the response body for the CreateInput request.
type CreateInputResponse struct {
	// ID of the newly created scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`
}

// Creates a new input, adding it as a scene item to the specified scene.
func (c *Client) CreateInput(params *CreateInputParams) (*CreateInputResponse, error) {
	data := &CreateInputResponse{}
	return data, c.SendRequest(params, data)
}
