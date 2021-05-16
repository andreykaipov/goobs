// This file has been automatically generated. Don't edit it.

package streaming

import requests "github.com/andreykaipov/goobs/api/requests"

/*
StartStopStreamingParams represents the params body for the "StartStopStreaming" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StartStopStreaming.
*/
type StartStopStreamingParams struct {
	requests.ParamsBasic
}

// Name just returns "StartStopStreaming".
func (o *StartStopStreamingParams) Name() string {
	return "StartStopStreaming"
}

/*
StartStopStreamingResponse represents the response body for the "StartStopStreaming" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StartStopStreaming.
*/
type StartStopStreamingResponse struct {
	requests.ResponseBasic
}

// StartStopStreaming sends the corresponding request to the connected OBS WebSockets server. Note
// the variadic arguments as this request doesn't require any parameters.
func (c *Client) StartStopStreaming(
	paramss ...*StartStopStreamingParams,
) (*StartStopStreamingResponse, error) {
	if len(paramss) == 0 {
		paramss = []*StartStopStreamingParams{{}}
	}
	params := paramss[0]
	data := &StartStopStreamingResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
