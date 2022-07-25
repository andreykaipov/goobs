// This file has been automatically generated. Don't edit it.

package general

import requests "github.com/andreykaipov/goobs/api/requests"

/*
BroadcastCustomEventParams represents the params body for the "BroadcastCustomEvent" request.
Broadcasts a `CustomEvent` to all WebSocket clients. Receivers are clients which are identified and subscribed.
*/
type BroadcastCustomEventParams struct {
	requests.ParamsBasic

	// Data payload to emit to all receivers
	EventData interface{} `json:"eventData,omitempty"`
}

// GetSelfName just returns "BroadcastCustomEvent".
func (o *BroadcastCustomEventParams) GetSelfName() string {
	return "BroadcastCustomEvent"
}

/*
BroadcastCustomEventResponse represents the response body for the "BroadcastCustomEvent" request.
Broadcasts a `CustomEvent` to all WebSocket clients. Receivers are clients which are identified and subscribed.
*/
type BroadcastCustomEventResponse struct {
	requests.ResponseBasic
}

// BroadcastCustomEvent sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) BroadcastCustomEvent(params *BroadcastCustomEventParams) (*BroadcastCustomEventResponse, error) {
	data := &BroadcastCustomEventResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
