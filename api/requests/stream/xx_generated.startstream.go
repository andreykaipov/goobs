// This file has been automatically generated. Don't edit it.

package stream

import requests "github.com/andreykaipov/goobs/api/requests"

/*
StartStreamParams represents the params body for the "StartStream" request.
Starts the stream output.
*/
type StartStreamParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "StartStream".
func (o *StartStreamParams) GetSelfName() string {
	return "StartStream"
}

/*
StartStreamResponse represents the response body for the "StartStream" request.
Starts the stream output.
*/
type StartStreamResponse struct {
	requests.ResponseBasic
}

// StartStream sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments as
// this request doesn't require any parameters.
func (c *Client) StartStream(paramss ...*StartStreamParams) (*StartStreamResponse, error) {
	if len(paramss) == 0 {
		paramss = []*StartStreamParams{{}}
	}
	params := paramss[0]
	data := &StartStreamResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
