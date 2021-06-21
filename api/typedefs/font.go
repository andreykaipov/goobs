package typedefs

// Font specifies the font for Text-specific requests.
type Font struct {
	// Font face.
	Face string `json:"face,omitempty"`

	// Font text styling flag. `Bold=1, Italic=2, Bold Italic=3,
	// Underline=5, Strikeout=8`
	Flags int `json:"flags,omitempty"`

	// Font text size.
	Size int `json:"size,omitempty"`

	// Font Style (unknown function).
	Style string `json:"style,omitempty"`
}
