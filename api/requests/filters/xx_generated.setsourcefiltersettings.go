// This file has been automatically generated. Don't edit it.

package filters

// Represents the request body for the SetSourceFilterSettings request.
type SetSourceFilterSettingsParams struct {
	// Name of the filter to set the settings of
	FilterName string `json:"filterName,omitempty"`

	// Object of settings to apply
	FilterSettings map[string]interface{} `json:"filterSettings,omitempty"`

	// True == apply the settings on top of existing ones, False == reset the input to its defaults, then apply
	// settings.
	Overlay *bool `json:"overlay,omitempty"`

	// Name of the source the filter is on
	SourceName string `json:"sourceName,omitempty"`
}

// Returns the associated request.
func (o *SetSourceFilterSettingsParams) GetRequestName() string {
	return "SetSourceFilterSettings"
}

// Represents the response body for the SetSourceFilterSettings request.
type SetSourceFilterSettingsResponse struct{}

// Sets the settings of a source filter.
func (c *Client) SetSourceFilterSettings(
	params *SetSourceFilterSettingsParams,
) (*SetSourceFilterSettingsResponse, error) {
	data := &SetSourceFilterSettingsResponse{}
	return data, c.SendRequest(params, data)
}
