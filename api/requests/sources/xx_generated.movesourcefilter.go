// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
MoveSourceFilterParams represents the params body for the "MoveSourceFilter" request.
Move a filter in the chain (relative positioning)
Since 4.5.0.
*/
type MoveSourceFilterParams struct {
	requests.ParamsBasic

	// Name of the filter to reorder
	FilterName string `json:"filterName,omitempty"`

	// How to move the filter around in the source's filter chain. Either "up", "down", "top" or "bottom".
	MovementType string `json:"movementType,omitempty"`

	// Name of the source to which the filter belongs
	SourceName string `json:"sourceName,omitempty"`
}

// GetSelfName just returns "MoveSourceFilter".
func (o *MoveSourceFilterParams) GetSelfName() string {
	return "MoveSourceFilter"
}

/*
MoveSourceFilterResponse represents the response body for the "MoveSourceFilter" request.
Move a filter in the chain (relative positioning)
Since v4.5.0.
*/
type MoveSourceFilterResponse struct {
	requests.ResponseBasic
}

// MoveSourceFilter sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) MoveSourceFilter(params *MoveSourceFilterParams) (*MoveSourceFilterResponse, error) {
	data := &MoveSourceFilterResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
