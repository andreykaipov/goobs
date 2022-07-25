// This file has been automatically generated. Don't edit it.

package filters

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetSourceFilterEnabledParams represents the params body for the "SetSourceFilterEnabled" request.
Sets the enable state of a source filter.
*/
type SetSourceFilterEnabledParams struct {
	requests.ParamsBasic

	// New enable state of the filter
	FilterEnabled bool `json:"filterEnabled,omitempty"`

	// Name of the filter
	FilterName string `json:"filterName,omitempty"`

	// Name of the source the filter is on
	SourceName string `json:"sourceName,omitempty"`
}

// GetSelfName just returns "SetSourceFilterEnabled".
func (o *SetSourceFilterEnabledParams) GetSelfName() string {
	return "SetSourceFilterEnabled"
}

/*
SetSourceFilterEnabledResponse represents the response body for the "SetSourceFilterEnabled" request.
Sets the enable state of a source filter.
*/
type SetSourceFilterEnabledResponse struct {
	requests.ResponseBasic
}

// SetSourceFilterEnabled sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetSourceFilterEnabled(params *SetSourceFilterEnabledParams) (*SetSourceFilterEnabledResponse, error) {
	data := &SetSourceFilterEnabledResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
