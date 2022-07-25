// This file has been automatically generated. Don't edit it.

package stream

import requests "github.com/andreykaipov/goobs/api/requests"

/*
StopStreamParams represents the params body for the "StopStream" request.
Stops the stream output.
*/
type StopStreamParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "StopStream".
func (o *StopStreamParams) GetSelfName() string {
	return "StopStream"
}

/*
StopStreamResponse represents the response body for the "StopStream" request.
Stops the stream output.
*/
type StopStreamResponse struct {
	requests.ResponseBasic
}

// StopStream sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments as
// this request doesn't require any parameters.
func (c *Client) StopStream(paramss ...*StopStreamParams) (*StopStreamResponse, error) {
	if len(paramss) == 0 {
		paramss = []*StopStreamParams{{}}
	}
	params := paramss[0]
	data := &StopStreamResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
