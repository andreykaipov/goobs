// This file has been automatically generated. Don't edit it.

package stream

// Represents the request body for the StartStream request.
type StartStreamParams struct{}

// Returns the associated request.
func (o *StartStreamParams) GetRequestName() string {
	return "StartStream"
}

// Represents the response body for the StartStream request.
type StartStreamResponse struct{}

// Starts the stream output.
func (c *Client) StartStream(paramss ...*StartStreamParams) (*StartStreamResponse, error) {
	if len(paramss) == 0 {
		paramss = []*StartStreamParams{{}}
	}
	params := paramss[0]
	data := &StartStreamResponse{}
	return data, c.SendRequest(params, data)
}
