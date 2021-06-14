// This file has been automatically generated. Don't edit it.

package general

import requests "github.com/andreykaipov/goobs/api/requests"

/*
BroadcastCustomMessageParams represents the params body for the "BroadcastCustomMessage" request.
Broadcast custom message to all connected WebSocket clients

Generated from https://github.com/Palakis/obs-websocket/blob/4.7.0/docs/generated/protocol.md#BroadcastCustomMessage.
*/
type BroadcastCustomMessageParams struct {
	requests.ParamsBasic

	// User-defined data
	Data map[string]interface{} `json:"data"`

	// Identifier to be choosen by the client
	Realm string `json:"realm"`
}

// Name just returns "BroadcastCustomMessage".
func (o *BroadcastCustomMessageParams) Name() string {
	return "BroadcastCustomMessage"
}

/*
BroadcastCustomMessageResponse represents the response body for the "BroadcastCustomMessage" request.
Broadcast custom message to all connected WebSocket clients

Generated from https://github.com/Palakis/obs-websocket/blob/4.7.0/docs/generated/protocol.md#BroadcastCustomMessage.
*/
type BroadcastCustomMessageResponse struct {
	requests.ResponseBasic
}

// BroadcastCustomMessage sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) BroadcastCustomMessage(
	params *BroadcastCustomMessageParams,
) (*BroadcastCustomMessageResponse, error) {
	data := &BroadcastCustomMessageResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
