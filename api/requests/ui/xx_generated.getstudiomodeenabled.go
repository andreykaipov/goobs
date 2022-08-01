// This file has been automatically generated. Don't edit it.

package ui

/*
GetStudioModeEnabledParams represents the params body for the "GetStudioModeEnabled" request.
Gets whether studio is enabled.
*/
type GetStudioModeEnabledParams struct{}

/*
GetStudioModeEnabledResponse represents the response body for the "GetStudioModeEnabled" request.
Gets whether studio is enabled.
*/
type GetStudioModeEnabledResponse struct {
	// Whether studio mode is enabled
	StudioModeEnabled bool `json:"studioModeEnabled,omitempty"`
}

// GetStudioModeEnabled sends the corresponding request to the connected OBS WebSockets server. Note the variadic
// arguments as this request doesn't require any parameters.
func (c *Client) GetStudioModeEnabled(paramss ...*GetStudioModeEnabledParams) (*GetStudioModeEnabledResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetStudioModeEnabledParams{{}}
	}
	params := paramss[0]
	resp, err := c.SendRequest("GetStudioModeEnabled", params)
	if err != nil {
		return nil, err
	}
	return resp.(*GetStudioModeEnabledResponse), nil
}
