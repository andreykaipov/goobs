// This file has been automatically generated. Don't edit it.

package general

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetFilenameFormattingParams represents the params body for the "SetFilenameFormatting" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetFilenameFormatting.
*/
type SetFilenameFormattingParams struct {
	requests.Params

	// Filename formatting string to set.
	FilenameFormatting string `json:"filename-formatting"`
}

// GetRequestType returns the RequestType of SetFilenameFormattingParams
func (o *SetFilenameFormattingParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of SetFilenameFormattingParams
func (o *SetFilenameFormattingParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on SetFilenameFormattingParams
func (o *SetFilenameFormattingParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
SetFilenameFormattingResponse represents the response body for the "SetFilenameFormatting" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetFilenameFormatting.
*/
type SetFilenameFormattingResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of SetFilenameFormattingResponse
func (o *SetFilenameFormattingResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of SetFilenameFormattingResponse
func (o *SetFilenameFormattingResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of SetFilenameFormattingResponse
func (o *SetFilenameFormattingResponse) GetError() string {
	return o.Error
}

// SetFilenameFormatting sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetFilenameFormatting(
	params *SetFilenameFormattingParams,
) (*SetFilenameFormattingResponse, error) {
	params.RequestType = "SetFilenameFormatting"
	data := &SetFilenameFormattingResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
