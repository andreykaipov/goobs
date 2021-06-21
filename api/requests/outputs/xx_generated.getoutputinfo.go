// This file has been automatically generated. Don't edit it.

package outputs

import (
	requests "github.com/andreykaipov/goobs/api/requests"
	typedefs "github.com/andreykaipov/goobs/api/typedefs"
)

/*
GetOutputInfoParams represents the params body for the "GetOutputInfo" request.
Get information about a single output
Since 4.7.0.
*/
type GetOutputInfoParams struct {
	requests.ParamsBasic

	// Output name
	OutputName string `json:"outputName,omitempty"`
}

// GetSelfName just returns "GetOutputInfo".
func (o *GetOutputInfoParams) GetSelfName() string {
	return "GetOutputInfo"
}

/*
GetOutputInfoResponse represents the response body for the "GetOutputInfo" request.
Get information about a single output
Since v4.7.0.
*/
type GetOutputInfoResponse struct {
	requests.ResponseBasic

	// Output info
	OutputInfo typedefs.Output `json:"outputInfo,omitempty"`
}

// GetOutputInfo sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetOutputInfo(params *GetOutputInfoParams) (*GetOutputInfoResponse, error) {
	data := &GetOutputInfoResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
