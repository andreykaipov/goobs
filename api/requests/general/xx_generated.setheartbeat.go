// This file has been automatically generated. Don't edit it.

package general

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetHeartbeatParams represents the params body for the "SetHeartbeat" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetHeartbeat.
*/
type SetHeartbeatParams struct {
	requests.Params

	// Starts/Stops emitting heartbeat messages
	Enable bool `json:"enable"`
}

// GetRequestType returns the RequestType of SetHeartbeatParams
func (o *SetHeartbeatParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of SetHeartbeatParams
func (o *SetHeartbeatParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on SetHeartbeatParams
func (o *SetHeartbeatParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
SetHeartbeatResponse represents the response body for the "SetHeartbeat" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetHeartbeat.
*/
type SetHeartbeatResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of SetHeartbeatResponse
func (o *SetHeartbeatResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of SetHeartbeatResponse
func (o *SetHeartbeatResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of SetHeartbeatResponse
func (o *SetHeartbeatResponse) GetError() string {
	return o.Error
}

// SetHeartbeat sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetHeartbeat(params *SetHeartbeatParams) (*SetHeartbeatResponse, error) {
	params.RequestType = "SetHeartbeat"
	data := &SetHeartbeatResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
