// This file has been automatically generated. Don't edit it.

package typedefs

// OBSStats represents the complex type for OBSStats.
type OBSStats struct {
	// Average frame render time (in milliseconds)
	AverageFrameTime float64 `json:"average-frame-time,omitempty"`

	// Current CPU usage (percentage)
	CpuUsage float64 `json:"cpu-usage,omitempty"`

	// Current framerate.
	Fps float64 `json:"fps,omitempty"`

	// Free recording disk space (in megabytes)
	FreeDiskSpace float64 `json:"free-disk-space,omitempty"`

	// Current RAM usage (in megabytes)
	MemoryUsage float64 `json:"memory-usage,omitempty"`

	// Number of frames skipped due to encoding lag
	OutputSkippedFrames int `json:"output-skipped-frames,omitempty"`

	// Number of frames outputted
	OutputTotalFrames int `json:"output-total-frames,omitempty"`

	// Number of frames missed due to rendering lag
	RenderMissedFrames int `json:"render-missed-frames,omitempty"`

	// Number of frames rendered
	RenderTotalFrames int `json:"render-total-frames,omitempty"`
}
