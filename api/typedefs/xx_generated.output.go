// This file has been automatically generated. Don't edit it.

package typedefs

// Output represents the complex type for Output.
type Output struct {
	// Output status (active or not)
	Active bool `json:"active"`

	// Output congestion
	Congestion float64 `json:"congestion"`

	// Number of frames dropped
	DroppedFrames int `json:"droppedFrames"`

	Flags struct {
		// Output uses audio
		Audio bool `json:"audio"`

		// Output is encoded
		Encoded bool `json:"encoded"`

		// Output uses several audio tracks
		MultiTrack bool `json:"multiTrack"`

		// Raw flags value
		RawValue int `json:"rawValue"`

		// Output uses a service
		Service bool `json:"service"`

		// Output uses video
		Video bool `json:"video"`
	} `json:"flags"`

	// Video output height
	Height int `json:"height"`

	// Output name
	Name string `json:"name"`

	// Output reconnection status (reconnecting or not)
	Reconnecting bool `json:"reconnecting"`

	// Output settings
	Settings map[string]interface{} `json:"settings"`

	// Total bytes sent
	TotalBytes int `json:"totalBytes"`

	// Number of frames sent
	TotalFrames int `json:"totalFrames"`

	// Output type/kind
	Type string `json:"type"`

	// Video output width
	Width int `json:"width"`
}
