// This file has been automatically generated. Don't edit it.

package stream

// Represents the request body for the StopStream request.
type StopStreamParams struct{}

// Returns the associated request.
func (o *StopStreamParams) GetRequestName() string {
	return "StopStream"
}

// Represents the response body for the StopStream request.
type StopStreamResponse struct{}

// Stops the stream output.
func (c *Client) StopStream(paramss ...*StopStreamParams) (*StopStreamResponse, error) {
	if len(paramss) == 0 {
		paramss = []*StopStreamParams{{}}
	}
	params := paramss[0]
	data := &StopStreamResponse{}
	return data, c.SendRequest(params, data)
}
