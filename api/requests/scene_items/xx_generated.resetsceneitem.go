// This file has been automatically generated. Don't edit it.

package sceneitems

import requests "github.com/andreykaipov/goobs/api/requests"

/*
ResetSceneItemParams represents the params body for the "ResetSceneItem" request.
Reset a scene item.
Since 4.2.0.
*/
type ResetSceneItemParams struct {
	requests.ParamsBasic

	Item struct {
		// Scene Item ID (if the `item` field is an object)
		Id int `json:"id"`

		// Scene Item name (if the `item` field is an object)
		Name string `json:"name"`
	} `json:"item"`

	// Name of the scene the scene item belongs to. Defaults to the current scene.
	SceneName string `json:"scene-name"`
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
