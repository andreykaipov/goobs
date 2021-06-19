package typedefs

// Crop defines a crop specification.
type Crop struct {
	// The number of pixels cropped off the bottom of the object before
	// scaling.
	Bottom int `json:"bottom"`

	// The number of pixels cropped off the left of the object before
	// scaling.
	Left int `json:"left"`

	// The number of pixels cropped off the right of the object before
	// scaling.
	Right int `json:"right"`

	// The number of pixels cropped off the top of the object before
	// scaling.
	Top int `json:"top"`
}
