// This file has been automatically generated. Don't edit it.

package typedefs

// Output represents the complex type for Output.
type Output struct {
	// Output status (active or not)
	Active bool `json:"active"`

	// Output congestion
	Congestion float64 `json:"congestion,omitempty"`

	// Number of frames dropped
	DroppedFrames int `json:"droppedFrames,omitempty"`

	Flags *Flags `json:"flags,omitempty"`

	// Video output height
	Height int `json:"height,omitempty"`

	// Output name
	Name string `json:"name,omitempty"`

	// Output reconnection status (reconnecting or not)
	Reconnecting bool `json:"reconnecting"`

	// Output settings
	Settings map[string]interface{} `json:"settings,omitempty"`

	// Total bytes sent
	TotalBytes int `json:"totalBytes,omitempty"`

	// Number of frames sent
	TotalFrames int `json:"totalFrames,omitempty"`

	// Output type/kind
	Type string `json:"type,omitempty"`

	// Video output width
	Width int `json:"width,omitempty"`
}

type Flags struct {
	// Output uses audio
	Audio bool `json:"audio"`

	// Output is encoded
	Encoded bool `json:"encoded"`

	// Output uses several audio tracks
	MultiTrack bool `json:"multiTrack"`

	// Raw flags value
	RawValue int `json:"rawValue,omitempty"`

	// Output uses a service
	Service bool `json:"service"`

	// Output uses video
	Video bool `json:"video"`
}
