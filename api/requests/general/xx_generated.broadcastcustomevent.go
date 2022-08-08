// This file has been automatically generated. Don't edit it.

package general

// Represents the request body for the BroadcastCustomEvent request.
type BroadcastCustomEventParams struct {
	// Data payload to emit to all receivers
	EventData interface{} `json:"eventData,omitempty"`
}

// Returns the associated request.
func (o *BroadcastCustomEventParams) GetRequestName() string {
	return "BroadcastCustomEvent"
}

// Represents the response body for the BroadcastCustomEvent request.
type BroadcastCustomEventResponse struct{}

// Broadcasts a `CustomEvent` to all WebSocket clients. Receivers are clients which are identified and subscribed.
func (c *Client) BroadcastCustomEvent(params *BroadcastCustomEventParams) (*BroadcastCustomEventResponse, error) {
	data := &BroadcastCustomEventResponse{}
	return data, c.SendRequest(params, data)
}
