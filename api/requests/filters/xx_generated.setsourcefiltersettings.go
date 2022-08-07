// This file has been automatically generated. Don't edit it.

package filters

/*
SetSourceFilterSettingsParams represents the params body for the "SetSourceFilterSettings" request.
Sets the settings of a source filter.
*/
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

/*
SetSourceFilterSettingsResponse represents the response body for the "SetSourceFilterSettings" request.
Sets the settings of a source filter.
*/
type SetSourceFilterSettingsResponse struct{}

// SetSourceFilterSettings sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetSourceFilterSettings(
	params *SetSourceFilterSettingsParams,
) (*SetSourceFilterSettingsResponse, error) {
	resp, err := c.SendRequest("SetSourceFilterSettings", params)
	if err != nil {
		return nil, err
	}
	return resp.(*SetSourceFilterSettingsResponse), nil
}
