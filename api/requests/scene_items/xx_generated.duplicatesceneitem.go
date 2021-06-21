// This file has been automatically generated. Don't edit it.

package sceneitems

import (
	requests "github.com/andreykaipov/goobs/api/requests"
	typedefs "github.com/andreykaipov/goobs/api/typedefs"
)

/*
DuplicateSceneItemParams represents the params body for the "DuplicateSceneItem" request.
Duplicates a scene item.
Since 4.5.0.
*/
type DuplicateSceneItemParams struct {
	requests.ParamsBasic

	// Name of the scene to copy the item from. Defaults to the current scene.
	FromScene string `json:"fromScene,omitempty"`

	// The item specification for this object.
	Item *typedefs.Item `json:"item,omitempty"`

	// Name of the scene to create the item in. Defaults to the current scene.
	ToScene string `json:"toScene,omitempty"`
}

// GetSelfName just returns "DuplicateSceneItem".
func (o *DuplicateSceneItemParams) GetSelfName() string {
	return "DuplicateSceneItem"
}

/*
DuplicateSceneItemResponse represents the response body for the "DuplicateSceneItem" request.
Duplicates a scene item.
Since v4.5.0.
*/
type DuplicateSceneItemResponse struct {
	requests.ResponseBasic

	// The item specification for this object.
	Item *typedefs.Item `json:"item,omitempty"`

	// Name of the scene where the new item was created
	Scene string `json:"scene,omitempty"`
}

// DuplicateSceneItem sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) DuplicateSceneItem(params *DuplicateSceneItemParams) (*DuplicateSceneItemResponse, error) {
	data := &DuplicateSceneItemResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
