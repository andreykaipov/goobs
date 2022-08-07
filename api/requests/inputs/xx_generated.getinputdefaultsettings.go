// This file has been automatically generated. Don't edit it.

package inputs

/*
GetInputDefaultSettingsParams represents the params body for the "GetInputDefaultSettings" request.
Gets the default settings for an input kind.
*/
type GetInputDefaultSettingsParams struct {
	// Input kind to get the default settings for
	InputKind string `json:"inputKind,omitempty"`
}

/*
GetInputDefaultSettingsResponse represents the response body for the "GetInputDefaultSettings" request.
Gets the default settings for an input kind.
*/
type GetInputDefaultSettingsResponse struct {
	// Object of default settings for the input kind
	DefaultInputSettings map[string]interface{} `json:"defaultInputSettings,omitempty"`
}

// GetInputDefaultSettings sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetInputDefaultSettings(
	params *GetInputDefaultSettingsParams,
) (*GetInputDefaultSettingsResponse, error) {
	resp, err := c.SendRequest("GetInputDefaultSettings", params)
	if err != nil {
		return nil, err
	}
	return resp.(*GetInputDefaultSettingsResponse), nil
}
