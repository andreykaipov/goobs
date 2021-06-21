package typedefs

// Bounds defines a bounding box.
type Bounds struct {
	// Alignment of the bounding box.
	Alignment int `json:"alignment,omitempty"`

	// Type of bounding box. Can be "OBS_BOUNDS_STRETCH",
	// "OBS_BOUNDS_SCALE_INNER", "OBS_BOUNDS_SCALE_OUTER",
	// "OBS_BOUNDS_SCALE_TO_WIDTH", "OBS_BOUNDS_SCALE_TO_HEIGHT",
	// "OBS_BOUNDS_MAX_ONLY" or "OBS_BOUNDS_NONE".
	Type string `json:"type,omitempty"`

	// Width of the bounding box.
	X float64 `json:"x,omitempty"`

	// Height of the bounding box.
	Y float64 `json:"y,omitempty"`
}
