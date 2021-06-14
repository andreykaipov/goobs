// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetSourceFilterVisibilityParams represents the params body for the "SetSourceFilterVisibility" request.
Change the visibility/enabled state of a filter

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#SetSourceFilterVisibility.
*/
type SetSourceFilterVisibilityParams struct {
	requests.ParamsBasic

	// New filter state
	FilterEnabled bool `json:"filterEnabled"`

	// Source filter name
	FilterName string `json:"filterName"`

	// Source name
	SourceName string `json:"sourceName"`
}

// GetSelfName just returns "SetSourceFilterVisibility".
func (o *SetSourceFilterVisibilityParams) GetSelfName() string {
	return "SetSourceFilterVisibility"
}

/*
SetSourceFilterVisibilityResponse represents the response body for the "SetSourceFilterVisibility" request.
Change the visibility/enabled state of a filter

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#SetSourceFilterVisibility.
*/
type SetSourceFilterVisibilityResponse struct {
	requests.ResponseBasic
}

// SetSourceFilterVisibility sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetSourceFilterVisibility(
	params *SetSourceFilterVisibilityParams,
) (*SetSourceFilterVisibilityResponse, error) {
	data := &SetSourceFilterVisibilityResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
