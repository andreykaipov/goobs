// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetMuteParams represents the params body for the "SetMute" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetMute.
*/
type SetMuteParams struct {
	requests.Params

	// Desired mute status.
	Mute bool `json:"mute"`

	// Source name.
	Source string `json:"source"`
}

// GetRequestType returns the RequestType of SetMuteParams
func (o *SetMuteParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of SetMuteParams
func (o *SetMuteParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on SetMuteParams
func (o *SetMuteParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
SetMuteResponse represents the response body for the "SetMute" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetMute.
*/
type SetMuteResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of SetMuteResponse
func (o *SetMuteResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of SetMuteResponse
func (o *SetMuteResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of SetMuteResponse
func (o *SetMuteResponse) GetError() string {
	return o.Error
}

// SetMute sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetMute(params *SetMuteParams) (*SetMuteResponse, error) {
	params.RequestType = "SetMute"
	data := &SetMuteResponse{}
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
