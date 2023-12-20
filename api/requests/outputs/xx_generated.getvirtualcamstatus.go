// This file has been automatically generated. Don't edit it.

package outputs

// Represents the request body for the GetVirtualCamStatus request.
type GetVirtualCamStatusParams struct{}

// Returns the associated request.
func (o *GetVirtualCamStatusParams) GetRequestName() string {
	return "GetVirtualCamStatus"
}

// Represents the response body for the GetVirtualCamStatus request.
type GetVirtualCamStatusResponse struct {
	_response

	// Whether the output is active
	OutputActive bool `json:"outputActive,omitempty"`
}

// Gets the status of the virtualcam output.
func (c *Client) GetVirtualCamStatus(paramss ...*GetVirtualCamStatusParams) (*GetVirtualCamStatusResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetVirtualCamStatusParams{{}}
	}
	params := paramss[0]
	data := &GetVirtualCamStatusResponse{}
	return data, c.client.SendRequest(params, data)
}
