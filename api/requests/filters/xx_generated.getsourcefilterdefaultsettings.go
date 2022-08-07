// This file has been automatically generated. Don't edit it.

package filters

/*
GetSourceFilterDefaultSettingsParams represents the params body for the "GetSourceFilterDefaultSettings" request.
Gets the default settings for a filter kind.
*/
type GetSourceFilterDefaultSettingsParams struct {
	// Filter kind to get the default settings for
	FilterKind string `json:"filterKind,omitempty"`
}

/*
GetSourceFilterDefaultSettingsResponse represents the response body for the "GetSourceFilterDefaultSettings" request.
Gets the default settings for a filter kind.
*/
type GetSourceFilterDefaultSettingsResponse struct {
	// Object of default settings for the filter kind
	DefaultFilterSettings map[string]interface{} `json:"defaultFilterSettings,omitempty"`
}

// GetSourceFilterDefaultSettings sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetSourceFilterDefaultSettings(
	params *GetSourceFilterDefaultSettingsParams,
) (*GetSourceFilterDefaultSettingsResponse, error) {
	resp, err := c.SendRequest("GetSourceFilterDefaultSettings", params)
	if err != nil {
		return nil, err
	}
	return resp.(*GetSourceFilterDefaultSettingsResponse), nil
}
