// This file has been automatically generated. Don't edit it.

package stream

/*
StartStreamParams represents the params body for the "StartStream" request.
Starts the stream output.
*/
type StartStreamParams struct{}

/*
StartStreamResponse represents the response body for the "StartStream" request.
Starts the stream output.
*/
type StartStreamResponse struct{}

// StartStream sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments as
// this request doesn't require any parameters.
func (c *Client) StartStream(paramss ...*StartStreamParams) (*StartStreamResponse, error) {
	if len(paramss) == 0 {
		paramss = []*StartStreamParams{{}}
	}
	params := paramss[0]
	resp, err := c.SendRequest("StartStream", params)
	if err != nil {
		return nil, err
	}
	return resp.(*StartStreamResponse), nil
}
