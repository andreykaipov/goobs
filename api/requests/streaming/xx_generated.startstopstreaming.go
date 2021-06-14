// This file has been automatically generated. Don't edit it.

package streaming

import requests "github.com/andreykaipov/goobs/api/requests"

/*
StartStopStreamingParams represents the params body for the "StartStopStreaming" request.
Toggle streaming on or off (depending on the current stream state).
Since 0.3.
*/
type StartStopStreamingParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "StartStopStreaming".
func (o *StartStopStreamingParams) GetSelfName() string {
	return "StartStopStreaming"
}

/*
StartStopStreamingResponse represents the response body for the "StartStopStreaming" request.
Toggle streaming on or off (depending on the current stream state).
Since v0.3.
*/
type StartStopStreamingResponse struct {
	requests.ResponseBasic
}

// StartStopStreaming sends the corresponding request to the connected OBS WebSockets server. Note the variadic
// arguments as this request doesn't require any parameters.
func (c *Client) StartStopStreaming(paramss ...*StartStopStreamingParams) (*StartStopStreamingResponse, error) {
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
