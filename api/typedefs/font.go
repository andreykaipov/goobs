package typedefs

// Font specifies the font for Text-specific requests.
type Font struct {
	// Font face.
	Face string `json:"face"`

	// Font text styling flag. `Bold=1, Italic=2, Bold Italic=3,
	// Underline=5, Strikeout=8`
	Flags int `json:"flags"`

	// Font text size.
	Size int `json:"size"`

	// Font Style (unknown function).
	Style string `json:"style"`
}
