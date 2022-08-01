// This file has been automatically generated. Don't edit it.

package mediainputs

/*
SetMediaInputCursorParams represents the params body for the "SetMediaInputCursor" request.
Sets the cursor position of a media input.

This request does not perform bounds checking of the cursor position.
*/
type SetMediaInputCursorParams struct {
	// Name of the media input
	InputName string `json:"inputName,omitempty"`

	// New cursor position to set
	MediaCursor float64 `json:"mediaCursor,omitempty"`
}

/*
SetMediaInputCursorResponse represents the response body for the "SetMediaInputCursor" request.
Sets the cursor position of a media input.

This request does not perform bounds checking of the cursor position.
*/
type SetMediaInputCursorResponse struct{}

// SetMediaInputCursor sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetMediaInputCursor(params *SetMediaInputCursorParams) (*SetMediaInputCursorResponse, error) {
	resp, err := c.SendRequest("SetMediaInputCursor", params)
	if err != nil {
		return nil, err
	}
	return resp.(*SetMediaInputCursorResponse), nil
}
