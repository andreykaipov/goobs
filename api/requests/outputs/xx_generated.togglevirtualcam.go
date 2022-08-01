// This file has been automatically generated. Don't edit it.

package outputs

/*
ToggleVirtualCamParams represents the params body for the "ToggleVirtualCam" request.
Toggles the state of the virtualcam output.
*/
type ToggleVirtualCamParams struct{}

/*
ToggleVirtualCamResponse represents the response body for the "ToggleVirtualCam" request.
Toggles the state of the virtualcam output.
*/
type ToggleVirtualCamResponse struct {
	// Whether the output is active
	OutputActive bool `json:"outputActive,omitempty"`
}

// ToggleVirtualCam sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments
// as this request doesn't require any parameters.
func (c *Client) ToggleVirtualCam(paramss ...*ToggleVirtualCamParams) (*ToggleVirtualCamResponse, error) {
	if len(paramss) == 0 {
		paramss = []*ToggleVirtualCamParams{{}}
	}
	params := paramss[0]
	resp, err := c.SendRequest("ToggleVirtualCam", params)
	if err != nil {
		return nil, err
	}
	return resp.(*ToggleVirtualCamResponse), nil
}
