// This file has been automatically generated. Don't edit it.

package outputs

// Represents the request body for the ToggleVirtualCam request.
type ToggleVirtualCamParams struct{}

// Returns the associated request.
func (o *ToggleVirtualCamParams) GetRequestName() string {
	return "ToggleVirtualCam"
}

// Represents the response body for the ToggleVirtualCam request.
type ToggleVirtualCamResponse struct {
	_response

	// Whether the output is active
	OutputActive bool `json:"outputActive,omitempty"`
}

// Toggles the state of the virtualcam output.
func (c *Client) ToggleVirtualCam(paramss ...*ToggleVirtualCamParams) (*ToggleVirtualCamResponse, error) {
	if len(paramss) == 0 {
		paramss = []*ToggleVirtualCamParams{{}}
	}
	params := paramss[0]
	data := &ToggleVirtualCamResponse{}
	return data, c.client.SendRequest(params, data)
}
