// This file has been automatically generated. Don't edit it.

package inputs

// Represents the request body for the GetInputDefaultSettings request.
type GetInputDefaultSettingsParams struct {
	// Input kind to get the default settings for
	InputKind *string `json:"inputKind,omitempty"`
}

func NewGetInputDefaultSettingsParams() *GetInputDefaultSettingsParams {
	return &GetInputDefaultSettingsParams{}
}
func (o *GetInputDefaultSettingsParams) WithInputKind(x string) *GetInputDefaultSettingsParams {
	o.InputKind = &x
	return o
}

// Returns the associated request.
func (o *GetInputDefaultSettingsParams) GetRequestName() string {
	return "GetInputDefaultSettings"
}

// Represents the response body for the GetInputDefaultSettings request.
type GetInputDefaultSettingsResponse struct {
	_response

	// Object of default settings for the input kind
	DefaultInputSettings map[string]any `json:"defaultInputSettings,omitempty"`
}

// Gets the default settings for an input kind.
func (c *Client) GetInputDefaultSettings(
	params *GetInputDefaultSettingsParams,
) (*GetInputDefaultSettingsResponse, error) {
	data := &GetInputDefaultSettingsResponse{}
	return data, c.client.SendRequest(params, data)
}
