// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
ToggleMuteParams represents the params body for the "ToggleMute" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ToggleMute.
*/
type ToggleMuteParams struct {
	requests.Params

	// Source name.
	Source string `json:"source"`
}

// GetRequestType returns the RequestType of ToggleMuteParams
func (o *ToggleMuteParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of ToggleMuteParams
func (o *ToggleMuteParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on ToggleMuteParams
func (o *ToggleMuteParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
ToggleMuteResponse represents the response body for the "ToggleMute" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ToggleMute.
*/
type ToggleMuteResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of ToggleMuteResponse
func (o *ToggleMuteResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of ToggleMuteResponse
func (o *ToggleMuteResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of ToggleMuteResponse
func (o *ToggleMuteResponse) GetError() string {
	return o.Error
}

// ToggleMute sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) ToggleMute(params *ToggleMuteParams) (*ToggleMuteResponse, error) {
	params.RequestType = "ToggleMute"
	data := &ToggleMuteResponse{}
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
