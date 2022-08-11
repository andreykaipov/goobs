// This file has been automatically generated. Don't edit it.

package mediainputs

// Represents the request body for the OffsetMediaInputCursor request.
type OffsetMediaInputCursorParams struct {
	// Name of the media input
	InputName string `json:"inputName,omitempty"`

	// Value to offset the current cursor position by
	MediaCursorOffset float64 `json:"mediaCursorOffset,omitempty"`
}

// Returns the associated request.
func (o *OffsetMediaInputCursorParams) GetRequestName() string {
	return "OffsetMediaInputCursor"
}

// Represents the response body for the OffsetMediaInputCursor request.
type OffsetMediaInputCursorResponse struct{}

/*
Offsets the current cursor position of a media input by the specified value.

This request does not perform bounds checking of the cursor position.
*/
func (c *Client) OffsetMediaInputCursor(params *OffsetMediaInputCursorParams) (*OffsetMediaInputCursorResponse, error) {
	data := &OffsetMediaInputCursorResponse{}
	return data, c.SendRequest(params, data)
}
