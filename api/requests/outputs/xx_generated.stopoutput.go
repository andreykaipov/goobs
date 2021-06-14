// This file has been automatically generated. Don't edit it.

package outputs

import requests "github.com/andreykaipov/goobs/api/requests"

/*
StopOutputParams represents the params body for the "StopOutput" request.
Stop an output

Generated from https://github.com/Palakis/obs-websocket/blob/4.7.0/docs/generated/protocol.md#StopOutput.
*/
type StopOutputParams struct {
	requests.ParamsBasic

	// Force stop (default: false)
	Force bool `json:"force"`

	// Output name
	OutputName string `json:"outputName"`
}

// Name just returns "StopOutput".
func (o *StopOutputParams) Name() string {
	return "StopOutput"
}

/*
StopOutputResponse represents the response body for the "StopOutput" request.
Stop an output

Generated from https://github.com/Palakis/obs-websocket/blob/4.7.0/docs/generated/protocol.md#StopOutput.
*/
type StopOutputResponse struct {
	requests.ResponseBasic
}

// StopOutput sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) StopOutput(params *StopOutputParams) (*StopOutputResponse, error) {
	data := &StopOutputResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
