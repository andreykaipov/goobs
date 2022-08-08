// This file has been automatically generated. Don't edit it.

package filters

// Represents the request body for the GetSourceFilterDefaultSettings request.
type GetSourceFilterDefaultSettingsParams struct {
	// Filter kind to get the default settings for
	FilterKind string `json:"filterKind,omitempty"`
}

// Returns the associated request.
func (o *GetSourceFilterDefaultSettingsParams) GetRequestName() string {
	return "GetSourceFilterDefaultSettings"
}

// Represents the response body for the GetSourceFilterDefaultSettings request.
type GetSourceFilterDefaultSettingsResponse struct {
	// Object of default settings for the filter kind
	DefaultFilterSettings map[string]interface{} `json:"defaultFilterSettings,omitempty"`
}

// Gets the default settings for a filter kind.
func (c *Client) GetSourceFilterDefaultSettings(
	params *GetSourceFilterDefaultSettingsParams,
) (*GetSourceFilterDefaultSettingsResponse, error) {
	data := &GetSourceFilterDefaultSettingsResponse{}
	return data, c.SendRequest(params, data)
}
