// This file has been automatically generated. Don't edit it.

package outputs

import requests "github.com/andreykaipov/goobs/api/requests"

/*
StartOutputParams represents the params body for the "StartOutput" request.
Start an output

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#StartOutput.
*/
type StartOutputParams struct {
	requests.ParamsBasic

	// Output name
	OutputName string `json:"outputName"`
}

// Name just returns "StartOutput".
func (o *StartOutputParams) Name() string {
	return "StartOutput"
}

/*
StartOutputResponse represents the response body for the "StartOutput" request.
Start an output

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#StartOutput.
*/
type StartOutputResponse struct {
	requests.ResponseBasic
}

// StartOutput sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) StartOutput(params *StartOutputParams) (*StartOutputResponse, error) {
	data := &StartOutputResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
