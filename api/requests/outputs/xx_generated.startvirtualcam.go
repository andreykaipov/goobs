// This file has been automatically generated. Don't edit it.

package outputs

/*
StartVirtualCamParams represents the params body for the "StartVirtualCam" request.
Starts the virtualcam output.
*/
type StartVirtualCamParams struct{}

/*
StartVirtualCamResponse represents the response body for the "StartVirtualCam" request.
Starts the virtualcam output.
*/
type StartVirtualCamResponse struct{}

// StartVirtualCam sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments
// as this request doesn't require any parameters.
func (c *Client) StartVirtualCam(paramss ...*StartVirtualCamParams) (*StartVirtualCamResponse, error) {
	if len(paramss) == 0 {
		paramss = []*StartVirtualCamParams{{}}
	}
	params := paramss[0]
	resp, err := c.SendRequest("StartVirtualCam", params)
	if err != nil {
		return nil, err
	}
	return resp.(*StartVirtualCamResponse), nil
}
