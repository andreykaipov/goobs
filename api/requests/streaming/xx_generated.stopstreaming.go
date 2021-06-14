// This file has been automatically generated. Don't edit it.

package streaming

import requests "github.com/andreykaipov/goobs/api/requests"

/*
StopStreamingParams represents the params body for the "StopStreaming" request.
Stop streaming.
Will return an `error` if streaming is not active.
Since 4.1.0.
*/
type StopStreamingParams struct {
	requests.ParamsBasic
}

// GetSelfName just returns "StopStreaming".
func (o *StopStreamingParams) GetSelfName() string {
	return "StopStreaming"
}

/*
StopStreamingResponse represents the response body for the "StopStreaming" request.
Stop streaming.
Will return an `error` if streaming is not active.
Since v4.1.0.
*/
type StopStreamingResponse struct {
	requests.ResponseBasic
}

// StopStreaming sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments as
// this request doesn't require any parameters.
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
