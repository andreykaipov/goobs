// This file has been automatically generated. Don't edit it.

package virtualcam

import requests "github.com/andreykaipov/goobs/api/requests"

/*
StartStopVirtualCamParams represents the params body for the "StartStopVirtualCam" request.
Toggle virtual cam on or off (depending on the current virtual cam state).
Since 4.9.1.
*/
type StartStopVirtualCamParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "StartStopVirtualCam".
func (o *StartStopVirtualCamParams) GetSelfName() string {
	return "StartStopVirtualCam"
}

/*
StartStopVirtualCamResponse represents the response body for the "StartStopVirtualCam" request.
Toggle virtual cam on or off (depending on the current virtual cam state).
Since v4.9.1.
*/
type StartStopVirtualCamResponse struct {
	requests.ResponseBasic
}

// StartStopVirtualCam sends the corresponding request to the connected OBS WebSockets server. Note the variadic
// arguments as this request doesn't require any parameters.
func (c *Client) StartStopVirtualCam(paramss ...*StartStopVirtualCamParams) (*StartStopVirtualCamResponse, error) {
	if len(paramss) == 0 {
		paramss = []*StartStopVirtualCamParams{{}}
	}
	params := paramss[0]
	data := &StartStopVirtualCamResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
