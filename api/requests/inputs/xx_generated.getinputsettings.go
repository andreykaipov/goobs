// This file has been automatically generated. Don't edit it.

package inputs

/*
GetInputSettingsParams represents the params body for the "GetInputSettings" request.
Gets the settings of an input.

Note: Does not include defaults. To create the entire settings object, overlay `inputSettings` over the `defaultInputSettings` provided by `GetInputDefaultSettings`.
*/
type GetInputSettingsParams struct {
	// Name of the input to get the settings of
	InputName string `json:"inputName,omitempty"`
}

/*
GetInputSettingsResponse represents the response body for the "GetInputSettings" request.
Gets the settings of an input.

Note: Does not include defaults. To create the entire settings object, overlay `inputSettings` over the `defaultInputSettings` provided by `GetInputDefaultSettings`.
*/
type GetInputSettingsResponse struct {
	// The kind of the input
	InputKind string `json:"inputKind,omitempty"`

	// Object of settings for the input
	InputSettings interface{} `json:"inputSettings,omitempty"`
}

// GetInputSettings sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetInputSettings(params *GetInputSettingsParams) (*GetInputSettingsResponse, error) {
	resp, err := c.SendRequest("GetInputSettings", params)
	if err != nil {
		return nil, err
	}
	return resp.(*GetInputSettingsResponse), nil
}
