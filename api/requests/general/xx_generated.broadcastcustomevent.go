// This file has been automatically generated. Don't edit it.

package general

/*
BroadcastCustomEventParams represents the params body for the "BroadcastCustomEvent" request.
Broadcasts a `CustomEvent` to all WebSocket clients. Receivers are clients which are identified and subscribed.
*/
type BroadcastCustomEventParams struct {
	// Data payload to emit to all receivers
	EventData interface{} `json:"eventData,omitempty"`
}

/*
BroadcastCustomEventResponse represents the response body for the "BroadcastCustomEvent" request.
Broadcasts a `CustomEvent` to all WebSocket clients. Receivers are clients which are identified and subscribed.
*/
type BroadcastCustomEventResponse struct{}

// BroadcastCustomEvent sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) BroadcastCustomEvent(params *BroadcastCustomEventParams) (*BroadcastCustomEventResponse, error) {
	resp, err := c.SendRequest("BroadcastCustomEvent", params)
	if err != nil {
		return nil, err
	}
	return resp.(*BroadcastCustomEventResponse), nil
}
