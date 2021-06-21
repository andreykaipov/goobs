package typedefs

// Position defines a position specification.
type Position struct {
	// The point on the object that the item is manipulated from. The sum of
	// 1=Left or 2=Right, and 4=Top or 8=Bottom, or omit to center on that
	// axis.
	Alignment int `json:"alignment,omitempty"`

	// The x position of the obejct from the left.
	X float64 `json:"x,omitempty"`

	// The y position of the object from the top.
	Y float64 `json:"y,omitempty"`
}
