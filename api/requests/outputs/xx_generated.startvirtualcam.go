// This file has been automatically generated. Don't edit it.

package outputs

// Represents the request body for the StartVirtualCam request.
type StartVirtualCamParams struct{}

// Returns the associated request.
func (o *StartVirtualCamParams) GetRequestName() string {
	return "StartVirtualCam"
}

// Represents the response body for the StartVirtualCam request.
type StartVirtualCamResponse struct {
	_response
}

// Starts the virtualcam output.
func (c *Client) StartVirtualCam(paramss ...*StartVirtualCamParams) (*StartVirtualCamResponse, error) {
	if len(paramss) == 0 {
		paramss = []*StartVirtualCamParams{{}}
	}
	params := paramss[0]
	data := &StartVirtualCamResponse{}
	return data, c.client.SendRequest(params, data)
}
