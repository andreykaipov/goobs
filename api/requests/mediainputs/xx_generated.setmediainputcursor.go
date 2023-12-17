// This file has been automatically generated. Don't edit it.

package mediainputs

// Represents the request body for the SetMediaInputCursor request.
type SetMediaInputCursorParams struct {
	// Name of the media input
	InputName string `json:"inputName,omitempty"`

	// New cursor position to set
	MediaCursor float64 `json:"mediaCursor,omitempty"`
}

func NewSetMediaInputCursorParams() *SetMediaInputCursorParams {
	return &SetMediaInputCursorParams{}
}
func (o *SetMediaInputCursorParams) WithInputName(x string) *SetMediaInputCursorParams {
	o.InputName = x
	return o
}
func (o *SetMediaInputCursorParams) WithMediaCursor(x float64) *SetMediaInputCursorParams {
	o.MediaCursor = x
	return o
}

// Returns the associated request.
func (o *SetMediaInputCursorParams) GetRequestName() string {
	return "SetMediaInputCursor"
}

// Represents the response body for the SetMediaInputCursor request.
type SetMediaInputCursorResponse struct{}

/*
Sets the cursor position of a media input.

This request does not perform bounds checking of the cursor position.
*/
func (c *Client) SetMediaInputCursor(params *SetMediaInputCursorParams) (*SetMediaInputCursorResponse, error) {
	data := &SetMediaInputCursorResponse{}
	return data, c.SendRequest(params, data)
}
