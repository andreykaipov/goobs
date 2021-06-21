// This file has been automatically generated. Don't edit it.

package events

import typedefs "github.com/andreykaipov/goobs/api/typedefs"

/*
Heartbeat represents the event body for the "Heartbeat" event.
Since vv0.3.
*/
type Heartbeat struct {
	EventBasic

	// Current active profile.
	CurrentProfile string `json:"current-profile,omitempty"`

	// Current active scene.
	CurrentScene string `json:"current-scene,omitempty"`

	// Toggles between every JSON message as an "I am alive" indicator.
	Pulse bool `json:"pulse"`

	// Current recording state.
	Recording bool `json:"recording"`

	// OBS Stats
	Stats typedefs.OBSStats `json:"stats,omitempty"`

	// Current streaming state.
	Streaming bool `json:"streaming"`

	// Total bytes recorded since the recording started.
	TotalRecordBytes int `json:"total-record-bytes,omitempty"`

	// Total frames recorded since the recording started.
	TotalRecordFrames int `json:"total-record-frames,omitempty"`

	// Total time (in seconds) since recording started.
	TotalRecordTime int `json:"total-record-time,omitempty"`

	// Total bytes sent since the stream started.
	TotalStreamBytes int `json:"total-stream-bytes,omitempty"`

	// Total frames streamed since the stream started.
	TotalStreamFrames int `json:"total-stream-frames,omitempty"`

	// Total time (in seconds) since the stream started.
	TotalStreamTime int `json:"total-stream-time,omitempty"`
}
