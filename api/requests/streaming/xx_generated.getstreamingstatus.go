// This file has been automatically generated. Don't edit it.

package streaming

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetStreamingStatusParams represents the params body for the "GetStreamingStatus" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetStreamingStatus.
*/
type GetStreamingStatusParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of GetStreamingStatusParams
func (o *GetStreamingStatusParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of GetStreamingStatusParams
func (o *GetStreamingStatusParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on GetStreamingStatusParams
func (o *GetStreamingStatusParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetStreamingStatusResponse represents the response body for the "GetStreamingStatus" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetStreamingStatus.
*/
type GetStreamingStatusResponse struct {
	requests.Response

	// Always false. Retrocompatibility with OBSRemote.
	PreviewOnly bool `json:"preview-only"`

	// Time elapsed since recording started (only present if currently recording).
	RecTimecode string `json:"rec-timecode"`

	// Current recording status.
	Recording bool `json:"recording"`

	// Time elapsed since streaming started (only present if currently streaming).
	StreamTimecode string `json:"stream-timecode"`

	// Current streaming status.
	Streaming bool `json:"streaming"`
}

// GetMessageID returns the MessageID of GetStreamingStatusResponse
func (o *GetStreamingStatusResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of GetStreamingStatusResponse
func (o *GetStreamingStatusResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of GetStreamingStatusResponse
func (o *GetStreamingStatusResponse) GetError() string {
	return o.Error
}

// GetStreamingStatus sends the corresponding request to the connected OBS WebSockets server. Note
// the variadic arguments as this request doesn't require any parameters.
func (c *Client) GetStreamingStatus(
	paramss ...*GetStreamingStatusParams,
) (*GetStreamingStatusResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetStreamingStatusParams{{}}
	}
	params := paramss[0]
	params.RequestType = "GetStreamingStatus"
	data := &GetStreamingStatusResponse{}
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
