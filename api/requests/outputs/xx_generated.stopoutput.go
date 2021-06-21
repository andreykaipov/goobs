// This file has been automatically generated. Don't edit it.

package outputs

import requests "github.com/andreykaipov/goobs/api/requests"

/*
StopOutputParams represents the params body for the "StopOutput" request.
Stop an output

Note: Controlling outputs is an experimental feature of obs-websocket. Some plugins which add outputs to OBS may not function properly when they are controlled in this way.
Since 4.7.0.
*/
type StopOutputParams struct {
	requests.ParamsBasic

	// Force stop (default: false)
	Force bool `json:"force"`

	// Output name
	OutputName string `json:"outputName,omitempty"`
}

// GetSelfName just returns "StopOutput".
func (o *StopOutputParams) GetSelfName() string {
	return "StopOutput"
}

/*
StopOutputResponse represents the response body for the "StopOutput" request.
Stop an output

Note: Controlling outputs is an experimental feature of obs-websocket. Some plugins which add outputs to OBS may not function properly when they are controlled in this way.
Since v4.7.0.
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
