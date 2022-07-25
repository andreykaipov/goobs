// This file has been automatically generated. Don't edit it.

package filters

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetSourceFilterSettingsParams represents the params body for the "SetSourceFilterSettings" request.
Sets the settings of a source filter.
*/
type SetSourceFilterSettingsParams struct {
	requests.ParamsBasic

	// Name of the filter to set the settings of
	FilterName string `json:"filterName,omitempty"`

	// Object of settings to apply
	FilterSettings interface{} `json:"filterSettings,omitempty"`

	// True == apply the settings on top of existing ones, False == reset the input to its defaults, then apply
	// settings.
	Overlay bool `json:"overlay,omitempty"`

	// Name of the source the filter is on
	SourceName string `json:"sourceName,omitempty"`
}

// GetSelfName just returns "SetSourceFilterSettings".
func (o *SetSourceFilterSettingsParams) GetSelfName() string {
	return "SetSourceFilterSettings"
}

/*
SetSourceFilterSettingsResponse represents the response body for the "SetSourceFilterSettings" request.
Sets the settings of a source filter.
*/
type SetSourceFilterSettingsResponse struct {
	requests.ResponseBasic
}

// SetSourceFilterSettings sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetSourceFilterSettings(
	params *SetSourceFilterSettingsParams,
) (*SetSourceFilterSettingsResponse, error) {
	data := &SetSourceFilterSettingsResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
