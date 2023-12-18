// This file has been automatically generated. Don't edit it.

package general

import api "github.com/andreykaipov/goobs/api"

// Represents the request body for the BroadcastCustomEvent request.
type BroadcastCustomEventParams struct {
	// Data payload to emit to all receivers
	EventData map[string]any `json:"eventData,omitempty"`
}

func NewBroadcastCustomEventParams() *BroadcastCustomEventParams {
	return &BroadcastCustomEventParams{}
}
func (o *BroadcastCustomEventParams) WithEventData(x map[string]any) *BroadcastCustomEventParams {
	o.EventData = x
	return o
}

// Returns the associated request.
func (o *BroadcastCustomEventParams) GetRequestName() string {
	return "BroadcastCustomEvent"
}

// Represents the response body for the BroadcastCustomEvent request.
type BroadcastCustomEventResponse struct {
	api.ResponseCommon
}

// Broadcasts a `CustomEvent` to all WebSocket clients. Receivers are clients which are identified and subscribed.
func (c *Client) BroadcastCustomEvent(params *BroadcastCustomEventParams) (*BroadcastCustomEventResponse, error) {
	data := &BroadcastCustomEventResponse{}
	return data, c.SendRequest(params, data)
}
