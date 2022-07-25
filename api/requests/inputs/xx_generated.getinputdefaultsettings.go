// This file has been automatically generated. Don't edit it.

package inputs

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetInputDefaultSettingsParams represents the params body for the "GetInputDefaultSettings" request.
Gets the default settings for an input kind.
*/
type GetInputDefaultSettingsParams struct {
	requests.ParamsBasic

	// Input kind to get the default settings for
	InputKind string `json:"inputKind,omitempty"`
}

// GetSelfName just returns "GetInputDefaultSettings".
func (o *GetInputDefaultSettingsParams) GetSelfName() string {
	return "GetInputDefaultSettings"
}

/*
GetInputDefaultSettingsResponse represents the response body for the "GetInputDefaultSettings" request.
Gets the default settings for an input kind.
*/
type GetInputDefaultSettingsResponse struct {
	requests.ResponseBasic

	// Object of default settings for the input kind
	DefaultInputSettings interface{} `json:"defaultInputSettings,omitempty"`
}

// GetInputDefaultSettings sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetInputDefaultSettings(
	params *GetInputDefaultSettingsParams,
) (*GetInputDefaultSettingsResponse, error) {
	data := &GetInputDefaultSettingsResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
