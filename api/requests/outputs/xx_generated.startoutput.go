// This file has been automatically generated. Don't edit it.

package outputs

import requests "github.com/andreykaipov/goobs/api/requests"

/*
StartOutputParams represents the params body for the "StartOutput" request.
Start an output

Note: Controlling outputs is an experimental feature of obs-websocket. Some plugins which add outputs to OBS may not function properly when they are controlled in this way.
Since 4.7.0.
*/
type StartOutputParams struct {
	requests.ParamsBasic

	// Output name
	OutputName string `json:"outputName,omitempty"`
}

// GetSelfName just returns "StartOutput".
func (o *StartOutputParams) GetSelfName() string {
	return "StartOutput"
}

/*
StartOutputResponse represents the response body for the "StartOutput" request.
Start an output

Note: Controlling outputs is an experimental feature of obs-websocket. Some plugins which add outputs to OBS may not function properly when they are controlled in this way.
Since v4.7.0.
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
