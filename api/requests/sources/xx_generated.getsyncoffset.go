// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetSyncOffsetParams represents the params body for the "GetSyncOffset" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSyncOffset.
*/
type GetSyncOffsetParams struct {
	requests.Params

	// Source name.
	Source string `json:"source"`
}

// GetRequestType returns the RequestType of GetSyncOffsetParams
func (o *GetSyncOffsetParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of GetSyncOffsetParams
func (o *GetSyncOffsetParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on GetSyncOffsetParams
func (o *GetSyncOffsetParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetSyncOffsetResponse represents the response body for the "GetSyncOffset" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSyncOffset.
*/
type GetSyncOffsetResponse struct {
	requests.Response

	// Source name.
	Name string `json:"name"`

	// The audio sync offset (in nanoseconds).
	Offset int `json:"offset"`
}

// GetMessageID returns the MessageID of GetSyncOffsetResponse
func (o *GetSyncOffsetResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of GetSyncOffsetResponse
func (o *GetSyncOffsetResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of GetSyncOffsetResponse
func (o *GetSyncOffsetResponse) GetError() string {
	return o.Error
}

// GetSyncOffset sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetSyncOffset(params *GetSyncOffsetParams) (*GetSyncOffsetResponse, error) {
	params.RequestType = "GetSyncOffset"
	data := &GetSyncOffsetResponse{}
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
