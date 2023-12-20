// This file has been automatically generated. Don't edit it.

package outputs

// Represents the request body for the StopVirtualCam request.
type StopVirtualCamParams struct{}

// Returns the associated request.
func (o *StopVirtualCamParams) GetRequestName() string {
	return "StopVirtualCam"
}

// Represents the response body for the StopVirtualCam request.
type StopVirtualCamResponse struct {
	_response
}

// Stops the virtualcam output.
func (c *Client) StopVirtualCam(paramss ...*StopVirtualCamParams) (*StopVirtualCamResponse, error) {
	if len(paramss) == 0 {
		paramss = []*StopVirtualCamParams{{}}
	}
	params := paramss[0]
	data := &StopVirtualCamResponse{}
	return data, c.client.SendRequest(params, data)
}
