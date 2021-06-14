// This file has been automatically generated. Don't edit it.

package typedefs

// OBSStats represents the complex type for OBSStats.
type OBSStats struct {
	// Average frame render time (in milliseconds)
	AverageFrameTime float64 `json:"average-frame-time"`

	// Current CPU usage (percentage)
	CpuUsage float64 `json:"cpu-usage"`

	// Current framerate.
	Fps float64 `json:"fps"`

	// Free recording disk space (in megabytes)
	FreeDiskSpace float64 `json:"free-disk-space"`

	// Current RAM usage (in megabytes)
	MemoryUsage float64 `json:"memory-usage"`

	// Number of frames skipped due to encoding lag
	OutputSkippedFrames int `json:"output-skipped-frames"`

	// Number of frames outputted
	OutputTotalFrames int `json:"output-total-frames"`

	// Number of frames missed due to rendering lag
	RenderMissedFrames int `json:"render-missed-frames"`

	// Number of frames rendered
	RenderTotalFrames int `json:"render-total-frames"`
}
