// This file has been automatically generated. Don't edit it.

package outputs

/*
StopVirtualCamParams represents the params body for the "StopVirtualCam" request.
Stops the virtualcam output.
*/
type StopVirtualCamParams struct{}

/*
StopVirtualCamResponse represents the response body for the "StopVirtualCam" request.
Stops the virtualcam output.
*/
type StopVirtualCamResponse struct{}

// StopVirtualCam sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments as
// this request doesn't require any parameters.
func (c *Client) StopVirtualCam(paramss ...*StopVirtualCamParams) (*StopVirtualCamResponse, error) {
	if len(paramss) == 0 {
		paramss = []*StopVirtualCamParams{{}}
	}
	params := paramss[0]
	resp, err := c.SendRequest("StopVirtualCam", params)
	if err != nil {
		return nil, err
	}
	return resp.(*StopVirtualCamResponse), nil
}
