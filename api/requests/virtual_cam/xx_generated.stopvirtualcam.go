// This file has been automatically generated. Don't edit it.

package virtualcam

import requests "github.com/andreykaipov/goobs/api/requests"

/*
StopVirtualCamParams represents the params body for the "StopVirtualCam" request.
Stop virtual cam.
Will return an `error` if virtual cam is not active.
Since 4.9.1.
*/
type StopVirtualCamParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "StopVirtualCam".
func (o *StopVirtualCamParams) GetSelfName() string {
	return "StopVirtualCam"
}

/*
StopVirtualCamResponse represents the response body for the "StopVirtualCam" request.
Stop virtual cam.
Will return an `error` if virtual cam is not active.
Since v4.9.1.
*/
type StopVirtualCamResponse struct {
	requests.ResponseBasic
}

// StopVirtualCam sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments as
// this request doesn't require any parameters.
func (c *Client) StopVirtualCam(paramss ...*StopVirtualCamParams) (*StopVirtualCamResponse, error) {
	if len(paramss) == 0 {
		paramss = []*StopVirtualCamParams{{}}
	}
	params := paramss[0]
	data := &StopVirtualCamResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
