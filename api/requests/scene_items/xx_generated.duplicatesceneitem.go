// This file has been automatically generated. Don't edit it.

package sceneitems

import requests "github.com/andreykaipov/goobs/api/requests"

/*
DuplicateSceneItemParams represents the params body for the "DuplicateSceneItem" request.
Duplicates a scene item.

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#DuplicateSceneItem.
*/
type DuplicateSceneItemParams struct {
	requests.ParamsBasic

	// Name of the scene to copy the item from. Defaults to the current scene.
	FromScene string `json:"fromScene"`

	Item struct {
		// Scene Item ID.
		Id int `json:"id"`

		// Scene Item name (prefer `id`, including both is acceptable).
		Name string `json:"name"`
	} `json:"item"`

	// Name of the scene to create the item in. Defaults to the current scene.
	ToScene string `json:"toScene"`
}

// GetSelfName just returns "DuplicateSceneItem".
func (o *DuplicateSceneItemParams) GetSelfName() string {
	return "DuplicateSceneItem"
}

/*
DuplicateSceneItemResponse represents the response body for the "DuplicateSceneItem" request.
Duplicates a scene item.

Generated from https://github.com/Palakis/obs-websocket/blob/4.8.0/docs/generated/protocol.md#DuplicateSceneItem.
*/
type DuplicateSceneItemResponse struct {
	requests.ResponseBasic

	Item struct {
		// New item ID
		Id int `json:"id"`

		// New item name
		Name string `json:"name"`
	} `json:"item"`

	// Name of the scene where the new item was created
	Scene string `json:"scene"`
}

// DuplicateSceneItem sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) DuplicateSceneItem(
	params *DuplicateSceneItemParams,
) (*DuplicateSceneItemResponse, error) {
	data := &DuplicateSceneItemResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
