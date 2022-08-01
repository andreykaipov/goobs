// This file has been automatically generated. Don't edit it.

package outputs

/*
GetVirtualCamStatusParams represents the params body for the "GetVirtualCamStatus" request.
Gets the status of the virtualcam output.
*/
type GetVirtualCamStatusParams struct{}

/*
GetVirtualCamStatusResponse represents the response body for the "GetVirtualCamStatus" request.
Gets the status of the virtualcam output.
*/
type GetVirtualCamStatusResponse struct {
	// Whether the output is active
	OutputActive bool `json:"outputActive,omitempty"`
}

// GetVirtualCamStatus sends the corresponding request to the connected OBS WebSockets server. Note the variadic
// arguments as this request doesn't require any parameters.
func (c *Client) GetVirtualCamStatus(paramss ...*GetVirtualCamStatusParams) (*GetVirtualCamStatusResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetVirtualCamStatusParams{{}}
	}
	params := paramss[0]
	resp, err := c.SendRequest("GetVirtualCamStatus", params)
	if err != nil {
		return nil, err
	}
	return resp.(*GetVirtualCamStatusResponse), nil
}
