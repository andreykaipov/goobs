// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetMuteParams represents the params body for the "GetMute" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetMute.
*/
type GetMuteParams struct {
	requests.Params

	// Source name.
	Source string `json:"source"`
}

// GetRequestType returns the RequestType of GetMuteParams
func (o *GetMuteParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of GetMuteParams
func (o *GetMuteParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on GetMuteParams
func (o *GetMuteParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetMuteResponse represents the response body for the "GetMute" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetMute.
*/
type GetMuteResponse struct {
	requests.Response

	// Mute status of the source.
	Muted bool `json:"muted"`

	// Source name.
	Name string `json:"name"`
}

// GetMessageID returns the MessageID of GetMuteResponse
func (o *GetMuteResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of GetMuteResponse
func (o *GetMuteResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of GetMuteResponse
func (o *GetMuteResponse) GetError() string {
	return o.Error
}

// GetMute sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetMute(params *GetMuteParams) (*GetMuteResponse, error) {
	params.RequestType = "GetMute"
	data := &GetMuteResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
