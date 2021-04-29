// This file has been automatically generated. Don't edit it.

package events

/*
Heartbeat represents the event body for the "Heartbeat" event.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#Heartbeat.
*/
type Heartbeat struct {
	EventBasic

	// Current active profile.
	CurrentProfile string `json:"current-profile"`

	// Current active scene.
	CurrentScene string `json:"current-scene"`

	// Toggles between every JSON message as an "I am alive" indicator.
	Pulse bool `json:"pulse"`

	// Current recording state.
	Recording bool `json:"recording"`

	// Current streaming state.
	Streaming bool `json:"streaming"`

	// Total bytes recorded since the recording started.
	TotalRecordBytes int `json:"total-record-bytes"`

	// Total frames recorded since the recording started.
	TotalRecordFrames int `json:"total-record-frames"`

	// Total time (in seconds) since recording started.
	TotalRecordTime int `json:"total-record-time"`

	// Total bytes sent since the stream started.
	TotalStreamBytes int `json:"total-stream-bytes"`

	// Total frames streamed since the stream started.
	TotalStreamFrames int `json:"total-stream-frames"`

	// Total time (in seconds) since the stream started.
	TotalStreamTime int `json:"total-stream-time"`
}
