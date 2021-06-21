package typedefs

// Scale defines a scale specification.
type Scale struct {
	// The scale filter of the source. Can be "OBS_SCALE_DISABLE",
	// "OBS_SCALE_POINT", "OBS_SCALE_BICUBIC", "OBS_SCALE_BILINEAR",
	// "OBS_SCALE_LANCZOS" or "OBS_SCALE_AREA".
	Filter string `json:"filter,omitempty"`

	// The x-scale factor of the object.
	X float64 `json:"x,omitempty"`

	// The y-scale factor of the object.
	Y float64 `json:"y,omitempty"`
}
