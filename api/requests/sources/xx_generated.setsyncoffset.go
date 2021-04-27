// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetSyncOffsetParams represents the params body for the "SetSyncOffset" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetSyncOffset.
*/
type SetSyncOffsetParams struct {
	requests.Params

	// The desired audio sync offset (in nanoseconds).
	Offset int `json:"offset"`

	// Source name.
	Source string `json:"source"`
}

// GetRequestType returns the RequestType of SetSyncOffsetParams
func (o *SetSyncOffsetParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of SetSyncOffsetParams
func (o *SetSyncOffsetParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on SetSyncOffsetParams
func (o *SetSyncOffsetParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
SetSyncOffsetResponse represents the response body for the "SetSyncOffset" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetSyncOffset.
*/
type SetSyncOffsetResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of SetSyncOffsetResponse
func (o *SetSyncOffsetResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of SetSyncOffsetResponse
func (o *SetSyncOffsetResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of SetSyncOffsetResponse
func (o *SetSyncOffsetResponse) GetError() string {
	return o.Error
}

// SetSyncOffset sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetSyncOffset(params *SetSyncOffsetParams) (*SetSyncOffsetResponse, error) {
	params.RequestType = "SetSyncOffset"
	data := &SetSyncOffsetResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
