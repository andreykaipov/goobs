// This file has been automatically generated. Don't edit it.

package mediainputs

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetMediaInputCursorParams represents the params body for the "SetMediaInputCursor" request.
Sets the cursor position of a media input.

This request does not perform bounds checking of the cursor position.
*/
type SetMediaInputCursorParams struct {
	requests.ParamsBasic

	// Name of the media input
	InputName string `json:"inputName,omitempty"`

	// New cursor position to set
	MediaCursor float64 `json:"mediaCursor,omitempty"`
}

// GetSelfName just returns "SetMediaInputCursor".
func (o *SetMediaInputCursorParams) GetSelfName() string {
	return "SetMediaInputCursor"
}

/*
SetMediaInputCursorResponse represents the response body for the "SetMediaInputCursor" request.
Sets the cursor position of a media input.

This request does not perform bounds checking of the cursor position.
*/
type SetMediaInputCursorResponse struct {
	requests.ResponseBasic
}

// SetMediaInputCursor sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetMediaInputCursor(params *SetMediaInputCursorParams) (*SetMediaInputCursorResponse, error) {
	data := &SetMediaInputCursorResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
