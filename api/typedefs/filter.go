package typedefs

// Filter specifies a filter.
type Filter struct {
	// Filter status (enabled or not)
	Enabled bool `json:"enabled"`

	// Filter name
	Name string `json:"name,omitempty"`

	// Filter settings
	Settings map[string]interface{} `json:"settings,omitempty"`

	// Filter type
	Type string `json:"type,omitempty"`
}
