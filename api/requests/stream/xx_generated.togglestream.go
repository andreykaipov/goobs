// This file has been automatically generated. Don't edit it.

package stream

import requests "github.com/andreykaipov/goobs/api/requests"

/*
ToggleStreamParams represents the params body for the "ToggleStream" request.
Toggles the status of the stream output.
*/
type ToggleStreamParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "ToggleStream".
func (o *ToggleStreamParams) GetSelfName() string {
	return "ToggleStream"
}

/*
ToggleStreamResponse represents the response body for the "ToggleStream" request.
Toggles the status of the stream output.
*/
type ToggleStreamResponse struct {
	requests.ResponseBasic

	// New state of the stream output
	OutputActive bool `json:"outputActive,omitempty"`
}

// ToggleStream sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments as
// this request doesn't require any parameters.
func (c *Client) ToggleStream(paramss ...*ToggleStreamParams) (*ToggleStreamResponse, error) {
	if len(paramss) == 0 {
		paramss = []*ToggleStreamParams{{}}
	}
	params := paramss[0]
	data := &ToggleStreamResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
