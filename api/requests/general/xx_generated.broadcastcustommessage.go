// This file has been automatically generated. Don't edit it.

package general

import requests "github.com/andreykaipov/goobs/api/requests"

/*
BroadcastCustomMessageParams represents the params body for the "BroadcastCustomMessage" request.
Broadcast custom message to all connected WebSocket clients
Since 4.7.0.
*/
type BroadcastCustomMessageParams struct {
	requests.ParamsBasic

	// User-defined data
	Data map[string]interface{} `json:"data,omitempty"`

	// Identifier to be choosen by the client
	Realm string `json:"realm,omitempty"`
}

// GetSelfName just returns "BroadcastCustomMessage".
func (o *BroadcastCustomMessageParams) GetSelfName() string {
	return "BroadcastCustomMessage"
}

/*
BroadcastCustomMessageResponse represents the response body for the "BroadcastCustomMessage" request.
Broadcast custom message to all connected WebSocket clients
Since v4.7.0.
*/
type BroadcastCustomMessageResponse struct {
	requests.ResponseBasic
}

// BroadcastCustomMessage sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) BroadcastCustomMessage(params *BroadcastCustomMessageParams) (*BroadcastCustomMessageResponse, error) {
	data := &BroadcastCustomMessageResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
