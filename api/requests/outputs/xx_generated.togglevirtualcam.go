// This file has been automatically generated. Don't edit it.

package outputs

import requests "github.com/andreykaipov/goobs/api/requests"

/*
ToggleVirtualCamParams represents the params body for the "ToggleVirtualCam" request.
Toggles the state of the virtualcam output.
*/
type ToggleVirtualCamParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "ToggleVirtualCam".
func (o *ToggleVirtualCamParams) GetSelfName() string {
	return "ToggleVirtualCam"
}

/*
ToggleVirtualCamResponse represents the response body for the "ToggleVirtualCam" request.
Toggles the state of the virtualcam output.
*/
type ToggleVirtualCamResponse struct {
	requests.ResponseBasic

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
	data := &ToggleVirtualCamResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
