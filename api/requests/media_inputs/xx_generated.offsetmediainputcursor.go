// This file has been automatically generated. Don't edit it.

package mediainputs

/*
OffsetMediaInputCursorParams represents the params body for the "OffsetMediaInputCursor" request.
Offsets the current cursor position of a media input by the specified value.

This request does not perform bounds checking of the cursor position.
*/
type OffsetMediaInputCursorParams struct {
	// Name of the media input
	InputName string `json:"inputName,omitempty"`

	// Value to offset the current cursor position by
	MediaCursorOffset float64 `json:"mediaCursorOffset,omitempty"`
}

/*
OffsetMediaInputCursorResponse represents the response body for the "OffsetMediaInputCursor" request.
Offsets the current cursor position of a media input by the specified value.

This request does not perform bounds checking of the cursor position.
*/
type OffsetMediaInputCursorResponse struct{}

// OffsetMediaInputCursor sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) OffsetMediaInputCursor(params *OffsetMediaInputCursorParams) (*OffsetMediaInputCursorResponse, error) {
	resp, err := c.SendRequest("OffsetMediaInputCursor", params)
	if err != nil {
		return nil, err
	}
	return resp.(*OffsetMediaInputCursorResponse), nil
}
