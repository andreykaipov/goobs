// This file has been automatically generated. Don't edit it.

package events

/*
StreamStatus represents the event body for the "StreamStatus" event.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#StreamStatus.
*/
type StreamStatus struct {
	EventBasic

	// Amount of data per second (in bytes) transmitted by the stream encoder.
	BytesPerSec int `json:"bytes-per-sec"`

	// Current framerate.
	Fps float64 `json:"fps"`

	// Amount of data per second (in kilobits) transmitted by the stream encoder.
	KbitsPerSec int `json:"kbits-per-sec"`

	// Number of frames dropped by the encoder since the stream started.
	NumDroppedFrames int `json:"num-dropped-frames"`

	// Total number of frames transmitted since the stream started.
	NumTotalFrames int `json:"num-total-frames"`

	// Always false (retrocompatibility).
	PreviewOnly bool `json:"preview-only"`

	// Current recording state.
	Recording bool `json:"recording"`

	// Percentage of dropped frames.
	Strain float64 `json:"strain"`

	// Current streaming state.
	Streaming bool `json:"streaming"`

	// Total time (in seconds) since the stream started.
	TotalStreamTime int `json:"total-stream-time"`
}
