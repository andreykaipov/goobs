// This file has been automatically generated. Don't edit it.

package ui

// Represents the request body for the GetStudioModeEnabled request.
type GetStudioModeEnabledParams struct{}

// Returns the associated request.
func (o *GetStudioModeEnabledParams) GetRequestName() string {
	return "GetStudioModeEnabled"
}

// Represents the response body for the GetStudioModeEnabled request.
type GetStudioModeEnabledResponse struct {
	_response

	// Whether studio mode is enabled
	StudioModeEnabled bool `json:"studioModeEnabled,omitempty"`
}

// Gets whether studio is enabled.
func (c *Client) GetStudioModeEnabled(paramss ...*GetStudioModeEnabledParams) (*GetStudioModeEnabledResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetStudioModeEnabledParams{{}}
	}
	params := paramss[0]
	data := &GetStudioModeEnabledResponse{}
	return data, c.client.SendRequest(params, data)
}
