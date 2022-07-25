// This file has been automatically generated. Don't edit it.

package mediainputs

import requests "github.com/andreykaipov/goobs/api/requests"

/*
OffsetMediaInputCursorParams represents the params body for the "OffsetMediaInputCursor" request.
Offsets the current cursor position of a media input by the specified value.

This request does not perform bounds checking of the cursor position.
*/
type OffsetMediaInputCursorParams struct {
	requests.ParamsBasic

	// Name of the media input
	InputName string `json:"inputName,omitempty"`

	// Value to offset the current cursor position by
	MediaCursorOffset float64 `json:"mediaCursorOffset,omitempty"`
}

// GetSelfName just returns "OffsetMediaInputCursor".
func (o *OffsetMediaInputCursorParams) GetSelfName() string {
	return "OffsetMediaInputCursor"
}

/*
OffsetMediaInputCursorResponse represents the response body for the "OffsetMediaInputCursor" request.
Offsets the current cursor position of a media input by the specified value.

This request does not perform bounds checking of the cursor position.
*/
type OffsetMediaInputCursorResponse struct {
	requests.ResponseBasic
}

// OffsetMediaInputCursor sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) OffsetMediaInputCursor(params *OffsetMediaInputCursorParams) (*OffsetMediaInputCursorResponse, error) {
	data := &OffsetMediaInputCursorResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
