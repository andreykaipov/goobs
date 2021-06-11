// This file has been automatically generated. Don't edit it.

package streaming

import requests "github.com/andreykaipov/goobs/api/requests"

/*
StopStreamingParams represents the params body for the "StopStreaming" request.
Stop streaming.
Will return an `error` if streaming is not active.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.1/docs/generated/protocol.md#StopStreaming.
*/
type StopStreamingParams struct {
	requests.ParamsBasic
}

// Name just returns "StopStreaming".
func (o *StopStreamingParams) Name() string {
	return "StopStreaming"
}

/*
StopStreamingResponse represents the response body for the "StopStreaming" request.
Stop streaming.
Will return an `error` if streaming is not active.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.1/docs/generated/protocol.md#StopStreaming.
*/
type StopStreamingResponse struct {
	requests.ResponseBasic
}

// StopStreaming sends the corresponding request to the connected OBS WebSockets server. Note the
// variadic arguments as this request doesn't require any parameters.
func (c *Client) StopStreaming(paramss ...*StopStreamingParams) (*StopStreamingResponse, error) {
	if len(paramss) == 0 {
		paramss = []*StopStreamingParams{{}}
	}
	params := paramss[0]
	data := &StopStreamingResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
