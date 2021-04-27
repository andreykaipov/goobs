// This file has been automatically generated. Don't edit it.

package general

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetFilenameFormattingParams represents the params body for the "GetFilenameFormatting" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetFilenameFormatting.
*/
type GetFilenameFormattingParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of GetFilenameFormattingParams
func (o *GetFilenameFormattingParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of GetFilenameFormattingParams
func (o *GetFilenameFormattingParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on GetFilenameFormattingParams
func (o *GetFilenameFormattingParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetFilenameFormattingResponse represents the response body for the "GetFilenameFormatting" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetFilenameFormatting.
*/
type GetFilenameFormattingResponse struct {
	requests.Response

	// Current filename formatting string.
	FilenameFormatting string `json:"filename-formatting"`
}

// GetMessageID returns the MessageID of GetFilenameFormattingResponse
func (o *GetFilenameFormattingResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of GetFilenameFormattingResponse
func (o *GetFilenameFormattingResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of GetFilenameFormattingResponse
func (o *GetFilenameFormattingResponse) GetError() string {
	return o.Error
}

// GetFilenameFormatting sends the corresponding request to the connected OBS WebSockets server.
// Note the variadic arguments as this request doesn't require any parameters.
func (c *Client) GetFilenameFormatting(
	paramss ...*GetFilenameFormattingParams,
) (*GetFilenameFormattingResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetFilenameFormattingParams{{}}
	}
	params := paramss[0]
	params.RequestType = "GetFilenameFormatting"
	data := &GetFilenameFormattingResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
