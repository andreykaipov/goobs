// This file has been automatically generated. Don't edit it.

package outputs

import requests "github.com/andreykaipov/goobs/api/requests"

/*
StartVirtualCamParams represents the params body for the "StartVirtualCam" request.
Starts the virtualcam output.
*/
type StartVirtualCamParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "StartVirtualCam".
func (o *StartVirtualCamParams) GetSelfName() string {
	return "StartVirtualCam"
}

/*
StartVirtualCamResponse represents the response body for the "StartVirtualCam" request.
Starts the virtualcam output.
*/
type StartVirtualCamResponse struct {
	requests.ResponseBasic
}

// StartVirtualCam sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments
// as this request doesn't require any parameters.
func (c *Client) StartVirtualCam(paramss ...*StartVirtualCamParams) (*StartVirtualCamResponse, error) {
	if len(paramss) == 0 {
		paramss = []*StartVirtualCamParams{{}}
	}
	params := paramss[0]
	data := &StartVirtualCamResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
