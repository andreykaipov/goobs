// This file has been automatically generated. Don't edit it.

package inputs

import api "github.com/andreykaipov/goobs/api"

// Represents the request body for the CreateInput request.
type CreateInputParams struct {
	// The kind of input to be created
	InputKind *string `json:"inputKind,omitempty"`

	// Name of the new input to created
	InputName *string `json:"inputName,omitempty"`

	// Settings object to initialize the input with
	InputSettings map[string]any `json:"inputSettings,omitempty"`

	// Whether to set the created scene item to enabled or disabled
	SceneItemEnabled *bool `json:"sceneItemEnabled,omitempty"`

	// Name of the scene to add the input to as a scene item
	SceneName *string `json:"sceneName,omitempty"`
}

func NewCreateInputParams() *CreateInputParams {
	return &CreateInputParams{}
}
func (o *CreateInputParams) WithInputKind(x string) *CreateInputParams {
	o.InputKind = &x
	return o
}
func (o *CreateInputParams) WithInputName(x string) *CreateInputParams {
	o.InputName = &x
	return o
}
func (o *CreateInputParams) WithInputSettings(x map[string]any) *CreateInputParams {
	o.InputSettings = x
	return o
}
func (o *CreateInputParams) WithSceneItemEnabled(x bool) *CreateInputParams {
	o.SceneItemEnabled = &x
	return o
}
func (o *CreateInputParams) WithSceneName(x string) *CreateInputParams {
	o.SceneName = &x
	return o
}

// Returns the associated request.
func (o *CreateInputParams) GetRequestName() string {
	return "CreateInput"
}

// Represents the response body for the CreateInput request.
type CreateInputResponse struct {
	api.ResponseCommon

	// ID of the newly created scene item
	SceneItemId int `json:"sceneItemId,omitempty"`
}

// Creates a new input, adding it as a scene item to the specified scene.
func (c *Client) CreateInput(params *CreateInputParams) (*CreateInputResponse, error) {
	data := &CreateInputResponse{}
	return data, c.SendRequest(params, data)
}
