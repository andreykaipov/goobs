// This file has been automatically generated. Don't edit it.

package scenes

import (
	requests "github.com/andreykaipov/goobs/api/requests"
	typedefs "github.com/andreykaipov/goobs/api/typedefs"
)

/*
ReorderSceneItemsParams represents the params body for the "ReorderSceneItems" request.
Changes the order of scene items in the requested scene.
Since 4.5.0.
*/
type ReorderSceneItemsParams struct {
	requests.ParamsBasic

	// The items for this object.
	Items []*typedefs.Item `json:"items,omitempty"`

	// Name of the scene to reorder (defaults to current).
	Scene string `json:"scene,omitempty"`
}

// GetSelfName just returns "ReorderSceneItems".
func (o *ReorderSceneItemsParams) GetSelfName() string {
	return "ReorderSceneItems"
}

/*
ReorderSceneItemsResponse represents the response body for the "ReorderSceneItems" request.
Changes the order of scene items in the requested scene.
Since v4.5.0.
*/
type ReorderSceneItemsResponse struct {
	requests.ResponseBasic
}

// ReorderSceneItems sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) ReorderSceneItems(params *ReorderSceneItemsParams) (*ReorderSceneItemsResponse, error) {
	data := &ReorderSceneItemsResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
