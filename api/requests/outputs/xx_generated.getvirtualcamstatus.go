// This file has been automatically generated. Don't edit it.

package outputs

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetVirtualCamStatusParams represents the params body for the "GetVirtualCamStatus" request.
Gets the status of the virtualcam output.
*/
type GetVirtualCamStatusParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "GetVirtualCamStatus".
func (o *GetVirtualCamStatusParams) GetSelfName() string {
	return "GetVirtualCamStatus"
}

/*
GetVirtualCamStatusResponse represents the response body for the "GetVirtualCamStatus" request.
Gets the status of the virtualcam output.
*/
type GetVirtualCamStatusResponse struct {
	requests.ResponseBasic

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
	data := &GetVirtualCamStatusResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
