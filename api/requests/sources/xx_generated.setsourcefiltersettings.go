// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetSourceFilterSettingsParams represents the params body for the "SetSourceFilterSettings" request.
Update settings of a filter
Since 4.5.0.
*/
type SetSourceFilterSettingsParams struct {
	requests.ParamsBasic

	// Name of the filter to reconfigure
	FilterName string `json:"filterName,omitempty"`

	// New settings. These will be merged to the current filter settings.
	FilterSettings map[string]interface{} `json:"filterSettings,omitempty"`

	// Name of the source to which the filter belongs
	SourceName string `json:"sourceName,omitempty"`
}

// GetSelfName just returns "SetSourceFilterSettings".
func (o *SetSourceFilterSettingsParams) GetSelfName() string {
	return "SetSourceFilterSettings"
}

/*
SetSourceFilterSettingsResponse represents the response body for the "SetSourceFilterSettings" request.
Update settings of a filter
Since v4.5.0.
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
