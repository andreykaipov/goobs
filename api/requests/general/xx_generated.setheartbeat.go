// This file has been automatically generated. Don't edit it.

package general

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetHeartbeatParams represents the params body for the "SetHeartbeat" request.
Enable/disable sending of the Heartbeat event

Generated from https://github.com/Palakis/obs-websocket/blob/4.6.1/docs/generated/protocol.md#SetHeartbeat.
*/
type SetHeartbeatParams struct {
	requests.ParamsBasic

	// Starts/Stops emitting heartbeat messages
	Enable bool `json:"enable"`
}

// Name just returns "SetHeartbeat".
func (o *SetHeartbeatParams) Name() string {
	return "SetHeartbeat"
}

/*
SetHeartbeatResponse represents the response body for the "SetHeartbeat" request.
Enable/disable sending of the Heartbeat event

Generated from https://github.com/Palakis/obs-websocket/blob/4.6.1/docs/generated/protocol.md#SetHeartbeat.
*/
type SetHeartbeatResponse struct {
	requests.ResponseBasic
}

// SetHeartbeat sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetHeartbeat(params *SetHeartbeatParams) (*SetHeartbeatResponse, error) {
	data := &SetHeartbeatResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
