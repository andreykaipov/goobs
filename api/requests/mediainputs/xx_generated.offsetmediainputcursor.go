// This file has been automatically generated. Don't edit it.

package mediainputs

import api "github.com/andreykaipov/goobs/api"

// Represents the request body for the OffsetMediaInputCursor request.
type OffsetMediaInputCursorParams struct {
	// Name of the media input
	InputName *string `json:"inputName,omitempty"`

	// Value to offset the current cursor position by
	MediaCursorOffset *float64 `json:"mediaCursorOffset,omitempty"`
}

func NewOffsetMediaInputCursorParams() *OffsetMediaInputCursorParams {
	return &OffsetMediaInputCursorParams{}
}
func (o *OffsetMediaInputCursorParams) WithInputName(x string) *OffsetMediaInputCursorParams {
	o.InputName = &x
	return o
}
func (o *OffsetMediaInputCursorParams) WithMediaCursorOffset(x float64) *OffsetMediaInputCursorParams {
	o.MediaCursorOffset = &x
	return o
}

// Returns the associated request.
func (o *OffsetMediaInputCursorParams) GetRequestName() string {
	return "OffsetMediaInputCursor"
}

// Represents the response body for the OffsetMediaInputCursor request.
type OffsetMediaInputCursorResponse struct {
	api.ResponseCommon
}

/*
Offsets the current cursor position of a media input by the specified value.

This request does not perform bounds checking of the cursor position.
*/
func (c *Client) OffsetMediaInputCursor(params *OffsetMediaInputCursorParams) (*OffsetMediaInputCursorResponse, error) {
	data := &OffsetMediaInputCursorResponse{}
	return data, c.SendRequest(params, data)
}
