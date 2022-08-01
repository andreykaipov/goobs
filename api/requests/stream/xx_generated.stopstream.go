// This file has been automatically generated. Don't edit it.

package stream

/*
StopStreamParams represents the params body for the "StopStream" request.
Stops the stream output.
*/
type StopStreamParams struct{}

/*
StopStreamResponse represents the response body for the "StopStream" request.
Stops the stream output.
*/
type StopStreamResponse struct{}

// StopStream sends the corresponding request to the connected OBS WebSockets server. Note the variadic arguments as
// this request doesn't require any parameters.
func (c *Client) StopStream(paramss ...*StopStreamParams) (*StopStreamResponse, error) {
	if len(paramss) == 0 {
		paramss = []*StopStreamParams{{}}
	}
	params := paramss[0]
	resp, err := c.SendRequest("StopStream", params)
	if err != nil {
		return nil, err
	}
	return resp.(*StopStreamResponse), nil
}
