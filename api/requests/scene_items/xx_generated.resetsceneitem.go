// This file has been automatically generated. Don't edit it.

package sceneitems

import (
	requests "github.com/andreykaipov/goobs/api/requests"
	typedefs "github.com/andreykaipov/goobs/api/typedefs"
)

/*
ResetSceneItemParams represents the params body for the "ResetSceneItem" request.
Reset a scene item.
Since 4.2.0.
*/
type ResetSceneItemParams struct {
	requests.ParamsBasic

	// The item specification for this object.
	Item *typedefs.Item `json:"item,omitempty"`

	// Name of the scene the scene item belongs to. Defaults to the current scene.
	SceneName string `json:"scene-name,omitempty"`
}

// GetSelfName just returns "ResetSceneItem".
func (o *ResetSceneItemParams) GetSelfName() string {
	return "ResetSceneItem"
}

/*
ResetSceneItemResponse represents the response body for the "ResetSceneItem" request.
Reset a scene item.
Since v4.2.0.
*/
type ResetSceneItemResponse struct {
	requests.ResponseBasic
}

// ResetSceneItem sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) ResetSceneItem(params *ResetSceneItemParams) (*ResetSceneItemResponse, error) {
	data := &ResetSceneItemResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
