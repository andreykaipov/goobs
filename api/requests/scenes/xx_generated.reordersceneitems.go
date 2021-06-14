// This file has been automatically generated. Don't edit it.

package scenes

import requests "github.com/andreykaipov/goobs/api/requests"

/*
ReorderSceneItemsParams represents the params body for the "ReorderSceneItems" request.
Changes the order of scene items in the requested scene.

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#ReorderSceneItems.
*/
type ReorderSceneItemsParams struct {
	requests.ParamsBasic

	Items []struct {
		// Id of a specific scene item. Unique on a scene by scene basis.
		Id int `json:"id"`

		// Name of a scene item. Sufficiently unique if no scene items share sources within the
		// scene.
		Name string `json:"name"`
	} `json:"items"`

	// Name of the scene to reorder (defaults to current).
	Scene string `json:"scene"`
}

// GetSelfName just returns "ReorderSceneItems".
func (o *ReorderSceneItemsParams) GetSelfName() string {
	return "ReorderSceneItems"
}

/*
ReorderSceneItemsResponse represents the response body for the "ReorderSceneItems" request.
Changes the order of scene items in the requested scene.

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#ReorderSceneItems.
*/
type ReorderSceneItemsResponse struct {
	requests.ResponseBasic
}

// ReorderSceneItems sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) ReorderSceneItems(
	params *ReorderSceneItemsParams,
) (*ReorderSceneItemsResponse, error) {
	data := &ReorderSceneItemsResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
