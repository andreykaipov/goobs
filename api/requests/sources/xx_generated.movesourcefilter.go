// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
MoveSourceFilterParams represents the params body for the "MoveSourceFilter" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#MoveSourceFilter.
*/
type MoveSourceFilterParams struct {
	requests.ParamsBasic

	// Name of the filter to reorder
	FilterName string `json:"filterName"`

	// How to move the filter around in the source's filter chain. Either "up", "down", "top" or
	// "bottom".
	MovementType string `json:"movementType"`

	// Name of the source to which the filter belongs
	SourceName string `json:"sourceName"`
}

// Name just returns "MoveSourceFilter".
func (o *MoveSourceFilterParams) Name() string {
	return "MoveSourceFilter"
}

/*
MoveSourceFilterResponse represents the response body for the "MoveSourceFilter" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#MoveSourceFilter.
*/
type MoveSourceFilterResponse struct {
	requests.ResponseBasic
}

// MoveSourceFilter sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) MoveSourceFilter(
	params *MoveSourceFilterParams,
) (*MoveSourceFilterResponse, error) {
	data := &MoveSourceFilterResponse{}
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
